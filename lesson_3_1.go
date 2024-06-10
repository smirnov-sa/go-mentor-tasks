package main

import "fmt"

func main() {
	//var list []int
	//fmt.Println(len(list) == 0)
	//
	//list = []int{}
	//fmt.Println(len(list) == 0)

	// Проверять слайс лучше по длине

	//list := make([]int, 5, 10)
	//fmt.Println(len(list), cap(list))

	list := []int{1, 2, 3, 4, 8}

	fmt.Println(double(list))
}
func double(nums []int) []int {
	//res := make([]int, 0, len(nums))
	res := make([]int, len(nums))

	//for _, num := range nums {
	//	res = append(res, num * 2)
	//}

	for i, num := range nums {
		res[i] = num * 2
	}
	return res
}
