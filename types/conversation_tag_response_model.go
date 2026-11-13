package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationTagResponseModel struct {
	TagId             string  `json:"tag_id" xml:"tag_id" required:"true"`
	WorkspaceId       string  `json:"workspace_id" xml:"workspace_id" required:"true"`
	OwnerUserId       string  `json:"owner_user_id" xml:"owner_user_id" required:"true"`
	Title             string  `json:"title" xml:"title" required:"true"`
	Description       *string `json:"description" xml:"description" required:"true"`
	CreatedAtUnixSecs int     `json:"created_at_unix_secs" xml:"created_at_unix_secs" required:"true"`
	extraProperties   map[string]interface{}
}

func (c *ConversationTagResponseModel) GetTagId() (value string) {
	if c == nil {
		return
	}
	return c.TagId
}

func (c *ConversationTagResponseModel) GetWorkspaceId() (value string) {
	if c == nil {
		return
	}
	return c.WorkspaceId
}

func (c *ConversationTagResponseModel) GetOwnerUserId() (value string) {
	if c == nil {
		return
	}
	return c.OwnerUserId
}

func (c *ConversationTagResponseModel) GetTitle() (value string) {
	if c == nil {
		return
	}
	return c.Title
}

func (c *ConversationTagResponseModel) GetDescription() (value *string) {
	if c == nil {
		return
	}
	return c.Description
}

func (c *ConversationTagResponseModel) GetCreatedAtUnixSecs() (value int) {
	if c == nil {
		return
	}
	return c.CreatedAtUnixSecs
}

func (c *ConversationTagResponseModel) SetTagId(tagId string) {
	c.TagId = tagId
}

func (c *ConversationTagResponseModel) SetWorkspaceId(workspaceId string) {
	c.WorkspaceId = workspaceId
}

func (c *ConversationTagResponseModel) SetOwnerUserId(ownerUserId string) {
	c.OwnerUserId = ownerUserId
}

func (c *ConversationTagResponseModel) SetTitle(title string) {
	c.Title = title
}

func (c *ConversationTagResponseModel) SetDescription(description *string) {
	c.Description = description
}

func (c *ConversationTagResponseModel) SetCreatedAtUnixSecs(createdAtUnixSecs int) {
	c.CreatedAtUnixSecs = createdAtUnixSecs
}

func (c *ConversationTagResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationTagResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationTagResponseModel to string"
	}
	return string(jsonData)
}

func (c *ConversationTagResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationTagResponseModel) MarshalJSON() ([]byte, error) {
	type embed ConversationTagResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
