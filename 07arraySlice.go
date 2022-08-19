package main

import (
	"fmt"
)

func main() {
	// 1 배열(array)
	// Go의 배열은 크기가 정해져 있는 정적 배열임
	numbers1 := [5]int{0, 1, 2, 3, 4}
	fmt.Println(numbers1, "\n")

	// rune -  rune타입으로 정의하면 UTF-8 인코딩으로 처리
	runes := [5]rune{
		'A',
		'B',
		'C',
		'D',
		'E',
	}
	for _, char := range runes {
		fmt.Println(string(char))
	}

	// 2 슬라이스(slice)
	// 슬라이스는 크기가 정해지지 않은 배열. 길이가 동적으로 변할 수 있다
	// 슬라이스를 생성하는 작업이 필요. make()비트인 합수를 사용
	slicenums1 := make([]int, 5)       // <-- 5개의 공간을 미리 확보한 슬라이스 생성
	slicenums1[0] = 1                  // <-- slicenums의 첫번째 원소 값 변경
	slicenums1 = append(slicenums1, 6) // <-- slicenums에 원소 추가
	fmt.Println("slicenums1 = ", slicenums1)

	// make() 사용하지지 않고 리터럴의 형태로 사용 가능
	slicenums2 := []int{1, 2, 3, 4, 5}
	fmt.Println(slicenums2)

	// make()를 사용하지 않고 타입만 있는 형태의 변수로 선언한 슬라이스의 제로값은 nil임
	var slicenums3 []int
	fmt.Println(slicenums3 == nil)

	// 이미 생성된 배열을 기반으로 슬라이스 생성 가능
	numbers2 := [5]int{0, 1, 2, 3, 4}
	slicenums4 := numbers2[1:4] // [1:] - 인덱스 1부터 끝, [:4] 처음부터 인덱스3, [:] 처음부터 끝
	fmt.Println("slicenums4-1 = ", slicenums4)
	// 위처럼 슬라이스 생성하면 슬라이스 값 변경시 배열도 변경, 배열 값 변경시 슬라이스 값도 변경된다
	slicenums4[1] = 20
	fmt.Println("slicenums4-2 = ", slicenums4)

	slicenums4 = append(slicenums4, 40) // <-- 배열의 길이와 같을때 까지 동기화 되는 듯(슬라이스는 배열의 일부분 보여주는 것)
	fmt.Println("slicenums4-3 = ", slicenums4)
	fmt.Println("numbers2 = ", numbers2)

	slicenums4 = append(slicenums4, 50) // <-- 배열의 길이 초과해서 값 넣으면 슬라이스에만 할당 됨
	fmt.Println("slicenums4-4 = ", slicenums4)
	fmt.Println("numbers2 = ", numbers2)

	// struct -> map 변경
	/*searchMap := make(map[int]map[string]interface{})

	// 배열에서 순서지키며 요소 삭제
	//list = append(list[:i], list[i+1:]...)

	items := reflect.ValueOf(searchList)
	if items.Kind() == reflect.Slice {
		for i := 0; i < items.Len(); i++ {
			item := items.Index(i)
			if searchMap[i] == nil {
				searchMap[i] = make(map[string]interface{})
			}
			if item.Kind() == reflect.Struct {
				v := reflect.Indirect(item)

				for j := 0; j < v.NumField(); j++ {
					tag := v.Type().Field(j).Tag.Get("json")
					searchMap[i][tag] = v.Field(j).Interface()
				}
			}
		}
	}*/

}
