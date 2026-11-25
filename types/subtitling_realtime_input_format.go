package types

import (
	"fmt"
)

// Audio encoding format of the source audio.
type SubtitlingRealtimeInputFormat string

const (
	SubtitlingRealtimeInputFormatPcm8000  SubtitlingRealtimeInputFormat = "pcm_8000"
	SubtitlingRealtimeInputFormatPcm16000 SubtitlingRealtimeInputFormat = "pcm_16000"
	SubtitlingRealtimeInputFormatPcm22050 SubtitlingRealtimeInputFormat = "pcm_22050"
	SubtitlingRealtimeInputFormatPcm24000 SubtitlingRealtimeInputFormat = "pcm_24000"
	SubtitlingRealtimeInputFormatPcm44100 SubtitlingRealtimeInputFormat = "pcm_44100"
	SubtitlingRealtimeInputFormatPcm48000 SubtitlingRealtimeInputFormat = "pcm_48000"
	SubtitlingRealtimeInputFormatUlaw8000 SubtitlingRealtimeInputFormat = "ulaw_8000"
)

func NewSubtitlingRealtimeInputFormatFromString(s string) (SubtitlingRealtimeInputFormat, error) {
	switch s {
	case "pcm_8000":
		return SubtitlingRealtimeInputFormatPcm8000, nil
	case "pcm_16000":
		return SubtitlingRealtimeInputFormatPcm16000, nil
	case "pcm_22050":
		return SubtitlingRealtimeInputFormatPcm22050, nil
	case "pcm_24000":
		return SubtitlingRealtimeInputFormatPcm24000, nil
	case "pcm_44100":
		return SubtitlingRealtimeInputFormatPcm44100, nil
	case "pcm_48000":
		return SubtitlingRealtimeInputFormatPcm48000, nil
	case "ulaw_8000":
		return SubtitlingRealtimeInputFormatUlaw8000, nil
	}
	var t SubtitlingRealtimeInputFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SubtitlingRealtimeInputFormat) Ptr() *SubtitlingRealtimeInputFormat {
	return &e
}
