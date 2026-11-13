package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SkipTurnToolResponseModel struct {
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
	Reason          *string `json:"reason,omitempty" xml:"reason,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SkipTurnToolResponseModel) GetStatus() (value *string) {
	if s == nil {
		return
	}
	return s.Status
}

func (s *SkipTurnToolResponseModel) GetReason() (value *string) {
	if s == nil {
		return
	}
	return s.Reason
}

func (s *SkipTurnToolResponseModel) SetStatus(status *string) {
	s.Status = status
}

func (s *SkipTurnToolResponseModel) SetReason(reason *string) {
	s.Reason = reason
}

func (s *SkipTurnToolResponseModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SkipTurnToolResponseModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SkipTurnToolResponseModel to string"
	}
	return string(jsonData)
}

func (s *SkipTurnToolResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "result_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SkipTurnToolResponseModel) MarshalJSON() ([]byte, error) {
	type embed SkipTurnToolResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
