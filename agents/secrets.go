package agents

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	big "math/big"
)

type PostWorkspaceSecretRequest struct {
	type_ string
	Name  string `json:"name" xml:"name" required:"true"`
	Value string `json:"value" xml:"value" required:"true"`
}

func (p *PostWorkspaceSecretRequest) GetName() (value string) {
	if p == nil {
		return
	}
	return p.Name
}

func (p *PostWorkspaceSecretRequest) GetValue() (value string) {
	if p == nil {
		return
	}
	return p.Value
}

func (p *PostWorkspaceSecretRequest) SetName(name string) {
	p.Name = name
}

func (p *PostWorkspaceSecretRequest) SetValue(value string) {
	p.Value = value
}

func (p *PostWorkspaceSecretRequest) Type() string {
	return p.type_
}

func (p *PostWorkspaceSecretRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (p PostWorkspaceSecretRequest) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PostWorkspaceSecretRequest to string"
	}
	return string(jsonData)
}

func (p *PostWorkspaceSecretRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PostWorkspaceSecretRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PostWorkspaceSecretRequest(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "new" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", p, "new", *literals.Type)
	}
	p.type_ = "new"
	return nil
}

func (p *PostWorkspaceSecretRequest) MarshalJSON() ([]byte, error) {
	type embed PostWorkspaceSecretRequest
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*p),
		Type:  "new",
	}
	return json.Marshal(marshaler)
}

type PostWorkspaceSecretRequestType = string

type PatchWorkspaceSecretRequest struct {
	SecretId string `json:"-" url:"-"`
	type_    string
	Name     string `json:"name" xml:"name" required:"true"`
	Value    string `json:"value" xml:"value" required:"true"`
}

func (p *PatchWorkspaceSecretRequest) GetSecretId() (value string) {
	if p == nil {
		return
	}
	return p.SecretId
}

func (p *PatchWorkspaceSecretRequest) GetName() (value string) {
	if p == nil {
		return
	}
	return p.Name
}

func (p *PatchWorkspaceSecretRequest) GetValue() (value string) {
	if p == nil {
		return
	}
	return p.Value
}

func (p *PatchWorkspaceSecretRequest) SetSecretId(secretId string) {
	p.SecretId = secretId
}

func (p *PatchWorkspaceSecretRequest) SetName(name string) {
	p.Name = name
}

func (p *PatchWorkspaceSecretRequest) SetValue(value string) {
	p.Value = value
}

func (p *PatchWorkspaceSecretRequest) Type() string {
	return p.type_
}

func (p *PatchWorkspaceSecretRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (p PatchWorkspaceSecretRequest) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PatchWorkspaceSecretRequest to string"
	}
	return string(jsonData)
}

func (p *PatchWorkspaceSecretRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PatchWorkspaceSecretRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PatchWorkspaceSecretRequest(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "update" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", p, "update", *literals.Type)
	}
	p.type_ = "update"
	return nil
}

func (p *PatchWorkspaceSecretRequest) MarshalJSON() ([]byte, error) {
	type embed PatchWorkspaceSecretRequest
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*p),
		Type:  "update",
	}
	return json.Marshal(marshaler)
}

type PatchWorkspaceSecretRequestType = string

var (
	listSecretsRequestFieldPageSize        = big.NewInt(1 << 0)
	listSecretsRequestFieldDependencyLimit = big.NewInt(1 << 1)
	listSecretsRequestFieldSearch          = big.NewInt(1 << 2)
	listSecretsRequestFieldCursor          = big.NewInt(1 << 3)
)

type ListSecretsRequest struct {
	PageSize        *int    `json:"-" url:"page_size,omitempty"`
	DependencyLimit *int    `json:"-" url:"dependency_limit,omitempty"`
	Search          *string `json:"-" url:"search,omitempty"`
	Cursor          *string `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListSecretsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSecretsRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listSecretsRequestFieldPageSize)
}

// SetDependencyLimit sets the DependencyLimit field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSecretsRequest) SetDependencyLimit(dependencyLimit *int) {
	l.DependencyLimit = dependencyLimit
	l.require(listSecretsRequestFieldDependencyLimit)
}

// SetSearch sets the Search field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSecretsRequest) SetSearch(search *string) {
	l.Search = search
	l.require(listSecretsRequestFieldSearch)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSecretsRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listSecretsRequestFieldCursor)
}

var (
	getDependenciesSecretsRequestFieldPageSize = big.NewInt(1 << 0)
	getDependenciesSecretsRequestFieldCursor   = big.NewInt(1 << 1)
)

type GetDependenciesSecretsRequest struct {
	PageSize *int    `json:"-" url:"page_size,omitempty"`
	Cursor   *string `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetDependenciesSecretsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetDependenciesSecretsRequest) SetPageSize(pageSize *int) {
	g.PageSize = pageSize
	g.require(getDependenciesSecretsRequestFieldPageSize)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetDependenciesSecretsRequest) SetCursor(cursor *string) {
	g.Cursor = cursor
	g.require(getDependenciesSecretsRequestFieldCursor)
}
