package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type EditVoiceSettingsResponseModel struct {
	// The status of the voice settings edit request. If the request was successful, the status will be 'ok'. Otherwise an error message with status 500 will be returned.
	Status          string `json:"status" xml:"status" required:"true"`
	extraProperties map[string]interface{}
}

func (e *EditVoiceSettingsResponseModel) GetStatus() (value string) {
	if e == nil {
		return
	}
	return e.Status
}

func (e *EditVoiceSettingsResponseModel) SetStatus(status string) {
	e.Status = status
}

func (e *EditVoiceSettingsResponseModel) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EditVoiceSettingsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EditVoiceSettingsResponseModel to string"
	}
	return string(jsonData)
}

func (e *EditVoiceSettingsResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, e); err != nil {
		return err
	}
	type alias EditVoiceSettingsResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*e = EditVoiceSettingsResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EditVoiceSettingsResponseModel) MarshalJSON() ([]byte, error) {
	type embed EditVoiceSettingsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
