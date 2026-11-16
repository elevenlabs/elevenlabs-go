package types

import (
	"fmt"
)

type BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPostFileFormat string

const (
	BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPostFileFormatPcmS16Le16 BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPostFileFormat = "pcm_s16le_16"
	BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPostFileFormatOther      BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPostFileFormat = "other"
)

func NewBodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPostFileFormatFromString(s string) (BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPostFileFormat, error) {
	switch s {
	case "pcm_s16le_16":
		return BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPostFileFormatPcmS16Le16, nil
	case "other":
		return BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPostFileFormatOther, nil
	}
	var t BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPostFileFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPostFileFormat) Ptr() *BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPostFileFormat {
	return &e
}
