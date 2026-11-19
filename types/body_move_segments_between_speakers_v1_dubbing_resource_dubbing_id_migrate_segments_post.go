package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyMoveSegmentsBetweenSpeakersV1DubbingResourceDubbingIdMigrateSegmentsPost struct {
	SegmentIds      []string `json:"segment_ids" xml:"segment_ids" required:"true"`
	SpeakerId       string   `json:"speaker_id" xml:"speaker_id" required:"true"`
	extraProperties map[string]interface{}
}

func (b *BodyMoveSegmentsBetweenSpeakersV1DubbingResourceDubbingIdMigrateSegmentsPost) GetSegmentIds() (value []string) {
	if b == nil {
		return
	}
	return b.SegmentIds
}

func (b *BodyMoveSegmentsBetweenSpeakersV1DubbingResourceDubbingIdMigrateSegmentsPost) GetSpeakerId() (value string) {
	if b == nil {
		return
	}
	return b.SpeakerId
}

func (b *BodyMoveSegmentsBetweenSpeakersV1DubbingResourceDubbingIdMigrateSegmentsPost) SetSegmentIds(segmentIds []string) {
	b.SegmentIds = segmentIds
}

func (b *BodyMoveSegmentsBetweenSpeakersV1DubbingResourceDubbingIdMigrateSegmentsPost) SetSpeakerId(speakerId string) {
	b.SpeakerId = speakerId
}

func (b *BodyMoveSegmentsBetweenSpeakersV1DubbingResourceDubbingIdMigrateSegmentsPost) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BodyMoveSegmentsBetweenSpeakersV1DubbingResourceDubbingIdMigrateSegmentsPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyMoveSegmentsBetweenSpeakersV1DubbingResourceDubbingIdMigrateSegmentsPost to string"
	}
	return string(jsonData)
}

func (b *BodyMoveSegmentsBetweenSpeakersV1DubbingResourceDubbingIdMigrateSegmentsPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyMoveSegmentsBetweenSpeakersV1DubbingResourceDubbingIdMigrateSegmentsPost
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyMoveSegmentsBetweenSpeakersV1DubbingResourceDubbingIdMigrateSegmentsPost(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BodyMoveSegmentsBetweenSpeakersV1DubbingResourceDubbingIdMigrateSegmentsPost) MarshalJSON() ([]byte, error) {
	type embed BodyMoveSegmentsBetweenSpeakersV1DubbingResourceDubbingIdMigrateSegmentsPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
