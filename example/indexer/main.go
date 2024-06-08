package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/baekhyunee7/ethidxer"
	"github.com/baekhyunee7/ethidxer/example/indexer/contract"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Indexed struct {
	Param1 common.Address
	Param2 int32
	Param3 *big.Int
}

type Data struct {
	Param4 common.Address
	Param5 *big.Int
}

func main() {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	_ = ethidxer.AutoMigrate(db)
	rpcCli, _ := ethclient.Dial("http://localhost:8545")
	abi, _ := abi.JSON(strings.NewReader(string(contract.ContractABI)))
	indexer := ethidxer.Builder("example").
		Db(db).
		RpcCli(rpcCli).
		Addrs([]string{"0x5fbdb2315678afecb367f032d93f642f64180aa3"}).
		Abi(&abi).
		Build(ethidxer.WithSleepDuration(time.Minute))
	ethidxer.Register(indexer, "ExampleEvent", func(t1 *Indexed, t2 *Data) error {
		fmt.Printf("%s %d %s\n", t1.Param1, t1.Param2, t1.Param3.String())
		fmt.Printf("%s %s\n", t2.Param4, t2.Param5.String())
		return nil
	})
	err := indexer.Run(context.Background())
	if err != nil {
		panic(err)
	}
}
