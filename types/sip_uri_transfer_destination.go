package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SipUriTransferDestination struct {
	SipUri          string `json:"sip_uri" xml:"sip_uri" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SipUriTransferDestination) GetSipUri() (value string) {
	if s == nil {
		return
	}
	return s.SipUri
}

func (s *SipUriTransferDestination) SetSipUri(sipUri string) {
	s.SipUri = sipUri
}

func (s *SipUriTransferDestination) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SipUriTransferDestination) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SipUriTransferDestination to string"
	}
	return string(jsonData)
}

func (s *SipUriTransferDestination) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SipUriTransferDestination
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SipUriTransferDestination(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SipUriTransferDestination) MarshalJSON() ([]byte, error) {
	type embed SipUriTransferDestination
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
