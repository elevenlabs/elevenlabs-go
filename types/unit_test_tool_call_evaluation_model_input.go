package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UnitTestToolCallEvaluationModelInput struct {
	// Parameters to evaluate for the agent's tool call. If empty, the tool call parameters are not evaluated.
	Parameters []*UnitTestToolCallParameter `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// The tool to evaluate a call against.
	ReferencedTool *ReferencedToolCommonModel `json:"referenced_tool,omitempty" xml:"referenced_tool,omitempty"`
	// Whether to verify that the tool was NOT called.
	VerifyAbsence *bool `json:"verify_absence,omitempty" xml:"verify_absence,omitempty"`
	// Configuration for testing workflow node transitions. When set, the test will verify the agent transitions to the specified workflow node.
	WorkflowNodeTransition *UnitTestWorkflowNodeTransitionEvaluationNodeId `json:"workflow_node_transition,omitempty" xml:"workflow_node_transition,omitempty"`
	// Evaluate guardrails triggered by the generated agent turn.
	GuardrailTrigger *UnitTestGuardrailTriggerEvaluationModel `json:"guardrail_trigger,omitempty" xml:"guardrail_trigger,omitempty"`
	extraProperties  map[string]interface{}
}

func (u *UnitTestToolCallEvaluationModelInput) GetParameters() (value []*UnitTestToolCallParameter) {
	if u == nil {
		return
	}
	return u.Parameters
}

func (u *UnitTestToolCallEvaluationModelInput) GetReferencedTool() (value *ReferencedToolCommonModel) {
	if u == nil {
		return
	}
	return u.ReferencedTool
}

func (u *UnitTestToolCallEvaluationModelInput) GetVerifyAbsence() (value *bool) {
	if u == nil {
		return
	}
	return u.VerifyAbsence
}

func (u *UnitTestToolCallEvaluationModelInput) GetWorkflowNodeTransition() (value *UnitTestWorkflowNodeTransitionEvaluationNodeId) {
	if u == nil {
		return
	}
	return u.WorkflowNodeTransition
}

func (u *UnitTestToolCallEvaluationModelInput) GetGuardrailTrigger() (value *UnitTestGuardrailTriggerEvaluationModel) {
	if u == nil {
		return
	}
	return u.GuardrailTrigger
}

func (u *UnitTestToolCallEvaluationModelInput) SetParameters(parameters []*UnitTestToolCallParameter) {
	u.Parameters = parameters
}

func (u *UnitTestToolCallEvaluationModelInput) SetReferencedTool(referencedTool *ReferencedToolCommonModel) {
	u.ReferencedTool = referencedTool
}

func (u *UnitTestToolCallEvaluationModelInput) SetVerifyAbsence(verifyAbsence *bool) {
	u.VerifyAbsence = verifyAbsence
}

func (u *UnitTestToolCallEvaluationModelInput) SetWorkflowNodeTransition(workflowNodeTransition *UnitTestWorkflowNodeTransitionEvaluationNodeId) {
	u.WorkflowNodeTransition = workflowNodeTransition
}

func (u *UnitTestToolCallEvaluationModelInput) SetGuardrailTrigger(guardrailTrigger *UnitTestGuardrailTriggerEvaluationModel) {
	u.GuardrailTrigger = guardrailTrigger
}

func (u *UnitTestToolCallEvaluationModelInput) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UnitTestToolCallEvaluationModelInput) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UnitTestToolCallEvaluationModelInput to string"
	}
	return string(jsonData)
}

func (u *UnitTestToolCallEvaluationModelInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, u); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UnitTestToolCallEvaluationModelInput) MarshalJSON() ([]byte, error) {
	type embed UnitTestToolCallEvaluationModelInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
