package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowToolLocator struct {
	ToolId string `json:"tool_id" xml:"tool_id" required:"true"`
	// Per-node parameter overrides applied on top of the tool's own configuration. Keys are dotted parameter paths (webhook tools prefix keys with path_params./query_params./request_body.). These take precedence over any overrides already defined on the tool itself.
	SchemaOverrides map[string]*WorkflowToolLocatorSchemaOverridesValue `json:"schema_overrides,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WorkflowToolLocator) GetToolId() (value string) {
	if w == nil {
		return
	}
	return w.ToolId
}

func (w *WorkflowToolLocator) GetSchemaOverrides() (value map[string]*WorkflowToolLocatorSchemaOverridesValue) {
	if w == nil {
		return
	}
	return w.SchemaOverrides
}

func (w *WorkflowToolLocator) SetToolId(toolId string) {
	w.ToolId = toolId
}

func (w *WorkflowToolLocator) SetSchemaOverrides(schemaOverrides map[string]*WorkflowToolLocatorSchemaOverridesValue) {
	w.SchemaOverrides = schemaOverrides
}

func (w *WorkflowToolLocator) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowToolLocator) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowToolLocator to string"
	}
	return string(jsonData)
}

func (w *WorkflowToolLocator) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowToolLocator) MarshalJSON() ([]byte, error) {
	type embed WorkflowToolLocator
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
