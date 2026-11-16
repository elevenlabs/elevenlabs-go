package types

import (
	"testing"
)

func TestOmitSchemaOverrideRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "omit"
}`
	assertRoundTrip(t, wire, new(OmitSchemaOverride))
}

func TestOmitSchemaOverrideRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "omit"
}`
	assertRoundTrip(t, wire, new(OmitSchemaOverride))
}

func TestOmitSchemaOverrideRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "omit"
}`
	assertRoundTrip(t, wire, new(OmitSchemaOverride))
}
