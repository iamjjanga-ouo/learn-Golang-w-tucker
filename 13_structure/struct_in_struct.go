package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"송하나", "hah", 23}
	vip := VIPUser{
		User{"화랑", "hwarang", 40},
		3,
		250,
	}

	fmt.Println(user.Name)
	fmt.Println(vip.UserInfo.Name)
}
