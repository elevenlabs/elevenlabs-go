package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CaptionStyleModel struct {
	Template                      *CaptionStyleTemplateModel            `json:"template,omitempty" xml:"template,omitempty"`
	TextFont                      *string                               `json:"text_font,omitempty" xml:"text_font,omitempty"`
	TextScale                     *float64                              `json:"text_scale,omitempty" xml:"text_scale,omitempty"`
	TextColor                     *string                               `json:"text_color,omitempty" xml:"text_color,omitempty"`
	TextAlign                     *CaptionStyleModelTextAlign           `json:"text_align,omitempty" xml:"text_align,omitempty"`
	TextStyle                     *CaptionStyleModelTextStyle           `json:"text_style,omitempty" xml:"text_style,omitempty"`
	TextWeight                    *CaptionStyleModelTextWeight          `json:"text_weight,omitempty" xml:"text_weight,omitempty"`
	TextTransform                 *CaptionStyleModelTextTransform       `json:"text_transform,omitempty" xml:"text_transform,omitempty"`
	TextBlendMode                 *CaptionStyleModelTextBlendMode       `json:"text_blend_mode,omitempty" xml:"text_blend_mode,omitempty"`
	TextShadow                    *StudioTextStyleShadowModel           `json:"text_shadow,omitempty" xml:"text_shadow,omitempty"`
	TextOutline                   *StudioTextStyleOutlineModel          `json:"text_outline,omitempty" xml:"text_outline,omitempty"`
	BackgroundEnabled             *bool                                 `json:"background_enabled,omitempty" xml:"background_enabled,omitempty"`
	BackgroundColor               *string                               `json:"background_color,omitempty" xml:"background_color,omitempty"`
	BackgroundOpacity             *float64                              `json:"background_opacity,omitempty" xml:"background_opacity,omitempty"`
	BackgroundBlur                *float64                              `json:"background_blur,omitempty" xml:"background_blur,omitempty"`
	BackgroundBorderRadius        *float64                              `json:"background_border_radius,omitempty" xml:"background_border_radius,omitempty"`
	WordHighlightsEnabled         *bool                                 `json:"word_highlights_enabled,omitempty" xml:"word_highlights_enabled,omitempty"`
	WordHighlightsColor           *string                               `json:"word_highlights_color,omitempty" xml:"word_highlights_color,omitempty"`
	WordHighlightsBackgroundColor *string                               `json:"word_highlights_background_color,omitempty" xml:"word_highlights_background_color,omitempty"`
	WordHighlightsOpacity         *float64                              `json:"word_highlights_opacity,omitempty" xml:"word_highlights_opacity,omitempty"`
	WordHighlightsBorderRadius    *float64                              `json:"word_highlights_border_radius,omitempty" xml:"word_highlights_border_radius,omitempty"`
	WordHighlightsBlur            *float64                              `json:"word_highlights_blur,omitempty" xml:"word_highlights_blur,omitempty"`
	SectionAnimation              *CaptionStyleSectionAnimationModel    `json:"section_animation,omitempty" xml:"section_animation,omitempty"`
	WordAnimation                 *CaptionStyleWordAnimationModel       `json:"word_animation,omitempty" xml:"word_animation,omitempty"`
	CharacterAnimation            *CaptionStyleCharacterAnimationModel  `json:"character_animation,omitempty" xml:"character_animation,omitempty"`
	CursorEnabled                 *bool                                 `json:"cursor_enabled,omitempty" xml:"cursor_enabled,omitempty"`
	WidthPct                      *float64                              `json:"width_pct,omitempty" xml:"width_pct,omitempty"`
	HorizontalPlacement           *CaptionStyleHorizontalPlacementModel `json:"horizontal_placement,omitempty" xml:"horizontal_placement,omitempty"`
	VerticalPlacement             *CaptionStyleVerticalPlacementModel   `json:"vertical_placement,omitempty" xml:"vertical_placement,omitempty"`
	AutoBreakEnabled              *bool                                 `json:"auto_break_enabled,omitempty" xml:"auto_break_enabled,omitempty"`
	MaxLinesPerSection            *int                                  `json:"max_lines_per_section,omitempty" xml:"max_lines_per_section,omitempty"`
	MaxWordsPerLine               *int                                  `json:"max_words_per_line,omitempty" xml:"max_words_per_line,omitempty"`
	extraProperties               map[string]interface{}
}

func (c *CaptionStyleModel) GetTemplate() (value *CaptionStyleTemplateModel) {
	if c == nil {
		return
	}
	return c.Template
}

func (c *CaptionStyleModel) GetTextFont() (value *string) {
	if c == nil {
		return
	}
	return c.TextFont
}

func (c *CaptionStyleModel) GetTextScale() (value *float64) {
	if c == nil {
		return
	}
	return c.TextScale
}

func (c *CaptionStyleModel) GetTextColor() (value *string) {
	if c == nil {
		return
	}
	return c.TextColor
}

func (c *CaptionStyleModel) GetTextAlign() (value *CaptionStyleModelTextAlign) {
	if c == nil {
		return
	}
	return c.TextAlign
}

func (c *CaptionStyleModel) GetTextStyle() (value *CaptionStyleModelTextStyle) {
	if c == nil {
		return
	}
	return c.TextStyle
}

func (c *CaptionStyleModel) GetTextWeight() (value *CaptionStyleModelTextWeight) {
	if c == nil {
		return
	}
	return c.TextWeight
}

func (c *CaptionStyleModel) GetTextTransform() (value *CaptionStyleModelTextTransform) {
	if c == nil {
		return
	}
	return c.TextTransform
}

func (c *CaptionStyleModel) GetTextBlendMode() (value *CaptionStyleModelTextBlendMode) {
	if c == nil {
		return
	}
	return c.TextBlendMode
}

func (c *CaptionStyleModel) GetTextShadow() (value *StudioTextStyleShadowModel) {
	if c == nil {
		return
	}
	return c.TextShadow
}

func (c *CaptionStyleModel) GetTextOutline() (value *StudioTextStyleOutlineModel) {
	if c == nil {
		return
	}
	return c.TextOutline
}

func (c *CaptionStyleModel) GetBackgroundEnabled() (value *bool) {
	if c == nil {
		return
	}
	return c.BackgroundEnabled
}

func (c *CaptionStyleModel) GetBackgroundColor() (value *string) {
	if c == nil {
		return
	}
	return c.BackgroundColor
}

func (c *CaptionStyleModel) GetBackgroundOpacity() (value *float64) {
	if c == nil {
		return
	}
	return c.BackgroundOpacity
}

func (c *CaptionStyleModel) GetBackgroundBlur() (value *float64) {
	if c == nil {
		return
	}
	return c.BackgroundBlur
}

func (c *CaptionStyleModel) GetBackgroundBorderRadius() (value *float64) {
	if c == nil {
		return
	}
	return c.BackgroundBorderRadius
}

func (c *CaptionStyleModel) GetWordHighlightsEnabled() (value *bool) {
	if c == nil {
		return
	}
	return c.WordHighlightsEnabled
}

func (c *CaptionStyleModel) GetWordHighlightsColor() (value *string) {
	if c == nil {
		return
	}
	return c.WordHighlightsColor
}

func (c *CaptionStyleModel) GetWordHighlightsBackgroundColor() (value *string) {
	if c == nil {
		return
	}
	return c.WordHighlightsBackgroundColor
}

func (c *CaptionStyleModel) GetWordHighlightsOpacity() (value *float64) {
	if c == nil {
		return
	}
	return c.WordHighlightsOpacity
}

func (c *CaptionStyleModel) GetWordHighlightsBorderRadius() (value *float64) {
	if c == nil {
		return
	}
	return c.WordHighlightsBorderRadius
}

func (c *CaptionStyleModel) GetWordHighlightsBlur() (value *float64) {
	if c == nil {
		return
	}
	return c.WordHighlightsBlur
}

func (c *CaptionStyleModel) GetSectionAnimation() (value *CaptionStyleSectionAnimationModel) {
	if c == nil {
		return
	}
	return c.SectionAnimation
}

func (c *CaptionStyleModel) GetWordAnimation() (value *CaptionStyleWordAnimationModel) {
	if c == nil {
		return
	}
	return c.WordAnimation
}

func (c *CaptionStyleModel) GetCharacterAnimation() (value *CaptionStyleCharacterAnimationModel) {
	if c == nil {
		return
	}
	return c.CharacterAnimation
}

func (c *CaptionStyleModel) GetCursorEnabled() (value *bool) {
	if c == nil {
		return
	}
	return c.CursorEnabled
}

func (c *CaptionStyleModel) GetWidthPct() (value *float64) {
	if c == nil {
		return
	}
	return c.WidthPct
}

func (c *CaptionStyleModel) GetHorizontalPlacement() (value *CaptionStyleHorizontalPlacementModel) {
	if c == nil {
		return
	}
	return c.HorizontalPlacement
}

func (c *CaptionStyleModel) GetVerticalPlacement() (value *CaptionStyleVerticalPlacementModel) {
	if c == nil {
		return
	}
	return c.VerticalPlacement
}

func (c *CaptionStyleModel) GetAutoBreakEnabled() (value *bool) {
	if c == nil {
		return
	}
	return c.AutoBreakEnabled
}

func (c *CaptionStyleModel) GetMaxLinesPerSection() (value *int) {
	if c == nil {
		return
	}
	return c.MaxLinesPerSection
}

func (c *CaptionStyleModel) GetMaxWordsPerLine() (value *int) {
	if c == nil {
		return
	}
	return c.MaxWordsPerLine
}

func (c *CaptionStyleModel) SetTemplate(template *CaptionStyleTemplateModel) {
	c.Template = template
}

func (c *CaptionStyleModel) SetTextFont(textFont *string) {
	c.TextFont = textFont
}

func (c *CaptionStyleModel) SetTextScale(textScale *float64) {
	c.TextScale = textScale
}

func (c *CaptionStyleModel) SetTextColor(textColor *string) {
	c.TextColor = textColor
}

func (c *CaptionStyleModel) SetTextAlign(textAlign *CaptionStyleModelTextAlign) {
	c.TextAlign = textAlign
}

func (c *CaptionStyleModel) SetTextStyle(textStyle *CaptionStyleModelTextStyle) {
	c.TextStyle = textStyle
}

func (c *CaptionStyleModel) SetTextWeight(textWeight *CaptionStyleModelTextWeight) {
	c.TextWeight = textWeight
}

func (c *CaptionStyleModel) SetTextTransform(textTransform *CaptionStyleModelTextTransform) {
	c.TextTransform = textTransform
}

func (c *CaptionStyleModel) SetTextBlendMode(textBlendMode *CaptionStyleModelTextBlendMode) {
	c.TextBlendMode = textBlendMode
}

func (c *CaptionStyleModel) SetTextShadow(textShadow *StudioTextStyleShadowModel) {
	c.TextShadow = textShadow
}

func (c *CaptionStyleModel) SetTextOutline(textOutline *StudioTextStyleOutlineModel) {
	c.TextOutline = textOutline
}

func (c *CaptionStyleModel) SetBackgroundEnabled(backgroundEnabled *bool) {
	c.BackgroundEnabled = backgroundEnabled
}

func (c *CaptionStyleModel) SetBackgroundColor(backgroundColor *string) {
	c.BackgroundColor = backgroundColor
}

func (c *CaptionStyleModel) SetBackgroundOpacity(backgroundOpacity *float64) {
	c.BackgroundOpacity = backgroundOpacity
}

func (c *CaptionStyleModel) SetBackgroundBlur(backgroundBlur *float64) {
	c.BackgroundBlur = backgroundBlur
}

func (c *CaptionStyleModel) SetBackgroundBorderRadius(backgroundBorderRadius *float64) {
	c.BackgroundBorderRadius = backgroundBorderRadius
}

func (c *CaptionStyleModel) SetWordHighlightsEnabled(wordHighlightsEnabled *bool) {
	c.WordHighlightsEnabled = wordHighlightsEnabled
}

func (c *CaptionStyleModel) SetWordHighlightsColor(wordHighlightsColor *string) {
	c.WordHighlightsColor = wordHighlightsColor
}

func (c *CaptionStyleModel) SetWordHighlightsBackgroundColor(wordHighlightsBackgroundColor *string) {
	c.WordHighlightsBackgroundColor = wordHighlightsBackgroundColor
}

func (c *CaptionStyleModel) SetWordHighlightsOpacity(wordHighlightsOpacity *float64) {
	c.WordHighlightsOpacity = wordHighlightsOpacity
}

func (c *CaptionStyleModel) SetWordHighlightsBorderRadius(wordHighlightsBorderRadius *float64) {
	c.WordHighlightsBorderRadius = wordHighlightsBorderRadius
}

func (c *CaptionStyleModel) SetWordHighlightsBlur(wordHighlightsBlur *float64) {
	c.WordHighlightsBlur = wordHighlightsBlur
}

func (c *CaptionStyleModel) SetSectionAnimation(sectionAnimation *CaptionStyleSectionAnimationModel) {
	c.SectionAnimation = sectionAnimation
}

func (c *CaptionStyleModel) SetWordAnimation(wordAnimation *CaptionStyleWordAnimationModel) {
	c.WordAnimation = wordAnimation
}

func (c *CaptionStyleModel) SetCharacterAnimation(characterAnimation *CaptionStyleCharacterAnimationModel) {
	c.CharacterAnimation = characterAnimation
}

func (c *CaptionStyleModel) SetCursorEnabled(cursorEnabled *bool) {
	c.CursorEnabled = cursorEnabled
}

func (c *CaptionStyleModel) SetWidthPct(widthPct *float64) {
	c.WidthPct = widthPct
}

func (c *CaptionStyleModel) SetHorizontalPlacement(horizontalPlacement *CaptionStyleHorizontalPlacementModel) {
	c.HorizontalPlacement = horizontalPlacement
}

func (c *CaptionStyleModel) SetVerticalPlacement(verticalPlacement *CaptionStyleVerticalPlacementModel) {
	c.VerticalPlacement = verticalPlacement
}

func (c *CaptionStyleModel) SetAutoBreakEnabled(autoBreakEnabled *bool) {
	c.AutoBreakEnabled = autoBreakEnabled
}

func (c *CaptionStyleModel) SetMaxLinesPerSection(maxLinesPerSection *int) {
	c.MaxLinesPerSection = maxLinesPerSection
}

func (c *CaptionStyleModel) SetMaxWordsPerLine(maxWordsPerLine *int) {
	c.MaxWordsPerLine = maxWordsPerLine
}

func (c *CaptionStyleModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CaptionStyleModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CaptionStyleModel to string"
	}
	return string(jsonData)
}

func (c *CaptionStyleModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CaptionStyleModel) MarshalJSON() ([]byte, error) {
	type embed CaptionStyleModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
