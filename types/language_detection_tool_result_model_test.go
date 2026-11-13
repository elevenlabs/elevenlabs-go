package types

import (
	"testing"
)

func TestLanguageDetectionToolResultModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "language_detection_success",
  "status": "success",
  "reason": "value",
  "language": "value"
}`
	assertRoundTrip(t, wire, new(LanguageDetectionToolResultModel))
}

func TestLanguageDetectionToolResultModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "language_detection_success",
  "status": "success",
  "reason": "value",
  "language": "value"
}`
	assertRoundTrip(t, wire, new(LanguageDetectionToolResultModel))
}
