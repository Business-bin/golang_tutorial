package main

import "fmt"

func main() {
	// 정수형 채널을 생성한다
	//ch := make(chan int)
	//go func() {
	//	ch <- 123 //채널에 123을 보낸다
	//}()
	//var i int
	//i = <-ch // 채널로부터 123을 받는다
	//fmt.Println(i)
	// -----------------------------------------------------
	//c := make(chan int)
	//c <- 1 //수신루틴이 없으므로 데드락
	//d := <-c
	//fmt.Println(d) //코멘트해도 데드락 (별도의 Go루틴없기 때문)
	// -----------------------------------------------------
	ch := make(chan int, 1)
	//수신자가 없더라도 보낼 수 있다.
	ch <- 101
	fmt.Println(<-ch)
	done := make(chan bool, 1)
	//수신자가 없더라도 보낼 수 있다.

	fmt.Println(<-done)
	fmt.Println("111")
}
