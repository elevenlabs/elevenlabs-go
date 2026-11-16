package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type GetConvAiDashboardSettingsResponseModelChartsItem struct {
	Type           string
	CallSuccess    *DashboardCallSuccessChartModel
	Criteria       *DashboardCriteriaChartModel
	DataCollection *DashboardDataCollectionChartModel
}

func (g *GetConvAiDashboardSettingsResponseModelChartsItem) GetCallSuccess() *DashboardCallSuccessChartModel {
	if g == nil {
		return nil
	}
	if g.CallSuccess == nil {
		return nil
	}
	return g.CallSuccess
}

func (g *GetConvAiDashboardSettingsResponseModelChartsItem) GetCriteria() *DashboardCriteriaChartModel {
	if g == nil {
		return nil
	}
	if g.Criteria == nil {
		return nil
	}
	return g.Criteria
}

func (g *GetConvAiDashboardSettingsResponseModelChartsItem) GetDataCollection() *DashboardDataCollectionChartModel {
	if g == nil {
		return nil
	}
	if g.DataCollection == nil {
		return nil
	}
	return g.DataCollection
}

func (g *GetConvAiDashboardSettingsResponseModelChartsItem) GetType() string {
	if g == nil {
		return ""
	}
	return g.Type
}

func (g *GetConvAiDashboardSettingsResponseModelChartsItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	g.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", g)
	}
	switch unmarshaler.Type {
	case "call_success":
		value := new(DashboardCallSuccessChartModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.CallSuccess = value
	case "criteria":
		value := new(DashboardCriteriaChartModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Criteria = value
	case "data_collection":
		value := new(DashboardDataCollectionChartModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.DataCollection = value
	}
	return nil
}

func (g GetConvAiDashboardSettingsResponseModelChartsItem) MarshalJSON() ([]byte, error) {
	if err := g.validate(); err != nil {
		return nil, err
	}
	if g.CallSuccess != nil {
		return internal.MarshalJSONWithDiscriminant(g.CallSuccess, "type", "call_success")
	}
	if g.Criteria != nil {
		return internal.MarshalJSONWithDiscriminant(g.Criteria, "type", "criteria")
	}
	if g.DataCollection != nil {
		return internal.MarshalJSONWithDiscriminant(g.DataCollection, "type", "data_collection")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", g)
}

type GetConvAiDashboardSettingsResponseModelChartsItemVisitor interface {
	VisitCallSuccess(*DashboardCallSuccessChartModel) error
	VisitCriteria(*DashboardCriteriaChartModel) error
	VisitDataCollection(*DashboardDataCollectionChartModel) error
}

func (g *GetConvAiDashboardSettingsResponseModelChartsItem) Accept(visitor GetConvAiDashboardSettingsResponseModelChartsItemVisitor) error {
	if g.CallSuccess != nil {
		return visitor.VisitCallSuccess(g.CallSuccess)
	}
	if g.Criteria != nil {
		return visitor.VisitCriteria(g.Criteria)
	}
	if g.DataCollection != nil {
		return visitor.VisitDataCollection(g.DataCollection)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", g)
}

func (g *GetConvAiDashboardSettingsResponseModelChartsItem) validate() error {
	if g == nil {
		return fmt.Errorf("type %T is nil", g)
	}
	var fields []string
	if g.CallSuccess != nil {
		fields = append(fields, "call_success")
	}
	if g.Criteria != nil {
		fields = append(fields, "criteria")
	}
	if g.DataCollection != nil {
		fields = append(fields, "data_collection")
	}
	if len(fields) == 0 {
		if g.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", g, g.Type)
		}
		return fmt.Errorf("type %T is empty", g)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", g, fields)
	}
	if g.Type != "" {
		field := fields[0]
		if g.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				g,
				g.Type,
				g,
			)
		}
	}
	return nil
}
