package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ExpiresAfter struct {
	Anchor          *string `json:"anchor,omitempty" xml:"anchor,omitempty"`
	Seconds         int     `json:"seconds" xml:"seconds" required:"true"`
	extraProperties map[string]interface{}
}

func (e *ExpiresAfter) GetAnchor() (value *string) {
	if e == nil {
		return
	}
	return e.Anchor
}

func (e *ExpiresAfter) GetSeconds() (value int) {
	if e == nil {
		return
	}
	return e.Seconds
}

func (e *ExpiresAfter) SetAnchor(anchor *string) {
	e.Anchor = anchor
}

func (e *ExpiresAfter) SetSeconds(seconds int) {
	e.Seconds = seconds
}

func (e *ExpiresAfter) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e ExpiresAfter) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ExpiresAfter to string"
	}
	return string(jsonData)
}

func (e *ExpiresAfter) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, e); err != nil {
		return err
	}
	type alias ExpiresAfter
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*e = ExpiresAfter(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *ExpiresAfter) MarshalJSON() ([]byte, error) {
	type embed ExpiresAfter
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
