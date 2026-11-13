package agents

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type BodyCreateOrUpdateDeploymentsV1ConvaiAgentsAgentIdDeploymentsPost struct {
	// The id of an agent. This is returned on agent creation.
	AgentId           string                        `json:"-" url:"-"`
	DeploymentRequest *types.AgentDeploymentRequest `json:"deployment_request" xml:"deployment_request" required:"true"`
}

func (b *BodyCreateOrUpdateDeploymentsV1ConvaiAgentsAgentIdDeploymentsPost) GetAgentId() (value string) {
	if b == nil {
		return
	}
	return b.AgentId
}

func (b *BodyCreateOrUpdateDeploymentsV1ConvaiAgentsAgentIdDeploymentsPost) GetDeploymentRequest() (value *types.AgentDeploymentRequest) {
	if b == nil {
		return
	}
	return b.DeploymentRequest
}

func (b *BodyCreateOrUpdateDeploymentsV1ConvaiAgentsAgentIdDeploymentsPost) SetAgentId(agentId string) {
	b.AgentId = agentId
}

func (b *BodyCreateOrUpdateDeploymentsV1ConvaiAgentsAgentIdDeploymentsPost) SetDeploymentRequest(deploymentRequest *types.AgentDeploymentRequest) {
	b.DeploymentRequest = deploymentRequest
}

func (b *BodyCreateOrUpdateDeploymentsV1ConvaiAgentsAgentIdDeploymentsPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreateOrUpdateDeploymentsV1ConvaiAgentsAgentIdDeploymentsPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreateOrUpdateDeploymentsV1ConvaiAgentsAgentIdDeploymentsPost to string"
	}
	return string(jsonData)
}

func (b *BodyCreateOrUpdateDeploymentsV1ConvaiAgentsAgentIdDeploymentsPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyCreateOrUpdateDeploymentsV1ConvaiAgentsAgentIdDeploymentsPost
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyCreateOrUpdateDeploymentsV1ConvaiAgentsAgentIdDeploymentsPost(tmp)
	return nil
}

func (b *BodyCreateOrUpdateDeploymentsV1ConvaiAgentsAgentIdDeploymentsPost) MarshalJSON() ([]byte, error) {
	type embed BodyCreateOrUpdateDeploymentsV1ConvaiAgentsAgentIdDeploymentsPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

var (
	listDeploymentsRequestFieldPage     = big.NewInt(1 << 0)
	listDeploymentsRequestFieldPageSize = big.NewInt(1 << 1)
)

type ListDeploymentsRequest struct {
	Page     *int `json:"-" url:"page,omitempty"`
	PageSize *int `json:"-" url:"page_size,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListDeploymentsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetPage sets the Page field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListDeploymentsRequest) SetPage(page *int) {
	l.Page = page
	l.require(listDeploymentsRequestFieldPage)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListDeploymentsRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listDeploymentsRequestFieldPageSize)
}
