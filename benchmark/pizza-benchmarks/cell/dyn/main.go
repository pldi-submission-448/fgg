package main;

import "fmt"

const iterations = 1000000;
type Any interface {};
type Int1ǂ struct {};
type Int1ǂ_meta struct {  };
type Int2ǂ struct {};
type Int2ǂ_meta struct {  };
type Anyǂ interface {};
type Cellǂ struct { elemǂ Any };
type Cellǂ_meta struct { _type_0 _type_metadata };
func (thisǂ *Cellǂ) addǂ(elemǂ Any) Any {
	return Cellǂ{elemǂ};
};
func (thisǂ *Cellǂ) getǂ() Any {
	return thisǂ.elemǂ;
};

type Dummyǂ struct {};
type Dummyǂ_meta struct {  };
func (thisǂ Dummyǂ) CellTestǂ(elemǂ Any) Any {
	cǂ := Cellǂ{42};
	cǂ.addǂ(42)
	for i := 0; i < iterations; i++ {
		cǂ.getǂ();
	}
	return thisǂ
};
type spec_metadata_0 struct { _type_0 _type_metadata };
type spec_metadata_1 struct { _type_0 _type_metadata; _type_1 _type_metadata };
type _type_metadata interface { tryCast(x Any) Any; assertEq(x _type_metadata) Any };

func main() {
	fmt.Println(1)
	_ = Dummyǂ{}.CellTestǂ(42)
	fmt.Println(2)
}
