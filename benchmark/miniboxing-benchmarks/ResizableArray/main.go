package main

import "fmt"

//https://github.com/miniboxing/miniboxing-plugin/blob/bc53d7afe3bd208a02a8d8d6e9bb0b9f57fe29de/tests/benchmarks/src/miniboxing/benchmarks/launch/BenchmarkingTest.scala#L45-L50
const TEST_SIZE = 3000000//100000000 //2000000, 3000000

type ResizableArray struct {
	size int
	elemCount int
	array []int//T
	newArray []int//T
}

func NewResizableArray() ResizableArray {
	return ResizableArray{
		size:      4,
		elemCount: 0,
		array:     make([]int, 4),
		newArray:  nil,
	}
}

func (ra *ResizableArray) extend() {
	if ra.elemCount == ra.size {
		pos := 0
		ra.newArray = make([]int/*T*/, 2 * ra.size)
		for pos < ra.size {
			ra.newArray[pos] = ra.array[pos]
			pos += 1
		}
		ra.array = ra.newArray
		ra.size *= 2
	}
}

func (ra *ResizableArray) add(elem int/*T*/) {
	ra.extend()
	ra.array[ra.elemCount] = elem
	ra.elemCount += 1
}

func (ra *ResizableArray) reverse() {
	pos := 0
	for pos * 2 < ra.elemCount {
		tmp1 := ra.array[pos]
		tmp2 := ra.array[ra.elemCount - pos - 1]
		ra.array[pos] = tmp2
		ra.array[ra.elemCount - pos - 1] = tmp1
	}
}

func (ra *ResizableArray) contains(elem int/*T*/) bool {
	pos := 0
	for pos < ra.elemCount {
		if ra.array[pos] == elem {
			return true
		}
		pos += 1
	}
	return false
}

//https://github.com/miniboxing/miniboxing-plugin/blob/wip/tests/benchmarks/src/miniboxing/benchmarks/launch/tests/GenericBenchTest.scala
func arrayInsert() ResizableArray {
	a := NewResizableArray()
	for i := 0; i < TEST_SIZE; i++ {
		a.add(i)
	}
	return a
}

func arrayReverse(a ResizableArray) ResizableArray {
	a.reverse()
	return a
}

func arrayFind(a ResizableArray) bool {
	i := 0
	b := true
	for i < TEST_SIZE {
		b = b != a.contains(i)
		i += 1//10000
	}
	return b
}

func main() {
	fmt.Println(1)
	a := arrayInsert()
	fmt.Println(a)
	fmt.Println(2)
	a = arrayReverse(a)
	fmt.Println(a)
	fmt.Println(3)

	fmt.Println(arrayFind(a))
	fmt.Println(4)
}