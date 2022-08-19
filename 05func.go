package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func main() {
	hello, err := sayHello01("Hello, Go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hello)

	// 이름이 있는 리턴 값
	hello02 := sayHello02()
	fmt.Println(hello02, "\n")

	// 익명함수
	// (1) p 변수에 익명함수 저장
	p := func(v string) {
		fmt.Println(v)
	}
	// (2) each함수 호출하며 배열과 함께 p 넘김
	each([]string{"Hello", "Go", "lang"}, p)

	/*
		defer
		- 함수가 종료될 때 호출 됨
		- 런타임 중 에러가 발생해도 defer키워드를 사용한 함수는 호출 보장 됨
		- 여러개 사용 가능하며, 스택스로 쌓이기 때문에 나중에 설정된 defer부터 실행 됨
	*/
	fmt.Println("------------")
	file, err2 := os.Open("./go.txt")
	if err2 != nil {
		log.Fatal(err2)
	}
	defer file.Close()
	fmt.Println(file.Name())

	/*
		go & go-routin
		- go 키워드는 동시성을 쓸 때 사용
		- Go언어의 동시성은 기본적으로 OS스레드를 한층 더 추상화한 코루틴(co-routin)이며
			비동기적 프로그래밍을 할 수 있게 해준다. Go에서는 이를 고루틴 이라 한다
	*/
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("i = ", i)
			time.Sleep(time.Second)
		}
	}()
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		for _, byte := range "Hello, Go!" {
			fmt.Printf("%c\n", byte)
			time.Sleep(time.Second)
		}
	}()
	fmt.Println("111111111111")
	wg1.Wait()
	fmt.Println("222222222222")
}

// Go함수의 구성요소 - func키워드 함수명 받는값 리턴값(리턴값으로 여러개의 값 가능)
// msg라는 string타입의 파라매터와 string, error타입의 리턴값
func sayHello01(msg string) (string, error) {
	return msg, nil
}

// 이름이 있는 리턴 값
func sayHello02() (hello string) {
	//hello02 = "Hello, Go 2"
	hello = "Hello, Go 2"
	return
}

// 익명함수
// 일급함수 : 함수를 값으로써 여기는 함수
// 고차함수 : 함수를 리턴하거나, 파라매터로 받거나 리턴하는 함수
func each(arr []string, iterator func(string)) {
	// (3) 배열길이만큼 반복하며 (1)에서 정의한 익명함수 실행
	for _, v := range arr {
		iterator(v)
	}
}
