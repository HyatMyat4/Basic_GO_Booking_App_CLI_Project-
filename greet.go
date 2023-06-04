package main

import (
	"fmt"
	"strings"
)


func Greet(userName  string , userEmail string , user_Ticket uint )  {
	fmt.Println("")
	fmt.Println("-------------------------------------------------")
	fmt.Println("Soorty We can't save You information .")
	fmt.Println("")
	fmt.Println("Please valid information!")
	if len(userName) <= 2 {
		fmt.Println("")
		fmt.Println("Please fill valide Name information! ")
	}else if user_Ticket <= 0 {
		fmt.Println("")
		fmt.Println("Please fill valid Ticket value ")
	} else if !strings.Contains(userEmail,"@gmail.com") {
		fmt.Println("")
		fmt.Println("Please fill valide Email information!")
	}
	fmt.Println("-------------------------------------------------")
	fmt.Println("")
}