package utils

import (
	"fmt"
)

// CreateVeerProfile | returns a newly created profile based on an index and virtual card information for Veer
func CreateVeerProfile(virtualCreditCard VirtualCCInfo, index int) Profile {
	var profile Profile

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

	veerAddress := CreateAddress(VeerAddress)

	profile.Shipping.Address = veerAddress
	profile.Shipping.Apt = VeerApt
	profile.Shipping.City = VeerCity
	profile.Shipping.State = VeerState
	profile.Shipping.Zip = VeerZip

	profile.Billing.Address = veerAddress
	profile.Billing.Apt = VeerApt
	profile.Billing.City = VeerCity
	profile.Billing.State = VeerState
	profile.Billing.Zip = VeerZip

	profile.Country = Country

	isCapitalOne := false
	if virtualCreditCard.CardCompany == captailOne {
		isCapitalOne = true
	}

	// Setup First Name & Last Name
	firstName, lastName := ManipulateVeerName(isCapitalOne)
	// TODO: ^the problem here is if what if there are more profiles for a specific site being created than manipulations being created for that site

	// TODO: create rules function based on what site the profile is being used
	site := virtualCreditCard.Site

	profile.Shipping.FirstName = firstName
	profile.Shipping.LastName = lastName
	profile.Billing.FirstName = firstName
	profile.Billing.LastName = lastName

	// Setup Email
	profile.Email = CreateRandomEmail(firstName, lastName)

	// Setup Profile Name
	if virtualCreditCard.CardCompany == captailOne {
		profile.Name = fmt.Sprintf("Profile_%03d_Kush_%s", index, site)
	} else {
		profile.Name = fmt.Sprintf("Profile_%03d_Veer_%s", index, site)
	}

	return profile
}

func resetVeerName(isCapitalOne bool) (string, string) {
	firstName := VeerFirstName
	lastName := VeerLastName

	if isCapitalOne {
		firstName = KushFirstName
	}

	return firstName, lastName
}

// ManipulateVeerName | Manipulates veers and kush name and keeps track of used names
func ManipulateVeerName(isCapitalOne bool) (string, string) {
	manipulationChoice := RandomIndex(0, 3)

	firstName, lastName := resetVeerName(isCapitalOne)

	switch manipulationChoice {
	case 0:
		// manipulates both the first name and last name
		firstName = ManipulateString(firstName)
		lastName = ManipulateString(lastName)
	case 1:
		// manipulates just the first name and not the last name
		firstName = ManipulateString(firstName)
	case 2:
		// manipulates not the first name and just the last name
		lastName = ManipulateString(lastName)
	default:
		firstName, lastName = resetVeerName(isCapitalOne)
	}

	fullName := fmt.Sprintf("%s %s", firstName, lastName)

	nameExists := FullNameMap[fullName]

	if !nameExists {
		FullNameMap[fullName] = true
		return firstName, lastName
	}

	return ManipulateVeerName(isCapitalOne)
}

/*
 * Veer's Information
 */

// VeerLastName | Veer's last name
const VeerLastName = "patel"

// KushFirstName | Kush's first name
const KushFirstName = "kush"

// VeerFirstName | Veer's first name
const VeerFirstName = "veer"

// VeerAddress | Veer's address
const VeerAddress = "1501 N 31st Street"

// VeerApt | Veer's apartment number address
const VeerApt = "304"

// VeerCity | Veer's city address
const VeerCity = "Philadelphia"

// VeerState | Veer's state address
const VeerState = "PA"

// VeerZip | Veer's zipcode address
const VeerZip = "19121"
