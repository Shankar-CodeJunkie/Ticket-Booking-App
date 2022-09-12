package main

import (
	"fmt"
	"strings"
)

func main() {
	var totalTickets int = 50
	var availableTickets int = 50
	var conferenceName string = "Go Learning by Shankar"
	var firstName string
	var lastName string
	var emailAddress string
	var tickets int
	var bookingList []string
	fmt.Println("Welcome to", conferenceName, "!")
	fmt.Println("You have", availableTickets, "available tickets out of ", totalTickets)

	for {
		if availableTickets > 0 {
			fmt.Printf("Welcome to %v \n", conferenceName)
			fmt.Printf("You have %v available tickets out of %v total tickets \n", availableTickets, totalTickets)

			fmt.Printf("Enter your firstName: \n")
			fmt.Scan(&firstName)
			fmt.Printf("Enter your lastName: \n")
			fmt.Scan(&lastName)
			fmt.Printf("Enter your email address: \n")
			fmt.Scan(&emailAddress)
			fmt.Printf("Enter your tickets: \n")
			fmt.Scan(&tickets)

			//var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
			isValidName := len(firstName) >= 2 && len(lastName) >= 2
			var isValidTicketCount bool = tickets > 0 && availableTickets >= tickets

			if isValidName && isValidTicketCount {
				availableTickets = availableTickets - tickets
				bookingList = append(bookingList, firstName+" "+lastName)
				fmt.Printf("Hi %v %v, Thanks for booking %v tickets with us. We still have %v tickets remaining \n", firstName, lastName, tickets, availableTickets)
				var firstNames []string

				for _, element := range bookingList {
					var splittedNameList []string = strings.Fields(element)
					firstNames = append(firstNames, splittedNameList[0])
				}
				fmt.Printf("Booked user List %v \n", firstNames)

			} else {
				fmt.Printf("Error while booking your tickets\n")
				if !isValidName {
					fmt.Printf("First or Last Name you entered is too short")
				} else if !isValidTicketCount {
					fmt.Printf("Your number of tickets for the booking is invalid")
				}
				break
			}

		} else {
			fmt.Printf("Apologies, we have been sold out")
			break
		}
	}
}
