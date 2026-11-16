package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Tracks a dynamic variable update that occurred during tool execution.
type DynamicVariableUpdateCommonModel struct {
	VariableName    string  `json:"variable_name" xml:"variable_name" required:"true"`
	OldValue        *string `json:"old_value" xml:"old_value" required:"true"`
	NewValue        string  `json:"new_value" xml:"new_value" required:"true"`
	UpdatedAt       float64 `json:"updated_at" xml:"updated_at" required:"true"`
	ToolName        string  `json:"tool_name" xml:"tool_name" required:"true"`
	ToolRequestId   string  `json:"tool_request_id" xml:"tool_request_id" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DynamicVariableUpdateCommonModel) GetVariableName() (value string) {
	if d == nil {
		return
	}
	return d.VariableName
}

func (d *DynamicVariableUpdateCommonModel) GetOldValue() (value *string) {
	if d == nil {
		return
	}
	return d.OldValue
}

func (d *DynamicVariableUpdateCommonModel) GetNewValue() (value string) {
	if d == nil {
		return
	}
	return d.NewValue
}

func (d *DynamicVariableUpdateCommonModel) GetUpdatedAt() (value float64) {
	if d == nil {
		return
	}
	return d.UpdatedAt
}

func (d *DynamicVariableUpdateCommonModel) GetToolName() (value string) {
	if d == nil {
		return
	}
	return d.ToolName
}

func (d *DynamicVariableUpdateCommonModel) GetToolRequestId() (value string) {
	if d == nil {
		return
	}
	return d.ToolRequestId
}

func (d *DynamicVariableUpdateCommonModel) SetVariableName(variableName string) {
	d.VariableName = variableName
}

func (d *DynamicVariableUpdateCommonModel) SetOldValue(oldValue *string) {
	d.OldValue = oldValue
}

func (d *DynamicVariableUpdateCommonModel) SetNewValue(newValue string) {
	d.NewValue = newValue
}

func (d *DynamicVariableUpdateCommonModel) SetUpdatedAt(updatedAt float64) {
	d.UpdatedAt = updatedAt
}

func (d *DynamicVariableUpdateCommonModel) SetToolName(toolName string) {
	d.ToolName = toolName
}

func (d *DynamicVariableUpdateCommonModel) SetToolRequestId(toolRequestId string) {
	d.ToolRequestId = toolRequestId
}

func (d *DynamicVariableUpdateCommonModel) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DynamicVariableUpdateCommonModel) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DynamicVariableUpdateCommonModel to string"
	}
	return string(jsonData)
}

func (d *DynamicVariableUpdateCommonModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, d); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DynamicVariableUpdateCommonModel) MarshalJSON() ([]byte, error) {
	type embed DynamicVariableUpdateCommonModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
