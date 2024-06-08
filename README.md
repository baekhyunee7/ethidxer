# Ethidxer

ethidxer is a library for synchronizing event data on the blockchain (EVM)

### Example
If a contract written in solidty is deployed
```solidity
contract ExampleContract {
    event ExampleEvent(
        address indexed param1,
        int32 indexed param2,
        int256 indexed param3,
        address param4,
        uint256 param5
    );

    function triggerEvent(
        address _param1,
        int32 _param2,
        int256 _param3,
        address _param4,
        uint256 _param5
    ) public {
        emit ExampleEvent(_param1, _param2, _param3, _param4, _param5);
    }
}
```

You can easily pull events from the blockchain
```go
import (
	"context"
	"fmt"
	"math/big"
	"strings"

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
		Db(db). // gorm instance
		RpcCli(rpcCli). // ethclient.Client instance
		Addrs([]string{"0x5fbdb2315678afecb367f032d93f642f64180aa3"}). // deployed contract address
		Abi(&abi). // abi generate from abigen
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
```

* `ethidxer.AutoMigrate` will migrate a table `check_points`. It is used to record the height of the currently synchronized blockchain.
* `ethidxer.Register` is used to register events and event callbacks.
* `ethidxer.TrasactionHook` hook a function that will ensure transactionality during database transaction execution.
* `Run` loop to pull events on the blockchain and process them.
* `Build` a `indexer` instance will be created and configuration options can also be passed through.

### Option
* `WithReserveBlockNum` specifies how many latest blocks will not be processed immediately. To prevent events from being canceled due to blockchain forks.
* `WithLogger` specifies logger.
* `WithBatchSize` used to determine how many blocks to pull at one time.
* `WithSleepDuration` used to configure how long it will sleep if no new blocks are generated.

Please see specific [examples](https://github.com/baekhyunee7/ethidxer/tree/main/example)
