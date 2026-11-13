package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Deliver the result to all of the workspace's configured flows webhooks.
type WebhookTargetAll struct {
	// Send the result to all of the workspace's configured flows webhooks.
	type_           string
	extraProperties map[string]interface{}
}

func (w *WebhookTargetAll) Type() string {
	return w.type_
}

func (w *WebhookTargetAll) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WebhookTargetAll) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WebhookTargetAll to string"
	}
	return string(jsonData)
}

func (w *WebhookTargetAll) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WebhookTargetAll
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WebhookTargetAll(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "all" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", w, "all", *literals.Type)
	}
	w.type_ = "all"
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WebhookTargetAll) MarshalJSON() ([]byte, error) {
	type embed WebhookTargetAll
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*w),
		Type:  "all",
	}
	return json.Marshal(marshaler)
}
