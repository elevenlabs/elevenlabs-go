package types

import (
	"testing"
)

func TestResourceAccessInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_creator": true,
  "creator_name": "value",
  "creator_email": "value",
  "role": "admin",
  "anonymous_access_level_override": "admin",
  "access_source": "creator"
}`
	assertRoundTrip(t, wire, new(ResourceAccessInfo))
}

func TestResourceAccessInfoRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_creator": true,
  "creator_name": "value",
  "creator_email": "value",
  "role": "admin",
  "anonymous_access_level_override": "admin",
  "access_source": "creator"
}`
	assertRoundTrip(t, wire, new(ResourceAccessInfo))
}

func TestResourceAccessInfoRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_creator": true,
  "creator_name": "value",
  "creator_email": "value",
  "role": "admin",
  "anonymous_access_level_override": "admin",
  "access_source": "creator"
}`
	assertRoundTrip(t, wire, new(ResourceAccessInfo))
}

func TestResourceAccessInfoRoundTrip4(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_creator": true,
  "creator_name": "value",
  "creator_email": "value",
  "role": "admin",
  "anonymous_access_level_override": "admin",
  "access_source": "creator"
}`
	assertRoundTrip(t, wire, new(ResourceAccessInfo))
}

func TestResourceAccessInfoRoundTrip5(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_creator": true,
  "creator_name": "value",
  "creator_email": "value",
  "role": "admin",
  "anonymous_access_level_override": "admin",
  "access_source": "creator"
}`
	assertRoundTrip(t, wire, new(ResourceAccessInfo))
}

func TestResourceAccessInfoRoundTrip6(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_creator": true,
  "creator_name": "value",
  "creator_email": "value",
  "role": "admin",
  "anonymous_access_level_override": "admin",
  "access_source": "creator"
}`
	assertRoundTrip(t, wire, new(ResourceAccessInfo))
}

func TestResourceAccessInfoRoundTrip7(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_creator": true,
  "creator_name": "value",
  "creator_email": "value",
  "role": "admin",
  "anonymous_access_level_override": "admin",
  "access_source": "creator"
}`
	assertRoundTrip(t, wire, new(ResourceAccessInfo))
}

func TestResourceAccessInfoRoundTrip8(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_creator": true,
  "creator_name": "value",
  "creator_email": "value",
  "role": "admin",
  "anonymous_access_level_override": "admin",
  "access_source": "creator"
}`
	assertRoundTrip(t, wire, new(ResourceAccessInfo))
}

func TestResourceAccessInfoRoundTrip9(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_creator": true,
  "creator_name": "value",
  "creator_email": "value",
  "role": "admin",
  "anonymous_access_level_override": "admin",
  "access_source": "creator"
}`
	assertRoundTrip(t, wire, new(ResourceAccessInfo))
}

func TestResourceAccessInfoRoundTrip10(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_creator": true,
  "creator_name": "value",
  "creator_email": "value",
  "role": "admin",
  "anonymous_access_level_override": "admin",
  "access_source": "creator"
}`
	assertRoundTrip(t, wire, new(ResourceAccessInfo))
}

func TestResourceAccessInfoRoundTrip11(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_creator": true,
  "creator_name": "value",
  "creator_email": "value",
  "role": "admin",
  "anonymous_access_level_override": "admin",
  "access_source": "creator"
}`
	assertRoundTrip(t, wire, new(ResourceAccessInfo))
}

func TestResourceAccessInfoRoundTrip12(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_creator": true,
  "creator_name": "value",
  "creator_email": "value",
  "role": "admin",
  "anonymous_access_level_override": "admin",
  "access_source": "creator"
}`
	assertRoundTrip(t, wire, new(ResourceAccessInfo))
}

func TestResourceAccessInfoRoundTrip13(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_creator": true,
  "creator_name": "value",
  "creator_email": "value",
  "role": "admin",
  "anonymous_access_level_override": "admin",
  "access_source": "creator"
}`
	assertRoundTrip(t, wire, new(ResourceAccessInfo))
}
