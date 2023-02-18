package main

import "fmt"
import "time"

func main() {
	a := make([]uint32, 10)
	fmt.Println(len(a), cap(a))
	a = append(a, 10)
	fmt.Println(len(a), cap(a))

	b := make([]uint32, 0, 10)
	fmt.Println(len(b), cap(b))
	b = append(b, 10)
	fmt.Println(len(b), cap(b))

	time1 := time.Now()
    fmt.Println(time1)
    time2 := time.Now()
    fmt.Println(time1.Before(time2))
}
