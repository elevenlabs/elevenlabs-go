package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SectionSource struct {
	// The ID of the song to source the section from. You can find the song ID in the response headers when you generate a song.
	SongId string     `json:"song_id" xml:"song_id" required:"true" maxLength:"100" minLength:"1"`
	Range  *TimeRange `json:"range" xml:"range" required:"true"`
	// The ranges to exclude from the 'range'.
	NegativeRanges  []*TimeRange `json:"negative_ranges,omitempty" xml:"negative_ranges,omitempty" maxItems:"10"`
	extraProperties map[string]interface{}
}

func (s *SectionSource) GetSongId() (value string) {
	if s == nil {
		return
	}
	return s.SongId
}

func (s *SectionSource) GetRange() (value *TimeRange) {
	if s == nil {
		return
	}
	return s.Range
}

func (s *SectionSource) GetNegativeRanges() (value []*TimeRange) {
	if s == nil {
		return
	}
	return s.NegativeRanges
}

func (s *SectionSource) SetSongId(songId string) {
	s.SongId = songId
}

func (s *SectionSource) SetRange(range_ *TimeRange) {
	s.Range = range_
}

func (s *SectionSource) SetNegativeRanges(negativeRanges []*TimeRange) {
	s.NegativeRanges = negativeRanges
}

func (s *SectionSource) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SectionSource) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SectionSource to string"
	}
	return string(jsonData)
}

func (s *SectionSource) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SectionSource
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SectionSource(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SectionSource) MarshalJSON() ([]byte, error) {
	type embed SectionSource
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
