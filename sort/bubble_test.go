package sort

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getRandNums(cnt int) (data []int) {
	for i := 0; i < cnt; i++ {
		data = append(data, rand.Intn(cnt+1))
	}
	log.Printf("待排序的列表为: %v\n", data)
	return
}

func TestBubbleSort(t *testing.T) {
	data := getRandNums(10)
	BubbleSort(data)
	log.Printf("冒泡排序结果为: %v\n", data)
}

func TestInsertSort(t *testing.T) {
	data := getRandNums(10)
	InsertSort(data)
	log.Printf("插入排序结果为: %v\n", data)
}

func TestQuickSort(t *testing.T) {
	data := getRandNums(10)
	QuickSort(data)
	log.Printf("快速排序结果为: %v\n", data)
}

func TestSelectSort(t *testing.T) {
	data := getRandNums(10)
	SelectSort(data)
	log.Printf("选择排序结果为: %v\n", data)
}

func TestMergeSortByRecursion(t *testing.T) {
	data := getRandNums(20)
	data = MergeSortByIter(data)
	log.Printf("归并排序结果为: %v\n", data)
}

func TestShellSort(t *testing.T) {
	data := getRandNums(20)
	ShellSort(data)
	log.Printf("希尔排序结果为: %v\n", data)
}

func TestCountingSort(t *testing.T) {
	data := getRandNums(20)
	var max int
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	CountingSort(data, max)
	log.Printf("计数排序结果为: %v\n", data)
}

func TestHeapSort(t *testing.T) {
	data := getRandNums(10)
	HeapSort(data)
	log.Printf("堆排序结果为: %v\n", data)
}

func TestBucketSort(t *testing.T) {
	data := getRandNums(10)
	BucketSort(data)
	log.Printf("桶排序结果为: %v\n", data)
}

func TestRadixSort(t *testing.T) {
	data := getRandNums(100)
	RadixSort(data)
	log.Printf("基数排序结果为: %v\n", data)
}

//
//
//
//
//
//
//
//

func BenchmarkBubbleSort(b *testing.B) {
	data := getRandNums(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BubbleSort(data)
	}
}

func BenchmarkInsertSort(b *testing.B) {
	data := getRandNums(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertSort(data)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	data := getRandNums(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(data)
	}
}

func BenchmarkSelectSort(b *testing.B) {
	data := getRandNums(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SelectSort(data)
	}
}

func BenchmarkMergeSortByRecursion(b *testing.B) {
	data := getRandNums(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeSortByRecursion(data)
	}
}
