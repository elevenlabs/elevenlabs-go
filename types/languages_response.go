package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type LanguagesResponse struct {
	Kind   string
	Pair   *PairedLanguagesResponse
	Single *SingleLanguagesResponse
}

func (l *LanguagesResponse) GetPair() *PairedLanguagesResponse {
	if l == nil {
		return nil
	}
	if l.Pair == nil {
		return nil
	}
	return l.Pair
}

func (l *LanguagesResponse) GetSingle() *SingleLanguagesResponse {
	if l == nil {
		return nil
	}
	if l.Single == nil {
		return nil
	}
	return l.Single
}

func (l *LanguagesResponse) GetKind() string {
	if l == nil {
		return ""
	}
	return l.Kind
}

func (l *LanguagesResponse) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Kind string `json:"kind"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	l.Kind = unmarshaler.Kind
	if unmarshaler.Kind == "" {
		return fmt.Errorf("%T did not include discriminant kind", l)
	}
	switch unmarshaler.Kind {
	case "pair":
		value := new(PairedLanguagesResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		l.Pair = value
	case "single":
		value := new(SingleLanguagesResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		l.Single = value
	}
	return nil
}

func (l LanguagesResponse) MarshalJSON() ([]byte, error) {
	if err := l.validate(); err != nil {
		return nil, err
	}
	if l.Pair != nil {
		return internal.MarshalJSONWithDiscriminant(l.Pair, "kind", "pair")
	}
	if l.Single != nil {
		return internal.MarshalJSONWithDiscriminant(l.Single, "kind", "single")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", l)
}

type LanguagesResponseVisitor interface {
	VisitPair(*PairedLanguagesResponse) error
	VisitSingle(*SingleLanguagesResponse) error
}

func (l *LanguagesResponse) Accept(visitor LanguagesResponseVisitor) error {
	if l.Pair != nil {
		return visitor.VisitPair(l.Pair)
	}
	if l.Single != nil {
		return visitor.VisitSingle(l.Single)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", l)
}

func (l *LanguagesResponse) validate() error {
	if l == nil {
		return fmt.Errorf("type %T is nil", l)
	}
	var fields []string
	if l.Pair != nil {
		fields = append(fields, "pair")
	}
	if l.Single != nil {
		fields = append(fields, "single")
	}
	if len(fields) == 0 {
		if l.Kind != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", l, l.Kind)
		}
		return fmt.Errorf("type %T is empty", l)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", l, fields)
	}
	if l.Kind != "" {
		field := fields[0]
		if l.Kind != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				l,
				l.Kind,
				l,
			)
		}
	}
	return nil
}
