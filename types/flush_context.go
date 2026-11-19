package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type FlushContext struct {
	ContextId       string  `json:"context_id" xml:"context_id" required:"true"`
	Text            *string `json:"text,omitempty" xml:"text,omitempty"`
	Flush           bool    `json:"flush" xml:"flush" required:"true"`
	extraProperties map[string]interface{}
}

func (f *FlushContext) GetContextId() (value string) {
	if f == nil {
		return
	}
	return f.ContextId
}

func (f *FlushContext) GetText() (value *string) {
	if f == nil {
		return
	}
	return f.Text
}

func (f *FlushContext) GetFlush() (value bool) {
	if f == nil {
		return
	}
	return f.Flush
}

func (f *FlushContext) SetContextId(contextId string) {
	f.ContextId = contextId
}

func (f *FlushContext) SetText(text *string) {
	f.Text = text
}

func (f *FlushContext) SetFlush(flush bool) {
	f.Flush = flush
}

func (f *FlushContext) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f FlushContext) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FlushContext to string"
	}
	return string(jsonData)
}

func (f *FlushContext) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, f); err != nil {
		return err
	}
	type alias FlushContext
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*f = FlushContext(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *FlushContext) MarshalJSON() ([]byte, error) {
	type embed FlushContext
	var marshaler = struct {
		embed
	}{
		embed: embed(*f),
	}
	return json.Marshal(marshaler)
}
