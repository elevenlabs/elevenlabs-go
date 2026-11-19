// Package webhooks manages workspace webhooks and verifies the webhooks that
// ElevenLabs sends to your server.
//
// ElevenLabs signs each webhook with an HMAC-SHA256 digest, sent in the
// ElevenLabs-Signature header as a compound value of the form
// "t=<unix_timestamp>,v0=<hex_digest>". The signed message is the timestamp and
// the raw request body joined with a period ("<timestamp>.<raw_body>").
// [ConstructEvent] (also reachable as client.Webhooks.ConstructEvent) checks that
// signature and returns the decoded event.
//
// The verifier is hand-written rather than generated, and listed in .fernignore:
// the Fern generator's built-in webhook verifier cannot parse a compound header
// that carries both the timestamp and the signature. It matches construct_event
// in the Python SDK and the other ElevenLabs SDKs.
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
	"time"
)

// Tolerance is how far in the past a webhook timestamp may be before
// [ConstructEvent] rejects it as a possible replay. Timestamps in the future are
// accepted, as in every other ElevenLabs SDK.
const Tolerance = 30 * time.Minute

// Errors returned by [ConstructEvent]. Match them with errors.Is. The messages
// are the ones every ElevenLabs SDK reports for the same input, so they are kept
// verbatim (including the leading capital letters).
var (
	// ErrMissingSignatureHeader is returned when the signature header is empty.
	ErrMissingSignatureHeader = errors.New("Missing signature header")
	// ErrSecretNotConfigured is returned when the webhook secret is empty.
	ErrSecretNotConfigured = errors.New("Webhook secret not configured")
	// ErrNoSignatureHash is returned when the header does not contain both a
	// t= timestamp and a v0= signature.
	ErrNoSignatureHash = errors.New("No signature hash found with expected scheme v0")
	// ErrTimestampOutsideTolerance is returned when the timestamp is more than
	// [Tolerance] in the past, or is not a number.
	ErrTimestampOutsideTolerance = errors.New("Timestamp outside the tolerance zone")
	// ErrSignatureMismatch is returned when the v0 signature does not match the
	// HMAC of the timestamp and raw body.
	ErrSignatureMismatch = errors.New("Signature hash does not match the expected signature hash for payload")
	// ErrInvalidPayload is returned, wrapping the JSON error, when the signature
	// is valid but the body is not a JSON object.
	ErrInvalidPayload = errors.New("Webhook payload is not a valid JSON object")
)

// ConstructEvent verifies an ElevenLabs webhook signature and returns the
// decoded event payload.
//
// rawBody must be the request body exactly as received: re-encoding a decoded
// body changes the bytes and breaks verification. sigHeader is the value of the
// ElevenLabs-Signature header, and secret is your webhook signing secret. The
// signature is compared in constant time.
//
// On failure it returns one of the sentinel errors above: the header or secret
// is missing, the header has no t=/v0= pair, the timestamp is more than
// [Tolerance] old, the signature does not match, or the body is not a JSON object.
func ConstructEvent(rawBody []byte, sigHeader, secret string) (map[string]any, error) {
	return constructEventAt(rawBody, sigHeader, secret, time.Now())
}

// ConstructEvent verifies an ElevenLabs webhook signature and returns the
// decoded event payload. It is the package-level [ConstructEvent], exposed on
// the client so it is reachable as client.Webhooks.ConstructEvent. It makes no
// API request.
func (c *Client) ConstructEvent(rawBody []byte, sigHeader, secret string) (map[string]any, error) {
	return ConstructEvent(rawBody, sigHeader, secret)
}

// constructEventAt is ConstructEvent with the current time injected, so the
// tolerance boundary can be tested.
func constructEventAt(rawBody []byte, sigHeader, secret string, now time.Time) (map[string]any, error) {
	if sigHeader == "" {
		return nil, ErrMissingSignatureHeader
	}
	if secret == "" {
		return nil, ErrSecretNotConfigured
	}

	// signature keeps its "v0=" prefix, so it is non-empty once found.
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

	// A timestamp that is not a number is treated as 0, and so rejected as
	// stale, matching the Rust, .NET, PHP and Ruby SDKs.
	requestTimestamp, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		requestTimestamp = 0
	}
	if requestTimestamp < now.Unix()-int64(Tolerance/time.Second) {
		return nil, ErrTimestampOutsideTolerance
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(timestamp))
	mac.Write([]byte("."))
	mac.Write(rawBody)
	expected := "v0=" + hex.EncodeToString(mac.Sum(nil))

	// Compare the whole "v0=<lowercase hex>" string, as the Python SDK does:
	// an uppercase digest does not match.
	if !hmac.Equal([]byte(expected), []byte(signature)) {
		return nil, ErrSignatureMismatch
	}

	var event map[string]any
	if err := json.Unmarshal(rawBody, &event); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInvalidPayload, err)
	}
	if event == nil {
		// The body was the JSON literal null.
		return nil, ErrInvalidPayload
	}
	return event, nil
}
