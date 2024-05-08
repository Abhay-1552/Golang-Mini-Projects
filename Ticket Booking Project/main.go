package main

import (
	"Ticket_Booking_Project/validate" // Import the validate package from the Ticket_Booking_Project folder
	"fmt"
	"sync"
)

const conferenceTicket = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50     // uint is an unsigned integer
var bookings = make([]userData, 0) // Array of maps of undefined length

var wg = sync.WaitGroup{} // WaitGroup is a struct type that waits for a collection of goroutines to finish

type userData struct {
	firstName  string
	lastName   string
	email      string
	userTicket uint
}

func main() {
	// function to greet the user
	greetUser()

	for {
		// function to get user input
		var firstName, lastName, email, userTicket = getUserInput()

		// function to validate user input
		// Here, we are using the ValidateUser function from the validate package and keeping the function name uppercase to make it public
		var isValidName, isValidEmail, isValidTicket = validate.ValidateUser(firstName, lastName, email, userTicket, remainingTickets)

		// If the user data is valid, book the tickets
		if isValidName && isValidEmail && isValidTicket {
			// function to book tickets
			bookTickets(firstName, lastName, userTicket, email)

			wg.Add(1) // Add 1 to the WaitGroup counter

			// function to send a confirmation email
			go sendTickets(firstName, lastName, userTicket, email)

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
		wg.Wait() // Wait for all the goroutines to finish
	}
}
