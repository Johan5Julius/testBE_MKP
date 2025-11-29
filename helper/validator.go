package helper

import (
	"errors"
	"regexp"
)

func ValidateEmail(email string) error {
	if email == "" {
		return errors.New("email tidak boleh kosong")
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(emailRegex, email)

	if !matched {
		return errors.New("format email tidak valid")
	}

	return nil
}

func ValidateLength(field string, value string, min int, max int) error {
	if value == "" {
		return errors.New(field + " tidak boleh kosong")
	}

	length := len(value)
	if length < min {
		return errors.New(field + " minimal " + string(rune(min)) + " karakter")
	}
	if length > max {
		return errors.New(field + " maksimal " + string(rune(max)) + " karakter")
	}
	return nil
}

func ValidateRequired(field string, value string) error {
	if value == "" {
		return errors.New(field + " tidak boleh kosong")
	}
	return nil
}
