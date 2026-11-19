package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type WorkspaceServiceAccountListResponseModel struct {
	ServiceAccounts []*WorkspaceServiceAccountResponseModel `json:"service-accounts" xml:"service-accounts" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkspaceServiceAccountListResponseModel) GetServiceAccounts() (value []*WorkspaceServiceAccountResponseModel) {
	if w == nil {
		return
	}
	return w.ServiceAccounts
}

func (w *WorkspaceServiceAccountListResponseModel) SetServiceAccounts(serviceAccounts []*WorkspaceServiceAccountResponseModel) {
	w.ServiceAccounts = serviceAccounts
}

func (w *WorkspaceServiceAccountListResponseModel) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkspaceServiceAccountListResponseModel) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkspaceServiceAccountListResponseModel to string"
	}
	return string(jsonData)
}

func (w *WorkspaceServiceAccountListResponseModel) UnmarshalJSON(data []byte) error {
	type unmarshaler WorkspaceServiceAccountListResponseModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WorkspaceServiceAccountListResponseModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkspaceServiceAccountListResponseModel) MarshalJSON() ([]byte, error) {
	type embed WorkspaceServiceAccountListResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
