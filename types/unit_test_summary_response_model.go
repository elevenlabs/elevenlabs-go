package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UnitTestSummaryResponseModel struct {
	// The ID of the test
	Id string `json:"id" xml:"id" required:"true"`
	// Name of the test
	Name string `json:"name" xml:"name" required:"true"`
	// The access information of the test
	AccessInfo *ResourceAccessInfo `json:"access_info,omitempty" xml:"access_info,omitempty"`
	// Creation time of the test in unix seconds
	CreatedAtUnixSecs int `json:"created_at_unix_secs" xml:"created_at_unix_secs" required:"true"`
	// Last update time of the test in unix seconds
	LastUpdatedAtUnixSecs int                  `json:"last_updated_at_unix_secs" xml:"last_updated_at_unix_secs" required:"true"`
	Type                  TestType             `json:"type" xml:"type" required:"true"`
	EntityType            *AgentTestEntityType `json:"entity_type,omitempty" xml:"entity_type,omitempty"`
	// The ID of the parent folder
	FolderParentId *string `json:"folder_parent_id,omitempty" xml:"folder_parent_id,omitempty"`
	// The folder path segments from root to this entity
	FolderPath []*AgentTestFolderPathSegmentResponseModel `json:"folder_path,omitempty" xml:"folder_path,omitempty"`
	// Number of direct children (tests and subfolders) for folders only
	ChildrenCount *int `json:"children_count,omitempty" xml:"children_count,omitempty"`
	// Channel the test simulates the conversation as. Null for folders or default behavior.
	ConversationInitiationSource *ConversationInitiationSource `json:"conversation_initiation_source,omitempty" xml:"conversation_initiation_source,omitempty"`
	extraProperties              map[string]interface{}
}

func (u *UnitTestSummaryResponseModel) GetId() (value string) {
	if u == nil {
		return
	}
	return u.Id
}

func (u *UnitTestSummaryResponseModel) GetName() (value string) {
	if u == nil {
		return
	}
	return u.Name
}

func (u *UnitTestSummaryResponseModel) GetAccessInfo() (value *ResourceAccessInfo) {
	if u == nil {
		return
	}
	return u.AccessInfo
}

func (u *UnitTestSummaryResponseModel) GetCreatedAtUnixSecs() (value int) {
	if u == nil {
		return
	}
	return u.CreatedAtUnixSecs
}

func (u *UnitTestSummaryResponseModel) GetLastUpdatedAtUnixSecs() (value int) {
	if u == nil {
		return
	}
	return u.LastUpdatedAtUnixSecs
}

func (u *UnitTestSummaryResponseModel) GetType() (value TestType) {
	if u == nil {
		return
	}
	return u.Type
}

func (u *UnitTestSummaryResponseModel) GetEntityType() (value *AgentTestEntityType) {
	if u == nil {
		return
	}
	return u.EntityType
}

func (u *UnitTestSummaryResponseModel) GetFolderParentId() (value *string) {
	if u == nil {
		return
	}
	return u.FolderParentId
}

func (u *UnitTestSummaryResponseModel) GetFolderPath() (value []*AgentTestFolderPathSegmentResponseModel) {
	if u == nil {
		return
	}
	return u.FolderPath
}

func (u *UnitTestSummaryResponseModel) GetChildrenCount() (value *int) {
	if u == nil {
		return
	}
	return u.ChildrenCount
}

func (u *UnitTestSummaryResponseModel) GetConversationInitiationSource() (value *ConversationInitiationSource) {
	if u == nil {
		return
	}
	return u.ConversationInitiationSource
}

func (u *UnitTestSummaryResponseModel) SetId(id string) {
	u.Id = id
}

func (u *UnitTestSummaryResponseModel) SetName(name string) {
	u.Name = name
}

func (u *UnitTestSummaryResponseModel) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	u.AccessInfo = accessInfo
}

func (u *UnitTestSummaryResponseModel) SetCreatedAtUnixSecs(createdAtUnixSecs int) {
	u.CreatedAtUnixSecs = createdAtUnixSecs
}

func (u *UnitTestSummaryResponseModel) SetLastUpdatedAtUnixSecs(lastUpdatedAtUnixSecs int) {
	u.LastUpdatedAtUnixSecs = lastUpdatedAtUnixSecs
}

func (u *UnitTestSummaryResponseModel) SetType(type_ TestType) {
	u.Type = type_
}

func (u *UnitTestSummaryResponseModel) SetEntityType(entityType *AgentTestEntityType) {
	u.EntityType = entityType
}

func (u *UnitTestSummaryResponseModel) SetFolderParentId(folderParentId *string) {
	u.FolderParentId = folderParentId
}

func (u *UnitTestSummaryResponseModel) SetFolderPath(folderPath []*AgentTestFolderPathSegmentResponseModel) {
	u.FolderPath = folderPath
}

func (u *UnitTestSummaryResponseModel) SetChildrenCount(childrenCount *int) {
	u.ChildrenCount = childrenCount
}

func (u *UnitTestSummaryResponseModel) SetConversationInitiationSource(conversationInitiationSource *ConversationInitiationSource) {
	u.ConversationInitiationSource = conversationInitiationSource
}

func (u *UnitTestSummaryResponseModel) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UnitTestSummaryResponseModel) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UnitTestSummaryResponseModel to string"
	}
	return string(jsonData)
}

func (u *UnitTestSummaryResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, u); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UnitTestSummaryResponseModel) MarshalJSON() ([]byte, error) {
	type embed UnitTestSummaryResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
