package types

import (
	"testing"
)

func TestQuoteInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "amount_usd": 9.99
}`
	assertRoundTrip(t, wire, new(QuoteInfo))
}

func TestQuoteInfoRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "amount_usd": 9.99
}`
	assertRoundTrip(t, wire, new(QuoteInfo))
}
