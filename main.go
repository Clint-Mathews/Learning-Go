package main

import (
	"fmt"
	"strings"
)

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

	for{
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// var bookings []string
	// var bookings =  []string{}
	bookings := []string{}

	// Ask for user input
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName);
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName);
	fmt.Println("Enter your email:")
	fmt.Scan(&email);
	fmt.Println("Enter your no of tickets:")
	fmt.Scan(&userTickets);

	isValidName :=len(firstName) >= 2 && len(lastName) >=2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketCount := userTickets > 0 && userTickets <= remaingTickets

	if isValidName && isValidEmail && isValidTicketCount  {

	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName + " " + lastName)
	remaingTickets = remaingTickets - userTickets

	// fmt.Printf("Whole Slice %v \n", bookings)
	// fmt.Printf("Bookings 1st value %v \n", bookings[0])
	// fmt.Printf("Bookings type %T \n", bookings)
	// fmt.Printf("Bookings size %v \n", len(bookings))

	fmt.Printf("%v %v with %v booked %v tickets \n",firstName,lastName,email ,userTickets)
	fmt.Println("Thank you! You will recive the confirmation email.")
	fmt.Printf("Total Tickets %v and Avaliable tickets %v \n", conferenceTickets, remaingTickets)

	firstNames := []string{}

	for _,name := range bookings {
		var firstName = strings.Fields(name)
		firstNames = append(firstNames,firstName[0]) 
	}

	fmt.Printf("All FirstName of Bookings : %v \n",firstNames)
	if remaingTickets <= 0 {
		// Breaks the loop
		fmt.Println("Conference booked out!")
		break
	}
// } else if userTickets == remaingTickets {

// 	// Do something
// } 
} else{
	if !isValidName {
		fmt.Println("First Name or Last Name is too short")
	}
	if !isValidEmail {
		fmt.Println("Email format is wrong")
	}
	if !isValidTicketCount {
		fmt.Println("Requested ticket number is wrong")
	}
			// fmt.Println("Invalid Input Data")
		// continue
}
// city := "London"
// switch city {
// 	case "London":
// 		//Execute on London city 
// 	case "Kochi","Aluva":
// 		//Execute on Kochi and Aluva
// 	default:
// 		// Default block
// 		fmt.Println("Invalid city")
// }
}
}
