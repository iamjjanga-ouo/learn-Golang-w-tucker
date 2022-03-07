package main

import "fmt"

func main() {
	slice1 := []int{1,2,3,4,5}
	slice2 := make([]int, len(slice1))  // slice1과 같은 길이의 슬라이스 생성

	for i,v := range slice1 {			// 값 복사
		slice2[i] = v
	}

	slice1[1] = 100
	fmt.Println(slice1)	// [1 100 3 4 5]
	fmt.Println(slice2) // [1 2 3 4 5]

	/* append() 코드로 개선 */
	slice3 := append([]int{}, slice1...)
	fmt.Println(slice3) // [1 100 3 4 5]

	/* copy() 코드로 개선 */
	slice4 := make([]int, 3, 10) // len:3, cap:10
	slice5 := make([]int, 10) // len:10, cap:10

	// copy(dst, src []Type) int : dst와 src 중 슬라이스 길이가 더 작은 개수만큼 복사됨.
	copy1 := copy(slice4, slice1)
	copy2 := copy(slice5, slice1)

	fmt.Println(copy1, slice4) // 3 [1 100 3]
	fmt.Println(copy2, slice5) // 5 [1 100 3 4 5 0 0 0 0 0]
}