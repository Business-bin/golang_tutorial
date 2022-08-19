package main

import "fmt"

func main() {
	// Go는 while문이 없음.
	var1 := 0
	for i := 0; i < 10; i++ {
		var1 += i
	}
	fmt.Println(var1)
	// while문 처럼 사용
	var2 := 1
	for var2 < 10 {
		var2 += var2
	}
	fmt.Println(var2, "\n")

	// for ~ range
	// idx에는 키 또는 인덱스가 str에는 값이 들어감
	strlist := []string{"A", "B", "C"}
	for idx, str := range strlist {
		fmt.Printf("값 = %#v 인덱스 = %#v \n", str, idx)
	}
}
