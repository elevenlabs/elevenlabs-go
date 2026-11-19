package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ExotelColdBlindTransferConfigOutput struct {
	TransferDestination *ExotelColdBlindTransferConfigOutputTransferDestination `json:"transfer_destination" xml:"transfer_destination" required:"true"`
	type_               string
	extraProperties     map[string]interface{}
}

func (e *ExotelColdBlindTransferConfigOutput) GetTransferDestination() (value *ExotelColdBlindTransferConfigOutputTransferDestination) {
	if e == nil {
		return
	}
	return e.TransferDestination
}

func (e *ExotelColdBlindTransferConfigOutput) SetTransferDestination(transferDestination *ExotelColdBlindTransferConfigOutputTransferDestination) {
	e.TransferDestination = transferDestination
}

func (e *ExotelColdBlindTransferConfigOutput) Type() string {
	return e.type_
}

func (e *ExotelColdBlindTransferConfigOutput) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e ExotelColdBlindTransferConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ExotelColdBlindTransferConfigOutput to string"
	}
	return string(jsonData)
}

func (e *ExotelColdBlindTransferConfigOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, e); err != nil {
		return err
	}
	type alias ExotelColdBlindTransferConfigOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*e = ExotelColdBlindTransferConfigOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "exotel_cold_blind_transfer" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", e, "exotel_cold_blind_transfer", *literals.Type)
	}
	e.type_ = "exotel_cold_blind_transfer"
	extraProperties, err := internal.ExtractExtraProperties(data, *e, "type")
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *ExotelColdBlindTransferConfigOutput) MarshalJSON() ([]byte, error) {
	type embed ExotelColdBlindTransferConfigOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*e),
		Type:  "exotel_cold_blind_transfer",
	}
	return json.Marshal(marshaler)
}
