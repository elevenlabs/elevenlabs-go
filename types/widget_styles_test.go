package types

import (
	"testing"
)

func TestWidgetStylesRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "base": "value",
  "base_hover": "value",
  "base_active": "value",
  "base_border": "value",
  "base_subtle": "value",
  "base_primary": "value",
  "base_error": "value",
  "accent": "value",
  "accent_hover": "value",
  "accent_active": "value",
  "accent_border": "value",
  "accent_subtle": "value",
  "accent_primary": "value",
  "overlay_padding": 9.99,
  "button_radius": 9.99,
  "input_radius": 9.99,
  "bubble_radius": 9.99,
  "sheet_radius": 9.99,
  "compact_sheet_radius": 9.99,
  "dropdown_sheet_radius": 9.99
}`
	assertRoundTrip(t, wire, new(WidgetStyles))
}

func TestWidgetStylesRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "base": "value",
  "base_hover": "value",
  "base_active": "value",
  "base_border": "value",
  "base_subtle": "value",
  "base_primary": "value",
  "base_error": "value",
  "accent": "value",
  "accent_hover": "value",
  "accent_active": "value",
  "accent_border": "value",
  "accent_subtle": "value",
  "accent_primary": "value",
  "overlay_padding": 9.99,
  "button_radius": 9.99,
  "input_radius": 9.99,
  "bubble_radius": 9.99,
  "sheet_radius": 9.99,
  "compact_sheet_radius": 9.99,
  "dropdown_sheet_radius": 9.99
}`
	assertRoundTrip(t, wire, new(WidgetStyles))
}
