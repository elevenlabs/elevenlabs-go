package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Definition for a tool the client can call.
type Tool struct {
	Name            string           `json:"name" xml:"name" required:"true"`
	Title           *string          `json:"title,omitempty" xml:"title,omitempty"`
	Description     *string          `json:"description,omitempty" xml:"description,omitempty"`
	InputSchema     map[string]any   `json:"inputSchema" required:"true"`
	Execution       *ToolExecution   `json:"execution,omitempty" xml:"execution,omitempty"`
	OutputSchema    map[string]any   `json:"outputSchema,omitempty"`
	Icons           []*Icon          `json:"icons,omitempty" xml:"icons,omitempty"`
	Annotations     *ToolAnnotations `json:"annotations,omitempty" xml:"annotations,omitempty"`
	Meta            map[string]any   `json:"_meta,omitempty"`
	extraProperties map[string]interface{}
}

func (t *Tool) GetName() (value string) {
	if t == nil {
		return
	}
	return t.Name
}

func (t *Tool) GetTitle() (value *string) {
	if t == nil {
		return
	}
	return t.Title
}

func (t *Tool) GetDescription() (value *string) {
	if t == nil {
		return
	}
	return t.Description
}

func (t *Tool) GetInputSchema() (value map[string]any) {
	if t == nil {
		return
	}
	return t.InputSchema
}

func (t *Tool) GetExecution() (value *ToolExecution) {
	if t == nil {
		return
	}
	return t.Execution
}

func (t *Tool) GetOutputSchema() (value map[string]any) {
	if t == nil {
		return
	}
	return t.OutputSchema
}

func (t *Tool) GetIcons() (value []*Icon) {
	if t == nil {
		return
	}
	return t.Icons
}

func (t *Tool) GetAnnotations() (value *ToolAnnotations) {
	if t == nil {
		return
	}
	return t.Annotations
}

func (t *Tool) GetMeta() (value map[string]any) {
	if t == nil {
		return
	}
	return t.Meta
}

func (t *Tool) SetName(name string) {
	t.Name = name
}

func (t *Tool) SetTitle(title *string) {
	t.Title = title
}

func (t *Tool) SetDescription(description *string) {
	t.Description = description
}

func (t *Tool) SetInputSchema(inputSchema map[string]any) {
	t.InputSchema = inputSchema
}

func (t *Tool) SetExecution(execution *ToolExecution) {
	t.Execution = execution
}

func (t *Tool) SetOutputSchema(outputSchema map[string]any) {
	t.OutputSchema = outputSchema
}

func (t *Tool) SetIcons(icons []*Icon) {
	t.Icons = icons
}

func (t *Tool) SetAnnotations(annotations *ToolAnnotations) {
	t.Annotations = annotations
}

func (t *Tool) SetMeta(meta map[string]any) {
	t.Meta = meta
}

func (t *Tool) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t Tool) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: Tool to string"
	}
	return string(jsonData)
}

func (t *Tool) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *Tool) MarshalJSON() ([]byte, error) {
	type embed Tool
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
