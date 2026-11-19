package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DeleteHistoryItemResponse struct {
	// The status of the deletion request. If the request was successful, the status will be 'ok'. Otherwise an error message with http code 500 will be returned.
	Status          string `json:"status" xml:"status" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DeleteHistoryItemResponse) GetStatus() (value string) {
	if d == nil {
		return
	}
	return d.Status
}

func (d *DeleteHistoryItemResponse) SetStatus(status string) {
	d.Status = status
}

func (d *DeleteHistoryItemResponse) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DeleteHistoryItemResponse) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DeleteHistoryItemResponse to string"
	}
	return string(jsonData)
}

func (d *DeleteHistoryItemResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DeleteHistoryItemResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DeleteHistoryItemResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DeleteHistoryItemResponse) MarshalJSON() ([]byte, error) {
	type embed DeleteHistoryItemResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
