package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ReaderResourceResponseModel struct {
	// The type of resource.
	ResourceType ReaderResourceResponseModelResourceType `json:"resource_type" xml:"resource_type" required:"true"`
	// The ID of the resource.
	ResourceId      string `json:"resource_id" xml:"resource_id" required:"true"`
	extraProperties map[string]interface{}
}

func (r *ReaderResourceResponseModel) GetResourceType() (value ReaderResourceResponseModelResourceType) {
	if r == nil {
		return
	}
	return r.ResourceType
}

func (r *ReaderResourceResponseModel) GetResourceId() (value string) {
	if r == nil {
		return
	}
	return r.ResourceId
}

func (r *ReaderResourceResponseModel) SetResourceType(resourceType ReaderResourceResponseModelResourceType) {
	r.ResourceType = resourceType
}

func (r *ReaderResourceResponseModel) SetResourceId(resourceId string) {
	r.ResourceId = resourceId
}

func (r *ReaderResourceResponseModel) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReaderResourceResponseModel) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReaderResourceResponseModel to string"
	}
	return string(jsonData)
}

func (r *ReaderResourceResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias ReaderResourceResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = ReaderResourceResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ReaderResourceResponseModel) MarshalJSON() ([]byte, error) {
	type embed ReaderResourceResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
