package tree

import (
	"log"
	"math/rand"
	"testing"
)

var (
	avlRoot = NewAVLTree(Int(10))
)

func TestAVLNode_AddNode(t *testing.T) {
	var toBeAddList = make([]Int, 10)
	for i := 0; i < 10; i++ {
		toBeAddList[i] = Int(rand.Intn(51))
	}
	//toBeAddList := []Int{21, 38, 6, 48, 22, 3, 49, 2, 27, 5}
	log.Printf("待增加的列表有: %v\n", toBeAddList)
	log.Printf("添加节点前的树为: %v\n", avlRoot)
	for _, v := range toBeAddList {
		avlRoot, _ = avlRoot.AddNode(v)
	}
	log.Printf("添加后的树为: %v\n", avlRoot)

	avlRoot, _ = avlRoot.RemoveNode(Int(25))
	log.Printf("删除节点后的树为: %v\n", avlRoot)
}
