package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName string = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		fmt.Printf("%v %v, your email is %v, you have requested %v tickets \n", firstName, lastName, email, userTickets)
		if isValidName && isValidEmail {
			if isValidTickets {
				fmt.Println("Thank you for your purchase!")
				remainingTickets = remainingTickets - userTickets
				// Adding the userdata to the bookings array
				var userData = UserData{
					firstName:       firstName,
					lastName:        lastName,
					email:           email,
					numberOfTickets: userTickets,
				}

				bookings = append(bookings, userData)
				fmt.Printf("list of bookings: %v \n", bookings)
			} else {
				fmt.Println("Sorry, we don't have enough tickets!")
				continue
			}
			fmt.Printf("We have %v remaining tickets for %v\n", remainingTickets, conferenceName)

			wg.Add(1)
			go sendTickets(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("All first names of bookings are: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				//end program
				fmt.Println("All tickets are sold out, Please try again next year!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First or last name is too short!")
			}
			if !isValidEmail {
				fmt.Println("Invalid email!")
			}
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking system \n", conferenceName)
	fmt.Printf("We have a total of %v tickets, out of which %v remain! \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket now!")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email:")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets you want:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##############################")
	fmt.Printf("Sending ticket: \n%v\n to %v\n", ticket, email)
	fmt.Println("##############################")
	wg.Done()
}
