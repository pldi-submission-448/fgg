package main

import "fmt"

type List interface {
	Reverse() List
}


type Cons struct {
	head int
	tail List
}

func (this Cons) Reverse() List {
	var xs List = this
	var ys List = Nil{}
	for {
		if _, ok := xs.(Nil); ok {
			return ys
		}
		if _xs, ok := xs.(Cons); ok {
			ys = Cons{
				head: _xs.head,
				tail: ys,
			}
			xs = _xs.tail
		}
	}
}


type Nil struct {}

func (this Nil) Reverse() List {
	return this
}

func StaticReverse(xs List) List {
	var ys List = Nil{}
	for {
		if _, ok := xs.(Nil); ok {
			return ys
		}
		if _xs, ok := xs.(Cons); ok {
			ys = Cons{
				head: _xs.head,
				tail: ys,
			}
			xs = _xs.tail
		}
	}
}


/*type List struct{
	head int//T
	tail *List
}*/

func (l Cons) append(last int) List {
	var newTail List
	if _, ok := l.tail.(Nil); ok {
		newTail = Cons{last, Nil{}}
	} else if _tail, ok := l.tail.(Cons); ok {
		newTail = _tail.append(last)
	}
	return Cons{l.head, newTail}
}

func MakeListFromArray(a []int) List {
	var xs List = Nil{}
	for i := len(a) - 1; i >= 0; i-- {
		xs = Cons{
			head: a[i],
			tail: xs,
		}
	}
	return xs
}

func BenchmarkStaticListReverse() {
	const iterations = 10000
	a := make([]int, iterations)
	for i := 0; i < iterations; i++ {
		a[i] = i
	}
	l := MakeListFromArray(a)
	for i := 0; i < iterations; i++ {
		if i %100 == 0 {
			fmt.Println(i)
		}
		l = StaticReverse(l)
	}
	fmt.Println(l.(Cons).head)
}

func BenchmarkInstanceListReverse() {
	const iterations = 10000
	a := make([]int, iterations)
	for i := 0; i < iterations; i++ {
		a[i] = i
	}
	l := MakeListFromArray(a)
	for i := 0; i < iterations; i++ {
		if i %100 == 0 {
			fmt.Println(i)
		}
		l = l.Reverse()
	}
	fmt.Println(l.(Cons).head)
}

func main()  {
	//BenchmarkStaticListReverse()
	BenchmarkInstanceListReverse()
}
//func Cons()
//func reverse(xs *List) List {
//	var ys List
//	for {
//		if xs == nil {
//			return ys
//		} else {
//			ys =
//		}
//	}
//}

type Cell struct {
	elem int
}

func (c Cell) add(elem int) Cell {
	c.elem = elem
	return c
}

func (c Cell) get() int {
	return c.elem
}

func CellTest(iterations int) {
	c := Cell{0}
	c = c.add(42)
	for i := 0; i < iterations; i++ {
		c.get()
	}
}

type Vector struct {
	elementData []int
	elementCount int
}

func (v Vector) reverseElements() {
	for i := 0; i < v.elementCount / 2; i++ {
		e := v.elementData[v.elementCount -i - 1]
		v.elementData[v.elementCount -i -1] = v.elementData[i]
		v.elementData[i] = e
	}
}

func (v *Vector) get(i int) int/*T*/ {
	return v.elementData[i]
}

func (v *Vector) set(i int, x int) {
	v.elementData[i] = x
}

func (v *Vector) size() int {
	return len(v.elementData)
}

func externalReverseElements(v *Vector) {
	for i := 0; i < v.size() / 2; i++ {
		e := v.get(v.size() -i -1)
		v.set(v.size() -i -1, v.get(i))
		v.set(i, e)
	}
}
