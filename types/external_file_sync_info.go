package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"time"
)

// Tracks the link back to the original file in an external source.
type ExternalFileSyncInfo struct {
	Type ExternalSyncProvider `json:"type" xml:"type" required:"true"`
	// Entity ID in the external system
	SourceEntityId string `json:"source_entity_id" xml:"source_entity_id" required:"true"`
	// Integration connection instance ID
	IntegrationConnectionId string `json:"integration_connection_id" xml:"integration_connection_id" required:"true"`
	// Folder ID in the external system this file was synced from
	SourceParentEntityId string `json:"source_parent_entity_id" xml:"source_parent_entity_id" required:"true"`
	// Original MIME type in the external system
	SourceMimeType string `json:"source_mime_type" xml:"source_mime_type" required:"true"`
	// Last modified time from the external system
	SourceModifiedTime time.Time `json:"source_modified_time" xml:"source_modified_time" required:"true"`
	// KB folder ID of the sync root, used to query all entities under a sync tree
	RootFolderId    *string `json:"root_folder_id,omitempty" xml:"root_folder_id,omitempty"`
	extraProperties map[string]interface{}
}

func (e *ExternalFileSyncInfo) GetType() (value ExternalSyncProvider) {
	if e == nil {
		return
	}
	return e.Type
}

func (e *ExternalFileSyncInfo) GetSourceEntityId() (value string) {
	if e == nil {
		return
	}
	return e.SourceEntityId
}

func (e *ExternalFileSyncInfo) GetIntegrationConnectionId() (value string) {
	if e == nil {
		return
	}
	return e.IntegrationConnectionId
}

func (e *ExternalFileSyncInfo) GetSourceParentEntityId() (value string) {
	if e == nil {
		return
	}
	return e.SourceParentEntityId
}

func (e *ExternalFileSyncInfo) GetSourceMimeType() (value string) {
	if e == nil {
		return
	}
	return e.SourceMimeType
}

func (e *ExternalFileSyncInfo) GetSourceModifiedTime() (value time.Time) {
	if e == nil {
		return
	}
	return e.SourceModifiedTime
}

func (e *ExternalFileSyncInfo) GetRootFolderId() (value *string) {
	if e == nil {
		return
	}
	return e.RootFolderId
}

func (e *ExternalFileSyncInfo) SetType(type_ ExternalSyncProvider) {
	e.Type = type_
}

func (e *ExternalFileSyncInfo) SetSourceEntityId(sourceEntityId string) {
	e.SourceEntityId = sourceEntityId
}

func (e *ExternalFileSyncInfo) SetIntegrationConnectionId(integrationConnectionId string) {
	e.IntegrationConnectionId = integrationConnectionId
}

func (e *ExternalFileSyncInfo) SetSourceParentEntityId(sourceParentEntityId string) {
	e.SourceParentEntityId = sourceParentEntityId
}

func (e *ExternalFileSyncInfo) SetSourceMimeType(sourceMimeType string) {
	e.SourceMimeType = sourceMimeType
}

func (e *ExternalFileSyncInfo) SetSourceModifiedTime(sourceModifiedTime time.Time) {
	e.SourceModifiedTime = sourceModifiedTime
}

func (e *ExternalFileSyncInfo) SetRootFolderId(rootFolderId *string) {
	e.RootFolderId = rootFolderId
}

func (e *ExternalFileSyncInfo) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e ExternalFileSyncInfo) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ExternalFileSyncInfo to string"
	}
	return string(jsonData)
}

func (e *ExternalFileSyncInfo) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, e); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *ExternalFileSyncInfo) MarshalJSON() ([]byte, error) {
	type embed ExternalFileSyncInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
