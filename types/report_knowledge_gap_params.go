package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ReportKnowledgeGapParams struct {
	extraProperties map[string]interface{}
}

func (r *ReportKnowledgeGapParams) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReportKnowledgeGapParams) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReportKnowledgeGapParams to string"
	}
	return string(jsonData)
}

func (r *ReportKnowledgeGapParams) UnmarshalJSON(data []byte) error {
	type unmarshaler ReportKnowledgeGapParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ReportKnowledgeGapParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r, "smb_tool_type")
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ReportKnowledgeGapParams) MarshalJSON() ([]byte, error) {
	type embed ReportKnowledgeGapParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
