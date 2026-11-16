package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ExotelColdBlindTransferConfigInput struct {
	TransferDestination *ExotelColdBlindTransferConfigInputTransferDestination `json:"transfer_destination" xml:"transfer_destination" required:"true"`
	extraProperties     map[string]interface{}
}

func (e *ExotelColdBlindTransferConfigInput) GetTransferDestination() (value *ExotelColdBlindTransferConfigInputTransferDestination) {
	if e == nil {
		return
	}
	return e.TransferDestination
}

func (e *ExotelColdBlindTransferConfigInput) SetTransferDestination(transferDestination *ExotelColdBlindTransferConfigInputTransferDestination) {
	e.TransferDestination = transferDestination
}

func (e *ExotelColdBlindTransferConfigInput) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e ExotelColdBlindTransferConfigInput) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ExotelColdBlindTransferConfigInput to string"
	}
	return string(jsonData)
}

func (e *ExotelColdBlindTransferConfigInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, e); err != nil {
		return err
	}
	type alias ExotelColdBlindTransferConfigInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*e = ExotelColdBlindTransferConfigInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *e, "type")
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *ExotelColdBlindTransferConfigInput) MarshalJSON() ([]byte, error) {
	type embed ExotelColdBlindTransferConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
