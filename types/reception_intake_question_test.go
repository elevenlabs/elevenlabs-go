package types

import (
	"testing"
)

func TestReceptionIntakeQuestionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "label": "value",
  "kind": "text",
  "options": [
    "value"
  ],
  "required": true,
  "is_active": true
}`
	assertRoundTrip(t, wire, new(ReceptionIntakeQuestion))
}

func TestReceptionIntakeQuestionRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "label": "value",
  "kind": "text",
  "options": [
    "value"
  ],
  "required": true,
  "is_active": true
}`
	assertRoundTrip(t, wire, new(ReceptionIntakeQuestion))
}
