package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type Feedback struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// ID of the agent response event being rated.
	EventId int `json:"event_id" xml:"event_id" required:"true"`
	// Feedback score for the referenced agent response. Null clears any previously submitted feedback for the event.
	Score           *FeedbackScore `json:"score,omitempty" xml:"score,omitempty"`
	extraProperties map[string]interface{}
}

func (f *Feedback) GetType() (value *string) {
	if f == nil {
		return
	}
	return f.Type
}

func (f *Feedback) GetEventId() (value int) {
	if f == nil {
		return
	}
	return f.EventId
}

func (f *Feedback) GetScore() (value *FeedbackScore) {
	if f == nil {
		return
	}
	return f.Score
}

func (f *Feedback) SetType(type_ *string) {
	f.Type = type_
}

func (f *Feedback) SetEventId(eventId int) {
	f.EventId = eventId
}

func (f *Feedback) SetScore(score *FeedbackScore) {
	f.Score = score
}

func (f *Feedback) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f Feedback) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: Feedback to string"
	}
	return string(jsonData)
}

func (f *Feedback) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, f); err != nil {
		return err
	}
	type alias Feedback
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*f = Feedback(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *Feedback) MarshalJSON() ([]byte, error) {
	type embed Feedback
	var marshaler = struct {
		embed
	}{
		embed: embed(*f),
	}
	return json.Marshal(marshaler)
}
