package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ExotelTransferInput struct {
	Config          *ExotelTransferInputConfig `json:"config" xml:"config" required:"true"`
	extraProperties map[string]interface{}
}

func (e *ExotelTransferInput) GetConfig() (value *ExotelTransferInputConfig) {
	if e == nil {
		return
	}
	return e.Config
}

func (e *ExotelTransferInput) SetConfig(config *ExotelTransferInputConfig) {
	e.Config = config
}

func (e *ExotelTransferInput) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e ExotelTransferInput) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ExotelTransferInput to string"
	}
	return string(jsonData)
}

func (e *ExotelTransferInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, e); err != nil {
		return err
	}
	type alias ExotelTransferInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*e = ExotelTransferInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *e, "type")
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *ExotelTransferInput) MarshalJSON() ([]byte, error) {
	type embed ExotelTransferInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
