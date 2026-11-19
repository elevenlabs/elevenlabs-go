package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type RealtimeConversationGroupResponseModel struct {
	Id          string              `json:"id" xml:"id" required:"true"`
	Title       string              `json:"title" xml:"title" required:"true"`
	Description *string             `json:"description" xml:"description" required:"true"`
	AccessInfo  *ResourceAccessInfo `json:"access_info" xml:"access_info" required:"true"`
	// How widely a group is shared, besides its owner and workspace admins (who always have access).
	Sharing           *RealtimeConversationGroupSharingSummary `json:"sharing" xml:"sharing" required:"true"`
	CreatedAtUnixSecs int                                      `json:"created_at_unix_secs" xml:"created_at_unix_secs" required:"true"`
	extraProperties   map[string]interface{}
}

func (r *RealtimeConversationGroupResponseModel) GetId() (value string) {
	if r == nil {
		return
	}
	return r.Id
}

func (r *RealtimeConversationGroupResponseModel) GetTitle() (value string) {
	if r == nil {
		return
	}
	return r.Title
}

func (r *RealtimeConversationGroupResponseModel) GetDescription() (value *string) {
	if r == nil {
		return
	}
	return r.Description
}

func (r *RealtimeConversationGroupResponseModel) GetAccessInfo() (value *ResourceAccessInfo) {
	if r == nil {
		return
	}
	return r.AccessInfo
}

func (r *RealtimeConversationGroupResponseModel) GetSharing() (value *RealtimeConversationGroupSharingSummary) {
	if r == nil {
		return
	}
	return r.Sharing
}

func (r *RealtimeConversationGroupResponseModel) GetCreatedAtUnixSecs() (value int) {
	if r == nil {
		return
	}
	return r.CreatedAtUnixSecs
}

func (r *RealtimeConversationGroupResponseModel) SetId(id string) {
	r.Id = id
}

func (r *RealtimeConversationGroupResponseModel) SetTitle(title string) {
	r.Title = title
}

func (r *RealtimeConversationGroupResponseModel) SetDescription(description *string) {
	r.Description = description
}

func (r *RealtimeConversationGroupResponseModel) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	r.AccessInfo = accessInfo
}

func (r *RealtimeConversationGroupResponseModel) SetSharing(sharing *RealtimeConversationGroupSharingSummary) {
	r.Sharing = sharing
}

func (r *RealtimeConversationGroupResponseModel) SetCreatedAtUnixSecs(createdAtUnixSecs int) {
	r.CreatedAtUnixSecs = createdAtUnixSecs
}

func (r *RealtimeConversationGroupResponseModel) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RealtimeConversationGroupResponseModel) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RealtimeConversationGroupResponseModel to string"
	}
	return string(jsonData)
}

func (r *RealtimeConversationGroupResponseModel) UnmarshalJSON(data []byte) error {
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

func (r *RealtimeConversationGroupResponseModel) MarshalJSON() ([]byte, error) {
	type embed RealtimeConversationGroupResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
