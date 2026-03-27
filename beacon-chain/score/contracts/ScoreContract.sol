// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

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
    mapping(bytes => ScoreData) public scores;
    mapping(uint256 => BlockAggregate) public blockAggregates;

    address public constant SYSTEM = address(0);

    event ValidatorRegistered(bytes pubKey);
    event ScoreUpdated(bytes pubKey, uint256 score, uint256 blockNumber);

    modifier onlySystem() {
        require(msg.sender == SYSTEM, "Only system can call");
        _;
    }

    function SetTargetValidatorsCount(uint256 _count) external onlySystem {
        targetValidatorsCount = _count;
    }

    function RegisterValidator(bytes calldata pubKey) external onlySystem {
        require(scores[pubKey].score == 0, "Validator already registered");

        scores[pubKey] = ScoreData({score: 1, lastUpdatedBlock: block.number});

        emit ValidatorRegistered(pubKey);
    }

    function UpdateScore(bytes calldata pubKey, uint256 score) external {
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

    function GetScore(bytes calldata pubKey) external view returns (uint256) {
        return scores[pubKey].score;
    }

    function GetScoreData(
        bytes calldata pubKey
    ) external view returns (uint256 score, uint256 lastUpdatedBlock) {
        ScoreData memory data = scores[pubKey];
        return (data.score, data.lastUpdatedBlock);
    }

    function IsValidatorRegistered(
        bytes calldata pubKey
    ) external view returns (bool) {
        return scores[pubKey].score != 0;
    }
}
