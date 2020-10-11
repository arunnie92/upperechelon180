package utils

// VirutalCCInfo | Store credit card information
type VirutalCCInfo struct {
	CCNumber string `json:"CCNumber"` // CCNumber | 16 digit credit card number
	CVV      string `json:"CVV"`      // CVV | 3 digit secruity code for credit card
	ExpMonth string `json:"ExpMonth"` // ExpMonth | expirtation month of credit card
	ExpYear  string `json:"ExpYear"`  // ExpYear | expiration year of credit card
	CardType string `json:"CardType"` // CardType | type of credit card for payment
	Site     string `json:"Site"`     // Site | the site the credit card belongs too
}

type address struct {
	FirstName string `json:"FirstName"` // FirstName | first name of profile of address
	LastName  string `json:"LastName"`  // LastName | first name of profile of address
	Address   string `json:"Address"`   // Address | address of profile
	Apt       string `json:"Apt"`       // Apt | apt number of address
	City      string `json:"City"`      // City | address city of profile
	State     string `json:"State"`     // State | state of address of profile
	Zip       string `json:"Zip"`       // Zip | state of zip of profile
}

// Profile | profile used at checkout
type Profile struct {
	CCNumber string  `json:"CCNumber"` // CCNumber | 16 digit credit card number
	CVV      string  `json:"CVV"`      // CVV | 3 digit secruity code for credit card
	ExpMonth string  `json:"ExpMonth"` // ExpMonth | expirtation month of credit card
	ExpYear  string  `json:"ExpYear"`  // ExpYear | expiration year of credit card
	CardType string  `json:"CardType"` // CardType | type of credit card for payment
	Email    string  `json:"Email"`    // Email | emaill address of profile
	Same     bool    `json:"Same"`     // Same | shipping and billing address are the same
	Shipping address `json:"Shipping"` // Shipping | shipping address
	Billing  address `json:"Billing"`  // Billing | billing address
	Phone    string  `json:"Phone"`    // Phone phone number
	Name     string  `json:"Name"`     // Name | name of profile
	Country  string  `json:"Country"`  // Country | country of address
}
