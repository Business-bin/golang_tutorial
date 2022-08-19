package main

import (
	"fmt"
	"unsafe"
)

type TestBsb1 struct {
	Idx     int32  `gorm:"AUTO_INCREMENT;primary_key;column:idx" json:"idx"`
	Col1    string `gorm:"column:col1" json:"col1"`
	Col2    string `gorm:"column:col2" json:"col2"`
	Col3    int32  `gorm:"column:col3" json:"col3"`
	Col4    string `gorm:"column:col4" json:"col4"`
	Col5    string `gorm:"column:col5" json:"col5"`
	RegTime int64  `gorm:"column:reg_time" json:"regTime"`
}

type TestBsb2 struct {
	Idx     int32  `gorm:"AUTO_INCREMENT;primary_key;column:idx" json:"idx"`
	Col3    int32  `gorm:"column:col3" json:"col3"`
	RegTime int64  `gorm:"column:reg_time" json:"regTime"`
	Col1    string `gorm:"column:col1" json:"col1"`
	Col2    string `gorm:"column:col2" json:"col2"`
	Col4    string `gorm:"column:col4" json:"col4"`
	Col5    string `gorm:"column:col5" json:"col5"`
}

type TestBsbInt8 struct {
	Col1 int8 `gorm:"column:col1" json:"col1"`
}

type TestBsbInt16 struct {
	Col1 int16 `gorm:"column:col1" json:"col1"`
}

type TestBsbInt32 struct {
	Col1 int32 `gorm:"column:col1" json:"col1"`
}

type TestBsbInt64 struct {
	Col1 int64 `gorm:"column:col1" json:"col1"`
}

type TestBsbInt struct {
	Col1 int `gorm:"column:col1" json:"col1"`
}

type TestBsbBool struct {
	Col1 bool `gorm:"column:col1" json:"col1"`
}

type TestBsbString struct {
	Col1 string `gorm:"column:col1" json:"col1"`
}

type TestBsbInterface struct {
	Col1 interface{} `gorm:"column:col1" json:"col1"`
}

type TestBsbArrayInterface struct {
	Col1 []interface{} `gorm:"column:col1" json:"col1"`
}

type TestBsbMap struct {
	Col1 map[interface{}]interface{} `gorm:"column:col1" json:"col1"`
}

type TestBsbStringMap struct {
	Col1 map[string]string `gorm:"column:col1" json:"col1"`
}

func main() {
	t1 := TestBsb1{}
	t2 := TestBsb2{}

	fmt.Println(unsafe.Sizeof(t1))
	fmt.Println(unsafe.Sizeof(t2))
	fmt.Println()

	testBsbInt8 := TestBsbInt8{}
	testBsbInt16 := TestBsbInt16{}
	testBsbInt32 := TestBsbInt32{}
	testBsbInt64 := TestBsbInt64{}
	testBsbInt := TestBsbInt{}
	testBsbBool := TestBsbBool{}
	testBsbString := TestBsbString{}
	testBsbInterface := TestBsbInterface{}
	testBsbArrayInterface := TestBsbArrayInterface{}
	testBsbMap := TestBsbMap{}
	testBsbStringMap := TestBsbStringMap{}

	fmt.Println("testBsbInt8 = ", unsafe.Sizeof(testBsbInt8))
	fmt.Println("testBsbInt16 = ", unsafe.Sizeof(testBsbInt16))
	fmt.Println("testBsbInt32 = ", unsafe.Sizeof(testBsbInt32))
	fmt.Println("testBsbInt64 = ", unsafe.Sizeof(testBsbInt64))
	fmt.Println("testBsbInt = ", unsafe.Sizeof(testBsbInt))
	fmt.Println("testBsbBool = ", unsafe.Sizeof(testBsbBool))
	fmt.Println("testBsbString = ", unsafe.Sizeof(testBsbString))
	fmt.Println("testBsbInterface = ", unsafe.Sizeof(testBsbInterface))
	fmt.Println("testBsbArrayInterface = ", unsafe.Sizeof(testBsbArrayInterface))
	fmt.Println("testBsbMap = ", unsafe.Sizeof(testBsbMap))
	fmt.Println("testBsbStringMap = ", unsafe.Sizeof(testBsbStringMap))
}
