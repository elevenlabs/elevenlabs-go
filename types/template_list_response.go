package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// One page of the templates the caller can run, most recently updated first.
type TemplateListResponse struct {
	// The templates on this page, most recently updated first. A page can hold fewer than `page_size` templates, or none at all, because templates you cannot run are filtered out after the page is read — keep paging while `has_more` is true.
	Templates []*TemplateSummary `json:"templates" xml:"templates" required:"true"`
	// Pass as `cursor` to fetch the next page. `null` when there is no further page.
	NextCursor *string `json:"next_cursor" xml:"next_cursor" required:"true"`
	// Whether more templates exist beyond this page.
	HasMore         *bool `json:"has_more,omitempty" xml:"has_more,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TemplateListResponse) GetTemplates() (value []*TemplateSummary) {
	if t == nil {
		return
	}
	return t.Templates
}

func (t *TemplateListResponse) GetNextCursor() (value *string) {
	if t == nil {
		return
	}
	return t.NextCursor
}

func (t *TemplateListResponse) GetHasMore() (value *bool) {
	if t == nil {
		return
	}
	return t.HasMore
}

func (t *TemplateListResponse) SetTemplates(templates []*TemplateSummary) {
	t.Templates = templates
}

func (t *TemplateListResponse) SetNextCursor(nextCursor *string) {
	t.NextCursor = nextCursor
}

func (t *TemplateListResponse) SetHasMore(hasMore *bool) {
	t.HasMore = hasMore
}

func (t *TemplateListResponse) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TemplateListResponse) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TemplateListResponse to string"
	}
	return string(jsonData)
}

func (t *TemplateListResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TemplateListResponse) MarshalJSON() ([]byte, error) {
	type embed TemplateListResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
