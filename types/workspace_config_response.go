package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkspaceConfigResponse struct {
	// Increases with every replacement; send it as `expected_version`.
	Version         int                   `json:"version" xml:"version" required:"true"`
	Config          *CampaignConfigOutput `json:"config" xml:"config" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkspaceConfigResponse) GetVersion() (value int) {
	if w == nil {
		return
	}
	return w.Version
}

func (w *WorkspaceConfigResponse) GetConfig() (value *CampaignConfigOutput) {
	if w == nil {
		return
	}
	return w.Config
}

func (w *WorkspaceConfigResponse) SetVersion(version int) {
	w.Version = version
}

func (w *WorkspaceConfigResponse) SetConfig(config *CampaignConfigOutput) {
	w.Config = config
}

func (w *WorkspaceConfigResponse) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkspaceConfigResponse) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkspaceConfigResponse to string"
	}
	return string(jsonData)
}

func (w *WorkspaceConfigResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WorkspaceConfigResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WorkspaceConfigResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkspaceConfigResponse) MarshalJSON() ([]byte, error) {
	type embed WorkspaceConfigResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
