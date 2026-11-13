package types

import (
	"fmt"
)

// Encoding and sample rate of output audio. PCM output is stereo; `ulaw_8000` output is mono.
type DubbingLiveOutputFormat string

const (
	DubbingLiveOutputFormatPcm8000  DubbingLiveOutputFormat = "pcm_8000"
	DubbingLiveOutputFormatPcm16000 DubbingLiveOutputFormat = "pcm_16000"
	DubbingLiveOutputFormatPcm22050 DubbingLiveOutputFormat = "pcm_22050"
	DubbingLiveOutputFormatPcm24000 DubbingLiveOutputFormat = "pcm_24000"
	DubbingLiveOutputFormatPcm44100 DubbingLiveOutputFormat = "pcm_44100"
	DubbingLiveOutputFormatPcm48000 DubbingLiveOutputFormat = "pcm_48000"
	DubbingLiveOutputFormatUlaw8000 DubbingLiveOutputFormat = "ulaw_8000"
)

func NewDubbingLiveOutputFormatFromString(s string) (DubbingLiveOutputFormat, error) {
	switch s {
	case "pcm_8000":
		return DubbingLiveOutputFormatPcm8000, nil
	case "pcm_16000":
		return DubbingLiveOutputFormatPcm16000, nil
	case "pcm_22050":
		return DubbingLiveOutputFormatPcm22050, nil
	case "pcm_24000":
		return DubbingLiveOutputFormatPcm24000, nil
	case "pcm_44100":
		return DubbingLiveOutputFormatPcm44100, nil
	case "pcm_48000":
		return DubbingLiveOutputFormatPcm48000, nil
	case "ulaw_8000":
		return DubbingLiveOutputFormatUlaw8000, nil
	}
	var t DubbingLiveOutputFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e DubbingLiveOutputFormat) Ptr() *DubbingLiveOutputFormat {
	return &e
}
