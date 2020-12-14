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

const bestBuy = "BestBuy" // site keyword
const walmart = "Walmart" // site keyword

// YeezySupply | site keyword
const YeezySupply = "YeezySupply"

const adidas = "AdidasUS" // site keyword

// FootLocker | site keyword
const FootLocker = "FootLockerUS"

// FootAction | site keyword
const FootAction = "FootAction"

// ChampsSports | site keyword
const ChampsSports = "ChampsSports"

// Eastbay | site keyword
const Eastbay = "Eastbay"

const All = "All" // site keyword to use for all sites

// SiteMap | map of sites
var SiteMap = make(map[string]bool)

// FootSiteMap | map of foot sites
var FootSiteMap = map[string]bool{
	FootLocker:   true,
	FootAction:   true,
	ChampsSports: true,
	Eastbay:      true,
}

// NonFootSiteMap | map of non foot sites
var NonFootSiteMap = map[string]bool{
	adidas:      true,
	YeezySupply: true,
	bestBuy:     true,
	walmart:     true,
}

// Skus | map of sites
var Skus = make(map[string]string)

/*
 * Veer's Information
 */

// VeerLastName | Veer's last name
const VeerLastName = "patel"

// KushFirstName | Kush's first name
const KushFirstName = "kush"

// VeerFirstName | Veer's first name
const VeerFirstName = "veer"

// VeerAddress | Veer's address
const VeerAddress = "1501 N 31st Street"

// VeerApt | Veer's apartment number address
const VeerApt = "304"

// VeerCity | Veer's city address
const VeerCity = "Philadelphia"

// VeerState | Veer's state address
const VeerState = "PA"

// VeerZip | Veer's zipcode address
const VeerZip = "19121"

const charlesSchwab = "CharlesSchwab"
