package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	fileName, fullURLFile string
)
var basePaht string

func main() {
	//fullURLFile = "https://storage.googleapis.com/krms-firebase-bucket/model/EVO3000GW.png"
	//fullURLFile = "https://storage.googleapis.com/krms-firebase-bucket/app/img_6.jpg"
	fullURLFile = "https://cdn.pixabay.com/photo/2020/09/16/04/02/skyline-5575251_960_720.jpg"

	fileName = "skyline-test.jpeg"

	// Create blank file
	//file := createFile()

	// Put content on file
	//putFile(file, httpClient())
	readFile(httpClient())
}

func createFile() *os.File {
	basePaht = "D:\\"
	file, err := os.Create(basePaht + fileName)

	checkError(err)
	return file
}

func putFile(file *os.File, client *http.Client) {
	resp, err := client.Get(fullURLFile)

	checkError(err)

	defer resp.Body.Close()
	fmt.Println()
	fmt.Println("resp Header = ", resp.Header)
	fmt.Println("resp Header = ", resp.Header["Content-Type"])
	fmt.Println()
	fmt.Println("resp Body = ", resp.Body)
	fmt.Println()
	fmt.Println("resp  = ", resp)
	fmt.Println()
	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	checkError(err)

	fmt.Printf("Just Downloaded a file %s with size %d\n", fileName, size)
}

func readFile(client *http.Client) {
	resp, err := client.Get(fullURLFile)

	checkError(err)
	defer resp.Body.Close()
	fmt.Println()
	fmt.Println("resp Header = ", resp.Header)
	fmt.Println("resp Header = ", resp.Header["Content-Type"])
	fmt.Println()
	fmt.Println("resp Body = ", resp.Body)
	fmt.Println()
	fmt.Println("resp  = ", resp)
	fmt.Println()

}

func httpClient() *http.Client {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	return &client
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
