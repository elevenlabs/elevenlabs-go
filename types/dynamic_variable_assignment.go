package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Configuration for extracting values from tool responses and assigning them to dynamic variables.
type DynamicVariableAssignment struct {
	// The source to extract the value from. Currently only 'response' is supported.
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The name of the dynamic variable to assign the extracted value to
	DynamicVariable string `json:"dynamic_variable" xml:"dynamic_variable" required:"true"`
	// Dot notation path to extract the value from the source (e.g., 'user.name' or 'data.0.id')
	ValuePath string `json:"value_path" xml:"value_path" required:"true"`
	// If true, this assignment's value will be removed from the tool response before sending to the LLM and transcript, but still processed for variable assignment.
	Sanitize *bool `json:"sanitize,omitempty" xml:"sanitize,omitempty"`
	// If true, non-scalar values (lists, objects) extracted from the tool response are stored as their native type instead of being stringified to JSON. Enable this to use extracted arrays directly as list dynamic variables.
	PreserveNativeType *bool `json:"preserve_native_type,omitempty" xml:"preserve_native_type,omitempty"`
	extraProperties    map[string]interface{}
}

func (d *DynamicVariableAssignment) GetSource() (value *string) {
	if d == nil {
		return
	}
	return d.Source
}

func (d *DynamicVariableAssignment) GetDynamicVariable() (value string) {
	if d == nil {
		return
	}
	return d.DynamicVariable
}

func (d *DynamicVariableAssignment) GetValuePath() (value string) {
	if d == nil {
		return
	}
	return d.ValuePath
}

func (d *DynamicVariableAssignment) GetSanitize() (value *bool) {
	if d == nil {
		return
	}
	return d.Sanitize
}

func (d *DynamicVariableAssignment) GetPreserveNativeType() (value *bool) {
	if d == nil {
		return
	}
	return d.PreserveNativeType
}

func (d *DynamicVariableAssignment) SetSource(source *string) {
	d.Source = source
}

func (d *DynamicVariableAssignment) SetDynamicVariable(dynamicVariable string) {
	d.DynamicVariable = dynamicVariable
}

func (d *DynamicVariableAssignment) SetValuePath(valuePath string) {
	d.ValuePath = valuePath
}

func (d *DynamicVariableAssignment) SetSanitize(sanitize *bool) {
	d.Sanitize = sanitize
}

func (d *DynamicVariableAssignment) SetPreserveNativeType(preserveNativeType *bool) {
	d.PreserveNativeType = preserveNativeType
}

func (d *DynamicVariableAssignment) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DynamicVariableAssignment) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DynamicVariableAssignment to string"
	}
	return string(jsonData)
}

func (d *DynamicVariableAssignment) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DynamicVariableAssignment
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DynamicVariableAssignment(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DynamicVariableAssignment) MarshalJSON() ([]byte, error) {
	type embed DynamicVariableAssignment
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
