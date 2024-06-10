// Реализация функции Append
package main

import "fmt"

func main() {
	list := make([]int, 4, 5)
	list = Append(list, 1)
	fmt.Println(list, len(list), cap(list))
	// размер слайса изменился на 1, а емкость удвоилась в 2
}

func Append(list []int, el int) []int {
	var res []int
	resLen := len(list) + 1

	if resLen <= cap(list) {
		res = list[:resLen]
	} else {
		resCap := resLen
		if resCap < 3*len(list) {
			resCap = 3 * len(list)
		}
		res = make([]int, resLen, resCap)
		copy(res, list)
	}
	res[len(list)] = el
	return res
}
