const MiMC = artifacts.require("MiMC");
const MerkleTree = artifacts.require("MerkleTree");
const Verifier = artifacts.require("Verifier");
const Rollup = artifacts.require("Rollup");

const BN = require('bn.js');
/* const fs = require('fs');
const path = require('path');

// Function to save deployment info to a separate file
function saveDeploymentInfo(contractName, instance, suffix) {
  const deploymentsDir = path.join(__dirname, '../build/contracts');
  if (!fs.existsSync(deploymentsDir)){
    fs.mkdirSync(deploymentsDir);
  }
  const filePath = path.join(deploymentsDir, `${contractName}-${suffix}.json`);
  const info = {
    address: instance.address,
    transactionHash: instance.transactionHash,
  };
  fs.writeFileSync(filePath, JSON.stringify(info, null, 2));
} */

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
            return deployer.deploy(Verifier).then(function(verifierInstance) {
              console.log("Deployed Verifier at address:", verifierInstance.address);
              // Correctly call deploy and then use the result within the .then() promise handler
              return deployer.deploy(Rollup, verifierInstance.address, 4, hexNumber); //.then(function(rollupInstance) {
              // Pass the deployer object to the saveDeploymentInfo function  
                //console.log("Deployed Rollup at address:", rollupInstance.address);
                //saveDeploymentInfo('Rollup', rollupInstance, '1');
              //});
            });
          });
        });
      });
    });
  });
};
