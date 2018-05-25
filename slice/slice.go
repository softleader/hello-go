package slice

import "fmt"

func Play() (n []int, m []int) {
	n = []int{1, 2, 3, 4, 5}
	fmt.Println(n)

	m = n[2:]
	fmt.Println(m)

	n[2] *= 10
	fmt.Println(m)

	m = append(m, 6)
	fmt.Println(m)
	fmt.Println(n)

	n[2] *= 10
	fmt.Println(m)
	fmt.Println(n)

	return
}
