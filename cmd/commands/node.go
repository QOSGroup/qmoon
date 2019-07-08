// Copyright 2018 The QOS Authors

package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/service/syncer"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/spf13/cobra"
)

// NodeCmd 数据库初始化命令
var NodeCmd = &cobra.Command{
	Use:   "node",
	Short: "node 管理",
}

var createNodeCmd = &cobra.Command{
	Use:   "create",
	Short: "添加node节点",
	RunE:  createNode,
}

var queryNodeCmd = &cobra.Command{
	Use:   "query",
	Short: "查询node节点",
	RunE:  queryNode,
}

var updateNodeCmd = &cobra.Command{
	Use:   "update nodeName",
	Short: "更新node节点",
	RunE:  updateNode,
}

var deleteNodeCmd = &cobra.Command{
	Use:   "delete nodeName",
	Short: "删除node节点",
	RunE:  deleteNode,
}

var syncNodeCmd = &cobra.Command{
	Use:   "sync nodeName",
	Short: "快速同步",
	RunE:  syncNode,
}

var (
	nodeName    string
	nodeUrl     string
	nodeType    string
	nodeVersion string
)

func init() {
	createNodeCmd.PersistentFlags().StringVar(&nodeName, "nodeName", "", "the name of node")
	createNodeCmd.PersistentFlags().StringVar(&nodeUrl, "nodeUrl", "", "the url of node")
	createNodeCmd.PersistentFlags().StringVar(&nodeType, "nodeType", "", fmt.Sprintf("节点类型:%s, %s, %s",
		types.NodeTypeQOS, types.NodeTypeQSC, types.NodeTypeCOSMOS))
	createNodeCmd.PersistentFlags().StringVar(&nodeVersion, "nodeVersion", "", "the version of node")
	createNodeCmd.Flags().String(types.FlagInfluxdbServer, "http://localhost:8086", "influxdb server")
	createNodeCmd.Flags().String(types.FlagInfluxdbUser, "", "influxdb user")
	createNodeCmd.Flags().String(types.FlagInfluxdbPassword, "", "influxdb password")

	updateNodeCmd.PersistentFlags().StringVar(&nodeName, "nodeName", "", "the name of node")
	updateNodeCmd.PersistentFlags().StringVar(&nodeUrl, "nodeUrl", "", "the url of node")

	queryNodeCmd.PersistentFlags().StringVar(&nodeName, "nodeName", "", "the name of node")

	deleteNodeCmd.Flags().String(types.FlagInfluxdbServer, "http://localhost:8086", "influxdb server")
	deleteNodeCmd.Flags().String(types.FlagInfluxdbUser, "", "influxdb user")
	deleteNodeCmd.Flags().String(types.FlagInfluxdbPassword, "", "influxdb password")

	registerFlagsDb(createNodeCmd)
	registerFlagsDb(queryNodeCmd)
	registerFlagsDb(updateNodeCmd)
	registerFlagsDb(deleteNodeCmd)
	registerFlagsDb(syncNodeCmd)

	NodeCmd.AddCommand(createNodeCmd, queryNodeCmd, deleteNodeCmd, syncNodeCmd)
}

func createNode(cmd *cobra.Command, args []string) error {
	if nodeName == "" {
		return errors.New("nodeName 不能为空")
	}

	if nodeUrl == "" {
		return errors.New("nodeUrl 不能为空")
	}

	if nodeType == "" {
		return errors.New("nodeType 不能为空")
	}

	if nodeVersion == "" {
		return errors.New("nodeVersion 不能为空")
	}

	if !types.CheckNodeType(types.NodeType(nodeType)) {
		return errors.New("nodeType 不支持")
	}

	if err := service.CreateNode(nodeName, nodeUrl, nodeType, nodeVersion, ""); err != nil {
		return err
	}

	return nil
}

func queryNode(cmd *cobra.Command, args []string) error {
	headers := []string{"name", "chain_id", "node_type", "url"}
	var datas [][]string
	if nodeName != "" {
		res, err := service.GetNodeByName(nodeName)
		if err != nil {
			return err
		}

		datas = append(datas, []string{res.Name, res.ChanID, res.NodeType, res.BaseURL})
		utils.PrintTable(cmd.OutOrStdout(), headers, datas)
	} else {
		res, err := service.AllNodes()
		if err != nil {
			return err
		}

		for _, v := range res {
			datas = append(datas, []string{v.Name, v.ChanID, v.NodeType, v.BaseURL})
		}

		utils.PrintTable(cmd.OutOrStdout(), headers, datas)
	}

	return nil
}

func updateNode(cmd *cobra.Command, args []string) error {

	return nil
}

func deleteNode(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("需要参数nodeName")
	}
	return service.DeleteNodeByName(args[0])
}

func syncNode(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("需要参数nodeName")
	}
	node, err := service.GetNodeByName(args[0])
	if err != nil {
		return err
	}
	syncer.NewSyncer(node).BlockLoop(context.Background())

	return nil
}
