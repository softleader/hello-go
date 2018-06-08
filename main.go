package main

import (
	"fmt"
	"github.com/softleader/hello-go/struc"
)

func main() {
	fmt.Println("Hello, Go!")

	// === slice ===
	// _, m := slice.Play()

	// === loop ===
	// loop.Go(m)

	// === struct ===

	//s := struc.Student{
	//	Name:   "Matt",
	//	Number: 1,
	//}

	s := struc.Student{}
	s.Name = "Matt"
	s.Number = 1
	fmt.Printf("%#v\n", s)
	s.Output()
	fmt.Println(struc.IsMale(s))

}

