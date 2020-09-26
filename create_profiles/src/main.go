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

	var creditCardInformation []utils.CCInfo

	unmarshalErr := json.Unmarshal(byteValue, &creditCardInformation)
	if unmarshalErr != nil {
		fmt.Println(unmarshalErr)
		return
	}
	fmt.Println("Successfully uploaded credit card info...")

	fmt.Println("Creating profiles...")

	var profiles []utils.Profile

	// TODO: export out to function
	for index := range creditCardInformation {
		var profile utils.Profile

		//////////////////////////////////////////////
		// KEEP THESE FOR FUTURE REFERENCE
		// rand.Seed(time.Now().UTC().UnixNano())
		// index := utils.RandomIndex(0, 3)
		/////////////////////////////////////////////

		// Setup Credit Card Information
		profile.CCNumber = creditCardInformation[index].CCNumber
		profile.CVV = creditCardInformation[index].CVV
		profile.ExpMonth = creditCardInformation[index].ExpMonth
		profile.ExpYear = creditCardInformation[index].ExpYear
		profile.CardType = creditCardInformation[index].CardType

		// Setup Phone Number
		profile.Phone = utils.CreateRandomPhoneNumber()
		if len(profile.Phone) != 10 {
			fmt.Println(fmt.Sprintf("wrong number at %s", profile.Name))
		}

		// Setup Address
		profile.Same = true

		address := utils.CreateAddress(utils.Address)

		profile.Shipping.Address = address
		profile.Shipping.Apt = utils.Apt
		profile.Shipping.City = utils.City
		profile.Shipping.State = utils.State
		profile.Shipping.Zip = utils.Zip

		profile.Billing.Address = address
		profile.Billing.Apt = utils.Apt
		profile.Billing.City = utils.City
		profile.Billing.State = utils.State
		profile.Billing.Zip = utils.Zip

		profile.Country = utils.Country

		firstName := utils.ManipulateName(utils.FirstName)
		lastName := utils.ManipulateName(utils.LastName)

		site := creditCardInformation[index].Site
		if strings.Compare(site, utils.BestBuy) == 0 {
			randomdata.FirstName(randomdata.Male)
			randomdata.LastName()
		}

		profile.Shipping.FirstName = firstName
		profile.Shipping.LastName = lastName
		profile.Billing.FirstName = firstName
		profile.Billing.LastName = lastName

		// Setup Email
		profile.Email = utils.CreateRandomEmail(profile.Shipping.FirstName, profile.Shipping.LastName)

		// Setup Profile Name
		profile.Name = fmt.Sprintf("Profile_%d_%s", index, site)

		// Add Complete Profile
		profiles = append(profiles, profile)
	}

	numberOfProfiles := len(creditCardInformation)
	fmt.Println(fmt.Sprintf("%d profiles created...", numberOfProfiles))

	file, marshallErr := json.MarshalIndent(profiles, "", " ")
	if marshallErr != nil {
		fmt.Println(marshallErr)
		return
	}

	_ = ioutil.WriteFile("../../../profiles/all_profiles.json", file, 0644)

	fmt.Println("Finished creating profiles...")
}
