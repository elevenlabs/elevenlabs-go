package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type RagDocumentIndexesResponseModel struct {
	Indexes         []*RagDocumentIndexResponseModel `json:"indexes" xml:"indexes" required:"true"`
	extraProperties map[string]interface{}
}

func (r *RagDocumentIndexesResponseModel) GetIndexes() (value []*RagDocumentIndexResponseModel) {
	if r == nil {
		return
	}
	return r.Indexes
}

func (r *RagDocumentIndexesResponseModel) SetIndexes(indexes []*RagDocumentIndexResponseModel) {
	r.Indexes = indexes
}

func (r *RagDocumentIndexesResponseModel) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RagDocumentIndexesResponseModel) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RagDocumentIndexesResponseModel to string"
	}
	return string(jsonData)
}

func (r *RagDocumentIndexesResponseModel) UnmarshalJSON(data []byte) error {
	type unmarshaler RagDocumentIndexesResponseModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RagDocumentIndexesResponseModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RagDocumentIndexesResponseModel) MarshalJSON() ([]byte, error) {
	type embed RagDocumentIndexesResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
