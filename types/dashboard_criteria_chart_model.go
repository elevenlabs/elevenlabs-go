package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DashboardCriteriaChartModel struct {
	Name            string `json:"name" xml:"name" required:"true"`
	CriteriaId      string `json:"criteria_id" xml:"criteria_id" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DashboardCriteriaChartModel) GetName() (value string) {
	if d == nil {
		return
	}
	return d.Name
}

func (d *DashboardCriteriaChartModel) GetCriteriaId() (value string) {
	if d == nil {
		return
	}
	return d.CriteriaId
}

func (d *DashboardCriteriaChartModel) SetName(name string) {
	d.Name = name
}

func (d *DashboardCriteriaChartModel) SetCriteriaId(criteriaId string) {
	d.CriteriaId = criteriaId
}

func (d *DashboardCriteriaChartModel) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DashboardCriteriaChartModel) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DashboardCriteriaChartModel to string"
	}
	return string(jsonData)
}

func (d *DashboardCriteriaChartModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DashboardCriteriaChartModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DashboardCriteriaChartModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DashboardCriteriaChartModel) MarshalJSON() ([]byte, error) {
	type embed DashboardCriteriaChartModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
