package main

import "log"

/*
	- 구조체는 클래스처럼 객체를 찍어내기 위한 판이라 생각하면 접근하기 좋음.
	- struct 카워드 사용
*/
// 구조체 타입 가진 변수 선언
type Struct01 struct {
	num  int
	str1 string
	str2 string
}

func main() {
	// 구조체 타입 변수에 값 할당 3가지 방법
	var struct01 Struct01
	// Case 1
	struct01.num = 1
	struct01.str1 = "str1 01"
	struct01.str2 = "str2 01"
	log.Printf("case 1 --- %v", struct01)
	// Case 2
	struct01 = Struct01{num: 2, str1: "str1 02", str2: "str2 02"}
	log.Printf("case 2 --- %v", struct01)
	// Case 3
	struct01 = Struct01{3, "str1 03", "str2 03"}
	log.Printf("case 3 --- %v", struct01)
}
