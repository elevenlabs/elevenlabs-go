package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DuplicateSpeechEngineRequest struct {
	// The speech engine ID (accepts seng_ or agent_ prefix)
	SpeechEngineId string `json:"-" url:"-"`
	// Name for the duplicated speech engine. Defaults to the source name.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (d *DuplicateSpeechEngineRequest) GetSpeechEngineId() (value string) {
	if d == nil {
		return
	}
	return d.SpeechEngineId
}

func (d *DuplicateSpeechEngineRequest) GetName() (value *string) {
	if d == nil {
		return
	}
	return d.Name
}

func (d *DuplicateSpeechEngineRequest) SetSpeechEngineId(speechEngineId string) {
	d.SpeechEngineId = speechEngineId
}

func (d *DuplicateSpeechEngineRequest) SetName(name *string) {
	d.Name = name
}

func (d *DuplicateSpeechEngineRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (d DuplicateSpeechEngineRequest) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DuplicateSpeechEngineRequest to string"
	}
	return string(jsonData)
}

func (d *DuplicateSpeechEngineRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, d)
}

func (d *DuplicateSpeechEngineRequest) MarshalJSON() ([]byte, error) {
	type embed DuplicateSpeechEngineRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
