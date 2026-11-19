package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type UsageCharactersResponseModel struct {
	// The time axis with unix timestamps for each day.
	Time []int `json:"time" xml:"time" required:"true"`
	// The usage of each breakdown type along the time axis.
	Usage           map[string][]float64 `json:"usage" required:"true"`
	extraProperties map[string]interface{}
}

func (u *UsageCharactersResponseModel) GetTime() (value []int) {
	if u == nil {
		return
	}
	return u.Time
}

func (u *UsageCharactersResponseModel) GetUsage() (value map[string][]float64) {
	if u == nil {
		return
	}
	return u.Usage
}

func (u *UsageCharactersResponseModel) SetTime(time []int) {
	u.Time = time
}

func (u *UsageCharactersResponseModel) SetUsage(usage map[string][]float64) {
	u.Usage = usage
}

func (u *UsageCharactersResponseModel) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UsageCharactersResponseModel) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UsageCharactersResponseModel to string"
	}
	return string(jsonData)
}

func (u *UsageCharactersResponseModel) UnmarshalJSON(data []byte) error {
	type unmarshaler UsageCharactersResponseModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UsageCharactersResponseModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UsageCharactersResponseModel) MarshalJSON() ([]byte, error) {
	type embed UsageCharactersResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
