package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetPhoneNumbersPageResponseModel struct {
	// The phone numbers on this page
	PhoneNumbers []*GetPhoneNumbersPageResponseModelPhoneNumbersItem `json:"phone_numbers" xml:"phone_numbers" required:"true"`
	// Pass this value as `cursor` to fetch the next page. Null when there are no more results.
	NextCursor *string `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	// Whether there are more results available
	HasMore         *bool `json:"has_more,omitempty" xml:"has_more,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GetPhoneNumbersPageResponseModel) GetPhoneNumbers() (value []*GetPhoneNumbersPageResponseModelPhoneNumbersItem) {
	if g == nil {
		return
	}
	return g.PhoneNumbers
}

func (g *GetPhoneNumbersPageResponseModel) GetNextCursor() (value *string) {
	if g == nil {
		return
	}
	return g.NextCursor
}

func (g *GetPhoneNumbersPageResponseModel) GetHasMore() (value *bool) {
	if g == nil {
		return
	}
	return g.HasMore
}

func (g *GetPhoneNumbersPageResponseModel) SetPhoneNumbers(phoneNumbers []*GetPhoneNumbersPageResponseModelPhoneNumbersItem) {
	g.PhoneNumbers = phoneNumbers
}

func (g *GetPhoneNumbersPageResponseModel) SetNextCursor(nextCursor *string) {
	g.NextCursor = nextCursor
}

func (g *GetPhoneNumbersPageResponseModel) SetHasMore(hasMore *bool) {
	g.HasMore = hasMore
}

func (g *GetPhoneNumbersPageResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetPhoneNumbersPageResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetPhoneNumbersPageResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetPhoneNumbersPageResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetPhoneNumbersPageResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetPhoneNumbersPageResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
