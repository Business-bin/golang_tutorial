package main

import "fmt"

/*
	- 구조체 임베딩은 클래스의 상속과 비슷
	- 변수명 없이 타입만 정의하면 임베딩 가능 ex) Node
*/
// ex1---------------

type Person struct {
	name string
	age  int
}
type Student struct {
	Person // 임베딩
	school string
	grade  int
}

func (p *Person) greeting() { // greeting 함수 Person구조체에 연결
	fmt.Println("Hello")
}
func (p *Student) greeting() { // 이름이 같은 메서드를 정의하면 오버라이딩 됨
	fmt.Println("Hello Students~")
}

// ------------------

// ex2---------------
// Note 예제에서 Blog
// Board 예제에서 Post
// Site 예제에서 TistoryBlog

type Note struct { // 예제에서 Blog
	num    int
	writer string
	board  []*Board
}
type Board struct { // 예제에서 Post
	title   string
	content string
}

func (n *Note) write(b *Board) {
	n.board = append(n.board, b)
}

type User struct {
	id    int
	email string
}

type Site struct {
	subscribers []*User // 구독자(이용자)
	Note                // 임베딩
}

func (s *Site) Subscribers() []*User {
	return s.subscribers
}
func (u *User) subscribe(s *Site) {
	s.subscribers = append(s.subscribers, u)
}

// ------------------

func main() {
	// ex1---------------
	var s Student
	s.greeting() // 오버라이딩된 greeting()이 없으면 Hello 찍힘
	s.Person.greeting()
	// ------------------

	// ex2---------------
	site := Site{
		[]*User{},
		Note{3, "3 bin", []*Board{}},
	}
	board := Board{"3 title", "3 content"}
	site.write(&board) // 임베딩 했으므로 Note와 연결된 write() 사용 가능

	user := User{1, "bin@aaa.aaa"}
	user.subscribe(&site)

	for _, v2 := range site.Subscribers() {
		fmt.Println("v2 = ", v2)
	}
	// ------------------
}
