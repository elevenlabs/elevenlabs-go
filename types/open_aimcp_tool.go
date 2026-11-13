package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type OpenAimcpTool struct {
	ServerLabel     string                        `json:"server_label" xml:"server_label" required:"true"`
	ServerUrl       string                        `json:"server_url" xml:"server_url" required:"true"`
	Authorization   *string                       `json:"authorization,omitempty" xml:"authorization,omitempty"`
	AllowedTools    []string                      `json:"allowed_tools,omitempty" xml:"allowed_tools,omitempty"`
	RequireApproval *OpenAimcpToolRequireApproval `json:"require_approval,omitempty" xml:"require_approval,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OpenAimcpTool) GetServerLabel() (value string) {
	if o == nil {
		return
	}
	return o.ServerLabel
}

func (o *OpenAimcpTool) GetServerUrl() (value string) {
	if o == nil {
		return
	}
	return o.ServerUrl
}

func (o *OpenAimcpTool) GetAuthorization() (value *string) {
	if o == nil {
		return
	}
	return o.Authorization
}

func (o *OpenAimcpTool) GetAllowedTools() (value []string) {
	if o == nil {
		return
	}
	return o.AllowedTools
}

func (o *OpenAimcpTool) GetRequireApproval() (value *OpenAimcpToolRequireApproval) {
	if o == nil {
		return
	}
	return o.RequireApproval
}

func (o *OpenAimcpTool) SetServerLabel(serverLabel string) {
	o.ServerLabel = serverLabel
}

func (o *OpenAimcpTool) SetServerUrl(serverUrl string) {
	o.ServerUrl = serverUrl
}

func (o *OpenAimcpTool) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

func (o *OpenAimcpTool) SetAllowedTools(allowedTools []string) {
	o.AllowedTools = allowedTools
}

func (o *OpenAimcpTool) SetRequireApproval(requireApproval *OpenAimcpToolRequireApproval) {
	o.RequireApproval = requireApproval
}

func (o *OpenAimcpTool) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OpenAimcpTool) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OpenAimcpTool to string"
	}
	return string(jsonData)
}

func (o *OpenAimcpTool) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, o); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *o, "type")
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OpenAimcpTool) MarshalJSON() ([]byte, error) {
	type embed OpenAimcpTool
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
