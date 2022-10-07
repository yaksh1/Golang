package main

import (
	"fmt"
	"strings"
)

// -----------------------------------------------
// function to validate the name of user
// -----------------------------------------------
func validateName() (string, string) {
	for len(firstName) < 2 || len(lastName) < 2 {
		fmt.Println("Your first name or last name is too short , please try again.")

		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)
	}
	return firstName, lastName
}

// -----------------------------------------------
// function to validate the email ID of user
// -----------------------------------------------
func validateEmail() string {
	for !strings.Contains(emailId, "@") {
		fmt.Println("Your email does not contain @ sign , please enter valid email.")
		fmt.Print("Enter your valid email ID: ")
		fmt.Scan(&emailId)
	}
	return emailId
}

// -----------------------------------------------
// function to validate the tickets of user
// -----------------------------------------------
func validateTickets() int {
	for tickets > remainingTickets || tickets < 0 {
		if tickets <= 0 {
			fmt.Println("\nPlease enter a valid number of tickets")
		} else {
			fmt.Printf("\nSorry we have only %v tickets remaining, so you cannot book %v tickets\n", remainingTickets, tickets)
		}
		fmt.Print("Enter number of tickets again: ")
		fmt.Scan(&tickets)
	}
	return tickets
}
