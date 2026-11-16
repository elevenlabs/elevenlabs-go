package types

import (
	"fmt"
)

type MusicGenerationMode string

const (
	MusicGenerationModeTrack        MusicGenerationMode = "track"
	MusicGenerationModeLoop         MusicGenerationMode = "loop"
	MusicGenerationModeAmbience     MusicGenerationMode = "ambience"
	MusicGenerationModeVideoToMusic MusicGenerationMode = "video_to_music"
)

func NewMusicGenerationModeFromString(s string) (MusicGenerationMode, error) {
	switch s {
	case "track":
		return MusicGenerationModeTrack, nil
	case "loop":
		return MusicGenerationModeLoop, nil
	case "ambience":
		return MusicGenerationModeAmbience, nil
	case "video_to_music":
		return MusicGenerationModeVideoToMusic, nil
	}
	var t MusicGenerationMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e MusicGenerationMode) Ptr() *MusicGenerationMode {
	return &e
}
