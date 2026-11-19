package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ListResponseMeta struct {
	Total           *int `json:"total,omitempty" xml:"total,omitempty"`
	Page            *int `json:"page,omitempty" xml:"page,omitempty"`
	PageSize        *int `json:"page_size,omitempty" xml:"page_size,omitempty"`
	extraProperties map[string]interface{}
}

func (l *ListResponseMeta) GetTotal() (value *int) {
	if l == nil {
		return
	}
	return l.Total
}

func (l *ListResponseMeta) GetPage() (value *int) {
	if l == nil {
		return
	}
	return l.Page
}

func (l *ListResponseMeta) GetPageSize() (value *int) {
	if l == nil {
		return
	}
	return l.PageSize
}

func (l *ListResponseMeta) SetTotal(total *int) {
	l.Total = total
}

func (l *ListResponseMeta) SetPage(page *int) {
	l.Page = page
}

func (l *ListResponseMeta) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
}

func (l *ListResponseMeta) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListResponseMeta) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListResponseMeta to string"
	}
	return string(jsonData)
}

func (l *ListResponseMeta) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, l); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListResponseMeta) MarshalJSON() ([]byte, error) {
	type embed ListResponseMeta
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
