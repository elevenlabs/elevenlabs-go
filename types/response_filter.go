package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Configuration for filtering tool responses before they are visible to the agent.
type ResponseFilter struct {
	// Controls how tool responses are filtered before being visible to the agent.
	Mode *ResponseFilterMode `json:"mode,omitempty" xml:"mode,omitempty"`
	// Dot notation paths to include when mode is 'allow' (e.g., ['ticket.id', 'ticket.status']).
	Filters []string `json:"filters,omitempty" xml:"filters,omitempty"`
	// Content type for response filtering. Only 'application/json' responses are filtered.
	ContentType     *string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ResponseFilter) GetMode() (value *ResponseFilterMode) {
	if r == nil {
		return
	}
	return r.Mode
}

func (r *ResponseFilter) GetFilters() (value []string) {
	if r == nil {
		return
	}
	return r.Filters
}

func (r *ResponseFilter) GetContentType() (value *string) {
	if r == nil {
		return
	}
	return r.ContentType
}

func (r *ResponseFilter) SetMode(mode *ResponseFilterMode) {
	r.Mode = mode
}

func (r *ResponseFilter) SetFilters(filters []string) {
	r.Filters = filters
}

func (r *ResponseFilter) SetContentType(contentType *string) {
	r.ContentType = contentType
}

func (r *ResponseFilter) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ResponseFilter) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ResponseFilter to string"
	}
	return string(jsonData)
}

func (r *ResponseFilter) UnmarshalJSON(data []byte) error {
	type unmarshaler ResponseFilter
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ResponseFilter(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ResponseFilter) MarshalJSON() ([]byte, error) {
	type embed ResponseFilter
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
