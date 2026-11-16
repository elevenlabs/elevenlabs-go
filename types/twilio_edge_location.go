package types

import (
	"fmt"
)

// Valid Twilio edge locations.
type TwilioEdgeLocation string

const (
	TwilioEdgeLocationAshburn   TwilioEdgeLocation = "ashburn"
	TwilioEdgeLocationDublin    TwilioEdgeLocation = "dublin"
	TwilioEdgeLocationFrankfurt TwilioEdgeLocation = "frankfurt"
	TwilioEdgeLocationSaoPaulo  TwilioEdgeLocation = "sao-paulo"
	TwilioEdgeLocationSingapore TwilioEdgeLocation = "singapore"
	TwilioEdgeLocationSydney    TwilioEdgeLocation = "sydney"
	TwilioEdgeLocationTokyo     TwilioEdgeLocation = "tokyo"
	TwilioEdgeLocationUmatilla  TwilioEdgeLocation = "umatilla"
	TwilioEdgeLocationRoaming   TwilioEdgeLocation = "roaming"
)

func NewTwilioEdgeLocationFromString(s string) (TwilioEdgeLocation, error) {
	switch s {
	case "ashburn":
		return TwilioEdgeLocationAshburn, nil
	case "dublin":
		return TwilioEdgeLocationDublin, nil
	case "frankfurt":
		return TwilioEdgeLocationFrankfurt, nil
	case "sao-paulo":
		return TwilioEdgeLocationSaoPaulo, nil
	case "singapore":
		return TwilioEdgeLocationSingapore, nil
	case "sydney":
		return TwilioEdgeLocationSydney, nil
	case "tokyo":
		return TwilioEdgeLocationTokyo, nil
	case "umatilla":
		return TwilioEdgeLocationUmatilla, nil
	case "roaming":
		return TwilioEdgeLocationRoaming, nil
	}
	var t TwilioEdgeLocation
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TwilioEdgeLocation) Ptr() *TwilioEdgeLocation {
	return &e
}
