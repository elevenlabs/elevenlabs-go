package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Status of a caller held in the concurrency wait queue. Sent once when the wait
// starts (`waiting`) and once when it ends: `admitted` when the caller is connected
// to the agent, or `timed_out` right before the connection is closed with code 4300.
// Hold audio arrives as regular `audio` events while the caller waits.
type QueueStatus struct {
	// Data for concurrency wait-queue status events.
	QueueStatusEvent *QueueStatusQueueStatusEvent `json:"queue_status_event" xml:"queue_status_event" required:"true"`
	Type             *string                      `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties  map[string]interface{}
}

func (q *QueueStatus) GetQueueStatusEvent() (value *QueueStatusQueueStatusEvent) {
	if q == nil {
		return
	}
	return q.QueueStatusEvent
}

func (q *QueueStatus) GetType() (value *string) {
	if q == nil {
		return
	}
	return q.Type
}

func (q *QueueStatus) SetQueueStatusEvent(queueStatusEvent *QueueStatusQueueStatusEvent) {
	q.QueueStatusEvent = queueStatusEvent
}

func (q *QueueStatus) SetType(type_ *string) {
	q.Type = type_
}

func (q *QueueStatus) GetExtraProperties() map[string]interface{} {
	if q == nil {
		return nil
	}
	return q.extraProperties
}

func (q QueueStatus) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: QueueStatus to string"
	}
	return string(jsonData)
}

func (q *QueueStatus) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, q); err != nil {
		return err
	}
	type alias QueueStatus
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*q = QueueStatus(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *q)
	if err != nil {
		return err
	}
	q.extraProperties = extraProperties
	return nil
}

func (q *QueueStatus) MarshalJSON() ([]byte, error) {
	type embed QueueStatus
	var marshaler = struct {
		embed
	}{
		embed: embed(*q),
	}
	return json.Marshal(marshaler)
}
