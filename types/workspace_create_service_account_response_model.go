package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkspaceCreateServiceAccountResponseModel struct {
	ServiceAccountUserId string `json:"service-account-user-id" xml:"service-account-user-id" required:"true"`
	extraProperties      map[string]interface{}
}

func (w *WorkspaceCreateServiceAccountResponseModel) GetServiceAccountUserId() (value string) {
	if w == nil {
		return
	}
	return w.ServiceAccountUserId
}

func (w *WorkspaceCreateServiceAccountResponseModel) SetServiceAccountUserId(serviceAccountUserId string) {
	w.ServiceAccountUserId = serviceAccountUserId
}

func (w *WorkspaceCreateServiceAccountResponseModel) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkspaceCreateServiceAccountResponseModel) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkspaceCreateServiceAccountResponseModel to string"
	}
	return string(jsonData)
}

func (w *WorkspaceCreateServiceAccountResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WorkspaceCreateServiceAccountResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WorkspaceCreateServiceAccountResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkspaceCreateServiceAccountResponseModel) MarshalJSON() ([]byte, error) {
	type embed WorkspaceCreateServiceAccountResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
