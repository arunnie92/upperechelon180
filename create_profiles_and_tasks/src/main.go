package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"../utils"
)

func createFootSiteProfiles() ([]utils.Profile, error) {
	jsonFile, jsonFilErr := os.Open(utils.VirutalCreditCardPath)
	if jsonFilErr != nil {
		fmt.Println(jsonFilErr)
		return nil, jsonFilErr
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var virutalCreditCardInformation []utils.VirutalCCInfo

	unmarshalErr := json.Unmarshal(byteValue, &virutalCreditCardInformation)
	if unmarshalErr != nil {
		fmt.Println(unmarshalErr)
		return nil, unmarshalErr
	}
	fmt.Println("Successfully uploaded virutal credit card information...")

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

		exportProfilesErr := utils.ExportJSON(exportPath, profileArr)
		if exportProfilesErr != nil {
			fmt.Println(exportProfilesErr)
			return nil, exportProfilesErr
		}

		numOfExports++

		fmt.Println(fmt.Sprintf("Finished creating %d %s profiles...", len(profileArr), siteKey))
	}

	// write all profiles to json file
	allProfilesPath := fmt.Sprintf("%s/%s", utils.ProfilesPath, "All_FootSite_Profiles.json")
	exportAllProfilesErr := utils.ExportJSON(allProfilesPath, profiles)
	if exportAllProfilesErr != nil {
		fmt.Println(exportAllProfilesErr)
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

func createFootSiteTasks(sku string) {
	profiles, createFootSiteProfilesErr := createFootSiteProfiles()
	if createFootSiteProfilesErr != nil {
		fmt.Println(createFootSiteProfilesErr)
		return
	}

	proxyListName := "proxy_list" // TODO: figure this out

	fmt.Println("")
	fmt.Println("Creating tasks...")
	tasks := []utils.Task{}
	for _, profile := range profiles {
		site := strings.Split(profile.Name, "_")[2]

		for i := 0; i < 5; i++ {
			newTask := utils.CreateTask(site, profile.Name, sku, proxyListName)

			tasks = append(tasks, newTask)
		}
	}
	fmt.Println(fmt.Sprintf("%d task created...", len(tasks)))

	allTasksPath := fmt.Sprintf("%s/%s", utils.TasksPath, "tasks.json")
	exportAllTasksErr := utils.ExportJSON(allTasksPath, tasks)
	if exportAllTasksErr != nil {
		fmt.Println(exportAllTasksErr)
		return
	}

	fmt.Println("All tasks have been exported...")
}

func createProfiles() {
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

		exportProfilesErr := utils.ExportJSON(exportPath, profileArr)
		if exportProfilesErr != nil {
			fmt.Println(exportProfilesErr)
			return
		}

		numOfExports++

		fmt.Println(fmt.Sprintf("Finished creating %d %s profiles...", len(profileArr), siteKey))
	}

	// write all profiles to json file
	allProfilesPath := fmt.Sprintf("%s/%s", utils.ProfilesPath, "All_Profiles.json")
	exportAllProfilesErr := utils.ExportJSON(allProfilesPath, profiles)
	if exportAllProfilesErr != nil {
		fmt.Println(exportAllProfilesErr)
		return
	}
	numOfExports++

	// write only foot site profiles to json file
	footSiteProfilesPath := fmt.Sprintf("%s/%s", utils.ProfilesPath, "FootSite_Profiles.json")
	exportFootSiteProfilesErr := utils.ExportJSON(footSiteProfilesPath, footSiteProfilesArr)
	if exportFootSiteProfilesErr != nil {
		fmt.Println(exportFootSiteProfilesErr)
		return
	}
	numOfExports++

	fmt.Println()
	fmt.Println(fmt.Sprintf("%d profiles created...", len(virutalCreditCardInformation))) // number of profiles created
	fmt.Println(fmt.Sprintf("%d exported profile files...", numOfExports))                // number of exported profiles
	fmt.Println(fmt.Sprintf("Finished creating profiles..."))
}

func main() {
	// create only footsite profiles
	createFootSiteTasks("sku")
}
