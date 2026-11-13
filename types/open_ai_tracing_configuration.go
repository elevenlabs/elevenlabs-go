package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Only “group_id“ is acted on; “workflow_name“ and “metadata“ are kept so the echoed session matches the request.
type OpenAiTracingConfiguration struct {
	// Adds conversation to the group (case-insensitive, up to 64 chars).
	GroupId         *string        `json:"group_id,omitempty" xml:"group_id,omitempty"`
	WorkflowName    *string        `json:"workflow_name,omitempty" xml:"workflow_name,omitempty"`
	Metadata        map[string]any `json:"metadata,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OpenAiTracingConfiguration) GetGroupId() (value *string) {
	if o == nil {
		return
	}
	return o.GroupId
}

func (o *OpenAiTracingConfiguration) GetWorkflowName() (value *string) {
	if o == nil {
		return
	}
	return o.WorkflowName
}

func (o *OpenAiTracingConfiguration) GetMetadata() (value map[string]any) {
	if o == nil {
		return
	}
	return o.Metadata
}

func (o *OpenAiTracingConfiguration) SetGroupId(groupId *string) {
	o.GroupId = groupId
}

func (o *OpenAiTracingConfiguration) SetWorkflowName(workflowName *string) {
	o.WorkflowName = workflowName
}

func (o *OpenAiTracingConfiguration) SetMetadata(metadata map[string]any) {
	o.Metadata = metadata
}

func (o *OpenAiTracingConfiguration) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OpenAiTracingConfiguration) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OpenAiTracingConfiguration to string"
	}
	return string(jsonData)
}

func (o *OpenAiTracingConfiguration) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, o); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OpenAiTracingConfiguration) MarshalJSON() ([]byte, error) {
	type embed OpenAiTracingConfiguration
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
