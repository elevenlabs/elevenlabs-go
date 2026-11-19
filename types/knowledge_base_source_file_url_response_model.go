package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type KnowledgeBaseSourceFileUrlResponseModel struct {
	// Signed URL to download the source file directly
	SignedUrl       string `json:"signed_url" xml:"signed_url" required:"true"`
	extraProperties map[string]interface{}
}

func (k *KnowledgeBaseSourceFileUrlResponseModel) GetSignedUrl() (value string) {
	if k == nil {
		return
	}
	return k.SignedUrl
}

func (k *KnowledgeBaseSourceFileUrlResponseModel) SetSignedUrl(signedUrl string) {
	k.SignedUrl = signedUrl
}

func (k *KnowledgeBaseSourceFileUrlResponseModel) GetExtraProperties() map[string]interface{} {
	if k == nil {
		return nil
	}
	return k.extraProperties
}

func (k KnowledgeBaseSourceFileUrlResponseModel) String() string {
	jsonData, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return "error converting struct: KnowledgeBaseSourceFileUrlResponseModel to string"
	}
	return string(jsonData)
}

func (k *KnowledgeBaseSourceFileUrlResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, k); err != nil {
		return err
	}
	type alias KnowledgeBaseSourceFileUrlResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*k = KnowledgeBaseSourceFileUrlResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *k)
	if err != nil {
		return err
	}
	k.extraProperties = extraProperties
	return nil
}

func (k *KnowledgeBaseSourceFileUrlResponseModel) MarshalJSON() ([]byte, error) {
	type embed KnowledgeBaseSourceFileUrlResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*k),
	}
	return json.Marshal(marshaler)
}
