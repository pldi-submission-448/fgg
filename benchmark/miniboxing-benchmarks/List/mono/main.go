package main;

import "fmt"

const TEST_SIZE = 3000000

type Dummyᐸᐳ struct {};
type intᐸᐳ int8;
type longᐸᐳ int32;
type doubleᐸᐳ struct {a int64; b int64;};
type bool_ᐸᐳ bool;
func (this Dummyᐸᐳ) x___Any___Any___Dummy__() Top { return this;
};
type Listᐸdoubleᐸᐳᐳ interface { Containsᐸᐳ(x doubleᐸᐳ) bool_ᐸᐳ; Contains___double___bool___() Top };
type Listᐸlongᐸᐳᐳ interface { Containsᐸᐳ(x longᐸᐳ) bool_ᐸᐳ; Contains___long___bool___() Top };
type Listᐸintᐸᐳᐳ interface { Containsᐸᐳ(x intᐸᐳ) bool_ᐸᐳ; Contains___int___bool___() Top };
type Consᐸlongᐸᐳᐳ struct { head longᐸᐳ; tail Listᐸlongᐸᐳᐳ };
type Consᐸintᐸᐳᐳ struct { head intᐸᐳ; tail Listᐸintᐸᐳᐳ };
type Consᐸdoubleᐸᐳᐳ struct { head doubleᐸᐳ; tail Listᐸdoubleᐸᐳᐳ };
func (c Consᐸlongᐸᐳᐳ) Containsᐸᐳ(x longᐸᐳ) bool_ᐸᐳ {
	var xs Listᐸlongᐸᐳᐳ = c;
	for {
		if xs_, ok := xs.(Consᐸlongᐸᐳᐳ); ok {
			if xs_.head == x {
				return true
			}else {
				xs = xs_.tail
			}
		}
		if _, ok := xs.(Nilᐸlongᐸᐳᐳ); ok {
			return false
		}
	}
};
func (c Consᐸintᐸᐳᐳ) Containsᐸᐳ(x intᐸᐳ) bool_ᐸᐳ {
	var xs Listᐸintᐸᐳᐳ = c;
	for {
		if xs_, ok := xs.(Consᐸintᐸᐳᐳ); ok {
			if xs_.head == x {
				return true
			}else {
				xs = xs_.tail
			}
		}
		if _, ok := xs.(Nilᐸintᐸᐳᐳ); ok {
			return false
		}
	}
};
func (c Consᐸdoubleᐸᐳᐳ) Containsᐸᐳ(x doubleᐸᐳ) bool_ᐸᐳ {
	var xs Listᐸdoubleᐸᐳᐳ = c;
	for {
		if xs_, ok := xs.(Consᐸdoubleᐸᐳᐳ); ok {
			if xs_.head == x {
				return true
			}else {
				xs = xs_.tail
			}
		}
		if _, ok := xs.(Nilᐸdoubleᐸᐳᐳ); ok {
			return false
		}
	}
};

func (c Consᐸdoubleᐸᐳᐳ) Contains___double___bool___() Top { return c;
};
func (c Consᐸlongᐸᐳᐳ) Contains___long___bool___() Top { return c;
};
func (c Consᐸintᐸᐳᐳ) Contains___int___bool___() Top { return c;
};
type Nilᐸintᐸᐳᐳ struct {};
type Nilᐸlongᐸᐳᐳ struct {};
type Nilᐸdoubleᐸᐳᐳ struct {};
func (n Nilᐸlongᐸᐳᐳ) Containsᐸᐳ(x longᐸᐳ) bool_ᐸᐳ {
	return false
};
func (n Nilᐸdoubleᐸᐳᐳ) Containsᐸᐳ(x doubleᐸᐳ) bool_ᐸᐳ {
	return false
};
func (n Nilᐸintᐸᐳᐳ) Containsᐸᐳ(x intᐸᐳ) bool_ᐸᐳ {
	return false;
};
func (n Nilᐸintᐸᐳᐳ) Contains___int___bool___() Top { return n;
};
func (n Nilᐸlongᐸᐳᐳ) Contains___long___bool___() Top { return n;
};
func (n Nilᐸdoubleᐸᐳᐳ) Contains___double___bool___() Top { return n;
};
func (d Dummyᐸᐳ) ListInsertᐸlongᐸᐳᐳ(x longᐸᐳ) Listᐸlongᐸᐳᐳ {
	var l Listᐸlongᐸᐳᐳ = Nilᐸlongᐸᐳᐳ{}
	for i := 0; i < TEST_SIZE; i++ {
		l = Consᐸlongᐸᐳᐳ{
			head: longᐸᐳ(x),
			tail: l,
		}
	}
	return l
};
func (d Dummyᐸᐳ) ListInsertᐸintᐸᐳᐳ(x intᐸᐳ) Listᐸintᐸᐳᐳ {
	var l Listᐸintᐸᐳᐳ = Nilᐸintᐸᐳᐳ{}
	for i := 0; i < TEST_SIZE; i++ {
		l = Consᐸintᐸᐳᐳ{
			head: intᐸᐳ(x),
			tail: l,
		}
	}
	return l
};
func (d Dummyᐸᐳ) ListInsertᐸdoubleᐸᐳᐳ(x doubleᐸᐳ) Listᐸdoubleᐸᐳᐳ {
	var l Listᐸdoubleᐸᐳᐳ = Nilᐸdoubleᐸᐳᐳ{}
	for i := 0; i < TEST_SIZE; i++ {
		l = Consᐸdoubleᐸᐳᐳ{
			head: doubleᐸᐳ(x),
			tail: l,
		}
	}
	return l
};
func (d Dummyᐸᐳ) ListInsert__β1_Any____β1_List_β1_() Top { return d;
};
func (d Dummyᐸᐳ) ListFindᐸdoubleᐸᐳᐳ(x doubleᐸᐳ, l Listᐸdoubleᐸᐳᐳ) bool_ᐸᐳ {
	b := true
	for i := 0; i < TEST_SIZE; i+=10000 {
		b = b != bool(l.Containsᐸᐳ(doubleᐸᐳ{int64(i),int64(i)}))
	}
	return bool_ᐸᐳ(b)
};
func (d Dummyᐸᐳ) ListFindᐸlongᐸᐳᐳ(x longᐸᐳ, l Listᐸlongᐸᐳᐳ) bool_ᐸᐳ {
	b := true
	for i := 0; i < TEST_SIZE; i+=10000 {
		b = b != bool(l.Containsᐸᐳ(longᐸᐳ(i)))
	}
	return bool_ᐸᐳ(b)
};
func (d Dummyᐸᐳ) ListFindᐸintᐸᐳᐳ(x intᐸᐳ, l Listᐸintᐸᐳᐳ) bool_ᐸᐳ {
	b := true
	for i := 0; i < TEST_SIZE; i+=10000 {
		b = b != bool(l.Containsᐸᐳ(intᐸᐳ(i)))
	}
	return bool_ᐸᐳ(b)
};
func (d Dummyᐸᐳ) ListFind__β1_Any____β1_List_β1__bool___() Top { return d;
};
func (this Dummyᐸᐳ) Main1ᐸᐳ() Dummyᐸᐳ { d := Dummyᐸᐳ{};
	fmt.Println("main1_1");
	t1 := d.ListInsertᐸintᐸᐳᐳ(1);
	fmt.Println("main1_2");
	_ = d.ListFindᐸintᐸᐳᐳ(1, t1);
	fmt.Println("main1_3");
	return this;
};
func (this Dummyᐸᐳ) Main1___Dummy__() Top { return this;
};
func (this Dummyᐸᐳ) Main2ᐸᐳ() Dummyᐸᐳ { d := Dummyᐸᐳ{};
	fmt.Println("main2_1");
	t1 := d.ListInsertᐸlongᐸᐳᐳ(1);
	fmt.Println("main2_2");
	_ = d.ListFindᐸlongᐸᐳᐳ(1, t1);
	fmt.Println("main2_3");
	return this;
};
func (this Dummyᐸᐳ) Main2___Dummy__() Top { return this;
};
func (this Dummyᐸᐳ) Main3ᐸᐳ() Dummyᐸᐳ { d := Dummyᐸᐳ{};
	fmt.Println("main3_1");
	t1 := d.ListInsertᐸdoubleᐸᐳᐳ(doubleᐸᐳ{1, 1});
	fmt.Println("main3_2");
	_ = d.ListFindᐸdoubleᐸᐳᐳ(doubleᐸᐳ{1, 1}, t1);
	fmt.Println("main3_3");
	return this;
};
func (this Dummyᐸᐳ) Main3___Dummy__() Top { return this;
};
type Top interface {};
func main() { _ = Dummyᐸᐳ{}.Main1ᐸᐳ().Main2ᐸᐳ().Main3ᐸᐳ() }
