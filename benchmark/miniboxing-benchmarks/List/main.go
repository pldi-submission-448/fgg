package main

import "fmt"

const TEST_SIZE = 100

type List interface {
	Contains(x int) bool
}


type Cons struct {
	head int
	tail List
}

func (c Cons) Contains(x int) bool {
	var xs List = c
	for {
		if xs_, ok := xs.(Cons); ok {
			if xs_.head == x {
				return true
			} else {
				xs = xs_.tail
			}
		}
		if _, ok := xs.(Nil); ok {
			return false
		}
	}
}


type Nil struct {}

func (n Nil) Contains(x int) bool {
	return false
}

func ListInsert() List {
	var l List = Nil{}
	for i := 0; i < TEST_SIZE; i++ {
		l = Cons{
			head: i,
			tail: l,
		}
	}
	return l
}

func ListFind(l List) bool {
	b := true
	for i := 0; i < TEST_SIZE; i+=10000 {
		b = b != l.Contains(i)
	}
	return b
}

func main() {
	fmt.Println(1)
	l := ListInsert()
	fmt.Println(2)
	ListFind(l)
	fmt.Println(3)
}