package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ListProceduresResponseModel struct {
	// Procedures on the branch with their draft-aware metadata.
	Procedures      []*ProcedureListItemResponseModel `json:"procedures" xml:"procedures" required:"true"`
	extraProperties map[string]interface{}
}

func (l *ListProceduresResponseModel) GetProcedures() (value []*ProcedureListItemResponseModel) {
	if l == nil {
		return
	}
	return l.Procedures
}

func (l *ListProceduresResponseModel) SetProcedures(procedures []*ProcedureListItemResponseModel) {
	l.Procedures = procedures
}

func (l *ListProceduresResponseModel) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListProceduresResponseModel) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListProceduresResponseModel to string"
	}
	return string(jsonData)
}

func (l *ListProceduresResponseModel) UnmarshalJSON(data []byte) error {
	type unmarshaler ListProceduresResponseModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListProceduresResponseModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListProceduresResponseModel) MarshalJSON() ([]byte, error) {
	type embed ListProceduresResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
