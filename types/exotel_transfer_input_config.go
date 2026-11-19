package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ExotelTransferInputConfig struct {
	Type                    string
	ExotelColdBlindTransfer *ExotelColdBlindTransferConfigInput
	ExotelWarmTransfer      *ExotelWarmTransferConfigInput
}

func (e *ExotelTransferInputConfig) GetExotelColdBlindTransfer() *ExotelColdBlindTransferConfigInput {
	if e == nil {
		return nil
	}
	if e.ExotelColdBlindTransfer == nil {
		return nil
	}
	return e.ExotelColdBlindTransfer
}

func (e *ExotelTransferInputConfig) GetExotelWarmTransfer() *ExotelWarmTransferConfigInput {
	if e == nil {
		return nil
	}
	if e.ExotelWarmTransfer == nil {
		return nil
	}
	return e.ExotelWarmTransfer
}

func (e *ExotelTransferInputConfig) GetType() string {
	if e == nil {
		return ""
	}
	return e.Type
}

func (e *ExotelTransferInputConfig) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	e.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", e)
	}
	switch unmarshaler.Type {
	case "exotel_cold_blind_transfer":
		value := new(ExotelColdBlindTransferConfigInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.ExotelColdBlindTransfer = value
	case "exotel_warm_transfer":
		value := new(ExotelWarmTransferConfigInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.ExotelWarmTransfer = value
	}
	return nil
}

func (e ExotelTransferInputConfig) MarshalJSON() ([]byte, error) {
	if err := e.validate(); err != nil {
		return nil, err
	}
	if e.ExotelColdBlindTransfer != nil {
		return internal.MarshalJSONWithDiscriminant(e.ExotelColdBlindTransfer, "type", "exotel_cold_blind_transfer")
	}
	if e.ExotelWarmTransfer != nil {
		return internal.MarshalJSONWithDiscriminant(e.ExotelWarmTransfer, "type", "exotel_warm_transfer")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", e)
}

type ExotelTransferInputConfigVisitor interface {
	VisitExotelColdBlindTransfer(*ExotelColdBlindTransferConfigInput) error
	VisitExotelWarmTransfer(*ExotelWarmTransferConfigInput) error
}

func (e *ExotelTransferInputConfig) Accept(visitor ExotelTransferInputConfigVisitor) error {
	if e.ExotelColdBlindTransfer != nil {
		return visitor.VisitExotelColdBlindTransfer(e.ExotelColdBlindTransfer)
	}
	if e.ExotelWarmTransfer != nil {
		return visitor.VisitExotelWarmTransfer(e.ExotelWarmTransfer)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", e)
}

func (e *ExotelTransferInputConfig) validate() error {
	if e == nil {
		return fmt.Errorf("type %T is nil", e)
	}
	var fields []string
	if e.ExotelColdBlindTransfer != nil {
		fields = append(fields, "exotel_cold_blind_transfer")
	}
	if e.ExotelWarmTransfer != nil {
		fields = append(fields, "exotel_warm_transfer")
	}
	if len(fields) == 0 {
		if e.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", e, e.Type)
		}
		return fmt.Errorf("type %T is empty", e)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", e, fields)
	}
	if e.Type != "" {
		field := fields[0]
		if e.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				e,
				e.Type,
				e,
			)
		}
	}
	return nil
}
