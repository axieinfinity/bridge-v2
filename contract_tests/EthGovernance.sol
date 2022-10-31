// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.0 <0.9.0;

contract EthGovernance {
    struct Signature {
        uint8 v;
        bytes32 r;
        bytes32 s;
    }
    mapping(uint256 => mapping(address => Signature)) signatures;

    function relayBridgeOperators(
        uint256 _period,
        address[] calldata _operators,
        Signature[] memory _signatures
    ) public {
        for (uint i = 0; i < _operators.length; i++) {
            signatures[_period][_operators[i]] = _signatures[i];
        }
    }

    function getBridgeOperatorVotingSignatures(
        uint256 _period,
        address[] calldata _voters
    ) public view returns (Signature[] memory) {
        Signature[] memory _signatures = new Signature[](_voters.length);
        for (uint i = 0; i < _voters.length; i++) {
            _signatures[i] = signatures[_period][_voters[i]];
        }

        return _signatures;
    }
}
