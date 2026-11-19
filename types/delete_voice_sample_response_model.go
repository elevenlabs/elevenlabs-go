package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DeleteVoiceSampleResponseModel struct {
	// The status of the voice sample deletion request. If the request was successful, the status will be 'ok'. Otherwise an error message with status 500 will be returned.
	Status          string `json:"status" xml:"status" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DeleteVoiceSampleResponseModel) GetStatus() (value string) {
	if d == nil {
		return
	}
	return d.Status
}

func (d *DeleteVoiceSampleResponseModel) SetStatus(status string) {
	d.Status = status
}

func (d *DeleteVoiceSampleResponseModel) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DeleteVoiceSampleResponseModel) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DeleteVoiceSampleResponseModel to string"
	}
	return string(jsonData)
}

func (d *DeleteVoiceSampleResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DeleteVoiceSampleResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DeleteVoiceSampleResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DeleteVoiceSampleResponseModel) MarshalJSON() ([]byte, error) {
	type embed DeleteVoiceSampleResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
