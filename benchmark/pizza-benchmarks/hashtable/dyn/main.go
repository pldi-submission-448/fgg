package main;

import "strconv"

const ITERATIONS = 10000

type Any interface {};
type Anyǂ interface {};
type Dummyǂ struct {};
type Dummyǂ_meta struct {  };
type Intǂ struct {};
type Intǂ_meta struct {  };
type Stringǂ struct {};
type Stringǂ_meta struct {  };
type _mapǂ struct {};
type _mapǂ_meta struct {  };
type Hashtableǂ struct { _innerǂ map[Any]Any; };
type Hashtableǂ_meta struct { _type_0 _type_metadata;_type_1 _type_metadata };
func (thisǂ *Hashtableǂ) getǂ(keyǂ Any) Any { return thisǂ._innerǂ[keyǂ];
};
func (thisǂ *Hashtableǂ) setǂ(keyǂ Any, valueǂ Any) { thisǂ._innerǂ[keyǂ] = valueǂ;
};
func (thisǂ Dummyǂ) Mainǂ() Any {
	xǂ := Hashtableǂ{make(map[Any]Any)};
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
type spec_metadata_0 struct { _type_0 _type_metadata };
type spec_metadata_2 struct { _type_0 _type_metadata; _type_1 _type_metadata; _type_2 _type_metadata };
type _type_metadata interface { tryCast(x Any) Any; assertEq(x _type_metadata) Any };

func main() { _ = Dummyǂ{}.Mainǂ() }
