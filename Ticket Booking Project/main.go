package main

import (
	"fmt"
)

const conferenceTicket = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50 // uint is an unsigned integer
var bookings []string          // Array of strings of undefined length

func main() {
	// function to greet the user
	greetUser()

	for {
		// function to get user input
		var firstName, lastName, email, userTicket = getUserInput()

		// function to validate user input
		var isValidName, isValidEmail, isValidTicket = validateUser(firstName, lastName, email, userTicket)

		// If the user data is valid, book the tickets
		if isValidName && isValidEmail && isValidTicket {
			// function to book tickets
			bookTickets(firstName, lastName, userTicket)

			// function to print the first name of the people who have booked tickets
			var firstNames = printFirstName()
			fmt.Println("\nThe following people have booked tickets for the", conferenceName, ":", firstNames)

			// If there are no tickets remaining, stop the loop
			if remainingTickets == 0 {
				fmt.Println("\nSorry, there are no tickets remaining for the", conferenceName)
				break
			}

		} else {
			if !isValidName {
				fmt.Println("\nSorry, your name is invalid. Please try again.")
			}
			if !isValidEmail {
				fmt.Println("\nSorry, your email is invalid. Please try again.")
			}
			if !isValidTicket {
				fmt.Println("\nSorry, the number of tickets you want to book is invalid. Please try again.")
			}
			continue
		}
	}
}
