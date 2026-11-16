package types

import (
	"fmt"
)

// Predefined background sound preset identifiers.
type BackgroundSoundPresetId string

const (
	BackgroundSoundPresetIdOffice2    BackgroundSoundPresetId = "office2"
	BackgroundSoundPresetIdOffice1    BackgroundSoundPresetId = "office1"
	BackgroundSoundPresetIdRestaurant BackgroundSoundPresetId = "restaurant"
	BackgroundSoundPresetIdCity       BackgroundSoundPresetId = "city"
	BackgroundSoundPresetIdTyping     BackgroundSoundPresetId = "typing"
	BackgroundSoundPresetIdElevator1  BackgroundSoundPresetId = "elevator1"
	BackgroundSoundPresetIdElevator2  BackgroundSoundPresetId = "elevator2"
	BackgroundSoundPresetIdElevator3  BackgroundSoundPresetId = "elevator3"
	BackgroundSoundPresetIdElevator4  BackgroundSoundPresetId = "elevator4"
)

func NewBackgroundSoundPresetIdFromString(s string) (BackgroundSoundPresetId, error) {
	switch s {
	case "office2":
		return BackgroundSoundPresetIdOffice2, nil
	case "office1":
		return BackgroundSoundPresetIdOffice1, nil
	case "restaurant":
		return BackgroundSoundPresetIdRestaurant, nil
	case "city":
		return BackgroundSoundPresetIdCity, nil
	case "typing":
		return BackgroundSoundPresetIdTyping, nil
	case "elevator1":
		return BackgroundSoundPresetIdElevator1, nil
	case "elevator2":
		return BackgroundSoundPresetIdElevator2, nil
	case "elevator3":
		return BackgroundSoundPresetIdElevator3, nil
	case "elevator4":
		return BackgroundSoundPresetIdElevator4, nil
	}
	var t BackgroundSoundPresetId
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BackgroundSoundPresetId) Ptr() *BackgroundSoundPresetId {
	return &e
}
