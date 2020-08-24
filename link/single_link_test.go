package link

import (
	"log"
	"testing"
)

var (
	head   = NewHead()
	circle = NewCircleHead()
)

func TestHeadInfo_AddNode(t *testing.T) {
	log.Printf("追加前链表为: %v\n\n", head.ListLink())
	for i := 0; i < 10; i++ {
		head.AppendNode(i)
	}
	log.Printf("追加后的链表为: %v\n\n", head.ListLink())

	head.AddNode(0, -1)
	head.AddNode(9, 11)
	head.AddNode(7, 100)
	log.Printf("添加后的链表为: %v\n\n", head.ListLink())

	node := head.GetNode(4)
	log.Printf("获取到4节点为: %v\n\n", node.DataString())

	head.AddAfter(node, 20)
	log.Printf("添加后的链表为: %v\n\n", head.ListLink())

	head.DeleteNode(0)
	log.Printf("删除节点[0]后的链表为: %v\n\n", head.ListLink())
	head.DeleteNode(10)
	log.Printf("删除节点[10]后的链表为: %v\n\n", head.ListLink())
	log.Printf("最后节点数为: %v\n", head.GetNodeCount())
}

func TestCircleHead_AddHead(t *testing.T) {
	log.Printf("添加节点前的循环链表为: %v %v\n", circle.Len(), circle)
	circle.AddHead(1)
	log.Printf("AddHead后: %v %v\n", circle.Len(), circle)
	circle.AppendNode(2)
	log.Printf("AppendNode后: %v %v\n", circle.Len(), circle)
	circle.AddAt(1, 3)
	circle.AddAt(-1, 4)
	log.Printf("AddAt后: %v %v\n", circle.Len(), circle)
	log.Printf("获取[2]的节点为: %v\n", circle.GetNode(2).Data)
	circle.DeleteNode(3)
	log.Printf("DeleteNode后: %v %v\n", circle.Len(), circle)
}
