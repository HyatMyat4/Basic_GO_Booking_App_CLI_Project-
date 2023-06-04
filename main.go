package main  

import (
	"fmt"
	"strings"	
	"booking-app/controller"
	"sync"
)
var wg = sync.WaitGroup{}

var name = "Go Conference"

const confrenceTickets = 50

var remainingTickets uint = 50

var booking_user = make([]User_data,0)

type User_data struct {
	userName    string
	userEmail   string
	user_Ticket uint
}

func main() {	    

	controller.System_Information(confrenceTickets ,remainingTickets , name , booking_user)

  for {     
	if remainingTickets == 0  {
		fmt.Println("Soorty We Don't Ticket Now , Come Back Next Week... ")
		break 
	}
    fmt.Println("")	
	fmt.Println("")	
	fmt.Println("Get your tickets here to attend Here 💻 ")	

	userName , userEmail , user_Ticket  := User_Input()

	 var userdata = User_data {
		userName    : userName,
		userEmail   : userEmail,
		user_Ticket : user_Ticket,
	}
	
	fmt.Println("")
	fmt.Printf("User %v Email %v booked %v Tickets.\n",userName , userEmail , user_Ticket)
	
	if len(userName) <= 2  || user_Ticket <= 0 || !strings.Contains(userEmail,"@gmail.com")  {
	    Greet(userName , userEmail , user_Ticket)
		continue
	}

	if user_Ticket > remainingTickets {
		Check_Ticket_Amount(remainingTickets)
		continue
	}

	remainingTickets = remainingTickets - user_Ticket
    
	booking_user = append(booking_user,userdata)

	controller.Print_booking_user(booking_user)
	
    Print_detail(remainingTickets , userName , userEmail , user_Ticket)
    
	wg.Add(1) //How much Muti Threads
	go SendTicket(userName , userEmail , user_Ticket)
}
    wg.Wait()
}








