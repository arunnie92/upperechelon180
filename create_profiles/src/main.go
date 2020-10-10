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

	footSiteProfilesArr := []utils.Profile{}

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

		if utils.IsFootSite(site) {
			footSiteProfilesArr = append(footSiteProfilesArr, newProfile)
		}
	}

	fmt.Println(fmt.Sprintf("Finished creating %d Foot Site profiles...", len(footSiteProfilesArr)))

	numOfExports := 0
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

		numOfExports++
		fmt.Println(fmt.Sprintf("Finished creating %d %s profiles...", len(arrValue), siteKey))
	}

	file, marshallErr := json.MarshalIndent(profiles, "", " ")
	if marshallErr != nil {
		fmt.Println(marshallErr)
		return
	}

	allProfilesPath := fmt.Sprintf("%s/%s", utils.ProfilesPath, "All_Profiles.json")
	allProfilesWriteFileErr := ioutil.WriteFile(allProfilesPath, file, 0644)
	if allProfilesWriteFileErr != nil {
		fmt.Println(allProfilesWriteFileErr)
		return
	}
	numOfExports++

	footSiteProfilesPath := fmt.Sprintf("%s/%s", utils.ProfilesPath, "FootSite_Profiles.json")
	footSiteProfilesWriteFileErr := ioutil.WriteFile(footSiteProfilesPath, file, 0644)
	if footSiteProfilesWriteFileErr != nil {
		fmt.Println(footSiteProfilesWriteFileErr)
		return
	}
	numOfExports++

	fmt.Println()
	fmt.Println(fmt.Sprintf("%d profiles created...", len(creditCardInformation))) // number of profiles created
	fmt.Println(fmt.Sprintf("%d exported profiles...", numOfExports))              // number of exported profiles
	fmt.Println(fmt.Sprintf("Finished creating profiles..."))
}
