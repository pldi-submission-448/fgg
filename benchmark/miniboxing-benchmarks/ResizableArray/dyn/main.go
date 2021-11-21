package main;

import "fmt"

const TEST_SIZE = 3000000//100000000 //2000000, 3000000

type Any interface {};
type intǂ int8;
type intǂ_meta struct {  };
type longǂ int32;
type longǂ_meta struct {  };
type doubleǂ int64;
type doubleǂ_meta struct {  };
type boolǂ struct {};
type boolǂ_meta struct {  };
type Anyǂ interface {};
type Dummyǂ struct {};
type Dummyǂ_meta struct {  };
type ResizableArrayǂ struct { sizeǂ Any;
elemCountǂ Any;
arrayǂ []Any;
	newArrayǂ []Any; };
type ResizableArrayǂ_meta struct { _type_0 _type_metadata };
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
func (thisǂ Dummyǂ) arrayInsertǂ(elemǂ Any) Any {
	a := ResizableArrayǂ{
		sizeǂ:      4,
		elemCountǂ: 0,
		arrayǂ:     make([]Any, 4),
		newArrayǂ:  nil,
	}
	for i := 0; i < TEST_SIZE; i++ {
		a.addǂ(elemǂ)
	}
	return a
};
func (thisǂ Dummyǂ) arrayReverseǂ(aǂ Any) Any {
	var t = aǂ.(ResizableArrayǂ)
	_ = (&t).reverseǂ();
	return aǂ;
};
func (thisǂ Dummyǂ) arrayFindǂ(xǂ Any, elemǂ Any) Any {
	i := 0
	b := true
	x := xǂ.(ResizableArrayǂ)
	for i < TEST_SIZE {
		b = b != x.containsǂ(i).(bool)
		i += 10000
	}
	return b
};
func (thisǂ Dummyǂ) Mainǂ(elemǂ Any) Any {
	fmt.Println(1)
	aǂ := thisǂ.arrayInsertǂ(elemǂ);
	fmt.Println(2)
	a2ǂ := Dummyǂ{}.arrayReverseǂ(aǂ);
	fmt.Println(3)
	_ = Dummyǂ{}.arrayFindǂ(a2ǂ, elemǂ);
	fmt.Println(4)
	return thisǂ;
};
type spec_metadata_0 struct { _type_0 _type_metadata };
type spec_metadata_1 struct { _type_0 _type_metadata; _type_1 _type_metadata };
type _type_metadata interface { tryCast(x Any) Any; assertEq(x _type_metadata) Any };

func main() {
	_ = Dummyǂ{}.Mainǂ(intǂ(4)).
	(Dummyǂ).Mainǂ(4).
	(Dummyǂ).Mainǂ(4)
}
