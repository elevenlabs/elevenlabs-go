package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type MemoryEntrySearchResult struct {
	EntryId         string                         `json:"entry_id" xml:"entry_id" required:"true"`
	Version         int                            `json:"version" xml:"version" required:"true"`
	Summary         *string                        `json:"summary,omitempty" xml:"summary,omitempty"`
	Text            *string                        `json:"text,omitempty" xml:"text,omitempty"`
	Source          *MemoryEntrySearchResultSource `json:"source,omitempty" xml:"source,omitempty"`
	extraProperties map[string]interface{}
}

func (m *MemoryEntrySearchResult) GetEntryId() (value string) {
	if m == nil {
		return
	}
	return m.EntryId
}

func (m *MemoryEntrySearchResult) GetVersion() (value int) {
	if m == nil {
		return
	}
	return m.Version
}

func (m *MemoryEntrySearchResult) GetSummary() (value *string) {
	if m == nil {
		return
	}
	return m.Summary
}

func (m *MemoryEntrySearchResult) GetText() (value *string) {
	if m == nil {
		return
	}
	return m.Text
}

func (m *MemoryEntrySearchResult) GetSource() (value *MemoryEntrySearchResultSource) {
	if m == nil {
		return
	}
	return m.Source
}

func (m *MemoryEntrySearchResult) SetEntryId(entryId string) {
	m.EntryId = entryId
}

func (m *MemoryEntrySearchResult) SetVersion(version int) {
	m.Version = version
}

func (m *MemoryEntrySearchResult) SetSummary(summary *string) {
	m.Summary = summary
}

func (m *MemoryEntrySearchResult) SetText(text *string) {
	m.Text = text
}

func (m *MemoryEntrySearchResult) SetSource(source *MemoryEntrySearchResultSource) {
	m.Source = source
}

func (m *MemoryEntrySearchResult) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MemoryEntrySearchResult) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MemoryEntrySearchResult to string"
	}
	return string(jsonData)
}

func (m *MemoryEntrySearchResult) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, m); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MemoryEntrySearchResult) MarshalJSON() ([]byte, error) {
	type embed MemoryEntrySearchResult
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
