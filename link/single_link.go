package link

import (
	"fmt"
	"sync"
)

/*
	单链表(如果尾节点的next指向头节点则表示循环链表, 循环链表的操作与单链表大致相同，不同的地方在于判断尾节点条件为尾节点next指向头节点而不是nil)
	1. 每个节点只有两个数据字段: 数据域和指针域, 数据域表示该节点存储的数据, 指针域存储下一个节点的地址
	2. 追加节点
	3. 在指定位置添加节点
	4. 在指定节点后添加节点
	5. 获取指定位置的节点(根据data来获取节点原理一样)
	6. 删除指定位置的节点
	7. 获取节点数量
	8. 遍历节点
*/

//表节点
type SingleNode struct {
	Data interface{} //数据
	Next *SingleNode //指向下一个节点
}

//表头信息
type SingleHead struct {
	mu     sync.RWMutex
	Length int         //链表长度
	Head   *SingleNode //表头指针
}

func NewLinkNode(data interface{}) *SingleNode {
	return &SingleNode{
		Data: data,
		Next: nil,
	}
}

func NewHead() *SingleHead {
	return &SingleHead{
		Length: 0,
		Head:   nil,
	}
}

func (h *SingleHead) String() string {
	var result string
	node := h.Head
	for node != nil {
		result += fmt.Sprintf("%v ", node.Data)
	}
	return result
}

func (n *SingleNode) DataString() string {
	var result string
	node := n
	for node != nil {
		result += fmt.Sprintf("%+v->", *node)
		node = node.Next
	}
	result += "NULL"
	return result
}

//表尾追加节点
func (h *SingleHead) AppendNode(nodeData interface{}) {
	//尾节点next指向nil
	node := &SingleNode{
		Data: nodeData,
		Next: nil,
	}
	h.mu.Lock()
	defer h.mu.Unlock()

	//如果是空链表
	if h.Head == nil {
		h.Head = node
		h.Length++
		return
	}
	//非空链表则先找到最后一个节点，然后追加
	tempNode := h.Head
	for tempNode.Next != nil {
		tempNode = tempNode.Next
	}

	//追加前的尾节点的next指向新的尾节点
	tempNode.Next = node
	//表头信息+1
	h.Length++
}

//新节点插入在指定位置的后面
func (h *SingleHead) AddNode(insertIndex int, nodeData interface{}) bool {
	if h.Head == nil {
		panic("unclear head info")
	}

	//判断插入位置是否合法，索引从0开始，所以插入位置最大不能超过length-1
	if insertIndex < 0 || insertIndex >= h.Length {
		panic("illegal insert position")
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	tempNode := h.Head
	for i := 0; i <= insertIndex; i++ {
		tempNode = tempNode.Next
		if tempNode == nil {
			return false
		}
	}

	node := &SingleNode{
		Data: nodeData,
		Next: tempNode.Next,
	}

	tempNode.Next = node
	h.Length += 1
	return true
}

//在某个节点后添加一个新节点
func (h *SingleHead) AddAfter(preNode *SingleNode, nodeData interface{}) bool {
	if h.Head == nil {
		panic("unclear head info")
	}
	if preNode == nil {
		panic("pre node is nil")
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	//从头节点开始遍历
	ptr := h.Head
	for ptr != nil {
		if ptr == preNode {
			next := ptr.Next
			node := &SingleNode{
				Data: nodeData,
				Next: next,
			}
			ptr.Next = node
			h.Length++
			return true
		}
		ptr = ptr.Next
	}
	return false
}

//获取指定位置的节点
func (h *SingleHead) GetNode(index int) *SingleNode {
	if h.Head == nil {
		panic("unclear head info")
	}

	if index < 0 || index >= h.Length {
		panic("illegal insert position")
	}

	h.mu.RLock()
	defer h.mu.RUnlock()

	tempNode := h.Head
	count := 0
	for tempNode != nil && count <= index {
		if count == index {
			break
		}
		tempNode = tempNode.Next
		count++
	}

	return tempNode
}

//删除指定位置的节点
func (h *SingleHead) DeleteNode(index int) bool {
	if h.Head == nil {
		panic("unclear head info")
	}

	if index < 0 || index >= h.Length {
		panic("illegal delete position")
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	//删除头节点
	if index == 0 {
		next := h.Head.Next
		h.Head = next
		h.Length--
		return true
	}

	tempNode := h.Head
	count := 0
	for tempNode != nil && count < index {
		if count == index-1 && tempNode.Next != nil {
			next := tempNode.Next.Next
			tempNode.Next = next
			h.Length--
			return true
		}
		tempNode = tempNode.Next
		count++
	}
	return false
}

//获取节点数
func (h *SingleHead) GetNodeCount() int {
	return h.Length
}

//遍历链表
func (h *SingleHead) ListLink() (result string) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	result += fmt.Sprintf("{Length=%+v; Head=%p}——>", h.Length, h.Head)
	tempNode := h.Head
	for tempNode != nil {
		result += fmt.Sprintf("{Data=%+v; Next=%p}——>", tempNode.Data, tempNode.Next)
		tempNode = tempNode.Next
	}
	result += "NULL"
	return
}
