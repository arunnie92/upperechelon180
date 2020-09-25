package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"../utils"
	"github.com/Pallinder/go-randomdata"
)

type creditCards []utils.CCInfo
type phantomProfiles []utils.Profile

// TODO: create 50 eno cards for each foot site, should have a total of 200 cards and 200 profiles
// TODO: add rules for specific sites
func main() {
	jsonFile, jsonFilErr := os.Open(utils.CreditCardPath)
	if jsonFilErr != nil {
		fmt.Println(jsonFilErr)
		return
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var creditCardInformation creditCards

	unmarshalErr := json.Unmarshal(byteValue, &creditCardInformation)
	if unmarshalErr != nil {
		fmt.Println(unmarshalErr)
		return
	}
	fmt.Println("Successfully uploaded credit card info...")

	fmt.Println("Creating profiles...")

	numberOfProfiles := len(creditCardInformation)

	var emptyProfiles phantomProfiles

	// TODO: export out to function
	// TODO: use range
	for creditCardIndex := 0; creditCardIndex < numberOfProfiles; creditCardIndex++ {
		var tempProfile utils.Profile

		//////////////////////////////////////////////
		// KEEP THESE FOR FUTURE REFERENCE
		// rand.Seed(time.Now().UTC().UnixNano())
		// creditCardIndex := utils.RandomIndex(0, 3)
		/////////////////////////////////////////////

		// Setup Credit Card Information
		tempProfile.CCNumber = creditCardInformation[creditCardIndex].CCNumber
		tempProfile.CVV = creditCardInformation[creditCardIndex].CVV
		tempProfile.ExpMonth = creditCardInformation[creditCardIndex].ExpMonth
		tempProfile.ExpYear = creditCardInformation[creditCardIndex].ExpYear
		tempProfile.CardType = creditCardInformation[creditCardIndex].CardType

		// Setup Phone Number
		tempProfile.Phone = utils.CreateRandomPhoneNumber()
		if len(tempProfile.Phone) != 10 {
			fmt.Println(fmt.Sprintf("wrong number at %s", tempProfile.Name))
		}

		// Setup Address
		tempProfile.Same = true

		address := utils.CreateAddress(utils.Address)

		tempProfile.Shipping.Address = address
		tempProfile.Shipping.Apt = utils.Apt
		tempProfile.Shipping.City = utils.City
		tempProfile.Shipping.State = utils.State
		tempProfile.Shipping.Zip = utils.Zip

		tempProfile.Billing.Address = address
		tempProfile.Billing.Apt = utils.Apt
		tempProfile.Billing.City = utils.City
		tempProfile.Billing.State = utils.State
		tempProfile.Billing.Zip = utils.Zip

		tempProfile.Country = utils.Country

		firstName := utils.ManipulateName(utils.FirstName)
		lastName := utils.ManipulateName(utils.LastName)

		site := creditCardInformation[creditCardIndex].Site
		if strings.Compare(site, utils.BestBuy) == 0 {
			randomdata.FirstName(randomdata.Male)
			randomdata.LastName()
		}

		tempProfile.Shipping.FirstName = firstName
		tempProfile.Shipping.LastName = lastName
		tempProfile.Billing.FirstName = firstName
		tempProfile.Billing.LastName = lastName

		// Setup Email
		tempProfile.Email = utils.CreateRandomEmail(tempProfile.Shipping.FirstName, tempProfile.Shipping.LastName)

		// Setup Profile Name
		tempProfile.Name = fmt.Sprintf("Profile_%d_%s", creditCardIndex, site)

		// Add Complete Profile
		emptyProfiles = append(emptyProfiles, tempProfile)
	}

	fmt.Println(fmt.Sprintf("%d profiles created...", numberOfProfiles))

	file, marshallErr := json.MarshalIndent(emptyProfiles, "", " ")
	if marshallErr != nil {
		fmt.Println(marshallErr)
		return
	}

	_ = ioutil.WriteFile("../../../real_profiles.json", file, 0644)

	fmt.Println("Finished creating profiles...")
}
