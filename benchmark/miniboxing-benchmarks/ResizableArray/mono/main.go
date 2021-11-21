package main;

import "fmt"

type intᐸᐳ int8;
type longᐸᐳ int32;
type doubleᐸᐳ int64;
type boolᐸᐳ bool;
type Dummyᐸᐳ struct {};
const TEST_SIZE = 3000000//100000000 //2000000, 3000000

type ResizableArrayᐸintᐸᐳᐳ struct {
	size int;
	elemCount int;
	array []intᐸᐳ;
	newArray []intᐸᐳ
};

type ResizableArrayᐸlongᐸᐳᐳ struct {
	size int;
	elemCount int;
	array []longᐸᐳ;
	newArray []longᐸᐳ
};

type ResizableArrayᐸdoubleᐸᐳᐳ struct {
	size int;
	elemCount int;
	array []doubleᐸᐳ;
	newArray []doubleᐸᐳ
};

func (ra *ResizableArrayᐸdoubleᐸᐳᐳ) extendᐸᐳ() Dummyᐸᐳ {
	var size int = ra.size
	if ra.elemCount == size {
		pos := 0
		ra.newArray = make([]doubleᐸᐳ, 2 * size)
		for pos < size {
			ra.newArray[pos] = ra.array[pos]
			pos += 1
		}
		ra.array = ra.newArray
		ra.size = size * 2
	}
	return Dummyᐸᐳ{};
};

func (ra *ResizableArrayᐸlongᐸᐳᐳ) extendᐸᐳ() Dummyᐸᐳ { _ = ra.elemCount;
	var size int = ra.size
	if int(ra.elemCount) == size {
		pos := 0
		ra.newArray = make([]longᐸᐳ, 2 * size)
		for pos < size {
			ra.newArray[pos] = ra.array[pos]
			pos += 1
		}
		ra.array = ra.newArray
		ra.size = size * 2
	}
	return Dummyᐸᐳ{};
};
func (ra *ResizableArrayᐸintᐸᐳᐳ) extendᐸᐳ() Dummyᐸᐳ { _ = ra.elemCount;
	var size int = ra.size
	if int(ra.elemCount) == size {
		pos := 0
		ra.newArray = make([]intᐸᐳ, 2 * size)
		for pos < size {
			ra.newArray[pos] = ra.array[pos]
			pos += 1
		}
		ra.array = ra.newArray
		ra.size = size * 2
	}
	return Dummyᐸᐳ{};
};
func (ra ResizableArrayᐸlongᐸᐳᐳ) extend___Dummy__() Top { return ra;
};
func (ra ResizableArrayᐸdoubleᐸᐳᐳ) extend___Dummy__() Top { return ra;
};
func (ra ResizableArrayᐸintᐸᐳᐳ) extend___Dummy__() Top { return ra;
};
func (ra *ResizableArrayᐸdoubleᐸᐳᐳ) addᐸᐳ(elem doubleᐸᐳ) Dummyᐸᐳ {
	_ = ra.extendᐸᐳ();
	cnt := int(ra.elemCount)
	ra.array[cnt] = elem
	ra.elemCount = cnt + 1
	return Dummyᐸᐳ{};
};
func (ra *ResizableArrayᐸintᐸᐳᐳ) addᐸᐳ(elem intᐸᐳ) Dummyᐸᐳ {
	_ = ra.extendᐸᐳ();
	cnt := int(ra.elemCount)
	ra.array[cnt] = elem
	ra.elemCount = cnt + 1
	return Dummyᐸᐳ{};
};
func (ra *ResizableArrayᐸlongᐸᐳᐳ) addᐸᐳ(elem longᐸᐳ) Dummyᐸᐳ {
	_ = ra.extendᐸᐳ();
	cnt := int(ra.elemCount)
	ra.array[cnt] = elem
	ra.elemCount = cnt + 1
	return Dummyᐸᐳ{};
};
func (ra ResizableArrayᐸlongᐸᐳᐳ) add___long___Dummy__() Top { return ra;
};
func (ra ResizableArrayᐸdoubleᐸᐳᐳ) add___double___Dummy__() Top { return ra;
};
func (ra ResizableArrayᐸintᐸᐳᐳ) add___int___Dummy__() Top { return ra;
};
func (ra *ResizableArrayᐸlongᐸᐳᐳ) reverseᐸᐳ() Dummyᐸᐳ {
	pos := 0
	cnt := int(ra.elemCount)
	for pos * 2 < cnt {
		tmp1 := ra.array[pos]
		tmp2 := ra.array[cnt - pos - 1]
		ra.array[pos] = tmp2
		ra.array[cnt - pos - 1] = tmp1
		pos += 1
	}
	return Dummyᐸᐳ{};
};
func (ra *ResizableArrayᐸdoubleᐸᐳᐳ) reverseᐸᐳ() Dummyᐸᐳ {
	pos := 0
	cnt := int(ra.elemCount)
	for pos * 2 < cnt {
		tmp1 := ra.array[pos]
		tmp2 := ra.array[cnt - pos - 1]
		ra.array[pos] = tmp2
		ra.array[cnt - pos - 1] = tmp1
		pos += 1
	}
	return Dummyᐸᐳ{};
};

func (ra *ResizableArrayᐸintᐸᐳᐳ) reverseᐸᐳ() Dummyᐸᐳ {
	pos := 0
	cnt := int(ra.elemCount)
	for pos * 2 < cnt {
		tmp1 := ra.array[pos]
		tmp2 := ra.array[cnt - pos - 1]
		ra.array[pos] = tmp2
		ra.array[cnt - pos - 1] = tmp1
		pos += 1
	}
	return Dummyᐸᐳ{};
};
func (ra ResizableArrayᐸlongᐸᐳᐳ) reverse___Dummy__() Top { return ra;
};
func (ra ResizableArrayᐸdoubleᐸᐳᐳ) reverse___Dummy__() Top { return ra;
};
func (ra ResizableArrayᐸintᐸᐳᐳ) reverse___Dummy__() Top { return ra;
};
func (ra *ResizableArrayᐸlongᐸᐳᐳ) containsᐸᐳ(elem longᐸᐳ) boolᐸᐳ {
	pos := 0
	cnt := int(ra.elemCount)
	for pos < cnt {
		if ra.array[pos] == elem {
			return true
		}
		pos += 1
	}
	return false
};

func (ra *ResizableArrayᐸintᐸᐳᐳ) containsᐸᐳ(elem intᐸᐳ) boolᐸᐳ {
	pos := 0
	cnt := int(ra.elemCount)
	for pos < cnt {
		if ra.array[pos] == elem {
			return true
		}
		pos += 1
	}
	return false
};
func (ra *ResizableArrayᐸdoubleᐸᐳᐳ) containsᐸᐳ(elem doubleᐸᐳ) boolᐸᐳ {
	pos := 0
	cnt := int(ra.elemCount)
	for pos < cnt {
		if ra.array[pos] == elem {
			return true
		}
		pos += 1
	}
	return false
};

func (ra ResizableArrayᐸdoubleᐸᐳᐳ) contains___double___bool__() Top { return ra;
};
func (ra ResizableArrayᐸintᐸᐳᐳ) contains___int___bool__() Top { return ra;
};
func (ra ResizableArrayᐸlongᐸᐳᐳ) contains___long___bool__() Top { return ra;
};
func (this Dummyᐸᐳ) arrayInsertᐸintᐸᐳᐳ(elem intᐸᐳ) ResizableArrayᐸintᐸᐳᐳ {
	a := ResizableArrayᐸintᐸᐳᐳ{
		size:      4,
		elemCount: 0,
		array:     make([]intᐸᐳ, 4),
		newArray:  nil,
	}
	for i := 0; i < TEST_SIZE; i++ {
		a.addᐸᐳ(elem)
	}
	return a
};
func (this Dummyᐸᐳ) arrayInsertᐸdoubleᐸᐳᐳ(elem doubleᐸᐳ) ResizableArrayᐸdoubleᐸᐳᐳ {
	a := ResizableArrayᐸdoubleᐸᐳᐳ{
		size:      4,
		elemCount: 0,
		array:     make([]doubleᐸᐳ, 4),
		newArray:  nil,
	}
	for i := 0; i < TEST_SIZE; i++ {
		a.addᐸᐳ(elem)
	}
	return a
};

func (this Dummyᐸᐳ) arrayInsertᐸlongᐸᐳᐳ(elem longᐸᐳ) ResizableArrayᐸlongᐸᐳᐳ {
	a := ResizableArrayᐸlongᐸᐳᐳ{
		size:      4,
		elemCount: 0,
		array:     make([]longᐸᐳ, 4),
		newArray:  nil,
	}
	for i := 0; i < TEST_SIZE; i++ {
		a.addᐸᐳ(elem)
	}
	return a
};

func (this Dummyᐸᐳ) arrayInsert__β1_Any____β1_ResizableArray_β1_() Top { return this;
};
func (this Dummyᐸᐳ) arrayReverseᐸdoubleᐸᐳᐳ(a ResizableArrayᐸdoubleᐸᐳᐳ) ResizableArrayᐸdoubleᐸᐳᐳ {
	_ = a.reverseᐸᐳ();
	return a;
};
func (this Dummyᐸᐳ) arrayReverseᐸintᐸᐳᐳ(a ResizableArrayᐸintᐸᐳᐳ) ResizableArrayᐸintᐸᐳᐳ {
	_ = a.reverseᐸᐳ();
	return a;
};
func (this Dummyᐸᐳ) arrayReverseᐸlongᐸᐳᐳ(a ResizableArrayᐸlongᐸᐳᐳ) ResizableArrayᐸlongᐸᐳᐳ {
	_ = a.reverseᐸᐳ();
	return a;
};
func (this Dummyᐸᐳ) arrayReverse__β1_Any____ResizableArray_β1__ResizableArray_β1_() Top { return this;
};
func (this Dummyᐸᐳ) arrayFindᐸlongᐸᐳᐳ(x ResizableArrayᐸlongᐸᐳᐳ, elem longᐸᐳ) boolᐸᐳ {
	i := 0
	b := true
	for i < TEST_SIZE {
		b = b != bool(x.containsᐸᐳ(longᐸᐳ(i)))
		i += 10000
	}
	return boolᐸᐳ(b)
};

func (this Dummyᐸᐳ) arrayFindᐸdoubleᐸᐳᐳ(x ResizableArrayᐸdoubleᐸᐳᐳ, elem doubleᐸᐳ) boolᐸᐳ { _ = x.containsᐸᐳ(elem);
	i := 0
	b := true
	for i < TEST_SIZE {
		b = b != bool(x.containsᐸᐳ(doubleᐸᐳ(i)))
		i += 10000
	}
	return boolᐸᐳ(b)
};

func (this Dummyᐸᐳ) arrayFindᐸintᐸᐳᐳ(x ResizableArrayᐸintᐸᐳᐳ, elem intᐸᐳ) boolᐸᐳ { _ = x.containsᐸᐳ(elem);
	i := 0
	b := true
	for i < TEST_SIZE {
		b = b != bool(x.containsᐸᐳ(intᐸᐳ(i)))
		i += 10000
	}
	return boolᐸᐳ(b)
};

func (this Dummyᐸᐳ) arrayFind__β1_Any____ResizableArray_β1__β1_bool__() Top { return this;
};
func (this Dummyᐸᐳ) Mainᐸintᐸᐳᐳ(elem intᐸᐳ) Dummyᐸᐳ {
	fmt.Println(1)
	a := this.arrayInsertᐸintᐸᐳᐳ(elem);
	fmt.Println(2)
	a2 := Dummyᐸᐳ{}.arrayReverseᐸintᐸᐳᐳ(a);
	fmt.Println(3)
	_ = Dummyᐸᐳ{}.arrayFindᐸintᐸᐳᐳ(a2, elem);
	fmt.Println(4)
	return this;
};
func (this Dummyᐸᐳ) Mainᐸdoubleᐸᐳᐳ(elem doubleᐸᐳ) Dummyᐸᐳ {
	fmt.Println(1)
	a := this.arrayInsertᐸdoubleᐸᐳᐳ(elem);
	fmt.Println(2)
	a2 := Dummyᐸᐳ{}.arrayReverseᐸdoubleᐸᐳᐳ(a);
	fmt.Println(3)
	_ = Dummyᐸᐳ{}.arrayFindᐸdoubleᐸᐳᐳ(a2, elem);
	fmt.Println(4)
	return this;
};
func (this Dummyᐸᐳ) Mainᐸlongᐸᐳᐳ(elem longᐸᐳ) Dummyᐸᐳ {
	fmt.Println(1)
	a := this.arrayInsertᐸlongᐸᐳᐳ(elem);
	fmt.Println(2)
	a2 := Dummyᐸᐳ{}.arrayReverseᐸlongᐸᐳᐳ(a);
	fmt.Println(3)
	_ = Dummyᐸᐳ{}.arrayFindᐸlongᐸᐳᐳ(a2, elem);
	fmt.Println(4)
	return this;
};
func (this Dummyᐸᐳ) Main__β1_Any____β1_Dummy__() Top { return this;
};
type Top interface {};
func main() { _ = Dummyᐸᐳ{}.Mainᐸintᐸᐳᐳ(4).Mainᐸlongᐸᐳᐳ(4).Mainᐸdoubleᐸᐳᐳ(4) }
