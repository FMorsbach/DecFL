pragma solidity >=0.4.0 <0.7.0;

contract Model {

    string public configurationAddress;
    string public weightsAddress;
    uint public epoch;

    uint private updatesTillAggregation;
    mapping(string => uint) private aggregationVotes;
    string private currentCandidate;
    uint private submittedVotes;
    string[] private votedAddresses;
    bool public aggregationReady;

    struct Submission {
        address trainer;
        string storageAddress;
    }

    Submission[] public localUpdates;

    constructor(string memory _configuration, string memory _weightsAddress, uint _updatesTillAggregation) public {
        configurationAddress = _configuration;
        weightsAddress = _weightsAddress;
        updatesTillAggregation = _updatesTillAggregation;
        epoch = 0;
        aggregationReady = false;
    }

    function submitLocalUpdate(string memory updateAddress) public {
        localUpdates.push(Submission(msg.sender, updateAddress));
        if (localUpdates.length >= updatesTillAggregation) {
            aggregationReady = true;
        }
    }

    function LocalUpdatesCount() public view returns (uint) {
        return localUpdates.length;
    }

    function submitLocalAggregation(string memory updateAddress) public {

        aggregationVotes[updateAddress] = aggregationVotes[updateAddress] + 1;
        submittedVotes = submittedVotes + 1;
        votedAddresses.push(updateAddress);

        if (aggregationVotes[updateAddress] > aggregationVotes[currentCandidate]) {
            currentCandidate = updateAddress;
        }

        if (submittedVotes >= updatesTillAggregation) {
            weightsAddress = currentCandidate;
            epoch = epoch + 1;
            submittedVotes = 0;
            aggregationReady = false;
            while(votedAddresses.length > 0) {
                aggregationVotes[votedAddresses[votedAddresses.length-1]] = 0;
                votedAddresses.pop();
            }
        }
    }
}