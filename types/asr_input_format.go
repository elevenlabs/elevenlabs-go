package types

import (
	"fmt"
)

type AsrInputFormat string

const (
	AsrInputFormatPcm8000  AsrInputFormat = "pcm_8000"
	AsrInputFormatPcm16000 AsrInputFormat = "pcm_16000"
	AsrInputFormatPcm22050 AsrInputFormat = "pcm_22050"
	AsrInputFormatPcm24000 AsrInputFormat = "pcm_24000"
	AsrInputFormatPcm44100 AsrInputFormat = "pcm_44100"
	AsrInputFormatPcm48000 AsrInputFormat = "pcm_48000"
	AsrInputFormatUlaw8000 AsrInputFormat = "ulaw_8000"
)

func NewAsrInputFormatFromString(s string) (AsrInputFormat, error) {
	switch s {
	case "pcm_8000":
		return AsrInputFormatPcm8000, nil
	case "pcm_16000":
		return AsrInputFormatPcm16000, nil
	case "pcm_22050":
		return AsrInputFormatPcm22050, nil
	case "pcm_24000":
		return AsrInputFormatPcm24000, nil
	case "pcm_44100":
		return AsrInputFormatPcm44100, nil
	case "pcm_48000":
		return AsrInputFormatPcm48000, nil
	case "ulaw_8000":
		return AsrInputFormatUlaw8000, nil
	}
	var t AsrInputFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AsrInputFormat) Ptr() *AsrInputFormat {
	return &e
}
