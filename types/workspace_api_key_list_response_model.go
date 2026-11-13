package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type WorkspaceApiKeyListResponseModel struct {
	ApiKeys         []*WorkspaceApiKeyResponseModel `json:"api-keys" xml:"api-keys" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkspaceApiKeyListResponseModel) GetApiKeys() (value []*WorkspaceApiKeyResponseModel) {
	if w == nil {
		return
	}
	return w.ApiKeys
}

func (w *WorkspaceApiKeyListResponseModel) SetApiKeys(apiKeys []*WorkspaceApiKeyResponseModel) {
	w.ApiKeys = apiKeys
}

func (w *WorkspaceApiKeyListResponseModel) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkspaceApiKeyListResponseModel) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkspaceApiKeyListResponseModel to string"
	}
	return string(jsonData)
}

func (w *WorkspaceApiKeyListResponseModel) UnmarshalJSON(data []byte) error {
	type unmarshaler WorkspaceApiKeyListResponseModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WorkspaceApiKeyListResponseModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkspaceApiKeyListResponseModel) MarshalJSON() ([]byte, error) {
	type embed WorkspaceApiKeyListResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
