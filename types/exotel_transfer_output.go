package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ExotelTransferOutput struct {
	type_           string
	Config          *ExotelTransferOutputConfig `json:"config" xml:"config" required:"true"`
	extraProperties map[string]interface{}
}

func (e *ExotelTransferOutput) GetConfig() (value *ExotelTransferOutputConfig) {
	if e == nil {
		return
	}
	return e.Config
}

func (e *ExotelTransferOutput) SetConfig(config *ExotelTransferOutputConfig) {
	e.Config = config
}

func (e *ExotelTransferOutput) Type() string {
	return e.type_
}

func (e *ExotelTransferOutput) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e ExotelTransferOutput) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ExotelTransferOutput to string"
	}
	return string(jsonData)
}

func (e *ExotelTransferOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, e); err != nil {
		return err
	}
	type alias ExotelTransferOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*e = ExotelTransferOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "exotel" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", e, "exotel", *literals.Type)
	}
	e.type_ = "exotel"
	extraProperties, err := internal.ExtractExtraProperties(data, *e, "type")
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *ExotelTransferOutput) MarshalJSON() ([]byte, error) {
	type embed ExotelTransferOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*e),
		Type:  "exotel",
	}
	return json.Marshal(marshaler)
}
