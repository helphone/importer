package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

type YAMLCountries struct {
	Countries map[string]Country
}

type Country struct {
	Translations map[string]string
}

type YAMLPhoneNumbersCategories struct {
	Categories map[string]PhoneNumbersCategory
}

type PhoneNumbersCategory struct {
	Translations map[string]string
}

type YAMLPhoneNumbers struct {
	Phonenumbers map[string][]Phonenumber
}

type Phonenumber struct {
	Number     string
	Categories []string
}

func cloneRepo() {
	key := os.Getenv("GIT_ACCESS")

	_, err := os.Stat("./data")
	if err == nil {
		return
	}

	cmd := exec.Command("git", "clone", "https://Swatto:"+key+"@github.com/helphone/data.git", "data")
	_, err = cmd.Output()

	if err != nil {
		panic(err)
	}
}

func pullRepo() bool {
	cmd := exec.Command("cd", "data", "&&", "git", "pull", "origin", "master")
	_, err := cmd.Output()

	return err != nil
}

func importCountries(db *Database) {
	filename, _ := filepath.Abs("./data/countries.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var config YAMLCountries

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	db.Connection.Exec("TRUNCATE TABLE countries_translations")
	for country_code, country := range config.Countries {
		if !db.IsCountryExist(&country_code) {
			return
		}

		for language_code, translation := range country.Translations {
			if !db.IsLanguageExist(&language_code) {
				return
			}

			db.Connection.Exec("INSERT INTO countries_translations (country_code, language_code, value) VALUES ($1, $2, $3)", country_code, language_code, translation)
		}
	}
}

func importPhoneNumbersCategories(db *Database) {
	filename, _ := filepath.Abs("./data/phone_numbers_categories.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var config YAMLPhoneNumbersCategories

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	db.Connection.Exec("TRUNCATE TABLE phone_numbers_to_phone_numbers_categories; TRUNCATE TABLE phone_numbers; TRUNCATE TABLE; TRUNCATE TABLE phone_numbers_categories_translations; TRUNCATE TABLE phone_numbers_categories")
	for category_code, category := range config.Categories {
		db.Connection.Exec("INSERT INTO phone_numbers_categories (name) VALUES ($1)", category_code)
		for language_code, translation := range category.Translations {
			if !db.IsLanguageExist(&language_code) {
				return
			}

			db.Connection.Exec("INSERT INTO phone_numbers_categories_translations (category_name, language_code, value) VALUES ($1, $2, $3)", category_code, language_code, translation)
		}
	}
}

func importPhoneNumbers(db *Database) {
	filename, _ := filepath.Abs("./data/phone_numbers.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var config YAMLPhoneNumbers

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	for country_code, phonenumbers := range config.Phonenumbers {
		if !db.IsCountryExist(&country_code) {
			return
		}

		for _, phonenumber := range phonenumbers {
			db.Connection.Exec("INSERT INTO phone_numbers (country_code, phone_number) VALUES ($1, $2)", country_code, phonenumber.Number)

			var id int32
			db.Connection.QueryRow("SELECT id FROM phone_numbers WHERE country_code = $1 AND phone_number = $2", country_code, phonenumber.Number).Scan(&id)

			for _, category_code := range phonenumber.Categories {
				if !db.IsCategoryExist(&category_code) {
					return
				}

				db.Connection.Exec("INSERT INTO phone_numbers_to_phone_numbers_categories (phone_number_id, phone_number_categories_name) VALUES ($1, $2)", id, category_code)
			}
		}
	}
}

func main() {
	db := GenerateDatabase()
	cloneRepo()

	for {
		success := pullRepo()
		if success == true {
			importCountries(db)
			importPhoneNumbersCategories(db)
			importPhoneNumbers(db)

			time.Sleep(time.Duration(1) * time.Hour)
		} else {
			time.Sleep(time.Duration(15) * time.Minute)
		}
	}

}
