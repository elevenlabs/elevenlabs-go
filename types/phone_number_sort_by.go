package types

import (
	"fmt"
)

type PhoneNumberSortBy string

const (
	PhoneNumberSortByLabel       PhoneNumberSortBy = "label"
	PhoneNumberSortByPhoneNumber PhoneNumberSortBy = "phone_number"
)

func NewPhoneNumberSortByFromString(s string) (PhoneNumberSortBy, error) {
	switch s {
	case "label":
		return PhoneNumberSortByLabel, nil
	case "phone_number":
		return PhoneNumberSortByPhoneNumber, nil
	}
	var t PhoneNumberSortBy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e PhoneNumberSortBy) Ptr() *PhoneNumberSortBy {
	return &e
}
