package tests

import (
	"testing"
)

func TestBodyCreateAgentTestFolderV1ConvaiAgentTestingFoldersPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "parent_folder_id": "value"
}`
	assertRoundTrip(t, wire, new(BodyCreateAgentTestFolderV1ConvaiAgentTestingFoldersPost))
}

func TestBodyUpdateAgentTestFolderV1ConvaiAgentTestingFoldersFolderIdPatchRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value"
}`
	assertRoundTrip(t, wire, new(BodyUpdateAgentTestFolderV1ConvaiAgentTestingFoldersFolderIdPatch))
}
