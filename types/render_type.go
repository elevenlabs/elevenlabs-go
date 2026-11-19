package types

import (
	"fmt"
)

type RenderType string

const (
	RenderTypeMp4       RenderType = "mp4"
	RenderTypeAac       RenderType = "aac"
	RenderTypeMp3       RenderType = "mp3"
	RenderTypeWav       RenderType = "wav"
	RenderTypeAaf       RenderType = "aaf"
	RenderTypeTracksZip RenderType = "tracks_zip"
	RenderTypeClipsZip  RenderType = "clips_zip"
	RenderTypeZip       RenderType = "zip"
)

func NewRenderTypeFromString(s string) (RenderType, error) {
	switch s {
	case "mp4":
		return RenderTypeMp4, nil
	case "aac":
		return RenderTypeAac, nil
	case "mp3":
		return RenderTypeMp3, nil
	case "wav":
		return RenderTypeWav, nil
	case "aaf":
		return RenderTypeAaf, nil
	case "tracks_zip":
		return RenderTypeTracksZip, nil
	case "clips_zip":
		return RenderTypeClipsZip, nil
	case "zip":
		return RenderTypeZip, nil
	}
	var t RenderType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e RenderType) Ptr() *RenderType {
	return &e
}
