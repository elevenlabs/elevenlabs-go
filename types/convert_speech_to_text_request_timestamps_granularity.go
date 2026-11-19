package types

import (
	"fmt"
)

// The granularity of the timestamps in the transcription. 'word' provides word-level timestamps and 'character' provides character-level timestamps per word.
type ConvertSpeechToTextRequestTimestampsGranularity string

const (
	ConvertSpeechToTextRequestTimestampsGranularityNone      ConvertSpeechToTextRequestTimestampsGranularity = "none"
	ConvertSpeechToTextRequestTimestampsGranularityWord      ConvertSpeechToTextRequestTimestampsGranularity = "word"
	ConvertSpeechToTextRequestTimestampsGranularityCharacter ConvertSpeechToTextRequestTimestampsGranularity = "character"
)

func NewConvertSpeechToTextRequestTimestampsGranularityFromString(s string) (ConvertSpeechToTextRequestTimestampsGranularity, error) {
	switch s {
	case "none":
		return ConvertSpeechToTextRequestTimestampsGranularityNone, nil
	case "word":
		return ConvertSpeechToTextRequestTimestampsGranularityWord, nil
	case "character":
		return ConvertSpeechToTextRequestTimestampsGranularityCharacter, nil
	}
	var t ConvertSpeechToTextRequestTimestampsGranularity
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConvertSpeechToTextRequestTimestampsGranularity) Ptr() *ConvertSpeechToTextRequestTimestampsGranularity {
	return &e
}
