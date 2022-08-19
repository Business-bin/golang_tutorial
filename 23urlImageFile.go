package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	fileName2, fullURLFile2 string
)

func main() {
	fullURLFile2 = "https://cdn.pixabay.com/photo/2022/03/30/14/01/rose-7101236__340.jpg"

	//fileName = "skyline-test.jpeg"

	fmt.Println("zzz")
	r, e := httpClient2().Get(fullURLFile2)
	fmt.Println("r : ", r)
	fmt.Println("e : ", e)
	if e != nil {
		panic(e)
	}
	defer r.Body.Close()
	buildFileName()
	// Create distination
	f, e := os.Create(fileName2) // "m1UIjW1.jpg"
	aa, _ := os.Readlink(fullURLFile2)
	fmt.Println("1111111111111 os.Readlink(fullURLFile2) = ", aa)
	fmt.Println("f : ", f)
	fmt.Println("e : ", e)
	os.Remove()
	if e != nil {
		panic(e)
	}
	defer f.Close()
	// Fill distination with content
	n, e := f.ReadFrom(r.Body)
	fmt.Println("n : ", n)
	fmt.Println("e : ", e)
	if e != nil {
		panic(e)
	}
	fmt.Println("File size: ", n)
}

func buildFileName() {
	fileUrl, e := url.Parse(fullURLFile2)
	if e != nil {
		panic(e)
	}

	path := fileUrl.Path
	segments := strings.Split(path, "/")

	fileName2 = segments[len(segments)-1]
	println(fileName2)
}

func httpClient2() *http.Client {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	return &client
}
