package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Paginated workspace audit log response.
type WorkspaceAuditLogsPageResponse struct {
	Entries         []*WorkspaceAuditLogEntryResponse `json:"entries" xml:"entries" required:"true"`
	HasMore         bool                              `json:"has_more" xml:"has_more" required:"true"`
	NextCursor      *string                           `json:"next_cursor" xml:"next_cursor" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkspaceAuditLogsPageResponse) GetEntries() (value []*WorkspaceAuditLogEntryResponse) {
	if w == nil {
		return
	}
	return w.Entries
}

func (w *WorkspaceAuditLogsPageResponse) GetHasMore() (value bool) {
	if w == nil {
		return
	}
	return w.HasMore
}

func (w *WorkspaceAuditLogsPageResponse) GetNextCursor() (value *string) {
	if w == nil {
		return
	}
	return w.NextCursor
}

func (w *WorkspaceAuditLogsPageResponse) SetEntries(entries []*WorkspaceAuditLogEntryResponse) {
	w.Entries = entries
}

func (w *WorkspaceAuditLogsPageResponse) SetHasMore(hasMore bool) {
	w.HasMore = hasMore
}

func (w *WorkspaceAuditLogsPageResponse) SetNextCursor(nextCursor *string) {
	w.NextCursor = nextCursor
}

func (w *WorkspaceAuditLogsPageResponse) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkspaceAuditLogsPageResponse) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkspaceAuditLogsPageResponse to string"
	}
	return string(jsonData)
}

func (w *WorkspaceAuditLogsPageResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkspaceAuditLogsPageResponse) MarshalJSON() ([]byte, error) {
	type embed WorkspaceAuditLogsPageResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
