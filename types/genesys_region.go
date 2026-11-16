package types

import (
	"fmt"
)

type GenesysRegion string

const (
	GenesysRegionUsEast1      GenesysRegion = "us_east_1"
	GenesysRegionEuWest1      GenesysRegion = "eu_west_1"
	GenesysRegionApSoutheast2 GenesysRegion = "ap_southeast_2"
	GenesysRegionApNortheast1 GenesysRegion = "ap_northeast_1"
	GenesysRegionEuCentral1   GenesysRegion = "eu_central_1"
	GenesysRegionUsWest2      GenesysRegion = "us_west_2"
	GenesysRegionCaCentral1   GenesysRegion = "ca_central_1"
	GenesysRegionApNortheast2 GenesysRegion = "ap_northeast_2"
	GenesysRegionEuWest2      GenesysRegion = "eu_west_2"
	GenesysRegionApSouth1     GenesysRegion = "ap_south_1"
	GenesysRegionUsEast2      GenesysRegion = "us_east_2"
	GenesysRegionSaEast1      GenesysRegion = "sa_east_1"
	GenesysRegionMeCentral1   GenesysRegion = "me_central_1"
	GenesysRegionApNortheast3 GenesysRegion = "ap_northeast_3"
	GenesysRegionEuCentral2   GenesysRegion = "eu_central_2"
	GenesysRegionMxCentral1   GenesysRegion = "mx_central_1"
	GenesysRegionApSoutheast1 GenesysRegion = "ap_southeast_1"
)

func NewGenesysRegionFromString(s string) (GenesysRegion, error) {
	switch s {
	case "us_east_1":
		return GenesysRegionUsEast1, nil
	case "eu_west_1":
		return GenesysRegionEuWest1, nil
	case "ap_southeast_2":
		return GenesysRegionApSoutheast2, nil
	case "ap_northeast_1":
		return GenesysRegionApNortheast1, nil
	case "eu_central_1":
		return GenesysRegionEuCentral1, nil
	case "us_west_2":
		return GenesysRegionUsWest2, nil
	case "ca_central_1":
		return GenesysRegionCaCentral1, nil
	case "ap_northeast_2":
		return GenesysRegionApNortheast2, nil
	case "eu_west_2":
		return GenesysRegionEuWest2, nil
	case "ap_south_1":
		return GenesysRegionApSouth1, nil
	case "us_east_2":
		return GenesysRegionUsEast2, nil
	case "sa_east_1":
		return GenesysRegionSaEast1, nil
	case "me_central_1":
		return GenesysRegionMeCentral1, nil
	case "ap_northeast_3":
		return GenesysRegionApNortheast3, nil
	case "eu_central_2":
		return GenesysRegionEuCentral2, nil
	case "mx_central_1":
		return GenesysRegionMxCentral1, nil
	case "ap_southeast_1":
		return GenesysRegionApSoutheast1, nil
	}
	var t GenesysRegion
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GenesysRegion) Ptr() *GenesysRegion {
	return &e
}
