package types

import (
	"fmt"
)

type AllowedOutputFormats string

const (
	AllowedOutputFormatsMp32205032   AllowedOutputFormats = "mp3_22050_32"
	AllowedOutputFormatsMp32400048   AllowedOutputFormats = "mp3_24000_48"
	AllowedOutputFormatsMp34410032   AllowedOutputFormats = "mp3_44100_32"
	AllowedOutputFormatsMp34410064   AllowedOutputFormats = "mp3_44100_64"
	AllowedOutputFormatsMp34410096   AllowedOutputFormats = "mp3_44100_96"
	AllowedOutputFormatsMp344100128  AllowedOutputFormats = "mp3_44100_128"
	AllowedOutputFormatsMp344100192  AllowedOutputFormats = "mp3_44100_192"
	AllowedOutputFormatsPcm8000      AllowedOutputFormats = "pcm_8000"
	AllowedOutputFormatsPcm16000     AllowedOutputFormats = "pcm_16000"
	AllowedOutputFormatsPcm22050     AllowedOutputFormats = "pcm_22050"
	AllowedOutputFormatsPcm24000     AllowedOutputFormats = "pcm_24000"
	AllowedOutputFormatsPcm32000     AllowedOutputFormats = "pcm_32000"
	AllowedOutputFormatsPcm44100     AllowedOutputFormats = "pcm_44100"
	AllowedOutputFormatsPcm48000     AllowedOutputFormats = "pcm_48000"
	AllowedOutputFormatsUlaw8000     AllowedOutputFormats = "ulaw_8000"
	AllowedOutputFormatsAlaw8000     AllowedOutputFormats = "alaw_8000"
	AllowedOutputFormatsOpus4800032  AllowedOutputFormats = "opus_48000_32"
	AllowedOutputFormatsOpus4800064  AllowedOutputFormats = "opus_48000_64"
	AllowedOutputFormatsOpus4800096  AllowedOutputFormats = "opus_48000_96"
	AllowedOutputFormatsOpus48000128 AllowedOutputFormats = "opus_48000_128"
	AllowedOutputFormatsOpus48000192 AllowedOutputFormats = "opus_48000_192"
)

func NewAllowedOutputFormatsFromString(s string) (AllowedOutputFormats, error) {
	switch s {
	case "mp3_22050_32":
		return AllowedOutputFormatsMp32205032, nil
	case "mp3_24000_48":
		return AllowedOutputFormatsMp32400048, nil
	case "mp3_44100_32":
		return AllowedOutputFormatsMp34410032, nil
	case "mp3_44100_64":
		return AllowedOutputFormatsMp34410064, nil
	case "mp3_44100_96":
		return AllowedOutputFormatsMp34410096, nil
	case "mp3_44100_128":
		return AllowedOutputFormatsMp344100128, nil
	case "mp3_44100_192":
		return AllowedOutputFormatsMp344100192, nil
	case "pcm_8000":
		return AllowedOutputFormatsPcm8000, nil
	case "pcm_16000":
		return AllowedOutputFormatsPcm16000, nil
	case "pcm_22050":
		return AllowedOutputFormatsPcm22050, nil
	case "pcm_24000":
		return AllowedOutputFormatsPcm24000, nil
	case "pcm_32000":
		return AllowedOutputFormatsPcm32000, nil
	case "pcm_44100":
		return AllowedOutputFormatsPcm44100, nil
	case "pcm_48000":
		return AllowedOutputFormatsPcm48000, nil
	case "ulaw_8000":
		return AllowedOutputFormatsUlaw8000, nil
	case "alaw_8000":
		return AllowedOutputFormatsAlaw8000, nil
	case "opus_48000_32":
		return AllowedOutputFormatsOpus4800032, nil
	case "opus_48000_64":
		return AllowedOutputFormatsOpus4800064, nil
	case "opus_48000_96":
		return AllowedOutputFormatsOpus4800096, nil
	case "opus_48000_128":
		return AllowedOutputFormatsOpus48000128, nil
	case "opus_48000_192":
		return AllowedOutputFormatsOpus48000192, nil
	}
	var t AllowedOutputFormats
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AllowedOutputFormats) Ptr() *AllowedOutputFormats {
	return &e
}
