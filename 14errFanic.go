package main

import (
	"fmt"
	"log"
	"os"
)

/*
	- Go에는 예외(Exception)가 없다
	- 일반적인 에러로는 프로그램이 종료되지 않는다
	- 패닉은 런타임중 발생하는 치명적인 에러로, 패닉이 발생하면 프로그램이 종료된다
*/

func main() {
	// 1 error
	_, err := os.Open("go1.txt") // 에러 발생시 err에 error타입의 값 할당
	if err != nil {
		log.Fatalln(err)
	}

	//2
	//fanic
	// 패닉 발생시 프로그램이 종료되는것을 방지하기 위해 정의
	/*defer func() {
		fmt.Println("defer")
		rec := recover()
		if rec != nil {
			fmt.Println("rec = ", rec)
		}
	}()*/
	// 정말 예상하지 못한 에러에 대해서는 패닉을 제대로 발생시켜줄 필요가 있다.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ERROR ", r)
		}
	}()
	fmt.Println("START")
	num := []int{1, 2}
	for i := 0; i <= len(num); i++ {
		fmt.Println(num[i])
		//panic("PANIC")	// <-- 패닉 발생 시키기
	}
	fmt.Println("END", len(num))

}
