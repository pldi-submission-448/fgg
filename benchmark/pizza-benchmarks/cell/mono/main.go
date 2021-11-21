package main;
//const iteration = 1000000;
const iteration = 1000000;

type Int1ᐸᐳ  int;
type Int2ᐸᐳ struct {};
type CellᐸInt2ᐸᐳᐳ struct { elem Int2ᐸᐳ };
type CellᐸInt1ᐸᐳᐳ struct { elem Int1ᐸᐳ };
func (this *CellᐸInt2ᐸᐳᐳ) addᐸᐳ(elem Int2ᐸᐳ) CellᐸInt2ᐸᐳᐳ { return CellᐸInt2ᐸᐳᐳ{elem};
};
func (this *CellᐸInt1ᐸᐳᐳ) addᐸᐳ(elem Int1ᐸᐳ) CellᐸInt1ᐸᐳᐳ { return CellᐸInt1ᐸᐳᐳ{elem};
};
func (this CellᐸInt2ᐸᐳᐳ) add___Int2___Cell_Int2___() Top { return this;
};
func (this CellᐸInt1ᐸᐳᐳ) add___Int1___Cell_Int1___() Top { return this;
};
func (this *CellᐸInt2ᐸᐳᐳ) getᐸᐳ() Int2ᐸᐳ { return this.elem;
};
func (this *CellᐸInt1ᐸᐳᐳ) getᐸᐳ() Int1ᐸᐳ { return this.elem;
};
func (this CellᐸInt2ᐸᐳᐳ) get___Int2__() Top { return this;
};
func (this CellᐸInt1ᐸᐳᐳ) get___Int1__() Top { return this;
};
type Dummyᐸᐳ struct {};
func (this Dummyᐸᐳ) CellTestᐸInt1ᐸᐳᐳ(elem Int1ᐸᐳ) Dummyᐸᐳ {
	c := &CellᐸInt1ᐸᐳᐳ{elem}
	c.addᐸᐳ(Int1ᐸᐳ(42))
	for i := 0; i < iteration; i++ {
		c.getᐸᐳ()
	}
	return Dummyᐸᐳ{};
};
/*func (this Dummyᐸᐳ) CellTestᐸInt2ᐸᐳᐳ(elem Int2ᐸᐳ) Dummyᐸᐳ { c := CellᐸInt2ᐸᐳᐳ{elem}.addᐸᐳ(elem);
	_ = c.getᐸᐳ();
	return Dummyᐸᐳ{};
};*/
func (this Dummyᐸᐳ) CellTest__β1_Any____β1_Dummy__() Top { return this;
};
type Top interface {};
func main() {
	Dummyᐸᐳ{}.CellTestᐸInt1ᐸᐳᐳ(Int1ᐸᐳ(1))
}

