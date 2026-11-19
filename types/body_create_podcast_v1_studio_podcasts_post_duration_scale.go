package types

import (
	"fmt"
)

// Duration of the generated podcast. Must be one of:
// short - produces podcasts shorter than 3 minutes.
// default - produces podcasts roughly between 3-7 minutes.
// long - produces podcasts longer than 7 minutes.
type BodyCreatePodcastV1StudioPodcastsPostDurationScale string

const (
	BodyCreatePodcastV1StudioPodcastsPostDurationScaleShort   BodyCreatePodcastV1StudioPodcastsPostDurationScale = "short"
	BodyCreatePodcastV1StudioPodcastsPostDurationScaleDefault BodyCreatePodcastV1StudioPodcastsPostDurationScale = "default"
	BodyCreatePodcastV1StudioPodcastsPostDurationScaleLong    BodyCreatePodcastV1StudioPodcastsPostDurationScale = "long"
)

func NewBodyCreatePodcastV1StudioPodcastsPostDurationScaleFromString(s string) (BodyCreatePodcastV1StudioPodcastsPostDurationScale, error) {
	switch s {
	case "short":
		return BodyCreatePodcastV1StudioPodcastsPostDurationScaleShort, nil
	case "default":
		return BodyCreatePodcastV1StudioPodcastsPostDurationScaleDefault, nil
	case "long":
		return BodyCreatePodcastV1StudioPodcastsPostDurationScaleLong, nil
	}
	var t BodyCreatePodcastV1StudioPodcastsPostDurationScale
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodyCreatePodcastV1StudioPodcastsPostDurationScale) Ptr() *BodyCreatePodcastV1StudioPodcastsPostDurationScale {
	return &e
}
