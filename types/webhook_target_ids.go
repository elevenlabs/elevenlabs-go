package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Deliver the result to specific configured flows webhooks.
type WebhookTargetIds struct {
	// Send the result to the listed flows webhooks.
	type_ string
	// The IDs of the workspace flows webhooks to deliver the result to. Each must be one of the workspace's configured flows webhooks.
	Ids             []string `json:"ids" xml:"ids" required:"true" minItems:"1"`
	extraProperties map[string]interface{}
}

func (w *WebhookTargetIds) GetIds() (value []string) {
	if w == nil {
		return
	}
	return w.Ids
}

func (w *WebhookTargetIds) SetIds(ids []string) {
	w.Ids = ids
}

func (w *WebhookTargetIds) Type() string {
	return w.type_
}

func (w *WebhookTargetIds) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WebhookTargetIds) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WebhookTargetIds to string"
	}
	return string(jsonData)
}

func (w *WebhookTargetIds) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WebhookTargetIds
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WebhookTargetIds(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "ids" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", w, "ids", *literals.Type)
	}
	w.type_ = "ids"
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WebhookTargetIds) MarshalJSON() ([]byte, error) {
	type embed WebhookTargetIds
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*w),
		Type:  "ids",
	}
	return json.Marshal(marshaler)
}
