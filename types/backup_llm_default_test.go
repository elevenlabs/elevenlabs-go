package types

import (
	"testing"
)

func TestBackupLlmDefaultRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "preference": "default"
}`
	assertRoundTrip(t, wire, new(BackupLlmDefault))
}

func TestBackupLlmDefaultRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "preference": "default"
}`
	assertRoundTrip(t, wire, new(BackupLlmDefault))
}
