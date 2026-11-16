package types

import (
	"fmt"
)

type ConvertSpeechToTextRequestNumbersFormat string

const (
	ConvertSpeechToTextRequestNumbersFormatDigits ConvertSpeechToTextRequestNumbersFormat = "digits"
	ConvertSpeechToTextRequestNumbersFormatWords  ConvertSpeechToTextRequestNumbersFormat = "words"
)

func NewConvertSpeechToTextRequestNumbersFormatFromString(s string) (ConvertSpeechToTextRequestNumbersFormat, error) {
	switch s {
	case "digits":
		return ConvertSpeechToTextRequestNumbersFormatDigits, nil
	case "words":
		return ConvertSpeechToTextRequestNumbersFormatWords, nil
	}
	var t ConvertSpeechToTextRequestNumbersFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConvertSpeechToTextRequestNumbersFormat) Ptr() *ConvertSpeechToTextRequestNumbersFormat {
	return &e
}
