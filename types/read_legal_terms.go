package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ReadLegalTerms struct {
	Terms           *string `json:"terms,omitempty" xml:"terms,omitempty"`
	StartDate       *string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	EndDate         *string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ReadLegalTerms) GetTerms() (value *string) {
	if r == nil {
		return
	}
	return r.Terms
}

func (r *ReadLegalTerms) GetStartDate() (value *string) {
	if r == nil {
		return
	}
	return r.StartDate
}

func (r *ReadLegalTerms) GetEndDate() (value *string) {
	if r == nil {
		return
	}
	return r.EndDate
}

func (r *ReadLegalTerms) SetTerms(terms *string) {
	r.Terms = terms
}

func (r *ReadLegalTerms) SetStartDate(startDate *string) {
	r.StartDate = startDate
}

func (r *ReadLegalTerms) SetEndDate(endDate *string) {
	r.EndDate = endDate
}

func (r *ReadLegalTerms) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReadLegalTerms) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReadLegalTerms to string"
	}
	return string(jsonData)
}

func (r *ReadLegalTerms) UnmarshalJSON(data []byte) error {
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

func (r *ReadLegalTerms) MarshalJSON() ([]byte, error) {
	type embed ReadLegalTerms
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
