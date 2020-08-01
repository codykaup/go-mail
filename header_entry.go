package gomail

import (
	"fmt"
	"strings"
)

type HeaderEntry struct {
	Key   string
	Value []string
}

// NewHeaderEntry generates a new HeaderEntry after format validation
//
// Validation is based on RFC-5322: https://tools.ietf.org/html/rfc5322
//
// fieldName - The name of the field (like Date/From/Subject)
// value     - The value for the field (use commas to signal multiple values)
func NewHeaderEntry(fieldName, value string) (HeaderEntry, error) {
	// ensure it's a valid field
	if !validField(fieldName) {
		return HeaderEntry{}, fmt.Errorf("invalid field name %s", fieldName)
	}

	// ensure we're not setting multiple values when only one value can be used
	splitValue := strings.Split(value, ",")
	if !multipleFieldsAllowed(fieldName) && len(splitValue) > 1 {
		return HeaderEntry{}, fmt.Errorf("only one value allowed for field %s (%d values given)", fieldName, len(splitValue))
	}

	return HeaderEntry{Key: fieldName, Value: []string{value}}, nil
}

func validField(fieldName string) bool {
	return existsInStringSlice(validFields(), fieldName)
}

func multipleFieldsAllowed(fieldName string) bool {
	return !existsInStringSlice(singleValueOnly(), fieldName)
}

func validFields() []string {
	return []string{
		"Date",
		"From",
		"Sender",
		"Reply-To",
		"To",
		"Cc",
		"Bcc",
		"Message-ID",
		"In-Reply-To",
		"References",
		"Subject",
		"Comments",
		"Keywords",
		"Resent-Date",
		"Resent-From",
		"Resent-Sender",
		"Resent-To",
		"Resent-Cc",
		"Resent-Bcc",
		"Resent-Message-ID",
		"Return-Path",
		"Received",
	}
}

func singleValueOnly() []string {
	return []string{
		"Date",
		"From",
		"Sender",
		"Message-ID",
		"In-Reply-To",
		"References",
		"Subject",
	}
}

func requiredFields() []string {
	return []string{
		"Date",
		"From",
	}
}

func existsInStringSlice(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}
