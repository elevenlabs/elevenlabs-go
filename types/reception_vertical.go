package types

import (
	"fmt"
)

type ReceptionVertical string

const (
	ReceptionVerticalSalonBarbershop ReceptionVertical = "salon_barbershop"
	ReceptionVerticalNailBeauty      ReceptionVertical = "nail_beauty"
	ReceptionVerticalSpa             ReceptionVertical = "spa"
	ReceptionVerticalMedical         ReceptionVertical = "medical"
	ReceptionVerticalVeterinary      ReceptionVertical = "veterinary"
	ReceptionVerticalFitness         ReceptionVertical = "fitness"
	ReceptionVerticalHomeServices    ReceptionVertical = "home_services"
	ReceptionVerticalPhotography     ReceptionVertical = "photography"
	ReceptionVerticalConsulting      ReceptionVertical = "consulting"
	ReceptionVerticalRestaurant      ReceptionVertical = "restaurant"
	ReceptionVerticalDental          ReceptionVertical = "dental"
	ReceptionVerticalAutoService     ReceptionVertical = "auto_service"
	ReceptionVerticalTutoring        ReceptionVertical = "tutoring"
	ReceptionVerticalLegal           ReceptionVertical = "legal"
	ReceptionVerticalRealEstate      ReceptionVertical = "real_estate"
	ReceptionVerticalOther           ReceptionVertical = "other"
)

func NewReceptionVerticalFromString(s string) (ReceptionVertical, error) {
	switch s {
	case "salon_barbershop":
		return ReceptionVerticalSalonBarbershop, nil
	case "nail_beauty":
		return ReceptionVerticalNailBeauty, nil
	case "spa":
		return ReceptionVerticalSpa, nil
	case "medical":
		return ReceptionVerticalMedical, nil
	case "veterinary":
		return ReceptionVerticalVeterinary, nil
	case "fitness":
		return ReceptionVerticalFitness, nil
	case "home_services":
		return ReceptionVerticalHomeServices, nil
	case "photography":
		return ReceptionVerticalPhotography, nil
	case "consulting":
		return ReceptionVerticalConsulting, nil
	case "restaurant":
		return ReceptionVerticalRestaurant, nil
	case "dental":
		return ReceptionVerticalDental, nil
	case "auto_service":
		return ReceptionVerticalAutoService, nil
	case "tutoring":
		return ReceptionVerticalTutoring, nil
	case "legal":
		return ReceptionVerticalLegal, nil
	case "real_estate":
		return ReceptionVerticalRealEstate, nil
	case "other":
		return ReceptionVerticalOther, nil
	}
	var t ReceptionVertical
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ReceptionVertical) Ptr() *ReceptionVertical {
	return &e
}
