package main;

import "fmt"

const TEST_SIZE = 3000000
type Any interface {};
type Anyǂ interface {};
type Dummyǂ struct {};
type Dummyǂ_meta struct {  };
type intǂ int8;
type intǂ_meta struct {  };
type longǂ int32;
type longǂ_meta struct {  };
type doubleǂ int64;
type doubleǂ_meta struct {  };
type bool_ǂ struct {};
type bool_ǂ_meta struct {  };
func (thisǂ Dummyǂ) xǂ(aǂ Any, bǂ Any) Any { return Dummyǂ{};
};
type Listǂ interface { Containsǂ(xǂ Any) Any };
type Consǂ struct { headǂ Any; tailǂ Any };
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
};
type Nilǂ struct {};
type Nilǂ_meta struct { _type_0 _type_metadata };
func (nǂ Nilǂ) Containsǂ(xǂ Any) Any { return false;
};
func (dǂ Dummyǂ) ListInsertǂ(xǂ Any) Any {
	var l Listǂ = Nilǂ{};
	for i := 0; i < TEST_SIZE; i++ {
		l = Consǂ{xǂ, l};
	}
	return l
};
func (dǂ Dummyǂ) ListFindǂ(xǂ Any, lǂ Any) Any {
	b := true
	l := lǂ.(Listǂ)
	for i := 0; i < TEST_SIZE; i+=10000 {
		b = b != l.Containsǂ(i)
	}
	return b
};

func (thisǂ Dummyǂ) Main1ǂ() Any {
	dǂ := Dummyǂ{};
	fmt.Println("main1_1");
	t1ǂ := dǂ.ListInsertǂ(1);
	fmt.Println("main1_2");
	_ = dǂ.ListFindǂ(1, t1ǂ);
	fmt.Println("main1_3");
	return thisǂ;
};

func (thisǂ Dummyǂ) Main2ǂ() Any {
	dǂ := Dummyǂ{};
	fmt.Println("main2_1");
	t1ǂ := dǂ.ListInsertǂ(1);
	fmt.Println("main2_2");
	_ = dǂ.ListFindǂ(1, t1ǂ);
	fmt.Println("main2_3");
	return thisǂ;
};
func (thisǂ Dummyǂ) Main3ǂ() Any {
	dǂ := Dummyǂ{};
	fmt.Println("main3_1");

	t1ǂ := dǂ.ListInsertǂ(1);
	fmt.Println("main3_2");

	_ = dǂ.ListFindǂ(1, t1ǂ);
	fmt.Println("main3_3");

	return thisǂ;
};
type spec_metadata_0 struct { _type_0 _type_metadata };
type spec_metadata_1 struct { _type_0 _type_metadata; _type_1 _type_metadata };
type _type_metadata interface { tryCast(x Any) Any; assertEq(x _type_metadata) Any };

func main() { _ = Dummyǂ{}.Main1ǂ().(Dummyǂ).Main2ǂ().(Dummyǂ).Main3ǂ() }
