package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SearchHighlightSegment struct {
	Value           string `json:"value" xml:"value" required:"true"`
	IsHit           bool   `json:"is_hit" xml:"is_hit" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SearchHighlightSegment) GetValue() (value string) {
	if s == nil {
		return
	}
	return s.Value
}

func (s *SearchHighlightSegment) GetIsHit() (value bool) {
	if s == nil {
		return
	}
	return s.IsHit
}

func (s *SearchHighlightSegment) SetValue(value string) {
	s.Value = value
}

func (s *SearchHighlightSegment) SetIsHit(isHit bool) {
	s.IsHit = isHit
}

func (s *SearchHighlightSegment) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SearchHighlightSegment) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SearchHighlightSegment to string"
	}
	return string(jsonData)
}

func (s *SearchHighlightSegment) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SearchHighlightSegment
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SearchHighlightSegment(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SearchHighlightSegment) MarshalJSON() ([]byte, error) {
	type embed SearchHighlightSegment
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
