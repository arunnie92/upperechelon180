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

	profiles := []utils.Profile{}
	profileMap := make(map[string][]utils.Profile)

	for index, virutalCard := range creditCardInformation {
		newProfile := utils.CreateProfile(virutalCard, index)

		site := virutalCard.Site

		arr := []utils.Profile{}

		if profileMap[site] == nil {
			arr = append(arr, newProfile)
		} else {
			arr = profileMap[site]
			arr = append(arr, newProfile)
		}

		profileMap[site] = arr

		profiles = append(profiles, newProfile)
	}

	for siteKey, arrValue := range profileMap {
		file, marshallErr := json.MarshalIndent(arrValue, "", " ")
		if marshallErr != nil {
			fmt.Println(marshallErr)
			return
		}

		exportPath := fmt.Sprintf("%s/_%s_Profiles.json", utils.ProfilesPath, siteKey)

		writeFileErr := ioutil.WriteFile(exportPath, file, 0644)
		if writeFileErr != nil {
			fmt.Println(writeFileErr)
			return
		}

		fmt.Println(fmt.Sprintf("Finished creating %d %s profiles...", len(arrValue), siteKey))
	}

	numberOfProfiles := len(creditCardInformation)
	fmt.Println(fmt.Sprintf("%d profiles created...", numberOfProfiles))

	file, marshallErr := json.MarshalIndent(profiles, "", " ")
	if marshallErr != nil {
		fmt.Println(marshallErr)
		return
	}

	allProfilesPath := fmt.Sprintf("%s/%s", utils.ProfilesPath, "All_Profiles.json")
	writeFileErr := ioutil.WriteFile(allProfilesPath, file, 0644)
	if writeFileErr != nil {
		fmt.Println(writeFileErr)
		return
	}

	fmt.Println(fmt.Sprintf("Finished creating all %d profiles...", numberOfProfiles))
}
