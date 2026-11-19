package music

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBodyCreateMusicFinetuneV1MusicFinetunesPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "primary_genre": "value",
  "tags": [
    "value"
  ],
  "visibility": "private",
  "model_id": "music_v1"
}`
	assertRoundTrip(t, wire, new(BodyCreateMusicFinetuneV1MusicFinetunesPost))
}

func TestUpdateMusicFinetuneRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "tags": [
    "value"
  ],
  "primary_genre": "value",
  "visibility": "private"
}`
	assertRoundTrip(t, wire, new(UpdateMusicFinetuneRequestModel))
}

func TestListFinetunesRequestSortVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ListFinetunesRequestSort
		wire    string
	}
	cases := []enumCase{
		{variant: ListFinetunesRequestSortCreatedAt, wire: "\"created_at\""},
		{variant: ListFinetunesRequestSortName, wire: "\"name\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ListFinetunesRequestSort
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

func TestListFinetunesRequestSortDirectionVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ListFinetunesRequestSortDirection
		wire    string
	}
	cases := []enumCase{
		{variant: ListFinetunesRequestSortDirectionAsc, wire: "\"asc\""},
		{variant: ListFinetunesRequestSortDirectionDesc, wire: "\"desc\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ListFinetunesRequestSortDirection
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

func TestCreateFinetunesRequestVisibilityVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant CreateFinetunesRequestVisibility
		wire    string
	}
	cases := []enumCase{
		{variant: CreateFinetunesRequestVisibilityPrivate, wire: "\"private\""},
		{variant: CreateFinetunesRequestVisibilityWorkspace, wire: "\"workspace\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded CreateFinetunesRequestVisibility
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

func TestUpdateMusicFinetuneRequestModelVisibilityVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant UpdateMusicFinetuneRequestModelVisibility
		wire    string
	}
	cases := []enumCase{
		{variant: UpdateMusicFinetuneRequestModelVisibilityPrivate, wire: "\"private\""},
		{variant: UpdateMusicFinetuneRequestModelVisibilityWorkspace, wire: "\"workspace\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded UpdateMusicFinetuneRequestModelVisibility
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
