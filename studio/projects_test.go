package studio

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBodyCreateStudioProjectV1StudioProjectsPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "default_title_voice_id": "value",
  "default_paragraph_voice_id": "value",
  "default_model_id": "value",
  "from_url": "value",
  "from_content_json": "value",
  "quality_preset": "standard",
  "title": "value",
  "author": "value",
  "description": "value",
  "genres": [
    "value"
  ],
  "target_audience": "children",
  "language": "va",
  "content_type": "value",
  "mature_content": true,
  "isbn_number": "value",
  "acx_volume_normalization": true,
  "volume_normalization": true,
  "pronunciation_dictionary_locators": [
    "value"
  ],
  "callback_url": "value",
  "fiction": "fiction",
  "apply_text_normalization": "auto",
  "auto_convert": true,
  "auto_assign_voices": true,
  "source_type": "blank",
  "voice_settings": [
    "value"
  ],
  "create_publishing_read": true
}`
	assertRoundTrip(t, wire, new(BodyCreateStudioProjectV1StudioProjectsPost))
}

func TestBodyUpdateStudioProjectV1StudioProjectsProjectIdPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "default_title_voice_id": "value",
  "default_paragraph_voice_id": "value",
  "title": "value",
  "author": "value",
  "isbn_number": "value",
  "volume_normalization": true
}`
	assertRoundTrip(t, wire, new(BodyUpdateStudioProjectV1StudioProjectsProjectIdPost))
}

func TestCreateProjectsRequestTargetAudienceVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant CreateProjectsRequestTargetAudience
		wire    string
	}
	cases := []enumCase{
		{variant: CreateProjectsRequestTargetAudienceChildren, wire: "\"children\""},
		{variant: CreateProjectsRequestTargetAudienceYoungAdult, wire: "\"young adult\""},
		{variant: CreateProjectsRequestTargetAudienceAdult, wire: "\"adult\""},
		{variant: CreateProjectsRequestTargetAudienceAllAges, wire: "\"all ages\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded CreateProjectsRequestTargetAudience
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

func TestCreateProjectsRequestFictionVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant CreateProjectsRequestFiction
		wire    string
	}
	cases := []enumCase{
		{variant: CreateProjectsRequestFictionFiction, wire: "\"fiction\""},
		{variant: CreateProjectsRequestFictionNonFiction, wire: "\"non-fiction\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded CreateProjectsRequestFiction
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

func TestCreateProjectsRequestApplyTextNormalizationVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant CreateProjectsRequestApplyTextNormalization
		wire    string
	}
	cases := []enumCase{
		{variant: CreateProjectsRequestApplyTextNormalizationAuto, wire: "\"auto\""},
		{variant: CreateProjectsRequestApplyTextNormalizationOn, wire: "\"on\""},
		{variant: CreateProjectsRequestApplyTextNormalizationOff, wire: "\"off\""},
		{variant: CreateProjectsRequestApplyTextNormalizationApplyEnglish, wire: "\"apply_english\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded CreateProjectsRequestApplyTextNormalization
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

func TestCreateProjectsRequestSourceTypeVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant CreateProjectsRequestSourceType
		wire    string
	}
	cases := []enumCase{
		{variant: CreateProjectsRequestSourceTypeBlank, wire: "\"blank\""},
		{variant: CreateProjectsRequestSourceTypeBook, wire: "\"book\""},
		{variant: CreateProjectsRequestSourceTypeArticle, wire: "\"article\""},
		{variant: CreateProjectsRequestSourceTypeGenfm, wire: "\"genfm\""},
		{variant: CreateProjectsRequestSourceTypeVideo, wire: "\"video\""},
		{variant: CreateProjectsRequestSourceTypeScreenplay, wire: "\"screenplay\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded CreateProjectsRequestSourceType
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
