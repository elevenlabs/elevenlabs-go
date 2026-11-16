package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TwilioColdBlindTransferConfigInput struct {
	TransferDestination *TwilioColdBlindTransferConfigInputTransferDestination `json:"transfer_destination" xml:"transfer_destination" required:"true"`
	// DTMF digits to send after the call connects (e.g. 'ww1234' for an extension). Static or a dynamic variable reference. Use 'w' for a 0.5s pause.
	PostDialDigits  *TwilioColdBlindTransferConfigInputPostDialDigits `json:"post_dial_digits,omitempty" xml:"post_dial_digits,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TwilioColdBlindTransferConfigInput) GetTransferDestination() (value *TwilioColdBlindTransferConfigInputTransferDestination) {
	if t == nil {
		return
	}
	return t.TransferDestination
}

func (t *TwilioColdBlindTransferConfigInput) GetPostDialDigits() (value *TwilioColdBlindTransferConfigInputPostDialDigits) {
	if t == nil {
		return
	}
	return t.PostDialDigits
}

func (t *TwilioColdBlindTransferConfigInput) SetTransferDestination(transferDestination *TwilioColdBlindTransferConfigInputTransferDestination) {
	t.TransferDestination = transferDestination
}

func (t *TwilioColdBlindTransferConfigInput) SetPostDialDigits(postDialDigits *TwilioColdBlindTransferConfigInputPostDialDigits) {
	t.PostDialDigits = postDialDigits
}

func (t *TwilioColdBlindTransferConfigInput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TwilioColdBlindTransferConfigInput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TwilioColdBlindTransferConfigInput to string"
	}
	return string(jsonData)
}

func (t *TwilioColdBlindTransferConfigInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TwilioColdBlindTransferConfigInput) MarshalJSON() ([]byte, error) {
	type embed TwilioColdBlindTransferConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
