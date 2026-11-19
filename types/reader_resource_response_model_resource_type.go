package types

import (
	"fmt"
)

// The type of resource.
type ReaderResourceResponseModelResourceType string

const (
	ReaderResourceResponseModelResourceTypeRead       ReaderResourceResponseModelResourceType = "read"
	ReaderResourceResponseModelResourceTypeCollection ReaderResourceResponseModelResourceType = "collection"
)

func NewReaderResourceResponseModelResourceTypeFromString(s string) (ReaderResourceResponseModelResourceType, error) {
	switch s {
	case "read":
		return ReaderResourceResponseModelResourceTypeRead, nil
	case "collection":
		return ReaderResourceResponseModelResourceTypeCollection, nil
	}
	var t ReaderResourceResponseModelResourceType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ReaderResourceResponseModelResourceType) Ptr() *ReaderResourceResponseModelResourceType {
	return &e
}
