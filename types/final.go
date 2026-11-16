package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type Final struct {
	// Marks the end of the closing flush sequence.
	IsFinal         bool `json:"is_final" xml:"is_final" required:"true"`
	extraProperties map[string]interface{}
}

func (f *Final) GetIsFinal() (value bool) {
	if f == nil {
		return
	}
	return f.IsFinal
}

func (f *Final) SetIsFinal(isFinal bool) {
	f.IsFinal = isFinal
}

func (f *Final) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f Final) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: Final to string"
	}
	return string(jsonData)
}

func (f *Final) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, f); err != nil {
		return err
	}
	type alias Final
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*f = Final(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *Final) MarshalJSON() ([]byte, error) {
	type embed Final
	var marshaler = struct {
		embed
	}{
		embed: embed(*f),
	}
	return json.Marshal(marshaler)
}
