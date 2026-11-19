package types

import (
	"testing"
)

func TestConvAiSecretLocatorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "secret_id": "value"
}`
	assertRoundTrip(t, wire, new(ConvAiSecretLocator))
}

func TestConvAiSecretLocatorRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "secret_id": "value"
}`
	assertRoundTrip(t, wire, new(ConvAiSecretLocator))
}

func TestConvAiSecretLocatorRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "secret_id": "value"
}`
	assertRoundTrip(t, wire, new(ConvAiSecretLocator))
}

func TestConvAiSecretLocatorRoundTrip4(t *testing.T) {
	t.Parallel()
	const wire = `{
  "secret_id": "value"
}`
	assertRoundTrip(t, wire, new(ConvAiSecretLocator))
}
