package voices

import (
	"testing"
)

func TestCreatePvcVoiceRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "language": "value",
  "description": "value",
  "labels": {
    "key": "value"
  }
}`
	assertRoundTrip(t, wire, new(CreatePvcVoiceRequest))
}

func TestBodyEditPvcVoiceV1VoicesPvcVoiceIdPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "language": "value",
  "description": "value",
  "labels": {
    "key": "value"
  }
}`
	assertRoundTrip(t, wire, new(BodyEditPvcVoiceV1VoicesPvcVoiceIdPost))
}

func TestBodyRunPvcTrainingV1VoicesPvcVoiceIdTrainPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "model_id": "value"
}`
	assertRoundTrip(t, wire, new(BodyRunPvcTrainingV1VoicesPvcVoiceIdTrainPost))
}
