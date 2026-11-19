package dubbing

import (
	"encoding/json"
	"fmt"
)

// Dubbing model (`dubbing_v1` or `dubbing_v2`) every language target of this project is dubbed with. Defaults to `dubbing_v2`. Fixed at create time — the source is prepared for this model, so neither the project nor an individual target can change it later.
type BodyCreateDubbingProjectV1DubbingProjectPostModelId struct {
	CreateProjectRequestModelIdZero CreateProjectRequestModelIdZero
	String                          string
	typ                             string
}

func (bodyCreateDubbingProjectV1DubbingProjectPostModelId *BodyCreateDubbingProjectV1DubbingProjectPostModelId) GetCreateProjectRequestModelIdZero() CreateProjectRequestModelIdZero {
	if bodyCreateDubbingProjectV1DubbingProjectPostModelId == nil {
		return ""
	}
	return bodyCreateDubbingProjectV1DubbingProjectPostModelId.CreateProjectRequestModelIdZero
}

func (bodyCreateDubbingProjectV1DubbingProjectPostModelId *BodyCreateDubbingProjectV1DubbingProjectPostModelId) GetString() string {
	if bodyCreateDubbingProjectV1DubbingProjectPostModelId == nil {
		return ""
	}
	return bodyCreateDubbingProjectV1DubbingProjectPostModelId.String
}

func (bodyCreateDubbingProjectV1DubbingProjectPostModelId *BodyCreateDubbingProjectV1DubbingProjectPostModelId) UnmarshalJSON(data []byte) error {
	var valueCreateProjectRequestModelIdZero CreateProjectRequestModelIdZero
	if err := json.Unmarshal(data, &valueCreateProjectRequestModelIdZero); err == nil {
		bodyCreateDubbingProjectV1DubbingProjectPostModelId.typ = "CreateProjectRequestModelIdZero"
		bodyCreateDubbingProjectV1DubbingProjectPostModelId.CreateProjectRequestModelIdZero = valueCreateProjectRequestModelIdZero
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		bodyCreateDubbingProjectV1DubbingProjectPostModelId.typ = "String"
		bodyCreateDubbingProjectV1DubbingProjectPostModelId.String = valueString
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, bodyCreateDubbingProjectV1DubbingProjectPostModelId)
}

func (bodyCreateDubbingProjectV1DubbingProjectPostModelId BodyCreateDubbingProjectV1DubbingProjectPostModelId) MarshalJSON() ([]byte, error) {
	if bodyCreateDubbingProjectV1DubbingProjectPostModelId.typ == "CreateProjectRequestModelIdZero" || bodyCreateDubbingProjectV1DubbingProjectPostModelId.CreateProjectRequestModelIdZero != "" {
		return json.Marshal(bodyCreateDubbingProjectV1DubbingProjectPostModelId.CreateProjectRequestModelIdZero)
	}
	if bodyCreateDubbingProjectV1DubbingProjectPostModelId.typ == "String" || bodyCreateDubbingProjectV1DubbingProjectPostModelId.String != "" {
		return json.Marshal(bodyCreateDubbingProjectV1DubbingProjectPostModelId.String)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", bodyCreateDubbingProjectV1DubbingProjectPostModelId)
}

type BodyCreateDubbingProjectV1DubbingProjectPostModelIdVisitor interface {
	VisitCreateProjectRequestModelIdZero(CreateProjectRequestModelIdZero) error
	VisitString(string) error
}

func (bodyCreateDubbingProjectV1DubbingProjectPostModelId *BodyCreateDubbingProjectV1DubbingProjectPostModelId) Accept(visitor BodyCreateDubbingProjectV1DubbingProjectPostModelIdVisitor) error {
	if bodyCreateDubbingProjectV1DubbingProjectPostModelId.typ == "CreateProjectRequestModelIdZero" || bodyCreateDubbingProjectV1DubbingProjectPostModelId.CreateProjectRequestModelIdZero != "" {
		return visitor.VisitCreateProjectRequestModelIdZero(bodyCreateDubbingProjectV1DubbingProjectPostModelId.CreateProjectRequestModelIdZero)
	}
	if bodyCreateDubbingProjectV1DubbingProjectPostModelId.typ == "String" || bodyCreateDubbingProjectV1DubbingProjectPostModelId.String != "" {
		return visitor.VisitString(bodyCreateDubbingProjectV1DubbingProjectPostModelId.String)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", bodyCreateDubbingProjectV1DubbingProjectPostModelId)
}
