// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract RoninTrustedOrganization {
    struct TrustedOrganization {
        address bridgeVoter;
    }
    TrustedOrganization[] trustedOrganizations;

    function getAllTrustedOrganizations() external view returns (TrustedOrganization[] memory _trustedOrganizations) {
        return trustedOrganizations;
    }

    function setTrustedOrganizations(address[] memory trusted) external {
        for (uint i = 0; i < trusted.length; i++) {
            TrustedOrganization memory voter = TrustedOrganization(trusted[i]);
            trustedOrganizations.push(voter);
        }
    }
}
