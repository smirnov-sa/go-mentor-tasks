package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4}

	fmt.Println("before", list)
	handle(list[:1]) //Берем только 1 элемент
	fmt.Println("after", list)
}

func handle(list []int) {
	newList := make([]int, len(list))
	copy(newList, list)
	list = append(newList, 5)
	fmt.Println("append", list)
}

// Если нет необьходимости менять исходный слайс, создадим копию.
