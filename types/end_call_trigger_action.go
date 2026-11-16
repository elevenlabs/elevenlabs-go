package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type EndCallTriggerAction struct {
	extraProperties map[string]interface{}
}

func (e *EndCallTriggerAction) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EndCallTriggerAction) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EndCallTriggerAction to string"
	}
	return string(jsonData)
}

func (e *EndCallTriggerAction) UnmarshalJSON(data []byte) error {
	type unmarshaler EndCallTriggerAction
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EndCallTriggerAction(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *e, "type")
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EndCallTriggerAction) MarshalJSON() ([]byte, error) {
	type embed EndCallTriggerAction
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
