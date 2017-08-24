package models

import (
	"errors"
	"regexp"

	"gopkg.in/mgo.v2/bson"
)

// Email helpers

type Email string

var (
	ErrEmptyEmail       = errors.New("empty email")
	ErrWrongEmailFormat = errors.New("wrong email format")
)

//Long and strange regexp to validate email format.
var emailRegex = regexp.MustCompile(`^(([^<>()\[\]\\.,;:\s@“]+(\.[^<>()\[\]\\.,;:\s@“]+)*)|(“.+“))@((\[[0-9]{1,3}\.
[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)

// Validates email
func (e Email) Validate() error {
	if len(e) == 0 {
		return ErrEmptyEmail
	}
	if !emailRegex.MatchString(string(e)) {
		return ErrWrongEmailFormat
	}
	return nil
}

// Password helpers

type Password string

var (
	ErrEmptyPassword       = errors.New("empty password")
	ErrWrongPasswordFormat = errors.New("wrong password format")
)

var passwordRegex = regexp.MustCompile(`^[0-9a-zA-Z\s\r\n@!#$^%&*()+=\-\[\]\\';,./{}|":<>?]{3,14}$`)

// Validates passowrd
func (p Password) Validate() error {
	if len(p) == 0 {
		return ErrEmptyPassword
	}
	if !passwordRegex.MatchString(string(p)) {
		return ErrWrongPasswordFormat
	}
	return nil
}

// Name

type Name string

var nameRegex = regexp.MustCompile(`^[a-zA-Z].{1,49}$`)

var (
	ErrEmptyName       = errors.New("empty name")
	ErrWrongNameFormat = errors.New("wrong name format")
)

// Validates names
func (n Name) Validate() error {
	if len(n) == 0 {
		return ErrEmptyName
	}
	if !nameRegex.MatchString(string(n)) {
		return ErrWrongNameFormat
	}
	return nil
}

// Text

type Text string

var (
	ErrTextTooLong = errors.New("text too long")

	MaxTextLen = 500
)

// Validate text
func (t Text) Validate() error {
	if len(t) > MaxTextLen {
		return ErrTextTooLong
	}
	return nil
}

// General Id helpers

type Id bson.ObjectId

var ErrInvalidId = errors.New("invalid id")

// Validates id
func ValidateId(id Id) error {
	// NOTE: By  default id.Valid() checks only id len
	// BTW we could pass id like: bson.ObjectId("12_bytes_len")
	if !bson.IsObjectIdHex(string(id)) {
		return ErrInvalidId
	}
	return nil
}

// AutoId helpers

type AutoId Id

var ErrIdMustBeOmitted = errors.New("id must be omitted")

// Validates generated id
func (id AutoId) Validate() error {
	if id != AutoId("") {
		return ErrIdMustBeOmitted
	}
	return nil
}

// RequiredId helpers

type RequiredId Id

var ErrIdMustBePresent = errors.New("id must be present")

// Validates required id
func (id RequiredId) Validate() error {
	if id == RequiredId("") {
		return ErrIdMustBePresent
	}

	if err := ValidateId(Id(id)); err != nil {
		return err
	}
	return nil
}

// Optional Id helpers

type OptionalId bson.ObjectId

// Validates optional id
func (id OptionalId) Validate() error {
	if err := AutoId(id).Validate(); err == nil {
		return nil
	}

	if err := RequiredId(id).Validate(); err != nil {
		return err
	}
	return nil
}