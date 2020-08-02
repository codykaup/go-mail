package gomail

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewHeaderEntry(t *testing.T) {
	r := require.New(t)

	expected := HeaderEntry{
		Key:   "From",
		Value: []string{"jane@example.com"},
	}

	actual, _ := NewHeaderEntry("From", "jane@example.com")

	r.Equal(expected, actual)
}

func TestNewHeaderEntryInvalidField(t *testing.T) {
	r := require.New(t)

	_, err := NewHeaderEntry("Asdf", "")

	r.EqualError(err, "invalid field name Asdf")
}

func TestNewHeaderEntryMultipleFieldError(t *testing.T) {
	r := require.New(t)

	_, err := NewHeaderEntry("From", "john@example.com, jane@example.com")

	r.EqualError(err, "only one value allowed for field From (2 values given)")
}

func TestValidField(t *testing.T) {
	r := require.New(t)

	r.True(validField("From"))
	r.False(validField("Asdf"))
}

func TestMultipleFieldsAllowed(t *testing.T) {
	r := require.New(t)

	r.True(multipleFieldsAllowed("To"))
	r.False(multipleFieldsAllowed("From"))
}
