package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"

	"github.com/Pallinder/go-randomdata"
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

// ManipulateName manipulates name string
func ManipulateName(name string) string {

	nameIndex := RandomIndex(0, len(name))

	runes := []rune(name)
	newName := fmt.Sprintf("%s%s%s",
		string(runes[0:nameIndex+1]),
		string(runes[nameIndex:nameIndex+1]),
		string(runes[nameIndex+1:len(name)]))

	return strings.Title(strings.ToLower(newName))
}

// ManipulateFullName | Manipulates full name and keeps track of used names
func ManipulateFullName() (string, string) {
	firstName := ManipulateName(FirstName)
	lastName := ManipulateName(LastName)

	fullName := fmt.Sprintf("%s %s", firstName, lastName)

	nameExists := FullNameMap[fullName]

	if !nameExists {
		FullNameMap[fullName] = true
		return firstName, lastName
	}

	return ManipulateFullName()
}

// CreateAddress creates an address with a concatenated alphanumerica value
func CreateAddress(address string) string {
	numeric := []string{"0", "1", "2"}
	alpha := []string{"A", "B", "C"}

	index := RandomIndex(0, 3)

	return fmt.Sprintf("%s %s%s%s", address, alpha[index], numeric[index], alpha[index])
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
	firstName, lastName := ManipulateFullName()
	// TODO: ^the problem here is if what if there are more profiles for a specific site being created than manipulations being created for that site

	// TODO: create rules function based on what site the profile is being used
	site := virtualCreditCard.Site
	if site == bestBuy {
		randomdata.FirstName(randomdata.Male)
		randomdata.LastName()
	}

	profile.Shipping.FirstName = firstName
	profile.Shipping.LastName = lastName
	profile.Billing.FirstName = firstName
	profile.Billing.LastName = lastName

	// Setup Email
	profile.Email = CreateRandomEmail(firstName, lastName)

	// Setup Profile Name
	profile.Name = fmt.Sprintf("Profile_%03d_%s", index, site)

	return profile
}

// IsFootSite | checks if the site a foot site
func IsFootSite(site string) bool {
	return SiteMap[site]
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

// CreateTask | initializes and creates a single task
func CreateTask(sku, site, profileName string, paypal bool) Task {
	return Task{
		URL:            sku,
		Size:           "R",
		Proxy:          "",
		Profile:        profileName,
		Site:           site,
		RandomEmail:    true, // TODO: Is this actually using a random email or the one I create?
		Desktop:        false,
		CheckoutMode:   "none",
		CaptchaSource:  "local",
		CartQuantity:   "",
		ProxyList:      "ProxyList",
		ManualCheckout: false,
		RepeatCheckout: false,
		MaxPrice:       "",
		PaypalCheckout: paypal,
	}
}

// CreateFiveTasks | creates five tasks
func CreateFiveTasks(sku, site, profileName string) []Task {
	tasks := []Task{}

	// TODO: Should I make all paypals?
	for i := 0; i < 3; i++ {
		task := CreateTask(sku, site, profileName, true)
		tasks = append(tasks, task)
	}

	for i := 0; i < 2; i++ {
		task := CreateTask(sku, site, profileName, false)
		tasks = append(tasks, task)
	}
	return tasks
}

// CreateAndExportTasks | creates 5 tasks per profile and exports all created tasks per sku
func CreateAndExportTasks(skus []string, profiles []Profile) {
	if len(skus) == 0 {
		fmt.Println("sku can not be empty")
		return
	}

	tasks := []Task{}

	for _, sku := range skus {
		fmt.Println(fmt.Sprintf("Creating tasks for %s", sku))

		for _, profile := range profiles {
			site := strings.Split(profile.Name, "_")[2]

			if site == all {
				for siteKey := range SiteMap {
					if siteKey == all {
						continue
					}

					newFiveTasks := CreateFiveTasks(sku, siteKey, profile.Name)
					tasks = append(tasks, newFiveTasks...)
				}
				continue
			}

			newFiveTasks := CreateFiveTasks(sku, site, profile.Name)
			tasks = append(tasks, newFiveTasks...)
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
