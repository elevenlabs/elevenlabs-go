package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingErrorResponse struct {
	// Stable identifier for the failure, safe to branch on. New codes are added over time, so we recommend treating an unrecognized value as `internal_error`.
	Code string `json:"code" xml:"code" required:"true"`
	// Human-readable description of the failure, for display. The wording may change at any time, so we recommend branching on `code` instead.
	Message string `json:"message" xml:"message" required:"true"`
	// Whether resubmitting the same input could succeed. A value of false means the failure describes the input or the account, so an identical retry will fail the same way.
	Retryable       bool `json:"retryable" xml:"retryable" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DubbingErrorResponse) GetCode() (value string) {
	if d == nil {
		return
	}
	return d.Code
}

func (d *DubbingErrorResponse) GetMessage() (value string) {
	if d == nil {
		return
	}
	return d.Message
}

func (d *DubbingErrorResponse) GetRetryable() (value bool) {
	if d == nil {
		return
	}
	return d.Retryable
}

func (d *DubbingErrorResponse) SetCode(code string) {
	d.Code = code
}

func (d *DubbingErrorResponse) SetMessage(message string) {
	d.Message = message
}

func (d *DubbingErrorResponse) SetRetryable(retryable bool) {
	d.Retryable = retryable
}

func (d *DubbingErrorResponse) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingErrorResponse) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingErrorResponse to string"
	}
	return string(jsonData)
}

func (d *DubbingErrorResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingErrorResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingErrorResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingErrorResponse) MarshalJSON() ([]byte, error) {
	type embed DubbingErrorResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
