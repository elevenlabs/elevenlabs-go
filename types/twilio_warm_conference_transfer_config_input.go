package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TwilioWarmConferenceTransferConfigInput struct {
	TransferDestination *TwilioWarmConferenceTransferConfigInputTransferDestination `json:"transfer_destination" xml:"transfer_destination" required:"true"`
	// DTMF digits to send after the call connects (e.g. 'ww1234' for an extension). Static or a dynamic variable reference. Use 'w' for a 0.5s pause.
	PostDialDigits  *TwilioWarmConferenceTransferConfigInputPostDialDigits `json:"post_dial_digits,omitempty" xml:"post_dial_digits,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TwilioWarmConferenceTransferConfigInput) GetTransferDestination() (value *TwilioWarmConferenceTransferConfigInputTransferDestination) {
	if t == nil {
		return
	}
	return t.TransferDestination
}

func (t *TwilioWarmConferenceTransferConfigInput) GetPostDialDigits() (value *TwilioWarmConferenceTransferConfigInputPostDialDigits) {
	if t == nil {
		return
	}
	return t.PostDialDigits
}

func (t *TwilioWarmConferenceTransferConfigInput) SetTransferDestination(transferDestination *TwilioWarmConferenceTransferConfigInputTransferDestination) {
	t.TransferDestination = transferDestination
}

func (t *TwilioWarmConferenceTransferConfigInput) SetPostDialDigits(postDialDigits *TwilioWarmConferenceTransferConfigInputPostDialDigits) {
	t.PostDialDigits = postDialDigits
}

func (t *TwilioWarmConferenceTransferConfigInput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TwilioWarmConferenceTransferConfigInput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TwilioWarmConferenceTransferConfigInput to string"
	}
	return string(jsonData)
}

func (t *TwilioWarmConferenceTransferConfigInput) UnmarshalJSON(data []byte) error {
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

func (t *TwilioWarmConferenceTransferConfigInput) MarshalJSON() ([]byte, error) {
	type embed TwilioWarmConferenceTransferConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
