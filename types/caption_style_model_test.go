package types

import (
	"testing"
)

func TestCaptionStyleModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "template": {
    "key": "value",
    "label": "value",
    "requires_high_fps": true
  },
  "text_font": "value",
  "text_scale": 9.99,
  "text_color": "value",
  "text_align": "start",
  "text_style": "normal",
  "text_weight": "normal",
  "text_transform": "none",
  "text_blend_mode": "normal",
  "text_shadow": {
    "enabled": true,
    "color": "value",
    "opacity": 9.99,
    "blur": 9.99,
    "offset_x": 9.99,
    "offset_y": 9.99
  },
  "text_outline": {
    "enabled": true,
    "color": "value",
    "opacity": 9.99,
    "width": 9.99
  },
  "background_enabled": true,
  "background_color": "value",
  "background_opacity": 9.99,
  "background_blur": 9.99,
  "background_border_radius": 9.99,
  "word_highlights_enabled": true,
  "word_highlights_color": "value",
  "word_highlights_background_color": "value",
  "word_highlights_opacity": 9.99,
  "word_highlights_border_radius": 9.99,
  "word_highlights_blur": 9.99,
  "section_animation": {
    "enter_type": "none",
    "exit_type": "none"
  },
  "word_animation": {
    "enter_type": "none",
    "exit_type": "none"
  },
  "character_animation": {
    "enter_type": "none",
    "exit_type": "none"
  },
  "cursor_enabled": true,
  "width_pct": 9.99,
  "horizontal_placement": {
    "align": "left",
    "translate_pct": 9.99
  },
  "vertical_placement": {
    "align": "top",
    "translate_pct": 9.99
  },
  "auto_break_enabled": true,
  "max_lines_per_section": 123,
  "max_words_per_line": 123
}`
	assertRoundTrip(t, wire, new(CaptionStyleModel))
}

func TestCaptionStyleModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "template": {
    "key": "value",
    "label": "value",
    "requires_high_fps": true
  },
  "text_font": "value",
  "text_scale": 9.99,
  "text_color": "value",
  "text_align": "start",
  "text_style": "normal",
  "text_weight": "normal",
  "text_transform": "none",
  "text_blend_mode": "normal",
  "text_shadow": {
    "enabled": true,
    "color": "value",
    "opacity": 9.99,
    "blur": 9.99,
    "offset_x": 9.99,
    "offset_y": 9.99
  },
  "text_outline": {
    "enabled": true,
    "color": "value",
    "opacity": 9.99,
    "width": 9.99
  },
  "background_enabled": true,
  "background_color": "value",
  "background_opacity": 9.99,
  "background_blur": 9.99,
  "background_border_radius": 9.99,
  "word_highlights_enabled": true,
  "word_highlights_color": "value",
  "word_highlights_background_color": "value",
  "word_highlights_opacity": 9.99,
  "word_highlights_border_radius": 9.99,
  "word_highlights_blur": 9.99,
  "section_animation": {
    "enter_type": "none",
    "exit_type": "none"
  },
  "word_animation": {
    "enter_type": "none",
    "exit_type": "none"
  },
  "character_animation": {
    "enter_type": "none",
    "exit_type": "none"
  },
  "cursor_enabled": true,
  "width_pct": 9.99,
  "horizontal_placement": {
    "align": "left",
    "translate_pct": 9.99
  },
  "vertical_placement": {
    "align": "top",
    "translate_pct": 9.99
  },
  "auto_break_enabled": true,
  "max_lines_per_section": 123,
  "max_words_per_line": 123
}`
	assertRoundTrip(t, wire, new(CaptionStyleModel))
}
