// Copyright 2018 The QOS Authors

package commands

import (
	"encoding/json"
	"errors"

	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/service"
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

var (
	nodeName string
	nodeUrl  string
)

func init() {
	createNodeCmd.PersistentFlags().StringVar(&nodeName, "nodeName", "", "the name of node")
	createNodeCmd.PersistentFlags().StringVar(&nodeUrl, "nodeUrl", "", "the url of node")

	updateNodeCmd.PersistentFlags().StringVar(&nodeName, "nodeName", "", "the name of node")
	updateNodeCmd.PersistentFlags().StringVar(&nodeUrl, "nodeUrl", "", "the url of node")

	queryNodeCmd.PersistentFlags().StringVar(&nodeName, "nodeName", "", "the name of node")

	registerFlagsDb(createNodeCmd)
	registerFlagsDb(queryNodeCmd)
	registerFlagsDb(updateNodeCmd)
	registerFlagsDb(deleteNodeCmd)

	NodeCmd.AddCommand(createNodeCmd, queryNodeCmd, deleteNodeCmd)
}

func createNode(cmd *cobra.Command, args []string) error {
	err := db.InitDb(config.DB, logger)
	if err != nil {
		return err
	}

	if nodeName == "" {
		return errors.New("nodeName 不能为空")
	}

	if nodeUrl == "" {
		return errors.New("nodeUrl 不能为空")
	}

	err = service.CreateNode(nodeName, nodeUrl, "", nil)
	if err != nil {
		return err
	}

	return nil
}

func queryNode(cmd *cobra.Command, args []string) error {
	err := db.InitDb(config.DB, logger)
	if err != nil {
		return err
	}

	if nodeName != "" {
		res, err := service.GetNodeByName(nodeName)
		if err != nil {
			return err
		}

		d, err := json.MarshalIndent(res, "", "    ")
		if err != nil {
			return err
		}

		cmd.Println(string(d))
	} else {
		res, err := service.AllNodes()
		if err != nil {
			return err
		}

		d, err := json.MarshalIndent(res, "", "    ")
		if err != nil {
			return err
		}

		cmd.Println(string(d))
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
	name := args[0]

	err := db.InitDb(config.DB, logger)
	if err != nil {
		return err
	}

	err = service.DeleteNodeByName(name)

	return err
}
