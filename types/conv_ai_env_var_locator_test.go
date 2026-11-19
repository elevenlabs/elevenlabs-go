package types

import (
	"testing"
)

func TestConvAiEnvVarLocatorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "env_var_label": "value"
}`
	assertRoundTrip(t, wire, new(ConvAiEnvVarLocator))
}

func TestConvAiEnvVarLocatorRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "env_var_label": "value"
}`
	assertRoundTrip(t, wire, new(ConvAiEnvVarLocator))
}

func TestConvAiEnvVarLocatorRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "env_var_label": "value"
}`
	assertRoundTrip(t, wire, new(ConvAiEnvVarLocator))
}
