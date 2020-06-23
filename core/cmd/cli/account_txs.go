/*
 * Copyright (c) 2019. Baidu Inc. All Rights Reserved.
 */

package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xuperchain/xuperchain/core/pb"
)

type AccountTxsQueryCommand struct {
	cli   *Cli
	cmd   *cobra.Command
	addr  string
	page  int64
	limit int64
}

func NewAccountTxsQueryCommand(cli *Cli) *cobra.Command {
	c := new(AccountTxsQueryCommand)
	c.cli = cli
	c.cmd = &cobra.Command{
		Use:   "list",
		Short: "Get tx records info of an user.",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.TODO()
			return c.queryAccountTxs(ctx)
		},
	}
	c.addFlags()
	return c.cmd
}

func (c *AccountTxsQueryCommand) addFlags() {
	c.cmd.Flags().StringVarP(&c.addr, "address", "A", "", "address")
	c.cmd.Flags().Int64VarP(&c.page, "page", "P", 1, "page number")
	c.cmd.Flags().Int64VarP(&c.limit, "limit", "L", 10, "txs limit")
}

func (c *AccountTxsQueryCommand) queryAccountTxs(ctx context.Context) error {
	client := c.cli.XchainClient()
	if c.addr == "" {
		c.addr, _ = readAddress(c.cli.RootOptions.Keys)
	}
	request := &pb.AccountTxs{
		Bcname:       c.cli.RootOptions.Name,
		AccountName:  c.addr,
		PageNum:      c.page,
		DisplayCount: c.limit,
	}
	reply, err := client.QueryAccountTxs(ctx, request)
	if err != nil {
		return err
	}

	txs := []*Transaction{}
	var output []byte
	for _, v := range reply.Txs {
		tx := FromSimpleTx(v)
		txs = append(txs, tx)
		output, err = json.MarshalIndent(txs, "", "  ")
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println(string(output))

	return nil
}
