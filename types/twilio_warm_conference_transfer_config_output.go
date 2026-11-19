package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TwilioWarmConferenceTransferConfigOutput struct {
	TransferDestination *TwilioWarmConferenceTransferConfigOutputTransferDestination `json:"transfer_destination" xml:"transfer_destination" required:"true"`
	// DTMF digits to send after the call connects (e.g. 'ww1234' for an extension). Static or a dynamic variable reference. Use 'w' for a 0.5s pause.
	PostDialDigits  *TwilioWarmConferenceTransferConfigOutputPostDialDigits `json:"post_dial_digits" xml:"post_dial_digits" required:"true"`
	type_           string
	extraProperties map[string]interface{}
}

func (t *TwilioWarmConferenceTransferConfigOutput) GetTransferDestination() (value *TwilioWarmConferenceTransferConfigOutputTransferDestination) {
	if t == nil {
		return
	}
	return t.TransferDestination
}

func (t *TwilioWarmConferenceTransferConfigOutput) GetPostDialDigits() (value *TwilioWarmConferenceTransferConfigOutputPostDialDigits) {
	if t == nil {
		return
	}
	return t.PostDialDigits
}

func (t *TwilioWarmConferenceTransferConfigOutput) SetTransferDestination(transferDestination *TwilioWarmConferenceTransferConfigOutputTransferDestination) {
	t.TransferDestination = transferDestination
}

func (t *TwilioWarmConferenceTransferConfigOutput) SetPostDialDigits(postDialDigits *TwilioWarmConferenceTransferConfigOutputPostDialDigits) {
	t.PostDialDigits = postDialDigits
}

func (t *TwilioWarmConferenceTransferConfigOutput) Type() string {
	return t.type_
}

func (t *TwilioWarmConferenceTransferConfigOutput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TwilioWarmConferenceTransferConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TwilioWarmConferenceTransferConfigOutput to string"
	}
	return string(jsonData)
}

func (t *TwilioWarmConferenceTransferConfigOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "twilio_warm_conference_transfer" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "twilio_warm_conference_transfer", *literals.Type)
	}
	t.type_ = "twilio_warm_conference_transfer"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TwilioWarmConferenceTransferConfigOutput) MarshalJSON() ([]byte, error) {
	type embed TwilioWarmConferenceTransferConfigOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*t),
		Type:  "twilio_warm_conference_transfer",
	}
	return json.Marshal(marshaler)
}
