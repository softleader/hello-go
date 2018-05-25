package loop

import "fmt"

func Go(m []int) {
	// m: [30 4 5 6]
	p := make(map[int]bool) // key: int, value: == even?

	var zero, even, odd int

	//Zero:
	for _, v := range m {
		//if v == 0 {
		//	zero += 1
		//	break Zero
		//} else if result := v % 2; result == 0 {
		//	even += 1
		//	continue
		//} else {
		//	odd += 1
		//	break
		//}

		switch {
		case v == 0:
			zero += 1
			break
		case v%2 == 0:
			even += 1
			p[v] = true
			break
		default:
			odd += 1
			p[v] = false
		}
	}

	fmt.Printf("zero: %v, even: %v, odd: %v\n", zero, even, odd)
	fmt.Printf("%#v\n", p)
}
