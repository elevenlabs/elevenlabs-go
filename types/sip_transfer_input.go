package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SipTransferInput struct {
	Config          *SipTransferInputConfig `json:"config" xml:"config" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SipTransferInput) GetConfig() (value *SipTransferInputConfig) {
	if s == nil {
		return
	}
	return s.Config
}

func (s *SipTransferInput) SetConfig(config *SipTransferInputConfig) {
	s.Config = config
}

func (s *SipTransferInput) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SipTransferInput) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SipTransferInput to string"
	}
	return string(jsonData)
}

func (s *SipTransferInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SipTransferInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SipTransferInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SipTransferInput) MarshalJSON() ([]byte, error) {
	type embed SipTransferInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
