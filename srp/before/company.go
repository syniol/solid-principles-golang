package before

import (
	"database/sql"
	"errors"
	"regexp"
	"unicode/utf8"
)

type Company struct {
	// todo: SQL lib
	dbClient *sql.DB
}

// FindCompanyWithName passes a lock by the value: type 'Company' contains 'sql.DB'
// contains 'atomic.Int64' contains 'interface{}' which is 'sync.Locker'
func (c *Company) FindCompanyWithName(
	name string,
) ([]string, error) {
	if utf8.RuneCountInString(name) == 0 {
		return nil, errors.New("")
	}

	isValidEnglishLetter := regexp.
		MustCompile(`^[a-zA-Z]+$`).
		MatchString(name)

	if !isValidEnglishLetter {
		return nil, errors.New("company name should be only English alphabets Aa-Zz")
	}

	res, err := c.dbClient.Query("SELECT name FROM companies WHERE name LIKE '?%'", name)
	if err != nil {
		return nil, err
	}

	var companyNames []string
	err = res.Scan(&companyNames)
	if err != nil {
		return nil, err
	}

	return companyNames, nil
}
