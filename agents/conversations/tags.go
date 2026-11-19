package conversations

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	big "math/big"
)

type AssignConversationTagsRequestModel struct {
	ConversationId string `json:"-" url:"-"`
	// Tag IDs to add to the conversation. Re-assigning an existing tag is a no-op.
	TagIds []string `json:"tag_ids" xml:"tag_ids" required:"true" minItems:"1" maxItems:"50"`
}

func (a *AssignConversationTagsRequestModel) GetConversationId() (value string) {
	if a == nil {
		return
	}
	return a.ConversationId
}

func (a *AssignConversationTagsRequestModel) GetTagIds() (value []string) {
	if a == nil {
		return
	}
	return a.TagIds
}

func (a *AssignConversationTagsRequestModel) SetConversationId(conversationId string) {
	a.ConversationId = conversationId
}

func (a *AssignConversationTagsRequestModel) SetTagIds(tagIds []string) {
	a.TagIds = tagIds
}

func (a *AssignConversationTagsRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (a AssignConversationTagsRequestModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AssignConversationTagsRequestModel to string"
	}
	return string(jsonData)
}

func (a *AssignConversationTagsRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AssignConversationTagsRequestModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AssignConversationTagsRequestModel(tmp)
	return nil
}

func (a *AssignConversationTagsRequestModel) MarshalJSON() ([]byte, error) {
	type embed AssignConversationTagsRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}

type CreateConversationTagRequestModel struct {
	// Display title of the tag.
	Title string `json:"title" xml:"title" required:"true" maxLength:"120" minLength:"1"`
	// Optional free-text description.
	Description *string `json:"description,omitempty" xml:"description,omitempty" maxLength:"1000"`
}

func (c *CreateConversationTagRequestModel) GetTitle() (value string) {
	if c == nil {
		return
	}
	return c.Title
}

func (c *CreateConversationTagRequestModel) GetDescription() (value *string) {
	if c == nil {
		return
	}
	return c.Description
}

func (c *CreateConversationTagRequestModel) SetTitle(title string) {
	c.Title = title
}

func (c *CreateConversationTagRequestModel) SetDescription(description *string) {
	c.Description = description
}

func (c *CreateConversationTagRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (c CreateConversationTagRequestModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateConversationTagRequestModel to string"
	}
	return string(jsonData)
}

func (c *CreateConversationTagRequestModel) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}

func (c *CreateConversationTagRequestModel) MarshalJSON() ([]byte, error) {
	type embed CreateConversationTagRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

type PatchConversationTagRequestModel struct {
	TagId string `json:"-" url:"-"`
	// If provided, replaces the tag title. Omit to leave unchanged.
	Title *string `json:"title,omitempty" xml:"title,omitempty" maxLength:"120" minLength:"1"`
	// If provided, replaces the tag description. Omit to leave unchanged.
	Description *string `json:"description,omitempty" xml:"description,omitempty" maxLength:"1000"`
}

func (p *PatchConversationTagRequestModel) GetTagId() (value string) {
	if p == nil {
		return
	}
	return p.TagId
}

func (p *PatchConversationTagRequestModel) GetTitle() (value *string) {
	if p == nil {
		return
	}
	return p.Title
}

func (p *PatchConversationTagRequestModel) GetDescription() (value *string) {
	if p == nil {
		return
	}
	return p.Description
}

func (p *PatchConversationTagRequestModel) SetTagId(tagId string) {
	p.TagId = tagId
}

func (p *PatchConversationTagRequestModel) SetTitle(title *string) {
	p.Title = title
}

func (p *PatchConversationTagRequestModel) SetDescription(description *string) {
	p.Description = description
}

func (p *PatchConversationTagRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (p PatchConversationTagRequestModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PatchConversationTagRequestModel to string"
	}
	return string(jsonData)
}

func (p *PatchConversationTagRequestModel) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, p)
}

func (p *PatchConversationTagRequestModel) MarshalJSON() ([]byte, error) {
	type embed PatchConversationTagRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}

var (
	listTagsRequestFieldPageSize = big.NewInt(1 << 0)
	listTagsRequestFieldCursor   = big.NewInt(1 << 1)
)

type ListTagsRequest struct {
	PageSize *int    `json:"-" url:"page_size,omitempty"`
	Cursor   *string `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListTagsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTagsRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listTagsRequestFieldPageSize)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTagsRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listTagsRequestFieldCursor)
}
