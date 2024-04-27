package company

import (
	"errors"
	"regexp"
	"unicode/utf8"
)

func ValidateCompanyName(name string) error {
	hasLengthValidationError := validateCompanyNameLength(name)
	if hasLengthValidationError != nil {
		return hasLengthValidationError
	}

	isValidEnglishLetter := validateCompanyNameAlphabets(name)
	if isValidEnglishLetter != nil {
		return isValidEnglishLetter
	}

	return nil
}

func validateCompanyNameAlphabets(name string) error {
	isValidEnglishLetter := regexp.
		MustCompile(`^[a-zA-Z]+$`).
		MatchString(name)

	if !isValidEnglishLetter {
		return errors.New("company name should be only English alphabets Aa-Zz")
	}

	return nil
}

func validateCompanyNameLength(name string) error {
	if utf8.RuneCountInString(name) == 0 {
		return errors.New("company name should be at least one character long")
	}

	return nil
}
