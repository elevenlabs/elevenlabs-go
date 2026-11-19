package types

import (
	"fmt"
)

// Audio encoding format for realtime translation output.
type TranslationRealtimeOutputFormat string

const (
	TranslationRealtimeOutputFormatMp32205032  TranslationRealtimeOutputFormat = "mp3_22050_32"
	TranslationRealtimeOutputFormatMp34410032  TranslationRealtimeOutputFormat = "mp3_44100_32"
	TranslationRealtimeOutputFormatMp34410064  TranslationRealtimeOutputFormat = "mp3_44100_64"
	TranslationRealtimeOutputFormatMp34410096  TranslationRealtimeOutputFormat = "mp3_44100_96"
	TranslationRealtimeOutputFormatMp344100128 TranslationRealtimeOutputFormat = "mp3_44100_128"
	TranslationRealtimeOutputFormatMp344100192 TranslationRealtimeOutputFormat = "mp3_44100_192"
	TranslationRealtimeOutputFormatPcm16000    TranslationRealtimeOutputFormat = "pcm_16000"
	TranslationRealtimeOutputFormatPcm22050    TranslationRealtimeOutputFormat = "pcm_22050"
	TranslationRealtimeOutputFormatPcm24000    TranslationRealtimeOutputFormat = "pcm_24000"
	TranslationRealtimeOutputFormatPcm44100    TranslationRealtimeOutputFormat = "pcm_44100"
	TranslationRealtimeOutputFormatUlaw8000    TranslationRealtimeOutputFormat = "ulaw_8000"
)

func NewTranslationRealtimeOutputFormatFromString(s string) (TranslationRealtimeOutputFormat, error) {
	switch s {
	case "mp3_22050_32":
		return TranslationRealtimeOutputFormatMp32205032, nil
	case "mp3_44100_32":
		return TranslationRealtimeOutputFormatMp34410032, nil
	case "mp3_44100_64":
		return TranslationRealtimeOutputFormatMp34410064, nil
	case "mp3_44100_96":
		return TranslationRealtimeOutputFormatMp34410096, nil
	case "mp3_44100_128":
		return TranslationRealtimeOutputFormatMp344100128, nil
	case "mp3_44100_192":
		return TranslationRealtimeOutputFormatMp344100192, nil
	case "pcm_16000":
		return TranslationRealtimeOutputFormatPcm16000, nil
	case "pcm_22050":
		return TranslationRealtimeOutputFormatPcm22050, nil
	case "pcm_24000":
		return TranslationRealtimeOutputFormatPcm24000, nil
	case "pcm_44100":
		return TranslationRealtimeOutputFormatPcm44100, nil
	case "ulaw_8000":
		return TranslationRealtimeOutputFormatUlaw8000, nil
	}
	var t TranslationRealtimeOutputFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TranslationRealtimeOutputFormat) Ptr() *TranslationRealtimeOutputFormat {
	return &e
}
