package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// CreateRandomEmail creates random email based on two strings
func CreateRandomEmail(firstName, lastName string) string {
	email := fmt.Sprintf("%s.%s@%s", firstName, lastName, domain)

	return strings.ToLower(email)
}

// CreateRandomPhoneNumber creates a random phone number as a string
func CreateRandomPhoneNumber() string {
	randomInteger := rand.Intn(9999999998-1000000001) + 1000000001

	return fmt.Sprintf("%d", randomInteger)
}

// RandomIndex returns a random number between max and min
func RandomIndex(min, max int) int {
	return rand.Intn(max-min) + min
}

// ManipulateString manipulates name string
func ManipulateString(name string) string {

	nameIndex := RandomIndex(0, len(name))

	runes := []rune(name)
	newName := fmt.Sprintf("%s%s%s",
		string(runes[0:nameIndex+1]),
		string(runes[nameIndex:nameIndex+1]),
		string(runes[nameIndex+1:len(name)]))

	return strings.ToLower(newName)
}

// ManipulateName |
func ManipulateName() (string, string) {
	manipulationChoice := RandomIndex(0, 3)

	firstName := FirstName
	lastName := LastName

	switch manipulationChoice {
	case 0:
		// manipulates both the first name and last name
		firstName = ManipulateString(FirstName)
		lastName = ManipulateString(LastName)
	case 1:
		// manipulates just the first name and not the last name
		firstName = ManipulateString(FirstName)
		lastName = LastName
	case 2:
		// manipulates not the first name and just the last name
		firstName = FirstName
		lastName = ManipulateString(LastName)
	default:
		// no manpulation
		firstName = FirstName
		lastName = LastName
	}

	fullName := fmt.Sprintf("%s %s", firstName, lastName)
	nameExists := FullNameMap[fullName]

	if !nameExists {
		FullNameMap[fullName] = true
		return firstName, lastName
	}

	return ManipulateName()
}

// CreateAddress creates an address with a concatenated alphanumerica value
func CreateAddress(address string) string {
	numeric := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	alpha := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}

	indexZero := RandomIndex(0, 10)
	indexOne := RandomIndex(0, 10)
	indexTwo := RandomIndex(0, 10)

	return fmt.Sprintf("%s %s%s%s", address, alpha[indexZero], numeric[indexOne], alpha[indexTwo])
}

// CreateProfile | returns a newly created profile based on an index and virtual card information
func CreateProfile(virtualCreditCard VirtualCCInfo, index int) Profile {
	var profile Profile

	//////////////////////////////////////////////
	// KEEP THESE FOR FUTURE REFERENCE
	// rand.Seed(time.Now().UTC().UnixNano())
	// index := RandomIndex(0, 3)
	/////////////////////////////////////////////

	// Setup Credit Card Information
	profile.CCNumber = virtualCreditCard.CCNumber
	profile.CVV = virtualCreditCard.CVV
	profile.ExpMonth = virtualCreditCard.ExpMonth
	profile.ExpYear = virtualCreditCard.ExpYear
	profile.CardType = virtualCreditCard.CardType

	// Setup Phone Number
	profile.Phone = CreateRandomPhoneNumber()
	if len(profile.Phone) != 10 {
		fmt.Println(fmt.Sprintf("wrong number at %s", profile.Name))
	}

	// Setup Address
	profile.Same = true

	address := CreateAddress(Address)

	profile.Shipping.Address = address
	profile.Shipping.Apt = Apt
	profile.Shipping.City = City
	profile.Shipping.State = State
	profile.Shipping.Zip = Zip

	profile.Billing.Address = address
	profile.Billing.Apt = Apt
	profile.Billing.City = City
	profile.Billing.State = State
	profile.Billing.Zip = Zip

	profile.Country = Country

	// Setup First Name & Last Name
	firstName, lastName := ManipulateName()
	// TODO: ^the problem here is if what if there are more profiles for a specific site being created than manipulations being created for that site

	// TODO: create rules function based on what site the profile is being used
	site := virtualCreditCard.Site
	// NOTE: Citi Cards work for all
	if site == All {
		// TODO: ADD specific site for Citi Virutal Cards
		site = "Site"
	}

	profile.Shipping.FirstName = firstName
	profile.Shipping.LastName = lastName
	profile.Billing.FirstName = firstName
	profile.Billing.LastName = lastName

	// Setup Email
	profile.Email = CreateRandomEmail(firstName, lastName)

	// Setup Profile Name
	profile.Name = fmt.Sprintf("Profile_%03d_Arunn_%s", index, site)

	return profile
}

// ExportData | exports data to json file
func ExportData(exportPath string, data interface{}) error {
	file, marshallErr := json.MarshalIndent(data, "", " ")
	if marshallErr != nil {
		return marshallErr
	}

	writeFileErr := ioutil.WriteFile(exportPath, file, 0644)
	if writeFileErr != nil {
		return writeFileErr
	}

	return nil
}

// CreateAndExportPhantomProlfileManager | creates and export profile manager
func CreateAndExportPhantomProlfileManager(profileArr []Profile) {
	fmt.Println(fmt.Sprintf("Creating Phantom's ProfileManager.json"))

	profileMap := make(map[string]Profile)

	for _, profile := range profileArr {
		profileMap[profile.Name] = profile
	}

	ExportData(ProfileManagerPath, profileMap)

	fmt.Println(fmt.Sprintf("Exported Phantom's ProfileManager.json"))
}

// CreateNonFootSiteTask | initializes and creates a single task for non foot sites
func CreateNonFootSiteTask(sku, site, profileName string) Task {
	return Task{
		Size:           "R",
		Proxy:          "",
		Profile:        profileName,
		Site:           site,
		URL:            sku,
		RandomEmail:    true, // TODO: Is this actually using a random email or the one I create?
		CheckoutMode:   "none",
		CaptchaSource:  "local",
		CartQuantity:   "1",
		ProxyList:      fmt.Sprintf("%s_ProxyList", site),
		ManualCheckout: false,
		RepeatCheckout: false,
		MaxPrice:       "",
		PaypalCheckout: false,
	}
}

// TODO: add comment for default values on why they are default
// CreateFootSiteTask | initializes and creates a single task for a foot site
func CreateFootSiteTask(sku, site, profileName string) Task {
	proxyList := fmt.Sprintf("%s_ProxyList", site)

	ProxyList[proxyList] = true

	return Task{
		URL:            sku,
		Size:           "R",
		Proxy:          "",
		Profile:        profileName,
		Site:           site,
		RandomEmail:    true,
		Desktop:        false,
		CheckoutMode:   "none",
		CaptchaSource:  "local",
		CartQuantity:   "",
		ProxyList:      proxyList,
		ManualCheckout: false,
		RepeatCheckout: false,
		MaxPrice:       "",
		PaypalCheckout: false,
	}
}

// CreateFiveTasks | creates five tasks
func CreateFiveTasks(sku, site, profileName string) []Task {
	tasks := []Task{}

	if FootSiteMap[site] {
		// TODO: Should I make all paypals?
		for i := 0; i < 5; i++ {
			task := CreateFootSiteTask(sku, site, profileName)
			tasks = append(tasks, task)
		}
	}

	if NonFootSiteMap[site] {
		for i := 0; i < 5; i++ {
			task := CreateNonFootSiteTask(sku, site, profileName)
			tasks = append(tasks, task)
		}
	}

	return tasks
}

// CreateAndExportTasks | creates 5 tasks per profile and exports all created tasks per sku
func CreateAndExportTasks(profiles []Profile) {
	if len(SiteSkusMap) == 0 {
		fmt.Println("sku can not be empty")
		return
	}

	tasks := []Task{}

	for siteKey, skusArr := range SiteSkusMap {
		for _, skuValue := range skusArr {
			fmt.Println(fmt.Sprintf("Creating tasks for sku %s for site %s", skuValue, siteKey))

			for _, profile := range profiles {
				profileSite := strings.Split(profile.Name, "_")[3]

				if profileSite == siteKey {
					newFiveTasks := CreateFiveTasks(skuValue, profileSite, profile.Name)
					tasks = append(tasks, newFiveTasks...)
				}
			}
		}
	}

	tasksPath := fmt.Sprintf("%s/master_tasks.json", absolutePath)
	exportTasksErr := ExportData(tasksPath, tasks)
	if exportTasksErr != nil {
		fmt.Println(exportTasksErr)
		return
	}

	fmt.Println(fmt.Sprintf("Created and exported all %d tasks", len(tasks)))
}

// ReadProfilesFromJSON | read json profiles
func ReadProfilesFromJSON(path string) ([]Profile, error) {
	jsonFile, jsonFilErr := os.Open(path)
	if jsonFilErr != nil {
		return nil, jsonFilErr
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var profiles []Profile

	unmarshalErr := json.Unmarshal(byteValue, &profiles)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return profiles, nil
}
