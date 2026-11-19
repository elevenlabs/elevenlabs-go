package types

import (
	"encoding/json"
	"fmt"
)

type GetSecretDependenciesResponseModelDependencies struct {
	GetSecretDependenciesResponseModelDependenciesZeroItemList []*GetSecretDependenciesResponseModelDependenciesZeroItem
	GetSecretDependenciesResponseModelDependenciesOneItemList  []*GetSecretDependenciesResponseModelDependenciesOneItem
	DependentPhoneNumberIdentifierList                         []*DependentPhoneNumberIdentifier
	typ                                                        string
}

func (g *GetSecretDependenciesResponseModelDependencies) GetGetSecretDependenciesResponseModelDependenciesZeroItemList() []*GetSecretDependenciesResponseModelDependenciesZeroItem {
	if g == nil {
		return nil
	}
	return g.GetSecretDependenciesResponseModelDependenciesZeroItemList
}

func (g *GetSecretDependenciesResponseModelDependencies) GetGetSecretDependenciesResponseModelDependenciesOneItemList() []*GetSecretDependenciesResponseModelDependenciesOneItem {
	if g == nil {
		return nil
	}
	return g.GetSecretDependenciesResponseModelDependenciesOneItemList
}

func (g *GetSecretDependenciesResponseModelDependencies) GetDependentPhoneNumberIdentifierList() []*DependentPhoneNumberIdentifier {
	if g == nil {
		return nil
	}
	return g.DependentPhoneNumberIdentifierList
}

func (g *GetSecretDependenciesResponseModelDependencies) UnmarshalJSON(data []byte) error {
	var valueGetSecretDependenciesResponseModelDependenciesZeroItemList []*GetSecretDependenciesResponseModelDependenciesZeroItem
	if err := json.Unmarshal(data, &valueGetSecretDependenciesResponseModelDependenciesZeroItemList); err == nil {
		g.typ = "GetSecretDependenciesResponseModelDependenciesZeroItemList"
		g.GetSecretDependenciesResponseModelDependenciesZeroItemList = valueGetSecretDependenciesResponseModelDependenciesZeroItemList
		return nil
	}
	var valueGetSecretDependenciesResponseModelDependenciesOneItemList []*GetSecretDependenciesResponseModelDependenciesOneItem
	if err := json.Unmarshal(data, &valueGetSecretDependenciesResponseModelDependenciesOneItemList); err == nil {
		g.typ = "GetSecretDependenciesResponseModelDependenciesOneItemList"
		g.GetSecretDependenciesResponseModelDependenciesOneItemList = valueGetSecretDependenciesResponseModelDependenciesOneItemList
		return nil
	}
	var valueDependentPhoneNumberIdentifierList []*DependentPhoneNumberIdentifier
	if err := json.Unmarshal(data, &valueDependentPhoneNumberIdentifierList); err == nil {
		g.typ = "DependentPhoneNumberIdentifierList"
		g.DependentPhoneNumberIdentifierList = valueDependentPhoneNumberIdentifierList
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, g)
}

func (g GetSecretDependenciesResponseModelDependencies) MarshalJSON() ([]byte, error) {
	if g.typ == "GetSecretDependenciesResponseModelDependenciesZeroItemList" || g.GetSecretDependenciesResponseModelDependenciesZeroItemList != nil {
		return json.Marshal(g.GetSecretDependenciesResponseModelDependenciesZeroItemList)
	}
	if g.typ == "GetSecretDependenciesResponseModelDependenciesOneItemList" || g.GetSecretDependenciesResponseModelDependenciesOneItemList != nil {
		return json.Marshal(g.GetSecretDependenciesResponseModelDependenciesOneItemList)
	}
	if g.typ == "DependentPhoneNumberIdentifierList" || g.DependentPhoneNumberIdentifierList != nil {
		return json.Marshal(g.DependentPhoneNumberIdentifierList)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", g)
}

type GetSecretDependenciesResponseModelDependenciesVisitor interface {
	VisitGetSecretDependenciesResponseModelDependenciesZeroItemList([]*GetSecretDependenciesResponseModelDependenciesZeroItem) error
	VisitGetSecretDependenciesResponseModelDependenciesOneItemList([]*GetSecretDependenciesResponseModelDependenciesOneItem) error
	VisitDependentPhoneNumberIdentifierList([]*DependentPhoneNumberIdentifier) error
}

func (g *GetSecretDependenciesResponseModelDependencies) Accept(visitor GetSecretDependenciesResponseModelDependenciesVisitor) error {
	if g.typ == "GetSecretDependenciesResponseModelDependenciesZeroItemList" || g.GetSecretDependenciesResponseModelDependenciesZeroItemList != nil {
		return visitor.VisitGetSecretDependenciesResponseModelDependenciesZeroItemList(g.GetSecretDependenciesResponseModelDependenciesZeroItemList)
	}
	if g.typ == "GetSecretDependenciesResponseModelDependenciesOneItemList" || g.GetSecretDependenciesResponseModelDependenciesOneItemList != nil {
		return visitor.VisitGetSecretDependenciesResponseModelDependenciesOneItemList(g.GetSecretDependenciesResponseModelDependenciesOneItemList)
	}
	if g.typ == "DependentPhoneNumberIdentifierList" || g.DependentPhoneNumberIdentifierList != nil {
		return visitor.VisitDependentPhoneNumberIdentifierList(g.DependentPhoneNumberIdentifierList)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", g)
}
