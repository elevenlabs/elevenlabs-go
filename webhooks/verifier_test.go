package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	core "github.com/elevenlabs/elevenlabs-go/core"
)

// Hand-written (listed in .fernignore). Mirrors the shared cross-SDK suite (see
// the Rust and .NET ports) case for case.

// Shared cross-SDK golden vector. The timestamp is far in the future so it stays
// inside the (past-only) tolerance window forever; every ElevenLabs SDK must
// accept this exact (secret, header, body) triple.
const (
	goldenSecret = "wsec_0123456789abcdef0123456789abcdef"
	goldenHeader = "t=4102444800,v0=01fc9faf23a5c155fbab279c4e0acb2ae059a567b14c10a19bd9f99c88315e99"
	goldenBody   = `{"type":"post_call_transcription","event_timestamp":1234567890,"data":{"agent_id":"abc123"}}`
)

// sign builds a valid header for the given timestamp and body.
func sign(timestamp, body string) string {
	mac := hmac.New(sha256.New, []byte(goldenSecret))
	mac.Write([]byte(timestamp + "." + body))
	return "t=" + timestamp + ",v0=" + hex.EncodeToString(mac.Sum(nil))
}

func assertRejected(t *testing.T, err, want error) {
	t.Helper()
	if !errors.Is(err, want) {
		t.Fatalf("expected %q, got %v", want, err)
	}
}

func TestGoldenVectorVerifiesAndDecodes(t *testing.T) {
	event, err := ConstructEvent([]byte(goldenBody), goldenHeader, goldenSecret)
	if err != nil {
		t.Fatalf("expected golden vector to verify, got error: %v", err)
	}
	if got := event["type"]; got != "post_call_transcription" {
		t.Errorf("type = %v, want post_call_transcription", got)
	}
	// JSON numbers decode to float64 in map[string]any.
	if got := event["event_timestamp"]; got != float64(1234567890) {
		t.Errorf("event_timestamp = %v, want 1234567890", got)
	}
	data, ok := event["data"].(map[string]any)
	if !ok {
		t.Fatalf("data is not an object: %T", event["data"])
	}
	if got := data["agent_id"]; got != "abc123" {
		t.Errorf("data.agent_id = %v, want abc123", got)
	}
}

func TestClientConstructEventDelegatesToVerifier(t *testing.T) {
	client := NewClient(&core.RequestOptions{})

	event, err := client.ConstructEvent([]byte(goldenBody), goldenHeader, goldenSecret)
	if err != nil {
		t.Fatalf("expected golden vector to verify, got error: %v", err)
	}
	if got := event["type"]; got != "post_call_transcription" {
		t.Errorf("type = %v, want post_call_transcription", got)
	}

	_, err = client.ConstructEvent([]byte(goldenBody), "", goldenSecret)
	assertRejected(t, err, ErrMissingSignatureHeader)
}

func TestCurrentTimestampVerifies(t *testing.T) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	body := `{"type":"ping"}`

	event, err := ConstructEvent([]byte(body), sign(timestamp, body), goldenSecret)
	if err != nil {
		t.Fatalf("expected current-timestamp event to verify, got error: %v", err)
	}
	if got := event["type"]; got != "ping" {
		t.Errorf("type = %v, want ping", got)
	}
}

func TestFutureTimestampVerifies(t *testing.T) {
	timestamp := strconv.FormatInt(time.Now().Add(24*time.Hour).Unix(), 10)
	body := `{"type":"ping"}`

	if _, err := ConstructEvent([]byte(body), sign(timestamp, body), goldenSecret); err != nil {
		t.Fatalf("expected future timestamp to verify, got error: %v", err)
	}
}

func TestTamperedSignatureIsRejected(t *testing.T) {
	header := "t=4102444800,v0=" + strings.Repeat("0", 64)

	_, err := ConstructEvent([]byte(goldenBody), header, goldenSecret)
	assertRejected(t, err, ErrSignatureMismatch)
	if !strings.Contains(err.Error(), "Signature hash does not match") {
		t.Errorf("message = %q", err.Error())
	}
}

func TestTamperedBodyIsRejected(t *testing.T) {
	_, err := ConstructEvent([]byte(goldenBody+" "), goldenHeader, goldenSecret)
	assertRejected(t, err, ErrSignatureMismatch)
}

func TestWrongSecretIsRejected(t *testing.T) {
	_, err := ConstructEvent([]byte(goldenBody), goldenHeader, "wsec_not_the_right_secret")
	assertRejected(t, err, ErrSignatureMismatch)
}

func TestUppercaseSignatureIsRejected(t *testing.T) {
	timestamp, digest, _ := strings.Cut(goldenHeader, ",v0=")
	header := timestamp + ",v0=" + strings.ToUpper(digest)

	_, err := ConstructEvent([]byte(goldenBody), header, goldenSecret)
	assertRejected(t, err, ErrSignatureMismatch)
}

func TestOldTimestampIsRejected(t *testing.T) {
	timestamp := strconv.FormatInt(time.Now().Add(-31*time.Minute).Unix(), 10)
	body := `{"type":"ping"}`

	_, err := ConstructEvent([]byte(body), sign(timestamp, body), goldenSecret)
	assertRejected(t, err, ErrTimestampOutsideTolerance)
	if got := err.Error(); got != "Timestamp outside the tolerance zone" {
		t.Errorf("message = %q", got)
	}
}

func TestToleranceBoundaryIsInclusive(t *testing.T) {
	now := time.Unix(1_700_000_000, 0)
	body := `{"type":"ping"}`

	atEdge := strconv.FormatInt(now.Unix()-30*60, 10)
	if _, err := constructEventAt([]byte(body), sign(atEdge, body), goldenSecret, now); err != nil {
		t.Fatalf("expected a timestamp exactly at the boundary to verify, got error: %v", err)
	}

	pastEdge := strconv.FormatInt(now.Unix()-30*60-1, 10)
	_, err := constructEventAt([]byte(body), sign(pastEdge, body), goldenSecret, now)
	assertRejected(t, err, ErrTimestampOutsideTolerance)
}

func TestNonNumericTimestampIsRejectedAsStale(t *testing.T) {
	body := `{"type":"ping"}`

	_, err := ConstructEvent([]byte(body), sign("abc", body), goldenSecret)
	assertRejected(t, err, ErrTimestampOutsideTolerance)
}

func TestMissingHeaderIsRejected(t *testing.T) {
	_, err := ConstructEvent([]byte(goldenBody), "", goldenSecret)
	assertRejected(t, err, ErrMissingSignatureHeader)
	if got := err.Error(); got != "Missing signature header" {
		t.Errorf("message = %q", got)
	}
}

func TestMissingSecretIsRejected(t *testing.T) {
	_, err := ConstructEvent([]byte(goldenBody), goldenHeader, "")
	assertRejected(t, err, ErrSecretNotConfigured)
	if got := err.Error(); got != "Webhook secret not configured" {
		t.Errorf("message = %q", got)
	}
}

func TestHeaderWithoutV0SchemeIsRejected(t *testing.T) {
	_, err := ConstructEvent([]byte(goldenBody), "t=4102444800", goldenSecret)
	assertRejected(t, err, ErrNoSignatureHash)
	if got := err.Error(); got != "No signature hash found with expected scheme v0" {
		t.Errorf("message = %q", got)
	}
}

func TestHeaderWithoutTimestampIsRejected(t *testing.T) {
	header := "v0=01fc9faf23a5c155fbab279c4e0acb2ae059a567b14c10a19bd9f99c88315e99"

	_, err := ConstructEvent([]byte(goldenBody), header, goldenSecret)
	assertRejected(t, err, ErrNoSignatureHash)
}

func TestValidSignatureOverNonJSONBodyIsAnInvalidPayload(t *testing.T) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	for _, body := range []string{"not json", "null", `["not","an","object"]`} {
		_, err := ConstructEvent([]byte(body), sign(timestamp, body), goldenSecret)
		assertRejected(t, err, ErrInvalidPayload)
	}

	// The JSON error is wrapped too, so callers can inspect it.
	_, err := ConstructEvent([]byte("not json"), sign(timestamp, "not json"), goldenSecret)
	var syntaxErr *json.SyntaxError
	if !errors.As(err, &syntaxErr) {
		t.Errorf("expected a wrapped *json.SyntaxError, got %v", err)
	}
}

func ExampleConstructEvent() {
	// In an http.Handler: rawBody, _ := io.ReadAll(r.Body)
	// and sigHeader := r.Header.Get("ElevenLabs-Signature").
	rawBody := []byte(goldenBody)
	sigHeader := goldenHeader

	event, err := ConstructEvent(rawBody, sigHeader, goldenSecret)
	if err != nil {
		// Missing or invalid signature, or a timestamp older than 30 minutes.
		fmt.Println("rejected:", err)
		return
	}
	fmt.Println(event["type"])
	// Output: post_call_transcription
}
