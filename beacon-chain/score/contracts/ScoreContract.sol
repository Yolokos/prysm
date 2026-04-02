// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

contract ScoreContract {
    struct ScoreData {
        uint256 score;
        uint256 lastUpdatedBlock;
    }

    struct BlockAggregate {
        uint256 totalScore;
        uint256 count;
    }

    uint256 public targetValidatorsCount;
    uint256 public validatorsCount;
    mapping(bytes32 => ScoreData) public scores;
    mapping(uint256 => BlockAggregate) public blockAggregates;
    mapping(bytes32 => bool) public isRegistered;

    address public constant SYSTEM = address(0);

    event ValidatorRegistered(bytes32 pubKey);
    event ScoreUpdated(bytes32 pubKey, uint256 score, uint256 blockNumber);

    modifier onlySystem() {
        require(msg.sender == SYSTEM, "Only system can call");
        _;
    }

    function SetTargetValidatorsCount(uint256 _count) external onlySystem {
        targetValidatorsCount = _count;
    }

    function RegisterValidator(bytes32 pubKey) external onlySystem {
        require(!isRegistered[pubKey], "Already registered");

        isRegistered[pubKey] = true;

        scores[pubKey] = ScoreData({score: 400, lastUpdatedBlock: block.number});
    }

    function UpdateScore(bytes32 pubKey, uint256 score) external {
        ScoreData storage data = scores[pubKey];

        uint256 prevScore = data.score;
        uint256 prevBlock = data.lastUpdatedBlock;

        if (prevBlock != 0) {
            BlockAggregate storage prevAgg = blockAggregates[prevBlock];
            if (prevAgg.count > 0) {
                prevAgg.totalScore -= prevScore;
                prevAgg.count -= 1;
            }
        }

        data.score = score;
        data.lastUpdatedBlock = block.number;

        BlockAggregate storage agg = blockAggregates[block.number];
        agg.totalScore += score;
        agg.count += 1;

        emit ScoreUpdated(pubKey, score, block.number);
    }

    function GetEpochRangeScore(
        uint256 startBlock,
        uint256 endBlock
    ) external view returns (uint256) {
        uint256 totalScore = 0;
        uint256 totalCount = 0;

        for (uint256 b = startBlock; b <= endBlock; b++) {
            BlockAggregate memory agg = blockAggregates[b];
            totalScore += agg.totalScore;
            totalCount += agg.count;
        }

        if (totalCount == 0) {
            return 0;
        }

        return totalScore / totalCount;
    }

    function GetScore(bytes32 pubKey) external view returns (uint256) {
        require(isRegistered[pubKey], "Validator not registered");
        return scores[pubKey].score;
    }

    function GetScoreData(
        bytes32 pubKey
    ) external view returns (uint256 score, uint256 lastUpdatedBlock) {
        ScoreData memory data = scores[pubKey];
        return (data.score, data.lastUpdatedBlock);
    }

    function IsValidatorRegistered(
        bytes32 pubKey
    ) external view returns (bool) {
        return isRegistered[pubKey];
    }
}
