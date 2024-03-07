const MiMC = artifacts.require("MiMC");
const MerkleTree = artifacts.require("MerkleTree");
// Assuming you have Verifier2 contract similar to Verifier
const Verifier2 = artifacts.require("Verifier2");
const Rollup = artifacts.require("Rollup");

const BN = require('bn.js');

let number = new BN('19929288186675197851976196030644791160893678128488844394573589927998971210120');
let hexNumber = '0x' + number.toString('hex');

module.exports = function(deployer) {
    deployer.then(async () => {
        await deployer.deploy(Verifier2);
        const verifier2Instance = await Verifier2.deployed();

        // Assuming the deployment of MiMC and MerkleTree has already been done in the first migration file
        // and they do not need to be redeployed or relinked for the second Rollup deployment.
        // Deploy the second Rollup with Verifier2
        await deployer.deploy(Rollup, verifier2Instance.address, 4, hexNumber);
    });
};
