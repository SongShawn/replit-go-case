package main

import "fmt"

func main() {
	a := make([]int, 0, 10)
	fmt.Println(len(a), cap(a))
	a = append(a, 100)
	fmt.Println(len(a), cap(a))
	fmt.Println(a)

	// b 和 a 共享同一个底层内存
	b := a[:0]
	fmt.Println(len(b), cap(b))
	b = append(b, 200)
	fmt.Println(a)
	fmt.Println(b)

    // 只能拷贝到目的切片c已经使用的空间中
	var c = make([]int, 10)
	// var c []int
	fmt.Println(copy(c, a))
	fmt.Println(c)
}
