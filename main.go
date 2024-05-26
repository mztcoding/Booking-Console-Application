package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	var totalTickets int = 50
	var remainingTickets int = 40
	var bookings []string

	fmt.Printf("Hello! Welcome to our %v ticket booking aplication.\n", conferenceName)
	fmt.Printf("Hurry up! We have only %v tickets left out of %v tickets.\n", remainingTickets, totalTickets)
	fmt.Println("Please book your tickets to attend the conference.")
	for {
		var fname string
		var lname string
		var email string
		var bookedtickets int
		var furtherProcess string

		fmt.Printf("Please enter your firstname:\n")
		fmt.Scan(&fname)

		fmt.Printf("Please enter your lastname:\n")
		fmt.Scan(&lname)

		fmt.Printf("Enter your email address:\n")
		fmt.Scan(&email)

		fmt.Printf("How many tickets do you want to book?\n")
		fmt.Scan(&bookedtickets)

		userName := fname + " " + lname

		bookings = append(bookings, userName)
		fmt.Printf("This is full array %v\n", bookings)
		fmt.Printf("This is first entry %v\n", bookings[0])
		fmt.Printf("The type is %T\n", bookings)
		fmt.Printf("This is length of array %v\n", len(bookings))

		remainingTickets = remainingTickets - bookedtickets
		fmt.Printf("Thankyou %v! for booking %v tickets. ", userName, bookedtickets)
		fmt.Printf("We have sent a confirmation code at %v\n", email)
		fmt.Printf("Now we have %v tickets left.\n", remainingTickets)

		firstNames := []string{}
		for _, booking := range bookings {
			var fnames = strings.Fields(booking)
			firstNames = append(firstNames, fnames[0])
		}
		if remainingTickets == 0 {
			fmt.Println("All tickets are sold out.")
			break
		} else if bookedtickets > remainingTickets {
			fmt.Printf("We are sorry! Only %v tickets are available.\n", remainingTickets)
			continue
		}
		fmt.Println("These are all the bookings in the application", firstNames)

		fmt.Printf("Do you want to proceed further? (type yes/no) \n")
		fmt.Scan(&furtherProcess)
		if furtherProcess == strings.ToLower("No") {
			fmt.Println("Thanks again!")
			break
		} else if furtherProcess == strings.ToLower("yes") {
			continue
		} else {
			break
		}
	}
}
