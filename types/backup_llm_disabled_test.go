package types

import (
	"testing"
)

func TestBackupLlmDisabledRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "preference": "disabled"
}`
	assertRoundTrip(t, wire, new(BackupLlmDisabled))
}

func TestBackupLlmDisabledRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "preference": "disabled"
}`
	assertRoundTrip(t, wire, new(BackupLlmDisabled))
}
