package main

import (
	"fmt"
)

func main() {
	initSlice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(&initSlice[1])

	//AppendVector
	//Создается новый слайс, длина 5, вместимость 8
	initSlice = append(initSlice, 5)
	fmt.Println(initSlice, len(initSlice), cap(initSlice))
	fmt.Println(&initSlice[1])
	fmt.Println("--------------------------------------------")

	//Copy
	//Чтобы не зависеть от изначального массива, можно сделать копию.
	copySlice := make([]int, len(initSlice))
	copy(copySlice, initSlice)
	copySlice = append(initSlice[:0:0], initSlice...)
	fmt.Println(copySlice)
	fmt.Println(&initSlice[1])
	fmt.Println(&copySlice[1])
	fmt.Println("--------------------------------------------")

	//Cut
	//Отрезаем 1 элемент, длина 7. Вместимость без изменений.
	initSlice = append(initSlice[:0], initSlice[1:]...)
	fmt.Println(initSlice, len(initSlice), cap(initSlice))
	fmt.Println(&initSlice[1])
	fmt.Println("--------------------------------------------")

	//Delete
	// Удаляем 2 элемента, длина 6. Вместимость без изменений.
	initSlice = append(initSlice[:1], initSlice[1+1:]...)
	fmt.Println(initSlice, len(initSlice), cap(initSlice))
	fmt.Println(&initSlice[1])
	fmt.Println("--------------------------------------------")

	//Delete without preserving order
	// Первый элемент становится 5, т.к. новая длина становится 5. Слайс уменьшается длина -1.
	initSlice[0] = initSlice[len(initSlice)-1]
	initSlice = initSlice[:len(initSlice)-1]
	fmt.Println(initSlice, len(initSlice), cap(initSlice))
	fmt.Println(&initSlice[1])
	fmt.Println("--------------------------------------------")

	//Expand
	//Берем 0 элемент + длина массива 5, состоящая из 0 + элементы от 0 до конца.
	initSlice = append(initSlice[:1], append(make([]int, 5), initSlice[0:]...)...)
	fmt.Println(initSlice, len(initSlice), cap(initSlice))
	fmt.Println(&initSlice[1])
	fmt.Println("--------------------------------------------")

	//Extend
	//Расширяем на 5 элементов, длина 16, вместимость увеличивается в 2 раза.
	initSlice = append(initSlice, make([]int, 5)...)
	fmt.Println(initSlice, len(initSlice), cap(initSlice))
	fmt.Println(&initSlice[1])
	fmt.Println("--------------------------------------------")

	//Insert
	//Вставляем 2 элемента в начало слайса. Длина 18, вместимость не меняется.
	initSlice = append(initSlice[:1], append([]int{2, 3}, initSlice[1:]...)...)
	fmt.Println(initSlice, len(initSlice), cap(initSlice))
	fmt.Println(&initSlice[1])
	fmt.Println("--------------------------------------------")

	//InsertVector
	//Вставляем элементы, равные кол-ву длины основного слайса. Вместимость в 2раза.
	insertVector := make([]int, len(initSlice))
	initSlice = append(initSlice[:0], append(insertVector, initSlice[0:]...)...)
	fmt.Println(initSlice, len(initSlice), cap(initSlice))
	fmt.Println(insertVector, len(insertVector), cap(insertVector))
	fmt.Println(&insertVector[1])
	fmt.Println(&initSlice[1])
	fmt.Println("--------------------------------------------")

	//Push
	// Точно не скажу, возможно похожее на append
	x := 5
	initSlice = append(initSlice, x)
	fmt.Println(initSlice, len(initSlice), cap(initSlice))
	fmt.Println(insertVector, len(insertVector), cap(insertVector))
	fmt.Println(&insertVector[1])
	fmt.Println(&initSlice[1])
	fmt.Println("--------------------------------------------")

	//Push Front/Unshift
	//Сокращение вместимости на 2 байта
	x = 7
	initSlice = append([]int{x}, initSlice...)
	fmt.Println(initSlice, len(initSlice), cap(initSlice))
	fmt.Println(&initSlice[1])
	fmt.Println("--------------------------------------------")

	//Pop Front/Shift
	//Сократит длину и вместимость.
	x = 7
	x, initSlice = initSlice[0], initSlice[1:]
	fmt.Println(initSlice, len(initSlice), cap(initSlice))
	fmt.Println(&initSlice[1])
	fmt.Println("--------------------------------------------")

	//Pop
	//Возможно уберет 7 индекс, судя по выводу.
	x = 7
	x, initSlice = initSlice[len(initSlice)-1], initSlice[:len(initSlice)-1]
	fmt.Println(initSlice, len(initSlice), cap(initSlice))
	fmt.Println(&initSlice[1])
	fmt.Println("--------------------------------------------")

	//Filter (in place)
	// Ошибка в keep, не знаю как пофиксить.
	//n := 0
	//
	//for _, x := range initSlice {
	//	if keep(x) {
	//		initSlice[n] = x
	//		n++
	//	}
	//}
	//initSlice = initSlice[:n]
	//fmt.Println(initSlice, len(initSlice), cap(initSlice))
	//fmt.Println(&initSlice[1])
	//fmt.Println("--------------------------------------------")
}
