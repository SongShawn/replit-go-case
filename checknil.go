package main

import "fmt"

type checkNilA struct {
	a int
}

var datas = make(map[string]*checkNilA)

func getOneData(key string) *checkNilA {
	return datas[key]
}

func getOneData1(key string) *checkNilA {
    if v,ok := datas[key]; ok {
        return v
    }else{
        return nil
    }
}

func main() {
    var x interface{}
	datas["a"] = &checkNilA{a: 10}

    // delete(datas, "a")
	x = getOneData1("b")
	if x == nil {
		fmt.Println("is nil")
	} else {
		fmt.Println("not nil")
	}
}
