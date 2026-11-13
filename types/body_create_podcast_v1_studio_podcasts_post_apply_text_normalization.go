package types

import (
	"fmt"
)

type BodyCreatePodcastV1StudioPodcastsPostApplyTextNormalization string

const (
	BodyCreatePodcastV1StudioPodcastsPostApplyTextNormalizationAuto         BodyCreatePodcastV1StudioPodcastsPostApplyTextNormalization = "auto"
	BodyCreatePodcastV1StudioPodcastsPostApplyTextNormalizationOn           BodyCreatePodcastV1StudioPodcastsPostApplyTextNormalization = "on"
	BodyCreatePodcastV1StudioPodcastsPostApplyTextNormalizationOff          BodyCreatePodcastV1StudioPodcastsPostApplyTextNormalization = "off"
	BodyCreatePodcastV1StudioPodcastsPostApplyTextNormalizationApplyEnglish BodyCreatePodcastV1StudioPodcastsPostApplyTextNormalization = "apply_english"
)

func NewBodyCreatePodcastV1StudioPodcastsPostApplyTextNormalizationFromString(s string) (BodyCreatePodcastV1StudioPodcastsPostApplyTextNormalization, error) {
	switch s {
	case "auto":
		return BodyCreatePodcastV1StudioPodcastsPostApplyTextNormalizationAuto, nil
	case "on":
		return BodyCreatePodcastV1StudioPodcastsPostApplyTextNormalizationOn, nil
	case "off":
		return BodyCreatePodcastV1StudioPodcastsPostApplyTextNormalizationOff, nil
	case "apply_english":
		return BodyCreatePodcastV1StudioPodcastsPostApplyTextNormalizationApplyEnglish, nil
	}
	var t BodyCreatePodcastV1StudioPodcastsPostApplyTextNormalization
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodyCreatePodcastV1StudioPodcastsPostApplyTextNormalization) Ptr() *BodyCreatePodcastV1StudioPodcastsPostApplyTextNormalization {
	return &e
}
