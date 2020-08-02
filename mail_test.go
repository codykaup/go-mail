package gomail

import (
	"net/mail"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadMessage(t *testing.T) {
	r := require.New(t)

	expected := Message{
		Header: mail.Header{
			"Subject": {"Test Email"},
			"To":      {"john@example.com"},
			"From":    {"jane@example.com"},
		},
		Body: "This is a test email.",
	}

	email := "Subject: Test Email\r\nTo: john@example.com\r\nFrom: jane@example.com\r\n\r\nThis is a test email."
	actual, _ := ReadMessage(strings.NewReader(email))

	r.Equal(expected.Header, actual.Header)
	r.Equal(expected.Body, actual.Body)
}

func TestAppendHeader(t *testing.T) {
	r := require.New(t)

	expected := Message{
		Header: mail.Header{
			"Subject": {"Test Email"},
			"To":      {"john@example.com"},
			"From":    {"jane@example.com"},
		},
		Body: "This is a test email.",
	}

	actual := Message{
		Header: mail.Header{
			"Subject": {"Test Email"},
			"To":      {"john@example.com"},
		},
		Body: "This is a test email.",
	}
	actual.AppendHeader(HeaderEntry{
		Key:   "From",
		Value: []string{"jane@example.com"},
	})

	r.Equal(expected.Header, actual.Header)
	r.Equal(expected.Body, actual.Body)
}

func TestJoin(t *testing.T) {
	r := require.New(t)

	email := "Subject: Test Email\r\nTo: john@example.com\r\nFrom: jane@example.com\r\n\r\nThis is a test email."
	expected := strings.Split(email, "\r\n")

	actualStruct := Message{
		Header: mail.Header{
			"Subject": {"Test Email"},
			"To":      {"john@example.com"},
			"From":    {"jane@example.com"},
		},
		Body: "This is a test email.",
	}
	actual := strings.Split(actualStruct.Join(), "\r\n")

	r.ElementsMatch(expected, actual)
}
