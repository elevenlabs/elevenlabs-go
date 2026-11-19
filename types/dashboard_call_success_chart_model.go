package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DashboardCallSuccessChartModel struct {
	Name            string `json:"name" xml:"name" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DashboardCallSuccessChartModel) GetName() (value string) {
	if d == nil {
		return
	}
	return d.Name
}

func (d *DashboardCallSuccessChartModel) SetName(name string) {
	d.Name = name
}

func (d *DashboardCallSuccessChartModel) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DashboardCallSuccessChartModel) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DashboardCallSuccessChartModel to string"
	}
	return string(jsonData)
}

func (d *DashboardCallSuccessChartModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DashboardCallSuccessChartModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DashboardCallSuccessChartModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DashboardCallSuccessChartModel) MarshalJSON() ([]byte, error) {
	type embed DashboardCallSuccessChartModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
