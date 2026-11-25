package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Data for concurrency wait-queue status events.
type QueueStatusQueueStatusEvent struct {
	Status          QueueStatusQueueStatusEventStatus `json:"status" xml:"status" required:"true"`
	extraProperties map[string]interface{}
}

func (q *QueueStatusQueueStatusEvent) GetStatus() (value QueueStatusQueueStatusEventStatus) {
	if q == nil {
		return
	}
	return q.Status
}

func (q *QueueStatusQueueStatusEvent) SetStatus(status QueueStatusQueueStatusEventStatus) {
	q.Status = status
}

func (q *QueueStatusQueueStatusEvent) GetExtraProperties() map[string]interface{} {
	if q == nil {
		return nil
	}
	return q.extraProperties
}

func (q QueueStatusQueueStatusEvent) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: QueueStatusQueueStatusEvent to string"
	}
	return string(jsonData)
}

func (q *QueueStatusQueueStatusEvent) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, q); err != nil {
		return err
	}
	type alias QueueStatusQueueStatusEvent
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*q = QueueStatusQueueStatusEvent(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *q)
	if err != nil {
		return err
	}
	q.extraProperties = extraProperties
	return nil
}

func (q *QueueStatusQueueStatusEvent) MarshalJSON() ([]byte, error) {
	type embed QueueStatusQueueStatusEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*q),
	}
	return json.Marshal(marshaler)
}
