package main

const PI = 3.14              // 타입이 없는 상수
const FloatPI float64 = 3.14 // float64 타입 상수

func main() {
	var a int = PI * 100      // 오류 발생 x
	var b int = FloatPI * 100 // Error
}
