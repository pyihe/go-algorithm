package link

import (
	"fmt"
	"sync"
)

/*
	双向链表(双向链表也可以是循环链表, 如果是循环链表则头节点的pre指向尾节点，尾节点的next指向头节点)
	1. 首尾追加节点
	2. 指定位置添加节点
	3. 获取指定位置的节点
	4. 删除指定位置的节点
	5. 获取节点个数
*/

//链表节点
type CircleNode struct {
	Data interface{} //数据域
	Pre  *CircleNode //指向前一个节点(对于头节点, pre指向尾节点)
	Next *CircleNode //指向后一个节点(对于尾节点, next指向头节点)
}

//头节点
type CircleHead struct {
	mu     sync.RWMutex
	Length int         //有多少个节点
	Guard  *CircleNode //哨兵节点
}

//new一个节点
func NewCircleNode(data interface{}) *CircleNode {
	return &CircleNode{
		Data: data,
		Pre:  nil,
		Next: nil,
	}
}

//new一个环形链表
func NewCircleHead() *CircleHead {
	guard := &CircleNode{
		Data: nil,
	}
	//初始化时哨兵节点的pre和next都指向它自己
	guard.Next = guard
	guard.Pre = guard

	return &CircleHead{
		Length: 0,
		Guard:  guard,
	}
}

//在尾节点后追加一个节点
func (c *CircleHead) AppendNode(value interface{}) *CircleNode {
	node := &CircleNode{
		Data: value,
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	tail := c.Guard.Pre //记录尾节点
	next := tail.Next   //尾节点的next
	tail.Next = node    //将node赋值给尾节点的next
	node.Pre = tail     //将tail赋值给node的pre
	node.Next = next    //将next赋值给node的next
	next.Pre = node     //将node赋值给next的pre
	c.Length++          //链表长度+1
	return node
}

//在指定位置添加一个节点
func (c *CircleHead) AddHead(value interface{}) *CircleNode {
	node := &CircleNode{
		Data: value,
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	head := c.Guard.Next
	pre := head.Pre
	node.Next = head
	head.Pre = node
	pre.Next = node
	node.Pre = pre
	c.Length++
	return node
}

//在指定位置添加节点, 如果是负数则从尾节点开始，反之从头节点开始
func (c *CircleHead) AddAt(at int, value interface{}) (*CircleNode, bool) {
	if at == 0 || at > c.Length || -at > c.Length {
		panic("illegal insert position")
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	temp := c.Guard
	n := at
	if at < 0 {
		at = -at
	}
	for i := 1; i <= at; i++ {
		if n > 0 {
			temp = temp.Next
		} else {
			temp = temp.Pre
		}
		if temp == c.Guard {
			return nil, false
		}
	}
	//找到at对应的节点后，执行插入
	var node *CircleNode
	var vNode = &CircleNode{
		Data: value,
	}
	if n > 0 {
		node = temp.Next
		temp.Next = vNode
		vNode.Pre = temp
		vNode.Next = node
		node.Pre = vNode
	} else {
		node = temp.Pre
		temp.Pre = vNode
		vNode.Next = temp
		node.Next = vNode
		vNode.Pre = node
	}
	c.Length++
	return vNode, true
}

func (c *CircleHead) GetNode(at int) *CircleNode {
	if at == 0 || at > c.Length || -at > c.Length {
		panic("illegal insert position")
	}
	temp := c.Guard
	n := at
	if at < 0 {
		at = -at
	}
	for i := 1; i <= at; i++ {
		if n > 0 {
			temp = temp.Next
		} else {
			temp = temp.Pre
		}
		if temp == c.Guard {
			return nil
		}
	}
	return temp
}

func (c *CircleHead) DeleteNode(pos int) bool {
	if pos == 0 || pos > c.Length || -pos > c.Length {
		panic("illegal insert position")
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	temp := c.Guard
	n := pos
	if pos < 0 {
		pos = -pos
	}
	for i := 1; i <= pos; i++ {
		if n > 0 {
			temp = temp.Next
		} else {
			temp = temp.Pre
		}
		if temp == c.Guard {
			return false
		}
	}
	next := temp.Next
	pre := temp.Pre
	pre.Next = next
	next.Pre = pre
	temp.Pre = nil
	temp.Data = nil
	temp.Next = nil
	c.Length--

	return true
}

func (c *CircleHead) Len() int {
	return c.Length
}

func (c *CircleHead) String() string {
	var result string
	node := c.Guard.Next
	for node != c.Guard {
		result += fmt.Sprintf("%v->", node.Data)
		node = node.Next
	}
	result += "guard"
	return result
}
