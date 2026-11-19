package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type PositionInput struct {
	X               *float64 `json:"x,omitempty" xml:"x,omitempty"`
	Y               *float64 `json:"y,omitempty" xml:"y,omitempty"`
	extraProperties map[string]interface{}
}

func (p *PositionInput) GetX() (value *float64) {
	if p == nil {
		return
	}
	return p.X
}

func (p *PositionInput) GetY() (value *float64) {
	if p == nil {
		return
	}
	return p.Y
}

func (p *PositionInput) SetX(x *float64) {
	p.X = x
}

func (p *PositionInput) SetY(y *float64) {
	p.Y = y
}

func (p *PositionInput) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PositionInput) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PositionInput to string"
	}
	return string(jsonData)
}

func (p *PositionInput) UnmarshalJSON(data []byte) error {
	type unmarshaler PositionInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PositionInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PositionInput) MarshalJSON() ([]byte, error) {
	type embed PositionInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
