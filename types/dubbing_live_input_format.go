package types

import (
	"fmt"
)

// Encoding and sample rate of each input audio chunk.
type DubbingLiveInputFormat string

const (
	DubbingLiveInputFormatPcm8000  DubbingLiveInputFormat = "pcm_8000"
	DubbingLiveInputFormatPcm16000 DubbingLiveInputFormat = "pcm_16000"
	DubbingLiveInputFormatPcm22050 DubbingLiveInputFormat = "pcm_22050"
	DubbingLiveInputFormatPcm24000 DubbingLiveInputFormat = "pcm_24000"
	DubbingLiveInputFormatPcm44100 DubbingLiveInputFormat = "pcm_44100"
	DubbingLiveInputFormatPcm48000 DubbingLiveInputFormat = "pcm_48000"
	DubbingLiveInputFormatUlaw8000 DubbingLiveInputFormat = "ulaw_8000"
)

func NewDubbingLiveInputFormatFromString(s string) (DubbingLiveInputFormat, error) {
	switch s {
	case "pcm_8000":
		return DubbingLiveInputFormatPcm8000, nil
	case "pcm_16000":
		return DubbingLiveInputFormatPcm16000, nil
	case "pcm_22050":
		return DubbingLiveInputFormatPcm22050, nil
	case "pcm_24000":
		return DubbingLiveInputFormatPcm24000, nil
	case "pcm_44100":
		return DubbingLiveInputFormatPcm44100, nil
	case "pcm_48000":
		return DubbingLiveInputFormatPcm48000, nil
	case "ulaw_8000":
		return DubbingLiveInputFormatUlaw8000, nil
	}
	var t DubbingLiveInputFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e DubbingLiveInputFormat) Ptr() *DubbingLiveInputFormat {
	return &e
}
