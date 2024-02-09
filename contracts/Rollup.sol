
// SPDX-License-Identifier: MIT

pragma solidity ^0.8.4;

import "./MerkleTree.sol";
import "./gnark_verifier.sol";
import "./MiMC.sol";

contract Rollup is MerkleTree {
    address public owner;

    // Mapping from Ethereum addresses to rollup addresses
    mapping(uint256 => address) public RollupToEth;

    //DepositQueue
    uint256[] public depositQueue;
    uint64 depositQueueIndex;

    // Reference to the ZK-Verifier contract
    Verifier verifier;

    struct Account {
        uint64 index;
        uint256 nonce;
        uint256 balance;
        PublicKey pubKey;
    }

    struct PublicKey {
        uint256 x;
        uint256 y;
    }

    mapping(uint256 => address) accounts;

    event Registered(
        address sender,
        uint256 index,
        uint256 nonce,
        uint256 balance,
        PublicKey pubkey
        
    );
    event Deposit(address indexed user, uint256 amount);
    event Withdraw(address indexed user, uint256 amount);
    event MerkleRootUpdated(uint256 newMerkleRoot);
    // Event for logging transaction data
    event TransactionData(uint256 indexed batchId, bytes indexed data);


    constructor(
        address _verifierAddress, 
        uint256 _treeLevels, 
        uint256 initRoot
    ) MerkleTree(_treeLevels) {
        levels = _treeLevels;
        verifier = Verifier(_verifierAddress);
        owner = msg.sender;
        setRoot(initRoot);
        depositQueueIndex = 0;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can submit batches");
        _;
    }

    // Function to submit a batch of transactions
    // The `external` modifier is used for functions that are expected to be called only from outside the contract
    // `calldata` is a special data location that contains the function arguments,
    // only available to external functions
    function submitVerifyBatch(
        uint256[8] calldata zkproof, 
        uint256[128] calldata input
        //uint256 batchId, 
        //bytes calldata transactions
    ) external onlyOwner {       
        uint256 preStateRoot = getRoot();
        //input[0] = preStateRoot;
        assert(preStateRoot == input[0]);
        //TODO: adjust for multiple transactions (BatchSize greather than 1)
        verifier.verifyProof(zkproof, input); // If the function hasn't reverted, zkProof is valid.
        
        // Emit an event containing the batchId and the transaction data
        // This data will be stored on the blockchain and can be accessed using blockchain explorers or Ethereum nodes
        //emit TransactionData(batchId, transactions);
        uint256 postStateRoot = input[input.length - 1];
        setRoot(postStateRoot); // New merkle root is part of the input
        emit MerkleRootUpdated(postStateRoot);
    }

    // Function to process the deposits in the depositQueue
    // by checking whether account is existent in the mapping, and if not,
    // adds entry to the mapping if leaf is empty
    function processDeposit(
        uint256[] memory merkleProof,
        uint256 amount,
        address ganacheAddress,
        Account memory account
    ) external onlyOwner {        
        require(verifyMerkleProof(merkleProof, account.index), "Merkle proof invalid");
        // set mapping from rollup account/index to Eth address 
        if (RollupToEth[account.pubKey.x] == address(0)) {
            // Check if leaf is empty 
            assert(merkleProof[0] == ZERO_VALUE);
            RollupToEth[account.pubKey.x] = ganacheAddress;
        }

        require(hashLeftRight(amount, uint256(uint160(ganacheAddress))) == depositQueue[depositQueueIndex], 
                "Deposit not correct one in queue");

        // this is problematic, as balance is already increased locally:
        // account.balance += amount;
        // Ensure that the balance update was successful and did not overflow
        assert(account.balance >= amount);

        // create account hash and recompute root from merkleproof
        uint256 newLeaf = hashAccount(account);
        merkleProof[0] = newLeaf;
        uint256 constructedRoot = computeMerkleRootFromPath(merkleProof, account.index);
    
        setRoot(constructedRoot); // New merkle root is part of the input
        emit MerkleRootUpdated(constructedRoot);
        
        // Processed deposit in queue and increase for next deposit
        depositQueueIndex++;
    }

    // Function to deposit funds to the depositQueue in the rollup and register
    function deposit() external payable {
        require(msg.value > 0, "Deposit value must be greater than 0");
        
        uint256 depositHash = hashLeftRight(msg.value, uint256(uint160(msg.sender)));
        depositQueue.push(depositHash);
        emit Deposit(msg.sender, msg.value);
    }

    // Function to withdraw funds from the rollup
    function withdraw(
        uint256 amount, 
        uint256[] memory merkleProof,
        Account memory account
    ) public {
        address ethAddress = RollupToEth[account.pubKey.x];
        require(ethAddress != address(0), "Rollup address does not exist");
        require(msg.sender == ethAddress, "Not rightful account address");
        require(account.balance >= amount, "Balance not sufficient");
        // construct leaf from account details here to verify Merkle Proof
        uint256 oldLeaf = hashAccount(account);
        merkleProof[0] = oldLeaf;
        require(verifyMerkleProof(merkleProof, account.index), "Invalid Merkle proof");

        payable(msg.sender).transfer(amount);
        emit Withdraw(ethAddress, amount);

        // Update the user's balance 
        account.balance -= amount;
        // This assert is used as a sanity check to ensure that the balance
        // subtraction did not underflow
        assert(account.balance <= account.balance + amount);

        uint256 newLeaf = hashAccount(account);
        merkleProof[0] = newLeaf;
        uint256 currentRoot = computeMerkleRootFromPath(merkleProof, account.index);
        setRoot(currentRoot);
        // TODO: update balance in rollup and check for identical root
        emit MerkleRootUpdated(currentRoot);
    }

    function getEthAddress(uint256 rollupAccountPKX) external view returns (address) {
        return RollupToEth[rollupAccountPKX];
    }

    function getMerkleRoot() external view returns (uint256) {
        return getRoot();
    }

    function hashAccount(Account memory account) public pure returns (uint256) {
        uint256[] memory input = new uint256[](5);

        input[0] = account.index;
        input[1] = account.nonce;
        input[2] = account.balance;
        input[3] = account.pubKey.x;
        input[4] = account.pubKey.y;
        
        return MiMC.hash(input);
    }
}
