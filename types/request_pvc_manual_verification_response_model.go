package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type RequestPvcManualVerificationResponseModel struct {
	// The status of the request PVC manual verification request. If the request was successful, the status will be 'ok'. Otherwise an error message with status 500 will be returned.
	Status          string `json:"status" xml:"status" required:"true"`
	extraProperties map[string]interface{}
}

func (r *RequestPvcManualVerificationResponseModel) GetStatus() (value string) {
	if r == nil {
		return
	}
	return r.Status
}

func (r *RequestPvcManualVerificationResponseModel) SetStatus(status string) {
	r.Status = status
}

func (r *RequestPvcManualVerificationResponseModel) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RequestPvcManualVerificationResponseModel) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RequestPvcManualVerificationResponseModel to string"
	}
	return string(jsonData)
}

func (r *RequestPvcManualVerificationResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias RequestPvcManualVerificationResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = RequestPvcManualVerificationResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RequestPvcManualVerificationResponseModel) MarshalJSON() ([]byte, error) {
	type embed RequestPvcManualVerificationResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
