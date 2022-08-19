package main

import "fmt"

func main() {
	calArr([]int{1, 3, 24, 26, 47, 49, 68, 70, 91, 93}, "kiwi")
	calArr([]int{2, 21, 23, 42, 44, 46, 65, 67, 69, 88}, "pear")
	calArr([]int{4, 6, 25, 29, 48, 50, 71, 73, 92, 94, 96}, "banana")
	calArr([]int{5, 7, 28, 30, 32, 51, 53, 74, 76, 95, 97}, "melon")
	calArr([]int{8, 10, 12, 31, 33, 52, 56, 75, 77, 79, 98, 100}, "pineapple")
	calArr([]int{9, 18, 27, 36, 45, 54, 63, 72, 81, 90, 99}, "apple")
	calArr([]int{11, 13, 34, 55, 57, 59, 78, 80}, "cucumber")
	calArr([]int{14, 16, 35, 37, 39, 58, 60, 83}, "orange")
	calArr([]int{15, 17, 19, 38, 40, 61, 82, 84, 86}, "grape")
	calArr([]int{20, 22, 41, 43, 62, 64, 66, 85, 87, 89}, "cherry")
}

func calArr(arr []int, plan string) []interface{} {
	resultArr := make([]interface{}, 0)
	for i, v := range arr {
		if i == len(arr)-1 {
			break
		}
		resultArr = append(resultArr, arr[i+1]-v)
	}
	fmt.Println(plan, " = ", resultArr, " 길이 = ", len(arr))
	return resultArr
}
