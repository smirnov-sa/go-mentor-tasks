package main

import (
	"os"
	"regexp"
)

//func main() {
//	list := make([]int, 4, 5)
//	list = append(list, 1)
//
//	list[0] = 5
//	//list2[0] = 9
//	fmt.Println(list)
//	//fmt.Println(list2)
//}

func main() {

}

var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
	b, _ := os.ReadFile(filename)

	b = digitRegexp.Find(b)

	res := make([]byte, len(b))
	for _, num := range b {
		res = append(res, num*2)
	}
	//copy(res, b)
	//
	return res

}

//Попытка реализации с copy на append
