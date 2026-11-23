package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type Interruption struct {
	InterruptionEvent *InterruptionInterruptionEvent `json:"interruption_event" xml:"interruption_event" required:"true"`
	Type              *string                        `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties   map[string]interface{}
}

func (i *Interruption) GetInterruptionEvent() (value *InterruptionInterruptionEvent) {
	if i == nil {
		return
	}
	return i.InterruptionEvent
}

func (i *Interruption) GetType() (value *string) {
	if i == nil {
		return
	}
	return i.Type
}

func (i *Interruption) SetInterruptionEvent(interruptionEvent *InterruptionInterruptionEvent) {
	i.InterruptionEvent = interruptionEvent
}

func (i *Interruption) SetType(type_ *string) {
	i.Type = type_
}

func (i *Interruption) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i Interruption) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: Interruption to string"
	}
	return string(jsonData)
}

func (i *Interruption) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, i); err != nil {
		return err
	}
	type alias Interruption
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*i = Interruption(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *Interruption) MarshalJSON() ([]byte, error) {
	type embed Interruption
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}
