package main

import "fmt"

func Check_Ticket_Amount(remainingTickets uint)  {
	fmt.Println("")
	fmt.Println("-------------------------------------------------")
	fmt.Println("Soorty We Don't Have That Amount of Ticket Now ")
	fmt.Printf("We Have Avilable : %v \n",remainingTickets)
	fmt.Println("-------------------------------------------------")
	fmt.Println("")
}