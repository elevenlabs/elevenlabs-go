package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type MetricRecord struct {
	ElapsedTime     float64 `json:"elapsed_time" xml:"elapsed_time" required:"true"`
	extraProperties map[string]interface{}
}

func (m *MetricRecord) GetElapsedTime() (value float64) {
	if m == nil {
		return
	}
	return m.ElapsedTime
}

func (m *MetricRecord) SetElapsedTime(elapsedTime float64) {
	m.ElapsedTime = elapsedTime
}

func (m *MetricRecord) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MetricRecord) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MetricRecord to string"
	}
	return string(jsonData)
}

func (m *MetricRecord) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias MetricRecord
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = MetricRecord(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MetricRecord) MarshalJSON() ([]byte, error) {
	type embed MetricRecord
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
