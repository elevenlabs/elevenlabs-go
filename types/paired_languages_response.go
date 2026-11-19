package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type PairedLanguagesResponse struct {
	// Indicates this response contains source-to-destination language pairs.
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The list of available source-to-destination language mappings.
	LanguagePairs   []*LanguagePairInfo `json:"language_pairs" xml:"language_pairs" required:"true"`
	extraProperties map[string]interface{}
}

func (p *PairedLanguagesResponse) GetKind() (value *string) {
	if p == nil {
		return
	}
	return p.Kind
}

func (p *PairedLanguagesResponse) GetLanguagePairs() (value []*LanguagePairInfo) {
	if p == nil {
		return
	}
	return p.LanguagePairs
}

func (p *PairedLanguagesResponse) SetKind(kind *string) {
	p.Kind = kind
}

func (p *PairedLanguagesResponse) SetLanguagePairs(languagePairs []*LanguagePairInfo) {
	p.LanguagePairs = languagePairs
}

func (p *PairedLanguagesResponse) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PairedLanguagesResponse) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PairedLanguagesResponse to string"
	}
	return string(jsonData)
}

func (p *PairedLanguagesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler PairedLanguagesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PairedLanguagesResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PairedLanguagesResponse) MarshalJSON() ([]byte, error) {
	type embed PairedLanguagesResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
