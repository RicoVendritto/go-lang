package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("get your tickets here to attend\n")

	bookings := []string{}
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	for {
		fmt.Println("\nWhat is your firstname!")
		fmt.Scan(&firstName)

		fmt.Println("\nWhat is your lastname!")
		fmt.Scan(&lastName)

		fmt.Println("\nWhat is your email!")
		fmt.Scan(&email)

		fmt.Println("\nHow many tickets do you want!")
		fmt.Scan(&userTickets)

		remainingTickets -= userTickets
		bookings = append(bookings, firstName+" "+lastName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("\nThank you %v %v for booking %v tickets.\nYou will receive a confirmation email at %v.\nNow bugger off!\n", firstName, lastName, userTickets, email)
		fmt.Printf("\nThere are %v tickets remaining for %v\n", remainingTickets, conferenceName)
		fmt.Printf("\nThese are all the bookings on fullname: %v\n", bookings)
		fmt.Printf("\nThese are all the bookings on firstname: %v\n", firstNames)
		fmt.Printf("\n\n********************************\nYou will now return to the homescreen\n********************************\n\n")
	}
}
