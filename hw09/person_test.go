package hw09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerialization(t *testing.T) {
	tests := map[string]struct {
		person Person
		result string
	}{
		"test case with empty fields": {
			result: "name=\nage=0\nmarried=false",
		},
		"test case with fields": {
			person: Person{
				Name:    "John Doe",
				Age:     30,
				Married: true,
			},
			result: "name=John Doe\nage=30\nmarried=true",
		},
		"test case with omitempty field": {
			person: Person{
				Name:    "John Doe",
				Age:     30,
				Married: true,
				Address: "Paris",
			},
			result: "name=John Doe\naddress=Paris\nage=30\nmarried=true",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := Serialize(test.person)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestSerialization_User(t *testing.T) {
	tests := map[string]struct {
		user   User
		result string
	}{
		"test case with empty fields": {
			result: "id=0\nlogin=\nactive=false\nscore=0",
		},
		"test case with fields": {
			user: User{
				Id:     123,
				Login:  "JohnDoe",
				Active: true,
				Score:  100.5,
			},
			result: "id=123\nlogin=JohnDoe\nactive=true\nscore=100.5",
		},
		"test case with omitempty field": {
			user: User{
				Id:     123456,
				Login:  "JaneDoe",
				Name:   "Jane",
				Email:  "jane@doe.com",
				Age:    28,
				Active: true,
				Score:  95.0,
			},
			result: "id=123456\nlogin=JaneDoe\nname=Jane\nemail=jane@doe.com\nage=28\nactive=true\nscore=95",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := Serialize(test.user)
			assert.Equal(t, test.result, result)
		})
	}
}
