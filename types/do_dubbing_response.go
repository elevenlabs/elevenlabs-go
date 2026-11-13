package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DoDubbingResponse struct {
	// The ID of the dubbing project.
	DubbingId string `json:"dubbing_id" xml:"dubbing_id" required:"true"`
	// The expected duration of the dubbing project in seconds.
	ExpectedDurationSec float64 `json:"expected_duration_sec" xml:"expected_duration_sec" required:"true"`
	extraProperties     map[string]interface{}
}

func (d *DoDubbingResponse) GetDubbingId() (value string) {
	if d == nil {
		return
	}
	return d.DubbingId
}

func (d *DoDubbingResponse) GetExpectedDurationSec() (value float64) {
	if d == nil {
		return
	}
	return d.ExpectedDurationSec
}

func (d *DoDubbingResponse) SetDubbingId(dubbingId string) {
	d.DubbingId = dubbingId
}

func (d *DoDubbingResponse) SetExpectedDurationSec(expectedDurationSec float64) {
	d.ExpectedDurationSec = expectedDurationSec
}

func (d *DoDubbingResponse) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DoDubbingResponse) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DoDubbingResponse to string"
	}
	return string(jsonData)
}

func (d *DoDubbingResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DoDubbingResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DoDubbingResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DoDubbingResponse) MarshalJSON() ([]byte, error) {
	type embed DoDubbingResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
