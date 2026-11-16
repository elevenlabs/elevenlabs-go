package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TwilioTransferInput struct {
	Config          *TwilioTransferInputConfig `json:"config" xml:"config" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TwilioTransferInput) GetConfig() (value *TwilioTransferInputConfig) {
	if t == nil {
		return
	}
	return t.Config
}

func (t *TwilioTransferInput) SetConfig(config *TwilioTransferInputConfig) {
	t.Config = config
}

func (t *TwilioTransferInput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TwilioTransferInput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TwilioTransferInput to string"
	}
	return string(jsonData)
}

func (t *TwilioTransferInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TwilioTransferInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TwilioTransferInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TwilioTransferInput) MarshalJSON() ([]byte, error) {
	type embed TwilioTransferInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
