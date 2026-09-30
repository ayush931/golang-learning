package auth

// this becomes private
func extractSession() string {
	return  "LoggedIn"
}

func GetSession() string {
	return extractSession()
}