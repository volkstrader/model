package model

import "github.com/volkstrader/girlfriend"

type Profile struct {
	Name string
	Age  int
}

func (v Profile) Dating() bool {
	gf := girlfriend.Profile{
		Age: v.Age,
	}
	return !gf.Breakup() || gf.Makeup()
}
