package types

import (
	"fmt"
)

type BodySpeechToSpeechV1SpeechToSpeechVoiceIdPostFileFormat string

const (
	BodySpeechToSpeechV1SpeechToSpeechVoiceIdPostFileFormatPcmS16Le16 BodySpeechToSpeechV1SpeechToSpeechVoiceIdPostFileFormat = "pcm_s16le_16"
	BodySpeechToSpeechV1SpeechToSpeechVoiceIdPostFileFormatOther      BodySpeechToSpeechV1SpeechToSpeechVoiceIdPostFileFormat = "other"
)

func NewBodySpeechToSpeechV1SpeechToSpeechVoiceIdPostFileFormatFromString(s string) (BodySpeechToSpeechV1SpeechToSpeechVoiceIdPostFileFormat, error) {
	switch s {
	case "pcm_s16le_16":
		return BodySpeechToSpeechV1SpeechToSpeechVoiceIdPostFileFormatPcmS16Le16, nil
	case "other":
		return BodySpeechToSpeechV1SpeechToSpeechVoiceIdPostFileFormatOther, nil
	}
	var t BodySpeechToSpeechV1SpeechToSpeechVoiceIdPostFileFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodySpeechToSpeechV1SpeechToSpeechVoiceIdPostFileFormat) Ptr() *BodySpeechToSpeechV1SpeechToSpeechVoiceIdPostFileFormat {
	return &e
}
