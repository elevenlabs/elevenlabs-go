package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type HistoryItemResponse struct {
	State           any `json:"state,omitempty" xml:"state,omitempty"`
	VoiceCategory   any `json:"voice_category,omitempty" xml:"voice_category,omitempty"`
	extraProperties map[string]interface{}
}

func (h *HistoryItemResponse) GetState() (value any) {
	if h == nil {
		return
	}
	return h.State
}

func (h *HistoryItemResponse) GetVoiceCategory() (value any) {
	if h == nil {
		return
	}
	return h.VoiceCategory
}

func (h *HistoryItemResponse) SetState(state any) {
	h.State = state
}

func (h *HistoryItemResponse) SetVoiceCategory(voiceCategory any) {
	h.VoiceCategory = voiceCategory
}

func (h *HistoryItemResponse) GetExtraProperties() map[string]interface{} {
	if h == nil {
		return nil
	}
	return h.extraProperties
}

func (h HistoryItemResponse) String() string {
	jsonData, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return "error converting struct: HistoryItemResponse to string"
	}
	return string(jsonData)
}

func (h *HistoryItemResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler HistoryItemResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*h = HistoryItemResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *h)
	if err != nil {
		return err
	}
	h.extraProperties = extraProperties
	return nil
}

func (h *HistoryItemResponse) MarshalJSON() ([]byte, error) {
	type embed HistoryItemResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*h),
	}
	return json.Marshal(marshaler)
}
