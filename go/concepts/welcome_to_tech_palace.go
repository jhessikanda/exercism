package concepts

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    var message string = ""
    
	for i := 0; i < numStarsPerLine; i++ {
    	message = message + "*"
	}

    message = message + "\n" + welcomeMsg + "\n"

    for i := 0; i < numStarsPerLine; i++ {
    	message = message + "*"
	}
	
    fmt.Println("Debug message :: " + message)
    
    return message
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	var msg = strings.Split(oldMsg, "\n")
    return strings.Trim(strings.Trim(msg[1], "*"), " ")
}
