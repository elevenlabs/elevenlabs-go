package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ArchitectDispatchResponseModel struct {
	ArchitectChatId      string `json:"architect_chat_id" xml:"architect_chat_id" required:"true"`
	UserId               string `json:"user_id" xml:"user_id" required:"true"`
	DispatchedAtUnixSecs int    `json:"dispatched_at_unix_secs" xml:"dispatched_at_unix_secs" required:"true"`
	extraProperties      map[string]interface{}
}

func (a *ArchitectDispatchResponseModel) GetArchitectChatId() (value string) {
	if a == nil {
		return
	}
	return a.ArchitectChatId
}

func (a *ArchitectDispatchResponseModel) GetUserId() (value string) {
	if a == nil {
		return
	}
	return a.UserId
}

func (a *ArchitectDispatchResponseModel) GetDispatchedAtUnixSecs() (value int) {
	if a == nil {
		return
	}
	return a.DispatchedAtUnixSecs
}

func (a *ArchitectDispatchResponseModel) SetArchitectChatId(architectChatId string) {
	a.ArchitectChatId = architectChatId
}

func (a *ArchitectDispatchResponseModel) SetUserId(userId string) {
	a.UserId = userId
}

func (a *ArchitectDispatchResponseModel) SetDispatchedAtUnixSecs(dispatchedAtUnixSecs int) {
	a.DispatchedAtUnixSecs = dispatchedAtUnixSecs
}

func (a *ArchitectDispatchResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a ArchitectDispatchResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: ArchitectDispatchResponseModel to string"
	}
	return string(jsonData)
}

func (a *ArchitectDispatchResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias ArchitectDispatchResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = ArchitectDispatchResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *ArchitectDispatchResponseModel) MarshalJSON() ([]byte, error) {
	type embed ArchitectDispatchResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
