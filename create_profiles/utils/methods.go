package utils

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/Pallinder/go-randomdata"
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

	return strings.Title(strings.ToLower(newName))
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

// CreateProfile | returns a newly created profile based on an index and virtual card information
func CreateProfile(virtualCreditCard VirutalCCInfo, index int) Profile {
	var profile Profile

	//////////////////////////////////////////////
	// KEEP THESE FOR FUTURE REFERENCE
	// rand.Seed(time.Now().UTC().UnixNano())
	// index := RandomIndex(0, 3)
	/////////////////////////////////////////////

	// Setup Credit Card Information
	profile.CCNumber = virtualCreditCard.CCNumber
	profile.CVV = virtualCreditCard.CVV
	profile.ExpMonth = virtualCreditCard.ExpMonth
	profile.ExpYear = virtualCreditCard.ExpYear
	profile.CardType = virtualCreditCard.CardType

	// Setup Phone Number
	profile.Phone = CreateRandomPhoneNumber()
	if len(profile.Phone) != 10 {
		fmt.Println(fmt.Sprintf("wrong number at %s", profile.Name))
	}

	// Setup Address
	profile.Same = true

	address := CreateAddress(Address)

	profile.Shipping.Address = address
	profile.Shipping.Apt = Apt
	profile.Shipping.City = City
	profile.Shipping.State = State
	profile.Shipping.Zip = Zip

	profile.Billing.Address = address
	profile.Billing.Apt = Apt
	profile.Billing.City = City
	profile.Billing.State = State
	profile.Billing.Zip = Zip

	profile.Country = Country

	// Setup First Name & Last Name
	firstName := ManipulateName(FirstName)
	lastName := ManipulateName(LastName)

	// TODO: create rules function based on what site the profile is being used
	site := virtualCreditCard.Site
	if strings.Compare(site, BestBuy) == 0 {
		randomdata.FirstName(randomdata.Male)
		randomdata.LastName()
	}

	profile.Shipping.FirstName = firstName
	profile.Shipping.LastName = lastName
	profile.Billing.FirstName = firstName
	profile.Billing.LastName = lastName

	// Setup Email
	profile.Email = CreateRandomEmail(firstName, lastName)

	// Setup Profile Name
	profile.Name = fmt.Sprintf("Profile_%d_%s", index, site)

	return profile
}

// IsFootSite | checks if the site a foot site
func IsFootSite(site string) bool {
	for _, footSite := range FootSites {
		if strings.Compare(site, footSite) == 0 {
			return true
		}
	}

	return false
}
