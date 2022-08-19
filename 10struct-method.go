package main

import "fmt"

/*
	- 클래스에서 매소드 만들 듯 해당 구조체에 속한 함수 만들 수 있으며, 그것도 메소드라 함
	- 일반적인 파라미터 외에 리시버 파라미터 사용
*/

type Note1 struct { // 예제에서 Blog
	num    int
	writer string
	board  []*Board1
}
type Board1 struct { // 예제에서 Post
	title   string
	content string
}

func main() {
	note := Note1{1, "bin", []*Board1{}}

	note.write(&Board1{"1 title", "1 content"})
	note.write(&Board1{"2 title", "2 content"})

	//fmt.Printf("%#v", note.Boards())
	for _, v := range note.board {
		fmt.Println("v = ", v.title)
	}
}

func (n *Note1) write(b *Board1) {
	n.board = append(n.board, b)
}

func (n *Note1) Boards() []*Board1 {
	//fmt.Println("n.board = ", n.board)
	return n.board
}
