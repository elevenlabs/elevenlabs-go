package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ColumnFilter struct {
	Column          string                    `json:"column" xml:"column" required:"true"`
	Operation       ColumnFilterOperation     `json:"operation" xml:"operation" required:"true"`
	Values          []*ColumnFilterValuesItem `json:"values" xml:"values" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ColumnFilter) GetColumn() (value string) {
	if c == nil {
		return
	}
	return c.Column
}

func (c *ColumnFilter) GetOperation() (value ColumnFilterOperation) {
	if c == nil {
		return
	}
	return c.Operation
}

func (c *ColumnFilter) GetValues() (value []*ColumnFilterValuesItem) {
	if c == nil {
		return
	}
	return c.Values
}

func (c *ColumnFilter) SetColumn(column string) {
	c.Column = column
}

func (c *ColumnFilter) SetOperation(operation ColumnFilterOperation) {
	c.Operation = operation
}

func (c *ColumnFilter) SetValues(values []*ColumnFilterValuesItem) {
	c.Values = values
}

func (c *ColumnFilter) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ColumnFilter) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ColumnFilter to string"
	}
	return string(jsonData)
}

func (c *ColumnFilter) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ColumnFilter
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ColumnFilter(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ColumnFilter) MarshalJSON() ([]byte, error) {
	type embed ColumnFilter
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
