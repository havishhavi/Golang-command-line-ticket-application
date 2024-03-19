// we want helper file to belong to  package
package helper

import (
	"fmt"
	"strings"
)

//this helper file can be common file or shared .go
//it has some functions which can be used to drive the application forward
//contains helper functions for the main application

//splitting code into multiple files

// this functio takes all the parameters as inputs and returs boolean values
// In Go function can return multiple values from the function
// the return type of all return are definde in the function decleration using()

// to export function checkUserInput we need to useupperCase first letter in the function
func CheckUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	//check for isvalidname type
	fmt.Printf("the type of isValidName is %T", isValidName)
	//checking if email (string) contains @ character in email is valid or not and returns bool true or false
	isValidEmail := strings.Contains(email, "@")
	//checking if entered tickets is less then remaining tickets and not 0
	//using && operator to check two conditions and  return true
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets

	// can return any number of variables using return keyword
	return isValidName, isValidEmail, isValidTickets
}
