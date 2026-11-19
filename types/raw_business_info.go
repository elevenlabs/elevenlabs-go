package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Business profile data extracted from a single page.
type RawBusinessInfo struct {
	CompanyName       LenientOptionalStr `json:"company_name,omitempty" xml:"company_name,omitempty"`
	Description       LenientOptionalStr `json:"description,omitempty" xml:"description,omitempty"`
	BusinessPhone     LenientOptionalStr `json:"business_phone,omitempty" xml:"business_phone,omitempty"`
	Locations         []*RawLocationInfo `json:"locations,omitempty" xml:"locations,omitempty"`
	BusinessCountry   LenientOptionalStr `json:"business_country,omitempty" xml:"business_country,omitempty"`
	BusinessState     LenientOptionalStr `json:"business_state,omitempty" xml:"business_state,omitempty"`
	SuggestedCurrency LenientOptionalStr `json:"suggested_currency,omitempty" xml:"suggested_currency,omitempty"`
	SuggestedLanguage LenientOptionalStr `json:"suggested_language,omitempty" xml:"suggested_language,omitempty"`
	extraProperties   map[string]interface{}
}

func (r *RawBusinessInfo) GetCompanyName() (value LenientOptionalStr) {
	if r == nil {
		return
	}
	return r.CompanyName
}

func (r *RawBusinessInfo) GetDescription() (value LenientOptionalStr) {
	if r == nil {
		return
	}
	return r.Description
}

func (r *RawBusinessInfo) GetBusinessPhone() (value LenientOptionalStr) {
	if r == nil {
		return
	}
	return r.BusinessPhone
}

func (r *RawBusinessInfo) GetLocations() (value []*RawLocationInfo) {
	if r == nil {
		return
	}
	return r.Locations
}

func (r *RawBusinessInfo) GetBusinessCountry() (value LenientOptionalStr) {
	if r == nil {
		return
	}
	return r.BusinessCountry
}

func (r *RawBusinessInfo) GetBusinessState() (value LenientOptionalStr) {
	if r == nil {
		return
	}
	return r.BusinessState
}

func (r *RawBusinessInfo) GetSuggestedCurrency() (value LenientOptionalStr) {
	if r == nil {
		return
	}
	return r.SuggestedCurrency
}

func (r *RawBusinessInfo) GetSuggestedLanguage() (value LenientOptionalStr) {
	if r == nil {
		return
	}
	return r.SuggestedLanguage
}

func (r *RawBusinessInfo) SetCompanyName(companyName LenientOptionalStr) {
	r.CompanyName = companyName
}

func (r *RawBusinessInfo) SetDescription(description LenientOptionalStr) {
	r.Description = description
}

func (r *RawBusinessInfo) SetBusinessPhone(businessPhone LenientOptionalStr) {
	r.BusinessPhone = businessPhone
}

func (r *RawBusinessInfo) SetLocations(locations []*RawLocationInfo) {
	r.Locations = locations
}

func (r *RawBusinessInfo) SetBusinessCountry(businessCountry LenientOptionalStr) {
	r.BusinessCountry = businessCountry
}

func (r *RawBusinessInfo) SetBusinessState(businessState LenientOptionalStr) {
	r.BusinessState = businessState
}

func (r *RawBusinessInfo) SetSuggestedCurrency(suggestedCurrency LenientOptionalStr) {
	r.SuggestedCurrency = suggestedCurrency
}

func (r *RawBusinessInfo) SetSuggestedLanguage(suggestedLanguage LenientOptionalStr) {
	r.SuggestedLanguage = suggestedLanguage
}

func (r *RawBusinessInfo) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RawBusinessInfo) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RawBusinessInfo to string"
	}
	return string(jsonData)
}

func (r *RawBusinessInfo) UnmarshalJSON(data []byte) error {
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

func (r *RawBusinessInfo) MarshalJSON() ([]byte, error) {
	type embed RawBusinessInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
