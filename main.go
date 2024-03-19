package main

import (
	"fmt"
	"sync"
	"time"

	"www.mylearning.com/helper"
)

// Writing a command line application to build a booking platform for conference

// package level variables can be accessed in any function
// const keyword is used to declarea variavle, cannot be changed
const conferenceTickets int = 50

// explicit decleration  (unit contains only positive whiole numbers)
var remainingTickets uint = 50

// Short Hand declaration (package level variables cannot be declared in short hand )
// conferenceName := "Go Conference"
var conferenceName string = "Go Conference"

// slice initialization
// slice is a datatype which can store any number of values. in this bookings it can store any number of string values
// var bookingNames = []string{}
// the bookingNames slice takes struct data type values
var bookingNames = make([]UserData, 0)

// collection of different types of data
// within the block we define keys and type of values
// type keyword denotes custom type
// type keyword crreats a new type and struct name
// structure of user type StructName
type UserData struct {
	firstName   string
	lastName    string
	email       string
	nooftickets uint
}

//synchronization using goroutines

var wg = sync.WaitGroup{}

func main() {

	//array decleration and using it further
	// var arr1 = [10]int{2, 3, 4, 5, 6, 7, 8, 9, 74, 10}
	// fmt.Println(arr1[1])
	// fmt.Println(arr1[2])

	//Using Scanner to get data from user for bookings
	// using the function to greet user with writing the code directly
	//as we declared package level variables we dont need to pass parameters to greetUser
	greetUser()
	//fmt.Printf("Welcome to %v booking application.!! \nWe have total of %v tickets and %v are remaining. ", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("\n")
	// infinite for loop can be directly used with for{ }
	for {
		// //initializng variables
		// var firstName string
		// var lastName string
		// var email string
		// //uint bc tickets can always be positive
		// var userTickets uint

		// //using Scan as taking user inputs
		// fmt.Println("\n Enter your First Name:")
		// //here we use &(pointer) bc the address at which the data is stored has to be changed
		// fmt.Scan(&firstName)

		// //using Scan as taking user inputs
		// fmt.Println("Enter your Last Name:")
		// //here we use &(pointer) bc the address at which the data is stored has to be changed
		// fmt.Scan(&lastName)

		// //using Scan as taking user inputs
		// fmt.Println("Enter your Email:")
		// //here we use &(pointer) bc the address at which the data is stored has to be changed
		// fmt.Scan(&email)

		// //using Scan as taking user inputs
		// fmt.Println("Enter number of tickets:")
		// //here we use &(pointer) bc the address at which the data is stored has to be changed
		// fmt.Scan(&userTickets)
		//function dosent take any inputs but returns 4 variables
		firstName, lastName, email, userTickets := userInput()

		//validating user inputs
		//byDefault bool is False
		//len function returns the length of the array/slice or no of characters in a string
		//checking if length of first and lastname is greater than 2 as a criteria and if yes true or false (boolean Type)
		//declared using bool type and below declared in short hand  both are same

		// var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		// //check for isvalidname type
		// fmt.Printf("the type of isValidName is %T", isValidName)
		// //checking if email (string) contains @ character in email is valid or not and returns bool true or false
		// isValidEmail := strings.Contains(email, "@")
		// //checking if entered tickets is less then remaining tickets and not 0
		// //using && operator to check two conditions and  return true
		// isValidTickets := userTickets > 0 && userTickets <= remainingTickets

		//to check between two cities hyderabad or secundrabad we use or operator and return true or false
		// isValidCity := city == "hyderabad" || city == "secundrabad"

		//function to check valid user input and collect multiple return values

		//while running the applicatiion go run main.go helper.go ) include helper.go file with main.go file
		isValidName, isValidEmail, isValidTickets := helper.CheckUserInput(firstName, lastName, email, userTickets, remainingTickets)

		//array is declared as ( var bookings = [50]int {}) or (var bookings [50] string) decleratio of an array
		// logical not (!) operator is used to reverse the boolean value (if condition is true it makes false)
		//using if else conditional statements checking if all 3 validations are true
		// && AND operator says all conditions must be true

		if isValidEmail && isValidName && isValidTickets {
			// {//book tickets according to the system
			// remainingTickets = remainingTickets - userTickets
			// //append is used to add first name and last name to the slice
			// bookingNames = append(bookingNames, firstName+" "+lastName)

			// fmt.Printf("Thank you %v for the booking %v, you will receive a confirmation email to %v \n", firstName, userTickets, email)
			// fmt.Printf("you still have %v for %v !!", remainingTickets, conferenceName)}

			//whole block of booking ticket is written in a function bookticket

			bookTicket(userTickets, firstName, lastName, email)
			// wg.Add  addes number of go routines used in the function, executed before creating a new thread
			wg.Add(1)
			//with go keyword we are executing the send ticket in a different thread concurrently and the normal; program execution is profromend as usual
			go sendTicket(userTickets, firstName, lastName, email)

			//{Print first names //declaring slice for storing firstnames
			//firstNames := []string{}
			//for each loop to store only first names and display them
			//for each loop goes through the list of bookingnames and gives each booking saperately
			// for _, booking := range bookingNames {
			// 	//strings.fields split the booking full name using space and stores them in a array/ slice.
			// 	var names = strings.Fields(booking)
			// 	//considering first name, just take value present at index 0
			// 	firstNames = append(firstNames, names[0])
			// }
			// fmt.Printf("the first names \n %v \n of people attending %v", firstNames, conferenceName)}

			//whole block of code is declared in a function and function is called inthe main function
			//accessing the package level variables therefore no need to mention them
			printFirstNames()

			//exiting thge application if remaining teckets in 0
			//using if else conditional statements we are confirming the condition
			if remainingTickets == 0 {
				fmt.Printf("The %v is House Full", conferenceName)
				//break statement used here ends the application and brings us out of for loop
				//break statement breaks helps to break the loop and continue executing existing program after the loop
				break
			} //continue causes loop to skip the reminder of its body  . restarts the condition

		} else {
			//only if is used because we want all the three statements to be executed
			//checking if email is valid or not and (! indicates false)
			if !isValidEmail {
				fmt.Printf("\nProvided Email %v is not valid please provide correct email ", email)
			}
			if !isValidName {
				fmt.Println("\nProvided FirstName and LastName are not valid")
			}
			if !isValidTickets {
				fmt.Println("request for right amount of tickets")
			}
		}

	}
	//end of main function
	//waits for all the threads to execute before the application is ended
	wg.Wait()

}

// assessing values from direct package level you dont need to give them saperately
func greetUser() {
	fmt.Printf("Welcome to %v booking application.!! \nWe have total of %v tickets and %v tickets are remaining. ", conferenceName, conferenceTickets, remainingTickets)
}

// here in function parameters as the bookingNames is slice it is declared with []
func printFirstNames() {
	//Print first names //declaring slice for storing firstnames
	firstNames := []string{}
	//for each loop to store only first names and display them
	//for each loop goes through the list of bookingnames and gives each booking saperately
	for _, booking := range bookingNames {
		//strings.fields split the booking full name using space and stores them in a array/ slice.
		//var names = strings.Fields(booking)
		//considering first name, just take value present at index 0
		//by implementing the struct we can directly access the firstname of the userData
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Printf("the first names \n %v \n of people attending %v", firstNames, conferenceName)
}

// // this functio takes all the parameters as inputs and returs boolean values
// // In Go function can return multiple values from the function
// // the return type of all return are definde in the function decleration using()
// func checkUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
// 	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
// 	//check for isvalidname type
// 	fmt.Printf("the type of isValidName is %T", isValidName)
// 	//checking if email (string) contains @ character in email is valid or not and returns bool true or false
// 	isValidEmail := strings.Contains(email, "@")
// 	//checking if entered tickets is less then remaining tickets and not 0
// 	//using && operator to check two conditions and  return true
// 	isValidTickets := userTickets > 0 && userTickets <= remainingTickets

// 	// can return any number of variables using return keyword
// 	return isValidName, isValidEmail, isValidTickets
// }

func userInput() (string, string, string, uint) {
	//initializng variables
	var firstName string
	var lastName string
	var email string
	//uint bc tickets can always be positive
	var userTickets uint

	//using Scan as taking user inputs
	fmt.Println("\n Enter your First Name:")
	//here we use &(pointer) bc the address at which the data is stored has to be changed
	fmt.Scan(&firstName)

	//using Scan as taking user inputs
	fmt.Println("Enter your Last Name:")
	//here we use &(pointer) bc the address at which the data is stored has to be changed
	fmt.Scan(&lastName)

	//using Scan as taking user inputs
	fmt.Println("Enter your Email:")
	//here we use &(pointer) bc the address at which the data is stored has to be changed
	fmt.Scan(&email)

	//using Scan as taking user inputs
	fmt.Println("Enter number of tickets:")
	//here we use" &"(pointer) bc the address at which the data is stored has to be changed
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	//book tickets according to the system
	remainingTickets = remainingTickets - userTickets
	// //append is used to add first name and last name to the slice
	// bookingNames = append(bookingNames, firstName+" "+lastName)
	//declaring a struct type to the bookings
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		nooftickets: userTickets,
	}
	bookingNames = append(bookingNames, userData)
	fmt.Printf("List of bookings is %v", bookingNames)

	fmt.Printf("Thank you %v for booking %v tickets, you will receive a confirmation email to %v \n", firstName, userTickets, email)
	fmt.Printf("you still have %v tickets left for %v !!", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	//time delay
	time.Sleep(20 * time.Second)
	//Sprintf function helps you put together a string just like printf but insted of printing it out we can save into a string variable
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)

	fmt.Printf("sending ticket: \n %v to email address %v ", ticket, email)
	//this function removes thread when the execution is done
	//decrementing counter
	wg.Done()
}

//for has 2 types
//infinite for loop where the loop repets endlessly as a condition is always true
//eg for true  {}  if no condition specifies its allways true orr for condition is true {execute this block of code}
// for each loop, syntac for index,list := range lists { block of code }

//if - else are very important as they decide what flow has to be followed in order to execute the program
