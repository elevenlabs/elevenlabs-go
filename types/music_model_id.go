package types

import (
	"fmt"
)

type MusicModelId string

const (
	MusicModelIdMusicV1  MusicModelId = "music_v1"
	MusicModelIdMusicV2  MusicModelId = "music_v2"
	MusicModelIdMusicV25 MusicModelId = "music_v2_5"
)

func NewMusicModelIdFromString(s string) (MusicModelId, error) {
	switch s {
	case "music_v1":
		return MusicModelIdMusicV1, nil
	case "music_v2":
		return MusicModelIdMusicV2, nil
	case "music_v2_5":
		return MusicModelIdMusicV25, nil
	}
	var t MusicModelId
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e MusicModelId) Ptr() *MusicModelId {
	return &e
}
