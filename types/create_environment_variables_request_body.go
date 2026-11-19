package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type CreateEnvironmentVariablesRequestBody struct {
	Type           string
	AuthConnection *CreateAuthConnectionEnvironmentVariableRequest
	Secret         *CreateSecretEnvironmentVariableRequest
	FieldString    *CreateStringEnvironmentVariableRequest
}

func (c *CreateEnvironmentVariablesRequestBody) GetAuthConnection() *CreateAuthConnectionEnvironmentVariableRequest {
	if c == nil {
		return nil
	}
	if c.AuthConnection == nil {
		return nil
	}
	return c.AuthConnection
}

func (c *CreateEnvironmentVariablesRequestBody) GetSecret() *CreateSecretEnvironmentVariableRequest {
	if c == nil {
		return nil
	}
	if c.Secret == nil {
		return nil
	}
	return c.Secret
}

func (c *CreateEnvironmentVariablesRequestBody) GetFieldString() *CreateStringEnvironmentVariableRequest {
	if c == nil {
		return nil
	}
	if c.FieldString == nil {
		return nil
	}
	return c.FieldString
}

func (c *CreateEnvironmentVariablesRequestBody) GetType() string {
	if c == nil {
		return ""
	}
	return c.Type
}

func (c *CreateEnvironmentVariablesRequestBody) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	c.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", c)
	}
	switch unmarshaler.Type {
	case "auth_connection":
		value := new(CreateAuthConnectionEnvironmentVariableRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.AuthConnection = value
	case "secret":
		value := new(CreateSecretEnvironmentVariableRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Secret = value
	case "string":
		value := new(CreateStringEnvironmentVariableRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.FieldString = value
	}
	return nil
}

func (c CreateEnvironmentVariablesRequestBody) MarshalJSON() ([]byte, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}
	if c.AuthConnection != nil {
		return internal.MarshalJSONWithDiscriminant(c.AuthConnection, "type", "auth_connection")
	}
	if c.Secret != nil {
		return internal.MarshalJSONWithDiscriminant(c.Secret, "type", "secret")
	}
	if c.FieldString != nil {
		return internal.MarshalJSONWithDiscriminant(c.FieldString, "type", "string")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", c)
}

type CreateEnvironmentVariablesRequestBodyVisitor interface {
	VisitAuthConnection(*CreateAuthConnectionEnvironmentVariableRequest) error
	VisitSecret(*CreateSecretEnvironmentVariableRequest) error
	VisitFieldString(*CreateStringEnvironmentVariableRequest) error
}

func (c *CreateEnvironmentVariablesRequestBody) Accept(visitor CreateEnvironmentVariablesRequestBodyVisitor) error {
	if c.AuthConnection != nil {
		return visitor.VisitAuthConnection(c.AuthConnection)
	}
	if c.Secret != nil {
		return visitor.VisitSecret(c.Secret)
	}
	if c.FieldString != nil {
		return visitor.VisitFieldString(c.FieldString)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", c)
}

func (c *CreateEnvironmentVariablesRequestBody) validate() error {
	if c == nil {
		return fmt.Errorf("type %T is nil", c)
	}
	var fields []string
	if c.AuthConnection != nil {
		fields = append(fields, "auth_connection")
	}
	if c.Secret != nil {
		fields = append(fields, "secret")
	}
	if c.FieldString != nil {
		fields = append(fields, "string")
	}
	if len(fields) == 0 {
		if c.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", c, c.Type)
		}
		return fmt.Errorf("type %T is empty", c)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", c, fields)
	}
	if c.Type != "" {
		field := fields[0]
		if c.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				c,
				c.Type,
				c,
			)
		}
	}
	return nil
}
