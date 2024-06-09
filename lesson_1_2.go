package main

import "fmt"

func main() {

	nums := make([]int, 1, 2) // Создадим слайс и зададим длину и вместимость.
	fmt.Println(nums)
	appendSlice(nums, 1024)
	nums = nums[:cap(nums)] // чтобы отобразилоль 1024. мы должны увеличить вместимость основого слайса, т.к. программа не знает, что внутри функции есть еще один элемент.
	fmt.Println(nums)

	//	mutateSlice(nums, 1, 512)
	//	fmt.Println(nums)
}

func appendSlice(sl []int, val int) {
	sl = append(sl, val)

}

//func mutateSlice(sl []int, idx, val int) {
//	sl[idx] = val
//}
