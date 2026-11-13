package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TicketSearchMatchResponseModel struct {
	Field TicketSearchMatchField `json:"field" xml:"field" required:"true"`
	// Whitespace-collapsed excerpt around the match, with an ellipsis where it was cut.
	Snippet string `json:"snippet" xml:"snippet" required:"true"`
	// Offset of the match in `snippet`.
	HighlightStart int `json:"highlight_start" xml:"highlight_start" required:"true"`
	// Exclusive end offset of the match in `snippet`.
	HighlightEnd int `json:"highlight_end" xml:"highlight_end" required:"true"`
	// The commented turn, set when `field` is 'turn_comment'.
	TurnIndex       *int `json:"turn_index,omitempty" xml:"turn_index,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TicketSearchMatchResponseModel) GetField() (value TicketSearchMatchField) {
	if t == nil {
		return
	}
	return t.Field
}

func (t *TicketSearchMatchResponseModel) GetSnippet() (value string) {
	if t == nil {
		return
	}
	return t.Snippet
}

func (t *TicketSearchMatchResponseModel) GetHighlightStart() (value int) {
	if t == nil {
		return
	}
	return t.HighlightStart
}

func (t *TicketSearchMatchResponseModel) GetHighlightEnd() (value int) {
	if t == nil {
		return
	}
	return t.HighlightEnd
}

func (t *TicketSearchMatchResponseModel) GetTurnIndex() (value *int) {
	if t == nil {
		return
	}
	return t.TurnIndex
}

func (t *TicketSearchMatchResponseModel) SetField(field TicketSearchMatchField) {
	t.Field = field
}

func (t *TicketSearchMatchResponseModel) SetSnippet(snippet string) {
	t.Snippet = snippet
}

func (t *TicketSearchMatchResponseModel) SetHighlightStart(highlightStart int) {
	t.HighlightStart = highlightStart
}

func (t *TicketSearchMatchResponseModel) SetHighlightEnd(highlightEnd int) {
	t.HighlightEnd = highlightEnd
}

func (t *TicketSearchMatchResponseModel) SetTurnIndex(turnIndex *int) {
	t.TurnIndex = turnIndex
}

func (t *TicketSearchMatchResponseModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TicketSearchMatchResponseModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TicketSearchMatchResponseModel to string"
	}
	return string(jsonData)
}

func (t *TicketSearchMatchResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TicketSearchMatchResponseModel) MarshalJSON() ([]byte, error) {
	type embed TicketSearchMatchResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
