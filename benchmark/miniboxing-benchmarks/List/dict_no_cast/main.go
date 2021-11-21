package main;

import "fmt"

type Any interface {};
type Anyǂ interface {};
type AnyǂDictǂ struct {  };
const TEST_SIZE = 3000000

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
type bool_ǂ struct {};
type bool_ǂ_meta struct {  };
func (this bool_ǂ_meta) tryCast(x Any) Any { _ = x.(bool_ǂ);
	return x;
};
func (this bool_ǂ_meta) assertEq(x _type_metadata) Any { return this;
};
func (thisǂ Dummyǂ) xǂ(aǂ Any, bǂ Any) Any { return Dummyǂ{};
};
func (thisǂ Dummyǂ) spec_x() spec_metadata_2 { return spec_metadata_2{Anyǂ_meta{}, Anyǂ_meta{}, Dummyǂ_meta{}};
};
type Listǂ interface { Containsǂ(xǂ Any) Any;  };
type ListǂDictǂ struct { Containsǂ func (receiver Any, xǂ Any) Any; _type _type_metadata };
type Listǂ_meta struct { _type_0 _type_metadata };


type Consǂ struct { headǂ Any; tailǂ Any; dict_0 AnyǂDictǂ };
type Consǂ_meta struct { _type_0 _type_metadata };

func (cǂ Consǂ) Containsǂ(xǂ Any) Any {
	var xs Listǂ = cǂ
	for {
		if xs_, ok := xs.(Consǂ); ok {
			if xs_.headǂ == xǂ {
				return true
			} else {
				xs = xs_.tailǂ.(Listǂ)
			}
		}
		if _, ok := xs.(Nilǂ); ok {
			return false
		}
	}
/*	_ = Consǂ_meta{cǂ.dict_0._type}.tryCast(cǂ.tailǂ);
	_ = Nilǂ_meta{cǂ.dict_0._type}.tryCast(cǂ.tailǂ);
	return bool_ǂ{};*/
};


type Nilǂ struct { dict_0 AnyǂDictǂ };
type Nilǂ_meta struct { _type_0 _type_metadata };


func (nǂ Nilǂ) Containsǂ(xǂ Any) Any { return false;
};

func (dǂ Dummyǂ) ListInsertǂ(dict_0 AnyǂDictǂ, xǂ Any) Any {
	var l Listǂ = Nilǂ{dict_0};
	for i := 0; i < TEST_SIZE; i++ {
		l = Consǂ{xǂ, l, dict_0};
	}
	return l
};

func (dǂ Dummyǂ) ListFindǂ(dict_0 AnyǂDictǂ, xǂ Any, lǂ Any) Any {
	b := true
	l := lǂ.(Listǂ)
	for i := 0; i < TEST_SIZE; i+=10000 {
		b = b != l.Containsǂ(i)
	}
	return b
};

func (thisǂ Dummyǂ) Main1ǂ() Any { dǂ := Dummyǂ{};
	fmt.Println("main1_1");
	t1ǂ := dǂ.ListInsertǂ(AnyǂDictǂ{}, 1);
	fmt.Println("main1_2");
	_ = dǂ.ListFindǂ(AnyǂDictǂ{}, 1, t1ǂ);
	fmt.Println("main1_3");
	return thisǂ;
};
func (thisǂ Dummyǂ) spec_Main1() spec_metadata_0 { return spec_metadata_0{Dummyǂ_meta{}};
};
func (thisǂ Dummyǂ) Main2ǂ() Any { dǂ := Dummyǂ{};
	fmt.Println("main2_1");
	t1ǂ := dǂ.ListInsertǂ(AnyǂDictǂ{}, 1);
	fmt.Println("main2_2");
	_ = dǂ.ListFindǂ(AnyǂDictǂ{}, 1, t1ǂ);
	fmt.Println("main2_3");
	return thisǂ;
};
func (thisǂ Dummyǂ) spec_Main2() spec_metadata_0 { return spec_metadata_0{Dummyǂ_meta{}};
};
func (thisǂ Dummyǂ) Main3ǂ() Any { dǂ := Dummyǂ{};
	fmt.Println("main3_1");

	t1ǂ := dǂ.ListInsertǂ(AnyǂDictǂ{}, 1);
	fmt.Println("main3_2");

	_ = dǂ.ListFindǂ(AnyǂDictǂ{}, 1, t1ǂ);
	fmt.Println("main3_3");

	return thisǂ;
};
func (thisǂ Dummyǂ) spec_Main3() spec_metadata_0 { return spec_metadata_0{Dummyǂ_meta{}};
};
type spec_metadata_3 struct { _type_0 _type_metadata; _type_1 _type_metadata; _type_2 _type_metadata; _type_3 _type_metadata };
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
func main() { _ = Dummyǂ{}.Main1ǂ().(Dummyǂ).Main2ǂ().(Dummyǂ).Main3ǂ() }
