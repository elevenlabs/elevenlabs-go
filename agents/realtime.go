package agents

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type CreateRealtimeConversationGroupRequestModel struct {
	// Name of the conversation group. Names are unique per workspace, ignoring case and extra whitespace, and at most 64 characters.
	Title string `json:"title" xml:"title" required:"true"`
	// Optional description of the conversation group.
	Description *string `json:"description,omitempty" xml:"description,omitempty" maxLength:"500"`
}

func (c *CreateRealtimeConversationGroupRequestModel) GetTitle() (value string) {
	if c == nil {
		return
	}
	return c.Title
}

func (c *CreateRealtimeConversationGroupRequestModel) GetDescription() (value *string) {
	if c == nil {
		return
	}
	return c.Description
}

func (c *CreateRealtimeConversationGroupRequestModel) SetTitle(title string) {
	c.Title = title
}

func (c *CreateRealtimeConversationGroupRequestModel) SetDescription(description *string) {
	c.Description = description
}

func (c *CreateRealtimeConversationGroupRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (c CreateRealtimeConversationGroupRequestModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateRealtimeConversationGroupRequestModel to string"
	}
	return string(jsonData)
}

func (c *CreateRealtimeConversationGroupRequestModel) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}

func (c *CreateRealtimeConversationGroupRequestModel) MarshalJSON() ([]byte, error) {
	type embed CreateRealtimeConversationGroupRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

// Only fields present in the request are applied; `description=None` clears it.
type UpdateRealtimeConversationGroupRequestModel struct {
	// The id of a realtime conversation group.
	ConversationGroupId string `json:"-" url:"-"`
	// Name of the conversation group. Names are unique per workspace, ignoring case and extra whitespace, and at most 64 characters.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// New description of the conversation group.
	Description *string `json:"description,omitempty" xml:"description,omitempty" maxLength:"500"`
}

func (u *UpdateRealtimeConversationGroupRequestModel) GetConversationGroupId() (value string) {
	if u == nil {
		return
	}
	return u.ConversationGroupId
}

func (u *UpdateRealtimeConversationGroupRequestModel) GetTitle() (value *string) {
	if u == nil {
		return
	}
	return u.Title
}

func (u *UpdateRealtimeConversationGroupRequestModel) GetDescription() (value *string) {
	if u == nil {
		return
	}
	return u.Description
}

func (u *UpdateRealtimeConversationGroupRequestModel) SetConversationGroupId(conversationGroupId string) {
	u.ConversationGroupId = conversationGroupId
}

func (u *UpdateRealtimeConversationGroupRequestModel) SetTitle(title *string) {
	u.Title = title
}

func (u *UpdateRealtimeConversationGroupRequestModel) SetDescription(description *string) {
	u.Description = description
}

func (u *UpdateRealtimeConversationGroupRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (u UpdateRealtimeConversationGroupRequestModel) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateRealtimeConversationGroupRequestModel to string"
	}
	return string(jsonData)
}

func (u *UpdateRealtimeConversationGroupRequestModel) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, u)
}

func (u *UpdateRealtimeConversationGroupRequestModel) MarshalJSON() ([]byte, error) {
	type embed UpdateRealtimeConversationGroupRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}

type SetConversationGroupRequestModel struct {
	// The id of a realtime conversation.
	ConversationId string `json:"-" url:"-"`
	// The conversation group to place the conversation in. A conversation belongs to at most one group, so this moves it out of any current group.
	ConversationGroupId string `json:"conversation_group_id" xml:"conversation_group_id" required:"true"`
}

func (s *SetConversationGroupRequestModel) GetConversationId() (value string) {
	if s == nil {
		return
	}
	return s.ConversationId
}

func (s *SetConversationGroupRequestModel) GetConversationGroupId() (value string) {
	if s == nil {
		return
	}
	return s.ConversationGroupId
}

func (s *SetConversationGroupRequestModel) SetConversationId(conversationId string) {
	s.ConversationId = conversationId
}

func (s *SetConversationGroupRequestModel) SetConversationGroupId(conversationGroupId string) {
	s.ConversationGroupId = conversationGroupId
}

func (s *SetConversationGroupRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (s SetConversationGroupRequestModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetConversationGroupRequestModel to string"
	}
	return string(jsonData)
}

func (s *SetConversationGroupRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SetConversationGroupRequestModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SetConversationGroupRequestModel(tmp)
	return nil
}

func (s *SetConversationGroupRequestModel) MarshalJSON() ([]byte, error) {
	type embed SetConversationGroupRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}

type CreateClientSecretRequest struct {
	Session      *types.OpenAiSessionConfigInput `json:"session,omitempty" xml:"session,omitempty"`
	ExpiresAfter *types.ExpiresAfter             `json:"expires_after,omitempty" xml:"expires_after,omitempty"`
}

func (c *CreateClientSecretRequest) GetSession() (value *types.OpenAiSessionConfigInput) {
	if c == nil {
		return
	}
	return c.Session
}

func (c *CreateClientSecretRequest) GetExpiresAfter() (value *types.ExpiresAfter) {
	if c == nil {
		return
	}
	return c.ExpiresAfter
}

func (c *CreateClientSecretRequest) SetSession(session *types.OpenAiSessionConfigInput) {
	c.Session = session
}

func (c *CreateClientSecretRequest) SetExpiresAfter(expiresAfter *types.ExpiresAfter) {
	c.ExpiresAfter = expiresAfter
}

func (c *CreateClientSecretRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (c CreateClientSecretRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateClientSecretRequest to string"
	}
	return string(jsonData)
}

func (c *CreateClientSecretRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}

func (c *CreateClientSecretRequest) MarshalJSON() ([]byte, error) {
	type embed CreateClientSecretRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

var (
	listConversationGroupsRealtimeRequestFieldPageSize = big.NewInt(1 << 0)
	listConversationGroupsRealtimeRequestFieldSearch   = big.NewInt(1 << 1)
	listConversationGroupsRealtimeRequestFieldCursor   = big.NewInt(1 << 2)
)

type ListConversationGroupsRealtimeRequest struct {
	PageSize *int    `json:"-" url:"page_size,omitempty"`
	Search   *string `json:"-" url:"search,omitempty"`
	Cursor   *string `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListConversationGroupsRealtimeRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationGroupsRealtimeRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listConversationGroupsRealtimeRequestFieldPageSize)
}

// SetSearch sets the Search field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationGroupsRealtimeRequest) SetSearch(search *string) {
	l.Search = search
	l.require(listConversationGroupsRealtimeRequestFieldSearch)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationGroupsRealtimeRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listConversationGroupsRealtimeRequestFieldCursor)
}
