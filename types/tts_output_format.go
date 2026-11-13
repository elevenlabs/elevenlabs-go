package types

import (
	"fmt"
)

type TtsOutputFormat string

const (
	TtsOutputFormatPcm8000  TtsOutputFormat = "pcm_8000"
	TtsOutputFormatPcm16000 TtsOutputFormat = "pcm_16000"
	TtsOutputFormatPcm22050 TtsOutputFormat = "pcm_22050"
	TtsOutputFormatPcm24000 TtsOutputFormat = "pcm_24000"
	TtsOutputFormatPcm44100 TtsOutputFormat = "pcm_44100"
	TtsOutputFormatPcm48000 TtsOutputFormat = "pcm_48000"
	TtsOutputFormatUlaw8000 TtsOutputFormat = "ulaw_8000"
)

func NewTtsOutputFormatFromString(s string) (TtsOutputFormat, error) {
	switch s {
	case "pcm_8000":
		return TtsOutputFormatPcm8000, nil
	case "pcm_16000":
		return TtsOutputFormatPcm16000, nil
	case "pcm_22050":
		return TtsOutputFormatPcm22050, nil
	case "pcm_24000":
		return TtsOutputFormatPcm24000, nil
	case "pcm_44100":
		return TtsOutputFormatPcm44100, nil
	case "pcm_48000":
		return TtsOutputFormatPcm48000, nil
	case "ulaw_8000":
		return TtsOutputFormatUlaw8000, nil
	}
	var t TtsOutputFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TtsOutputFormat) Ptr() *TtsOutputFormat {
	return &e
}
