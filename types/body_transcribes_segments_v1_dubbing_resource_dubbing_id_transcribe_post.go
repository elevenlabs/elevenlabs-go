package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type BodyTranscribesSegmentsV1DubbingResourceDubbingIdTranscribePost struct {
	// Transcribe this specific list of segments.
	Segments        []string `json:"segments" xml:"segments" required:"true"`
	extraProperties map[string]interface{}
}

func (b *BodyTranscribesSegmentsV1DubbingResourceDubbingIdTranscribePost) GetSegments() (value []string) {
	if b == nil {
		return
	}
	return b.Segments
}

func (b *BodyTranscribesSegmentsV1DubbingResourceDubbingIdTranscribePost) SetSegments(segments []string) {
	b.Segments = segments
}

func (b *BodyTranscribesSegmentsV1DubbingResourceDubbingIdTranscribePost) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BodyTranscribesSegmentsV1DubbingResourceDubbingIdTranscribePost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyTranscribesSegmentsV1DubbingResourceDubbingIdTranscribePost to string"
	}
	return string(jsonData)
}

func (b *BodyTranscribesSegmentsV1DubbingResourceDubbingIdTranscribePost) UnmarshalJSON(data []byte) error {
	type unmarshaler BodyTranscribesSegmentsV1DubbingResourceDubbingIdTranscribePost
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BodyTranscribesSegmentsV1DubbingResourceDubbingIdTranscribePost(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BodyTranscribesSegmentsV1DubbingResourceDubbingIdTranscribePost) MarshalJSON() ([]byte, error) {
	type embed BodyTranscribesSegmentsV1DubbingResourceDubbingIdTranscribePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
