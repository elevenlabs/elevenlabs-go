package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentSummaryBatchSuccessfulResponseModel struct {
	status          string
	Data            *AgentSummaryResponseModel `json:"data" xml:"data" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AgentSummaryBatchSuccessfulResponseModel) GetData() (value *AgentSummaryResponseModel) {
	if a == nil {
		return
	}
	return a.Data
}

func (a *AgentSummaryBatchSuccessfulResponseModel) SetData(data *AgentSummaryResponseModel) {
	a.Data = data
}

func (a *AgentSummaryBatchSuccessfulResponseModel) Status() string {
	return a.status
}

func (a *AgentSummaryBatchSuccessfulResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentSummaryBatchSuccessfulResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentSummaryBatchSuccessfulResponseModel to string"
	}
	return string(jsonData)
}

func (a *AgentSummaryBatchSuccessfulResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentSummaryBatchSuccessfulResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentSummaryBatchSuccessfulResponseModel(tmp)
	var literals struct {
		Status *string `json:"status"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Status != nil && *literals.Status != "success" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", a, "success", *literals.Status)
	}
	a.status = "success"
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "status")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentSummaryBatchSuccessfulResponseModel) MarshalJSON() ([]byte, error) {
	type embed AgentSummaryBatchSuccessfulResponseModel
	var marshaler = struct {
		embed
		Status string `json:"status"`
	}{
		embed:  embed(*a),
		Status: "success",
	}
	return json.Marshal(marshaler)
}
