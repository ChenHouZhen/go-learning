package main

import "fmt"

type User struct {
	name	string
}

func main() {
	u := User{name:"小明"}
	fmt.Printf("fun main :%+v, %p\n", u, &u)
	object(u)
	fmt.Printf("fun main :%+v, %p\n", u, &u)
	pointer(&u)
	fmt.Printf("fun main :%+v, %p\n", u, &u)
}

func object(objUser User)  {
	objUser.name = "小红"
	fmt.Printf("fun object :%+v, %p\n", objUser, &objUser)
}

func pointer(ptrUser *User)  {
	ptrUser.name = "小蓝"
	fmt.Printf("fun object :%+v, %p\n", ptrUser, ptrUser)
}
