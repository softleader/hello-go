package main

import (
	"fmt"
	"github.com/softleader/hello-go/struc"
	"github.com/softleader/hello-go/loop"
	"github.com/softleader/hello-go/slice"
)

func main() {
	fmt.Println("Hello, Go!")
	_, m := slice.Play()
	loop.Go(m)

	//s := struc.Student{
	//	Name:   "Matt",
	//	Number: 1,
	//}

	s := struc.Student{}
	s.Name = "Matt"
	s.Number = 1
	fmt.Printf("%#v\n", s)

	s.UpdateName("Abc")
	s.Output()

	fmt.Println(struc.IsMale(s))
}
