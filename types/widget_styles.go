package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WidgetStyles struct {
	// The base background color.
	Base *string `json:"base,omitempty" xml:"base,omitempty"`
	// The color of the base background when hovered.
	BaseHover *string `json:"base_hover,omitempty" xml:"base_hover,omitempty"`
	// The color of the base background when active (clicked).
	BaseActive *string `json:"base_active,omitempty" xml:"base_active,omitempty"`
	// The color of the border against the base background.
	BaseBorder *string `json:"base_border,omitempty" xml:"base_border,omitempty"`
	// The color of subtle text against the base background.
	BaseSubtle *string `json:"base_subtle,omitempty" xml:"base_subtle,omitempty"`
	// The color of primary text against the base background.
	BasePrimary *string `json:"base_primary,omitempty" xml:"base_primary,omitempty"`
	// The color of error text against the base background.
	BaseError *string `json:"base_error,omitempty" xml:"base_error,omitempty"`
	// The accent background color.
	Accent *string `json:"accent,omitempty" xml:"accent,omitempty"`
	// The color of the accent background when hovered.
	AccentHover *string `json:"accent_hover,omitempty" xml:"accent_hover,omitempty"`
	// The color of the accent background when active (clicked).
	AccentActive *string `json:"accent_active,omitempty" xml:"accent_active,omitempty"`
	// The color of the border against the accent background.
	AccentBorder *string `json:"accent_border,omitempty" xml:"accent_border,omitempty"`
	// The color of subtle text against the accent background.
	AccentSubtle *string `json:"accent_subtle,omitempty" xml:"accent_subtle,omitempty"`
	// The color of primary text against the accent background.
	AccentPrimary *string `json:"accent_primary,omitempty" xml:"accent_primary,omitempty"`
	// The padding around the edges of the viewport.
	OverlayPadding *float64 `json:"overlay_padding,omitempty" xml:"overlay_padding,omitempty"`
	// The radius of the buttons.
	ButtonRadius *float64 `json:"button_radius,omitempty" xml:"button_radius,omitempty"`
	// The radius of the input fields.
	InputRadius *float64 `json:"input_radius,omitempty" xml:"input_radius,omitempty"`
	// The radius of the chat bubbles.
	BubbleRadius *float64 `json:"bubble_radius,omitempty" xml:"bubble_radius,omitempty"`
	// The default radius of sheets.
	SheetRadius *float64 `json:"sheet_radius,omitempty" xml:"sheet_radius,omitempty"`
	// The radius of the sheet in compact mode.
	CompactSheetRadius *float64 `json:"compact_sheet_radius,omitempty" xml:"compact_sheet_radius,omitempty"`
	// The radius of the dropdown sheet.
	DropdownSheetRadius *float64 `json:"dropdown_sheet_radius,omitempty" xml:"dropdown_sheet_radius,omitempty"`
	extraProperties     map[string]interface{}
}

func (w *WidgetStyles) GetBase() (value *string) {
	if w == nil {
		return
	}
	return w.Base
}

func (w *WidgetStyles) GetBaseHover() (value *string) {
	if w == nil {
		return
	}
	return w.BaseHover
}

func (w *WidgetStyles) GetBaseActive() (value *string) {
	if w == nil {
		return
	}
	return w.BaseActive
}

func (w *WidgetStyles) GetBaseBorder() (value *string) {
	if w == nil {
		return
	}
	return w.BaseBorder
}

func (w *WidgetStyles) GetBaseSubtle() (value *string) {
	if w == nil {
		return
	}
	return w.BaseSubtle
}

func (w *WidgetStyles) GetBasePrimary() (value *string) {
	if w == nil {
		return
	}
	return w.BasePrimary
}

func (w *WidgetStyles) GetBaseError() (value *string) {
	if w == nil {
		return
	}
	return w.BaseError
}

func (w *WidgetStyles) GetAccent() (value *string) {
	if w == nil {
		return
	}
	return w.Accent
}

func (w *WidgetStyles) GetAccentHover() (value *string) {
	if w == nil {
		return
	}
	return w.AccentHover
}

func (w *WidgetStyles) GetAccentActive() (value *string) {
	if w == nil {
		return
	}
	return w.AccentActive
}

func (w *WidgetStyles) GetAccentBorder() (value *string) {
	if w == nil {
		return
	}
	return w.AccentBorder
}

func (w *WidgetStyles) GetAccentSubtle() (value *string) {
	if w == nil {
		return
	}
	return w.AccentSubtle
}

func (w *WidgetStyles) GetAccentPrimary() (value *string) {
	if w == nil {
		return
	}
	return w.AccentPrimary
}

func (w *WidgetStyles) GetOverlayPadding() (value *float64) {
	if w == nil {
		return
	}
	return w.OverlayPadding
}

func (w *WidgetStyles) GetButtonRadius() (value *float64) {
	if w == nil {
		return
	}
	return w.ButtonRadius
}

func (w *WidgetStyles) GetInputRadius() (value *float64) {
	if w == nil {
		return
	}
	return w.InputRadius
}

func (w *WidgetStyles) GetBubbleRadius() (value *float64) {
	if w == nil {
		return
	}
	return w.BubbleRadius
}

func (w *WidgetStyles) GetSheetRadius() (value *float64) {
	if w == nil {
		return
	}
	return w.SheetRadius
}

func (w *WidgetStyles) GetCompactSheetRadius() (value *float64) {
	if w == nil {
		return
	}
	return w.CompactSheetRadius
}

func (w *WidgetStyles) GetDropdownSheetRadius() (value *float64) {
	if w == nil {
		return
	}
	return w.DropdownSheetRadius
}

func (w *WidgetStyles) SetBase(base *string) {
	w.Base = base
}

func (w *WidgetStyles) SetBaseHover(baseHover *string) {
	w.BaseHover = baseHover
}

func (w *WidgetStyles) SetBaseActive(baseActive *string) {
	w.BaseActive = baseActive
}

func (w *WidgetStyles) SetBaseBorder(baseBorder *string) {
	w.BaseBorder = baseBorder
}

func (w *WidgetStyles) SetBaseSubtle(baseSubtle *string) {
	w.BaseSubtle = baseSubtle
}

func (w *WidgetStyles) SetBasePrimary(basePrimary *string) {
	w.BasePrimary = basePrimary
}

func (w *WidgetStyles) SetBaseError(baseError *string) {
	w.BaseError = baseError
}

func (w *WidgetStyles) SetAccent(accent *string) {
	w.Accent = accent
}

func (w *WidgetStyles) SetAccentHover(accentHover *string) {
	w.AccentHover = accentHover
}

func (w *WidgetStyles) SetAccentActive(accentActive *string) {
	w.AccentActive = accentActive
}

func (w *WidgetStyles) SetAccentBorder(accentBorder *string) {
	w.AccentBorder = accentBorder
}

func (w *WidgetStyles) SetAccentSubtle(accentSubtle *string) {
	w.AccentSubtle = accentSubtle
}

func (w *WidgetStyles) SetAccentPrimary(accentPrimary *string) {
	w.AccentPrimary = accentPrimary
}

func (w *WidgetStyles) SetOverlayPadding(overlayPadding *float64) {
	w.OverlayPadding = overlayPadding
}

func (w *WidgetStyles) SetButtonRadius(buttonRadius *float64) {
	w.ButtonRadius = buttonRadius
}

func (w *WidgetStyles) SetInputRadius(inputRadius *float64) {
	w.InputRadius = inputRadius
}

func (w *WidgetStyles) SetBubbleRadius(bubbleRadius *float64) {
	w.BubbleRadius = bubbleRadius
}

func (w *WidgetStyles) SetSheetRadius(sheetRadius *float64) {
	w.SheetRadius = sheetRadius
}

func (w *WidgetStyles) SetCompactSheetRadius(compactSheetRadius *float64) {
	w.CompactSheetRadius = compactSheetRadius
}

func (w *WidgetStyles) SetDropdownSheetRadius(dropdownSheetRadius *float64) {
	w.DropdownSheetRadius = dropdownSheetRadius
}

func (w *WidgetStyles) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WidgetStyles) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WidgetStyles to string"
	}
	return string(jsonData)
}

func (w *WidgetStyles) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WidgetStyles) MarshalJSON() ([]byte, error) {
	type embed WidgetStyles
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
