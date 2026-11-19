package types

import (
	"fmt"
)

type ConvertAudioIsolationRequestFileFormat string

const (
	ConvertAudioIsolationRequestFileFormatPcmS16Le16 ConvertAudioIsolationRequestFileFormat = "pcm_s16le_16"
	ConvertAudioIsolationRequestFileFormatOther      ConvertAudioIsolationRequestFileFormat = "other"
)

func NewConvertAudioIsolationRequestFileFormatFromString(s string) (ConvertAudioIsolationRequestFileFormat, error) {
	switch s {
	case "pcm_s16le_16":
		return ConvertAudioIsolationRequestFileFormatPcmS16Le16, nil
	case "other":
		return ConvertAudioIsolationRequestFileFormatOther, nil
	}
	var t ConvertAudioIsolationRequestFileFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConvertAudioIsolationRequestFileFormat) Ptr() *ConvertAudioIsolationRequestFileFormat {
	return &e
}
