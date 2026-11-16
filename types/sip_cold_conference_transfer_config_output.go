package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SipColdConferenceTransferConfigOutput struct {
	TransferDestination *SipColdConferenceTransferConfigOutputTransferDestination `json:"transfer_destination" xml:"transfer_destination" required:"true"`
	type_               string
	extraProperties     map[string]interface{}
}

func (s *SipColdConferenceTransferConfigOutput) GetTransferDestination() (value *SipColdConferenceTransferConfigOutputTransferDestination) {
	if s == nil {
		return
	}
	return s.TransferDestination
}

func (s *SipColdConferenceTransferConfigOutput) SetTransferDestination(transferDestination *SipColdConferenceTransferConfigOutputTransferDestination) {
	s.TransferDestination = transferDestination
}

func (s *SipColdConferenceTransferConfigOutput) Type() string {
	return s.type_
}

func (s *SipColdConferenceTransferConfigOutput) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SipColdConferenceTransferConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SipColdConferenceTransferConfigOutput to string"
	}
	return string(jsonData)
}

func (s *SipColdConferenceTransferConfigOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SipColdConferenceTransferConfigOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SipColdConferenceTransferConfigOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "sip_cold_conference_transfer" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "sip_cold_conference_transfer", *literals.Type)
	}
	s.type_ = "sip_cold_conference_transfer"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SipColdConferenceTransferConfigOutput) MarshalJSON() ([]byte, error) {
	type embed SipColdConferenceTransferConfigOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*s),
		Type:  "sip_cold_conference_transfer",
	}
	return json.Marshal(marshaler)
}
