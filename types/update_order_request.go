package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UpdateOrderRequest struct {
	// The new name for the order.
	Name            string `json:"name" xml:"name" required:"true" maxLength:"120" minLength:"1"`
	extraProperties map[string]interface{}
}

func (u *UpdateOrderRequest) GetName() (value string) {
	if u == nil {
		return
	}
	return u.Name
}

func (u *UpdateOrderRequest) SetName(name string) {
	u.Name = name
}

func (u *UpdateOrderRequest) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateOrderRequest) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateOrderRequest to string"
	}
	return string(jsonData)
}

func (u *UpdateOrderRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, u); err != nil {
		return err
	}
	type alias UpdateOrderRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*u = UpdateOrderRequest(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateOrderRequest) MarshalJSON() ([]byte, error) {
	type embed UpdateOrderRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
