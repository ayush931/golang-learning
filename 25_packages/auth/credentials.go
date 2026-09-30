package auth

import "fmt"

// Captial letter function let you import the function anywhere, in small can only use or export in same package
func LoginWithCredentials(username string, password string) {
	fmt.Println("Login user using", username, password)
}