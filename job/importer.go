package job

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/helphone/importer/model"
	"github.com/helphone/importer/transaction"
	"gopkg.in/yaml.v2"
)

// Refresh will launch the process of updating the database
func Refresh() {
	log.Info("Refresh started")
	conn, err := transaction.CreateConnection()

	if err != nil {
		log.Info("Error in database connection")
		return
	}

	defer conn.Finish(err)

	path, _ := filepath.Abs("/etc/data/countries.yml")
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	var countries model.YAMLCountries
	err = yaml.Unmarshal(yamlFile, &countries)
	if err != nil {
		return
	}

	if err = importCountries(conn, countries); err != nil {
		log.Infof("Error in import of countries, err: %v", err)
		return
	}

	path, _ = filepath.Abs("/etc/data/phone_numbers_categories.yml")
	yamlFile, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}

	var categories model.YAMLPhoneNumbersCategories
	err = yaml.Unmarshal(yamlFile, &categories)
	if err != nil {
		return
	}

	if err = importPhoneNumbersCategories(conn, categories); err != nil {
		log.Infof("Error in import of categories, err: %v", err)
		return
	}

	path, _ = filepath.Abs("/etc/data/phone_numbers.yml")
	yamlFile, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}

	var phonenumbers model.YAMLPhoneNumbers
	err = yaml.Unmarshal(yamlFile, &phonenumbers)
	if err != nil {
		return
	}

	if err = importPhoneNumbers(conn, phonenumbers); err != nil {
		log.Infof("Error in import of phonenumbers, err: %v", err)
		return
	}
	log.Info("Refresh finished")
}

func importCountries(conn *transaction.Connection, config model.YAMLCountries) (err error) {
	log.Info("Import countries started")

	if err = conn.RemoveCountriesTranslations(); err != nil {
		return
	}
	for countryCode, country := range config.Countries {
		if conn.IsCountryExist(countryCode) == false {
			err = errors.New("Missing country")
			break
		}

		errorInsideTranslation := false
		for languageCode, translation := range country.Translations {
			if conn.IsLanguageExist(languageCode) == false {
				errorInsideTranslation = true
				err = errors.New("No language")
				break
			}

			if err = conn.CreateCountryTranslation(countryCode, languageCode, translation); err != nil {
				errorInsideTranslation = true
				break
			}
		}

		if errorInsideTranslation == true {
			break
		}
	}
	log.Info("Import countries finished")
	return
}

func importPhoneNumbersCategories(conn *transaction.Connection, config model.YAMLPhoneNumbersCategories) (err error) {
	log.Info("Import phonenumber categories started")

	if err = conn.RemovePhonenumbersAndCategories(); err != nil {
		return
	}

	for categoryCode, category := range config.Categories {
		if err = conn.CreatePhonenumberCategory(categoryCode); err != nil {
			break
		}

		errorInsideTranslation := false
		for languageCode, translation := range category.Translations {
			if conn.IsLanguageExist(languageCode) == false {
				errorInsideTranslation = true
				err = errors.New("No language")
				break
			}

			if err = conn.CreatePhonenumberCategoryTranslation(categoryCode, languageCode, translation); err != nil {
				errorInsideTranslation = true
				break
			}
		}

		if errorInsideTranslation == true {
			break
		}
	}
	log.Info("Import phonenumber categories finished")
	return
}

func importPhoneNumbers(conn *transaction.Connection, config model.YAMLPhoneNumbers) (err error) {
	log.Info("Import phonenumbers started")

	for countryCode, phonenumbers := range config.Phonenumbers {
		if conn.IsCountryExist(countryCode) == false {
			err = errors.New("No country")
			break
		}

		errorInsidePhonenumber := false
		for _, phonenumber := range phonenumbers {
			if err = conn.CreatePhonenumber(countryCode, phonenumber.Number); err != nil {
				errorInsidePhonenumber = true
				break
			}

			errorInsideAssignement := false
			for _, categoryCode := range phonenumber.Categories {
				if conn.IsCategoryExist(categoryCode) == false {
					err = errors.New("No category " + categoryCode)
					errorInsideAssignement = true
					break
				}

				if err = conn.AssignPhonenumberToCategory(countryCode, phonenumber.Number, categoryCode); err != nil {
					errorInsideAssignement = true
					break
				}
			}

			if errorInsideAssignement == true {
				break
			}
		}

		if errorInsidePhonenumber == true {
			break
		}
	}
	log.Info("Import phonenumbers finished")
	return
}
