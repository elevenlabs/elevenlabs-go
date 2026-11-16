package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type RagDocumentIndexResponseModel struct {
	Id                      string                 `json:"id" xml:"id" required:"true"`
	Model                   EmbeddingModelEnum     `json:"model" xml:"model" required:"true"`
	Status                  RagIndexStatus         `json:"status" xml:"status" required:"true"`
	ProgressPercentage      float64                `json:"progress_percentage" xml:"progress_percentage" required:"true"`
	DocumentModelIndexUsage *RagDocumentIndexUsage `json:"document_model_index_usage" xml:"document_model_index_usage" required:"true"`
	extraProperties         map[string]interface{}
}

func (r *RagDocumentIndexResponseModel) GetId() (value string) {
	if r == nil {
		return
	}
	return r.Id
}

func (r *RagDocumentIndexResponseModel) GetModel() (value EmbeddingModelEnum) {
	if r == nil {
		return
	}
	return r.Model
}

func (r *RagDocumentIndexResponseModel) GetStatus() (value RagIndexStatus) {
	if r == nil {
		return
	}
	return r.Status
}

func (r *RagDocumentIndexResponseModel) GetProgressPercentage() (value float64) {
	if r == nil {
		return
	}
	return r.ProgressPercentage
}

func (r *RagDocumentIndexResponseModel) GetDocumentModelIndexUsage() (value *RagDocumentIndexUsage) {
	if r == nil {
		return
	}
	return r.DocumentModelIndexUsage
}

func (r *RagDocumentIndexResponseModel) SetId(id string) {
	r.Id = id
}

func (r *RagDocumentIndexResponseModel) SetModel(model EmbeddingModelEnum) {
	r.Model = model
}

func (r *RagDocumentIndexResponseModel) SetStatus(status RagIndexStatus) {
	r.Status = status
}

func (r *RagDocumentIndexResponseModel) SetProgressPercentage(progressPercentage float64) {
	r.ProgressPercentage = progressPercentage
}

func (r *RagDocumentIndexResponseModel) SetDocumentModelIndexUsage(documentModelIndexUsage *RagDocumentIndexUsage) {
	r.DocumentModelIndexUsage = documentModelIndexUsage
}

func (r *RagDocumentIndexResponseModel) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RagDocumentIndexResponseModel) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RagDocumentIndexResponseModel to string"
	}
	return string(jsonData)
}

func (r *RagDocumentIndexResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias RagDocumentIndexResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = RagDocumentIndexResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RagDocumentIndexResponseModel) MarshalJSON() ([]byte, error) {
	type embed RagDocumentIndexResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
