package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// In-flight/last refresh state for an externally-synced KB file.
type FileRefreshStatus struct {
	Status          *CrawlStatus `json:"status,omitempty" xml:"status,omitempty"`
	EnqueuedAt      *int         `json:"enqueued_at,omitempty" xml:"enqueued_at,omitempty"`
	StartedAt       *int         `json:"started_at,omitempty" xml:"started_at,omitempty"`
	CompletedAt     *int         `json:"completed_at,omitempty" xml:"completed_at,omitempty"`
	LastSyncedAt    *int         `json:"last_synced_at,omitempty" xml:"last_synced_at,omitempty"`
	ErrorMessage    *string      `json:"error_message,omitempty" xml:"error_message,omitempty"`
	extraProperties map[string]interface{}
}

func (f *FileRefreshStatus) GetStatus() (value *CrawlStatus) {
	if f == nil {
		return
	}
	return f.Status
}

func (f *FileRefreshStatus) GetEnqueuedAt() (value *int) {
	if f == nil {
		return
	}
	return f.EnqueuedAt
}

func (f *FileRefreshStatus) GetStartedAt() (value *int) {
	if f == nil {
		return
	}
	return f.StartedAt
}

func (f *FileRefreshStatus) GetCompletedAt() (value *int) {
	if f == nil {
		return
	}
	return f.CompletedAt
}

func (f *FileRefreshStatus) GetLastSyncedAt() (value *int) {
	if f == nil {
		return
	}
	return f.LastSyncedAt
}

func (f *FileRefreshStatus) GetErrorMessage() (value *string) {
	if f == nil {
		return
	}
	return f.ErrorMessage
}

func (f *FileRefreshStatus) SetStatus(status *CrawlStatus) {
	f.Status = status
}

func (f *FileRefreshStatus) SetEnqueuedAt(enqueuedAt *int) {
	f.EnqueuedAt = enqueuedAt
}

func (f *FileRefreshStatus) SetStartedAt(startedAt *int) {
	f.StartedAt = startedAt
}

func (f *FileRefreshStatus) SetCompletedAt(completedAt *int) {
	f.CompletedAt = completedAt
}

func (f *FileRefreshStatus) SetLastSyncedAt(lastSyncedAt *int) {
	f.LastSyncedAt = lastSyncedAt
}

func (f *FileRefreshStatus) SetErrorMessage(errorMessage *string) {
	f.ErrorMessage = errorMessage
}

func (f *FileRefreshStatus) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f FileRefreshStatus) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FileRefreshStatus to string"
	}
	return string(jsonData)
}

func (f *FileRefreshStatus) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, f); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *FileRefreshStatus) MarshalJSON() ([]byte, error) {
	type embed FileRefreshStatus
	var marshaler = struct {
		embed
	}{
		embed: embed(*f),
	}
	return json.Marshal(marshaler)
}
