package utils

import (
	"fmt"
	"math/rand"
	"strings"
)

// CreateRandomEmail creates random email based on two strings
func CreateRandomEmail(firstName, lastName string) string {
	email := fmt.Sprintf("%s.%s@%s", firstName, lastName, domain)

	return strings.ToLower(email)
}

// CreateRandomPhoneNumber creates a random phone number as a string
func CreateRandomPhoneNumber() string {
	randomInteger := rand.Intn(9999999998-1000000001) + 1000000001

	return fmt.Sprintf("%d", randomInteger)
}

// RandomIndex returns a random number between max and min
func RandomIndex(min, max int) int {
	return rand.Intn(max-min) + min
}

// ManipulateName manipulates name
func ManipulateName(name string) string {
	nameIndex := RandomIndex(0, len(name))

	runes := []rune(name)
	newName := fmt.Sprintf("%s%s%s",
		string(runes[0:nameIndex+1]),
		string(runes[nameIndex:nameIndex+1]),
		string(runes[nameIndex+1:len(name)]))

	return newName
}

// CreateAddress creates an address with a concatenated alphanumerica value
func CreateAddress(address string) string {
	numeric := []string{"0", "1", "2"}
	alpha := []string{"A", "B", "C"}

	index := RandomIndex(0, 3)

	return fmt.Sprintf("%s %s%s%s", address, alpha[index], numeric[index], alpha[index])
}

// GetFootSite | returns foot site based on input
func GetFootSite(num int) string {
	return FootSites[num%len(FootSites)]
}
