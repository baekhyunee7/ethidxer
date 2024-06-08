package ethidxer

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type handler func(*types.Log) error

type indexer struct {
	db              *gorm.DB
	rpcCli          *ethclient.Client
	cfg             *config
	name            string
	handlers        map[string]handler
	addrs           []string
	abi             *abi.ABI
	transactionHook func(*gorm.DB) error
}

type CheckPoint struct {
	ID        int64  `gorm:"column:id; primaryKey"`
	Height    int64  `gorm:"column:height; not null; type:bigint; default:0"`
	Name      string `gorm:"column:name; not null; type:varchar(255); uniqueIndex"`
	CreatedAt int64  `gorm:"column:created_at; not null; type:bigint; autoCreateTime"`
	UpdatedAt int64  `gorm:"column:updated_at; not null; type:bigint; autoUpdateTime"`
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&CheckPoint{})
}

type config struct {
	reserveBlockNum uint64
	logger          Logger
	batchSize       int64
	sleepDuration   time.Duration
}

func newConfig() *config {
	return &config{
		reserveBlockNum: 0,
		logger:          NewDefaultLogger(),
		batchSize:       10,
		sleepDuration:   time.Minute,
	}
}

func (x *indexer) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			latestSyncBlock, err := x.getLatestSyncBlock(ctx, x.name)
			if err != nil {
				return err
			}

			endBlock, err := x.rpcCli.BlockNumber(ctx)
			if err != nil {
				return err
			}
			if endBlock >= x.cfg.reserveBlockNum {
				endBlock -= x.cfg.reserveBlockNum
			} else {
				endBlock = 0
			}
			if latestSyncBlock >= int64(endBlock) {
				x.cfg.logger.Infof("no new block created, latest sync block: %d", latestSyncBlock)
				sleepContext(ctx, time.Minute)
				continue
			}
			addrs := []common.Address{}
			for _, addr := range x.addrs {
				addrs = append(addrs, common.HexToAddress(addr))
			}
			eventHashs := []common.Hash{}
			eventHashMap := map[string]string{}
			for event := range x.handlers {
				h := x.abi.Events[event].ID
				eventHashs = append(eventHashs, h)
				eventHashMap[event] = h.Hex()
			}

			batchSize := x.cfg.batchSize
			for s := latestSyncBlock + 1; s <= int64(endBlock); s += batchSize {
				fromBlock := big.NewInt(s)
				toBlock := big.NewInt(s + batchSize - 1)

				if toBlock.Int64() > int64(endBlock) {
					toBlock = big.NewInt(int64(endBlock))
				}

				query := ethereum.FilterQuery{
					FromBlock: fromBlock,
					ToBlock:   toBlock,
					Addresses: addrs,
					Topics: [][]common.Hash{
						eventHashs,
					},
				}
				logs, err := x.rpcCli.FilterLogs(ctx, query)
				if err != nil {
					return err
				}
				for _, vLog := range logs {
					eventHashHex := vLog.Topics[0].Hex()
					for event, hashHex := range eventHashMap {
						if hashHex == eventHashHex {
							handler := x.handlers[event]
							err := handler(&vLog)
							if err != nil {
								return err
							}
							break
						}
					}
				}
				err = x.createLiquidityMinePairs(ctx, toBlock.Int64())
				if err != nil {
					return err
				}
				x.cfg.logger.Infof("sync to %d successed", toBlock.Int64())
			}
		}
	}
}

func sleepContext(ctx context.Context, d time.Duration) {
	select {
	case <-ctx.Done():
		return
	case <-time.After(d):
		return
	}
}

func Register[TIndex any, TData any](x *indexer, event string, f func(*TIndex, *TData) error) {
	x.handlers[event] = func(l *types.Log) error {
		var indexed TIndex
		err := unmarshalLogTopic(l, &indexed)
		if err != nil {
			return err
		}
		var data TData
		err = x.abi.UnpackIntoInterface(&data, event, l.Data)
		if err != nil {
			return err
		}
		return f(&indexed, &data)
	}
}

func (x *indexer) getLatestSyncBlock(ctx context.Context, name string) (int64, error) {
	var cp CheckPoint
	err := x.db.WithContext(ctx).Model(&CheckPoint{}).Where("name = ?", name).First(&cp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = x.db.WithContext(ctx).Create(&CheckPoint{
				Height: 0,
				Name:   name,
			}).Error
			if err != nil {
				return 0, err
			}
			return 0, nil
		} else {
			return 0, err
		}
	}
	return cp.Height, nil
}

func (x *indexer) createLiquidityMinePairs(ctx context.Context, endblock int64) error {
	return x.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if x.transactionHook != nil {
			err := x.transactionHook(tx)
			if err != nil {
				return err
			}
		}
		return tx.Model(&CheckPoint{}).Where("name = ?", x.name).
			Update("height", endblock).Error
	})
}

func (x *indexer) TrasactionHook(f func(*gorm.DB) error) {
	x.transactionHook = f
}
