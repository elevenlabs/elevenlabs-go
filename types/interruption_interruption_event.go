package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type InterruptionInterruptionEvent struct {
	EventId         int `json:"event_id" xml:"event_id" required:"true"`
	extraProperties map[string]interface{}
}

func (i *InterruptionInterruptionEvent) GetEventId() (value int) {
	if i == nil {
		return
	}
	return i.EventId
}

func (i *InterruptionInterruptionEvent) SetEventId(eventId int) {
	i.EventId = eventId
}

func (i *InterruptionInterruptionEvent) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i InterruptionInterruptionEvent) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InterruptionInterruptionEvent to string"
	}
	return string(jsonData)
}

func (i *InterruptionInterruptionEvent) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, i); err != nil {
		return err
	}
	type alias InterruptionInterruptionEvent
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*i = InterruptionInterruptionEvent(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *InterruptionInterruptionEvent) MarshalJSON() ([]byte, error) {
	type embed InterruptionInterruptionEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}
