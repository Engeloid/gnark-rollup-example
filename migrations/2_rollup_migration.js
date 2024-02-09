const MiMC = artifacts.require("MiMC");
const MerkleTree = artifacts.require("MerkleTree");
const Verifier = artifacts.require("Verifier");
const Rollup = artifacts.require("Rollup");

const BN = require('bn.js');

let number = new BN('19929288186675197851976196030644791160893678128488844394573589927998971210120');
let hexNumber = '0x' + number.toString('hex');

module.exports = function(deployer) {
  deployer.deploy(MiMC).then(function() {
    // Link MiMC to MerkleTree
    return deployer.link(MiMC, MerkleTree).then(function() {
      // Deploy MerkleTree
      return deployer.deploy(MerkleTree, 4).then(function() {
        // Link MiMC and MerkleTree to Rollup
        return deployer.link(MiMC, Rollup).then(function() {
          return deployer.link(MerkleTree, Rollup).then(function() {
            // Deploy Verifier
            return deployer.deploy(Verifier).then(function() {
              // Finally, deploy Rollup with the address of the Verifier
              return deployer.deploy(Rollup, Verifier.address, 4, hexNumber);
            });
          });
        });
      });
    });
  });
};
