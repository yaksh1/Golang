package main

import (
	"fmt"
	"sync"
	"time"
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

var (
	movieName        = "Avatar"
	movieTickets     int
	remainingTickets int
	bookings         = make([]UserData, 0)
	n                int
	firstName        string
	lastName         string
	emailId          string
	tickets          int
)

var wg = sync.WaitGroup{}

func main() {
	//taking input of total tickets
	fmt.Print("Total number of tickets for the movie: ")
	fmt.Scan(&n)

	//initializing movie tickets and remaining tickets
	movieTickets = n
	remainingTickets = n

	//calling to greet users
	greetUsers()
	greetUsers()

	fmt.Println("--------------------------------------------------------------")

	for remainingTickets > 0 && len(bookings) < n {

		//input of name from user
		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)

		//validating name of the user
		firstName, lastName = validateName()

		//input of email ID of user
		fmt.Print("Enter your Email ID: ")
		fmt.Scan(&emailId)

		//validating email ID
		emailId = validateEmail()

		//input of number of tickets user wants to buy
		fmt.Print("Enter how many tickets you want to book: ")
		fmt.Scan(&tickets)

		//validating tickets
		tickets = validateTickets()

		//booking the tickets and appending the slice
		bookings, remainingTickets = bookTicket()

		// generating and sending the ticket
		wg.Add(1) // number of go threads
		go sendTicket()

		//calling the function to print only first name of the user for privacy
		printFirstNames()
	}
	//end program
	fmt.Println("!!The tickets for the movie are sold out!!")

	wg.Wait()
}

// ----------------------------------------
// function to greet the users
// ----------------------------------------
func greetUsers() {
	fmt.Println("\nWelcome to", movieName, "booking application")
	fmt.Println("\nWe have total of", movieTickets, "tickets and", remainingTickets, "tickets are still available")
	fmt.Println("\nBook your tickets here to watch the movie")
}

// -----------------------------------------------
// function to print the first names of booking
// -----------------------------------------------
func printFirstNames() {
	firstNamesList := []string{}
	for _, name := range bookings {
		firstNamesList = append(firstNamesList, name.firstName)
	}
	fmt.Printf("\nThe first names of bookings are: %v\n\n", firstNamesList)
}

// ---------------------------------------------------------------
// function to update remaining tickets and append bookings array
// ---------------------------------------------------------------
func bookTicket() ([]UserData, int) {
	//updating the remaining tickets
	remainingTickets -= tickets

	//create a struct for user data
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           emailId,
		numberOfTickets: tickets,
	}

	//updating the bookings slice
	bookings = append(bookings, userData)
	fmt.Printf("\nList of bookings is %v", bookings)

	fmt.Printf("\nThank you %v %v for booking %v tickets, you will receive a confirmation email at %v\n", firstName, lastName, tickets, emailId)
	fmt.Printf("\n%v tickets remaining for %v.\n", remainingTickets, movieName)

	return bookings, remainingTickets
}

func sendTicket() {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", tickets, firstName, lastName)

	//sending ticket
	fmt.Println("----------------------------------------------------")
	fmt.Printf("Sending ticket: \n%v\n to email address %v\n", ticket, emailId)
	fmt.Println("----------------------------------------------------")
	wg.Done()
}
