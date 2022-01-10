package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50;
	var remaingTickets uint = 50;
	fmt.Printf("Welcome to our %v booking application.\n",conferenceName)
	fmt.Printf("Total Tickets %v and Avaliable tickets %v \n", conferenceTickets, remaingTickets)
	fmt.Println("Get your tickets here!")
	// fmt.Printf("conferenceTickets Type %T \n",conferenceTickets)
	// fmt.Printf("%v \n",remaingTickets)
	// fmt.Println(&remaingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// Ask for user input
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName);
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName);
	fmt.Println("Enter your email:")
	fmt.Scan(&email);
	fmt.Println("Enter your no of tickets:")
	fmt.Scan(&userTickets);

	remaingTickets = remaingTickets - userTickets

	fmt.Printf("%v %v with %v booked %v tickets \n",firstName,lastName,email ,userTickets)
	fmt.Println("Thank you! You will recive the confirmation email.")
	fmt.Printf("Total Tickets %v and Avaliable tickets %v \n", conferenceTickets, remaingTickets)

}
