package main

import (
	"fmt"
)

type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[string]string)
	m["이화랑"] = "서울시 광진구"
	m["송하나"] = "서울시 강남구"
	m["백두산"] = "부산시 사하구"
	m["최번개"] = "전주시 덕진구"

	m["최번개"] = "전주시 상담구" // 값 변경

	fmt.Printf("송하나의 주소는 %s입니다.\n", m["송하나"])
	fmt.Printf("최번개의 주소는 %s입니다.\n", m["최번개"])

	fmt.Println()

	// map 순회
	mp := make(map[int]Product)
	mp[16] = Product{"볼펜", 500}
	mp[46] = Product{"지우개", 200}
	mp[78] = Product{"자", 1200}
	mp[345] = Product{"샤프", 3000}
	mp[897] = Product{"샤프심", 500}

	for k, v := range mp {
		fmt.Println(k, v)
	}

	fmt.Println()

	// 요소 삭제와 없는 요소 확인
	// delete(m: 맵변수, key: 삭제키)
	delete(mp, 16)
	delete(mp, 46)
	fmt.Println(mp[16]) // 삭제된 요소 출력
	v, ok := mp[46]
	fmt.Println("delete(m, key)에 의해 삭제된 map인지 확인 ok = false", v, ok)
	for k, v := range mp {
		fmt.Println(k, v)
	}
}
