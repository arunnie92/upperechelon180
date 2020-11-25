package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"../utils"
)

func createProfiles() ([]utils.Profile, error) {
	jsonFile, jsonFilErr := os.Open(utils.VirtualCreditCardPath)
	if jsonFilErr != nil {
		return nil, jsonFilErr
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var virtualCreditCardInformation []utils.VirtualCCInfo

	unmarshalErr := json.Unmarshal(byteValue, &virtualCreditCardInformation)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	fmt.Println("Successfully ingested virtual credit card information...")

	fmt.Println("Creating profiles...")

	profiles := []utils.Profile{}
	profileMap := make(map[string][]utils.Profile)

	previousSite := virtualCreditCardInformation[0].Site

	profilesCreated := 0

	for index, virtualCard := range virtualCreditCardInformation {
		if virtualCard.IsVeer {
			continue
		}

		currentSite := virtualCard.Site

		if !utils.SiteMap[currentSite] {
			continue
		}

		// instantiate map again from reusing full name
		if currentSite != previousSite {
			utils.FullNameMap = make(map[string]bool)
			previousSite = virtualCard.Site
		}

		newProfile := utils.CreateProfile(virtualCard, index)

		arr := []utils.Profile{}

		if profileMap[currentSite] == nil {
			arr = append(arr, newProfile)
		} else {
			arr = profileMap[currentSite]
			arr = append(arr, newProfile)
		}

		profileMap[currentSite] = arr

		profiles = append(profiles, newProfile)

		profilesCreated++
	}

	numOfExports := 0
	for siteKey, profileArr := range profileMap {
		exportPath := fmt.Sprintf("%s/_%s_Profiles.json", utils.ProfilesPath, siteKey)

		exportProfilesErr := utils.ExportData(exportPath, profileArr)
		if exportProfilesErr != nil {
			return nil, exportProfilesErr
		}

		numOfExports++

		fmt.Println(fmt.Sprintf("Finished creating %d %s profiles...", len(profileArr), siteKey))
	}

	// write all profiles to json file
	allProfilesPath := fmt.Sprintf("%s/%s", utils.ProfilesPath, "All_Profiles.json")
	exportAllProfilesErr := utils.ExportData(allProfilesPath, profiles)
	if exportAllProfilesErr != nil {
		return nil, exportAllProfilesErr
	}
	numOfExports++

	fmt.Println()
	fmt.Println(fmt.Sprintf("%d profiles created...", profilesCreated))
	fmt.Println(fmt.Sprintf("%d exported profile files...", numOfExports))
	fmt.Println(fmt.Sprintf("Finished creating profiles..."))

	fmt.Println()
	utils.CreateAndExportPhantomProlfileManager(profiles)

	return profiles, nil
}

func main() {
	/**
	 * TODO: MAKE SURE YOU UPDATE utils.SiteMap BEFORE YOU RUN SCRIPT
	 **/

	// create only footsite profiles
	profiles, profilesErr := createProfiles()
	if profilesErr != nil {
		fmt.Println(profilesErr)
		return
	}

	fmt.Println()

	skus := []string{
		"",
		"",
	}

	utils.CreateAndExportTasks(skus, profiles)
}
