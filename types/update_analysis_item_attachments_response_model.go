package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type UpdateAnalysisItemAttachmentsResponseModel struct {
	Results         []*UpdateAnalysisItemAttachmentsResponseModelResultsItem `json:"results" xml:"results" required:"true"`
	extraProperties map[string]interface{}
}

func (u *UpdateAnalysisItemAttachmentsResponseModel) GetResults() (value []*UpdateAnalysisItemAttachmentsResponseModelResultsItem) {
	if u == nil {
		return
	}
	return u.Results
}

func (u *UpdateAnalysisItemAttachmentsResponseModel) SetResults(results []*UpdateAnalysisItemAttachmentsResponseModelResultsItem) {
	u.Results = results
}

func (u *UpdateAnalysisItemAttachmentsResponseModel) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateAnalysisItemAttachmentsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateAnalysisItemAttachmentsResponseModel to string"
	}
	return string(jsonData)
}

func (u *UpdateAnalysisItemAttachmentsResponseModel) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateAnalysisItemAttachmentsResponseModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateAnalysisItemAttachmentsResponseModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateAnalysisItemAttachmentsResponseModel) MarshalJSON() ([]byte, error) {
	type embed UpdateAnalysisItemAttachmentsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
