package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Represents a calendar date as individual numeric components.
type DateModel struct {
	// Year
	Year int `json:"year" xml:"year" required:"true" min:"2026" max:"2100"`
	// Month (1-12)
	Month int `json:"month" xml:"month" required:"true" min:"1" max:"12"`
	// Day of month (1-31)
	Day             int `json:"day" xml:"day" required:"true" min:"1" max:"31"`
	extraProperties map[string]interface{}
}

func (d *DateModel) GetYear() (value int) {
	if d == nil {
		return
	}
	return d.Year
}

func (d *DateModel) GetMonth() (value int) {
	if d == nil {
		return
	}
	return d.Month
}

func (d *DateModel) GetDay() (value int) {
	if d == nil {
		return
	}
	return d.Day
}

func (d *DateModel) SetYear(year int) {
	d.Year = year
}

func (d *DateModel) SetMonth(month int) {
	d.Month = month
}

func (d *DateModel) SetDay(day int) {
	d.Day = day
}

func (d *DateModel) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DateModel) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DateModel to string"
	}
	return string(jsonData)
}

func (d *DateModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DateModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DateModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DateModel) MarshalJSON() ([]byte, error) {
	type embed DateModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
