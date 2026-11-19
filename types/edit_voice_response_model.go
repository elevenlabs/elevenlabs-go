package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type EditVoiceResponseModel struct {
	// The status of the voice edit request. If the request was successful, the status will be 'ok'. Otherwise an error message with status 500 will be returned.
	Status          string `json:"status" xml:"status" required:"true"`
	extraProperties map[string]interface{}
}

func (e *EditVoiceResponseModel) GetStatus() (value string) {
	if e == nil {
		return
	}
	return e.Status
}

func (e *EditVoiceResponseModel) SetStatus(status string) {
	e.Status = status
}

func (e *EditVoiceResponseModel) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EditVoiceResponseModel) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EditVoiceResponseModel to string"
	}
	return string(jsonData)
}

func (e *EditVoiceResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, e); err != nil {
		return err
	}
	type alias EditVoiceResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*e = EditVoiceResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EditVoiceResponseModel) MarshalJSON() ([]byte, error) {
	type embed EditVoiceResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
