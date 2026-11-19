package types

import (
	"fmt"
)

// The audio encoding of the output, as `codec_sampleRateHz_bitrateKbps`. `mp3_44100_192` requires the Creator tier or above.
type ElevenFlashV25RequestOutputFormat string

const (
	ElevenFlashV25RequestOutputFormatMp32205032  ElevenFlashV25RequestOutputFormat = "mp3_22050_32"
	ElevenFlashV25RequestOutputFormatMp32400048  ElevenFlashV25RequestOutputFormat = "mp3_24000_48"
	ElevenFlashV25RequestOutputFormatMp34410032  ElevenFlashV25RequestOutputFormat = "mp3_44100_32"
	ElevenFlashV25RequestOutputFormatMp34410064  ElevenFlashV25RequestOutputFormat = "mp3_44100_64"
	ElevenFlashV25RequestOutputFormatMp34410096  ElevenFlashV25RequestOutputFormat = "mp3_44100_96"
	ElevenFlashV25RequestOutputFormatMp344100128 ElevenFlashV25RequestOutputFormat = "mp3_44100_128"
	ElevenFlashV25RequestOutputFormatMp344100192 ElevenFlashV25RequestOutputFormat = "mp3_44100_192"
)

func NewElevenFlashV25RequestOutputFormatFromString(s string) (ElevenFlashV25RequestOutputFormat, error) {
	switch s {
	case "mp3_22050_32":
		return ElevenFlashV25RequestOutputFormatMp32205032, nil
	case "mp3_24000_48":
		return ElevenFlashV25RequestOutputFormatMp32400048, nil
	case "mp3_44100_32":
		return ElevenFlashV25RequestOutputFormatMp34410032, nil
	case "mp3_44100_64":
		return ElevenFlashV25RequestOutputFormatMp34410064, nil
	case "mp3_44100_96":
		return ElevenFlashV25RequestOutputFormatMp34410096, nil
	case "mp3_44100_128":
		return ElevenFlashV25RequestOutputFormatMp344100128, nil
	case "mp3_44100_192":
		return ElevenFlashV25RequestOutputFormatMp344100192, nil
	}
	var t ElevenFlashV25RequestOutputFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ElevenFlashV25RequestOutputFormat) Ptr() *ElevenFlashV25RequestOutputFormat {
	return &e
}
