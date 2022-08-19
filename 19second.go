package main

import (
	"fmt"
	"time"
)

func main() {
	//miliSecond := 1645160831851 // 22-02-18 14:07:11
	//miliSecond := 1647932400000 // 22-02-18 14:07:11
	/*day := miliSecond / (1000 * 60 * 60 * 24)
	hour := (miliSecond / (1000 * 60 * 60)) % 24
	minute := (miliSecond / (1000 * 60)) % 60
	second := (miliSecond / (1000)) % 60

	fmt.Println(day, "일 ", hour, "시 ", minute, "분 ", second, "초")*/

	//nanoSecond := miliSecond * 1000000
	nanoSecond := 1657872180000000000
	//nanoSecond := 1657846340000000000

	// nanosecond 날짜 포맷으로 변경
	nowTimea := time.Unix(0, int64(nanoSecond))
	fmt.Println("==== nowTimea = ", nowTimea)

	// 문자 날짜포맷으로 변경
	nowTime := time.Date(2022, 04, 27, 11, 4, 57, 98, time.Local)
	fmt.Println("nowTime = ", nowTime)

	// 날짜 nanosecond 변경
	fmt.Println("나노초 = ", nowTime.UnixNano())
	// 날짜 milisecond 변경
	fmt.Println("밀리초 = ", nowTime.UnixNano()/1000000)
	fmt.Println()

	dateTime1 := time.Unix(0, int64(1656310770883*1000000))
	dateTime2 := time.Unix(0, int64(1656479785979*1000000))
	fmt.Println("dateTime1 = ", dateTime1)
	fmt.Println("dateTime2 = ", dateTime2)

	// 날짜 포맷 변경
	//formatNow := time.Now()
	customDate := nowTime.Format("2006-01-02")
	fmt.Println("customDate = ", customDate)
	fmt.Println("customDate = ", nowTime.Year())
	fmt.Println("customDate = ", nowTime.YearDay())

	//timeT, _ := time.Parse("2006-01-02", "2020-04-14")
	timeT, _ := time.Parse("2006-01-02 15:04:05", "2020-04-14 00:00:00")
	timeT2, _ := time.Parse("15:04:05", "00:00:00")

	fmt.Println("11111 timeT = ", timeT)
	fmt.Println("11111 timeT2 = ", timeT2)
	fmt.Println(timeT.Format("2006-01-02 15:04:05"))
	fmt.Println(timeT.Format("2006-01-02"))

	// 날짜 더하고 뺴기
	addDay := nowTime.AddDate(0, 0, 1).Format("2006-01-02")
	diffDay := nowTime.AddDate(0, 0, -1).Format("2006-01-02")
	fmt.Println("addDay = ", addDay)
	fmt.Println("diffDay = ", diffDay)

	testNowTime := time.Now()
	fmt.Println("testNowTime 111 ", testNowTime)
	fmt.Println("testNowTime 222 ", testNowTime.Format("2006-01-02"))
	fmt.Println("testNowTime 333 ", testNowTime.Format("15:04:05"))
	aa, _ := time.Parse("15:04:05", testNowTime.Format("15:04:05"))
	aa.Unix()
	fmt.Println("testNowTime 444 ", aa.Hour())
	fmt.Println("testNowTime 444 ", aa.Minute())
	fmt.Println("testNowTime 444 ", aa.Second())
	time222 := time.Date(0, 0, 0, aa.Hour(), aa.Minute(), aa.Second(), 00, time.Local)
	fmt.Println("time222 1111 ", time222)
	//startDateFormat, _ := time.Parse("2006-01-02", startDate)
	//endDateFormat, _ := time.Parse("2006-01-02", endDate)
	//days := int(endDateFormat.Sub(startDateFormat).Hours() / 24)

}
