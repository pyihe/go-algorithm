package tree

import (
	"log"
	"math/rand"
	"testing"
)

var (
	redBlackTree = &RedBlackTree{}
	numCnt       = 10
)

func TestRedBlackTree_AddNode(t *testing.T) {
	var toBeAddList = make([]Int, numCnt)
	for i := 0; i < numCnt; i++ {
		toBeAddList[i] = Int(rand.Intn(51))
	}
	//toBeAddList := []Int{42, 22, 49, 39, 8, 1, 19, 39, 1, 50}
	log.Printf("待增加的列表有: %v\n\n", toBeAddList)
	for _, v := range toBeAddList {
		redBlackTree.AddNode(v)
		log.Printf("添加节点[%v]后的树为: %v\n\n", v, redBlackTree)
	}
	log.Printf("添加后的树为: %v\n\n", redBlackTree)

	var index int
	for i := 0; i < numCnt; i++ {
		index = rand.Intn(numCnt)
		log.Printf("删除节点[%v]前的树为: %v\n\n", toBeAddList[index], redBlackTree)
		redBlackTree.RemoveNode(toBeAddList[index])
		log.Printf("删除节点[%v]后的树为: %v\n\n", toBeAddList[index], redBlackTree)
	}
}
