package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Metadata about the project's source media.
type DubbingSourceMediaInfo struct {
	// Original filename of the uploaded source media (null for URL sources).
	Filename *string `json:"filename,omitempty" xml:"filename,omitempty"`
	// Duration of the source media, in seconds.
	DurationS *float64 `json:"duration_s,omitempty" xml:"duration_s,omitempty"`
	// Whether the source media contains a video stream.
	HasVideo *bool `json:"has_video,omitempty" xml:"has_video,omitempty"`
	// MIME type of the uploaded source media (null for URL sources).
	MimeType        *string `json:"mime_type,omitempty" xml:"mime_type,omitempty"`
	extraProperties map[string]interface{}
}

func (d *DubbingSourceMediaInfo) GetFilename() (value *string) {
	if d == nil {
		return
	}
	return d.Filename
}

func (d *DubbingSourceMediaInfo) GetDurationS() (value *float64) {
	if d == nil {
		return
	}
	return d.DurationS
}

func (d *DubbingSourceMediaInfo) GetHasVideo() (value *bool) {
	if d == nil {
		return
	}
	return d.HasVideo
}

func (d *DubbingSourceMediaInfo) GetMimeType() (value *string) {
	if d == nil {
		return
	}
	return d.MimeType
}

func (d *DubbingSourceMediaInfo) SetFilename(filename *string) {
	d.Filename = filename
}

func (d *DubbingSourceMediaInfo) SetDurationS(durationS *float64) {
	d.DurationS = durationS
}

func (d *DubbingSourceMediaInfo) SetHasVideo(hasVideo *bool) {
	d.HasVideo = hasVideo
}

func (d *DubbingSourceMediaInfo) SetMimeType(mimeType *string) {
	d.MimeType = mimeType
}

func (d *DubbingSourceMediaInfo) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingSourceMediaInfo) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingSourceMediaInfo to string"
	}
	return string(jsonData)
}

func (d *DubbingSourceMediaInfo) UnmarshalJSON(data []byte) error {
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

func (d *DubbingSourceMediaInfo) MarshalJSON() ([]byte, error) {
	type embed DubbingSourceMediaInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
