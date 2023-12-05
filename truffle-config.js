module.exports = {
  networks: {
    local_ganache: {
      network_id: "*",
      port: 7545,
      host: "127.0.0.1"
    },
    loc_development_development: {
      network_id: "*",
      port: 8545,
      host: "127.0.0.1"
    }
  },
// Set default mocha options here, use special reporters etc.
 mocha: {
  // timeout: 100000
},

// Configure your compilers
compilers: {
  solc: {
    // version: "0.8.0" // use this version to make it work
    version: "0.8.0", // Fetch exact version from solc-bin (default: truffle's version)
    docker: true,        // Use "0.5.1" you've installed locally with docker (default: false)
    // settings: {          // See the solidity docs for advice about optimization and evmVersion
    //  optimizer: {
    //    enabled: false,
    //    runs: 200
    //  },
    //  evmVersion: "byzantium"
    // }
  },
},
};
