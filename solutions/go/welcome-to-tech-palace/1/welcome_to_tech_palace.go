package techpalace
import "strings"


// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    return "Welcome to the Tech Palace, " + strings.ToUpper(customer) 
	panic("Please implement the WelcomeMessage() function")
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    repeat := strings.Repeat("*", numStarsPerLine)
    return repeat + "\n" + welcomeMsg + "\n" + repeat
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	oldMsg = strings.ReplaceAll(oldMsg, "*", "")
    oldMsg = strings.TrimSpace(oldMsg)
    oldMsg = strings.ReplaceAll(oldMsg, "\n", "")
    return oldMsg
}
