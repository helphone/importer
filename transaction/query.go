package transaction

import "errors"

// IsDatabaseEmpty tell you if the database is at his first pass without data inside
func (c *Connection) IsDatabaseEmpty() (bool, error) {
	var count int
	err := c.Tx.QueryRow("SELECT COUNT(DISTINCT country_code) from countries_translations").Scan(&count)
	return (count == 0), err
}

// RemoveCountriesTranslations does a truncate on the database
func (c *Connection) RemoveCountriesTranslations() error {
	_, err := c.Tx.Exec("TRUNCATE TABLE countries_translations")
	return err
}

// RemovePhonenumbersAndCategories does a truncate on the database
func (c *Connection) RemovePhonenumbersAndCategories() error {
	_, err := c.Tx.Exec("TRUNCATE TABLE phone_numbers_to_phone_numbers_categories, phone_numbers, phone_numbers_categories_translations, phone_numbers_categories")
	return err
}

// CreateCountryTranslation create an entry inside the database with the giving country code,
// the language code and the value of the translation
func (c *Connection) CreateCountryTranslation(countryCode string, languageCode string, value string) error {
	_, err := c.Tx.Exec("INSERT INTO countries_translations (country_code, language_code, value) VALUES ($1, $2, $3)", countryCode, languageCode, value)
	return err
}

// CreatePhonenumberCategory create an entry inside the database with the giving name
func (c *Connection) CreatePhonenumberCategory(name string) error {
	_, err := c.Tx.Exec("INSERT INTO phone_numbers_categories (name) VALUES ($1)", name)
	return err
}

// CreatePhonenumberCategoryTranslation create an entry inside the database with the giving category name,
// the language code and the value of the translation
func (c *Connection) CreatePhonenumberCategoryTranslation(categoryName string, languageCode string, value string) error {
	_, err := c.Tx.Exec("INSERT INTO phone_numbers_categories_translations (category_name, language_code, value) VALUES ($1, $2, $3)", categoryName, languageCode, value)
	return err
}

// CreatePhonenumber create an entry inside the database with the giving country code
// and the number to call
func (c *Connection) CreatePhonenumber(countryCode string, number string) error {
	_, err := c.Tx.Exec("INSERT INTO phone_numbers (country_code, phone_number) VALUES ($1, $2)", countryCode, number)
	return err
}

// AssignPhonenumberToCategory create an entry inside the database to link a phonenumber to
// a category. You must provide the country code, the value of the number and the category to assign
func (c *Connection) AssignPhonenumberToCategory(countryCode string, phonenumber string, categoryName string) error {
	id, err := c.getPhonenumberID(countryCode, phonenumber)
	if err != nil {
		return err
	}

	_, err = c.Tx.Exec("INSERT INTO phone_numbers_to_phone_numbers_categories (phone_number_id, phone_number_categories_name) VALUES ($1, $2)", id, categoryName)
	return err
}

// IsCountryExist returns a boolean regard of the country exist inside the database
func (c *Connection) IsCountryExist(code string) bool {
	err := c.Tx.QueryRow("SELECT code FROM countries WHERE code=$1", code).Scan(new(string))
	return err == nil
}

// IsLanguageExist returns a boolean regard of the language exist inside the database
func (c *Connection) IsLanguageExist(code string) bool {
	err := c.Tx.QueryRow("SELECT code FROM languages WHERE code=$1", code).Scan(new(string))
	return err == nil
}

// IsCategoryExist returns a boolean regard of the category exist inside the database
func (c *Connection) IsCategoryExist(name string) bool {
	err := c.Tx.QueryRow("SELECT name FROM phone_numbers_categories WHERE name = $1", name).Scan(new(string))
	return err == nil
}

func (c *Connection) getPhonenumberID(countryCode string, phonenumber string) (int32, error) {
	var id int32
	c.Tx.QueryRow("SELECT id FROM phone_numbers WHERE country_code = $1 AND phone_number = $2", countryCode, phonenumber).Scan(&id)
	if id == 0 {
		return id, errors.New("No phonenumber")
	}
	return id, nil
}
