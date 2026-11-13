package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AllowlistItem struct {
	// The hostname of the allowed origin
	Hostname        string `json:"hostname" xml:"hostname" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AllowlistItem) GetHostname() (value string) {
	if a == nil {
		return
	}
	return a.Hostname
}

func (a *AllowlistItem) SetHostname(hostname string) {
	a.Hostname = hostname
}

func (a *AllowlistItem) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AllowlistItem) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AllowlistItem to string"
	}
	return string(jsonData)
}

func (a *AllowlistItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AllowlistItem
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AllowlistItem(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AllowlistItem) MarshalJSON() ([]byte, error) {
	type embed AllowlistItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
