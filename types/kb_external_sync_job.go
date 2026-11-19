package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type KbExternalSyncJob struct {
	Type                    ExternalSyncProvider   `json:"type" xml:"type" required:"true"`
	FolderId                string                 `json:"folder_id" xml:"folder_id" required:"true"`
	IntegrationConnectionId string                 `json:"integration_connection_id" xml:"integration_connection_id" required:"true"`
	TriggeredBy             ExternalSyncJobTrigger `json:"triggered_by" xml:"triggered_by" required:"true"`
	Status                  *CrawlStatus           `json:"status,omitempty" xml:"status,omitempty"`
	SyncType                *ExternalSyncJobType   `json:"sync_type,omitempty" xml:"sync_type,omitempty"`
	ItemsIdentified         *int                   `json:"items_identified,omitempty" xml:"items_identified,omitempty"`
	ItemsProcessed          *int                   `json:"items_processed,omitempty" xml:"items_processed,omitempty"`
	ErrorMessage            *string                `json:"error_message,omitempty" xml:"error_message,omitempty"`
	StartedAt               *int                   `json:"started_at,omitempty" xml:"started_at,omitempty"`
	CompletedAt             *int                   `json:"completed_at,omitempty" xml:"completed_at,omitempty"`
	UpdatedAt               int                    `json:"updated_at" xml:"updated_at" required:"true"`
	Id                      string                 `json:"id" xml:"id" required:"true"`
	CreatedAt               int                    `json:"created_at" xml:"created_at" required:"true"`
	extraProperties         map[string]interface{}
}

func (k *KbExternalSyncJob) GetType() (value ExternalSyncProvider) {
	if k == nil {
		return
	}
	return k.Type
}

func (k *KbExternalSyncJob) GetFolderId() (value string) {
	if k == nil {
		return
	}
	return k.FolderId
}

func (k *KbExternalSyncJob) GetIntegrationConnectionId() (value string) {
	if k == nil {
		return
	}
	return k.IntegrationConnectionId
}

func (k *KbExternalSyncJob) GetTriggeredBy() (value ExternalSyncJobTrigger) {
	if k == nil {
		return
	}
	return k.TriggeredBy
}

func (k *KbExternalSyncJob) GetStatus() (value *CrawlStatus) {
	if k == nil {
		return
	}
	return k.Status
}

func (k *KbExternalSyncJob) GetSyncType() (value *ExternalSyncJobType) {
	if k == nil {
		return
	}
	return k.SyncType
}

func (k *KbExternalSyncJob) GetItemsIdentified() (value *int) {
	if k == nil {
		return
	}
	return k.ItemsIdentified
}

func (k *KbExternalSyncJob) GetItemsProcessed() (value *int) {
	if k == nil {
		return
	}
	return k.ItemsProcessed
}

func (k *KbExternalSyncJob) GetErrorMessage() (value *string) {
	if k == nil {
		return
	}
	return k.ErrorMessage
}

func (k *KbExternalSyncJob) GetStartedAt() (value *int) {
	if k == nil {
		return
	}
	return k.StartedAt
}

func (k *KbExternalSyncJob) GetCompletedAt() (value *int) {
	if k == nil {
		return
	}
	return k.CompletedAt
}

func (k *KbExternalSyncJob) GetUpdatedAt() (value int) {
	if k == nil {
		return
	}
	return k.UpdatedAt
}

func (k *KbExternalSyncJob) GetId() (value string) {
	if k == nil {
		return
	}
	return k.Id
}

func (k *KbExternalSyncJob) GetCreatedAt() (value int) {
	if k == nil {
		return
	}
	return k.CreatedAt
}

func (k *KbExternalSyncJob) SetType(type_ ExternalSyncProvider) {
	k.Type = type_
}

func (k *KbExternalSyncJob) SetFolderId(folderId string) {
	k.FolderId = folderId
}

func (k *KbExternalSyncJob) SetIntegrationConnectionId(integrationConnectionId string) {
	k.IntegrationConnectionId = integrationConnectionId
}

func (k *KbExternalSyncJob) SetTriggeredBy(triggeredBy ExternalSyncJobTrigger) {
	k.TriggeredBy = triggeredBy
}

func (k *KbExternalSyncJob) SetStatus(status *CrawlStatus) {
	k.Status = status
}

func (k *KbExternalSyncJob) SetSyncType(syncType *ExternalSyncJobType) {
	k.SyncType = syncType
}

func (k *KbExternalSyncJob) SetItemsIdentified(itemsIdentified *int) {
	k.ItemsIdentified = itemsIdentified
}

func (k *KbExternalSyncJob) SetItemsProcessed(itemsProcessed *int) {
	k.ItemsProcessed = itemsProcessed
}

func (k *KbExternalSyncJob) SetErrorMessage(errorMessage *string) {
	k.ErrorMessage = errorMessage
}

func (k *KbExternalSyncJob) SetStartedAt(startedAt *int) {
	k.StartedAt = startedAt
}

func (k *KbExternalSyncJob) SetCompletedAt(completedAt *int) {
	k.CompletedAt = completedAt
}

func (k *KbExternalSyncJob) SetUpdatedAt(updatedAt int) {
	k.UpdatedAt = updatedAt
}

func (k *KbExternalSyncJob) SetId(id string) {
	k.Id = id
}

func (k *KbExternalSyncJob) SetCreatedAt(createdAt int) {
	k.CreatedAt = createdAt
}

func (k *KbExternalSyncJob) GetExtraProperties() map[string]interface{} {
	if k == nil {
		return nil
	}
	return k.extraProperties
}

func (k KbExternalSyncJob) String() string {
	jsonData, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return "error converting struct: KbExternalSyncJob to string"
	}
	return string(jsonData)
}

func (k *KbExternalSyncJob) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, k); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *k)
	if err != nil {
		return err
	}
	k.extraProperties = extraProperties
	return nil
}

func (k *KbExternalSyncJob) MarshalJSON() ([]byte, error) {
	type embed KbExternalSyncJob
	var marshaler = struct {
		embed
	}{
		embed: embed(*k),
	}
	return json.Marshal(marshaler)
}
