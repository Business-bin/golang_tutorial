package main

import "fmt"

// 인터페이스 정의
type Develop interface {
	Coding()
}

type Language struct {
	languages []string
}

func (l Language) Coding() {
	for _, lan := range l.languages {
		fmt.Println(lan)
	}
}

func Work(dev Develop) {
	dev.Coding()
}

func main() {
	l := Language{[]string{"java", "javascript", "golang"}}
	//l.Coding()
	Work(l)
}
