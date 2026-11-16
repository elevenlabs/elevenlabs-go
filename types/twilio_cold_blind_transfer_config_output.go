package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TwilioColdBlindTransferConfigOutput struct {
	TransferDestination *TwilioColdBlindTransferConfigOutputTransferDestination `json:"transfer_destination" xml:"transfer_destination" required:"true"`
	// DTMF digits to send after the call connects (e.g. 'ww1234' for an extension). Static or a dynamic variable reference. Use 'w' for a 0.5s pause.
	PostDialDigits  *TwilioColdBlindTransferConfigOutputPostDialDigits `json:"post_dial_digits" xml:"post_dial_digits" required:"true"`
	type_           string
	extraProperties map[string]interface{}
}

func (t *TwilioColdBlindTransferConfigOutput) GetTransferDestination() (value *TwilioColdBlindTransferConfigOutputTransferDestination) {
	if t == nil {
		return
	}
	return t.TransferDestination
}

func (t *TwilioColdBlindTransferConfigOutput) GetPostDialDigits() (value *TwilioColdBlindTransferConfigOutputPostDialDigits) {
	if t == nil {
		return
	}
	return t.PostDialDigits
}

func (t *TwilioColdBlindTransferConfigOutput) SetTransferDestination(transferDestination *TwilioColdBlindTransferConfigOutputTransferDestination) {
	t.TransferDestination = transferDestination
}

func (t *TwilioColdBlindTransferConfigOutput) SetPostDialDigits(postDialDigits *TwilioColdBlindTransferConfigOutputPostDialDigits) {
	t.PostDialDigits = postDialDigits
}

func (t *TwilioColdBlindTransferConfigOutput) Type() string {
	return t.type_
}

func (t *TwilioColdBlindTransferConfigOutput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TwilioColdBlindTransferConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TwilioColdBlindTransferConfigOutput to string"
	}
	return string(jsonData)
}

func (t *TwilioColdBlindTransferConfigOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "twilio_cold_blind_transfer" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "twilio_cold_blind_transfer", *literals.Type)
	}
	t.type_ = "twilio_cold_blind_transfer"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TwilioColdBlindTransferConfigOutput) MarshalJSON() ([]byte, error) {
	type embed TwilioColdBlindTransferConfigOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*t),
		Type:  "twilio_cold_blind_transfer",
	}
	return json.Marshal(marshaler)
}
