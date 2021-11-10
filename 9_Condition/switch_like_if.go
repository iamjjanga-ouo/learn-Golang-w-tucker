package main

import "fmt"

func main() {
	temp := 18

	switch true {
	case temp < 10, temp > 30:
		fmt.Println("바깥 활동하기 좋은 날씨")
	case temp >= 10 && temp < 20:
		fmt.Println("약간 추울 지도..")
	case temp > 15 && temp < 25:
		fmt.Println("야외 활동에 좋을듯.")
	default:
		fmt.Print("따뜻합니다.")
	}
}
