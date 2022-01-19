package helper

import "strings"

func ValidateUser(firstName string, lastName string, email string, userTickets uint, remaingTickets uint) (bool, bool, bool) {
	return ValidateUserInput(firstName, lastName), ValidateUserEmail(email), ValidateUserTicketCount(userTickets, remaingTickets)
}

func ValidateUserInput(firstName string, lastName string) bool {
	return len(firstName) >= 2 && len(lastName) >= 2
}

func ValidateUserEmail(email string) bool {
	return strings.Contains(email, "@")
}

func ValidateUserTicketCount(userTickets uint, remaingTickets uint) bool {
	return userTickets > 0 && userTickets <= remaingTickets
}
