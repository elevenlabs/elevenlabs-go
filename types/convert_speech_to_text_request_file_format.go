package types

import (
	"fmt"
)

// The format of input audio. Options are 'pcm_s16le_16' or 'other' For `pcm_s16le_16`, the input audio must be 16-bit PCM at a 16kHz sample rate, single channel (mono), and little-endian byte order. Latency will be lower than with passing an encoded waveform.
type ConvertSpeechToTextRequestFileFormat string

const (
	ConvertSpeechToTextRequestFileFormatPcmS16Le16 ConvertSpeechToTextRequestFileFormat = "pcm_s16le_16"
	ConvertSpeechToTextRequestFileFormatOther      ConvertSpeechToTextRequestFileFormat = "other"
)

func NewConvertSpeechToTextRequestFileFormatFromString(s string) (ConvertSpeechToTextRequestFileFormat, error) {
	switch s {
	case "pcm_s16le_16":
		return ConvertSpeechToTextRequestFileFormatPcmS16Le16, nil
	case "other":
		return ConvertSpeechToTextRequestFileFormatOther, nil
	}
	var t ConvertSpeechToTextRequestFileFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConvertSpeechToTextRequestFileFormat) Ptr() *ConvertSpeechToTextRequestFileFormat {
	return &e
}
