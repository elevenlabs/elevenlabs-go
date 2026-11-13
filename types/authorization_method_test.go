package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAuthorizationMethodVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant AuthorizationMethod
		wire    string
	}
	cases := []enumCase{
		{variant: AuthorizationMethodInvalid, wire: "\"invalid\""},
		{variant: AuthorizationMethodPublic, wire: "\"public\""},
		{variant: AuthorizationMethodAuthorizationHeader, wire: "\"authorization_header\""},
		{variant: AuthorizationMethodSignedUrl, wire: "\"signed_url\""},
		{variant: AuthorizationMethodShareableLink, wire: "\"shareable_link\""},
		{variant: AuthorizationMethodLivekitToken, wire: "\"livekit_token\""},
		{variant: AuthorizationMethodLivekitTokenWebsite, wire: "\"livekit_token_website\""},
		{variant: AuthorizationMethodGenesysApiKey, wire: "\"genesys_api_key\""},
		{variant: AuthorizationMethodAvayaApiKey, wire: "\"avaya_api_key\""},
		{variant: AuthorizationMethodAudiocodesApiKey, wire: "\"audiocodes_api_key\""},
		{variant: AuthorizationMethodWhatsapp, wire: "\"whatsapp\""},
		{variant: AuthorizationMethodSms, wire: "\"sms\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded AuthorizationMethod
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
