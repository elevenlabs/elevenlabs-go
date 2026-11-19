package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// An optionally-sized icon for display in a user interface (2025-11-25+).
type Icon struct {
	Src             string     `json:"src" xml:"src" required:"true"`
	MimeType        *string    `json:"mimeType,omitempty" xml:"mimeType,omitempty"`
	Sizes           []string   `json:"sizes,omitempty" xml:"sizes,omitempty"`
	Theme           *IconTheme `json:"theme,omitempty" xml:"theme,omitempty"`
	extraProperties map[string]interface{}
}

func (i *Icon) GetSrc() (value string) {
	if i == nil {
		return
	}
	return i.Src
}

func (i *Icon) GetMimeType() (value *string) {
	if i == nil {
		return
	}
	return i.MimeType
}

func (i *Icon) GetSizes() (value []string) {
	if i == nil {
		return
	}
	return i.Sizes
}

func (i *Icon) GetTheme() (value *IconTheme) {
	if i == nil {
		return
	}
	return i.Theme
}

func (i *Icon) SetSrc(src string) {
	i.Src = src
}

func (i *Icon) SetMimeType(mimeType *string) {
	i.MimeType = mimeType
}

func (i *Icon) SetSizes(sizes []string) {
	i.Sizes = sizes
}

func (i *Icon) SetTheme(theme *IconTheme) {
	i.Theme = theme
}

func (i *Icon) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i Icon) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: Icon to string"
	}
	return string(jsonData)
}

func (i *Icon) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, i); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *Icon) MarshalJSON() ([]byte, error) {
	type embed Icon
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}
