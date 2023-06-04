package main

import (
	"fmt"
	"time"
)

func SendTicket(userName  string , userEmail string , user_Ticket uint)  {
	fmt.Printf("Start sending tickets to user %v...\n", userName)	
	fmt.Println("----------------------------------------------------> Sending tickets...")	
	time.Sleep(20 * time.Second)
	fmt.Println("")	
	fmt.Println("************************************************************************************************")
	fmt.Println("")	
	fmt.Printf("Completed ticket bookings for User %v  \n",userName)
	fmt.Printf("Total bookings ticket is %v  \n",user_Ticket)
	fmt.Printf("Completed Sending to %v ...\n ",userEmail)
	fmt.Println("")
	fmt.Println("************************************************************************************************")
	wg.Done()
}