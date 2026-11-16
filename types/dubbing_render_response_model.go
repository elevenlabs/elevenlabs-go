package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingRenderResponseModel struct {
	Version         int    `json:"version" xml:"version" required:"true"`
	RenderId        string `json:"render_id" xml:"render_id" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DubbingRenderResponseModel) GetVersion() (value int) {
	if d == nil {
		return
	}
	return d.Version
}

func (d *DubbingRenderResponseModel) GetRenderId() (value string) {
	if d == nil {
		return
	}
	return d.RenderId
}

func (d *DubbingRenderResponseModel) SetVersion(version int) {
	d.Version = version
}

func (d *DubbingRenderResponseModel) SetRenderId(renderId string) {
	d.RenderId = renderId
}

func (d *DubbingRenderResponseModel) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingRenderResponseModel) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingRenderResponseModel to string"
	}
	return string(jsonData)
}

func (d *DubbingRenderResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingRenderResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingRenderResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingRenderResponseModel) MarshalJSON() ([]byte, error) {
	type embed DubbingRenderResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
