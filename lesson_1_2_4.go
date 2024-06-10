package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
	"unsafe"
)

func main() {
	stats := new(runtime.MemStats)
	runtime.ReadMemStats(stats)
	fmt.Println("-----------------------------------------------------")
	fmt.Printf("| stats: #{stats.HeapAlloc} |\n")

	bytes := make([]byte, math.MaxInt32)

	runtime.ReadMemStats(stats)
	fmt.Println("-----------------------------------------------------")
	fmt.Printf("| stats after bytes created: %d |\n", stats.HeapAlloc)

	bytesAsString := string(bytes) // представление строки в байтах
	runtime.ReadMemStats(stats)

	fmt.Println("----------------------------------------------------")
	fmt.Printf("| bytes len: %d, cap: %d, size: %d |\n| string len: %d, size: %d |\n",
		len(bytes),
		cap(bytes),
		unsafe.Sizeof(bytes),
		len(bytesAsString),
		unsafe.Sizeof(bytesAsString),
	)
	fmt.Printf("| stats agter string created %d |\n", stats.HeapAlloc)
	// у строки нет вместимости, она иммутабельна

	go func(bs string, memStats *runtime.MemStats) {
		stats := new(runtime.MemStats)
		runtime.ReadMemStats(stats)
		fmt.Printf("| stats in lambda: %d |\n", stats.HeapAlloc)
		_ = bs
	}(bytesAsString, stats)
	fmt.Println("----------------------------------------------------")
	<-time.After(time.Second * 2)

	//конвертация байт в строку
	someBytes := []byte("Hello")
	stringUnsafe := *(*string)(unsafe.Pointer(&someBytes))
	fmt.Printf("| bytes: %v, %T. string: %s, %T |\n", someBytes, someBytes, stringUnsafe, stringUnsafe)
	someBytes[0] = 'Y'
	fmt.Printf("| bytes: %v, %T. string: %s, %T |\n", someBytes, someBytes, stringUnsafe, stringUnsafe)
	time.Sleep(time.Hour)
	fmt.Println(bytesAsString, bytes)

	//SliceHeader 24 байта
	//StringHeader 16 байт

}
