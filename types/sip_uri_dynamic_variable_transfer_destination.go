package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SipUriDynamicVariableTransferDestination struct {
	SipUri          string `json:"sip_uri" xml:"sip_uri" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SipUriDynamicVariableTransferDestination) GetSipUri() (value string) {
	if s == nil {
		return
	}
	return s.SipUri
}

func (s *SipUriDynamicVariableTransferDestination) SetSipUri(sipUri string) {
	s.SipUri = sipUri
}

func (s *SipUriDynamicVariableTransferDestination) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SipUriDynamicVariableTransferDestination) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SipUriDynamicVariableTransferDestination to string"
	}
	return string(jsonData)
}

func (s *SipUriDynamicVariableTransferDestination) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SipUriDynamicVariableTransferDestination
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SipUriDynamicVariableTransferDestination(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SipUriDynamicVariableTransferDestination) MarshalJSON() ([]byte, error) {
	type embed SipUriDynamicVariableTransferDestination
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
