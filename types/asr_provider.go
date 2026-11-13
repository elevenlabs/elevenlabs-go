package types

import (
	"fmt"
)

type AsrProvider string

const (
	AsrProviderElevenlabs     AsrProvider = "elevenlabs"
	AsrProviderScribeRealtime AsrProvider = "scribe_realtime"
	AsrProviderScribeV2Turbo  AsrProvider = "scribe_v2_turbo"
)

func NewAsrProviderFromString(s string) (AsrProvider, error) {
	switch s {
	case "elevenlabs":
		return AsrProviderElevenlabs, nil
	case "scribe_realtime":
		return AsrProviderScribeRealtime, nil
	case "scribe_v2_turbo":
		return AsrProviderScribeV2Turbo, nil
	}
	var t AsrProvider
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AsrProvider) Ptr() *AsrProvider {
	return &e
}
