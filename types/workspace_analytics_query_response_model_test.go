package types

import (
	"testing"
)

func TestWorkspaceAnalyticsQueryResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "columns": [
    "value"
  ],
  "column_types": [
    "String"
  ],
  "rows": [
    [
      "value"
    ]
  ],
  "column_units": [
    "ms"
  ]
}`
	assertRoundTrip(t, wire, new(WorkspaceAnalyticsQueryResponseModel))
}

func TestWorkspaceAnalyticsQueryResponseModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "columns": [
    "value"
  ],
  "column_types": [
    "String"
  ],
  "rows": [
    [
      "value"
    ]
  ],
  "column_units": [
    "ms"
  ]
}`
	assertRoundTrip(t, wire, new(WorkspaceAnalyticsQueryResponseModel))
}
