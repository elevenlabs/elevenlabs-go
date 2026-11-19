package dashboard

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

type PatchConvAiDashboardSettingsRequest struct {
	Charts []*PatchConvAiDashboardSettingsRequestChartsItem `json:"charts,omitempty" xml:"charts,omitempty" maxItems:"4"`
}

func (p *PatchConvAiDashboardSettingsRequest) GetCharts() (value []*PatchConvAiDashboardSettingsRequestChartsItem) {
	if p == nil {
		return
	}
	return p.Charts
}

func (p *PatchConvAiDashboardSettingsRequest) SetCharts(charts []*PatchConvAiDashboardSettingsRequestChartsItem) {
	p.Charts = charts
}

func (p *PatchConvAiDashboardSettingsRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (p PatchConvAiDashboardSettingsRequest) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PatchConvAiDashboardSettingsRequest to string"
	}
	return string(jsonData)
}

func (p *PatchConvAiDashboardSettingsRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler PatchConvAiDashboardSettingsRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PatchConvAiDashboardSettingsRequest(value)
	return nil
}

func (p *PatchConvAiDashboardSettingsRequest) MarshalJSON() ([]byte, error) {
	type embed PatchConvAiDashboardSettingsRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}

type PatchConvAiDashboardSettingsRequestChartsItem struct {
	Type           string
	CallSuccess    *types.DashboardCallSuccessChartModel
	Criteria       *types.DashboardCriteriaChartModel
	DataCollection *types.DashboardDataCollectionChartModel
}

func (p *PatchConvAiDashboardSettingsRequestChartsItem) GetCallSuccess() *types.DashboardCallSuccessChartModel {
	if p == nil {
		return nil
	}
	if p.CallSuccess == nil {
		return nil
	}
	return p.CallSuccess
}

func (p *PatchConvAiDashboardSettingsRequestChartsItem) GetCriteria() *types.DashboardCriteriaChartModel {
	if p == nil {
		return nil
	}
	if p.Criteria == nil {
		return nil
	}
	return p.Criteria
}

func (p *PatchConvAiDashboardSettingsRequestChartsItem) GetDataCollection() *types.DashboardDataCollectionChartModel {
	if p == nil {
		return nil
	}
	if p.DataCollection == nil {
		return nil
	}
	return p.DataCollection
}

func (p *PatchConvAiDashboardSettingsRequestChartsItem) GetType() string {
	if p == nil {
		return ""
	}
	return p.Type
}

func (p *PatchConvAiDashboardSettingsRequestChartsItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	p.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", p)
	}
	switch unmarshaler.Type {
	case "call_success":
		value := new(types.DashboardCallSuccessChartModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.CallSuccess = value
	case "criteria":
		value := new(types.DashboardCriteriaChartModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Criteria = value
	case "data_collection":
		value := new(types.DashboardDataCollectionChartModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.DataCollection = value
	}
	return nil
}

func (p PatchConvAiDashboardSettingsRequestChartsItem) MarshalJSON() ([]byte, error) {
	if err := p.validate(); err != nil {
		return nil, err
	}
	if p.CallSuccess != nil {
		return internal.MarshalJSONWithDiscriminant(p.CallSuccess, "type", "call_success")
	}
	if p.Criteria != nil {
		return internal.MarshalJSONWithDiscriminant(p.Criteria, "type", "criteria")
	}
	if p.DataCollection != nil {
		return internal.MarshalJSONWithDiscriminant(p.DataCollection, "type", "data_collection")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", p)
}

type PatchConvAiDashboardSettingsRequestChartsItemVisitor interface {
	VisitCallSuccess(*types.DashboardCallSuccessChartModel) error
	VisitCriteria(*types.DashboardCriteriaChartModel) error
	VisitDataCollection(*types.DashboardDataCollectionChartModel) error
}

func (p *PatchConvAiDashboardSettingsRequestChartsItem) Accept(visitor PatchConvAiDashboardSettingsRequestChartsItemVisitor) error {
	if p.CallSuccess != nil {
		return visitor.VisitCallSuccess(p.CallSuccess)
	}
	if p.Criteria != nil {
		return visitor.VisitCriteria(p.Criteria)
	}
	if p.DataCollection != nil {
		return visitor.VisitDataCollection(p.DataCollection)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", p)
}

func (p *PatchConvAiDashboardSettingsRequestChartsItem) validate() error {
	if p == nil {
		return fmt.Errorf("type %T is nil", p)
	}
	var fields []string
	if p.CallSuccess != nil {
		fields = append(fields, "call_success")
	}
	if p.Criteria != nil {
		fields = append(fields, "criteria")
	}
	if p.DataCollection != nil {
		fields = append(fields, "data_collection")
	}
	if len(fields) == 0 {
		if p.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", p, p.Type)
		}
		return fmt.Errorf("type %T is empty", p)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", p, fields)
	}
	if p.Type != "" {
		field := fields[0]
		if p.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				p,
				p.Type,
				p,
			)
		}
	}
	return nil
}
