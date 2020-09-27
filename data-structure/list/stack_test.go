package list

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

var (
	queue = NewList(ListTypeStack, 1024)
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestList_Push(t *testing.T) {
	var toBePush []int
	for i := 0; i < 100; i++ {
		toBePush = append(toBePush, rand.Intn(101))
	}
	log.Printf("待添加的数据为: %v %v\n", len(toBePush), toBePush)

	for _, v := range toBePush {
		queue.Push(v)
	}
	log.Printf("添加数据个数为: %v\n", queue.Len())
}

func TestList_Pop(t *testing.T) {
	for i := 0; i<100;i++ {
		data, ok := queue.Pop()
		if ok {
			log.Printf("%v\n", data)
		}
	}
	log.Printf("Pop后剩余数据个数为: %v\n", queue.Len())
}