package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Which `campaign__` dynamic variables follow-up calls get.
type FollowupContextPolicy struct {
	// Pass the call type as `campaign__call_type`.
	CallType *bool `json:"call_type,omitempty" xml:"call_type,omitempty"`
	// Pass the last conversation's ID as `campaign__previous_conversation_id`.
	PreviousConversationId *bool `json:"previous_conversation_id,omitempty" xml:"previous_conversation_id,omitempty"`
	// Pass the last conversation's summary as `campaign__previous_summary`.
	PreviousSummary *bool `json:"previous_summary,omitempty" xml:"previous_summary,omitempty"`
	// Pass the last conversation's transcript as `campaign__previous_transcript`.
	PreviousTranscript *bool `json:"previous_transcript,omitempty" xml:"previous_transcript,omitempty"`
	extraProperties    map[string]interface{}
}

func (f *FollowupContextPolicy) GetCallType() (value *bool) {
	if f == nil {
		return
	}
	return f.CallType
}

func (f *FollowupContextPolicy) GetPreviousConversationId() (value *bool) {
	if f == nil {
		return
	}
	return f.PreviousConversationId
}

func (f *FollowupContextPolicy) GetPreviousSummary() (value *bool) {
	if f == nil {
		return
	}
	return f.PreviousSummary
}

func (f *FollowupContextPolicy) GetPreviousTranscript() (value *bool) {
	if f == nil {
		return
	}
	return f.PreviousTranscript
}

func (f *FollowupContextPolicy) SetCallType(callType *bool) {
	f.CallType = callType
}

func (f *FollowupContextPolicy) SetPreviousConversationId(previousConversationId *bool) {
	f.PreviousConversationId = previousConversationId
}

func (f *FollowupContextPolicy) SetPreviousSummary(previousSummary *bool) {
	f.PreviousSummary = previousSummary
}

func (f *FollowupContextPolicy) SetPreviousTranscript(previousTranscript *bool) {
	f.PreviousTranscript = previousTranscript
}

func (f *FollowupContextPolicy) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f FollowupContextPolicy) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FollowupContextPolicy to string"
	}
	return string(jsonData)
}

func (f *FollowupContextPolicy) UnmarshalJSON(data []byte) error {
	type unmarshaler FollowupContextPolicy
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = FollowupContextPolicy(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *FollowupContextPolicy) MarshalJSON() ([]byte, error) {
	type embed FollowupContextPolicy
	var marshaler = struct {
		embed
	}{
		embed: embed(*f),
	}
	return json.Marshal(marshaler)
}
