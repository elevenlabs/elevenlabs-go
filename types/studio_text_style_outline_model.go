package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type StudioTextStyleOutlineModel struct {
	Enabled         bool    `json:"enabled" xml:"enabled" required:"true"`
	Color           string  `json:"color" xml:"color" required:"true"`
	Opacity         float64 `json:"opacity" xml:"opacity" required:"true"`
	Width           float64 `json:"width" xml:"width" required:"true"`
	extraProperties map[string]interface{}
}

func (s *StudioTextStyleOutlineModel) GetEnabled() (value bool) {
	if s == nil {
		return
	}
	return s.Enabled
}

func (s *StudioTextStyleOutlineModel) GetColor() (value string) {
	if s == nil {
		return
	}
	return s.Color
}

func (s *StudioTextStyleOutlineModel) GetOpacity() (value float64) {
	if s == nil {
		return
	}
	return s.Opacity
}

func (s *StudioTextStyleOutlineModel) GetWidth() (value float64) {
	if s == nil {
		return
	}
	return s.Width
}

func (s *StudioTextStyleOutlineModel) SetEnabled(enabled bool) {
	s.Enabled = enabled
}

func (s *StudioTextStyleOutlineModel) SetColor(color string) {
	s.Color = color
}

func (s *StudioTextStyleOutlineModel) SetOpacity(opacity float64) {
	s.Opacity = opacity
}

func (s *StudioTextStyleOutlineModel) SetWidth(width float64) {
	s.Width = width
}

func (s *StudioTextStyleOutlineModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s StudioTextStyleOutlineModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: StudioTextStyleOutlineModel to string"
	}
	return string(jsonData)
}

func (s *StudioTextStyleOutlineModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias StudioTextStyleOutlineModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = StudioTextStyleOutlineModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *StudioTextStyleOutlineModel) MarshalJSON() ([]byte, error) {
	type embed StudioTextStyleOutlineModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
