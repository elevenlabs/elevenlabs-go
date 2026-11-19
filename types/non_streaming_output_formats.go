package types

import (
	"fmt"
)

type NonStreamingOutputFormats string

const (
	NonStreamingOutputFormatsWav8000  NonStreamingOutputFormats = "wav_8000"
	NonStreamingOutputFormatsWav16000 NonStreamingOutputFormats = "wav_16000"
	NonStreamingOutputFormatsWav22050 NonStreamingOutputFormats = "wav_22050"
	NonStreamingOutputFormatsWav24000 NonStreamingOutputFormats = "wav_24000"
	NonStreamingOutputFormatsWav32000 NonStreamingOutputFormats = "wav_32000"
	NonStreamingOutputFormatsWav44100 NonStreamingOutputFormats = "wav_44100"
	NonStreamingOutputFormatsWav48000 NonStreamingOutputFormats = "wav_48000"
)

func NewNonStreamingOutputFormatsFromString(s string) (NonStreamingOutputFormats, error) {
	switch s {
	case "wav_8000":
		return NonStreamingOutputFormatsWav8000, nil
	case "wav_16000":
		return NonStreamingOutputFormatsWav16000, nil
	case "wav_22050":
		return NonStreamingOutputFormatsWav22050, nil
	case "wav_24000":
		return NonStreamingOutputFormatsWav24000, nil
	case "wav_32000":
		return NonStreamingOutputFormatsWav32000, nil
	case "wav_44100":
		return NonStreamingOutputFormatsWav44100, nil
	case "wav_48000":
		return NonStreamingOutputFormatsWav48000, nil
	}
	var t NonStreamingOutputFormats
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e NonStreamingOutputFormats) Ptr() *NonStreamingOutputFormats {
	return &e
}
