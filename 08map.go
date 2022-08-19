package main

import "fmt"

// Go는 제네릭이 없다(하나의 값이 여러 데이터타입을 담을 수 없음)
func main() {
	// 맵 생성시 make()를 통해 생성
	var m1 map[string]string = make(map[string]string)
	fmt.Printf("%v\n", m1)

	// make()사용 없이 리터럴로 생성 가능
	m2 := map[string]string{"aaa": "111", "bbb": "222"}
	fmt.Printf("%v\n", m2)
	fmt.Printf("%v\n", m2["bbb"])
	// 값을 할당하지 않은 키에 접근해도 해당 값의 타입에 해당하는 제로값이 할당되기때문에 에러 안남
	fmt.Printf("%#v\n", m2["java"])

	m2["ccc"] = "333" // <-- 맵에 추가
	fmt.Printf("%v\n", m2)
	delete(m2, "ccc") // <-- 맵에서 제거
	fmt.Printf("%v\n", m2)
	// 키 존재여부 확인. 해당 키 존재하면 ok에 true 할당
	if a, ok := m2["bbb"]; ok {
		fmt.Println("aaa = ", a)
		fmt.Println(ok)
	}
	// 맵 순회
	for k, v := range m2 {
		fmt.Printf("key = %s , value = %s \n", k, v)
	}

}
