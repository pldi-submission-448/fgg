package main;

import "fmt"

const TEST_SIZE = 3000000//100000000 //2000000, 3000000

type Any interface {};
type intǂ int8;
type intǂ_meta struct {  };
func (this intǂ_meta) tryCast(x Any) Any { _ = x.(intǂ);
	return x;
};
func (this intǂ_meta) assertEq(x _type_metadata) Any { return this;
};
type longǂ int32;
type longǂ_meta struct {  };
func (this longǂ_meta) tryCast(x Any) Any { _ = x.(longǂ);
	return x;
};
func (this longǂ_meta) assertEq(x _type_metadata) Any { return this;
};
type doubleǂ int64;
type doubleǂ_meta struct {  };
func (this doubleǂ_meta) tryCast(x Any) Any { _ = x.(doubleǂ);
	return x;
};
func (this doubleǂ_meta) assertEq(x _type_metadata) Any { return this;
};
type boolǂ struct {};
type boolǂ_meta struct {  };
func (this boolǂ_meta) tryCast(x Any) Any { _ = x.(boolǂ);
	return x;
};
func (this boolǂ_meta) assertEq(x _type_metadata) Any { return this;
};
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
type ResizableArrayǂ struct {
	sizeǂ Any;
	elemCountǂ Any;
	arrayǂ []Any;
	newArrayǂ []Any;
	dict_0 AnyǂDictǂ
};
type ResizableArrayǂ_meta struct { _type_0 _type_metadata };
func (this ResizableArrayǂ_meta) tryCast(x Any) Any { _ = x.(ResizableArrayǂ);
	return x;
};
func (this ResizableArrayǂ_meta) assertEq(x _type_metadata) Any { this._type_0.assertEq(x.(ResizableArrayǂ_meta)._type_0);
	return this;
};

func (raǂ *ResizableArrayǂ) extendǂ() Any {
	var size int = int(raǂ.sizeǂ.(int))
	if raǂ.elemCountǂ == size {
		pos := 0
		raǂ.newArrayǂ = make([]Any/*T*/, 2 * size)
		for pos < size {
			raǂ.newArrayǂ[pos] = raǂ.arrayǂ[pos]
			pos += 1
		}
		raǂ.arrayǂ = raǂ.newArrayǂ
		raǂ.sizeǂ = size * 2
	}
	return Dummyǂ{};
};

func (raǂ ResizableArrayǂ) spec_extend() spec_metadata_0 { return spec_metadata_0{Dummyǂ_meta{}};
};
func (raǂ *ResizableArrayǂ) addǂ(elemǂ Any) Any {
	_ = raǂ.extendǂ();
	cnt := int(raǂ.elemCountǂ.(int))
	raǂ.arrayǂ[cnt] = elemǂ
	raǂ.elemCountǂ = cnt + 1
	return Dummyǂ{};
};

func (raǂ *ResizableArrayǂ) reverseǂ() Any {
	pos := 0
	cnt := int(raǂ.elemCountǂ.(int))
	for pos * 2 < cnt {
		tmp1 := raǂ.arrayǂ[pos]
		tmp2 := raǂ.arrayǂ[cnt - pos - 1]
		raǂ.arrayǂ[pos] = tmp2
		raǂ.arrayǂ[cnt - pos - 1] = tmp1
		pos += 1
	}
	return raǂ
};

func (raǂ ResizableArrayǂ) spec_reverse() spec_metadata_0 { return spec_metadata_0{Dummyǂ_meta{}};
};
func (raǂ *ResizableArrayǂ) containsǂ(elemǂ Any) Any {
	pos := 0
	cnt := int(raǂ.elemCountǂ.(int))
	for pos < cnt {
		if raǂ.arrayǂ[pos] == elemǂ {
			return true
		}
		pos += 1
	}
	return false
};

func (thisǂ Dummyǂ) arrayInsertǂ(dict_0 AnyǂDictǂ, elemǂ Any) Any {
	a := ResizableArrayǂ{
		sizeǂ:      4,
		elemCountǂ: 0,
		arrayǂ:     make([]Any, 4),
		newArrayǂ:  nil,
		dict_0: dict_0,
	}
	for i := 0; i < TEST_SIZE; i++ {
		a.addǂ(elemǂ)
	}
	return a
};
func (thisǂ Dummyǂ) spec_arrayInsert() spec_metadata_2 { return spec_metadata_2{Anyǂ_meta{}, param_index_0{}, ResizableArrayǂ_meta{param_index_0{}}};
};
func (thisǂ Dummyǂ) arrayReverseǂ(dict_0 AnyǂDictǂ, aǂ Any) Any {
	var t = aǂ.(ResizableArrayǂ)
	_ = (&t).reverseǂ();
	return aǂ;
};
func (thisǂ Dummyǂ) spec_arrayReverse() spec_metadata_2 { return spec_metadata_2{Anyǂ_meta{}, ResizableArrayǂ_meta{param_index_0{}}, ResizableArrayǂ_meta{param_index_0{}}};
};
func (thisǂ Dummyǂ) arrayFindǂ(dict_0 AnyǂDictǂ, xǂ Any, elemǂ Any) Any {
	i := 0
	b := true
	x := xǂ.(ResizableArrayǂ)
	for i < TEST_SIZE {
		b = b != x.containsǂ(i).(bool)
		i += 10000
	}
	return b
};
func (thisǂ Dummyǂ) spec_arrayFind() spec_metadata_3 { return spec_metadata_3{Anyǂ_meta{}, ResizableArrayǂ_meta{param_index_0{}}, param_index_0{}, boolǂ_meta{}};
};
func (thisǂ Dummyǂ) Mainǂ(dict_0 AnyǂDictǂ, elemǂ Any) Any {
	fmt.Println(1)
	aǂ := thisǂ.arrayInsertǂ(dict_0, elemǂ);
	fmt.Println(2)
	a2ǂ := Dummyǂ{}.arrayReverseǂ(dict_0, aǂ);
	fmt.Println(3)
	_ = Dummyǂ{}.arrayFindǂ(dict_0, a2ǂ, elemǂ);
	fmt.Println(4)
	return thisǂ;
};
func (thisǂ Dummyǂ) spec_Main() spec_metadata_2 { return spec_metadata_2{Anyǂ_meta{}, param_index_0{}, Dummyǂ_meta{}};
};
type spec_metadata_0 struct { _type_0 _type_metadata };
type spec_metadata_1 struct { _type_0 _type_metadata; _type_1 _type_metadata };
type spec_metadata_2 struct { _type_0 _type_metadata; _type_1 _type_metadata; _type_2 _type_metadata };
type spec_metadata_3 struct { _type_0 _type_metadata; _type_1 _type_metadata; _type_2 _type_metadata; _type_3 _type_metadata };
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
func main() {
	_ = Dummyǂ{}.Mainǂ(AnyǂDictǂ{}, intǂ(4)).
		(Dummyǂ).Mainǂ(AnyǂDictǂ{}, 4).
		(Dummyǂ).Mainǂ(AnyǂDictǂ{}, 4)
}
