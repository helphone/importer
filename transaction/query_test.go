package transaction

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

var (
	nilMockResult           = sqlmock.NewResult(0, 0)
	oneRowCreatedMockResult = sqlmock.NewResult(1, 1)
)

func TestIsDatabaseEmpty(t *testing.T) {
	c, mock, err := generateConnection()
	if err != nil {
		t.Errorf("Error in the generation of the mock: err %v", err)
	}

	rows := sqlmock.NewRows([]string{"count"}).AddRow(0)
	mock.ExpectQuery("SELECT (.+) from countries_translations").WillReturnRows(rows)
	isIt, err := c.IsDatabaseEmpty()
	if err != nil {
		t.Errorf("Error in SQL: err %v", err)
	}
	if isIt != false {
		t.Errorf("The database should be empty: err %v", err)
	}

	generateResult(t, mock, err)
}

func TestRemoveCountriesTranslations(t *testing.T) {
	c, mock, err := generateConnection()
	if err != nil {
		t.Errorf("Error in the generation of the mock: err %v", err)
	}

	mock.ExpectExec("TRUNCATE TABLE countries_translations").WillReturnResult(nilMockResult)
	err = c.RemoveCountriesTranslations()
	if err != nil {
		t.Errorf("Error in SQL: err %v", err)
	}

	generateResult(t, mock, err)
}

func TestRemovePhonenumbersAndCategories(t *testing.T) {
	c, mock, err := generateConnection()
	if err != nil {
		t.Errorf("Error in the generation of the mock: err %v", err)
	}

	mock.ExpectExec("TRUNCATE TABLE phone_numbers_to_phone_numbers_categories, phone_numbers, phone_numbers_categories_translations, phone_numbers_categories").WillReturnResult(nilMockResult)
	err = c.RemovePhonenumbersAndCategories()
	if err != nil {
		t.Errorf("Error in SQL: err %v", err)
	}

	generateResult(t, mock, err)
}

func TestCreateCountryTranslation(t *testing.T) {
	c, mock, err := generateConnection()
	if err != nil {
		t.Errorf("Error in the generation of the mock: err %v", err)
	}

	const countryCode string = "BE"
	const languageCode string = "fr"
	const value string = "Belgique"

	mock.ExpectExec("INSERT INTO countries_translations").WithArgs(countryCode, languageCode, value).WillReturnResult(oneRowCreatedMockResult)
	err = c.CreateCountryTranslation(countryCode, languageCode, value)
	if err != nil {
		t.Errorf("Error in SQL: err %v", err)
	}

	generateResult(t, mock, err)
}

func TestCreatePhonenumberCategory(t *testing.T) {
	c, mock, err := generateConnection()
	if err != nil {
		t.Errorf("Error in the generation of the mock: err %v", err)
	}

	const name string = "police"

	mock.ExpectExec("INSERT INTO phone_numbers_categories").WithArgs(name).WillReturnResult(oneRowCreatedMockResult)
	err = c.CreatePhonenumberCategory(name)
	if err != nil {
		t.Errorf("Error in SQL: err %v", err)
	}

	generateResult(t, mock, err)
}

func TestCreatePhonenumberCategoryTranslation(t *testing.T) {
	c, mock, err := generateConnection()
	if err != nil {
		t.Errorf("Error in the generation of the mock: err %v", err)
	}

	const categoryName string = "urgency"
	const languageCode string = "fr"
	const value string = "urgence"

	mock.ExpectExec("INSERT INTO phone_numbers_categories_translations").WithArgs(categoryName, languageCode, value).WillReturnResult(oneRowCreatedMockResult)
	err = c.CreatePhonenumberCategoryTranslation(categoryName, languageCode, value)
	if err != nil {
		t.Errorf("Error in SQL: err %v", err)
	}

	generateResult(t, mock, err)
}

func TestCreatePhonenumber(t *testing.T) {
	c, mock, err := generateConnection()
	if err != nil {
		t.Errorf("Error in the generation of the mock: err %v", err)
	}

	const countryCode string = "BE"
	const number string = "112"

	mock.ExpectExec("INSERT INTO phone_numbers").WithArgs(countryCode, number).WillReturnResult(oneRowCreatedMockResult)
	err = c.CreatePhonenumber(countryCode, number)
	if err != nil {
		t.Errorf("Error in SQL: err %v", err)
	}

	generateResult(t, mock, err)
}

func TestAssignPhonenumberToCategory(t *testing.T) {
	c, mock, err := generateConnection()
	if err != nil {
		t.Errorf("Error in the generation of the mock: err %v", err)
	}

	const countryCode string = "BE"
	const number string = "112"
	const categoryName string = "urgency"

	mock.ExpectExec("INSERT INTO phone_numbers").WithArgs(countryCode, number).WillReturnResult(oneRowCreatedMockResult)
	err = c.CreatePhonenumber(countryCode, number)
	if err != nil {
		t.Errorf("Error in SQL preparation: err %v", err)
	}

	mock.ExpectQuery("SELECT (.+) FROM phone_numbers").WithArgs(countryCode, number).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectExec("INSERT INTO phone_numbers_to_phone_numbers_categories").WithArgs(1, categoryName).WillReturnResult(oneRowCreatedMockResult)
	err = c.AssignPhonenumberToCategory(countryCode, number, categoryName)
	if err != nil {
		t.Errorf("Error in SQL: err %v", err)
	}

	generateResult(t, mock, err)
}

func generateConnection() (c *Connection, mock sqlmock.Sqlmock, err error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return
	}

	mock.ExpectBegin()
	tx, err := db.Begin()
	if err != nil {
		return
	}

	c = &Connection{
		Tx: tx,
	}
	return
}

func generateResult(t *testing.T, mock sqlmock.Sqlmock, err error) {
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}
