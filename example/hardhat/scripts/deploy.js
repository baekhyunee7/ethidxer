async function main() {
    const hre = require("hardhat");
  
    // We get the contract to deploy
    const Example = await hre.ethers.getContractFactory("ExampleContract");
    const example = await Example.deploy();
  
    console.log("Example deployed to:", example.address);
  
    // Trigger the event by calling the function
    const tx = await example.triggerEvent(
      "0x0000000000000000000000000000000000000001", // param1
      12345, // param2
      "1234567890123456789012345678901234567890", // param3
      "0x0000000000000000000000000000000000000002", // param4
      "9876543210987654321" // param5
    );
  
    // Wait for the transaction to be mined
    await tx.wait();
  
    console.log("Event triggered");
  }
  
  main()
    .then(() => process.exit(0))
    .catch(error => {
      console.error(error);
      process.exit(1);
    });
  