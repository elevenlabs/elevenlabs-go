package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type WorkspaceAnalyticsQueryResponseModel struct {
	Columns         []string                                              `json:"columns" xml:"columns" required:"true"`
	ColumnTypes     []WorkspaceAnalyticsQueryResponseModelColumnTypesItem `json:"column_types" xml:"column_types" required:"true"`
	Rows            [][]*WorkspaceAnalyticsQueryResponseModelRowsItemItem `json:"rows" xml:"rows" required:"true"`
	ColumnUnits     []*ColumnUnit                                         `json:"column_units" xml:"column_units" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkspaceAnalyticsQueryResponseModel) GetColumns() (value []string) {
	if w == nil {
		return
	}
	return w.Columns
}

func (w *WorkspaceAnalyticsQueryResponseModel) GetColumnTypes() (value []WorkspaceAnalyticsQueryResponseModelColumnTypesItem) {
	if w == nil {
		return
	}
	return w.ColumnTypes
}

func (w *WorkspaceAnalyticsQueryResponseModel) GetRows() (value [][]*WorkspaceAnalyticsQueryResponseModelRowsItemItem) {
	if w == nil {
		return
	}
	return w.Rows
}

func (w *WorkspaceAnalyticsQueryResponseModel) GetColumnUnits() (value []*ColumnUnit) {
	if w == nil {
		return
	}
	return w.ColumnUnits
}

func (w *WorkspaceAnalyticsQueryResponseModel) SetColumns(columns []string) {
	w.Columns = columns
}

func (w *WorkspaceAnalyticsQueryResponseModel) SetColumnTypes(columnTypes []WorkspaceAnalyticsQueryResponseModelColumnTypesItem) {
	w.ColumnTypes = columnTypes
}

func (w *WorkspaceAnalyticsQueryResponseModel) SetRows(rows [][]*WorkspaceAnalyticsQueryResponseModelRowsItemItem) {
	w.Rows = rows
}

func (w *WorkspaceAnalyticsQueryResponseModel) SetColumnUnits(columnUnits []*ColumnUnit) {
	w.ColumnUnits = columnUnits
}

func (w *WorkspaceAnalyticsQueryResponseModel) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkspaceAnalyticsQueryResponseModel) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkspaceAnalyticsQueryResponseModel to string"
	}
	return string(jsonData)
}

func (w *WorkspaceAnalyticsQueryResponseModel) UnmarshalJSON(data []byte) error {
	type unmarshaler WorkspaceAnalyticsQueryResponseModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WorkspaceAnalyticsQueryResponseModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkspaceAnalyticsQueryResponseModel) MarshalJSON() ([]byte, error) {
	type embed WorkspaceAnalyticsQueryResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
