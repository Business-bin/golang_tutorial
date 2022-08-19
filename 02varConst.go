package main

import (
	"fmt"
	"reflect"
)

// 전역에서 변수 선언
var message5 string
var message6 string = "Hello, message6"

// message6 := "" <<< error, 함수 외부에서는 모든 문장이 키워드로 시작해야 함

func main() {
	// go는 선언한 변수 사용하지 않으면 에러
	var message1 string
	message1 = "Hello, message1"
	fmt.Println(message1)

	var message2 string = "Hello, message2"
	fmt.Println(message2)

	var message3 = "Hello, message3"
	fmt.Println(message3)
	fmt.Println(reflect.TypeOf(message3))
	// 단축변수 선언,타입이 명시되지 않은경우 컴파일시 자동으로 타입 추론
	message4 := "Hello, message4"
	fmt.Println(message4)

	message5 = "Hello, message5"
	fmt.Println(message5)

	fmt.Println(message6)

	// 여러개 변수 선언(타입 달라도 됨)
	str1, str2, str3, bool1 := "sss", "ttt", "rrr", true
	fmt.Println(str1, str2, str3, bool1)

	// 제로 값 - 변수를 선언만 하고 명적으로 초기화 하지 않은 상태에서 디폴트로 할당되는 값
	var (
		msg7  string  // ""
		msg8  int     // 0
		msg9  float64 // 0
		msg10 bool    // false
	)
	msg7 = "여기"
	fmt.Printf("%#v %#v %#v %#v \n", msg7, msg8, msg9, msg10)

	// 상수는 선언과 동시에 초기화 해줘야 함
	const con1 int = 10
	fmt.Println(con1)
	const con2, con3 = 20, "golang"
	fmt.Println(con2, con3)
	const (
		con4 int    = 30
		con5 string = "con5"
	)
	fmt.Println(con4, con5)

}
