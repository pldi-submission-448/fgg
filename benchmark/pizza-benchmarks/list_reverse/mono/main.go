package main;

import (
	"fmt"
	"strconv"
)

type Anyᐸᐳ interface {};
type Dummyᐸᐳ struct {};
//type intᐸᐳ struct {};
type intᐸᐳ int;
type string_ᐸᐳ string;

//type string_ᐸᐳ struct {};
func (this Dummyᐸᐳ) xᐸᐳ(a Anyᐸᐳ, b Anyᐸᐳ) Dummyᐸᐳ { return Dummyᐸᐳ{};
};
func (this Dummyᐸᐳ) x___Any___Any___Dummy__() Top { return this;
};
type Listᐸintᐸᐳᐳ interface { Reverseᐸᐳ() Listᐸintᐸᐳᐳ; Reverse___List_int___() Top };
type Listᐸstring_ᐸᐳᐳ interface { Reverseᐸᐳ() Listᐸstring_ᐸᐳᐳ; Reverse___List_string____() Top };

type Consᐸstring_ᐸᐳᐳ struct { head string_ᐸᐳ; tail Listᐸstring_ᐸᐳᐳ };
type Consᐸintᐸᐳᐳ struct { head intᐸᐳ; tail Listᐸintᐸᐳᐳ };
func (this Consᐸstring_ᐸᐳᐳ) Reverseᐸᐳ() Listᐸstring_ᐸᐳᐳ {
	var xs Listᐸstring_ᐸᐳᐳ = this
	var ys Listᐸstring_ᐸᐳᐳ = Nilᐸstring_ᐸᐳᐳ{}
	for {
		if _, ok := xs.(Nilᐸstring_ᐸᐳᐳ); ok {
			return ys
		}
		if _xs, ok := xs.(Consᐸstring_ᐸᐳᐳ); ok {
			ys = Consᐸstring_ᐸᐳᐳ{
				head: _xs.head,
				tail: ys,
			}
			xs = _xs.tail
		}
	}
};
func (this Consᐸintᐸᐳᐳ) Reverseᐸᐳ() Listᐸintᐸᐳᐳ {
	var xs Listᐸintᐸᐳᐳ = this
	var ys Listᐸintᐸᐳᐳ = Nilᐸintᐸᐳᐳ{}
	for {
		if _, ok := xs.(Nilᐸintᐸᐳᐳ); ok {
			return ys
		}
		if _xs, ok := xs.(Consᐸintᐸᐳᐳ); ok {
			ys = Consᐸintᐸᐳᐳ{
				head: _xs.head,
				tail: ys,
			}
			xs = _xs.tail
		}
	}
};

func (this Consᐸstring_ᐸᐳᐳ) Reverse___List_string____() Top { return this;
};
func (this Consᐸintᐸᐳᐳ) Reverse___List_int___() Top { return this;
};
type Nilᐸintᐸᐳᐳ struct {};
type Nilᐸstring_ᐸᐳᐳ struct {};
func (this Nilᐸstring_ᐸᐳᐳ) Reverseᐸᐳ() Listᐸstring_ᐸᐳᐳ { return this;
};
func (this Nilᐸintᐸᐳᐳ) Reverseᐸᐳ() Listᐸintᐸᐳᐳ { return this;
};
func (this Nilᐸintᐸᐳᐳ) Reverse___List_int___() Top { return this;
};
func (this Nilᐸstring_ᐸᐳᐳ) Reverse___List_string____() Top { return this;
};
func (this Dummyᐸᐳ) StaticReverseᐸintᐸᐳᐳ(xs Listᐸintᐸᐳᐳ) Listᐸintᐸᐳᐳ {
	var ys Listᐸintᐸᐳᐳ = Nilᐸintᐸᐳᐳ{}
	for {
		if _, ok := xs.(Nilᐸintᐸᐳᐳ); ok {
			return ys
		}
		if _xs, ok := xs.(Consᐸintᐸᐳᐳ); ok {
			ys = Consᐸintᐸᐳᐳ{
				head: _xs.head,
				tail: ys,
			}
			xs = _xs.tail
		}
	}
};
func (this Dummyᐸᐳ) StaticReverseᐸstring_ᐸᐳᐳ(xs Listᐸstring_ᐸᐳᐳ) Listᐸstring_ᐸᐳᐳ {
	var ys Listᐸstring_ᐸᐳᐳ = Nilᐸstring_ᐸᐳᐳ{}
	for {
		if _, ok := xs.(Nilᐸstring_ᐸᐳᐳ); ok {
			return ys
		}
		if _xs, ok := xs.(Consᐸstring_ᐸᐳᐳ); ok {
			ys = Consᐸstring_ᐸᐳᐳ{
				head: _xs.head,
				tail: ys,
			}
			xs = _xs.tail
		}
	}
};

func MakeListFromArray_int(a []intᐸᐳ) Listᐸintᐸᐳᐳ {
	var xs Listᐸintᐸᐳᐳ = Nilᐸintᐸᐳᐳ{}
	for i := len(a) - 1; i >= 0; i-- {
		xs = Consᐸintᐸᐳᐳ{
			head: a[i],
			tail: xs,
		}
	}
	return xs
}

func MakeListFromArray_string(a []string_ᐸᐳ) Listᐸstring_ᐸᐳᐳ {
	var xs Listᐸstring_ᐸᐳᐳ = Nilᐸstring_ᐸᐳᐳ{}
	for i := len(a) - 1; i >= 0; i-- {
		xs = Consᐸstring_ᐸᐳᐳ{
			head: a[i],
			tail: xs,
		}
	}
	return xs
}

func BenchmarkStaticListReverse_int() {
	const iterations = 10000
	a := make([]intᐸᐳ, iterations)
	for i := 0; i < iterations; i++ {
		a[i] = intᐸᐳ(i)
	}
	l := MakeListFromArray_int(a)
	for i := 0; i < iterations; i++ {
		if i %100 == 0 {
			//fmt.Println(i)
		}
		l = Dummyᐸᐳ{}.StaticReverseᐸintᐸᐳᐳ(l)
	}
	fmt.Println(l.(Consᐸintᐸᐳᐳ).head)
}

func BenchmarkStaticListReverse_string() {
	const iterations = 10000
	a := make([]string_ᐸᐳ, iterations)
	for i := 0; i < iterations; i++ {
		a[i] = string_ᐸᐳ(strconv.Itoa(i))
	}
	l := MakeListFromArray_string(a)
	for i := 0; i < iterations; i++ {
		if i %100 == 0 {
			//fmt.Println(i)
		}
		l = Dummyᐸᐳ{}.StaticReverseᐸstring_ᐸᐳᐳ(l)
	}
	fmt.Println(l.(Consᐸstring_ᐸᐳᐳ).head)
}

func BenchmarkInstanceListReverse_int() {
	const iterations = 10000
	a := make([]intᐸᐳ, iterations)
	for i := 0; i < iterations; i++ {
		a[i] = intᐸᐳ(i)
	}
	l := MakeListFromArray_int(a)
	for i := 0; i < iterations; i++ {
		if i %100 == 0 {
			//fmt.Println(i)
		}
		l = l.Reverseᐸᐳ()
	}
	fmt.Println(l.(Consᐸintᐸᐳᐳ).head)
}

func BenchmarkInstanceListReverse_string() {
	const iterations = 10000
	a := make([]string_ᐸᐳ, iterations)
	for i := 0; i < iterations; i++ {
		a[i] = string_ᐸᐳ(strconv.Itoa(i))
	}
	l := MakeListFromArray_string(a)
	for i := 0; i < iterations; i++ {
		if i %100 == 0 {
			//fmt.Println(i)
		}
		l = l.Reverseᐸᐳ()
	}
	fmt.Println(l.(Consᐸstring_ᐸᐳᐳ).head)
}


func (this Dummyᐸᐳ) StaticReverse__β1_Any____List_β1__List_β1_() Top { return this;
};

func (this Dummyᐸᐳ) Main1___Dummy__() Top { return this;
};
func (this Dummyᐸᐳ) Main2___Dummy__() Top { return this;
};
type Top interface {};

func main() {
	fmt.Println("1")
	BenchmarkInstanceListReverse_int()
	fmt.Println("2")
	BenchmarkInstanceListReverse_string()
	fmt.Println("3")
	BenchmarkStaticListReverse_int()
	fmt.Println("4")
	BenchmarkInstanceListReverse_string()
	fmt.Println("5")
}
