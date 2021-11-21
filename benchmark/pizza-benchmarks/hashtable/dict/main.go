package main;

import "strconv"
const ITERATIONS = 10000

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
type Intǂ struct {};
type Intǂ_meta struct {  };
func (this Intǂ_meta) tryCast(x Any) Any { _ = x.(Intǂ);
	return x;
};
func (this Intǂ_meta) assertEq(x _type_metadata) Any { return this;
};
type Stringǂ struct {};
type Stringǂ_meta struct {  };
func (this Stringǂ_meta) tryCast(x Any) Any { _ = x.(Stringǂ);
	return x;
};
func (this Stringǂ_meta) assertEq(x _type_metadata) Any { return this;
};
type _mapǂ struct {};
type _mapǂ_meta struct {  };
func (this _mapǂ_meta) tryCast(x Any) Any { _ = x.(_mapǂ);
	return x;
};
func (this _mapǂ_meta) assertEq(x _type_metadata) Any { return this;
};
type Hashtableǂ struct { _innerǂ map[Any]Any; dict_0 AnyǂDictǂ; dict_1 AnyǂDictǂ };
type Hashtableǂ_meta struct { _type_0 _type_metadata;_type_1 _type_metadata };
func (this Hashtableǂ_meta) tryCast(x Any) Any { _ = x.(Hashtableǂ);
	this._type_0.assertEq(x.(Hashtableǂ).dict_0._type);
	this._type_1.assertEq(x.(Hashtableǂ).dict_1._type);
	return x;
};
func (this Hashtableǂ_meta) assertEq(x _type_metadata) Any { this._type_0.assertEq(x.(Hashtableǂ_meta)._type_0);
	this._type_1.assertEq(x.(Hashtableǂ_meta)._type_1);
	return this;
};
func (thisǂ *Hashtableǂ) getǂ(keyǂ Any) Any { return thisǂ._innerǂ[keyǂ];
};
func (thisǂ Hashtableǂ) spec_get() spec_metadata_1 { return spec_metadata_1{thisǂ.dict_0._type, thisǂ.dict_1._type};
};
func (thisǂ *Hashtableǂ) setǂ(keyǂ Any, valueǂ Any) { thisǂ._innerǂ[keyǂ] = valueǂ;
};
func (thisǂ Hashtableǂ) spec_set() spec_metadata_2 { return spec_metadata_2{thisǂ.dict_0._type, thisǂ.dict_1._type, Dummyǂ_meta{}};
};
func (thisǂ Dummyǂ) Mainǂ() Any {
	xǂ := Hashtableǂ{make(map[Any]Any), AnyǂDictǂ{Intǂ_meta{}}, AnyǂDictǂ{Stringǂ_meta{}}};
	for i := 0; i < 1000; i++ {
		xǂ.setǂ(i, "A" + strconv.Itoa(i))
	}
	for j := 0; j < ITERATIONS; j++ {
		for i := 0; i < 1000; i++ {
			xǂ.getǂ(i)
		}
	}
	return thisǂ;
};
func (thisǂ Dummyǂ) spec_Main() spec_metadata_0 { return spec_metadata_0{Dummyǂ_meta{}};
};
type spec_metadata_1 struct { _type_0 _type_metadata; _type_1 _type_metadata };
type spec_metadata_0 struct { _type_0 _type_metadata };
type spec_metadata_2 struct { _type_0 _type_metadata; _type_1 _type_metadata; _type_2 _type_metadata };
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
func main() { _ = Dummyǂ{}.Mainǂ() }
