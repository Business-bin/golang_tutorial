package hello

import "fmt"

/*
	go mod init 모듈명 (여기서는 go-mod)
	go.mod에 다른 모듈들을 import할 수 있고, import하게 되면 해당 모듈들을 프로젝트 내에서 자유롭게 사용 할 수 있게 된다
*/
/*
	접근 제어
	- 식별자가 대문자로 시작하면 패키지 외부에서 접근 가능
	- 식별자가 소문자로 시작하면 패키지 외부에서 접근 불가능
*/
func SayHello() {
	fmt.Println("Hello, Go!!!")
}
