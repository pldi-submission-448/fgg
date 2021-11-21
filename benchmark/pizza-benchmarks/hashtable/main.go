package main

import "strconv"

const ITERATIONS = 10000
type Hashtable struct {
	_inner map[interface{}]interface{}
}

func (this Hashtable) get(key interface{}) interface{} {
	return this._inner[key]
}

func (this Hashtable) set(key interface{}, value interface{}) {
	this._inner[key] = value
}

func main() {
	ht := Hashtable{make(map[interface{}]interface{})}
	for i := 1; i < 1000; i++ {
		ht.set(i, "A" + strconv.Itoa(i))
	}
	for j := 0; j < ITERATIONS; j++ {
		for i := 0; i < 1000; i++ {
			ht.get(i)
		}
	}


}