package outboundcampaigns

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

type ReplaceWorkspaceConfigRequest struct {
	// The `version` of the configuration being replaced. Omit it only to create the first configuration.
	ExpectedVersion *int                       `json:"expected_version,omitempty" xml:"expected_version,omitempty" min:"1"`
	Config          *types.CampaignConfigInput `json:"config" xml:"config" required:"true"`
}

func (r *ReplaceWorkspaceConfigRequest) GetExpectedVersion() (value *int) {
	if r == nil {
		return
	}
	return r.ExpectedVersion
}

func (r *ReplaceWorkspaceConfigRequest) GetConfig() (value *types.CampaignConfigInput) {
	if r == nil {
		return
	}
	return r.Config
}

func (r *ReplaceWorkspaceConfigRequest) SetExpectedVersion(expectedVersion *int) {
	r.ExpectedVersion = expectedVersion
}

func (r *ReplaceWorkspaceConfigRequest) SetConfig(config *types.CampaignConfigInput) {
	r.Config = config
}

func (r *ReplaceWorkspaceConfigRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReplaceWorkspaceConfigRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReplaceWorkspaceConfigRequest to string"
	}
	return string(jsonData)
}

func (r *ReplaceWorkspaceConfigRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}

func (r *ReplaceWorkspaceConfigRequest) MarshalJSON() ([]byte, error) {
	type embed ReplaceWorkspaceConfigRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
