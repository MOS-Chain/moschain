package xchaincore

import (
	"fmt"
	"github.com/xuperchain/xuperchain/core/consensus/tdpos"
	"github.com/xuperchain/xuperchain/core/pb"
	"math/big"
	"strconv"
)

//生成区块的投票奖励
func (xc *XChainCore) GenerateVoteAward() ([]*pb.Transaction, error) {

	//奖励交易
	txs := make([]*pb.Transaction, 0)

	//选民的票数统计
	ballots := make(map[string]int64)
	tdpos.VoterBallots.Range(func(k, v interface{}) bool {
		voter, err := tdpos.ParseVoterKey(k.(string))
		if err != nil {
			return false
		}
		ballots[voter] = v.(int64)
		ballots["all"] += v.(int64)
		return true
	})

	//打印票数
	for voter, ballot := range ballots {
		if voter == "all" {
			xc.log.Info("[Vote_Award] all ballots count", "ballots", ballots["all"])
			continue
		}
		xc.log.Info("[Vote_Award] voter ballots count", "voter", voter, "ballot", ballot)
	}

	//生成奖励
	for voter, ballot := range ballots {
		if voter == "all" {
			continue
		}

		//投票占比
		r := new(big.Rat)
		r.SetString(fmt.Sprintf("%d/%d", ballot, ballots["all"]))
		ratio, err := strconv.ParseFloat(r.FloatString(16), 10)
		if err != nil {
			xc.log.Error("[Vote_Award] fail to ratio parse float64", "err", err)
			return nil, err
		}

		//投票奖励
		voteAward := xc.Ledger.GenesisBlock.CalcVoteAward(tdpos.VoteAward, ratio)
		ratioStr := fmt.Sprintf("%.16f", ratio)
		xc.log.Info("[Vote_Award] calc vote award success", "voter", voter, "ratio", ratioStr, "award", voteAward)

		//奖励为0的不生成交易
		if voteAward.Int64() == 0 {
			continue
		}

		//生成交易
		voteawardtx, err := xc.Utxovm.GenerateVoteAwardTx([]byte(voter), voteAward.String(), []byte{'1'})
		if err != nil {
			xc.log.Error("[Vote_Award] fail to generate vote award tx", "err", err)
			return nil, err
		}
		txs = append(txs, voteawardtx)
	}

	return txs, nil
}
