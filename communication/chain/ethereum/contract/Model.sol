pragma solidity >=0.4.0 <0.7.0;

contract Model {

    string public configurationAddress;
    string public weightsAddress;

    struct Submission {
        address trainer;
        string storageAddress;
    }

    Submission[] public localUpdates;
    Submission[] private localAggregations;

    constructor(string memory _configuration, string memory _weightsAddress) public {
        configurationAddress = _configuration;
        weightsAddress = _weightsAddress;
    }

    function submitLocalUpdate(string memory updateAddress) public {
        localUpdates.push(Submission(msg.sender, updateAddress));
    }

    function LocalUpdatesCount() public view returns (uint) {
        return localUpdates.length;
    }

    function submitLocalAggregation(string memory updateAddress) public {
        localAggregations.push(Submission(msg.sender, updateAddress));
    }
}