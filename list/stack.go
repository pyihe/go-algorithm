package list

import (
	"sync"
)

/*
	栈或队列(go标准包包含有采用双链表实现的list), 这里只是自己利用切片模拟栈或者队列的实现
*/

type (
	List interface {
		Len() int
		UnsafeLen() int
		Push(ele interface{}) bool
		Pop() (interface{}, bool)
		Del(i int) bool
		UnsafeDel(i int) bool
		Get(i int) interface{}
		UnsafeGet(i int) interface{}
		Set(index int, v interface{}) bool
		UnsafeSet(index int, v interface{}) bool
		Range(func(i int, v interface{}))
		UnsafeRange(func(i int, v interface{}))
		Index(data interface{}) (i int, ok bool)
		UnsafeIndex(data interface{}) (i int, ok bool)
	}

	list struct {
		mu    *sync.RWMutex
		t     ListType
		count int
		data  []interface{}
	}

	ListType int
)

const (
	ListTypeQueue ListType = iota + 1
	ListTypeStack
)

var _ List = &list{}

func NewList(t ListType, defaultCap int) List {
	if t != ListTypeStack && t != ListTypeQueue {
		panic("unknown list type")
	}
	q := &list{
		mu:    &sync.RWMutex{},
		count: 0,
		t:     t,
		data:  make([]interface{}, 0, defaultCap),
	}
	return q
}

func (q *list) init() {
	q.count = 0
	q.data = make([]interface{}, 0)
}

func (q *list) checkLen(i int) {
	if q == nil {
		q.init()
	}
	if q.count-1 < i {
		panic("out of range")
	}
}

func (q *list) Len() int {
	if q == nil {
		return 0
	}
	q.mu.RLock()
	defer q.mu.RUnlock()
	return q.count
}

func (q *list) UnsafeLen() int {
	if q == nil {
		return 0
	}
	return q.count
}

//add
func (q *list) Push(ele interface{}) bool {
	if q == nil {
		return false
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	q.count++
	q.data = append(q.data, ele)
	return true
}

//get&remove
func (q *list) Pop() (interface{}, bool) {
	if q == nil {
		return nil, false
	}
	var data interface{}
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.count == 0 {
		return nil, false
	}
	switch q.t {
	case ListTypeQueue:
		data = q.data[0]
		q.data = q.data[1:]
	case ListTypeStack:
		data = q.data[q.count-1]
		q.data = q.data[:q.count-1]
	}
	q.count--
	return data, true
}

//update
func (q *list) Set(i int, v interface{}) bool {
	if q == nil {
		return false
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	q.checkLen(i)
	q.data[i] = v
	return true
}

func (q *list) UnsafeSet(i int, v interface{}) bool {
	if q == nil {
		return false
	}
	q.checkLen(i)
	q.data[i] = v
	return true
}

//del
func (q *list) Del(i int) bool {
	if q == nil {
		return false
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	q.checkLen(i)
	q.data = append(q.data[:i], q.data[i+1:]...)
	q.count--
	return true
}

func (q *list) UnsafeDel(i int) bool {
	if q == nil {
		return false
	}
	q.checkLen(i)
	q.data = append(q.data[:i], q.data[i+1:]...)
	q.count--
	return true
}

//Get
func (q *list) Get(i int) interface{} {
	if q == nil {
		return nil
	}
	q.mu.RLock()
	defer q.mu.RUnlock()
	q.checkLen(i)
	d := q.data[i]
	return d
}

func (q *list) UnsafeGet(i int) interface{} {
	if q == nil {
		return nil
	}
	q.checkLen(i)
	d := q.data[i]
	return d
}

//safe range
func (q *list) Range(f func(i int, v interface{})) {
	q.mu.Lock()
	defer q.mu.Unlock()
	switch q.t {
	case ListTypeQueue:
		for i, v := range q.data {
			f(i, v)
		}
	case ListTypeStack:
		for i := q.count - 1; i >= 0; i-- {
			f(i, q.data[i])
		}
	}

}

//unsafe range
func (q *list) UnsafeRange(f func(i int, v interface{})) {
	switch q.t {
	case ListTypeQueue:
		for i, v := range q.data {
			f(i, v)
		}
	case ListTypeStack:
		for i := q.count - 1; i >= 0; i-- {
			f(i, q.data[i])
		}
	}
}

func (q *list) Index(data interface{}) (int, bool) {
	q.mu.RLock()
	defer q.mu.RUnlock()
	for i, v := range q.data {
		if v == data {
			return i, true
		}
	}
	return 0, false
}

func (q *list) UnsafeIndex(data interface{}) (int, bool) {
	for i, v := range q.data {
		if v == data {
			return i, true
		}
	}
	return 0, false
}
