package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ExotelWarmTransferConfigOutput struct {
	TransferDestination *ExotelWarmTransferConfigOutputTransferDestination `json:"transfer_destination" xml:"transfer_destination" required:"true"`
	type_               string
	extraProperties     map[string]interface{}
}

func (e *ExotelWarmTransferConfigOutput) GetTransferDestination() (value *ExotelWarmTransferConfigOutputTransferDestination) {
	if e == nil {
		return
	}
	return e.TransferDestination
}

func (e *ExotelWarmTransferConfigOutput) SetTransferDestination(transferDestination *ExotelWarmTransferConfigOutputTransferDestination) {
	e.TransferDestination = transferDestination
}

func (e *ExotelWarmTransferConfigOutput) Type() string {
	return e.type_
}

func (e *ExotelWarmTransferConfigOutput) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e ExotelWarmTransferConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ExotelWarmTransferConfigOutput to string"
	}
	return string(jsonData)
}

func (e *ExotelWarmTransferConfigOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, e); err != nil {
		return err
	}
	type alias ExotelWarmTransferConfigOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*e = ExotelWarmTransferConfigOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "exotel_warm_transfer" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", e, "exotel_warm_transfer", *literals.Type)
	}
	e.type_ = "exotel_warm_transfer"
	extraProperties, err := internal.ExtractExtraProperties(data, *e, "type")
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *ExotelWarmTransferConfigOutput) MarshalJSON() ([]byte, error) {
	type embed ExotelWarmTransferConfigOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*e),
		Type:  "exotel_warm_transfer",
	}
	return json.Marshal(marshaler)
}
