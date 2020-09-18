package tree

import (
	"testing"
)

var (
	bt = NewBTree(5)
)

func init() {
	var num = []int{39, 22, 97, 41, 53, 13, 21, 40, 30, 27, 33, 36, 35, 34, 24, 29, 26, 17, 28, 29, 31, 32}
	for _, v := range num {
		bt.Insert(v, nil)
	}
}

func TestBTree_Insert(t *testing.T) {
	t.Logf("original: %v\n", bt)
	var num = []int{3, 8, 39, 11, 23, 29, 50, 28}
	for _, v := range num {
		bt.Insert(v, nil)
	}
	t.Logf("insert result: %v\n", bt)
	searchResult := bt.Search(27)
	if searchResult != nil {
		t.Logf("search result: %v\n", searchResult)
	}
}

func TestBTree_Remove(t *testing.T) {
	t.Logf("original: %v\n", bt)
	t.Logf("remove 21: %v\n", bt.Remove(21))
	t.Logf("after remove 21: %v\n", bt)
	t.Logf("remove 27: %v\n", bt.Remove(27))
	t.Logf("after remove 27: %v\n", bt)
	t.Logf("remove 33: %v\n", bt.Remove(33))
	t.Logf("after remove 33: %v\n", bt)
}
