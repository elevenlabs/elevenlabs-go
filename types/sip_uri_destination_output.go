package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SipUriDestinationOutput struct {
	// Phone number, SIP URI, or the dynamic variable name that resolves to one.
	Destination     string `json:"destination" xml:"destination" required:"true" minLength:"1"`
	type_           string
	extraProperties map[string]interface{}
}

func (s *SipUriDestinationOutput) GetDestination() (value string) {
	if s == nil {
		return
	}
	return s.Destination
}

func (s *SipUriDestinationOutput) SetDestination(destination string) {
	s.Destination = destination
}

func (s *SipUriDestinationOutput) Type() string {
	return s.type_
}

func (s *SipUriDestinationOutput) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SipUriDestinationOutput) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SipUriDestinationOutput to string"
	}
	return string(jsonData)
}

func (s *SipUriDestinationOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SipUriDestinationOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SipUriDestinationOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "sip_uri" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "sip_uri", *literals.Type)
	}
	s.type_ = "sip_uri"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SipUriDestinationOutput) MarshalJSON() ([]byte, error) {
	type embed SipUriDestinationOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*s),
		Type:  "sip_uri",
	}
	return json.Marshal(marshaler)
}
