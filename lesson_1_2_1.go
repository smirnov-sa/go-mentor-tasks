package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	intSlice := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		intSlice = append(intSlice, i)
	}
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Printf(
		"| intSlice slice %v, len: %d, cap: %d\n",
		intSlice,
		len(intSlice),
		cap(intSlice),
	)
	// Создаем слайс с 0 длиной, после прогоняем цикл, чтобы было 10/10.

	fmt.Println("---------------------------------------------------------------------------")
	cut := intSlice[2:4]
	fmt.Printf(
		"| cut slice %v, len: %d, cap: %d\n",
		cut,
		len(cut),
		cap(cut),
	)
	// Создаем новый слайс из intSlice с длиной 2, вместимость 8. Старый слайс имел вместимость 10, убрали 2 элемента.

	fmt.Println("---------------------------------------------------------------------------")
	cut = cut[:cap(cut)]
	//_ = cut[2] Если закомментировать предыдущую строку, то мы не сможет обратиться к слайсу, т.к. превышаем вместимость.
	fmt.Printf(
		"| cut slice %v, len: %d, cap: %d\n",
		cut,
		len(cut),
		cap(cut),
	)
	fmt.Println("---------------------------------------------------------------------------")
	// Увеличиваем длину с 0 до конца.

	fmt.Printf(
		"| intSlice %d, cut: %d |\n"+
			"| intSlice %v, cut: %v |\n",
		reflect.ValueOf(intSlice).Pointer(),
		reflect.ValueOf(cut).Pointer(),
		*(*reflect.SliceHeader)(unsafe.Pointer(&intSlice)),
		*(*reflect.SliceHeader)(unsafe.Pointer(&cut)),
	)
	// Просмотр поинтеров в области памяти основного слайса и обрезанного.

	fmt.Println("---------------------------------------------------------------------------")

	cut[0] = 1 << 32
	fmt.Printf("| intSlice: #{intSlice}, len: #{len(intSlice)}, cap: #{cap(intSlice)}", intSlice, len(intSlice), cap(intSlice))
	fmt.Printf("| cut: %v, len: %d, cap: %d |\n", cut, len(cut), cap(cut))
	fmt.Println("---------------------------------------------------------------------------")

	cut = append(cut, (1<<32)+1)
	cut[0] = 1 << 10
	fmt.Printf("| intSlice: %v, intSlice: %d, intSlice: %d |\n", intSlice, len(intSlice), cap(intSlice))
	fmt.Printf("| cut: %v, len: %d, cap: %d |\n", cut, len(cut), cap(cut))
	fmt.Println("---------------------------------------------------------------------------")

	fmt.Printf(
		"| intSlice %d, cut: %d |\n"+
			"| intSlice %v, cut: %v |\n",
		reflect.ValueOf(intSlice).Pointer(),
		reflect.ValueOf(cut).Pointer(),
		*(*reflect.SliceHeader)(unsafe.Pointer(&intSlice)),
		*(*reflect.SliceHeader)(unsafe.Pointer(&cut)),
	)
	// Просмотр поинтеров в области памяти основного слайса и обрезанного.
}
