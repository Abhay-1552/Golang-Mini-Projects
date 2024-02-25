package main

import (
	"fmt"
	"strings"
	"time"
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

// function for booking tickets
func bookTickets(firstName string, lastName string, userTicket uint, email string) {
	fmt.Println("\nThank you", firstName, lastName, "for booking", userTicket, "tickets for the", conferenceName)

	// reduce the number of tickets
	remainingTickets = remainingTickets - userTicket

	fmt.Println("We have", remainingTickets, "tickets remaining for the", conferenceName, "out of", conferenceTicket, "tickets.")

	// Create a struct to store the user data
	// Struct is a user-defined data type that contains a collection of named fields/properties
	var userData = userData{
		firstName:  firstName,
		lastName:   lastName,
		email:      email,
		userTicket: userTicket,
	}

	bookings = append(bookings, userData)
}

// Print the first name of the people who have booked tickets
func printFirstName() []string {
	firstNames := []string{}

	// _ is a blank identifier used to ignore the index
	for _, name := range bookings {
		firstNames = append(firstNames, name.firstName) // Take the first name from the struct and append it to the firstNames array
	}

	return firstNames
}

// function to send a confirmation email
func sendTickets(firstName string, lastName string, userTicket uint, email string) {
	// time.Sleep is used to simulate a delay
	time.Sleep(5 * time.Second)

	fmt.Println("\n----------------------------------------")
	fmt.Println("Sending tickets to the attendees...")
	fmt.Println("SENDING TICKETS:", userTicket, "number of tickets to", firstName, lastName, "at", email)
	fmt.Println("----------------------------------------")

	wg.Done() // Subtract 1 from the WaitGroup counter
}
