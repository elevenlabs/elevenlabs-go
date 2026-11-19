package types

import (
	"fmt"
)

type OutputFormat string

const (
	OutputFormatMp32205032  OutputFormat = "mp3_22050_32"
	OutputFormatMp34410032  OutputFormat = "mp3_44100_32"
	OutputFormatMp34410064  OutputFormat = "mp3_44100_64"
	OutputFormatMp34410096  OutputFormat = "mp3_44100_96"
	OutputFormatMp344100128 OutputFormat = "mp3_44100_128"
	OutputFormatMp344100192 OutputFormat = "mp3_44100_192"
	OutputFormatPcm16000    OutputFormat = "pcm_16000"
	OutputFormatPcm22050    OutputFormat = "pcm_22050"
	OutputFormatPcm24000    OutputFormat = "pcm_24000"
	OutputFormatPcm44100    OutputFormat = "pcm_44100"
	OutputFormatUlaw8000    OutputFormat = "ulaw_8000"
)

func NewOutputFormatFromString(s string) (OutputFormat, error) {
	switch s {
	case "mp3_22050_32":
		return OutputFormatMp32205032, nil
	case "mp3_44100_32":
		return OutputFormatMp34410032, nil
	case "mp3_44100_64":
		return OutputFormatMp34410064, nil
	case "mp3_44100_96":
		return OutputFormatMp34410096, nil
	case "mp3_44100_128":
		return OutputFormatMp344100128, nil
	case "mp3_44100_192":
		return OutputFormatMp344100192, nil
	case "pcm_16000":
		return OutputFormatPcm16000, nil
	case "pcm_22050":
		return OutputFormatPcm22050, nil
	case "pcm_24000":
		return OutputFormatPcm24000, nil
	case "pcm_44100":
		return OutputFormatPcm44100, nil
	case "ulaw_8000":
		return OutputFormatUlaw8000, nil
	}
	var t OutputFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e OutputFormat) Ptr() *OutputFormat {
	return &e
}
