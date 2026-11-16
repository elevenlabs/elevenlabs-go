package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestIngestFailureVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant IngestFailure
		wire    string
	}
	cases := []enumCase{
		{variant: IngestFailureBlockedUrl, wire: "\"blocked_url\""},
		{variant: IngestFailureDnsFailure, wire: "\"dns_failure\""},
		{variant: IngestFailureHostUnreachable, wire: "\"host_unreachable\""},
		{variant: IngestFailureAccessRefused, wire: "\"access_refused\""},
		{variant: IngestFailureCrawlTimeout, wire: "\"crawl_timeout\""},
		{variant: IngestFailureNoPages, wire: "\"no_pages\""},
		{variant: IngestFailureNoContent, wire: "\"no_content\""},
		{variant: IngestFailureUpstreamTimeout, wire: "\"upstream_timeout\""},
		{variant: IngestFailureUpstreamError, wire: "\"upstream_error\""},
		{variant: IngestFailureUnknown, wire: "\"unknown\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded IngestFailure
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
