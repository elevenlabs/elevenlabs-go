// Hand-written (listed in .fernignore).

package speechengine

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"testing"
	"time"
)

// Shared cross-SDK golden vector: this exact (API key, token) pair must verify in every
// ElevenLabs SDK.
const (
	goldenAPIKey = "sk_0123456789abcdef0123456789abcdef"
	goldenToken  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2FwaS5lbGV2ZW5sYWJzLmlvL2NvbnZhaS9zcGVlY2gtZW5naW5lIiwic3ViIjoiY29udmFpX3NwZWVjaF9lbmdpbmVfdXBzdHJlYW0iLCJpYXQiOjEyMzQ1Njc4OTAsImV4cCI6NDEwMjQ0NDgwMH0.UObM4hx1W8NvGViU3aqPv1-IEXhmXgWDgoS9RZZWILk"
)

// signToken signs payload the way the API does, for exercising the claim checks.
func signToken(t *testing.T, apiKey string, payload map[string]any) string {
	t.Helper()
	enc := base64.RawURLEncoding
	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}
	signingInput := enc.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`)) + "." + enc.EncodeToString(body)
	secret := sha256.Sum256([]byte(normalizeAPIKeyForSigning(apiKey)))
	mac := hmac.New(sha256.New, secret[:])
	mac.Write([]byte(signingInput))
	return signingInput + "." + enc.EncodeToString(mac.Sum(nil))
}

func claims(iat, exp int64) map[string]any {
	return map[string]any{"iss": jwtIssuer, "sub": jwtSubject, "iat": iat, "exp": exp}
}

func expectJWTError(t *testing.T, err error, fragment string) {
	t.Helper()
	var jwtErr *JWTError
	if !errors.As(err, &jwtErr) {
		t.Fatalf("expected *JWTError, got %v", err)
	}
	if !strings.Contains(err.Error(), fragment) {
		t.Fatalf("expected error containing %q, got %q", fragment, err.Error())
	}
}

func TestGoldenVectorVerifies(t *testing.T) {
	got, err := VerifySpeechEngineJWT(goldenToken, goldenAPIKey)
	if err != nil {
		t.Fatalf("golden vector must verify: %v", err)
	}
	if got["iss"] != jwtIssuer || got["sub"] != jwtSubject {
		t.Fatalf("unexpected claims: %v", got)
	}
	if got["iat"] != float64(1234567890) || got["exp"] != float64(4102444800) {
		t.Fatalf("unexpected iat/exp: %v", got)
	}
}

func TestResidencySuffixIsStripped(t *testing.T) {
	if _, err := VerifySpeechEngineJWT(goldenToken, goldenAPIKey+"_residency_us"); err != nil {
		t.Fatalf("residency key must hash to the base-key secret: %v", err)
	}
}

func TestBearerPrefixIsTolerated(t *testing.T) {
	if _, err := VerifySpeechEngineJWT("Bearer "+goldenToken, goldenAPIKey); err != nil {
		t.Fatal(err)
	}
	if _, err := VerifySpeechEngineJWT("  bearer "+goldenToken+" ", goldenAPIKey); err != nil {
		t.Fatal(err)
	}
}

func TestPaddedSignatureIsTolerated(t *testing.T) {
	if _, err := VerifySpeechEngineJWT(goldenToken+"=", goldenAPIKey); err != nil {
		t.Fatal(err)
	}
}

func TestEmptyAPIKeysAreRejected(t *testing.T) {
	// A key that normalizes to "" would hash to SHA-256(""), a publicly known secret.
	for _, key := range []string{"", "   ", "_residency_us"} {
		_, err := VerifySpeechEngineJWT(goldenToken, key)
		expectJWTError(t, err, "No API key configured")
	}
}

func TestWrongAPIKeyIsASignatureMismatch(t *testing.T) {
	_, err := VerifySpeechEngineJWT(goldenToken, "sk_wrongwrongwrongwrong")
	expectJWTError(t, err, "signature mismatch")
}

func TestTamperedPayloadIsRejected(t *testing.T) {
	parts := strings.Split(goldenToken, ".")
	tampered := parts[0] + "." + parts[1] + "x." + parts[2]
	if _, err := VerifySpeechEngineJWT(tampered, goldenAPIKey); err == nil {
		t.Fatal("tampered token must not verify")
	}
}

func TestMalformedTokensAreRejected(t *testing.T) {
	_, err := VerifySpeechEngineJWT("a.b", goldenAPIKey)
	expectJWTError(t, err, "expected 3 parts")
	_, err = VerifySpeechEngineJWT("a.!!!.c", goldenAPIKey)
	expectJWTError(t, err, "failed to decode payload")
}

func TestClaimChecks(t *testing.T) {
	now := time.Now().Unix()
	cases := []struct {
		name     string
		payload  map[string]any
		fragment string
	}{
		{"expired", claims(now-7200, now-3600), "has expired"},
		{"future iat", claims(now+3600, now+7200), "iat is in the future"},
		{"issuer", map[string]any{"iss": "https://evil.example.com", "sub": jwtSubject, "iat": now, "exp": now + 60}, "unexpected issuer"},
		{"subject", map[string]any{"iss": jwtIssuer, "sub": "something_else", "iat": now, "exp": now + 60}, "unexpected subject"},
		{"missing exp", map[string]any{"iss": jwtIssuer, "sub": jwtSubject, "iat": now}, "missing exp claim"},
		{"missing iat", map[string]any{"iss": jwtIssuer, "sub": jwtSubject, "exp": now + 60}, "missing iat claim"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := VerifySpeechEngineJWT(signToken(t, goldenAPIKey, tc.payload), goldenAPIKey)
			expectJWTError(t, err, tc.fragment)
		})
	}
}

func TestLeewayAllowsAJustExpiredToken(t *testing.T) {
	now := time.Now().Unix()
	if _, err := VerifySpeechEngineJWT(signToken(t, goldenAPIKey, claims(now-120, now-30)), goldenAPIKey); err != nil {
		t.Fatalf("a token expired within the leeway must verify: %v", err)
	}
}

func TestVerifyRequest(t *testing.T) {
	resource := NewSpeechEngineResource("seng_1", goldenAPIKey, nil)
	req, _ := http.NewRequest(http.MethodGet, "http://localhost/ws", nil)
	if err := resource.VerifyRequest(req); err == nil {
		t.Fatal("a request without the header must be rejected")
	}
	// Header names are case-insensitive.
	req.Header.Set("X-ELEVENLABS-SPEECH-ENGINE-AUTHORIZATION", goldenToken)
	if err := resource.VerifyRequest(req); err != nil {
		t.Fatalf("golden token must verify: %v", err)
	}
	if err := NewSpeechEngineResource("seng_1", "", nil).VerifyRequest(req); err == nil {
		t.Fatal("a resource without an API key must reject every request")
	}
}
