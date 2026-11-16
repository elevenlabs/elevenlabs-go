package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PatchWorkspaceWebhookResponseModel struct {
	// The status of the workspace webhook patch request. If the request was successful, the status will be 'ok'. Otherwise an error message with status 500 will be returned.
	Status          string `json:"status" xml:"status" required:"true"`
	extraProperties map[string]interface{}
}

func (p *PatchWorkspaceWebhookResponseModel) GetStatus() (value string) {
	if p == nil {
		return
	}
	return p.Status
}

func (p *PatchWorkspaceWebhookResponseModel) SetStatus(status string) {
	p.Status = status
}

func (p *PatchWorkspaceWebhookResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PatchWorkspaceWebhookResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PatchWorkspaceWebhookResponseModel to string"
	}
	return string(jsonData)
}

func (p *PatchWorkspaceWebhookResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PatchWorkspaceWebhookResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PatchWorkspaceWebhookResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PatchWorkspaceWebhookResponseModel) MarshalJSON() ([]byte, error) {
	type embed PatchWorkspaceWebhookResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
