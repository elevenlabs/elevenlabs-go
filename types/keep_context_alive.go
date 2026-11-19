package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type KeepContextAlive struct {
	// An empty string. Ignored by the server but resets the inactivity timeout for the context.
	text            string
	ContextId       string `json:"context_id" xml:"context_id" required:"true"`
	extraProperties map[string]interface{}
}

func (k *KeepContextAlive) GetContextId() (value string) {
	if k == nil {
		return
	}
	return k.ContextId
}

func (k *KeepContextAlive) SetContextId(contextId string) {
	k.ContextId = contextId
}

func (k *KeepContextAlive) Text() string {
	return k.text
}

func (k *KeepContextAlive) GetExtraProperties() map[string]interface{} {
	if k == nil {
		return nil
	}
	return k.extraProperties
}

func (k KeepContextAlive) String() string {
	jsonData, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return "error converting struct: KeepContextAlive to string"
	}
	return string(jsonData)
}

func (k *KeepContextAlive) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, k); err != nil {
		return err
	}
	type alias KeepContextAlive
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*k = KeepContextAlive(tmp)
	var literals struct {
		Text *string `json:"text"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Text != nil && *literals.Text != "" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", k, "", *literals.Text)
	}
	k.text = ""
	extraProperties, err := internal.ExtractExtraProperties(data, *k, "text")
	if err != nil {
		return err
	}
	k.extraProperties = extraProperties
	return nil
}

func (k *KeepContextAlive) MarshalJSON() ([]byte, error) {
	type embed KeepContextAlive
	var marshaler = struct {
		embed
		Text string `json:"text"`
	}{
		embed: embed(*k),
		Text:  "",
	}
	return json.Marshal(marshaler)
}
