package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SendTextMulti struct {
	Text            string  `json:"text" xml:"text" required:"true"`
	ContextId       *string `json:"context_id,omitempty" xml:"context_id,omitempty"`
	Flush           *bool   `json:"flush,omitempty" xml:"flush,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SendTextMulti) GetText() (value string) {
	if s == nil {
		return
	}
	return s.Text
}

func (s *SendTextMulti) GetContextId() (value *string) {
	if s == nil {
		return
	}
	return s.ContextId
}

func (s *SendTextMulti) GetFlush() (value *bool) {
	if s == nil {
		return
	}
	return s.Flush
}

func (s *SendTextMulti) SetText(text string) {
	s.Text = text
}

func (s *SendTextMulti) SetContextId(contextId *string) {
	s.ContextId = contextId
}

func (s *SendTextMulti) SetFlush(flush *bool) {
	s.Flush = flush
}

func (s *SendTextMulti) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SendTextMulti) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SendTextMulti to string"
	}
	return string(jsonData)
}

func (s *SendTextMulti) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SendTextMulti
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SendTextMulti(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SendTextMulti) MarshalJSON() ([]byte, error) {
	type embed SendTextMulti
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
