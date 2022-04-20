package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	httpGetBaiduPage()
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
