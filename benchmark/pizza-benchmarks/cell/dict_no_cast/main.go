package main;

import "fmt"

const iterations = 1000000;
//const iterations = 50000000;

type Any interface {};
type Int1ǂ int;
type Int1ǂ_meta struct {  };
func (this Int1ǂ_meta) tryCast(x Any) Any { _ = x.(Int1ǂ);
	return x;
};
func (this Int1ǂ_meta) assertEq(x _type_metadata) Any { return this;
};
type Int2ǂ struct {};
type Int2ǂ_meta struct {  };
func (this Int2ǂ_meta) tryCast(x Any) Any { _ = x.(Int2ǂ);
	return x;
};
func (this Int2ǂ_meta) assertEq(x _type_metadata) Any { return this;
};
type Anyǂ interface {};
type AnyǂDictǂ struct { _type _type_metadata };
type Anyǂ_meta struct {  };
func (this Anyǂ_meta) tryCast(x Any) Any { x_ := x.(Anyǂ);
	return x_;
};
func (this Anyǂ_meta) assertEq(x _type_metadata) Any { x_ := x.(Anyǂ_meta);
	return x_;
};
type Cellǂ struct { elemǂ Any };
type Cellǂ_meta struct { _type_0 _type_metadata };

func (thisǂ *Cellǂ) addǂ(elemǂ Any) Any { return Cellǂ{elemǂ};
};

func (thisǂ *Cellǂ) getǂ() Any { return thisǂ.elemǂ;
};

type Dummyǂ struct {};
type Dummyǂ_meta struct {  };
func (this Dummyǂ_meta) tryCast(x Any) Any { _ = x.(Dummyǂ);
	return x;
};
func (this Dummyǂ_meta) assertEq(x _type_metadata) Any { return this;
};
func (thisǂ Dummyǂ) CellTestǂ( elemǂ Any) Any {
	cǂ := Cellǂ{42};
	cǂ.addǂ(42)
	for i := 0; i < iterations; i++ {
		cǂ.getǂ();
	}
	return thisǂ
};
func (thisǂ Dummyǂ) spec_CellTest() spec_metadata_2 { return spec_metadata_2{Anyǂ_meta{}, param_index_0{}, Dummyǂ_meta{}};
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
func main() {
	fmt.Println(1)
	_ = Dummyǂ{}.CellTestǂ( 42)
	fmt.Println(2)
}
