package gomail

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/mail"
	"strings"
)

type Message struct {
	Header mail.Header
	Body   string
}

// ReadMessage reads the message and returns the Message value
func ReadMessage(src io.Reader) (Message, error) {
	// Read message to parse out Header/Body
	m, err := mail.ReadMessage(src)
	if err != nil {
		return Message{}, err
	}

	body, err := ioutil.ReadAll(m.Body)
	if err != nil {
		return Message{Header: m.Header}, err
	}

	return Message{
		Header: m.Header,
		Body:   string(body),
	}, nil
}

// AppendHeader adds the new entry to the message header
func (m *Message) AppendHeader(entry HeaderEntry) error {
	keyExists := m.keyExists(entry.Key)

	if !multipleFieldsAllowed(entry.Key) && keyExists {
		return fmt.Errorf("key %s already exists in header", entry.Key)
	}

	// multiple fields are allowed so we can append if it exists
	if keyExists {
		s := strings.Join(entry.Value, ", ")
		newValue := []string{fmt.Sprintf("%s, %s", m.Header[entry.Key], s)}
		m.Header[entry.Key] = newValue

		return nil
	}

	m.Header[entry.Key] = entry.Value
	return nil
}

// Join appends the body of the message to the header to display the full
// message
func (m *Message) Join() string {
	var result []string

	for key, value := range m.Header {
		// value is stored as a slice within the first element
		result = append(result, fmt.Sprintf("%s: %s", key, value[0]))
	}

	// Ensure there's a blank line before starting the body of the message
	result = append(result, "")
	result = append(result, m.Body)

	return strings.Join(result, "\r\n")
}

func (m *Message) keyExists(key string) bool {
	return m.Header.Get(key) != ""
}
