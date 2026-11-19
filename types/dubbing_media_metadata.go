package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingMediaMetadata struct {
	// The content type of the media.
	ContentType string `json:"content_type" xml:"content_type" required:"true"`
	// The duration of the media in seconds.
	Duration        float64 `json:"duration" xml:"duration" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DubbingMediaMetadata) GetContentType() (value string) {
	if d == nil {
		return
	}
	return d.ContentType
}

func (d *DubbingMediaMetadata) GetDuration() (value float64) {
	if d == nil {
		return
	}
	return d.Duration
}

func (d *DubbingMediaMetadata) SetContentType(contentType string) {
	d.ContentType = contentType
}

func (d *DubbingMediaMetadata) SetDuration(duration float64) {
	d.Duration = duration
}

func (d *DubbingMediaMetadata) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingMediaMetadata) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingMediaMetadata to string"
	}
	return string(jsonData)
}

func (d *DubbingMediaMetadata) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingMediaMetadata
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingMediaMetadata(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingMediaMetadata) MarshalJSON() ([]byte, error) {
	type embed DubbingMediaMetadata
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
