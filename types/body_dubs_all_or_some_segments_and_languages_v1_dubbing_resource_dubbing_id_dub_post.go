package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyDubsAllOrSomeSegmentsAndLanguagesV1DubbingResourceDubbingIdDubPost struct {
	// Dub only this list of segments.
	Segments []string `json:"segments" xml:"segments" required:"true"`
	// Dub only these languages for each segment.
	Languages       []string `json:"languages" xml:"languages" required:"true"`
	extraProperties map[string]interface{}
}

func (b *BodyDubsAllOrSomeSegmentsAndLanguagesV1DubbingResourceDubbingIdDubPost) GetSegments() (value []string) {
	if b == nil {
		return
	}
	return b.Segments
}

func (b *BodyDubsAllOrSomeSegmentsAndLanguagesV1DubbingResourceDubbingIdDubPost) GetLanguages() (value []string) {
	if b == nil {
		return
	}
	return b.Languages
}

func (b *BodyDubsAllOrSomeSegmentsAndLanguagesV1DubbingResourceDubbingIdDubPost) SetSegments(segments []string) {
	b.Segments = segments
}

func (b *BodyDubsAllOrSomeSegmentsAndLanguagesV1DubbingResourceDubbingIdDubPost) SetLanguages(languages []string) {
	b.Languages = languages
}

func (b *BodyDubsAllOrSomeSegmentsAndLanguagesV1DubbingResourceDubbingIdDubPost) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BodyDubsAllOrSomeSegmentsAndLanguagesV1DubbingResourceDubbingIdDubPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyDubsAllOrSomeSegmentsAndLanguagesV1DubbingResourceDubbingIdDubPost to string"
	}
	return string(jsonData)
}

func (b *BodyDubsAllOrSomeSegmentsAndLanguagesV1DubbingResourceDubbingIdDubPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, b); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BodyDubsAllOrSomeSegmentsAndLanguagesV1DubbingResourceDubbingIdDubPost) MarshalJSON() ([]byte, error) {
	type embed BodyDubsAllOrSomeSegmentsAndLanguagesV1DubbingResourceDubbingIdDubPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
