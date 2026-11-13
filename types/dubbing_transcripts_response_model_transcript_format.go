package types

import (
	"fmt"
)

type DubbingTranscriptsResponseModelTranscriptFormat string

const (
	DubbingTranscriptsResponseModelTranscriptFormatSrt    DubbingTranscriptsResponseModelTranscriptFormat = "srt"
	DubbingTranscriptsResponseModelTranscriptFormatWebvtt DubbingTranscriptsResponseModelTranscriptFormat = "webvtt"
	DubbingTranscriptsResponseModelTranscriptFormatJson   DubbingTranscriptsResponseModelTranscriptFormat = "json"
)

func NewDubbingTranscriptsResponseModelTranscriptFormatFromString(s string) (DubbingTranscriptsResponseModelTranscriptFormat, error) {
	switch s {
	case "srt":
		return DubbingTranscriptsResponseModelTranscriptFormatSrt, nil
	case "webvtt":
		return DubbingTranscriptsResponseModelTranscriptFormatWebvtt, nil
	case "json":
		return DubbingTranscriptsResponseModelTranscriptFormatJson, nil
	}
	var t DubbingTranscriptsResponseModelTranscriptFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e DubbingTranscriptsResponseModelTranscriptFormat) Ptr() *DubbingTranscriptsResponseModelTranscriptFormat {
	return &e
}
