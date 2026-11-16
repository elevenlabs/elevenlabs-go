package types

import (
	"fmt"
)

type MusicOnlyOutputFormats string

const (
	MusicOnlyOutputFormatsMp348000128 MusicOnlyOutputFormats = "mp3_48000_128"
	MusicOnlyOutputFormatsMp348000192 MusicOnlyOutputFormats = "mp3_48000_192"
	MusicOnlyOutputFormatsMp348000240 MusicOnlyOutputFormats = "mp3_48000_240"
	MusicOnlyOutputFormatsMp348000320 MusicOnlyOutputFormats = "mp3_48000_320"
)

func NewMusicOnlyOutputFormatsFromString(s string) (MusicOnlyOutputFormats, error) {
	switch s {
	case "mp3_48000_128":
		return MusicOnlyOutputFormatsMp348000128, nil
	case "mp3_48000_192":
		return MusicOnlyOutputFormatsMp348000192, nil
	case "mp3_48000_240":
		return MusicOnlyOutputFormatsMp348000240, nil
	case "mp3_48000_320":
		return MusicOnlyOutputFormatsMp348000320, nil
	}
	var t MusicOnlyOutputFormats
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e MusicOnlyOutputFormats) Ptr() *MusicOnlyOutputFormats {
	return &e
}
