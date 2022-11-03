package main

import "fmt"

func main() {
    s := "aaaaaa"
    for i, v := range s {
        fmt.Println(i, v)
    }
}