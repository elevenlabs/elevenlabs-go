package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type RemoveOrderItemResponse struct {
	// Whether the item was successfully removed.
	Success         bool `json:"success" xml:"success" required:"true"`
	extraProperties map[string]interface{}
}

func (r *RemoveOrderItemResponse) GetSuccess() (value bool) {
	if r == nil {
		return
	}
	return r.Success
}

func (r *RemoveOrderItemResponse) SetSuccess(success bool) {
	r.Success = success
}

func (r *RemoveOrderItemResponse) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RemoveOrderItemResponse) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RemoveOrderItemResponse to string"
	}
	return string(jsonData)
}

func (r *RemoveOrderItemResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias RemoveOrderItemResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = RemoveOrderItemResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RemoveOrderItemResponse) MarshalJSON() ([]byte, error) {
	type embed RemoveOrderItemResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
