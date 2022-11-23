package main

import (
	"fmt"
	"strconv"
)

const conferenceTickets uint = 50
var conferenceName = "Go Conference"
var remainingTickets = uint(50)
var bookings = make([]map[string]string, 0)

func main() {

	greetUser()

	for remainingTickets > 0 {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidNumTicket := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidNumTicket {
			
			BookTicket(userTickets, firstName, lastName, email)

			// A function to print all the firstName of the bookings
			firstNames := getFirstNames()
			fmt.Printf("The firstName of bookings are : %v\n", firstNames)

			fmt.Println("=================")

			if remainingTickets == 0 {
				fmt.Println("Tickets are fully booked. See you next year.")
				fmt.Println("=================")
				break
			}
		} else {
			fmt.Println("=================")
			fmt.Println("Invalid input : ")
			if !isValidName {
				fmt.Println("- Your firstName or lastName input are too short. Please input more than 2 characters.")
			} 
			if !isValidEmail {
				fmt.Println("- Email address you entered doesn't contain @ sign.")
			} 
			if !isValidNumTicket {
				fmt.Printf("- Please input less than %v remaining ticket.\n", remainingTickets)
			}
			fmt.Println("Try again.")
			fmt.Println("=================")
		}
	}
}

func greetUser() {
	
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have %v remaining ticket of %v total ticket\n", remainingTickets, conferenceTickets)
}

func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}		
	return firstNames
}

func getUserInput() (string, string, string, uint) {

	var firstName, lastName, email string
	var userTickets uint

	fmt.Printf("Enter your firstName : ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your lastName : ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email address : ")
	fmt.Scan(&email)

	fmt.Printf("Enter amount of ticket : ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func BookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
		
	fmt.Println("=================")
	fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("List of booking is %v\n", bookings)
}