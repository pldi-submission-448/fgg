package main;

import "strconv"
const ITERATIONS = 10000

type Dummyᐸᐳ struct {};
type Intᐸᐳ int;
type Stringᐸᐳ string;
type _mapᐸᐳ struct {};
type HashtableᐸIntᐸᐳᐨStringᐸᐳᐳ struct { _inner map[Intᐸᐳ]Stringᐸᐳ };
func (this *HashtableᐸIntᐸᐳᐨStringᐸᐳᐳ) getᐸᐳ(key Intᐸᐳ) Stringᐸᐳ { return this._inner[key]
};
func (this HashtableᐸIntᐸᐳᐨStringᐸᐳᐳ) get___Int___String__() Top { return this;
};
func (this *HashtableᐸIntᐸᐳᐨStringᐸᐳᐳ) setᐸᐳ(key Intᐸᐳ, value Stringᐸᐳ) Dummyᐸᐳ {
	this._inner[key] = value
	return Dummyᐸᐳ{};
};
func (this HashtableᐸIntᐸᐳᐨStringᐸᐳᐳ) set___Int___String___Dummy__() Top { return this;
};
func (this Dummyᐸᐳ) Mainᐸᐳ() Dummyᐸᐳ {
	ht := HashtableᐸIntᐸᐳᐨStringᐸᐳᐳ{make(map[Intᐸᐳ]Stringᐸᐳ)}
	for i := 0; i < 1000; i++ {
		ht.setᐸᐳ(Intᐸᐳ(i), Stringᐸᐳ("A" + strconv.Itoa(i)))
	}
	for j := 0; j < ITERATIONS; j++ {
		for i := 0; i < 1000; i++ {
			ht.getᐸᐳ(Intᐸᐳ(i))
		}
	}
	return this;
};
func (this Dummyᐸᐳ) Main___Dummy__() Top { return this;
};
type Top interface {};
func main() { _ = Dummyᐸᐳ{}.Mainᐸᐳ() }
