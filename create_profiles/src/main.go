package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"../utils"
)

func main() {
	jsonFile, jsonFilErr := os.Open(utils.VirutalCreditCardPath)
	if jsonFilErr != nil {
		fmt.Println(jsonFilErr)
		return
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var virutalCreditCardInformation []utils.VirutalCCInfo

	unmarshalErr := json.Unmarshal(byteValue, &virutalCreditCardInformation)
	if unmarshalErr != nil {
		fmt.Println(unmarshalErr)
		return
	}
	fmt.Println("Successfully uploaded virutal credit card information...")

	fmt.Println("Creating profiles...")

	profiles := []utils.Profile{}
	profileMap := make(map[string][]utils.Profile)

	footSiteProfilesArr := []utils.Profile{}

	for index, virutalCard := range virutalCreditCardInformation {
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

	// write all profiles to json file
	allProfilesFile, marshallErr := json.MarshalIndent(profiles, "", " ")
	if marshallErr != nil {
		fmt.Println(marshallErr)
		return
	}
	allProfilesPath := fmt.Sprintf("%s/%s", utils.ProfilesPath, "All_Profiles.json")
	allProfilesWriteFileErr := ioutil.WriteFile(allProfilesPath, allProfilesFile, 0644)
	if allProfilesWriteFileErr != nil {
		fmt.Println(allProfilesWriteFileErr)
		return
	}
	numOfExports++

	// write only foot site profiles to json file
	footSiteProfilesFile, marshallErr := json.MarshalIndent(footSiteProfilesArr, "", " ")
	if marshallErr != nil {
		fmt.Println(marshallErr)
		return
	}
	footSiteProfilesPath := fmt.Sprintf("%s/%s", utils.ProfilesPath, "FootSite_Profiles.json")
	footSiteProfilesWriteFileErr := ioutil.WriteFile(footSiteProfilesPath, footSiteProfilesFile, 0644)
	if footSiteProfilesWriteFileErr != nil {
		fmt.Println(footSiteProfilesWriteFileErr)
		return
	}
	numOfExports++

	fmt.Println()
	fmt.Println(fmt.Sprintf("%d profiles created...", len(virutalCreditCardInformation))) // number of profiles created
	fmt.Println(fmt.Sprintf("%d exported profile files...", numOfExports))                // number of exported profiles
	fmt.Println(fmt.Sprintf("Finished creating profiles..."))
}
