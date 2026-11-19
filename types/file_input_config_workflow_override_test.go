package types

import (
	"testing"
)

func TestFileInputConfigWorkflowOverrideRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "enabled": true,
  "max_files_in_memory": 123,
  "max_files_per_conversation": 123
}`
	assertRoundTrip(t, wire, new(FileInputConfigWorkflowOverride))
}
