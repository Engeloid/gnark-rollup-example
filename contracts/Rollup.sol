// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Rollup {
    mapping(address => address) public ethToRollupAddress;
    mapping(address => uint256) public balances;
    bytes32 public merkleRoot;

    function deposit() external payable {
        balances[msg.sender] += msg.value;
    }

    function withdraw(uint256 amount) external {
        require(balances[msg.sender] >= amount, "Insufficient balance");
        require(address(this).balance >= amount, "Contract balance too low");

        balances[msg.sender] -= amount;
        payable(msg.sender).transfer(amount);
    }

    function registerRollupAddress(address rollupAddress) external {
        require(ethToRollupAddress[msg.sender] == address(0), "Address already registered");
        ethToRollupAddress[msg.sender] = rollupAddress;
    }

    function getRollupAddress(address ethAddress) external view returns (address) {
        return ethToRollupAddress[ethAddress];
    }

    function setMerkleRoot(bytes32 _root) external {
        merkleRoot = _root;
    }

    function getMerkleRoot() external view returns (bytes32) {
        return merkleRoot;
    }
}
