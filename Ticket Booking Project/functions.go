package main

import (
	"fmt"
	"strings"
)

// Greet the user
func greetUser() {
	fmt.Println("Welcome to the", conferenceName, "booking system.")
	fmt.Println("We have", remainingTickets, "tickets remaining for the", conferenceName, "out of", conferenceTicket, "tickets.")

	// Use to repeat a string multiple times
	fmt.Println(strings.Repeat("-", 75))
}

// Get the user's input
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Print("\nEnter your First Name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your Last Name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your Email: ")
	fmt.Scan(&email)

	fmt.Print("Enter the number of tickets you want to book: ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

// Validate the user's input
func validateUser(firstName string, lastName string, email string, userTicket uint) (bool, bool, bool) {
	var isValidName = len(firstName) > 2 && len(lastName) > 2

	// Use to check if a string contains a substring
	var isValidEmail = strings.Contains(email, "@") && strings.Contains(email, ".")

	var isValidTicket = userTicket > 0 && userTicket <= remainingTickets

	return isValidName, isValidEmail, isValidTicket
}

// function for booking tickets
func bookTickets(firstName string, lastName string, userTicket uint) ([]string, uint) {
	fmt.Println("\nThank you", firstName, lastName, "for booking", userTicket, "tickets for the", conferenceName)

	// reduce the number of tickets
	remainingTickets = remainingTickets - userTicket

	fmt.Println("We have", remainingTickets, "tickets remaining for the", conferenceName, "out of", conferenceTicket, "tickets.")

	// Add the user to the bookings array
	var fullName = firstName + " " + lastName
	bookings = append(bookings, fullName)

	return bookings, remainingTickets
}

// Print the first name of the people who have booked tickets
func printFirstName() []string {
	firstNames := []string{}

	// _ is a blank identifier used to ignore the index
	for _, name := range bookings {

		// Use to split a string into substrings
		var first = strings.Fields(name)

		firstNames = append(firstNames, first[0])
	}
	return firstNames
}
