package main

import "fmt"

type User struct {
	Name  string
	ID    string
	Age   int
	Level int
}

type VIPUser struct {
	User
	Price int
	Level int
}

func main() {
	user := User{"송하나", "hah", 23, 10}
	vip := VIPUser{
		User{"화랑", "hwarang", 40, 10},
		250,
		100,
	}

	fmt.Println(vip.Level)      // 100
	fmt.Println(vip.User.Level) // 10
}
