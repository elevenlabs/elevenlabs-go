package types

import (
	"testing"
)

func TestListResponseMetaRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "total": 123,
  "page": 123,
  "page_size": 123
}`
	assertRoundTrip(t, wire, new(ListResponseMeta))
}

func TestListResponseMetaRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "total": 123,
  "page": 123,
  "page_size": 123
}`
	assertRoundTrip(t, wire, new(ListResponseMeta))
}

func TestListResponseMetaRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "total": 123,
  "page": 123,
  "page_size": 123
}`
	assertRoundTrip(t, wire, new(ListResponseMeta))
}

func TestListResponseMetaRoundTrip4(t *testing.T) {
	t.Parallel()
	const wire = `{
  "total": 123,
  "page": 123,
  "page_size": 123
}`
	assertRoundTrip(t, wire, new(ListResponseMeta))
}
