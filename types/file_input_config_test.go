package types

import (
	"testing"
)

func TestFileInputConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "enabled": true,
  "max_files_in_memory": 30,
  "max_files_per_conversation": 123
}`
	assertRoundTrip(t, wire, new(FileInputConfig))
}

func TestFileInputConfigRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "enabled": true,
  "max_files_in_memory": 30,
  "max_files_per_conversation": 123
}`
	assertRoundTrip(t, wire, new(FileInputConfig))
}

func TestFileInputConfigRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "enabled": true,
  "max_files_in_memory": 30,
  "max_files_per_conversation": 123
}`
	assertRoundTrip(t, wire, new(FileInputConfig))
}
