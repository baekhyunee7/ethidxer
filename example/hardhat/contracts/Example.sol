// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

// Uncomment this line to use console.log
// import "hardhat/console.sol";

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