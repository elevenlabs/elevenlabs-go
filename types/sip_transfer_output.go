package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SipTransferOutput struct {
	type_           string
	Config          *SipTransferOutputConfig `json:"config" xml:"config" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SipTransferOutput) GetConfig() (value *SipTransferOutputConfig) {
	if s == nil {
		return
	}
	return s.Config
}

func (s *SipTransferOutput) SetConfig(config *SipTransferOutputConfig) {
	s.Config = config
}

func (s *SipTransferOutput) Type() string {
	return s.type_
}

func (s *SipTransferOutput) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SipTransferOutput) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SipTransferOutput to string"
	}
	return string(jsonData)
}

func (s *SipTransferOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SipTransferOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SipTransferOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "sip_trunk" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "sip_trunk", *literals.Type)
	}
	s.type_ = "sip_trunk"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SipTransferOutput) MarshalJSON() ([]byte, error) {
	type embed SipTransferOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*s),
		Type:  "sip_trunk",
	}
	return json.Marshal(marshaler)
}
