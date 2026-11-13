package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DashboardDataCollectionChartModel struct {
	Name             string `json:"name" xml:"name" required:"true"`
	DataCollectionId string `json:"data_collection_id" xml:"data_collection_id" required:"true"`
	extraProperties  map[string]interface{}
}

func (d *DashboardDataCollectionChartModel) GetName() (value string) {
	if d == nil {
		return
	}
	return d.Name
}

func (d *DashboardDataCollectionChartModel) GetDataCollectionId() (value string) {
	if d == nil {
		return
	}
	return d.DataCollectionId
}

func (d *DashboardDataCollectionChartModel) SetName(name string) {
	d.Name = name
}

func (d *DashboardDataCollectionChartModel) SetDataCollectionId(dataCollectionId string) {
	d.DataCollectionId = dataCollectionId
}

func (d *DashboardDataCollectionChartModel) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DashboardDataCollectionChartModel) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DashboardDataCollectionChartModel to string"
	}
	return string(jsonData)
}

func (d *DashboardDataCollectionChartModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DashboardDataCollectionChartModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DashboardDataCollectionChartModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DashboardDataCollectionChartModel) MarshalJSON() ([]byte, error) {
	type embed DashboardDataCollectionChartModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
