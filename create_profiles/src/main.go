package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"../utils"
)

func createFootSiteProfiles() {
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
	fmt.Println("Successfully ingested virutal credit card information...")

	fmt.Println("Creating profiles...")

	profiles := []utils.Profile{}
	profileMap := make(map[string][]utils.Profile)

	previousSite := virutalCreditCardInformation[0].Site

	profilesCreated := 0

	for index, virutalCard := range virutalCreditCardInformation {
		currentSite := virutalCard.Site

		if !utils.FootSitesMap[currentSite] {
			continue
		}

		// instantiate map again from reusing full name
		if currentSite != previousSite {
			utils.FullNameMap = make(map[string]bool)
			previousSite = virutalCard.Site
		}

		newProfile := utils.CreateProfile(virutalCard, index)

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

		exportProfilesErr := utils.ExportProfiles(exportPath, profileArr)
		if exportProfilesErr != nil {
			fmt.Println(exportProfilesErr)
			return
		}

		numOfExports++

		fmt.Println(fmt.Sprintf("Finished creating %d %s profiles...", len(profileArr), siteKey))
	}

	// write all profiles to json file
	allProfilesPath := fmt.Sprintf("%s/%s", utils.ProfilesPath, "All_FootSite_Profiles.json")
	exportAllProfilesErr := utils.ExportProfiles(allProfilesPath, profiles)
	if exportAllProfilesErr != nil {
		fmt.Println(exportAllProfilesErr)
		return
	}
	numOfExports++

	fmt.Println()
	fmt.Println(fmt.Sprintf("%d profiles created...", profilesCreated))
	fmt.Println(fmt.Sprintf("%d exported profile files...", numOfExports))
	fmt.Println(fmt.Sprintf("Finished creating profiles..."))

	fmt.Println()
	utils.CreateAndExportPhantomProlfileManager(profiles)
}

func createAllProfiles() {
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
	fmt.Println("Successfully ingested virutal credit card information...")

	fmt.Println("Creating profiles...")

	profiles := []utils.Profile{}
	profileMap := make(map[string][]utils.Profile)

	footSiteProfilesArr := []utils.Profile{}

	previousSite := virutalCreditCardInformation[0].Site

	for index, virutalCard := range virutalCreditCardInformation {
		currentSite := virutalCard.Site

		// instantiate map again from reusing full name
		if currentSite != previousSite {
			utils.FullNameMap = make(map[string]bool)
			previousSite = virutalCard.Site
		}

		newProfile := utils.CreateProfile(virutalCard, index)

		arr := []utils.Profile{}

		if profileMap[currentSite] == nil {
			arr = append(arr, newProfile)
		} else {
			arr = profileMap[currentSite]
			arr = append(arr, newProfile)
		}

		profileMap[currentSite] = arr

		profiles = append(profiles, newProfile)

		if utils.IsFootSite(currentSite) {
			footSiteProfilesArr = append(footSiteProfilesArr, newProfile)
		}
	}

	fmt.Println(fmt.Sprintf("Finished creating %d Foot Site profiles...", len(footSiteProfilesArr)))

	numOfExports := 0
	for siteKey, profileArr := range profileMap {
		exportPath := fmt.Sprintf("%s/_%s_Profiles.json", utils.ProfilesPath, siteKey)

		exportProfilesErr := utils.ExportProfiles(exportPath, profileArr)
		if exportProfilesErr != nil {
			fmt.Println(exportProfilesErr)
			return
		}

		numOfExports++

		fmt.Println(fmt.Sprintf("Finished creating %d %s profiles...", len(profileArr), siteKey))
	}

	// write all profiles to json file
	allProfilesPath := fmt.Sprintf("%s/%s", utils.ProfilesPath, "All_Profiles.json")
	exportAllProfilesErr := utils.ExportProfiles(allProfilesPath, profiles)
	if exportAllProfilesErr != nil {
		fmt.Println(exportAllProfilesErr)
		return
	}
	numOfExports++

	// write only foot site profiles to json file
	footSiteProfilesPath := fmt.Sprintf("%s/%s", utils.ProfilesPath, "FootSite_Profiles.json")
	exportFootSiteProfilesErr := utils.ExportProfiles(footSiteProfilesPath, footSiteProfilesArr)
	if exportFootSiteProfilesErr != nil {
		fmt.Println(exportFootSiteProfilesErr)
		return
	}
	numOfExports++

	fmt.Println()
	fmt.Println(fmt.Sprintf("%d profiles created...", len(virutalCreditCardInformation))) // number of profiles created
	fmt.Println(fmt.Sprintf("%d exported profile files...", numOfExports))                // number of exported profiles
	fmt.Println(fmt.Sprintf("Finished creating profiles..."))

	fmt.Println()
	utils.CreateAndExportPhantomProlfileManager(profiles)
}

func main() {
	// create all profiles
	// createAllProfiles()

	// create only footsite profiles
	createFootSiteProfiles()
}
