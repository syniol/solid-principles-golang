package before

import (
	"errors"
	"fmt"
	"regexp"
	"unicode/utf8"
)

type Company struct {
	// todo: SQL lib
	dbClient any
}

func (c Company) FindCompanyByName(name string) ([]string, error) {
	if utf8.RuneCountInString(name) == 0 {
		return nil, errors.New("")
	}

	isValidEnglishLetter := regexp.
		MustCompile(`^[a-zA-Z]+$`).
		MatchString(name)

	if !isValidEnglishLetter {
		return nil, errors.New("company name should be only English alphabets Aa-Zz")
	}

	c.dbClient(fmt.Sprintf("SELECT name FROM companies WHERE name LIKE '%%s%'", name))
	// todo: return names as slice

	return nil, nil
}
