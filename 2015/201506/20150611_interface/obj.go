package main

import (
	"fmt"
)

type Ob struct {
	Id int
}

func (this *Ob) Just_Do_It() {
	fmt.Printf("in just do it\n")
	fmt.Printf("%d\n", this.Id)
}

func main() {
	var aif IF

	aif = &Ob{11}
	ts(aif)

	ka := Ob{13}
	ts(&ka)

	kb := &Ob{14}
	ts(kb)

	// ob := Ob{15}
	// ts(ob)

	os := Ob{16}
	is(&os)

	ik := Ob{17}
	ip(&ik)
}

func ts(in interface{}) {
	// in.(IF).Just_Do_It()

	ab := in.(IF)
	ab.Just_Do_It()
}

func is(in IF) {
	in.Just_Do_It()
}

func ip(in *IF) {
	(*in).Just_Do_It()
}
