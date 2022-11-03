package main

import "fmt"

func main() {
    var a = 10
A_LABEL:
    for ;a < 20; a++ {
        for b:=0;b<5;b++ {
            if b == 2 {
                continue A_LABEL
            }
            fmt.Println(a, b)
        }
    }
}