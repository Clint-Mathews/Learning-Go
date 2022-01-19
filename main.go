package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint

}

const conferenceTickets int = 50;
var conferenceName string= "Go Conference"
var remaingTickets uint = 50;
// var bookings []string = []string{}
// var bookings = make([]map[string]string, 0);
var bookings = make([]UserData, 0);

func main() {
	greatUsers();
	// fmt.Printf("conferenceTickets Type %T \n",conferenceTickets)
	// fmt.Printf("%v \n",remaingTickets)
	// fmt.Println(&remaingTickets)

	for{
	// var firstName string
	// var lastName string
	// var email string
	// var userTickets uint
	// // var bookings []string
	// // var bookings =  []string{}
	// // Ask for user input
	// fmt.Println("Enter your first name:")
	// fmt.Scan(&firstName);
	// fmt.Println("Enter your last name:")
	// fmt.Scan(&lastName);
	// fmt.Println("Enter your email:")
	// fmt.Scan(&email);
	// fmt.Println("Enter your no of tickets:")
	// fmt.Scan(&userTickets);
	firstName, lastName, email,userTickets := getUserInput()
	isValidUser, isValidEmail, isVaildTickets := helper.ValidateUser(firstName, lastName, email, userTickets, remaingTickets)


	if  isValidUser && isValidEmail && isVaildTickets{

	// bookings[0] = firstName + " " + lastName
    bookings,remaingTickets := bookTicket(firstName, lastName, userTickets , email )
	go sendTicket(userTickets, firstName, lastName, email )
	// fmt.Printf("Whole Slice %v \n", bookings)
	// fmt.Printf("Bookings 1st value %v \n", bookings[0])
	// fmt.Printf("Bookings type %T \n", bookings)
	// fmt.Printf("Bookings size %v \n", len(bookings))

	// fmt.Printf("%v %v with %v booked %v tickets \n",firstName,lastName,email ,userTickets)
	// fmt.Println("Thank you! You will recive the confirmation email.")
	// fmt.Printf("Total Tickets %v and Avaliable tickets %v \n", conferenceTickets, remaingTickets)
	firstNames := getFirstNames(bookings)
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
	if !helper.ValidateUserInput(firstName,lastName) {
		fmt.Println("First Name or Last Name is too short")
	}
	if !helper.ValidateUserEmail(email) {
		fmt.Println("Email format is wrong")
	}
	if !helper.ValidateUserTicketCount(userTickets, remaingTickets) {
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

func greatUsers() {
	fmt.Printf("Welcome to our %v booking application.\n",conferenceName)
	fmt.Printf("Total Tickets %v and Avaliable tickets %v \n", conferenceTickets, remaingTickets)
	fmt.Println("Get your tickets here!")
}


func getFirstNames(bookings []UserData) []string {
		firstNames := []string{}

	for _,name := range bookings {
		// var firstName = strings.Fields(name["firstName"])
		// firstNames = append(firstNames,firstName[0]) 
		// firstNames = append(firstNames,name["firstName"]) 

		firstNames = append(firstNames,name.firstName) 
	}
	return firstNames;
}


func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName);
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName);
	fmt.Println("Enter your email:")
	fmt.Scan(&email);
	fmt.Println("Enter your no of tickets:")
	fmt.Scan(&userTickets);
	return firstName, lastName, email, userTickets
}

func bookTicket (firstName string, lastName string, userTickets uint, email string) ([]UserData, uint) {
	
	// create map for user tickets
	// 	var userData = make(map[string]string);
	// userData["firstName"] = firstName;
	// userData["lastName"] = lastName;
	// userData["email"] = email;
	// userData["userTickets"] = strconv.FormatUint(uint64(userTickets),10);
	var userData =  UserData{
			firstName : firstName,
	lastName: lastName,
	email: email,
	numberOfTickets: userTickets,
	}


	bookings = append(bookings, userData)
	remaingTickets = remaingTickets - userTickets
	fmt.Printf("List of bookings : %v\n",bookings)
	fmt.Printf("%v %v with %v booked %v tickets \n",firstName,lastName,email ,userTickets)
	fmt.Println("Thank you! You will recive the confirmation email.")
	fmt.Printf("Total Tickets %v and Avaliable tickets %v \n", conferenceTickets, remaingTickets)

	return bookings, remaingTickets
}

func sendTicket(numberOfTickets uint, firstName string, lastName string, email string){
	// Needs time
	time.Sleep(10* time.Second)
	var ticket =fmt.Sprintf("%v tickets for %v %v", numberOfTickets, firstName, lastName)
	fmt.Println("-------------------------------------------")
	fmt.Printf("Sending Ticket: \n%v to %v\n",ticket, email)
	fmt.Println("-------------------------------------------")
}