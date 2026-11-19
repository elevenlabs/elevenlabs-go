package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SipUriDestinationInput struct {
	// Phone number, SIP URI, or the dynamic variable name that resolves to one.
	Destination     string `json:"destination" xml:"destination" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (s *SipUriDestinationInput) GetDestination() (value string) {
	if s == nil {
		return
	}
	return s.Destination
}

func (s *SipUriDestinationInput) SetDestination(destination string) {
	s.Destination = destination
}

func (s *SipUriDestinationInput) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SipUriDestinationInput) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SipUriDestinationInput to string"
	}
	return string(jsonData)
}

func (s *SipUriDestinationInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SipUriDestinationInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SipUriDestinationInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SipUriDestinationInput) MarshalJSON() ([]byte, error) {
	type embed SipUriDestinationInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
