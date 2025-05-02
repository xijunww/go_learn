package component

type Inner struct {
}

func (i Inner) DoSomething() {

}

type Outer struct {
	Inner
}

type OuterPtr struct {
	*Inner
}
type OOOOOuter struct {
	Outer
}

func UserInner() {
	var o Outer
	o.DoSomething()

	var o2 *OuterPtr
	o2.DoSomething()

	o1 := Outer{
		Inner: Inner{},
	}
	o1.DoSomething()
}
