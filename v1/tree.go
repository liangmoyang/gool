package tree

import (
	"encoding/json"
	"fmt"
	"testing"
)

type ItemType struct {
	Id      int64
	Pid     int64
	Content string
}

type TreeItem struct {
	ItemType
	Children []*TreeItem
}

type IdMapTreeType map[int64]*TreeItem

func Tree1(t *testing.T) {
	var items = []ItemType{
		{Id: 1, Pid: 0, Content: "1"},
		{Id: 2, Pid: 1, Content: "1-2"},
		{Id: 3, Pid: 1, Content: "1-3"},
		{Id: 4, Pid: 3, Content: "1-3-4"},
		{Id: 5, Pid: 0, Content: "5"},
		{Id: 6, Pid: 5, Content: "5-6"},
		{Id: 7, Pid: 6, Content: "5-6-7"},
		{Id: 8, Pid: 6, Content: "5-6-8"},
	}
	var tree []*TreeItem
	idMapTreeItem := make(IdMapTreeType)

	for _, item := range items {
		var treeItem TreeItem
		treeItem.Id = item.Id
		treeItem.Pid = item.Pid
		treeItem.Content = item.Content

		if item.Pid == 0 {
			tree = append(tree, &treeItem)
		} else {
			// 子节点收集
			idMapTreeItem[item.Pid].Children = append(idMapTreeItem[item.Pid].Children, &treeItem)
		}

		idMapTreeItem[item.Id] = &treeItem
		jsonRes, _ := json.Marshal(tree)
		fmt.Println(string(jsonRes))
	}

	fmt.Println(idMapTreeItem)
}
