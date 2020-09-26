package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"../utils"
)

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

	for index, virutalCard := range creditCardInformation {
		newProfile := utils.CreateProfile(virutalCard, index)

		profiles = append(profiles, newProfile)
	}

	numberOfProfiles := len(creditCardInformation)
	fmt.Println(fmt.Sprintf("%d profiles created...", numberOfProfiles))

	file, marshallErr := json.MarshalIndent(profiles, "", " ")
	if marshallErr != nil {
		fmt.Println(marshallErr)
		return
	}

	writeFileErr := ioutil.WriteFile(utils.AllProfilesPath, file, 0644)
	if writeFileErr != nil {
		fmt.Println(writeFileErr)
		return
	}

	fmt.Println("Finished creating profiles...")
}
