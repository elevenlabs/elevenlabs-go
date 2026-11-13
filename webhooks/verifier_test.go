package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"testing"
	"time"
)

// Shared cross-SDK golden vector. The timestamp is far in the future so it stays
// inside the (past-only) tolerance window forever; every ElevenLabs SDK must
// accept this exact (secret, header, body) triple byte-for-byte.
const (
	goldenSecret = "wsec_0123456789abcdef0123456789abcdef"
	goldenHeader = "t=4102444800,v0=01fc9faf23a5c155fbab279c4e0acb2ae059a567b14c10a19bd9f99c88315e99"
	goldenBody   = `{"type":"post_call_transcription","event_timestamp":1234567890,"data":{"agent_id":"abc123"}}`
)

// sign builds a valid header for the given timestamp and body.
func sign(t *testing.T, timestamp string, body, secret string) string {
	t.Helper()
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(timestamp + "." + body))
	return "t=" + timestamp + ",v0=" + hex.EncodeToString(mac.Sum(nil))
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

func TestCurrentTimestampVerifies(t *testing.T) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	body := `{"type":"ping"}`
	header := sign(t, timestamp, body, goldenSecret)

	event, err := ConstructEvent([]byte(body), header, goldenSecret)
	if err != nil {
		t.Fatalf("expected current-timestamp event to verify, got error: %v", err)
	}
	if got := event["type"]; got != "ping" {
		t.Errorf("type = %v, want ping", got)
	}
}

func TestTamperedSignatureIsRejected(t *testing.T) {
	header := "t=4102444800,v0=0000000000000000000000000000000000000000000000000000000000000000"

	_, err := ConstructEvent([]byte(goldenBody), header, goldenSecret)
	if !errors.Is(err, ErrSignatureMismatch) {
		t.Fatalf("expected ErrSignatureMismatch, got %v", err)
	}
}

func TestTamperedBodyIsRejected(t *testing.T) {
	_, err := ConstructEvent([]byte(goldenBody+" "), goldenHeader, goldenSecret)
	if !errors.Is(err, ErrSignatureMismatch) {
		t.Fatalf("expected ErrSignatureMismatch, got %v", err)
	}
}

func TestOldTimestampIsRejected(t *testing.T) {
	timestamp := strconv.FormatInt(time.Now().Add(-31*time.Minute).Unix(), 10)
	body := `{"type":"ping"}`
	header := sign(t, timestamp, body, goldenSecret)

	_, err := ConstructEvent([]byte(body), header, goldenSecret)
	if !errors.Is(err, ErrTimestampOutsideTolerance) {
		t.Fatalf("expected ErrTimestampOutsideTolerance, got %v", err)
	}
}

func TestMissingHeaderIsRejected(t *testing.T) {
	_, err := ConstructEvent([]byte(goldenBody), "", goldenSecret)
	if !errors.Is(err, ErrMissingSignatureHeader) {
		t.Fatalf("expected ErrMissingSignatureHeader, got %v", err)
	}
}

func TestMissingSecretIsRejected(t *testing.T) {
	_, err := ConstructEvent([]byte(goldenBody), goldenHeader, "")
	if !errors.Is(err, ErrSecretNotConfigured) {
		t.Fatalf("expected ErrSecretNotConfigured, got %v", err)
	}
}

func TestHeaderWithoutV0SchemeIsRejected(t *testing.T) {
	_, err := ConstructEvent([]byte(goldenBody), "t=4102444800", goldenSecret)
	if !errors.Is(err, ErrNoSignatureHash) {
		t.Fatalf("expected ErrNoSignatureHash, got %v", err)
	}
}
