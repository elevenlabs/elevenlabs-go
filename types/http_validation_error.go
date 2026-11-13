package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type HttpValidationError struct {
	Detail          []*ValidationError `json:"detail,omitempty" xml:"detail,omitempty"`
	extraProperties map[string]interface{}
}

func (h *HttpValidationError) GetDetail() (value []*ValidationError) {
	if h == nil {
		return
	}
	return h.Detail
}

func (h *HttpValidationError) SetDetail(detail []*ValidationError) {
	h.Detail = detail
}

func (h *HttpValidationError) GetExtraProperties() map[string]interface{} {
	if h == nil {
		return nil
	}
	return h.extraProperties
}

func (h HttpValidationError) String() string {
	jsonData, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return "error converting struct: HttpValidationError to string"
	}
	return string(jsonData)
}

func (h *HttpValidationError) UnmarshalJSON(data []byte) error {
	type unmarshaler HttpValidationError
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*h = HttpValidationError(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *h)
	if err != nil {
		return err
	}
	h.extraProperties = extraProperties
	return nil
}

func (h *HttpValidationError) MarshalJSON() ([]byte, error) {
	type embed HttpValidationError
	var marshaler = struct {
		embed
	}{
		embed: embed(*h),
	}
	return json.Marshal(marshaler)
}
