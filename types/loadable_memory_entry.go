package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type LoadableMemoryEntry struct {
	EntryId         string `json:"entry_id" xml:"entry_id" required:"true"`
	Summary         string `json:"summary" xml:"summary" required:"true"`
	Text            string `json:"text" xml:"text" required:"true"`
	extraProperties map[string]interface{}
}

func (l *LoadableMemoryEntry) GetEntryId() (value string) {
	if l == nil {
		return
	}
	return l.EntryId
}

func (l *LoadableMemoryEntry) GetSummary() (value string) {
	if l == nil {
		return
	}
	return l.Summary
}

func (l *LoadableMemoryEntry) GetText() (value string) {
	if l == nil {
		return
	}
	return l.Text
}

func (l *LoadableMemoryEntry) SetEntryId(entryId string) {
	l.EntryId = entryId
}

func (l *LoadableMemoryEntry) SetSummary(summary string) {
	l.Summary = summary
}

func (l *LoadableMemoryEntry) SetText(text string) {
	l.Text = text
}

func (l *LoadableMemoryEntry) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LoadableMemoryEntry) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LoadableMemoryEntry to string"
	}
	return string(jsonData)
}

func (l *LoadableMemoryEntry) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, l); err != nil {
		return err
	}
	type alias LoadableMemoryEntry
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*l = LoadableMemoryEntry(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LoadableMemoryEntry) MarshalJSON() ([]byte, error) {
	type embed LoadableMemoryEntry
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
