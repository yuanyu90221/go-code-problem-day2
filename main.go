package main

import (
	"github.com/yuanyu90221/go-code-problem-day2/sol"
)

func main() {
	array := []int{6, 5, 4, 3, 2, 1}
	list := sol.ConvertList(&array)
	sol.Transerval(list)
	list = sol.ReverseBetween(list, 2, 6)
	sol.Transerval(list)
}
