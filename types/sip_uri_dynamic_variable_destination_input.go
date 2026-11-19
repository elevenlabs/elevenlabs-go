package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SipUriDynamicVariableDestinationInput struct {
	// Phone number, SIP URI, or the dynamic variable name that resolves to one.
	Destination     string `json:"destination" xml:"destination" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (s *SipUriDynamicVariableDestinationInput) GetDestination() (value string) {
	if s == nil {
		return
	}
	return s.Destination
}

func (s *SipUriDynamicVariableDestinationInput) SetDestination(destination string) {
	s.Destination = destination
}

func (s *SipUriDynamicVariableDestinationInput) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SipUriDynamicVariableDestinationInput) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SipUriDynamicVariableDestinationInput to string"
	}
	return string(jsonData)
}

func (s *SipUriDynamicVariableDestinationInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SipUriDynamicVariableDestinationInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SipUriDynamicVariableDestinationInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SipUriDynamicVariableDestinationInput) MarshalJSON() ([]byte, error) {
	type embed SipUriDynamicVariableDestinationInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
