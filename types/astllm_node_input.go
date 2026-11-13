package types

import (
	"encoding/json"
	"fmt"
)

type AstllmNodeInput struct {
	AstllmNodeInputZero *AstllmNodeInputZero
	AstllmNodeInputOne  *AstllmNodeInputOne
	typ                 string
}

func (a *AstllmNodeInput) GetAstllmNodeInputZero() *AstllmNodeInputZero {
	if a == nil {
		return nil
	}
	if a.AstllmNodeInputZero == nil {
		return nil
	}
	return a.AstllmNodeInputZero
}

func (a *AstllmNodeInput) GetAstllmNodeInputOne() *AstllmNodeInputOne {
	if a == nil {
		return nil
	}
	if a.AstllmNodeInputOne == nil {
		return nil
	}
	return a.AstllmNodeInputOne
}

func (a *AstllmNodeInput) UnmarshalJSON(data []byte) error {
	var unionKeys map[string]json.RawMessage
	_ = json.Unmarshal(data, &unionKeys)
	hasLiteral := func(key string, want interface{}) bool {
		raw, ok := unionKeys[key]
		if !ok {
			return false
		}
		var got interface{}
		return json.Unmarshal(raw, &got) == nil && got == want
	}
	matchedLiteral := false
	if hasLiteral("type", "llm") {
		matchedLiteral = true
		var valueAstllmNodeInputZeroLiteral *AstllmNodeInputZero
		if err := json.Unmarshal(data, &valueAstllmNodeInputZeroLiteral); err == nil {
			a.typ = "AstllmNodeInputZero"
			a.AstllmNodeInputZero = valueAstllmNodeInputZeroLiteral
			return nil
		}
	}
	if hasLiteral("type", "llm") {
		matchedLiteral = true
		var valueAstllmNodeInputOneLiteral *AstllmNodeInputOne
		if err := json.Unmarshal(data, &valueAstllmNodeInputOneLiteral); err == nil {
			a.typ = "AstllmNodeInputOne"
			a.AstllmNodeInputOne = valueAstllmNodeInputOneLiteral
			return nil
		}
	}
	if matchedLiteral {
		return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
	}
	var valueAstllmNodeInputZero *AstllmNodeInputZero
	if err := json.Unmarshal(data, &valueAstllmNodeInputZero); err == nil {
		a.typ = "AstllmNodeInputZero"
		a.AstllmNodeInputZero = valueAstllmNodeInputZero
		return nil
	}
	var valueAstllmNodeInputOne *AstllmNodeInputOne
	if err := json.Unmarshal(data, &valueAstllmNodeInputOne); err == nil {
		a.typ = "AstllmNodeInputOne"
		a.AstllmNodeInputOne = valueAstllmNodeInputOne
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a AstllmNodeInput) MarshalJSON() ([]byte, error) {
	if a.typ == "AstllmNodeInputZero" || a.AstllmNodeInputZero != nil {
		return json.Marshal(a.AstllmNodeInputZero)
	}
	if a.typ == "AstllmNodeInputOne" || a.AstllmNodeInputOne != nil {
		return json.Marshal(a.AstllmNodeInputOne)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type AstllmNodeInputVisitor interface {
	VisitAstllmNodeInputZero(*AstllmNodeInputZero) error
	VisitAstllmNodeInputOne(*AstllmNodeInputOne) error
}

func (a *AstllmNodeInput) Accept(visitor AstllmNodeInputVisitor) error {
	if a.typ == "AstllmNodeInputZero" || a.AstllmNodeInputZero != nil {
		return visitor.VisitAstllmNodeInputZero(a.AstllmNodeInputZero)
	}
	if a.typ == "AstllmNodeInputOne" || a.AstllmNodeInputOne != nil {
		return visitor.VisitAstllmNodeInputOne(a.AstllmNodeInputOne)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}
