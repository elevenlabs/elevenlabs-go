package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type MetronomeUsageColumnFilter struct {
	Column          MetronomeUsageColumnFilterColumn        `json:"column" xml:"column" required:"true"`
	Operation       MetronomeUsageColumnFilterOperation     `json:"operation" xml:"operation" required:"true"`
	Values          []*MetronomeUsageColumnFilterValuesItem `json:"values" xml:"values" required:"true"`
	extraProperties map[string]interface{}
}

func (m *MetronomeUsageColumnFilter) GetColumn() (value MetronomeUsageColumnFilterColumn) {
	if m == nil {
		return
	}
	return m.Column
}

func (m *MetronomeUsageColumnFilter) GetOperation() (value MetronomeUsageColumnFilterOperation) {
	if m == nil {
		return
	}
	return m.Operation
}

func (m *MetronomeUsageColumnFilter) GetValues() (value []*MetronomeUsageColumnFilterValuesItem) {
	if m == nil {
		return
	}
	return m.Values
}

func (m *MetronomeUsageColumnFilter) SetColumn(column MetronomeUsageColumnFilterColumn) {
	m.Column = column
}

func (m *MetronomeUsageColumnFilter) SetOperation(operation MetronomeUsageColumnFilterOperation) {
	m.Operation = operation
}

func (m *MetronomeUsageColumnFilter) SetValues(values []*MetronomeUsageColumnFilterValuesItem) {
	m.Values = values
}

func (m *MetronomeUsageColumnFilter) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MetronomeUsageColumnFilter) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MetronomeUsageColumnFilter to string"
	}
	return string(jsonData)
}

func (m *MetronomeUsageColumnFilter) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias MetronomeUsageColumnFilter
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = MetronomeUsageColumnFilter(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MetronomeUsageColumnFilter) MarshalJSON() ([]byte, error) {
	type embed MetronomeUsageColumnFilter
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
