package types

import (
	"fmt"
)

type MediaCodec string

const (
	MediaCodecG7228000 MediaCodec = "G722/8000"
	MediaCodecPcmu8000 MediaCodec = "PCMU/8000"
	MediaCodecPcma8000 MediaCodec = "PCMA/8000"
)

func NewMediaCodecFromString(s string) (MediaCodec, error) {
	switch s {
	case "G722/8000":
		return MediaCodecG7228000, nil
	case "PCMU/8000":
		return MediaCodecPcmu8000, nil
	case "PCMA/8000":
		return MediaCodecPcma8000, nil
	}
	var t MediaCodec
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e MediaCodec) Ptr() *MediaCodec {
	return &e
}
