package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// How the website widget looks. A null label follows the receptionist's language.
type ReceptionWidgetAppearance struct {
	// First color of the widget's orb gradient, as a hex color.
	OrbColor1 string `json:"orb_color_1" xml:"orb_color_1" required:"true" pattern:"^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"`
	// Second color of the widget's orb gradient, as a hex color.
	OrbColor2 string `json:"orb_color_2" xml:"orb_color_2" required:"true" pattern:"^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"`
	// Widget size, from the smallest ('tiny') to the largest ('full').
	Variant ReceptionWidgetAppearanceVariant `json:"variant" xml:"variant" required:"true"`
	// Where on the page the widget sits.
	Placement ReceptionWidgetAppearancePlacement `json:"placement" xml:"placement" required:"true"`
	// Whether visitors can close the widget.
	Dismissible bool `json:"dismissible" xml:"dismissible" required:"true"`
	// Text shown on the widget; null uses the default for the receptionist's language.
	MainLabel *string `json:"main_label,omitempty" xml:"main_label,omitempty" maxLength:"100"`
	// Text on the button that starts a call; null uses the default for the receptionist's language.
	StartCall       *string `json:"start_call,omitempty" xml:"start_call,omitempty" maxLength:"100"`
	extraProperties map[string]interface{}
}

func (r *ReceptionWidgetAppearance) GetOrbColor1() (value string) {
	if r == nil {
		return
	}
	return r.OrbColor1
}

func (r *ReceptionWidgetAppearance) GetOrbColor2() (value string) {
	if r == nil {
		return
	}
	return r.OrbColor2
}

func (r *ReceptionWidgetAppearance) GetVariant() (value ReceptionWidgetAppearanceVariant) {
	if r == nil {
		return
	}
	return r.Variant
}

func (r *ReceptionWidgetAppearance) GetPlacement() (value ReceptionWidgetAppearancePlacement) {
	if r == nil {
		return
	}
	return r.Placement
}

func (r *ReceptionWidgetAppearance) GetDismissible() (value bool) {
	if r == nil {
		return
	}
	return r.Dismissible
}

func (r *ReceptionWidgetAppearance) GetMainLabel() (value *string) {
	if r == nil {
		return
	}
	return r.MainLabel
}

func (r *ReceptionWidgetAppearance) GetStartCall() (value *string) {
	if r == nil {
		return
	}
	return r.StartCall
}

func (r *ReceptionWidgetAppearance) SetOrbColor1(orbColor1 string) {
	r.OrbColor1 = orbColor1
}

func (r *ReceptionWidgetAppearance) SetOrbColor2(orbColor2 string) {
	r.OrbColor2 = orbColor2
}

func (r *ReceptionWidgetAppearance) SetVariant(variant ReceptionWidgetAppearanceVariant) {
	r.Variant = variant
}

func (r *ReceptionWidgetAppearance) SetPlacement(placement ReceptionWidgetAppearancePlacement) {
	r.Placement = placement
}

func (r *ReceptionWidgetAppearance) SetDismissible(dismissible bool) {
	r.Dismissible = dismissible
}

func (r *ReceptionWidgetAppearance) SetMainLabel(mainLabel *string) {
	r.MainLabel = mainLabel
}

func (r *ReceptionWidgetAppearance) SetStartCall(startCall *string) {
	r.StartCall = startCall
}

func (r *ReceptionWidgetAppearance) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionWidgetAppearance) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionWidgetAppearance to string"
	}
	return string(jsonData)
}

func (r *ReceptionWidgetAppearance) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ReceptionWidgetAppearance) MarshalJSON() ([]byte, error) {
	type embed ReceptionWidgetAppearance
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
