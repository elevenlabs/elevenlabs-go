package types

import (
	"fmt"
)

// Audio encoding format for realtime translation input.
type TranslationRealtimeInputFormat string

const (
	TranslationRealtimeInputFormatPcm8000  TranslationRealtimeInputFormat = "pcm_8000"
	TranslationRealtimeInputFormatPcm16000 TranslationRealtimeInputFormat = "pcm_16000"
	TranslationRealtimeInputFormatPcm22050 TranslationRealtimeInputFormat = "pcm_22050"
	TranslationRealtimeInputFormatPcm24000 TranslationRealtimeInputFormat = "pcm_24000"
	TranslationRealtimeInputFormatPcm44100 TranslationRealtimeInputFormat = "pcm_44100"
	TranslationRealtimeInputFormatPcm48000 TranslationRealtimeInputFormat = "pcm_48000"
	TranslationRealtimeInputFormatUlaw8000 TranslationRealtimeInputFormat = "ulaw_8000"
)

func NewTranslationRealtimeInputFormatFromString(s string) (TranslationRealtimeInputFormat, error) {
	switch s {
	case "pcm_8000":
		return TranslationRealtimeInputFormatPcm8000, nil
	case "pcm_16000":
		return TranslationRealtimeInputFormatPcm16000, nil
	case "pcm_22050":
		return TranslationRealtimeInputFormatPcm22050, nil
	case "pcm_24000":
		return TranslationRealtimeInputFormatPcm24000, nil
	case "pcm_44100":
		return TranslationRealtimeInputFormatPcm44100, nil
	case "pcm_48000":
		return TranslationRealtimeInputFormatPcm48000, nil
	case "ulaw_8000":
		return TranslationRealtimeInputFormatUlaw8000, nil
	}
	var t TranslationRealtimeInputFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TranslationRealtimeInputFormat) Ptr() *TranslationRealtimeInputFormat {
	return &e
}
