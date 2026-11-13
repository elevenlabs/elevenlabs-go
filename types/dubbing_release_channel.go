package types

import (
	"fmt"
)

type DubbingReleaseChannel string

const (
	DubbingReleaseChannelStable       DubbingReleaseChannel = "stable"
	DubbingReleaseChannelRelease      DubbingReleaseChannel = "release"
	DubbingReleaseChannelExperimental DubbingReleaseChannel = "experimental"
)

func NewDubbingReleaseChannelFromString(s string) (DubbingReleaseChannel, error) {
	switch s {
	case "stable":
		return DubbingReleaseChannelStable, nil
	case "release":
		return DubbingReleaseChannelRelease, nil
	case "experimental":
		return DubbingReleaseChannelExperimental, nil
	}
	var t DubbingReleaseChannel
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e DubbingReleaseChannel) Ptr() *DubbingReleaseChannel {
	return &e
}
