// Hand-written (listed in .fernignore): the generator only emits the client that dials out to
// /speech-engine/upstream and has no concept of the inbound server handshake.

package speechengine

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// AuthorizationHeader is the request header that carries the JWT the ElevenLabs Speech Engine
// API presents when it opens a WebSocket connection to your server.
const AuthorizationHeader = "X-Elevenlabs-Speech-Engine-Authorization"

const (
	jwtIssuer  = "https://api.elevenlabs.io/convai/speech-engine"
	jwtSubject = "convai_speech_engine_upstream"
	// Clock skew, in seconds, allowed on the exp and iat claims.
	jwtLeewaySeconds = 60
)

// Data-residency keys are signed with the base key, so this suffix is stripped before hashing.
var residencyKeySuffix = regexp.MustCompile(`_residency_[a-z0-9]+$`)

// JWTError reports why an inbound Speech Engine token was rejected. The messages match the
// Python, TypeScript, PHP, .NET and Rust SDKs.
type JWTError struct {
	Message string
}

func (e *JWTError) Error() string { return e.Message }

func jwtError(message string) error { return &JWTError{Message: message} }

// VerifySpeechEngineJWT verifies the token the ElevenLabs Speech Engine API sends in the
// X-Elevenlabs-Speech-Engine-Authorization header and returns its decoded claims.
//
// value is the header value, with or without a "Bearer " prefix. apiKey is the API key the
// engine was created with. The token is an HS256 JWT whose HMAC secret is the SHA-256 hash of
// the API key, with any "_residency_<region>" suffix removed first.
//
// It returns a *JWTError when the key is empty, the token is malformed, the signature does not
// match, or the issuer, subject, exp or iat claim is invalid.
func VerifySpeechEngineJWT(value string, apiKey string) (map[string]any, error) {
	// Reject an empty key before deriving the secret: a key that normalizes to "" would hash to
	// SHA-256(""), a publicly known secret that would accept forged tokens.
	normalizedKey := normalizeAPIKeyForSigning(apiKey)
	if normalizedKey == "" {
		return nil, jwtError("No API key configured for Speech Engine verification")
	}

	token := strings.TrimSpace(value)
	if len(token) >= 7 && strings.EqualFold(token[:7], "bearer ") {
		token = strings.TrimSpace(token[7:])
	}

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, jwtError("Invalid JWT: expected 3 parts")
	}
	headerB64, payloadB64, signatureB64 := parts[0], parts[1], parts[2]

	payloadBytes, err := base64URLDecode(payloadB64)
	if err != nil {
		return nil, jwtError("Invalid JWT: failed to decode payload")
	}
	var payload map[string]any
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return nil, jwtError("Invalid JWT: failed to decode payload")
	}
	if payload == nil {
		return nil, jwtError("Invalid JWT: payload is not an object")
	}

	secret := sha256.Sum256([]byte(normalizedKey))
	mac := hmac.New(sha256.New, secret[:])
	mac.Write([]byte(headerB64 + "." + payloadB64))
	expected := mac.Sum(nil)

	actual, err := base64URLDecode(signatureB64)
	if err != nil || !hmac.Equal(expected, actual) {
		return nil, jwtError("Invalid JWT: signature mismatch — make sure the API key used by your Speech Engine server matches the one used to create the engine.")
	}

	if iss, _ := payload["iss"].(string); iss != jwtIssuer {
		return nil, jwtError("Invalid JWT: unexpected issuer")
	}
	if sub, _ := payload["sub"].(string); sub != jwtSubject {
		return nil, jwtError("Invalid JWT: unexpected subject")
	}
	exp, ok := payload["exp"].(float64)
	if !ok {
		return nil, jwtError("Invalid JWT: missing exp claim")
	}
	iat, ok := payload["iat"].(float64)
	if !ok {
		return nil, jwtError("Invalid JWT: missing iat claim")
	}
	now := float64(time.Now().Unix())
	if exp+jwtLeewaySeconds < now {
		return nil, jwtError("Invalid JWT: token has expired")
	}
	if iat-jwtLeewaySeconds > now {
		return nil, jwtError("Invalid JWT: iat is in the future")
	}
	return payload, nil
}

// verifyRequestHeader checks the authorization header of an inbound request.
func verifyRequestHeader(header http.Header, apiKey string) error {
	if strings.TrimSpace(apiKey) == "" {
		return jwtError("No API key configured for Speech Engine verification")
	}
	value := header.Get(AuthorizationHeader)
	if value == "" {
		return jwtError("missing " + AuthorizationHeader + " header")
	}
	_, err := VerifySpeechEngineJWT(value, apiKey)
	return err
}

func normalizeAPIKeyForSigning(apiKey string) string {
	return residencyKeySuffix.ReplaceAllString(strings.TrimSpace(apiKey), "")
}

// base64URLDecode decodes base64url (RFC 7515), with or without padding.
func base64URLDecode(input string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(strings.TrimRight(input, "="))
}
