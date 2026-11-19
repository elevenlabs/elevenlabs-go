package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkspaceServiceAccountResponseModel struct {
	ServiceAccountUserId string                              `json:"service_account_user_id" xml:"service_account_user_id" required:"true"`
	Name                 string                              `json:"name" xml:"name" required:"true"`
	CreatedAtUnix        *int                                `json:"created_at_unix,omitempty" xml:"created_at_unix,omitempty"`
	ApiKeys              []*WorkspaceApiKeyResponseModel     `json:"api-keys" xml:"api-keys" required:"true"`
	DefaultSharingGroups []*DefaultSharingGroupResponseModel `json:"default_sharing_groups,omitempty" xml:"default_sharing_groups,omitempty"`
	extraProperties      map[string]interface{}
}

func (w *WorkspaceServiceAccountResponseModel) GetServiceAccountUserId() (value string) {
	if w == nil {
		return
	}
	return w.ServiceAccountUserId
}

func (w *WorkspaceServiceAccountResponseModel) GetName() (value string) {
	if w == nil {
		return
	}
	return w.Name
}

func (w *WorkspaceServiceAccountResponseModel) GetCreatedAtUnix() (value *int) {
	if w == nil {
		return
	}
	return w.CreatedAtUnix
}

func (w *WorkspaceServiceAccountResponseModel) GetApiKeys() (value []*WorkspaceApiKeyResponseModel) {
	if w == nil {
		return
	}
	return w.ApiKeys
}

func (w *WorkspaceServiceAccountResponseModel) GetDefaultSharingGroups() (value []*DefaultSharingGroupResponseModel) {
	if w == nil {
		return
	}
	return w.DefaultSharingGroups
}

func (w *WorkspaceServiceAccountResponseModel) SetServiceAccountUserId(serviceAccountUserId string) {
	w.ServiceAccountUserId = serviceAccountUserId
}

func (w *WorkspaceServiceAccountResponseModel) SetName(name string) {
	w.Name = name
}

func (w *WorkspaceServiceAccountResponseModel) SetCreatedAtUnix(createdAtUnix *int) {
	w.CreatedAtUnix = createdAtUnix
}

func (w *WorkspaceServiceAccountResponseModel) SetApiKeys(apiKeys []*WorkspaceApiKeyResponseModel) {
	w.ApiKeys = apiKeys
}

func (w *WorkspaceServiceAccountResponseModel) SetDefaultSharingGroups(defaultSharingGroups []*DefaultSharingGroupResponseModel) {
	w.DefaultSharingGroups = defaultSharingGroups
}

func (w *WorkspaceServiceAccountResponseModel) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkspaceServiceAccountResponseModel) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkspaceServiceAccountResponseModel to string"
	}
	return string(jsonData)
}

func (w *WorkspaceServiceAccountResponseModel) UnmarshalJSON(data []byte) error {
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

func (w *WorkspaceServiceAccountResponseModel) MarshalJSON() ([]byte, error) {
	type embed WorkspaceServiceAccountResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
