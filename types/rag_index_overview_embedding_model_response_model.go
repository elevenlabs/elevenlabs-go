package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type RagIndexOverviewEmbeddingModelResponseModel struct {
	Model           EmbeddingModelEnum `json:"model" xml:"model" required:"true"`
	UsedBytes       int                `json:"used_bytes" xml:"used_bytes" required:"true"`
	extraProperties map[string]interface{}
}

func (r *RagIndexOverviewEmbeddingModelResponseModel) GetModel() (value EmbeddingModelEnum) {
	if r == nil {
		return
	}
	return r.Model
}

func (r *RagIndexOverviewEmbeddingModelResponseModel) GetUsedBytes() (value int) {
	if r == nil {
		return
	}
	return r.UsedBytes
}

func (r *RagIndexOverviewEmbeddingModelResponseModel) SetModel(model EmbeddingModelEnum) {
	r.Model = model
}

func (r *RagIndexOverviewEmbeddingModelResponseModel) SetUsedBytes(usedBytes int) {
	r.UsedBytes = usedBytes
}

func (r *RagIndexOverviewEmbeddingModelResponseModel) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RagIndexOverviewEmbeddingModelResponseModel) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RagIndexOverviewEmbeddingModelResponseModel to string"
	}
	return string(jsonData)
}

func (r *RagIndexOverviewEmbeddingModelResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias RagIndexOverviewEmbeddingModelResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = RagIndexOverviewEmbeddingModelResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RagIndexOverviewEmbeddingModelResponseModel) MarshalJSON() ([]byte, error) {
	type embed RagIndexOverviewEmbeddingModelResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
