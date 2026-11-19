package types

import (
	"fmt"
)

type StreamAudioIsolationRequestFileFormat string

const (
	StreamAudioIsolationRequestFileFormatPcmS16Le16 StreamAudioIsolationRequestFileFormat = "pcm_s16le_16"
	StreamAudioIsolationRequestFileFormatOther      StreamAudioIsolationRequestFileFormat = "other"
)

func NewStreamAudioIsolationRequestFileFormatFromString(s string) (StreamAudioIsolationRequestFileFormat, error) {
	switch s {
	case "pcm_s16le_16":
		return StreamAudioIsolationRequestFileFormatPcmS16Le16, nil
	case "other":
		return StreamAudioIsolationRequestFileFormatOther, nil
	}
	var t StreamAudioIsolationRequestFileFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e StreamAudioIsolationRequestFileFormat) Ptr() *StreamAudioIsolationRequestFileFormat {
	return &e
}
