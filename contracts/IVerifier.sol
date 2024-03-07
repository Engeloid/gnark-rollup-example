// SPDX-License-Identifier: MIT

pragma solidity ^0.8.9;

interface IVerifier {
    function verifyProof(
        uint256[8] calldata proof,
        uint256[8] calldata input
    ) external view;
}