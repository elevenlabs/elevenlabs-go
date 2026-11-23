package types

import (
	"fmt"
)

type AsrProvider string

const (
	AsrProviderElevenlabs          AsrProvider = "elevenlabs"
	AsrProviderScribeRealtime      AsrProvider = "scribe_realtime"
	AsrProviderScribeV2Turbo       AsrProvider = "scribe_v2_turbo"
	AsrProviderScribeV3Live        AsrProvider = "scribe_v3_live"
	AsrProviderScribeV3LiveMedical AsrProvider = "scribe_v3_live_medical"
)

func NewAsrProviderFromString(s string) (AsrProvider, error) {
	switch s {
	case "elevenlabs":
		return AsrProviderElevenlabs, nil
	case "scribe_realtime":
		return AsrProviderScribeRealtime, nil
	case "scribe_v2_turbo":
		return AsrProviderScribeV2Turbo, nil
	case "scribe_v3_live":
		return AsrProviderScribeV3Live, nil
	case "scribe_v3_live_medical":
		return AsrProviderScribeV3LiveMedical, nil
	}
	var t AsrProvider
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AsrProvider) Ptr() *AsrProvider {
	return &e
}
