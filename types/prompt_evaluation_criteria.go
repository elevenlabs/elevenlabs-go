package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// An evaluation using the transcript and a prompt for a yes/no achieved answer
type PromptEvaluationCriteria struct {
	// The unique identifier for the evaluation criteria
	Id   string `json:"id" xml:"id" required:"true"`
	Name string `json:"name" xml:"name" required:"true"`
	// The type of evaluation criteria
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The prompt that the agent should use to evaluate the conversation
	ConversationGoalPrompt string `json:"conversation_goal_prompt" xml:"conversation_goal_prompt" required:"true" maxLength:"2000"`
	// When evaluating the prompt, should the agent's knowledge base be used.
	UseKnowledgeBase *bool          `json:"use_knowledge_base,omitempty" xml:"use_knowledge_base,omitempty"`
	Scope            *AnalysisScope `json:"scope,omitempty" xml:"scope,omitempty"`
	// LLM model to use for this evaluation criteria. If not set, uses agent's analysis_llm default.
	Llm         *Llm                 `json:"llm,omitempty" xml:"llm,omitempty"`
	ScoringMode *CriteriaScoringMode `json:"scoring_mode,omitempty" xml:"scoring_mode,omitempty"`
	// Maximum value of the numeric score scale (minimum is always 0). Only used when scoring_mode is 'numeric_uniform'.
	MaxScore *int `json:"max_score,omitempty" xml:"max_score,omitempty"`
	// Optional free-text instructions describing how to assign values on the numeric scale. Only used when scoring_mode is 'numeric_uniform'.
	ScoreInstructions *string `json:"score_instructions,omitempty" xml:"score_instructions,omitempty" maxLength:"2000"`
	extraProperties   map[string]interface{}
}

func (p *PromptEvaluationCriteria) GetId() (value string) {
	if p == nil {
		return
	}
	return p.Id
}

func (p *PromptEvaluationCriteria) GetName() (value string) {
	if p == nil {
		return
	}
	return p.Name
}

func (p *PromptEvaluationCriteria) GetType() (value *string) {
	if p == nil {
		return
	}
	return p.Type
}

func (p *PromptEvaluationCriteria) GetConversationGoalPrompt() (value string) {
	if p == nil {
		return
	}
	return p.ConversationGoalPrompt
}

func (p *PromptEvaluationCriteria) GetUseKnowledgeBase() (value *bool) {
	if p == nil {
		return
	}
	return p.UseKnowledgeBase
}

func (p *PromptEvaluationCriteria) GetScope() (value *AnalysisScope) {
	if p == nil {
		return
	}
	return p.Scope
}

func (p *PromptEvaluationCriteria) GetLlm() (value *Llm) {
	if p == nil {
		return
	}
	return p.Llm
}

func (p *PromptEvaluationCriteria) GetScoringMode() (value *CriteriaScoringMode) {
	if p == nil {
		return
	}
	return p.ScoringMode
}

func (p *PromptEvaluationCriteria) GetMaxScore() (value *int) {
	if p == nil {
		return
	}
	return p.MaxScore
}

func (p *PromptEvaluationCriteria) GetScoreInstructions() (value *string) {
	if p == nil {
		return
	}
	return p.ScoreInstructions
}

func (p *PromptEvaluationCriteria) SetId(id string) {
	p.Id = id
}

func (p *PromptEvaluationCriteria) SetName(name string) {
	p.Name = name
}

func (p *PromptEvaluationCriteria) SetType(type_ *string) {
	p.Type = type_
}

func (p *PromptEvaluationCriteria) SetConversationGoalPrompt(conversationGoalPrompt string) {
	p.ConversationGoalPrompt = conversationGoalPrompt
}

func (p *PromptEvaluationCriteria) SetUseKnowledgeBase(useKnowledgeBase *bool) {
	p.UseKnowledgeBase = useKnowledgeBase
}

func (p *PromptEvaluationCriteria) SetScope(scope *AnalysisScope) {
	p.Scope = scope
}

func (p *PromptEvaluationCriteria) SetLlm(llm *Llm) {
	p.Llm = llm
}

func (p *PromptEvaluationCriteria) SetScoringMode(scoringMode *CriteriaScoringMode) {
	p.ScoringMode = scoringMode
}

func (p *PromptEvaluationCriteria) SetMaxScore(maxScore *int) {
	p.MaxScore = maxScore
}

func (p *PromptEvaluationCriteria) SetScoreInstructions(scoreInstructions *string) {
	p.ScoreInstructions = scoreInstructions
}

func (p *PromptEvaluationCriteria) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PromptEvaluationCriteria) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PromptEvaluationCriteria to string"
	}
	return string(jsonData)
}

func (p *PromptEvaluationCriteria) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, p); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PromptEvaluationCriteria) MarshalJSON() ([]byte, error) {
	type embed PromptEvaluationCriteria
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
