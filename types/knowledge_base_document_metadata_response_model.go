package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type KnowledgeBaseDocumentMetadataResponseModel struct {
	CreatedAtUnixSecs     int `json:"created_at_unix_secs" xml:"created_at_unix_secs" required:"true"`
	LastUpdatedAtUnixSecs int `json:"last_updated_at_unix_secs" xml:"last_updated_at_unix_secs" required:"true"`
	SizeBytes             int `json:"size_bytes" xml:"size_bytes" required:"true"`
	extraProperties       map[string]interface{}
}

func (k *KnowledgeBaseDocumentMetadataResponseModel) GetCreatedAtUnixSecs() (value int) {
	if k == nil {
		return
	}
	return k.CreatedAtUnixSecs
}

func (k *KnowledgeBaseDocumentMetadataResponseModel) GetLastUpdatedAtUnixSecs() (value int) {
	if k == nil {
		return
	}
	return k.LastUpdatedAtUnixSecs
}

func (k *KnowledgeBaseDocumentMetadataResponseModel) GetSizeBytes() (value int) {
	if k == nil {
		return
	}
	return k.SizeBytes
}

func (k *KnowledgeBaseDocumentMetadataResponseModel) SetCreatedAtUnixSecs(createdAtUnixSecs int) {
	k.CreatedAtUnixSecs = createdAtUnixSecs
}

func (k *KnowledgeBaseDocumentMetadataResponseModel) SetLastUpdatedAtUnixSecs(lastUpdatedAtUnixSecs int) {
	k.LastUpdatedAtUnixSecs = lastUpdatedAtUnixSecs
}

func (k *KnowledgeBaseDocumentMetadataResponseModel) SetSizeBytes(sizeBytes int) {
	k.SizeBytes = sizeBytes
}

func (k *KnowledgeBaseDocumentMetadataResponseModel) GetExtraProperties() map[string]interface{} {
	if k == nil {
		return nil
	}
	return k.extraProperties
}

func (k KnowledgeBaseDocumentMetadataResponseModel) String() string {
	jsonData, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return "error converting struct: KnowledgeBaseDocumentMetadataResponseModel to string"
	}
	return string(jsonData)
}

func (k *KnowledgeBaseDocumentMetadataResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, k); err != nil {
		return err
	}
	type alias KnowledgeBaseDocumentMetadataResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*k = KnowledgeBaseDocumentMetadataResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *k)
	if err != nil {
		return err
	}
	k.extraProperties = extraProperties
	return nil
}

func (k *KnowledgeBaseDocumentMetadataResponseModel) MarshalJSON() ([]byte, error) {
	type embed KnowledgeBaseDocumentMetadataResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*k),
	}
	return json.Marshal(marshaler)
}
