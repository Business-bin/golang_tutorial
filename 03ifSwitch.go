package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	// 1 if ~ else if ~ else
	// Go는 조건문의 경우 중괄호가 없다
	var1 := 20
	if var1 == 10 {
		fmt.Println("if")
	} else if var1 == 20 {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}
	// 변수와 함께 사용 - 조건문의 조건에서 변수 선언
	if var2, err := http.Get("https://example.com.com"); err == nil { // 세미콜론을 구본자로 값을 받아온 후 비교연산
		fmt.Println(var2.StatusCode)
	}

	// 2 switch ~ case
	var3, _ := http.Get("https://example.com.com")
	// break문이 필요 없다
	switch var3.StatusCode {
	case 200:
		fmt.Println("Successful")
	case 404:
		fmt.Println("Not Found")
	default:
		fmt.Println("Good")
	}

	// fallthrough 키워드를 사용 해 다음으로 넘어갈 수 있다(부합한 조건 뒤로는 모두 실행 됨)
	switch var3.StatusCode {
	case 200:
		fmt.Println("Successful fallthrough")
		fallthrough
	case 404:
		fmt.Println("Not Found fallthrough")
		fallthrough
	default:
		fmt.Println("Good fallthrough")
	}

	// 조건 없이 사용 가능
	switch {
	case var3.StatusCode == 200:
		fmt.Println("Successful")
	case var3.StatusCode == 404:
		fmt.Println("Not Found")
	default:
		fmt.Println("Good")
	}

	// 변수와 함께 사용
	switch var4, _ := http.Get("https://example.com.com"); var4.StatusCode {
	case 200:
		fmt.Println("Successful var4")
		fallthrough
	case 404:
		fmt.Println("Not Found var4")
		fallthrough
	default:
		fmt.Println("Good var4")
	}

	// 조건문
	str1 := ""
	if str1 = "1"; str1 == "0" {
		str1 = "0"
	}
	fmt.Println("조건문 str1  =", str1)

	str2 := "Go "
	str3 := "go"

	if strings.EqualFold(str2, str3) {
		fmt.Println("같음")
	} else {
		fmt.Println("다름")
	}

}
