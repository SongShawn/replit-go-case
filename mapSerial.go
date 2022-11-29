// 多个map的值指向同一个内存，序列化反序列化

package main

import (
	"encoding/json"
	"fmt"
)

type someValue struct {
	A int
}

type dataBase struct {
	AA map[int]*someValue
	BB map[int]*someValue
}

func main() {
	a := &someValue{
		A: 10,
	}

	data := &dataBase{}
	data.AA = make(map[int]*someValue)
	data.BB = make(map[int]*someValue)

	data.AA[1] = a
	data.BB[1] = a

	fmt.Printf("AA %p, BB %p\r\n", data.AA[1], data.BB[1])
	marshaledData, _ := json.Marshal(data)
	fmt.Println(string(marshaledData))

	data1 := data
	json.Unmarshal(marshaledData, &data1)
	fmt.Printf("AA %p, BB %p\r\n", data1.AA[1], data1.BB[1])
	fmt.Println(string(marshaledData))
}
