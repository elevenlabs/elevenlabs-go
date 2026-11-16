package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkspaceWebhookUsageResponseModel struct {
	UsageType       WebhookUsageType `json:"usage_type" xml:"usage_type" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkspaceWebhookUsageResponseModel) GetUsageType() (value WebhookUsageType) {
	if w == nil {
		return
	}
	return w.UsageType
}

func (w *WorkspaceWebhookUsageResponseModel) SetUsageType(usageType WebhookUsageType) {
	w.UsageType = usageType
}

func (w *WorkspaceWebhookUsageResponseModel) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkspaceWebhookUsageResponseModel) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkspaceWebhookUsageResponseModel to string"
	}
	return string(jsonData)
}

func (w *WorkspaceWebhookUsageResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WorkspaceWebhookUsageResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WorkspaceWebhookUsageResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkspaceWebhookUsageResponseModel) MarshalJSON() ([]byte, error) {
	type embed WorkspaceWebhookUsageResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
