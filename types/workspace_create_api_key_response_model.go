package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkspaceCreateApiKeyResponseModel struct {
	XiApiKey        string `json:"xi-api-key" xml:"xi-api-key" required:"true"`
	KeyId           string `json:"key_id" xml:"key_id" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkspaceCreateApiKeyResponseModel) GetXiApiKey() (value string) {
	if w == nil {
		return
	}
	return w.XiApiKey
}

func (w *WorkspaceCreateApiKeyResponseModel) GetKeyId() (value string) {
	if w == nil {
		return
	}
	return w.KeyId
}

func (w *WorkspaceCreateApiKeyResponseModel) SetXiApiKey(xiApiKey string) {
	w.XiApiKey = xiApiKey
}

func (w *WorkspaceCreateApiKeyResponseModel) SetKeyId(keyId string) {
	w.KeyId = keyId
}

func (w *WorkspaceCreateApiKeyResponseModel) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkspaceCreateApiKeyResponseModel) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkspaceCreateApiKeyResponseModel to string"
	}
	return string(jsonData)
}

func (w *WorkspaceCreateApiKeyResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WorkspaceCreateApiKeyResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WorkspaceCreateApiKeyResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkspaceCreateApiKeyResponseModel) MarshalJSON() ([]byte, error) {
	type embed WorkspaceCreateApiKeyResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
