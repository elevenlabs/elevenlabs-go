package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type MusicFinetunePageResponseModel struct {
	// The finetunes in this page.
	Finetunes []*MusicFinetuneResponseModel `json:"finetunes" xml:"finetunes" required:"true"`
	// Cursor to pass as `cursor` to fetch the next page; `null` when there are no more results.
	NextCursor *string `json:"next_cursor" xml:"next_cursor" required:"true"`
	// Whether more finetunes are available beyond this page.
	HasMore         bool `json:"has_more" xml:"has_more" required:"true"`
	extraProperties map[string]interface{}
}

func (m *MusicFinetunePageResponseModel) GetFinetunes() (value []*MusicFinetuneResponseModel) {
	if m == nil {
		return
	}
	return m.Finetunes
}

func (m *MusicFinetunePageResponseModel) GetNextCursor() (value *string) {
	if m == nil {
		return
	}
	return m.NextCursor
}

func (m *MusicFinetunePageResponseModel) GetHasMore() (value bool) {
	if m == nil {
		return
	}
	return m.HasMore
}

func (m *MusicFinetunePageResponseModel) SetFinetunes(finetunes []*MusicFinetuneResponseModel) {
	m.Finetunes = finetunes
}

func (m *MusicFinetunePageResponseModel) SetNextCursor(nextCursor *string) {
	m.NextCursor = nextCursor
}

func (m *MusicFinetunePageResponseModel) SetHasMore(hasMore bool) {
	m.HasMore = hasMore
}

func (m *MusicFinetunePageResponseModel) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MusicFinetunePageResponseModel) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MusicFinetunePageResponseModel to string"
	}
	return string(jsonData)
}

func (m *MusicFinetunePageResponseModel) UnmarshalJSON(data []byte) error {
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

func (m *MusicFinetunePageResponseModel) MarshalJSON() ([]byte, error) {
	type embed MusicFinetunePageResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
