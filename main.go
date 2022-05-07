package main

import (
	"fmt"
	"io"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type A struct {
	a int
}

func main() {
	a := &A{}

	var b *A

	checkNil(a)
	checkNil(b)
	checkNil(nil)
	checkNil(returnIntfNil())
    checkNil(returnRealNil())
}

func returnIntfNil() interface{} {
	var b *A

	return b
}

func returnRealNil() interface{} {
    return nil
}

func checkNil(val interface{}) {
	if val == nil {
		fmt.Println("is nil")
	} else {
		fmt.Println("not nil")
	}
}

func sliceInit() {
	a1 := make([]int, 10)
	fmt.Printf("len=%d cap=%d content=%#v\n", len(a1), cap(a1), a1)

	a2 := make([]int, 0, 10)
	fmt.Printf("len=%d cap=%d content=%#v\n", len(a2), cap(a2), a2)
}
func mathDiv() {
	fmt.Println(math.Ceil(1 * 1.0 / 3))
}

func ginCase() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong by sxn",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func testSwitchCase1() {
	var a int

	a = 7

	switch {
	case a > 0 && a < 4:
		fmt.Println("a>0 and a<4")

	case a >= 5 && a < 10:
		fmt.Println("a>=5 and a<10")

	case a >= 10 && a < 20:
		fmt.Println("a>=10 and a < 20")

	default:
		fmt.Println("default")
	}
}

func httpGetBaiduPage() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("http get err ", err)
		return
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("resp read err ", err)
		return
	}

	fmt.Println(string(data))
}
