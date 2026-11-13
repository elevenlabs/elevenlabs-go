package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type StudioTextStyleShadowModel struct {
	Enabled         bool    `json:"enabled" xml:"enabled" required:"true"`
	Color           string  `json:"color" xml:"color" required:"true"`
	Opacity         float64 `json:"opacity" xml:"opacity" required:"true"`
	Blur            float64 `json:"blur" xml:"blur" required:"true"`
	OffsetX         float64 `json:"offset_x" xml:"offset_x" required:"true"`
	OffsetY         float64 `json:"offset_y" xml:"offset_y" required:"true"`
	extraProperties map[string]interface{}
}

func (s *StudioTextStyleShadowModel) GetEnabled() (value bool) {
	if s == nil {
		return
	}
	return s.Enabled
}

func (s *StudioTextStyleShadowModel) GetColor() (value string) {
	if s == nil {
		return
	}
	return s.Color
}

func (s *StudioTextStyleShadowModel) GetOpacity() (value float64) {
	if s == nil {
		return
	}
	return s.Opacity
}

func (s *StudioTextStyleShadowModel) GetBlur() (value float64) {
	if s == nil {
		return
	}
	return s.Blur
}

func (s *StudioTextStyleShadowModel) GetOffsetX() (value float64) {
	if s == nil {
		return
	}
	return s.OffsetX
}

func (s *StudioTextStyleShadowModel) GetOffsetY() (value float64) {
	if s == nil {
		return
	}
	return s.OffsetY
}

func (s *StudioTextStyleShadowModel) SetEnabled(enabled bool) {
	s.Enabled = enabled
}

func (s *StudioTextStyleShadowModel) SetColor(color string) {
	s.Color = color
}

func (s *StudioTextStyleShadowModel) SetOpacity(opacity float64) {
	s.Opacity = opacity
}

func (s *StudioTextStyleShadowModel) SetBlur(blur float64) {
	s.Blur = blur
}

func (s *StudioTextStyleShadowModel) SetOffsetX(offsetX float64) {
	s.OffsetX = offsetX
}

func (s *StudioTextStyleShadowModel) SetOffsetY(offsetY float64) {
	s.OffsetY = offsetY
}

func (s *StudioTextStyleShadowModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s StudioTextStyleShadowModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: StudioTextStyleShadowModel to string"
	}
	return string(jsonData)
}

func (s *StudioTextStyleShadowModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias StudioTextStyleShadowModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = StudioTextStyleShadowModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *StudioTextStyleShadowModel) MarshalJSON() ([]byte, error) {
	type embed StudioTextStyleShadowModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
