package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AssetTranscription struct {
	Status          AssetTranscriptionStatus `json:"status" xml:"status" required:"true"`
	Data            *AssetTranscriptionData  `json:"data" xml:"data" required:"true"`
	UpdatedAtMs     *int                     `json:"updated_at_ms,omitempty" xml:"updated_at_ms,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AssetTranscription) GetStatus() (value AssetTranscriptionStatus) {
	if a == nil {
		return
	}
	return a.Status
}

func (a *AssetTranscription) GetData() (value *AssetTranscriptionData) {
	if a == nil {
		return
	}
	return a.Data
}

func (a *AssetTranscription) GetUpdatedAtMs() (value *int) {
	if a == nil {
		return
	}
	return a.UpdatedAtMs
}

func (a *AssetTranscription) SetStatus(status AssetTranscriptionStatus) {
	a.Status = status
}

func (a *AssetTranscription) SetData(data *AssetTranscriptionData) {
	a.Data = data
}

func (a *AssetTranscription) SetUpdatedAtMs(updatedAtMs *int) {
	a.UpdatedAtMs = updatedAtMs
}

func (a *AssetTranscription) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AssetTranscription) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AssetTranscription to string"
	}
	return string(jsonData)
}

func (a *AssetTranscription) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AssetTranscription) MarshalJSON() ([]byte, error) {
	type embed AssetTranscription
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
