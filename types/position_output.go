package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PositionOutput struct {
	X               float64 `json:"x" xml:"x" required:"true"`
	Y               float64 `json:"y" xml:"y" required:"true"`
	extraProperties map[string]interface{}
}

func (p *PositionOutput) GetX() (value float64) {
	if p == nil {
		return
	}
	return p.X
}

func (p *PositionOutput) GetY() (value float64) {
	if p == nil {
		return
	}
	return p.Y
}

func (p *PositionOutput) SetX(x float64) {
	p.X = x
}

func (p *PositionOutput) SetY(y float64) {
	p.Y = y
}

func (p *PositionOutput) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PositionOutput) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PositionOutput to string"
	}
	return string(jsonData)
}

func (p *PositionOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PositionOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PositionOutput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PositionOutput) MarshalJSON() ([]byte, error) {
	type embed PositionOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
