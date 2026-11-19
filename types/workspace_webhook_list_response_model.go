package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type WorkspaceWebhookListResponseModel struct {
	// List of webhooks currently configured for the workspace
	Webhooks        []*WorkspaceWebhookResponseModel `json:"webhooks" xml:"webhooks" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkspaceWebhookListResponseModel) GetWebhooks() (value []*WorkspaceWebhookResponseModel) {
	if w == nil {
		return
	}
	return w.Webhooks
}

func (w *WorkspaceWebhookListResponseModel) SetWebhooks(webhooks []*WorkspaceWebhookResponseModel) {
	w.Webhooks = webhooks
}

func (w *WorkspaceWebhookListResponseModel) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkspaceWebhookListResponseModel) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkspaceWebhookListResponseModel to string"
	}
	return string(jsonData)
}

func (w *WorkspaceWebhookListResponseModel) UnmarshalJSON(data []byte) error {
	type unmarshaler WorkspaceWebhookListResponseModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WorkspaceWebhookListResponseModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkspaceWebhookListResponseModel) MarshalJSON() ([]byte, error) {
	type embed WorkspaceWebhookListResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
