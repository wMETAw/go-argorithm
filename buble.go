package main

import "fmt"

func bubleSort(a []int) []int {
	c := 0
	for i := 0; i < len(a); i++ {
		for n := 0; n < len(a)-i-1; n++ {
			c++
			if a[n] > a[n+1] {
				a[n], a[n+1] = a[n+1], a[n]
			}
		}
	}
	fmt.Println(c)
	return a
}

func main() {

	a := []int{1, 5, 6, 2, 4, 8, 7, 9, 6, 3}
	fmt.Println(bubleSort(a))
}
