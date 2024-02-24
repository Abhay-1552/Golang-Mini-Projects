package validate

import "strings"

// Keep the function name uppercase to make it public or to export it explicitly

// Validate the user's input
func ValidateUser(firstName string, lastName string, email string, userTicket uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName = len(firstName) > 2 && len(lastName) > 2

	// Use to check if a string contains a substring
	var isValidEmail = strings.Contains(email, "@") && strings.Contains(email, ".")

	var isValidTicket = userTicket > 0 && userTicket <= remainingTickets

	return isValidName, isValidEmail, isValidTicket
}
