package main

import (
	"fmt"
	"strings"
	"time"
)

func sleep() {
	time.Sleep(1 * time.Second)
}

// func valNames(name) {
// 	return len(name) >= 2
// }

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	sleep()
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	sleep()
	fmt.Printf("get your tickets here to attend\n")
	sleep()

	bookings := []string{}
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	for {
		fmt.Println("\nWhat is your firstname!")
		fmt.Scan(&firstName)
		fmt.Println("...")
		sleep()

		fmt.Println("\nWhat is your lastname!")
		fmt.Scan(&lastName)
		fmt.Println("...")
		sleep()

		fmt.Println("\nWhat is your email!")
		fmt.Scan(&email)
		fmt.Println("...")
		sleep()

		fmt.Println("\nHow many tickets do you want!")
		fmt.Scan(&userTickets)
		fmt.Println("processing...")
		sleep()

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("\nThank you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
			sleep()
			fmt.Printf("You will receive a confirmation email at %v.\n", email)
			sleep()
			fmt.Println("Now bugger off!")
			sleep()
			fmt.Printf("\nThere are %v tickets remaining for %v\n", remainingTickets, conferenceName)
			sleep()
			fmt.Printf("\nThese are all the bookings on fullname: %v\n", bookings)
			sleep()
			fmt.Printf("\nThese are all the bookings on firstname: %v\n", firstNames)
			sleep()

			noTicketsRemaining := remainingTickets <= 0
			if noTicketsRemaining {
				fmt.Println("\n\n********************************\nOur conference is SOLD OUT!\nWe hope to see you next year!\n********************************")
				sleep()
				break
			} else {
				fmt.Println("\n\n********************************\nYou will now return to the homescreen\n********************************")
				sleep()
				continue
			}

		} else {
			if !isValidName {
				fmt.Println("Firstname or lastname you entered is too short")
				sleep()
			}
			if !isValidEmail {
				fmt.Println("You did not provide a valid email address")
				sleep()
			}
			if !isValidTicketNumber {
				fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
				sleep()
			}
			if isValidName && isValidEmail && isValidTicketNumber {
				fmt.Println("Sorry, something went wrong")
				sleep()
			}

			fmt.Println("\n\n********************************\nYou will now return to the homescreen\n********************************")
			sleep()
			continue
		}
	}
}
