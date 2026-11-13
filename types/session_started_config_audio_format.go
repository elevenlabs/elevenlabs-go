package types

import (
	"fmt"
)

type SessionStartedConfigAudioFormat string

const (
	SessionStartedConfigAudioFormatPcm8000  SessionStartedConfigAudioFormat = "pcm_8000"
	SessionStartedConfigAudioFormatPcm16000 SessionStartedConfigAudioFormat = "pcm_16000"
	SessionStartedConfigAudioFormatPcm22050 SessionStartedConfigAudioFormat = "pcm_22050"
	SessionStartedConfigAudioFormatPcm24000 SessionStartedConfigAudioFormat = "pcm_24000"
	SessionStartedConfigAudioFormatPcm44100 SessionStartedConfigAudioFormat = "pcm_44100"
	SessionStartedConfigAudioFormatPcm48000 SessionStartedConfigAudioFormat = "pcm_48000"
	SessionStartedConfigAudioFormatUlaw8000 SessionStartedConfigAudioFormat = "ulaw_8000"
)

func NewSessionStartedConfigAudioFormatFromString(s string) (SessionStartedConfigAudioFormat, error) {
	switch s {
	case "pcm_8000":
		return SessionStartedConfigAudioFormatPcm8000, nil
	case "pcm_16000":
		return SessionStartedConfigAudioFormatPcm16000, nil
	case "pcm_22050":
		return SessionStartedConfigAudioFormatPcm22050, nil
	case "pcm_24000":
		return SessionStartedConfigAudioFormatPcm24000, nil
	case "pcm_44100":
		return SessionStartedConfigAudioFormatPcm44100, nil
	case "pcm_48000":
		return SessionStartedConfigAudioFormatPcm48000, nil
	case "ulaw_8000":
		return SessionStartedConfigAudioFormatUlaw8000, nil
	}
	var t SessionStartedConfigAudioFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SessionStartedConfigAudioFormat) Ptr() *SessionStartedConfigAudioFormat {
	return &e
}
