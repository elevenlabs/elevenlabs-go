// Package webhooks provides helpers for working with ElevenLabs webhooks,
// including signature verification.
//
// ElevenLabs signs each webhook with an HMAC-SHA256 digest, sent in the
// ElevenLabs-Signature header as a compound value of the form
// "t=<unix_timestamp>,v0=<hex_digest>". The signed message is the timestamp and
// the raw request body joined with a period ("<timestamp>.<raw_body>").
//
// ConstructEvent is hand-written rather than generated: the Fern generator's
// built-in webhook verifier cannot parse a compound header that carries both the
// timestamp and the signature. The file is listed in .fernignore so regeneration
// leaves it in place.
package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"
)

// toleranceSeconds is how far in the past a webhook timestamp may be before it
// is rejected as a possible replay. Future timestamps are always accepted.
const toleranceSeconds = 30 * 60

// Errors returned by ConstructEvent. They are sentinel values so callers can
// match them with errors.Is; the messages are shared verbatim across every
// ElevenLabs SDK.
var (
	// ErrMissingSignatureHeader is returned when the signature header is empty.
	ErrMissingSignatureHeader = errors.New("Missing signature header")
	// ErrSecretNotConfigured is returned when no webhook secret is provided.
	ErrSecretNotConfigured = errors.New("Webhook secret not configured")
	// ErrNoSignatureHash is returned when the header has no usable t=/v0= parts.
	ErrNoSignatureHash = errors.New("No signature hash found with expected scheme v0")
	// ErrTimestampOutsideTolerance is returned when the timestamp is too old.
	ErrTimestampOutsideTolerance = errors.New("Timestamp outside the tolerance zone")
	// ErrSignatureMismatch is returned when the computed digest does not match.
	ErrSignatureMismatch = errors.New("Signature hash does not match the expected signature hash for payload")
)

// ConstructEvent verifies an ElevenLabs webhook signature and returns the
// decoded event payload.
//
// rawBody must be the raw request body exactly as received: re-encoding a decoded
// body changes the bytes and breaks verification. sigHeader is the value of the
// ElevenLabs-Signature header, and secret is your webhook signing secret.
//
// It returns one of the sentinel errors above if the header or secret is missing,
// the timestamp is outside the tolerance window, or the signature does not match.
func ConstructEvent(rawBody []byte, sigHeader, secret string) (map[string]any, error) {
	if sigHeader == "" {
		return nil, ErrMissingSignatureHeader
	}
	if secret == "" {
		return nil, ErrSecretNotConfigured
	}

	var timestamp, signature string
	for _, part := range strings.Split(sigHeader, ",") {
		switch {
		case strings.HasPrefix(part, "t="):
			timestamp = part[len("t="):]
		case strings.HasPrefix(part, "v0="):
			signature = part
		}
	}

	if timestamp == "" || signature == "" {
		return nil, ErrNoSignatureHash
	}

	ts, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return nil, ErrNoSignatureHash
	}

	// Past-only tolerance: reject timestamps older than the window, allow future ones.
	reqTimestampMs := ts * 1000
	toleranceMs := time.Now().UnixMilli() - toleranceSeconds*1000
	if reqTimestampMs < toleranceMs {
		return nil, ErrTimestampOutsideTolerance
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(timestamp))
	mac.Write([]byte("."))
	mac.Write(rawBody)
	digest := "v0=" + hex.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(digest), []byte(signature)) {
		return nil, ErrSignatureMismatch
	}

	var event map[string]any
	if err := json.Unmarshal(rawBody, &event); err != nil {
		return nil, err
	}
	return event, nil
}
