package tree

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

type Int int

func (v Int) Value() interface{} {
	return v
}

func (v Int) Compare(target Element) int {
	return int(v - target.Value().(Int))
}

var (
	root = NewTree(Int(10))
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestBinaryTreeNode_AddNode(t *testing.T) {
	var toBeAddList = make([]Int, 50)
	for i := 0; i < 50; i++ {
		toBeAddList[i] = Int(rand.Intn(101))
	}
	log.Printf("待增加的列表有: %v\n", toBeAddList)
	log.Printf("添加节点前的树为: %v\n", root)
	for _, v := range toBeAddList {
		root.AddNode(v)
	}
	log.Printf("添加节点后的树为: %v\n", root)
}

func TestBinaryTreeNode_RemoveNode(t *testing.T) {
	//删除节点
	var toBeRemoveList []Int
	for i := 0; i < 10; i++ {
		toBeRemoveList = append(toBeRemoveList, Int(rand.Intn(101)))
	}
	log.Printf("待删除的列表有: %v\n", toBeRemoveList)
	for _, v := range toBeRemoveList {
		root.RemoveNode(v)
	}
	log.Printf("删除后的树为: %v\n", root)
}

func TestBinaryTreeNode_Search(t *testing.T) {
	target := Int(rand.Intn(101))
	result := root.Search(target)
	if result != nil {
		log.Printf("搜索[%d]的结果为: %v\n", target, result.Data)
	} else {
		log.Printf("搜索[%d]的结果为空\n", target)
	}
}

func TestBinaryTreeNode_Max(t *testing.T) {
	log.Printf("最大值为: %v\n", root.Max().Data)
	log.Printf("最小值为: %v\n", root.Min().Data)
}

func TestBinaryTreeNode_BFSTraverse(t *testing.T) {
	var toBeInsert []Int
	for i := 0; i < 10; i++ {
		toBeInsert = append(toBeInsert, Int(rand.Intn(11)))
	}
	log.Printf("待添加的列表为: %v\n", toBeInsert)
	log.Printf("添加节点前的根节点为: %v\n", root)
	for _, v := range toBeInsert {
		root.AddNode(v)
	}
	log.Printf("添加节点后的根节点为: %v\n", root)

	result := root.BFSTraverse()
	log.Printf("广度优先遍历结果为: %v\n", result)
}

func TestBinaryTreeNode_DFSTraverse(t *testing.T) {
	var toBeInsert []Int
	for i := 0; i < 20; i++ {
		toBeInsert = append(toBeInsert, Int(rand.Intn(21)))
	}
	log.Printf("待添加的列表为: %v\n", toBeInsert)
	log.Printf("添加节点前的根节点为: %v\n", root)
	for _, v := range toBeInsert {
		root.AddNode(v)
	}
	log.Printf("添加节点后的根节点为: %v\n", root)

	result := root.DFSTraverse()
	log.Printf("深度优先遍历结果为: %v\n", result)
}