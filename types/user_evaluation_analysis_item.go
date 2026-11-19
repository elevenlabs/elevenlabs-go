package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UserEvaluationAnalysisItem struct {
	Name   string  `json:"name" xml:"name" required:"true"`
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// LLM to use for this item. If unset, the attaching agent's analysis_llm default is used.
	Llm  *Llm    `json:"llm,omitempty" xml:"llm,omitempty"`
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The prompt used to evaluate the conversation.
	ConversationGoalPrompt string `json:"conversation_goal_prompt" xml:"conversation_goal_prompt" required:"true" maxLength:"2000"`
	// Whether the attaching agent's knowledge base is used when evaluating this criterion.
	UseKnowledgeBase *bool `json:"use_knowledge_base,omitempty" xml:"use_knowledge_base,omitempty"`
	// How this criterion is scored (binary or numeric_uniform).
	Scoring         *UserEvaluationAnalysisItemScoring `json:"scoring,omitempty" xml:"scoring,omitempty"`
	extraProperties map[string]interface{}
}

func (u *UserEvaluationAnalysisItem) GetName() (value string) {
	if u == nil {
		return
	}
	return u.Name
}

func (u *UserEvaluationAnalysisItem) GetSource() (value *string) {
	if u == nil {
		return
	}
	return u.Source
}

func (u *UserEvaluationAnalysisItem) GetLlm() (value *Llm) {
	if u == nil {
		return
	}
	return u.Llm
}

func (u *UserEvaluationAnalysisItem) GetKind() (value *string) {
	if u == nil {
		return
	}
	return u.Kind
}

func (u *UserEvaluationAnalysisItem) GetConversationGoalPrompt() (value string) {
	if u == nil {
		return
	}
	return u.ConversationGoalPrompt
}

func (u *UserEvaluationAnalysisItem) GetUseKnowledgeBase() (value *bool) {
	if u == nil {
		return
	}
	return u.UseKnowledgeBase
}

func (u *UserEvaluationAnalysisItem) GetScoring() (value *UserEvaluationAnalysisItemScoring) {
	if u == nil {
		return
	}
	return u.Scoring
}

func (u *UserEvaluationAnalysisItem) SetName(name string) {
	u.Name = name
}

func (u *UserEvaluationAnalysisItem) SetSource(source *string) {
	u.Source = source
}

func (u *UserEvaluationAnalysisItem) SetLlm(llm *Llm) {
	u.Llm = llm
}

func (u *UserEvaluationAnalysisItem) SetKind(kind *string) {
	u.Kind = kind
}

func (u *UserEvaluationAnalysisItem) SetConversationGoalPrompt(conversationGoalPrompt string) {
	u.ConversationGoalPrompt = conversationGoalPrompt
}

func (u *UserEvaluationAnalysisItem) SetUseKnowledgeBase(useKnowledgeBase *bool) {
	u.UseKnowledgeBase = useKnowledgeBase
}

func (u *UserEvaluationAnalysisItem) SetScoring(scoring *UserEvaluationAnalysisItemScoring) {
	u.Scoring = scoring
}

func (u *UserEvaluationAnalysisItem) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UserEvaluationAnalysisItem) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UserEvaluationAnalysisItem to string"
	}
	return string(jsonData)
}

func (u *UserEvaluationAnalysisItem) UnmarshalJSON(data []byte) error {
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

func (u *UserEvaluationAnalysisItem) MarshalJSON() ([]byte, error) {
	type embed UserEvaluationAnalysisItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
