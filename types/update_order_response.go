package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UpdateOrderResponse struct {
	// The updated order name.
	Name            string `json:"name" xml:"name" required:"true"`
	extraProperties map[string]interface{}
}

func (u *UpdateOrderResponse) GetName() (value string) {
	if u == nil {
		return
	}
	return u.Name
}

func (u *UpdateOrderResponse) SetName(name string) {
	u.Name = name
}

func (u *UpdateOrderResponse) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateOrderResponse) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateOrderResponse to string"
	}
	return string(jsonData)
}

func (u *UpdateOrderResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, u); err != nil {
		return err
	}
	type alias UpdateOrderResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*u = UpdateOrderResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateOrderResponse) MarshalJSON() ([]byte, error) {
	type embed UpdateOrderResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
