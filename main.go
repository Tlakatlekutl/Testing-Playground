package main

import "fmt"

var coins = []int{10, 5 ,2 , 1}

func calc_change(N int) {
	for _, c := range coins {
		if n := N/c; n != 0 {
			N %= c
			fmt.Printf("%d: %d\n", c, n)
		}
	}
}

func main() {
	calc_change(18)
}
