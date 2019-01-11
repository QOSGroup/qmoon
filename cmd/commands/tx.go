// Copyright 2018 The QOS Authors

package commands

import (
	"github.com/spf13/cobra"
)

// TxCmd 交易解析
var TxCmd = &cobra.Command{
	Use:   "tx",
	Short: "解析块里的交易",
	RunE:  txParse,
}

var (
	height  int64
	txindex string
)

func init() {
	TxCmd.PersistentFlags().StringVar(&txindex, "txindex", "", "块中交易序号，逗号分割 1,2")
	TxCmd.PersistentFlags().StringVar(&nodeUrl, "remote", "127.0.0.1:26657", "node地址")
	TxCmd.PersistentFlags().Int64Var(&height, "height", 0, "the name of node")
}

func txParse(cmd *cobra.Command, args []string) error {
	//tmc := lib.TendermintClient(nodeUrl)
	//b, err := tmc.Block(&height)
	//if err != nil {
	//	return err
	//}
	//
	//if b.Block.NumTxs == 0 {
	//	fmt.Printf("没有交易\n")
	//	return nil
	//}
	//
	//cdc := lib.MakeCodec()
	//for k, v := range b.Block.Data.Txs {
	//	var tx cosmostypes.Msg
	//	err := cdc.UnmarshalBinaryBare(v, &tx)
	//	if err != nil {
	//		return err
	//	}
	//
	//	log.Printf("tx%d, %+v", k, tx)
	//
	//}

	return nil
}
