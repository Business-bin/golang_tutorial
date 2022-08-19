package main

import (
	"fmt"
	blog "go-test/example.com/encapsulation"
)

func main() {
	b := blog.NewBlog(1, "binBlog", "http://", []*blog.Post{}) // 구조체에 값 할당
	helloPost := blog.NewPost("go hello go", "this is golang") // 구조체에 값 할당

	helloPost.SetTitle("GO : Hello, Go!") // setter, Post구조체 title에 값 새로 할당
	b.Wirte(helloPost)

	for _, post := range b.Posts() { // b.Posts() <-- getter Blog구조체의 posts값 읽어옴
		fmt.Println(*post)
	}
}
