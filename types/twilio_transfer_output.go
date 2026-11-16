package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TwilioTransferOutput struct {
	type_           string
	Config          *TwilioTransferOutputConfig `json:"config" xml:"config" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TwilioTransferOutput) GetConfig() (value *TwilioTransferOutputConfig) {
	if t == nil {
		return
	}
	return t.Config
}

func (t *TwilioTransferOutput) SetConfig(config *TwilioTransferOutputConfig) {
	t.Config = config
}

func (t *TwilioTransferOutput) Type() string {
	return t.type_
}

func (t *TwilioTransferOutput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TwilioTransferOutput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TwilioTransferOutput to string"
	}
	return string(jsonData)
}

func (t *TwilioTransferOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TwilioTransferOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TwilioTransferOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "twilio" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "twilio", *literals.Type)
	}
	t.type_ = "twilio"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TwilioTransferOutput) MarshalJSON() ([]byte, error) {
	type embed TwilioTransferOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*t),
		Type:  "twilio",
	}
	return json.Marshal(marshaler)
}
