package validation

import "regexp"

func IsValidEmail(email string) bool {
	// Define a regular expression pattern for a valid email address
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression
	re := regexp.MustCompile(emailPattern)

	// Use the MatchString function to check if the email matches the pattern
	return re.MatchString(email)
}

func IsValidPassword(password string) bool {
    return len(password) > 4
}