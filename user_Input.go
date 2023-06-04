package main

import "fmt"


func User_Input() (string ,string , uint)  {

	var userName string
	var userEmail string
	var user_Ticket uint

	fmt.Println("")
	fmt.Printf("Enter Your Name : ")
    fmt.Scan(&userName)	

	fmt.Println("")
	fmt.Printf("Enter Your Email : ")
    fmt.Scan(&userEmail)	

	fmt.Println("")
	fmt.Printf("Enter Your Ticket : ")
    fmt.Scan(&user_Ticket)

	return userName , userEmail , user_Ticket
}