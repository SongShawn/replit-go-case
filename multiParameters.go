package main

import "fmt"

func logWrap1(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func logWrap(format string, args ...interface{}) {
    format += ", %s.\n"
    args = append(args, "100x")
    
    logWrap1(format, args...)
}

func main() {
	logWrap("%d %d", 1, 2)

    var a map[int]string

    fmt.Printf("len=%d\n", len(a))
}
