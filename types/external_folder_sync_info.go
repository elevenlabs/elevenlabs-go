package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Metadata for a KB folder that mirrors an external source folder.
type ExternalFolderSyncInfo struct {
	Type ExternalSyncProvider `json:"type" xml:"type" required:"true"`
	// Entity ID in the external system
	SourceEntityId string `json:"source_entity_id" xml:"source_entity_id" required:"true"`
	// Integration connection instance ID
	IntegrationConnectionId string `json:"integration_connection_id" xml:"integration_connection_id" required:"true"`
	// KB folder ID of the sync root. None means this folder is the root.
	RootFolderId *string `json:"root_folder_id,omitempty" xml:"root_folder_id,omitempty"`
	// Opaque cursor for incremental sync, interpreted by the provider
	SyncCursor *string `json:"sync_cursor,omitempty" xml:"sync_cursor,omitempty"`
	// Unix timestamp of last completed sync
	LastSyncAt      *int `json:"last_sync_at,omitempty" xml:"last_sync_at,omitempty"`
	extraProperties map[string]interface{}
}

func (e *ExternalFolderSyncInfo) GetType() (value ExternalSyncProvider) {
	if e == nil {
		return
	}
	return e.Type
}

func (e *ExternalFolderSyncInfo) GetSourceEntityId() (value string) {
	if e == nil {
		return
	}
	return e.SourceEntityId
}

func (e *ExternalFolderSyncInfo) GetIntegrationConnectionId() (value string) {
	if e == nil {
		return
	}
	return e.IntegrationConnectionId
}

func (e *ExternalFolderSyncInfo) GetRootFolderId() (value *string) {
	if e == nil {
		return
	}
	return e.RootFolderId
}

func (e *ExternalFolderSyncInfo) GetSyncCursor() (value *string) {
	if e == nil {
		return
	}
	return e.SyncCursor
}

func (e *ExternalFolderSyncInfo) GetLastSyncAt() (value *int) {
	if e == nil {
		return
	}
	return e.LastSyncAt
}

func (e *ExternalFolderSyncInfo) SetType(type_ ExternalSyncProvider) {
	e.Type = type_
}

func (e *ExternalFolderSyncInfo) SetSourceEntityId(sourceEntityId string) {
	e.SourceEntityId = sourceEntityId
}

func (e *ExternalFolderSyncInfo) SetIntegrationConnectionId(integrationConnectionId string) {
	e.IntegrationConnectionId = integrationConnectionId
}

func (e *ExternalFolderSyncInfo) SetRootFolderId(rootFolderId *string) {
	e.RootFolderId = rootFolderId
}

func (e *ExternalFolderSyncInfo) SetSyncCursor(syncCursor *string) {
	e.SyncCursor = syncCursor
}

func (e *ExternalFolderSyncInfo) SetLastSyncAt(lastSyncAt *int) {
	e.LastSyncAt = lastSyncAt
}

func (e *ExternalFolderSyncInfo) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e ExternalFolderSyncInfo) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ExternalFolderSyncInfo to string"
	}
	return string(jsonData)
}

func (e *ExternalFolderSyncInfo) UnmarshalJSON(data []byte) error {
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

func (e *ExternalFolderSyncInfo) MarshalJSON() ([]byte, error) {
	type embed ExternalFolderSyncInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
