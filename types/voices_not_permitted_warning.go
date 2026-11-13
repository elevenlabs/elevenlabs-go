package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VoicesNotPermittedWarning struct {
	// Identifies this warning; branch on it to read the other fields.
	type_ string
	// Speakers whose voices were not permitted for cloning. The dub used a replacement voice for each of them; all other speakers are unaffected.
	SpeakerIds []string `json:"speaker_ids" xml:"speaker_ids" required:"true"`
	// Human-readable description of the warning, for display. The wording may change at any time, so we recommend branching on `type` instead.
	Message         string `json:"message" xml:"message" required:"true"`
	extraProperties map[string]interface{}
}

func (v *VoicesNotPermittedWarning) GetSpeakerIds() (value []string) {
	if v == nil {
		return
	}
	return v.SpeakerIds
}

func (v *VoicesNotPermittedWarning) GetMessage() (value string) {
	if v == nil {
		return
	}
	return v.Message
}

func (v *VoicesNotPermittedWarning) SetSpeakerIds(speakerIds []string) {
	v.SpeakerIds = speakerIds
}

func (v *VoicesNotPermittedWarning) SetMessage(message string) {
	v.Message = message
}

func (v *VoicesNotPermittedWarning) Type() string {
	return v.type_
}

func (v *VoicesNotPermittedWarning) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VoicesNotPermittedWarning) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VoicesNotPermittedWarning to string"
	}
	return string(jsonData)
}

func (v *VoicesNotPermittedWarning) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, v); err != nil {
		return err
	}
	type alias VoicesNotPermittedWarning
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*v = VoicesNotPermittedWarning(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "voices_not_permitted" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", v, "voices_not_permitted", *literals.Type)
	}
	v.type_ = "voices_not_permitted"
	extraProperties, err := internal.ExtractExtraProperties(data, *v, "type")
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VoicesNotPermittedWarning) MarshalJSON() ([]byte, error) {
	type embed VoicesNotPermittedWarning
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*v),
		Type:  "voices_not_permitted",
	}
	return json.Marshal(marshaler)
}
