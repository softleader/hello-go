package struc

import (
	"fmt"
)

type People interface {
	Sex() string
}

type Student struct {
	Name   string
	Number int
}

func (s Student) Output() {
	fmt.Printf("My name is %v\n", s.Name)
}

func (s Student) Sex() string {
	return "M"
}

func IsMale(p People) bool {
	return p.Sex() == "M"
}
