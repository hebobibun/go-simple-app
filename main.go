package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have %v remaining ticket of %v total ticket\n", remainingTickets, conferenceTickets)

	var bookings []string

	for remainingTickets > 0 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Printf("Enter your firstName : ")
		fmt.Scan(&firstName)

		for len(firstName) < 2 {
			fmt.Println("Your first name input is too short")
			fmt.Printf("Enter your firstName : ")
			fmt.Scan(&firstName)
		}

		fmt.Printf("Enter your lastName : ")
		fmt.Scan(&lastName)

		for len(lastName) < 2 {
			fmt.Println("Your last name input is too short")
			fmt.Printf("Enter your lastName : ")
			fmt.Scan(&lastName)
		}

		fmt.Printf("Enter your email address : ")
		fmt.Scan(&email)

		for !strings.Contains(email, "@") {
			fmt.Println("Your email address format is wrong (please include : @")
			fmt.Printf("Enter your email address : ")
			fmt.Scan(&email)
		}

		fmt.Printf("Enter amount of ticket : ")
		fmt.Scan(&userTickets)

		for userTickets > remainingTickets {
			fmt.Printf("We have only %v remaining, so you can't book %v tikects\n", remainingTickets, userTickets)
			fmt.Printf("Enter amount of ticket : ")
			fmt.Scan(&userTickets)
		}

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName + " " + lastName)
			
		fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
				
		firstNames := []string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
				
		fmt.Printf("These are all our bookings : %v\n", firstNames)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		if remainingTickets == 0 {
			fmt.Println("Tickets are fully booked. See you next year.")
			break
		}
	}
}