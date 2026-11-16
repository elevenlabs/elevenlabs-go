package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetSpeechHistoryResponse struct {
	// A list of speech history items.
	History []*SpeechHistoryItemResponse `json:"history" xml:"history" required:"true"`
	// The ID of the last history item.
	LastHistoryItemId *string `json:"last_history_item_id,omitempty" xml:"last_history_item_id,omitempty"`
	// Whether there are more history items to fetch.
	HasMore bool `json:"has_more" xml:"has_more" required:"true"`
	// The timestamp of the last history item.
	ScannedUntil    *int `json:"scanned_until,omitempty" xml:"scanned_until,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GetSpeechHistoryResponse) GetHistory() (value []*SpeechHistoryItemResponse) {
	if g == nil {
		return
	}
	return g.History
}

func (g *GetSpeechHistoryResponse) GetLastHistoryItemId() (value *string) {
	if g == nil {
		return
	}
	return g.LastHistoryItemId
}

func (g *GetSpeechHistoryResponse) GetHasMore() (value bool) {
	if g == nil {
		return
	}
	return g.HasMore
}

func (g *GetSpeechHistoryResponse) GetScannedUntil() (value *int) {
	if g == nil {
		return
	}
	return g.ScannedUntil
}

func (g *GetSpeechHistoryResponse) SetHistory(history []*SpeechHistoryItemResponse) {
	g.History = history
}

func (g *GetSpeechHistoryResponse) SetLastHistoryItemId(lastHistoryItemId *string) {
	g.LastHistoryItemId = lastHistoryItemId
}

func (g *GetSpeechHistoryResponse) SetHasMore(hasMore bool) {
	g.HasMore = hasMore
}

func (g *GetSpeechHistoryResponse) SetScannedUntil(scannedUntil *int) {
	g.ScannedUntil = scannedUntil
}

func (g *GetSpeechHistoryResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetSpeechHistoryResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetSpeechHistoryResponse to string"
	}
	return string(jsonData)
}

func (g *GetSpeechHistoryResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetSpeechHistoryResponse) MarshalJSON() ([]byte, error) {
	type embed GetSpeechHistoryResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
