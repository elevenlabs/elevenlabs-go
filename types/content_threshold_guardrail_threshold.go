package types

import (
	"encoding/json"
	"fmt"
)

type ContentThresholdGuardrailThreshold struct {
	Double                                float64
	ContentThresholdGuardrailThresholdOne ContentThresholdGuardrailThresholdOne
	typ                                   string
}

func (c *ContentThresholdGuardrailThreshold) GetDouble() float64 {
	if c == nil {
		return 0
	}
	return c.Double
}

func (c *ContentThresholdGuardrailThreshold) GetContentThresholdGuardrailThresholdOne() ContentThresholdGuardrailThresholdOne {
	if c == nil {
		return ""
	}
	return c.ContentThresholdGuardrailThresholdOne
}

func (c *ContentThresholdGuardrailThreshold) UnmarshalJSON(data []byte) error {
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		c.typ = "Double"
		c.Double = valueDouble
		return nil
	}
	var valueContentThresholdGuardrailThresholdOne ContentThresholdGuardrailThresholdOne
	if err := json.Unmarshal(data, &valueContentThresholdGuardrailThresholdOne); err == nil {
		c.typ = "ContentThresholdGuardrailThresholdOne"
		c.ContentThresholdGuardrailThresholdOne = valueContentThresholdGuardrailThresholdOne
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c ContentThresholdGuardrailThreshold) MarshalJSON() ([]byte, error) {
	if c.typ == "Double" || c.Double != 0 {
		return json.Marshal(c.Double)
	}
	if c.typ == "ContentThresholdGuardrailThresholdOne" || c.ContentThresholdGuardrailThresholdOne != "" {
		return json.Marshal(c.ContentThresholdGuardrailThresholdOne)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", c)
}

type ContentThresholdGuardrailThresholdVisitor interface {
	VisitDouble(float64) error
	VisitContentThresholdGuardrailThresholdOne(ContentThresholdGuardrailThresholdOne) error
}

func (c *ContentThresholdGuardrailThreshold) Accept(visitor ContentThresholdGuardrailThresholdVisitor) error {
	if c.typ == "Double" || c.Double != 0 {
		return visitor.VisitDouble(c.Double)
	}
	if c.typ == "ContentThresholdGuardrailThresholdOne" || c.ContentThresholdGuardrailThresholdOne != "" {
		return visitor.VisitContentThresholdGuardrailThresholdOne(c.ContentThresholdGuardrailThresholdOne)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", c)
}
