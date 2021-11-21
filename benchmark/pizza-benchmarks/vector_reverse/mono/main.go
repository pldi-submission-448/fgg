package main;

import (
	"fmt"
	"strconv"
)

type Dummyᐸᐳ struct {};
type intᐸᐳ int;
type string_ᐸᐳ string;
type Vectorᐸstring_ᐸᐳᐳ struct { elementData []string_ᐸᐳ; elementCount intᐸᐳ };
type Vectorᐸintᐸᐳᐳ struct { elementData []intᐸᐳ; elementCount intᐸᐳ };
func (v *Vectorᐸintᐸᐳᐳ) reverseElementsᐸᐳ() Dummyᐸᐳ {
	count := len(v.elementData)
	for i := 0; i < count / 2; i++ {
		e := v.elementData[count -i - 1]
		v.elementData[count -i -1] = v.elementData[i]
		v.elementData[i] = e
	}
	return Dummyᐸᐳ{};
};
func (v *Vectorᐸstring_ᐸᐳᐳ) reverseElementsᐸᐳ() Dummyᐸᐳ { _ = v.elementData;
	count := len(v.elementData)
	for i := 0; i < count / 2; i++ {
		e := v.elementData[count -i - 1]
		v.elementData[count -i -1] = v.elementData[i]
		v.elementData[i] = e
	}
	return Dummyᐸᐳ{};
};
func (v Vectorᐸintᐸᐳᐳ) reverseElements___Dummy__() Top { return v;
};
func (v Vectorᐸstring_ᐸᐳᐳ) reverseElements___Dummy__() Top { return v;
};
func (v *Vectorᐸintᐸᐳᐳ) getᐸᐳ(i intᐸᐳ) intᐸᐳ { return v.elementData[i];
};
func (v *Vectorᐸstring_ᐸᐳᐳ) getᐸᐳ(i intᐸᐳ) string_ᐸᐳ { return v.elementData[i];
};
func (v Vectorᐸintᐸᐳᐳ) get___int___int__() Top { return v;
};
func (v Vectorᐸstring_ᐸᐳᐳ) get___int___string___() Top { return v;
};
func (v *Vectorᐸintᐸᐳᐳ) setᐸᐳ(i intᐸᐳ, x intᐸᐳ) Dummyᐸᐳ {
	v.elementData[i] = x;
	return Dummyᐸᐳ{};
};
func (v *Vectorᐸstring_ᐸᐳᐳ) setᐸᐳ(i intᐸᐳ, x string_ᐸᐳ) Dummyᐸᐳ {
	v.elementData[i] = x;
	return Dummyᐸᐳ{};
};
func (v Vectorᐸintᐸᐳᐳ) set___int___int___Dummy__() Top { return v;
};
func (v Vectorᐸstring_ᐸᐳᐳ) set___int___string____Dummy__() Top { return v;
};
func (v Vectorᐸstring_ᐸᐳᐳ) sizeᐸᐳ() intᐸᐳ {
	return intᐸᐳ(len(v.elementData))
};
func (v Vectorᐸintᐸᐳᐳ) sizeᐸᐳ() intᐸᐳ {
	return intᐸᐳ(len(v.elementData))
};
func (v Vectorᐸintᐸᐳᐳ) size___int__() Top { return v;
};
func (v Vectorᐸstring_ᐸᐳᐳ) size___int__() Top { return v;
};
func (this Dummyᐸᐳ) externalReverseElementsᐸintᐸᐳᐳ(v Vectorᐸintᐸᐳᐳ) Dummyᐸᐳ {
	for i := 0; i < int(v.sizeᐸᐳ()) / 2; i++ {
		e := v.getᐸᐳ(intᐸᐳ(int(v.sizeᐸᐳ()) - i -1))
		v.setᐸᐳ(intᐸᐳ(int(v.sizeᐸᐳ()) - i -1), v.getᐸᐳ(intᐸᐳ(i)))
		v.setᐸᐳ(intᐸᐳ(i), e)
	}
	return Dummyᐸᐳ{};
};
func (this Dummyᐸᐳ) externalReverseElementsᐸstring_ᐸᐳᐳ(a Vectorᐸstring_ᐸᐳᐳ) Dummyᐸᐳ {
	for i := 0; i < int(a.sizeᐸᐳ()) / 2; i++ {
		e := a.getᐸᐳ(intᐸᐳ(int(a.sizeᐸᐳ()) - i -1))
		a.setᐸᐳ(intᐸᐳ(int(a.sizeᐸᐳ()) - i -1), a.getᐸᐳ(intᐸᐳ(i)))
		a.setᐸᐳ(intᐸᐳ(i), e)
	}
	return Dummyᐸᐳ{};
};
func (this Dummyᐸᐳ) externalReverseElements__β1_Any____Vector_β1__Dummy__() Top { return this;
};
func (this Dummyᐸᐳ) x___Any___Any___Dummy__() Top { return this;
};
func BenchmarkVectorReverse_int() {
	const iterations = 10000
	v := Vectorᐸintᐸᐳᐳ{
		elementData:  make([]intᐸᐳ, iterations),
		elementCount: iterations,
	}
	for i := 0; i < iterations; i++ {
		i_ := intᐸᐳ(i)
		v.setᐸᐳ(i_, i_)
	}
	for i := 0; i < iterations; i++ {
		if i %100 == 0 {
			//fmt.Println(i)
		}
		v.reverseElementsᐸᐳ()
	}
}

func BenchmarkVectorReverse_string() {
	const iterations = 10000
	v := Vectorᐸstring_ᐸᐳᐳ{
		elementData:  make([]string_ᐸᐳ, iterations),
		elementCount: iterations,
	}
	for i := 0; i < iterations; i++ {
		i_ := intᐸᐳ(i)
		v.setᐸᐳ(i_, string_ᐸᐳ(strconv.Itoa(i)))
	}
	for i := 0; i < iterations; i++ {
		if i %100 == 0 {
			//fmt.Println(i)
		}
		v.reverseElementsᐸᐳ()
	}
}

func BenchmarkStaticVectorReverse_int() {
	const iterations = 10000
	v := Vectorᐸintᐸᐳᐳ{
		elementData:  make([]intᐸᐳ, iterations),
		elementCount: iterations,
	}
	for i := 0; i < iterations; i++ {
		i_ := intᐸᐳ(i)
		v.setᐸᐳ(i_, i_)
	}
	for i := 0; i < iterations; i++ {
		if i %100 == 0 {
			//fmt.Println(i)
		}
		Dummyᐸᐳ{}.externalReverseElementsᐸintᐸᐳᐳ(v)
	}
}

func BenchmarkStaticVectorReverse_string() {
	const iterations = 10000
	v := Vectorᐸstring_ᐸᐳᐳ{
		elementData:  make([]string_ᐸᐳ, iterations),
		elementCount: iterations,
	}
	for i := 0; i < iterations; i++ {
		i_ := intᐸᐳ(i)
		v.setᐸᐳ(i_, string_ᐸᐳ(strconv.Itoa(i)))
	}
	for i := 0; i < iterations; i++ {
		if i %100 == 0 {
			//fmt.Println(i)
		}
		Dummyᐸᐳ{}.externalReverseElementsᐸstring_ᐸᐳᐳ(v)
	}
}


func (this Dummyᐸᐳ) Main1___Dummy__() Top { return this;
};
func (this Dummyᐸᐳ) Main2___Dummy__() Top { return this;
};
type Top interface {};
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
