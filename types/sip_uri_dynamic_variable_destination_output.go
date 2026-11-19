package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SipUriDynamicVariableDestinationOutput struct {
	// Phone number, SIP URI, or the dynamic variable name that resolves to one.
	Destination     string `json:"destination" xml:"destination" required:"true" minLength:"1"`
	type_           string
	extraProperties map[string]interface{}
}

func (s *SipUriDynamicVariableDestinationOutput) GetDestination() (value string) {
	if s == nil {
		return
	}
	return s.Destination
}

func (s *SipUriDynamicVariableDestinationOutput) SetDestination(destination string) {
	s.Destination = destination
}

func (s *SipUriDynamicVariableDestinationOutput) Type() string {
	return s.type_
}

func (s *SipUriDynamicVariableDestinationOutput) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SipUriDynamicVariableDestinationOutput) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SipUriDynamicVariableDestinationOutput to string"
	}
	return string(jsonData)
}

func (s *SipUriDynamicVariableDestinationOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SipUriDynamicVariableDestinationOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SipUriDynamicVariableDestinationOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "sip_uri_dynamic_variable" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "sip_uri_dynamic_variable", *literals.Type)
	}
	s.type_ = "sip_uri_dynamic_variable"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SipUriDynamicVariableDestinationOutput) MarshalJSON() ([]byte, error) {
	type embed SipUriDynamicVariableDestinationOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*s),
		Type:  "sip_uri_dynamic_variable",
	}
	return json.Marshal(marshaler)
}
