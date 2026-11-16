package types

import (
	"testing"
)

func TestBackupLlmOverrideRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "preference": "override",
  "order": [
    "gpt-4o-mini"
  ]
}`
	assertRoundTrip(t, wire, new(BackupLlmOverride))
}

func TestBackupLlmOverrideRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "preference": "override",
  "order": [
    "gpt-4o-mini"
  ]
}`
	assertRoundTrip(t, wire, new(BackupLlmOverride))
}
