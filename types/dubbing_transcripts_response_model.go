package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingTranscriptsResponseModel struct {
	TranscriptFormat DubbingTranscriptsResponseModelTranscriptFormat `json:"transcript_format" xml:"transcript_format" required:"true"`
	Srt              *string                                         `json:"srt,omitempty" xml:"srt,omitempty"`
	Webvtt           *string                                         `json:"webvtt,omitempty" xml:"webvtt,omitempty"`
	Json             *DubbingTranscript                              `json:"json,omitempty" xml:"json,omitempty"`
	extraProperties  map[string]interface{}
}

func (d *DubbingTranscriptsResponseModel) GetTranscriptFormat() (value DubbingTranscriptsResponseModelTranscriptFormat) {
	if d == nil {
		return
	}
	return d.TranscriptFormat
}

func (d *DubbingTranscriptsResponseModel) GetSrt() (value *string) {
	if d == nil {
		return
	}
	return d.Srt
}

func (d *DubbingTranscriptsResponseModel) GetWebvtt() (value *string) {
	if d == nil {
		return
	}
	return d.Webvtt
}

func (d *DubbingTranscriptsResponseModel) GetJson() (value *DubbingTranscript) {
	if d == nil {
		return
	}
	return d.Json
}

func (d *DubbingTranscriptsResponseModel) SetTranscriptFormat(transcriptFormat DubbingTranscriptsResponseModelTranscriptFormat) {
	d.TranscriptFormat = transcriptFormat
}

func (d *DubbingTranscriptsResponseModel) SetSrt(srt *string) {
	d.Srt = srt
}

func (d *DubbingTranscriptsResponseModel) SetWebvtt(webvtt *string) {
	d.Webvtt = webvtt
}

func (d *DubbingTranscriptsResponseModel) SetJson(json *DubbingTranscript) {
	d.Json = json
}

func (d *DubbingTranscriptsResponseModel) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingTranscriptsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingTranscriptsResponseModel to string"
	}
	return string(jsonData)
}

func (d *DubbingTranscriptsResponseModel) UnmarshalJSON(data []byte) error {
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

func (d *DubbingTranscriptsResponseModel) MarshalJSON() ([]byte, error) {
	type embed DubbingTranscriptsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
