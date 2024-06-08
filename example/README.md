# Example

### Start local node
```
cd hardhat && npx hardhat node
```

### compile contract
```
npx hardhat compile
```

### Deploy contract and emit event

```
npx hardhat run ./scripts/deploy.js --network localhost
```

### Run Indxer
find Contract address and replace in `main.go`
```
go run ./indxer/main.go
```