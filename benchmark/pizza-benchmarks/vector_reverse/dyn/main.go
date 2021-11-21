package main;

import (
	"fmt"
	"strconv"
	"unsafe"
)

type Any interface {};
type Anyǂ interface {};
type Dummyǂ struct {};
type Dummyǂ_meta struct {  };
type intǂ int;
type intǂ_meta struct {  };
type string_ǂ string;
type string_ǂ_meta struct {  };
type Vectorǂ struct { elementDataǂ []Any; elementCountǂ Any };
type Vectorǂ_meta struct { _type_0 _type_metadata };
func (vǂ *Vectorǂ) reverseElementsǂ() Any {
	count := len(vǂ.elementDataǂ)
	for i := 0; i < count / 2; i++ {
		e := vǂ.elementDataǂ[count -i - 1]
		vǂ.elementDataǂ[count -i -1] = vǂ.elementDataǂ[i]
		vǂ.elementDataǂ[i] = e
	}
	return Dummyǂ{};
};
func (vǂ *Vectorǂ) getǂ(iǂ unsafe.Pointer) Any { return vǂ.elementDataǂ[*((*int)(iǂ))];
};
func (vǂ *Vectorǂ) setǂ(iǂ unsafe.Pointer, xǂ Any) Any {	vǂ.elementDataǂ[*((*int)(iǂ))] = xǂ;
	return Dummyǂ{};
};
func (vǂ Vectorǂ) sizeǂ() int { 	return len(vǂ.elementDataǂ)

};
func (thisǂ Dummyǂ) externalReverseElementsǂ(aǂ Any) Any {	a := aǂ.(Vectorǂ)
	for i := 0; i < a.sizeǂ() / 2; i++ {
		ind1 := a.sizeǂ() - i - 1
		e := a.getǂ(unsafe.Pointer(&ind1))
		a.setǂ(unsafe.Pointer(&ind1), a.getǂ(unsafe.Pointer(&i)))
		a.setǂ(unsafe.Pointer(&i), e)
	}
	return Dummyǂ{};
};
func (thisǂ Dummyǂ) xǂ(aǂ Any, bǂ Any) Any { return Dummyǂ{};
};

type spec_metadata_0 struct { _type_0 _type_metadata };
type spec_metadata_1 struct { _type_0 _type_metadata; _type_1 _type_metadata };
type _type_metadata interface { tryCast(x Any) Any; assertEq(x _type_metadata) Any };

func BenchmarkVectorReverse_int() {
	const iterations = 10000
	v := Vectorǂ{
		elementDataǂ:  make([]Any, iterations),
		elementCountǂ: iterations,
	}
	for i := 0; i < iterations; i++ {
		i_ := intǂ(i)
		v.setǂ(unsafe.Pointer(&i), i_)
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
	}
	for i := 0; i < iterations; i++ {
		v.setǂ(unsafe.Pointer(&i), string_ǂ(strconv.Itoa(i)))
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
	}
	for i := 0; i < iterations; i++ {
		i_ := intǂ(i)
		v.setǂ(unsafe.Pointer(&i), i_)
	}
	for i := 0; i < iterations; i++ {
		if i %100 == 0 {
			//fmt.Println(i)
		}
		Dummyǂ{}.externalReverseElementsǂ(v)
	}
}
func BenchmarkStaticVectorReverse_string() {
	const iterations = 10000
	v := Vectorǂ{
		elementDataǂ:  make([]Any, iterations),
		elementCountǂ: iterations,
	}
	for i := 0; i < iterations; i++ {
		v.setǂ(unsafe.Pointer(&i), string_ǂ(strconv.Itoa(i)))
	}
	for i := 0; i < iterations; i++ {
		if i %100 == 0 {
			//fmt.Println(i)
		}
		Dummyǂ{}.externalReverseElementsǂ(v)
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