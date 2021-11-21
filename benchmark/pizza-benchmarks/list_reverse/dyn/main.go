package main;

import (
	"fmt"
	"strconv"
)

type Any interface {};
type Anyǂ interface {};
type Dummyǂ struct {};
type Dummyǂ_meta struct {  };
type intǂ struct {};
type intǂ_meta struct {  };
type string_ǂ struct {};
type string_ǂ_meta struct {  };
func (thisǂ Dummyǂ) xǂ(aǂ Any, bǂ Any) Any { return Dummyǂ{};
};
type Listǂ interface { Reverseǂ() Any };
type Consǂ struct { headǂ Any; tailǂ Any };
type Consǂ_meta struct { _type_0 _type_metadata };
func (thisǂ Consǂ) Reverseǂ() Any {
	var xs Any = thisǂ
	var ys Any = Nilǂ{}
	for {
		if _, ok := xs.(Nilǂ); ok {
			return ys
		}
		if _xs, ok := xs.(Consǂ); ok {
			ys = Consǂ{
				headǂ: _xs.headǂ,
				tailǂ: ys,
			}
			xs = _xs.tailǂ
		}
	}
};
type Nilǂ struct {};
type Nilǂ_meta struct { _type_0 _type_metadata };
func (thisǂ Nilǂ) Reverseǂ() Any { return thisǂ;
};
func (thisǂ Dummyǂ) StaticReverseǂ(xsǂ Any) Any {
	var ys Any = Nilǂ{}
	for {
		if _, ok := xsǂ.(Nilǂ); ok {
			return ys
		}
		if _xs, ok := xsǂ.(Consǂ); ok {
			ys = Consǂ{
				headǂ: _xs.headǂ,
				tailǂ: ys,
			}
			xsǂ = _xs.tailǂ
		}
	}
};
func (thisǂ Dummyǂ) Main1ǂ() Any {
	return thisǂ.xǂ(Consǂ{intǂ{}, Nilǂ{}}.Reverseǂ(), Dummyǂ{}.StaticReverseǂ(Nilǂ{}));
};
func (thisǂ Dummyǂ) Main2ǂ() Any { return thisǂ.xǂ(Consǂ{string_ǂ{}, Nilǂ{}}.Reverseǂ(), Dummyǂ{}.StaticReverseǂ(Nilǂ{}));
};
type spec_metadata_0 struct { _type_0 _type_metadata };
type spec_metadata_1 struct { _type_0 _type_metadata; _type_1 _type_metadata };
type _type_metadata interface { tryCast(x Any) Any; assertEq(x _type_metadata) Any };

func MakeListFromArray(a []Any) Listǂ {
	var xs Listǂ = Nilǂ{}
	for i := len(a) - 1; i >= 0; i-- {
		xs = Consǂ{
			headǂ:  a[i],
			tailǂ:  xs,
		}
	}
	return xs
}

func BenchmarkStaticListReverse_int() {
	const iterations = 10000
	a := make([]Any, iterations)
	for i := 0; i < iterations; i++ {
		a[i] = i
	}
	var l Any = MakeListFromArray(a)
	for i := 0; i < iterations; i++ {
		if i %100 == 0 {
			//fmt.Println(i)
		}
		l = Dummyǂ{}.StaticReverseǂ(l)
	}
	fmt.Println(l.(Consǂ).headǂ)
}

func BenchmarkStaticListReverse_string() {
	const iterations = 10000
	a := make([]Any, iterations)
	for i := 0; i < iterations; i++ {
		a[i] = strconv.Itoa(i)
	}
	var l Any = MakeListFromArray(a)
	for i := 0; i < iterations; i++ {
		if i %100 == 0 {
			//fmt.Println(i)
		}
		l = Dummyǂ{}.StaticReverseǂ(l)
	}
}

func BenchmarkInstanceListReverse_int() {
	const iterations = 10000
	a := make([]Any, iterations)
	for i := 0; i < iterations; i++ {
		a[i] = i
	}
	var l Any = MakeListFromArray(a)
	for i := 0; i < iterations; i++ {
		if i %100 == 0 {
			//fmt.Println(i)
		}
		l = l.(Listǂ).Reverseǂ()
	}
	fmt.Println(l.(Consǂ).headǂ)
}

func BenchmarkInstanceListReverse_string() {
	const iterations = 10000
	a := make([]Any, iterations)
	for i := 0; i < iterations; i++ {
		a[i] = strconv.Itoa(i)
	}
	var l Any = MakeListFromArray(a)
	for i := 0; i < iterations; i++ {
		if i %100 == 0 {
			//fmt.Println(i)
		}
		l = l.(Listǂ).Reverseǂ()
	}
	fmt.Println(l.(Consǂ).headǂ)
}

func main() {
	fmt.Println("1")
	BenchmarkInstanceListReverse_int()
	fmt.Println("2")
	BenchmarkInstanceListReverse_string()
	fmt.Println("3")
	BenchmarkStaticListReverse_int()
	fmt.Println("4")
	BenchmarkStaticListReverse_string()
	fmt.Println("5")

}