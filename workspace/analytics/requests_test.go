package analytics

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBodyListApiRequestsV1WorkspaceAnalyticsRequestsPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "start_time": 1577836800000,
  "end_time": 1577836800000,
  "limit": 123,
  "sort": "asc",
  "filters": [
    {
      "column": "value",
      "operation": "in",
      "values": [
        "value"
      ]
    }
  ],
  "search": "value"
}`
	assertRoundTrip(t, wire, new(BodyListApiRequestsV1WorkspaceAnalyticsRequestsPost))
}

func TestBodyListApiRequestsV1WorkspaceAnalyticsRequestsPostSortVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant BodyListApiRequestsV1WorkspaceAnalyticsRequestsPostSort
		wire    string
	}
	cases := []enumCase{
		{variant: BodyListApiRequestsV1WorkspaceAnalyticsRequestsPostSortAsc, wire: "\"asc\""},
		{variant: BodyListApiRequestsV1WorkspaceAnalyticsRequestsPostSortDesc, wire: "\"desc\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded BodyListApiRequestsV1WorkspaceAnalyticsRequestsPostSort
			if err := json.Unmarshal([]byte(tc.wire), &decoded); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if decoded != tc.variant {
				t.Fatalf("decoded mismatch: got %#v want %#v", decoded, tc.variant)
			}
			var want any
			var got any
			if err := json.Unmarshal([]byte(tc.wire), &want); err != nil {
				t.Fatalf("unmarshal want: %v", err)
			}
			if err := json.Unmarshal(encoded, &got); err != nil {
				t.Fatalf("unmarshal got: %v", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("round-trip mismatch\n got: %#v\nwant: %#v", got, want)
			}
		})
	}
}
