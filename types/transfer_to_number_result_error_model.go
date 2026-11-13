package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TransferToNumberResultErrorModel struct {
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
	Error           string  `json:"error" xml:"error" required:"true"`
	Details         *string `json:"details,omitempty" xml:"details,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TransferToNumberResultErrorModel) GetStatus() (value *string) {
	if t == nil {
		return
	}
	return t.Status
}

func (t *TransferToNumberResultErrorModel) GetError() (value string) {
	if t == nil {
		return
	}
	return t.Error
}

func (t *TransferToNumberResultErrorModel) GetDetails() (value *string) {
	if t == nil {
		return
	}
	return t.Details
}

func (t *TransferToNumberResultErrorModel) SetStatus(status *string) {
	t.Status = status
}

func (t *TransferToNumberResultErrorModel) SetError(error string) {
	t.Error = error
}

func (t *TransferToNumberResultErrorModel) SetDetails(details *string) {
	t.Details = details
}

func (t *TransferToNumberResultErrorModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TransferToNumberResultErrorModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TransferToNumberResultErrorModel to string"
	}
	return string(jsonData)
}

func (t *TransferToNumberResultErrorModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "result_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TransferToNumberResultErrorModel) MarshalJSON() ([]byte, error) {
	type embed TransferToNumberResultErrorModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
