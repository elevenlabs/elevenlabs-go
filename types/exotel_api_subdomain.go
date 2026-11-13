package types

import (
	"fmt"
)

type ExotelApiSubdomain string

const (
	ExotelApiSubdomainApiInExotelCom ExotelApiSubdomain = "api.in.exotel.com"
	ExotelApiSubdomainApiExotelCom   ExotelApiSubdomain = "api.exotel.com"
)

func NewExotelApiSubdomainFromString(s string) (ExotelApiSubdomain, error) {
	switch s {
	case "api.in.exotel.com":
		return ExotelApiSubdomainApiInExotelCom, nil
	case "api.exotel.com":
		return ExotelApiSubdomainApiExotelCom, nil
	}
	var t ExotelApiSubdomain
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ExotelApiSubdomain) Ptr() *ExotelApiSubdomain {
	return &e
}
