package main

import (
	"booking-app/helper"
	"fmt"
)

type UserData struct {
	firstName    string
	lastName     string
	emailAddress string
	ticketCount  int
}

func main() {
	var totalTickets int = 50
	var availableTickets int = 50
	var conferenceName string = "Go Learning by Shankar"
	var firstName string
	var lastName string
	var emailAddress string
	var tickets int
	//var bookingList = make([]map[string]string, 0)
	var bookingList []UserData
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
			isValidName := helper.IsValidUserName(firstName, lastName)
			var isValidTicketCount bool = tickets > 0 && availableTickets >= tickets

			if isValidName && isValidTicketCount {
				availableTickets = availableTickets - tickets
				/*var userData = make(map[string](string))
				userData["firstName"] = firstName
				userData["lastName"] = lastName
				userData["ticketCount"] = strconv.Itoa(tickets)*/
				var userData UserData
				userData.firstName = firstName
				userData.lastName = lastName
				userData.emailAddress = emailAddress
				userData.ticketCount = tickets
				bookingList = append(bookingList, userData)
				fmt.Printf("Hi %v %v, Thanks for booking %v tickets with us. We still have %v tickets remaining \n", firstName, lastName, tickets, availableTickets)
				fmt.Printf("The booking map list so far %v", bookingList)
				firstNames := bookTickets(bookingList)
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

func bookTickets(bookingList []UserData) []string {
	var firstNames []string
	for _, element := range bookingList {
		//var splittedNameList []string = strings.Fields(element)
		//firstNames = append(firstNames, element["firstName"])
		firstNames = append(firstNames, element.firstName)
	}

	return firstNames
}
