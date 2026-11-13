package types

import (
	"encoding/json"
	"fmt"
)

type UpdateEnvironmentVariableRequestValuesValue struct {
	String                                        string
	EnvironmentVariableSecretValueRequest         *EnvironmentVariableSecretValueRequest
	EnvironmentVariableAuthConnectionValueRequest *EnvironmentVariableAuthConnectionValueRequest
	typ                                           string
}

func (u *UpdateEnvironmentVariableRequestValuesValue) GetString() string {
	if u == nil {
		return ""
	}
	return u.String
}

func (u *UpdateEnvironmentVariableRequestValuesValue) GetEnvironmentVariableSecretValueRequest() *EnvironmentVariableSecretValueRequest {
	if u == nil {
		return nil
	}
	if u.EnvironmentVariableSecretValueRequest == nil {
		return nil
	}
	return u.EnvironmentVariableSecretValueRequest
}

func (u *UpdateEnvironmentVariableRequestValuesValue) GetEnvironmentVariableAuthConnectionValueRequest() *EnvironmentVariableAuthConnectionValueRequest {
	if u == nil {
		return nil
	}
	if u.EnvironmentVariableAuthConnectionValueRequest == nil {
		return nil
	}
	return u.EnvironmentVariableAuthConnectionValueRequest
}

func (u *UpdateEnvironmentVariableRequestValuesValue) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		u.typ = "String"
		u.String = valueString
		return nil
	}
	var valueEnvironmentVariableSecretValueRequest *EnvironmentVariableSecretValueRequest
	if err := json.Unmarshal(data, &valueEnvironmentVariableSecretValueRequest); err == nil {
		u.typ = "EnvironmentVariableSecretValueRequest"
		u.EnvironmentVariableSecretValueRequest = valueEnvironmentVariableSecretValueRequest
		return nil
	}
	var valueEnvironmentVariableAuthConnectionValueRequest *EnvironmentVariableAuthConnectionValueRequest
	if err := json.Unmarshal(data, &valueEnvironmentVariableAuthConnectionValueRequest); err == nil {
		u.typ = "EnvironmentVariableAuthConnectionValueRequest"
		u.EnvironmentVariableAuthConnectionValueRequest = valueEnvironmentVariableAuthConnectionValueRequest
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, u)
}

func (u UpdateEnvironmentVariableRequestValuesValue) MarshalJSON() ([]byte, error) {
	if u.typ == "String" || u.String != "" {
		return json.Marshal(u.String)
	}
	if u.typ == "EnvironmentVariableSecretValueRequest" || u.EnvironmentVariableSecretValueRequest != nil {
		return json.Marshal(u.EnvironmentVariableSecretValueRequest)
	}
	if u.typ == "EnvironmentVariableAuthConnectionValueRequest" || u.EnvironmentVariableAuthConnectionValueRequest != nil {
		return json.Marshal(u.EnvironmentVariableAuthConnectionValueRequest)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", u)
}

type UpdateEnvironmentVariableRequestValuesValueVisitor interface {
	VisitString(string) error
	VisitEnvironmentVariableSecretValueRequest(*EnvironmentVariableSecretValueRequest) error
	VisitEnvironmentVariableAuthConnectionValueRequest(*EnvironmentVariableAuthConnectionValueRequest) error
}

func (u *UpdateEnvironmentVariableRequestValuesValue) Accept(visitor UpdateEnvironmentVariableRequestValuesValueVisitor) error {
	if u.typ == "String" || u.String != "" {
		return visitor.VisitString(u.String)
	}
	if u.typ == "EnvironmentVariableSecretValueRequest" || u.EnvironmentVariableSecretValueRequest != nil {
		return visitor.VisitEnvironmentVariableSecretValueRequest(u.EnvironmentVariableSecretValueRequest)
	}
	if u.typ == "EnvironmentVariableAuthConnectionValueRequest" || u.EnvironmentVariableAuthConnectionValueRequest != nil {
		return visitor.VisitEnvironmentVariableAuthConnectionValueRequest(u.EnvironmentVariableAuthConnectionValueRequest)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", u)
}
