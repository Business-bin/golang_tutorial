package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

// sync.WaitGroup 사용시 main() 고루틴은 다른 고루틴이 끝날 때까지 대기하고,
// 등록한 모든 고루틴이 종료되면 main()고루틴도 종료됨
var wg sync.WaitGroup

func getWebPaceContent(url string, msg string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer wg.Done()
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(msg)
	return string(body)
}

func main() {
	/*
		wg.Add() : 실행할 고루틴의 수
		wg.Done() : 해당 고루틴 종료
		wg.Add() 에 넣은 고루틴 수와 wg.Done() 을 호출한 고루틴 수가 같아야 한다
		wg.Wait() : main()고루틴 이외에 다른 고루틴 끝날 때까지 대기
	*/
	wg.Add(3)
	go getWebPaceContent("https://example.com", "res1 KO") // go 키워드 붙임
	go getWebPaceContent("https://golang.org", "res2 KO")
	go getWebPaceContent("https://golang.org/doc", "res3 KO")

	wg.Wait()

}
