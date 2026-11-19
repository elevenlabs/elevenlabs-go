package types

import (
	"fmt"
)

type RealtimeModel string

const (
	RealtimeModelElevenRealtimeV1Mini RealtimeModel = "eleven_realtime_v1_mini"
	RealtimeModelElevenRealtimeV1Max  RealtimeModel = "eleven_realtime_v1_max"
)

func NewRealtimeModelFromString(s string) (RealtimeModel, error) {
	switch s {
	case "eleven_realtime_v1_mini":
		return RealtimeModelElevenRealtimeV1Mini, nil
	case "eleven_realtime_v1_max":
		return RealtimeModelElevenRealtimeV1Max, nil
	}
	var t RealtimeModel
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e RealtimeModel) Ptr() *RealtimeModel {
	return &e
}
