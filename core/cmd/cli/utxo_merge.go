// Copyright (c) 2019. Baidu Inc. All Rights Reserved.

package main

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/spf13/cobra"

	"github.com/xuperchain/xuperchain/core/global"
	"github.com/xuperchain/xuperchain/core/pb"
	"github.com/xuperchain/xuperchain/core/permission/acl"
	"github.com/xuperchain/xuperchain/core/utxo"
	"github.com/xuperchain/xuperchain/core/utxo/txhash"
)

// MergeUtxoCommand necessary parameter for merge utxo
type MergeUtxoCommand struct {
	cli *Cli
	cmd *cobra.Command
	// account will be merged
	account string
	// white merge an contract account, it can not be null
	accountPath string

	fee string //手续费
}

// NewMergeUtxoCommand new an instance of merge utxo command
func NewMergeUtxoCommand(cli *Cli) *cobra.Command {
	c := new(MergeUtxoCommand)
	c.cli = cli
	c.cmd = &cobra.Command{
		Use:   "merge ",
		Short: "merge the utxo of an account or address.",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.TODO()
			return c.mergeUtxo(ctx)
		},
	}
	c.addFlags()
	return c.cmd
}

func (c *MergeUtxoCommand) addFlags() {
	c.cmd.Flags().StringVarP(&c.account, "account", "A", "", "The account/address to be merged (default ./data/keys/address).")
	c.cmd.Flags().StringVarP(&c.accountPath, "accountPath", "P", "", "The account path, which is required for an account.")
	c.cmd.Flags().StringVar(&c.fee, "fee", "0", "fee of one tx")
}

func (c *MergeUtxoCommand) mergeUtxo(ctx context.Context) error {
	if acl.IsAccount(c.account) == 1 && c.accountPath == "" {
		return errors.New("accountPath can not be null because account is an Account name")
	}

	initAk, _ := readAddress(c.cli.RootOptions.Keys)
	if c.account == "" {
		c.account = initAk
	}

	tx := &pb.Transaction{
		Version:   utxo.TxVersion,
		Coinbase:  false,
		Nonce:     global.GenNonce(),
		Timestamp: time.Now().UnixNano(),
		Initiator: initAk,
	}

	ct := &CommTrans{
		FrozenHeight: 0,
		Version:      utxo.TxVersion,
		From:         c.account,
		Args:         make(map[string][]byte),
		IsQuick:      false,
		ChainName:    c.cli.RootOptions.Name,
		Keys:         c.cli.RootOptions.Keys,
		XchainClient: c.cli.XchainClient(),
		CryptoType:   c.cli.RootOptions.CryptoType,
	}

	txInputs, txOutput, err := ct.GenTxInputsWithMergeUTXO(context.Background())
	if err != nil {
		return err
	}

	tx.TxInputs = txInputs
	// validation check
	if len(tx.TxInputs) == 0 {
		return errors.New("not enough available utxo to merge")
	}

	txOutputs := []*pb.TxOutput{}

	//设置手续费的交易
	fee, ok := big.NewInt(0).SetString(c.fee, 10)
	if !ok {
		return errors.New("can't get fee")
	}
	txOutputs = append(txOutputs, &pb.TxOutput{
		ToAddr: []byte(utxo.FeePlaceholder),
		Amount: fee.Bytes(),
	})

	//转账的金额要减去手续费
	if txOutput != nil {
		amount := big.NewInt(0).SetBytes(txOutput.Amount)
		amount = big.NewInt(0).Sub(amount, fee)
		txOutput.Amount = amount.Bytes()
		txOutputs = append(txOutputs, txOutput)
	}

	tx.TxOutputs = txOutputs

	tx.AuthRequire, err = genAuthRequire(c.account, c.accountPath)
	if err != nil {
		return errors.New("genAuthRequire error")
	}

	// preExe
	preExeRPCReq := &pb.InvokeRPCRequest{
		Bcname: c.cli.RootOptions.Name,
		Requests: []*pb.InvokeRequest{
			{ModuleName: "transfer", Amount: c.fee},
		},
		Header:      global.GHeader(),
		Initiator:   initAk,
		AuthRequire: tx.AuthRequire,
	}
	preExeRes, err := ct.XchainClient.PreExec(context.Background(), preExeRPCReq)
	if err != nil {
		return err
	}
	tx.ContractRequests = preExeRes.GetResponse().GetRequests()
	tx.TxInputsExt = preExeRes.GetResponse().GetInputs()
	tx.TxOutputsExt = preExeRes.GetResponse().GetOutputs()

	tx.InitiatorSigns, err = ct.genInitSign(tx)
	if err != nil {
		return err
	}
	tx.AuthRequireSigns, err = ct.genAuthRequireSignsFromPath(tx, c.accountPath)
	if err != nil {
		return err
	}

	// calculate txid
	tx.Txid, err = txhash.MakeTransactionID(tx)
	if err != nil {
		return err
	}
	txid, err := ct.postTx(context.Background(), tx)
	if err != nil {
		return err
	}
	fmt.Println(txid)

	return nil
}
