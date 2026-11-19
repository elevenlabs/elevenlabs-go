package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UpdateMemberUsageLimitResponseModel struct {
	Status          string `json:"status" xml:"status" required:"true"`
	extraProperties map[string]interface{}
}

func (u *UpdateMemberUsageLimitResponseModel) GetStatus() (value string) {
	if u == nil {
		return
	}
	return u.Status
}

func (u *UpdateMemberUsageLimitResponseModel) SetStatus(status string) {
	u.Status = status
}

func (u *UpdateMemberUsageLimitResponseModel) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateMemberUsageLimitResponseModel) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateMemberUsageLimitResponseModel to string"
	}
	return string(jsonData)
}

func (u *UpdateMemberUsageLimitResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, u); err != nil {
		return err
	}
	type alias UpdateMemberUsageLimitResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*u = UpdateMemberUsageLimitResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateMemberUsageLimitResponseModel) MarshalJSON() ([]byte, error) {
	type embed UpdateMemberUsageLimitResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
