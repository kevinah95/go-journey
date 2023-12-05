package examples

import (
	"fmt"
	"golang.org/x/exp/slices"
)

func slices_() {
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)
	fmt.Println(arr[2])

	var s []int //slice
	fmt.Println(s)
	s = []int{1, 2, 3}
	fmt.Println(s[1])
	s[1] = 99
	fmt.Println(s)

	s = append(s, 5, 10, 15)
	fmt.Println(s)

	s = slices.Delete(s, 1, 3)
	fmt.Println(s)

}
