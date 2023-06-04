package controller

import (
	"fmt"
	
)



func System_Information(confrenceTickets uint , remainingTickets uint , name string , booking_user any ){
    fmt.Println("")
	fmt.Printf("Total BookingUser : %v  \n",booking_user)
	//fmt.Printf("ConferenceTickets is %T, remainingTickets is %T, ConferenceName is %T\n   ",confrenceTickets ,remainingTickets , userName)
    fmt.Println("")
	fmt.Printf("Welcome to our %v booking application\n", name  )
	fmt.Println("")
	fmt.Printf("We have total of %v Tickes and %v are staill avabled.\n",confrenceTickets , remainingTickets )
	fmt.Println("")
}