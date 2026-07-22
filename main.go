package main

import "fmt"

type User struct {
	Name string // property or member variable
	Age  int
}

func printUserDetails(usr User) {
	fmt.Println("Name", usr.Name)
	fmt.Println("Age", usr.Age)
}

func main() {
	var user1 User

	user1 = User{
		Name: "Habib",
		Age:  30,
	}

	printUserDetails(user1)

	user2 := User{ // instance or object
		Name: "User second",
		Age:  50,
	}

	printUserDetails(user2)

}

func init() {
	fmt.Println("This will be invoked first")
}
