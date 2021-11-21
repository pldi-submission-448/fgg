package main;

import (
	"fmt"
	"strconv"
)

type Any interface {};
type Anyǂ interface {};
type AnyǂDictǂ struct {  };
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
type intǂ int;//struct {};
type intǂ_meta struct {  };
func (this intǂ_meta) tryCast(x Any) Any { _ = x.(intǂ);
	return x;
};
func (this intǂ_meta) assertEq(x _type_metadata) Any { return this;
};
type string_ǂ string;//struct {};
type string_ǂ_meta struct {  };
func (this string_ǂ_meta) tryCast(x Any) Any { _ = x.(string_ǂ);
	return x;
};
func (this string_ǂ_meta) assertEq(x _type_metadata) Any { return this;
};
type Vectorǂ struct { elementDataǂ []Any; elementCountǂ Any; dict_0 AnyǂDictǂ };
type Vectorǂ_meta struct { _type_0 _type_metadata };
func (this Vectorǂ_meta) tryCast(x Any) Any { _ = x.(Vectorǂ);
	return x;
};
func (this Vectorǂ_meta) assertEq(x _type_metadata) Any { this._type_0.assertEq(x.(Vectorǂ_meta)._type_0);
	return this;
};
func (vǂ *Vectorǂ) reverseElementsǂ() Any {
	count := len(vǂ.elementDataǂ)
	for i := 0; i < count / 2; i++ {
		e := vǂ.elementDataǂ[count -i - 1]
		vǂ.elementDataǂ[count -i -1] = vǂ.elementDataǂ[i]
		vǂ.elementDataǂ[i] = e
	}
	return Dummyǂ{};
};
func (vǂ Vectorǂ) spec_reverseElements() spec_metadata_0 { return spec_metadata_0{Dummyǂ_meta{}};
};
func (vǂ *Vectorǂ) getǂ(iǂ Any) Any { return vǂ.elementDataǂ[iǂ.(int)];
};

func (vǂ *Vectorǂ) setǂ(iǂ Any, xǂ Any) Any {
	vǂ.elementDataǂ[iǂ.(int)] = xǂ;
	return Dummyǂ{};
};

func (vǂ Vectorǂ) sizeǂ() Any {
	return len(vǂ.elementDataǂ)
};
func (vǂ Vectorǂ) spec_size() spec_metadata_0 { return spec_metadata_0{intǂ_meta{}};
};
func (thisǂ Dummyǂ) externalReverseElementsǂ(dict_0 AnyǂDictǂ, aǂ Any) Any {
	a := aǂ.(Vectorǂ)
	for i := 0; i < a.sizeǂ().(int) / 2; i++ {
		e := a.getǂ(a.sizeǂ().(int) - i -1)
		a.setǂ(a.sizeǂ().(int) - i -1, a.getǂ(i))
		a.setǂ(i, e)
	}
	return Dummyǂ{};
};
func (thisǂ Dummyǂ) spec_externalReverseElements() spec_metadata_2 { return spec_metadata_2{Anyǂ_meta{}, Vectorǂ_meta{param_index_0{}}, Dummyǂ_meta{}};
};
func (thisǂ Dummyǂ) xǂ(aǂ Any, bǂ Any) Any { return Dummyǂ{};
};
func (thisǂ Dummyǂ) spec_x() spec_metadata_2 { return spec_metadata_2{Anyǂ_meta{}, Anyǂ_meta{}, Dummyǂ_meta{}};
};

func (thisǂ Dummyǂ) spec_Main1() spec_metadata_0 { return spec_metadata_0{Dummyǂ_meta{}};
};

func (thisǂ Dummyǂ) spec_Main2() spec_metadata_0 { return spec_metadata_0{Dummyǂ_meta{}};
};
type spec_metadata_0 struct { _type_0 _type_metadata };
type spec_metadata_1 struct { _type_0 _type_metadata; _type_1 _type_metadata };
type spec_metadata_2 struct { _type_0 _type_metadata; _type_1 _type_metadata; _type_2 _type_metadata };
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
func BenchmarkVectorReverse_int() {
	const iterations = 10000
	v := Vectorǂ{
		elementDataǂ:  make([]Any, iterations),
		elementCountǂ: iterations,
		dict_0: AnyǂDictǂ{},
	}
	for i := 0; i < iterations; i++ {
		i_ := intǂ(i)
		v.setǂ(i, i_)
	}
	for i := 0; i < iterations; i++ {
		if i %100 == 0 {
			//fmt.Println(i)
		}
		v.reverseElementsǂ()
	}
}
func BenchmarkVectorReverse_string() {
	const iterations = 10000
	v := Vectorǂ{
		elementDataǂ:  make([]Any, iterations),
		elementCountǂ: iterations,
		dict_0: AnyǂDictǂ{},
	}
	for i := 0; i < iterations; i++ {
		v.setǂ(i, string_ǂ(strconv.Itoa(i)))
	}
	for i := 0; i < iterations; i++ {
		if i %100 == 0 {
			//fmt.Println(i)
		}
		v.reverseElementsǂ()
	}
}
func BenchmarkStaticVectorReverse_int() {
	const iterations = 10000
	v := Vectorǂ{
		elementDataǂ:  make([]Any, iterations),
		elementCountǂ: iterations,
		dict_0: AnyǂDictǂ{},
	}
	for i := 0; i < iterations; i++ {
		i_ := intǂ(i)
		v.setǂ(i, i_)
	}
	for i := 0; i < iterations; i++ {
		if i %100 == 0 {
			//fmt.Println(i)
		}
		Dummyǂ{}.externalReverseElementsǂ(v.dict_0, v)
	}
}
func BenchmarkStaticVectorReverse_string() {
	const iterations = 10000
	v := Vectorǂ{
		elementDataǂ:  make([]Any, iterations),
		elementCountǂ: iterations,
		dict_0: AnyǂDictǂ{},
	}
	for i := 0; i < iterations; i++ {
		v.setǂ(i, string_ǂ(strconv.Itoa(i)))
	}
	for i := 0; i < iterations; i++ {
		if i %100 == 0 {
			//fmt.Println(i)
		}
		Dummyǂ{}.externalReverseElementsǂ(v.dict_0, v)
	}
}
func main() {
	fmt.Println(1)
	BenchmarkVectorReverse_int()
	fmt.Println(2)
	BenchmarkVectorReverse_string()
	fmt.Println(3)
	BenchmarkStaticVectorReverse_int()
	fmt.Println(4)
	BenchmarkStaticVectorReverse_string()
	fmt.Println(5)
}
