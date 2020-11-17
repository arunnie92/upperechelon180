package utils

const absolutePath = "/Users/arunnie92/Documents/upper_echelon_180"

// VirtualCreditCardPath | path of virtual credit card data
const VirtualCreditCardPath = absolutePath + "/upperechelon180/create_profiles/data/eno.json"

// ProfilesPath | path to export all profiles
const ProfilesPath = absolutePath + "/profiles"

// ProfileManagerPath | absolute path to export Phantom Profile Manager
const ProfileManagerPath = "/Users/arunnie92/Library/ApplicationSupport/Phantom/ProfileManager.json"

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
const walmart = "Walmart"

const yeezySupply = "YeezySupply"
const adidas = "AdidasUS"
const footLocker = "FootLockerUS"
const footAction = "FootAction"
const champsSports = "ChampsSports"
const eastbay = "Eastbay"

// FootSitesArr | array of foot sites
var FootSitesArr = []string{footLocker, footAction, champsSports, eastbay}

// FootSitesMap | map of foot sitres
var FootSitesMap = map[string]bool{
	footLocker:   true,
	footAction:   true,
	champsSports: true,
	eastbay:      true,
}
