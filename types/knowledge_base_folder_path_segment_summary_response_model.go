package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type KnowledgeBaseFolderPathSegmentSummaryResponseModel struct {
	Id              string `json:"id" xml:"id" required:"true"`
	extraProperties map[string]interface{}
}

func (k *KnowledgeBaseFolderPathSegmentSummaryResponseModel) GetId() (value string) {
	if k == nil {
		return
	}
	return k.Id
}

func (k *KnowledgeBaseFolderPathSegmentSummaryResponseModel) SetId(id string) {
	k.Id = id
}

func (k *KnowledgeBaseFolderPathSegmentSummaryResponseModel) GetExtraProperties() map[string]interface{} {
	if k == nil {
		return nil
	}
	return k.extraProperties
}

func (k KnowledgeBaseFolderPathSegmentSummaryResponseModel) String() string {
	jsonData, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return "error converting struct: KnowledgeBaseFolderPathSegmentSummaryResponseModel to string"
	}
	return string(jsonData)
}

func (k *KnowledgeBaseFolderPathSegmentSummaryResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, k); err != nil {
		return err
	}
	type alias KnowledgeBaseFolderPathSegmentSummaryResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*k = KnowledgeBaseFolderPathSegmentSummaryResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *k)
	if err != nil {
		return err
	}
	k.extraProperties = extraProperties
	return nil
}

func (k *KnowledgeBaseFolderPathSegmentSummaryResponseModel) MarshalJSON() ([]byte, error) {
	type embed KnowledgeBaseFolderPathSegmentSummaryResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*k),
	}
	return json.Marshal(marshaler)
}
