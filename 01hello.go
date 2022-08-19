// main = 실행 가능한 소스코드 패키지
package main

import (
	"fmt"
)

// message6 := "" <<< error, 함수 외부에서는 모든 문장이 키워드로 시작해야 함

// go가 실행되면서 가장 먼저 실행되는 함수
func main() {
	fmt.Println("Hello, Go!")
}
