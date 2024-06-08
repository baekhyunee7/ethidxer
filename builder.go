package ethidxer

import (
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type indexerbuilderForName struct {
	name string
}

func Builder(name string) *indexerbuilderForName {
	return &indexerbuilderForName{
		name: name,
	}
}

func (b *indexerbuilderForName) Db(db *gorm.DB) *indexerBuilderForDb {
	return &indexerBuilderForDb{
		name: b.name,
		db:   db,
	}
}

type indexerBuilderForDb struct {
	name string
	db   *gorm.DB
}

func (b *indexerBuilderForDb) RpcCli(cli *ethclient.Client) *indexerBuilderForRpcCli {
	return &indexerBuilderForRpcCli{
		name:   b.name,
		db:     b.db,
		rpcCli: cli,
	}
}

type indexerBuilderForRpcCli struct {
	name   string
	db     *gorm.DB
	rpcCli *ethclient.Client
}

func (b *indexerBuilderForRpcCli) Addrs(addrs []string) *indexerBuilderForAddr {
	return &indexerBuilderForAddr{
		name:   b.name,
		db:     b.db,
		rpcCli: b.rpcCli,
		addrs:  addrs,
	}
}

type indexerBuilderForAddr struct {
	name   string
	db     *gorm.DB
	rpcCli *ethclient.Client
	addrs  []string
}

func (b *indexerBuilderForAddr) Abi(abi *abi.ABI) *indexerBuilderForAbi {
	return &indexerBuilderForAbi{
		name:   b.name,
		db:     b.db,
		rpcCli: b.rpcCli,
		addrs:  b.addrs,
		abi:    abi,
	}
}

type indexerBuilderForAbi struct {
	name   string
	db     *gorm.DB
	rpcCli *ethclient.Client
	addrs  []string
	abi    *abi.ABI
}

func (b *indexerBuilderForAbi) Build(opts ...Option) *indexer {
	cfg := newConfig()
	for _, opt := range opts {
		opt(cfg)
	}
	return &indexer{
		db:       b.db,
		rpcCli:   b.rpcCli,
		cfg:      cfg,
		name:     b.name,
		addrs:    b.addrs,
		abi:      b.abi,
		handlers: map[string]handler{},
	}
}

type Option func(c *config)

func WithReserveBlockNum(num uint64) Option {
	return func(c *config) {
		c.reserveBlockNum = num
	}
}

func WithLogger(logger Logger) Option {
	return func(c *config) {
		c.logger = logger
	}
}

func WithBatchSize(bs int64) Option {
	return func(c *config) {
		c.batchSize = bs
	}
}

func WithSleepDuration(d time.Duration) Option {
	return func(c *config) {
		c.sleepDuration = d
	}
}
