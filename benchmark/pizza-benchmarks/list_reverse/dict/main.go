package main;

import (
	"fmt"
	"strconv"
)

type Any interface {};
type Anyǂ interface {};
type AnyǂDictǂ struct { _type _type_metadata };
type Anyǂ_meta struct {  };
func (this Anyǂ_meta) tryCast(x Any) Any { x_ := x.(Anyǂ);
	return x_;
};
func (this Anyǂ_meta) assertEq(x _type_metadata) Any { x_ := x.(Anyǂ_meta);
	return x_;
};
type Dummyǂ struct {};
type Dummyǂ_meta struct {  };
func (this Dummyǂ_meta) tryCast(x Any) Any { _ = x.(Dummyǂ);
	return x;
};
func (this Dummyǂ_meta) assertEq(x _type_metadata) Any { return this;
};
type intǂ struct {};
type intǂ_meta struct {  };
func (this intǂ_meta) tryCast(x Any) Any { _ = x.(intǂ);
	return x;
};
func (this intǂ_meta) assertEq(x _type_metadata) Any { return this;
};
type string_ǂ struct {};
type string_ǂ_meta struct {  };
func (this string_ǂ_meta) tryCast(x Any) Any { _ = x.(string_ǂ);
	return x;
};
func (this string_ǂ_meta) assertEq(x _type_metadata) Any { return this;
};
func (thisǂ Dummyǂ) xǂ(aǂ Any, bǂ Any) Any { return Dummyǂ{};
};
func (thisǂ Dummyǂ) spec_x() spec_metadata_2 { return spec_metadata_2{Anyǂ_meta{}, Anyǂ_meta{}, Dummyǂ_meta{}};
};
type Listǂ interface { Reverseǂ() Any; spec_Reverse() spec_metadata_0 };
type ListǂDictǂ struct { Reverseǂ func (receiver Any) Any; _type _type_metadata };
type Listǂ_meta struct { _type_0 _type_metadata };
func (this Listǂ_meta) tryCast(x Any) Any { x_ := x.(Listǂ);
	Reverse_actual := x_.spec_Reverse()
	;
	Reverse_actual._type_0.assertEq(Listǂ_meta{this._type_0});
	return x_;
};
func (this Listǂ_meta) assertEq(x _type_metadata) Any { x_ := x.(Listǂ_meta);
	this._type_0.assertEq(x_._type_0);
	return x_;
};
type Consǂ struct { headǂ Any; tailǂ Any; dict_0 AnyǂDictǂ };
type Consǂ_meta struct { _type_0 _type_metadata };
func (this Consǂ_meta) tryCast(x Any) Any { _ = x.(Consǂ);
	this._type_0.assertEq(x.(Consǂ).dict_0._type);
	return x;
};
func (this Consǂ_meta) assertEq(x _type_metadata) Any { this._type_0.assertEq(x.(Consǂ_meta)._type_0);
	return this;
};
func (thisǂ Consǂ) Reverseǂ() Any {
	var xs Any = thisǂ
	var ys Any = Nilǂ{thisǂ.dict_0}
	for {
		if _, ok := xs.(Nilǂ); ok {
			return ys
		}
		if _xs, ok := xs.(Consǂ); ok {
			ys = Consǂ{
				headǂ: _xs.headǂ,
				tailǂ: ys,
				dict_0: thisǂ.dict_0,
			}
			xs = _xs.tailǂ
		}
	}
};
func (thisǂ Consǂ) spec_Reverse() spec_metadata_0 { return spec_metadata_0{Listǂ_meta{thisǂ.dict_0._type}};
};
type Nilǂ struct { dict_0 AnyǂDictǂ };
type Nilǂ_meta struct { _type_0 _type_metadata };
func (this Nilǂ_meta) tryCast(x Any) Any { _ = x.(Nilǂ);
	this._type_0.assertEq(x.(Nilǂ).dict_0._type);
	return x;
};
func (this Nilǂ_meta) assertEq(x _type_metadata) Any { this._type_0.assertEq(x.(Nilǂ_meta)._type_0);
	return this;
};
func (thisǂ Nilǂ) Reverseǂ() Any { return thisǂ;
};
func (thisǂ Nilǂ) spec_Reverse() spec_metadata_0 { return spec_metadata_0{Listǂ_meta{thisǂ.dict_0._type}};
};
func (thisǂ Dummyǂ) StaticReverseǂ(dict_0 AnyǂDictǂ, xsǂ Any) Any {
	var ys Any = Nilǂ{dict_0}
	for {
		if _, ok := xsǂ.(Nilǂ); ok {
			return ys
		}
		if _xs, ok := xsǂ.(Consǂ); ok {
			ys = Consǂ{
				headǂ: _xs.headǂ,
				tailǂ: ys,
				dict_0: dict_0,
			}
			xsǂ = _xs.tailǂ
		}
	}
};
func (thisǂ Dummyǂ) spec_StaticReverse() spec_metadata_2 { return spec_metadata_2{Anyǂ_meta{}, Listǂ_meta{param_index_0{}}, Listǂ_meta{param_index_0{}}};
};
func (thisǂ Dummyǂ) Main1ǂ() Any { return thisǂ.xǂ(Consǂ{intǂ{}, Nilǂ{AnyǂDictǂ{intǂ_meta{}}}, AnyǂDictǂ{intǂ_meta{}}}.Reverseǂ(), Dummyǂ{}.StaticReverseǂ(AnyǂDictǂ{intǂ_meta{}}, Nilǂ{AnyǂDictǂ{intǂ_meta{}}}));
};
func (thisǂ Dummyǂ) spec_Main1() spec_metadata_0 { return spec_metadata_0{Dummyǂ_meta{}};
};
func (thisǂ Dummyǂ) Main2ǂ() Any { return thisǂ.xǂ(Consǂ{string_ǂ{}, Nilǂ{AnyǂDictǂ{string_ǂ_meta{}}}, AnyǂDictǂ{string_ǂ_meta{}}}.Reverseǂ(), Dummyǂ{}.StaticReverseǂ(AnyǂDictǂ{string_ǂ_meta{}}, Nilǂ{AnyǂDictǂ{string_ǂ_meta{}}}));
};
func (thisǂ Dummyǂ) spec_Main2() spec_metadata_0 { return spec_metadata_0{Dummyǂ_meta{}};
};
type spec_metadata_0 struct { _type_0 _type_metadata };
type spec_metadata_2 struct { _type_0 _type_metadata; _type_1 _type_metadata; _type_2 _type_metadata };
type spec_metadata_1 struct { _type_0 _type_metadata; _type_1 _type_metadata };
type param_index_0 struct {};
func (this param_index_0) tryCast(x Any) Any {_ = x.(param_index_0); return x;};
func (this param_index_0) assertEq(x _type_metadata) Any {_ = x.(param_index_0); return x;};
type _type_metadata interface { tryCast(x Any) Any; assertEq(x _type_metadata) Any };
type __BidirectionalChannel struct {};
func (this __BidirectionalChannel) assertEq(x Any) Any { return x.(__BidirectionalChannel);
};
type __sendOnlyChannel struct {};
func (this __sendOnlyChannel) assertEq(x Any) Any { return x.(__sendOnlyChannel);
};
type __receiveOnlyChannel struct {};
func (this __receiveOnlyChannel) assertEq(x Any) Any { return x.(__receiveOnlyChannel);
};
type chan_wrapper struct { ch chan Any; _type _type_metadata };
type chan_direction interface { assertEq(x Any) Any };
type chan_metadata struct { dir chan_direction; _type _type_metadata };
func (this chan_metadata) tryCast(x Any) Any {
	_ = x.(chan_wrapper)._type.assertEq(this);
	return x
};
func (this chan_metadata) assertEq(x _type_metadata) Any {
	_ = x.(chan_metadata)._type.assertEq(this._type);
	return x.(chan_metadata).dir.assertEq(this.dir)
};

func MakeListFromArray(dict_0 AnyǂDictǂ, a []Any) Listǂ {
	var xs Listǂ = Nilǂ{dict_0: dict_0}
	for i := len(a) - 1; i >= 0; i-- {
		xs = Consǂ{
			headǂ:  a[i],
			tailǂ:  xs,
			dict_0: dict_0,
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
	var l Any = MakeListFromArray(AnyǂDictǂ{_type: intǂ_meta{}}, a)
	for i := 0; i < iterations; i++ {
		if i %100 == 0 {
			//fmt.Println(i)
		}
		l = Dummyǂ{}.StaticReverseǂ(AnyǂDictǂ{_type: intǂ_meta{}}, l)
	}
	fmt.Println(l.(Consǂ).headǂ)
}

func BenchmarkStaticListReverse_string() {
	const iterations = 10000
	a := make([]Any, iterations)
	for i := 0; i < iterations; i++ {
		a[i] = strconv.Itoa(i)
	}
	var l Any = MakeListFromArray(AnyǂDictǂ{_type: string_ǂ_meta{}}, a)
	for i := 0; i < iterations; i++ {
		if i %100 == 0 {
			//fmt.Println(i)
		}
		l = Dummyǂ{}.StaticReverseǂ(AnyǂDictǂ{_type: string_ǂ_meta{}}, l)
	}
}

func BenchmarkInstanceListReverse_int() {
	const iterations = 10000
	a := make([]Any, iterations)
	for i := 0; i < iterations; i++ {
		a[i] = i
	}
	var l Any = MakeListFromArray(AnyǂDictǂ{_type: intǂ_meta{}}, a)
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
	var l Any = MakeListFromArray(AnyǂDictǂ{_type: string_ǂ_meta{}}, a)
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