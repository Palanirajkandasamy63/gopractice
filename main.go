package main

import (
	"booking-app/helper"
	"fmt"
	"time"
	// "strconv"
)

const conferenceTickets int = 50

var remainingTickets uint = 50
var conferenceName = "Go Conference"

// var	bookings = []string{} list of string
// var userData= make(map[string]string) map
// var bookings = make([]map[string]string, 0) //list of maps which list contains maps 
var bookings =make([]UserData,0)

type UserData struct{
	firstName string 
	lastName string
	email string 
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		// validate user input

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			// book ticket in system
			bookTicket( userTickets, firstName, lastName, email)
			 go sendTicket( userTickets, firstName, lastName, email)


			// print only first names
			firstNames := getFirstNames()
			fmt.Printf("The first names are :%v\n", firstNames)

			// exit application if no tickets are left
			if remainingTickets == 0 {	
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("firt name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
}

//before map we used strings field check in readme.md
// func getFirstNames(bookings []string)[]string{
// 	firstNames := []string{}
// 			for _, booking := range bookings {
// 					var names = strings.Fields(booking)
// 					firstNames = append(firstNames, names[0])
// 				}
// 	        return firstNames
// 	}

func getFirstNames()[]string{
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}
// func getFirstNames(bookings []map[string]string) []string {
// 	firstNames := []string{}
// 	for _, booking := range bookings {
// 		// firstNames = append(firstNames, booking["firstName"])
// 		firstNames=append(firstNames, booking.firstName)
// 	}
// 	return firstNames
// }

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// asking for user input
	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	//create a map for a user

	var UserData =UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}


	// before struct

	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// //convert typescat check in documnetation or google
	// //strconv.FormatUint(uint64(userTickets),10)
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// bookings = append(bookings, firstName+" "+lastName) for appending into list

	bookings = append(bookings, UserData)
	fmt.Printf("list of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
func sendTicket(userTickets uint, firstName string, lastName string,email string ){
	time.Sleep(10 *time.Second)
	fmt.Printf("%v tickets for  %v %v ",userTickets,firstName,lastName)
	fmt.Println("######")
	var ticket=fmt.Sprintf("%v tickets for  %v %v ",userTickets,firstName,lastName)
	fmt.Printf("sending ticket:\n  %v to email address %v \n",ticket,email)
	fmt.Println("######")
}
