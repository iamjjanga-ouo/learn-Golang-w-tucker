package main

func main() {
	slice1 := []int{1,2,3,4,5,6}
	idx := 2 // 삭제할 인덱스

	for i := idx+1; i < len(slice1); i++ { // 요소 앞당기기
		slice1[i-1] = slice1[i]
	}

	/* append 함수로 개선 */
	slice1 = append(slice1[:idx], slice1[idx+1:]...)
}