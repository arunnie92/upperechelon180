package utils

const absolutePath = "/Users/arunnie92/Documents/upper_echelon_180"

// VirutalCreditCardPath | path of virutal credit card data
const VirutalCreditCardPath = absolutePath + "/upperechelon180/create_profiles_and_tasks/data/eno.json"

// ProfilesPath | path to export all profiles
const ProfilesPath = absolutePath + "/profiles"

// ProfileManagerPath | absolute path to export Phantom Profile Manager
const ProfileManagerPath = "/Users/arunnie92/Library/ApplicationSupport/Phantom/ProfileManager.json"

// TasksPath | path to export tasks
const TasksPath = absolutePath + "/tasks"

// LastName | last name of profile
const LastName = "chanthirakanthan"

// FirstName | first name of profile
const FirstName = "arunn"

// FullNameMap | keeps track of used full names for profiles 
var FullNameMap = make(map[string]bool)

// Address | first line of address
const Address = "105 Esplanade Ave"

// Apt | apt number of address
const Apt = "45"

// City | city of address
const City = "Pacifica"

// State | state of address
const State = "CA"

// Zip | zipcode of address
const Zip = "94044"

// Country | country of address
const Country = "US"

const domain = "upperechelon180.com"

const bestBuy = "BestBuy"
const yeezySupply = "YeezySupply"
const adidas = "AdidasUS"
const footLocker = "FootLockerUS"
const footAction = "FootAction"
const champsSports = "ChampsSports"
const eastbay = "Eastbay"

// FootSites | array of foot sites
var FootSitesArr = []string{footLocker, footAction, champsSports, eastbay}

// FootSitesMap | mp of foot sitres
var FootSitesMap = map[string]bool{
	footLocker:   true,
	footAction:   true,
	champsSports: true,
	eastbay:      true,
}

// objects used in Tasks
const size = "R"
const checkoutModeNone = "none"
const captchaSourceLocal = "local"
