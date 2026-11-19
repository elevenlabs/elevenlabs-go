package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbedSegment struct {
	StartTime       float64                 `json:"start_time" xml:"start_time" required:"true"`
	EndTime         float64                 `json:"end_time" xml:"end_time" required:"true"`
	Text            *string                 `json:"text" xml:"text" required:"true"`
	Subtitles       []*SegmentSubtitleFrame `json:"subtitles" xml:"subtitles" required:"true"`
	AudioStale      bool                    `json:"audio_stale" xml:"audio_stale" required:"true"`
	MediaRef        *DubbingMediaReference  `json:"media_ref" xml:"media_ref" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DubbedSegment) GetStartTime() (value float64) {
	if d == nil {
		return
	}
	return d.StartTime
}

func (d *DubbedSegment) GetEndTime() (value float64) {
	if d == nil {
		return
	}
	return d.EndTime
}

func (d *DubbedSegment) GetText() (value *string) {
	if d == nil {
		return
	}
	return d.Text
}

func (d *DubbedSegment) GetSubtitles() (value []*SegmentSubtitleFrame) {
	if d == nil {
		return
	}
	return d.Subtitles
}

func (d *DubbedSegment) GetAudioStale() (value bool) {
	if d == nil {
		return
	}
	return d.AudioStale
}

func (d *DubbedSegment) GetMediaRef() (value *DubbingMediaReference) {
	if d == nil {
		return
	}
	return d.MediaRef
}

func (d *DubbedSegment) SetStartTime(startTime float64) {
	d.StartTime = startTime
}

func (d *DubbedSegment) SetEndTime(endTime float64) {
	d.EndTime = endTime
}

func (d *DubbedSegment) SetText(text *string) {
	d.Text = text
}

func (d *DubbedSegment) SetSubtitles(subtitles []*SegmentSubtitleFrame) {
	d.Subtitles = subtitles
}

func (d *DubbedSegment) SetAudioStale(audioStale bool) {
	d.AudioStale = audioStale
}

func (d *DubbedSegment) SetMediaRef(mediaRef *DubbingMediaReference) {
	d.MediaRef = mediaRef
}

func (d *DubbedSegment) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbedSegment) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbedSegment to string"
	}
	return string(jsonData)
}

func (d *DubbedSegment) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, d); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbedSegment) MarshalJSON() ([]byte, error) {
	type embed DubbedSegment
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
