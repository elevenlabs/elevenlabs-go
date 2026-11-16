package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ExotelWarmTransferConfigInput struct {
	TransferDestination *ExotelWarmTransferConfigInputTransferDestination `json:"transfer_destination" xml:"transfer_destination" required:"true"`
	extraProperties     map[string]interface{}
}

func (e *ExotelWarmTransferConfigInput) GetTransferDestination() (value *ExotelWarmTransferConfigInputTransferDestination) {
	if e == nil {
		return
	}
	return e.TransferDestination
}

func (e *ExotelWarmTransferConfigInput) SetTransferDestination(transferDestination *ExotelWarmTransferConfigInputTransferDestination) {
	e.TransferDestination = transferDestination
}

func (e *ExotelWarmTransferConfigInput) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e ExotelWarmTransferConfigInput) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ExotelWarmTransferConfigInput to string"
	}
	return string(jsonData)
}

func (e *ExotelWarmTransferConfigInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, e); err != nil {
		return err
	}
	type alias ExotelWarmTransferConfigInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*e = ExotelWarmTransferConfigInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *e, "type")
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *ExotelWarmTransferConfigInput) MarshalJSON() ([]byte, error) {
	type embed ExotelWarmTransferConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
