package types

import (
	"fmt"
)

// Controls the response shape when use_multi_channel is enabled. 'separate' (default) returns one transcript per channel under 'transcripts'. 'combined' merges all channels into a single transcript whose words are sorted by start time, each carrying a 'channel_index' - matching the single-channel response shape. 'combined' requires timestamps (timestamps_granularity must not be 'none') and does not support entity detection or redaction.
type ConvertSpeechToTextRequestMultichannelOutputStyle string

const (
	ConvertSpeechToTextRequestMultichannelOutputStyleSeparate ConvertSpeechToTextRequestMultichannelOutputStyle = "separate"
	ConvertSpeechToTextRequestMultichannelOutputStyleCombined ConvertSpeechToTextRequestMultichannelOutputStyle = "combined"
)

func NewConvertSpeechToTextRequestMultichannelOutputStyleFromString(s string) (ConvertSpeechToTextRequestMultichannelOutputStyle, error) {
	switch s {
	case "separate":
		return ConvertSpeechToTextRequestMultichannelOutputStyleSeparate, nil
	case "combined":
		return ConvertSpeechToTextRequestMultichannelOutputStyleCombined, nil
	}
	var t ConvertSpeechToTextRequestMultichannelOutputStyle
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConvertSpeechToTextRequestMultichannelOutputStyle) Ptr() *ConvertSpeechToTextRequestMultichannelOutputStyle {
	return &e
}
