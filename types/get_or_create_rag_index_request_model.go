package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetOrCreateRagIndexRequestModel struct {
	// ID of the knowledgebase document for which to retrieve the index
	DocumentId string `json:"document_id" xml:"document_id" required:"true"`
	// Whether to create the RAG index if it does not exist
	CreateIfMissing bool               `json:"create_if_missing" xml:"create_if_missing" required:"true"`
	Model           EmbeddingModelEnum `json:"model" xml:"model" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetOrCreateRagIndexRequestModel) GetDocumentId() (value string) {
	if g == nil {
		return
	}
	return g.DocumentId
}

func (g *GetOrCreateRagIndexRequestModel) GetCreateIfMissing() (value bool) {
	if g == nil {
		return
	}
	return g.CreateIfMissing
}

func (g *GetOrCreateRagIndexRequestModel) GetModel() (value EmbeddingModelEnum) {
	if g == nil {
		return
	}
	return g.Model
}

func (g *GetOrCreateRagIndexRequestModel) SetDocumentId(documentId string) {
	g.DocumentId = documentId
}

func (g *GetOrCreateRagIndexRequestModel) SetCreateIfMissing(createIfMissing bool) {
	g.CreateIfMissing = createIfMissing
}

func (g *GetOrCreateRagIndexRequestModel) SetModel(model EmbeddingModelEnum) {
	g.Model = model
}

func (g *GetOrCreateRagIndexRequestModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetOrCreateRagIndexRequestModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetOrCreateRagIndexRequestModel to string"
	}
	return string(jsonData)
}

func (g *GetOrCreateRagIndexRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GetOrCreateRagIndexRequestModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GetOrCreateRagIndexRequestModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetOrCreateRagIndexRequestModel) MarshalJSON() ([]byte, error) {
	type embed GetOrCreateRagIndexRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
