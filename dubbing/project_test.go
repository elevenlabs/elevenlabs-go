package dubbing

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBodyCreateDubbingProjectV1DubbingProjectPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source_url": "value",
  "reference": "value",
  "source_language": "value",
  "model_id": "dubbing_v1",
  "keyterms": [
    "value"
  ],
  "webhook_ids": [
    "value"
  ],
  "target_language": "value"
}`
	assertRoundTrip(t, wire, new(BodyCreateDubbingProjectV1DubbingProjectPost))
}

func TestListProjectRequestSortDirectionVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ListProjectRequestSortDirection
		wire    string
	}
	cases := []enumCase{
		{variant: ListProjectRequestSortDirectionAscending, wire: "\"ASCENDING\""},
		{variant: ListProjectRequestSortDirectionDescending, wire: "\"DESCENDING\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ListProjectRequestSortDirection
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

func TestCreateProjectRequestModelIdZeroVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant CreateProjectRequestModelIdZero
		wire    string
	}
	cases := []enumCase{
		{variant: CreateProjectRequestModelIdZeroDubbingV1, wire: "\"dubbing_v1\""},
		{variant: CreateProjectRequestModelIdZeroDubbingV2, wire: "\"dubbing_v2\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded CreateProjectRequestModelIdZero
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
