package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type KnowledgeBaseSummaryBatchSuccessfulResponseModel struct {
	status          string
	Data            *KnowledgeBaseSummaryBatchSuccessfulResponseModelData `json:"data" xml:"data" required:"true"`
	extraProperties map[string]interface{}
}

func (k *KnowledgeBaseSummaryBatchSuccessfulResponseModel) GetData() (value *KnowledgeBaseSummaryBatchSuccessfulResponseModelData) {
	if k == nil {
		return
	}
	return k.Data
}

func (k *KnowledgeBaseSummaryBatchSuccessfulResponseModel) SetData(data *KnowledgeBaseSummaryBatchSuccessfulResponseModelData) {
	k.Data = data
}

func (k *KnowledgeBaseSummaryBatchSuccessfulResponseModel) Status() string {
	return k.status
}

func (k *KnowledgeBaseSummaryBatchSuccessfulResponseModel) GetExtraProperties() map[string]interface{} {
	if k == nil {
		return nil
	}
	return k.extraProperties
}

func (k KnowledgeBaseSummaryBatchSuccessfulResponseModel) String() string {
	jsonData, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return "error converting struct: KnowledgeBaseSummaryBatchSuccessfulResponseModel to string"
	}
	return string(jsonData)
}

func (k *KnowledgeBaseSummaryBatchSuccessfulResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, k); err != nil {
		return err
	}
	type alias KnowledgeBaseSummaryBatchSuccessfulResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*k = KnowledgeBaseSummaryBatchSuccessfulResponseModel(tmp)
	var literals struct {
		Status *string `json:"status"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Status != nil && *literals.Status != "success" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", k, "success", *literals.Status)
	}
	k.status = "success"
	extraProperties, err := internal.ExtractExtraProperties(data, *k, "status")
	if err != nil {
		return err
	}
	k.extraProperties = extraProperties
	return nil
}

func (k *KnowledgeBaseSummaryBatchSuccessfulResponseModel) MarshalJSON() ([]byte, error) {
	type embed KnowledgeBaseSummaryBatchSuccessfulResponseModel
	var marshaler = struct {
		embed
		Status string `json:"status"`
	}{
		embed:  embed(*k),
		Status: "success",
	}
	return json.Marshal(marshaler)
}
