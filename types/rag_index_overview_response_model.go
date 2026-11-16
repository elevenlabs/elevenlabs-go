package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type RagIndexOverviewResponseModel struct {
	TotalUsedBytes  int                                            `json:"total_used_bytes" xml:"total_used_bytes" required:"true"`
	TotalMaxBytes   int                                            `json:"total_max_bytes" xml:"total_max_bytes" required:"true"`
	Models          []*RagIndexOverviewEmbeddingModelResponseModel `json:"models" xml:"models" required:"true"`
	extraProperties map[string]interface{}
}

func (r *RagIndexOverviewResponseModel) GetTotalUsedBytes() (value int) {
	if r == nil {
		return
	}
	return r.TotalUsedBytes
}

func (r *RagIndexOverviewResponseModel) GetTotalMaxBytes() (value int) {
	if r == nil {
		return
	}
	return r.TotalMaxBytes
}

func (r *RagIndexOverviewResponseModel) GetModels() (value []*RagIndexOverviewEmbeddingModelResponseModel) {
	if r == nil {
		return
	}
	return r.Models
}

func (r *RagIndexOverviewResponseModel) SetTotalUsedBytes(totalUsedBytes int) {
	r.TotalUsedBytes = totalUsedBytes
}

func (r *RagIndexOverviewResponseModel) SetTotalMaxBytes(totalMaxBytes int) {
	r.TotalMaxBytes = totalMaxBytes
}

func (r *RagIndexOverviewResponseModel) SetModels(models []*RagIndexOverviewEmbeddingModelResponseModel) {
	r.Models = models
}

func (r *RagIndexOverviewResponseModel) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RagIndexOverviewResponseModel) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RagIndexOverviewResponseModel to string"
	}
	return string(jsonData)
}

func (r *RagIndexOverviewResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias RagIndexOverviewResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = RagIndexOverviewResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RagIndexOverviewResponseModel) MarshalJSON() ([]byte, error) {
	type embed RagIndexOverviewResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
