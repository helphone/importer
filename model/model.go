package model

// YAMLCountries is a type that contains the list of countries
// from a YAML file YAMLCountries
type YAMLCountries struct {
	Countries map[string]Country
}

// Country is a type that contains translations of a country
type Country struct {
	Translations map[string]string
}

// YAMLPhoneNumbersCategories is a type that contains the list of phonenumber's
// category from a YAML file YAMLCountries
type YAMLPhoneNumbersCategories struct {
	Categories map[string]PhoneNumbersCategory
}

// PhoneNumbersCategory is a type that contains translations of a category
type PhoneNumbersCategory struct {
	Translations map[string]string
}

// YAMLPhoneNumbers is a type that contains the list of phonenumbers
// from a YAML file YAMLCountries
type YAMLPhoneNumbers struct {
	Phonenumbers map[string][]Phonenumber
}

// Phonenumber is a type that contains information about a phonenumber
// like his number and categories of it
type Phonenumber struct {
	Number     string
	Categories []string
}
