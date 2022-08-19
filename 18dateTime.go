package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	now := time.Now()
	nowUTC := time.Now().UTC()
	nowUNIX := time.Now().Unix()
	nowMili := time.Now().UnixNano() / 1000000
	//nowMili := 1649210942878389200 / 1000000
	nowNano := time.Now().UnixNano()
	fmt.Println(now)
	fmt.Println(nowUTC)
	fmt.Println(nowUNIX)
	fmt.Println("nowMili = ", nowMili)
	fmt.Println("nowNano = ", nowNano)

	//now := time.Now().UnixNano() / 1000000
	//now := time.Now().UnixNano()
	//var aa int64
	//aa = 1.6466343892791263e+18 // 1646634389279126272 1646634389279126300 1646619146559456468
	//var bb int64
	//bb = 1646619146559456468
	dateModifiedF, _ := strconv.ParseFloat("1.6466343892791263e+18", 64)
	dateModifiedI, _ := strconv.Atoi("1.6466343892791263e+18")
	fmt.Println("dateModifiedF = ", dateModifiedF)
	fmt.Println("dateModifiedI = ", dateModifiedI)

	//nowTimea := time.Unix(0, aa)
	//nowTimeb := time.Unix(0, bb)
	//nowTimetest1 := time.Unix(0, 1646634389279126300)
	//nowTimetest2 := time.Unix(0, 1646634707101703700)
	//fmt.Println("nowTimea = ", nowTimea)
	//fmt.Println("nowTimeb = ", nowTimeb)
	//fmt.Println("nowTimetest = ", nowTimetest1)
	//fmt.Println("nowTimetest = ", nowTimetest2)
}
