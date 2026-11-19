package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SipColdConferenceTransferConfigInput struct {
	TransferDestination *SipColdConferenceTransferConfigInputTransferDestination `json:"transfer_destination" xml:"transfer_destination" required:"true"`
	extraProperties     map[string]interface{}
}

func (s *SipColdConferenceTransferConfigInput) GetTransferDestination() (value *SipColdConferenceTransferConfigInputTransferDestination) {
	if s == nil {
		return
	}
	return s.TransferDestination
}

func (s *SipColdConferenceTransferConfigInput) SetTransferDestination(transferDestination *SipColdConferenceTransferConfigInputTransferDestination) {
	s.TransferDestination = transferDestination
}

func (s *SipColdConferenceTransferConfigInput) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SipColdConferenceTransferConfigInput) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SipColdConferenceTransferConfigInput to string"
	}
	return string(jsonData)
}

func (s *SipColdConferenceTransferConfigInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SipColdConferenceTransferConfigInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SipColdConferenceTransferConfigInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SipColdConferenceTransferConfigInput) MarshalJSON() ([]byte, error) {
	type embed SipColdConferenceTransferConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
