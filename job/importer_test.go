package job

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/helphone/importer/model"
	"github.com/helphone/importer/transaction"
	"gopkg.in/yaml.v2"
)

func TestBadImportCountries(t *testing.T) {
	conn, err := transaction.CreateConnection()
	if err != nil {
		t.Errorf("Connection creation failed, err: %v", err)
	}

	var countries model.YAMLCountries
	path, _ := filepath.Abs("/etc/data/bad_countries.yml")
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		t.Errorf("Error during the reading of the YAML, err: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &countries)
	if err != nil {
		t.Errorf("Error during the reading of the YAML file, err: %v", err)
	}

	err = importCountries(conn, countries)
	if err == nil {
		t.Error("The import should fails with missing country and translation")
	}

	conn.Finish(err)
}

func TestGoodImportCountries(t *testing.T) {
	conn, err := transaction.CreateConnection()
	if err != nil {
		t.Errorf("Connection creation failed, err: %v", err)
	}

	var countries model.YAMLCountries
	path, _ := filepath.Abs("/etc/data/good_countries.yml")
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		t.Errorf("Error during the reading of the YAML, err: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &countries)
	if err != nil {
		t.Errorf("Error during the reading of the YAML file, err: %v", err)
	}

	err = importCountries(conn, countries)
	if err != nil {
		t.Errorf("The import should be ok, err: %v", err)
	}

	conn.Finish(err)
}

func TestBadImportPhoneNumbersCategories(t *testing.T) {
	conn, err := transaction.CreateConnection()
	if err != nil {
		t.Errorf("Connection creation failed, err: %v", err)
	}

	var categories model.YAMLPhoneNumbersCategories
	path, _ := filepath.Abs("/etc/data/bad_phone_numbers_categories.yml")
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		t.Errorf("Error during the reading of the YAML, err: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &categories)
	if err != nil {
		t.Errorf("Error during the reading of the YAML file, err: %v", err)
	}

	err = importPhoneNumbersCategories(conn, categories)
	if err == nil {
		t.Error("The import should fails with missing language and repeating name")
	}

	conn.Finish(err)
}

func TestGoodImportPhoneNumbersCategories(t *testing.T) {
	conn, err := transaction.CreateConnection()
	if err != nil {
		t.Errorf("Connection creation failed, err: %v", err)
	}

	var categories model.YAMLPhoneNumbersCategories
	path, _ := filepath.Abs("/etc/data/good_phone_numbers_categories.yml")
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		t.Errorf("Error during the reading of the YAML, err: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &categories)
	if err != nil {
		t.Errorf("Error during the reading of the YAML file, err: %v", err)
	}

	err = importPhoneNumbersCategories(conn, categories)
	if err != nil {
		t.Errorf("The import should be ok, err: %v", err)
	}

	conn.Finish(err)
}

func TestBadImportPhoneNumbers(t *testing.T) {
	conn, err := transaction.CreateConnection()
	if err != nil {
		t.Errorf("Connection creation failed, err: %v", err)
	}

	var phonenumbers model.YAMLPhoneNumbers
	path, _ := filepath.Abs("/etc/data/bad_phone_numbers.yml")
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		t.Errorf("Error during the reading of the YAML, err: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &phonenumbers)
	if err != nil {
		t.Errorf("Error during the reading of the YAML file, err: %v", err)
	}

	err = importPhoneNumbers(conn, phonenumbers)
	if err == nil {
		t.Error("The import should fails with missing country and category")
	}

	conn.Finish(err)
}

func TestGoodImportPhoneNumbers(t *testing.T) {
	conn, err := transaction.CreateConnection()
	if err != nil {
		t.Errorf("Connection creation failed, err: %v", err)
	}

	var phonenumbers model.YAMLPhoneNumbers
	path, _ := filepath.Abs("/etc/data/good_phone_numbers.yml")
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		t.Errorf("Error during the reading of the YAML, err: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &phonenumbers)
	if err != nil {
		t.Errorf("Error during the reading of the YAML file, err: %v", err)
	}

	err = importPhoneNumbers(conn, phonenumbers)
	if err != nil {
		t.Errorf("The import should be ok, err: %v", err)
	}

	conn.Finish(err)
}
