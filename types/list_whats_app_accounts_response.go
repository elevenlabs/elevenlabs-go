package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ListWhatsAppAccountsResponse struct {
	Items           []*GetWhatsAppAccountResponse `json:"items" xml:"items" required:"true"`
	extraProperties map[string]interface{}
}

func (l *ListWhatsAppAccountsResponse) GetItems() (value []*GetWhatsAppAccountResponse) {
	if l == nil {
		return
	}
	return l.Items
}

func (l *ListWhatsAppAccountsResponse) SetItems(items []*GetWhatsAppAccountResponse) {
	l.Items = items
}

func (l *ListWhatsAppAccountsResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListWhatsAppAccountsResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListWhatsAppAccountsResponse to string"
	}
	return string(jsonData)
}

func (l *ListWhatsAppAccountsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListWhatsAppAccountsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListWhatsAppAccountsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListWhatsAppAccountsResponse) MarshalJSON() ([]byte, error) {
	type embed ListWhatsAppAccountsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
