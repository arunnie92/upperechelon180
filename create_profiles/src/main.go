package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"../utils"
)

var useVeer bool
var useArunn bool

func createProfiles(useVeer, useArunn bool) ([]utils.Profile, error) {
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

	index := 0

	for _, virtualCard := range virtualCreditCardInformation {
		currentSite := virtualCard.Site

		if !utils.SiteMap[currentSite] {
			continue
		}

		// instantiate map again from reusing full name
		if currentSite != previousSite {
			utils.FullNameMap = make(map[string]bool)
			previousSite = virtualCard.Site
		}

		var newProfile utils.Profile

		if useArunn && useVeer {
			if virtualCard.IsVeer {
				newProfile = utils.CreateVeerProfile(virtualCard, index)
			} else {
				newProfile = utils.CreateProfile(virtualCard, index)
			}
		} else if useVeer && virtualCard.IsVeer {
			newProfile = utils.CreateVeerProfile(virtualCard, index)
		} else if useArunn && !virtualCard.IsVeer {
			newProfile = utils.CreateProfile(virtualCard, index)
		}	

		// TODO: is the a better way to check if a profile has been created?
		if len(newProfile.Name) > 0 {
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
			index++
		}
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

// TODO: MAKE SURE YOU SETUP FLAGS BEFORE YOU RUN SCRIPT
func init() {
	utils.SiteMap = map[string]bool{
		utils.FootLocker: true,
	}

	useVeer = true
	useArunn = true

	utils.Skus = map[string]string{
		utils.FootLocker: "1",
	}
}

func main() {
	profiles, profilesErr := createProfiles(useVeer, useArunn)
	if profilesErr != nil {
		fmt.Println(profilesErr)
		return
	}

	fmt.Println()

	utils.CreateAndExportTasks(profiles)
}
