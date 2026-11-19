package dubbing

import (
	"fmt"
)

// Format to return transcript in. For subtitles use either 'srt' or 'webvtt', and for a full transcript use 'json'. The 'json' format is not yet supported for Dubbing Studio.
type GetTranscriptsRequestFormatType string

const (
	GetTranscriptsRequestFormatTypeSrt    GetTranscriptsRequestFormatType = "srt"
	GetTranscriptsRequestFormatTypeWebvtt GetTranscriptsRequestFormatType = "webvtt"
	GetTranscriptsRequestFormatTypeJson   GetTranscriptsRequestFormatType = "json"
)

func NewGetTranscriptsRequestFormatTypeFromString(s string) (GetTranscriptsRequestFormatType, error) {
	switch s {
	case "srt":
		return GetTranscriptsRequestFormatTypeSrt, nil
	case "webvtt":
		return GetTranscriptsRequestFormatTypeWebvtt, nil
	case "json":
		return GetTranscriptsRequestFormatTypeJson, nil
	}
	var t GetTranscriptsRequestFormatType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GetTranscriptsRequestFormatType) Ptr() *GetTranscriptsRequestFormatType {
	return &e
}
