package pwcheck

import (
	"errors"
	"fmt"
	"strings"
)

// Validate returns nil if the password has greater than or
// equal to the minimum entropy. If not, an error is returned
// that explains how the password can be strengthened. This error
// is safe to show the client
func Validate(password string, minEntropy float64) error {
	entropy := getEntropy(password)
	if entropy >= minEntropy {
		return nil
	}

	hasReplace := false
	hasSep := false
	hasOtherSpecial := false
	hasLower := false
	hasUpper := false
	hasDigits := false

	for _, c := range password {
		switch {
		case strings.ContainsRune(ReplaceChars, c):
			hasReplace = true
		case strings.ContainsRune(SepChars, c):
			hasSep = true
		case strings.ContainsRune(OtherSpecialChars, c):
			hasOtherSpecial = true
		case strings.ContainsRune(LowerChars, c):
			hasLower = true
		case strings.ContainsRune(UpperChars, c):
			hasUpper = true
		case strings.ContainsRune(DigitsChars, c):
			hasDigits = true
		}
	}

	allMessages := []string{}

	if !hasOtherSpecial || !hasSep || !hasReplace {
		allMessages = append(allMessages, "including more special characters")
	}
	if !hasLower {
		allMessages = append(allMessages, "using lowercase letters")
	}
	if !hasUpper {
		allMessages = append(allMessages, "using uppercase letters")
	}
	if !hasDigits {
		allMessages = append(allMessages, "using numbers")
	}

	if len(allMessages) > 0 {
		return fmt.Errorf(
			"insecure password, try %v or using a longer password",
			strings.Join(allMessages, ", "),
		)
	}

	return errors.New("insecure password, try using a longer password")
}

// ValidateSlice takes a slice of passwords and returns a slice of errors
// that correspond to the passwords. If the password is valid, the error
// will be nil. If the password is invalid, the error will be the same
// as the Validate function
func ValidateSlice(passwords []string, minEntropy float64) []error {
	errors := make([]error, len(passwords))
	for i, password := range passwords {
		errors[i] = Validate(password, minEntropy)
	}
	return errors
}
