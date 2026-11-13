package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type SystemToolConfigInputParams struct {
	SystemToolType      string
	EndCall             *EndCallToolConfig
	EndProcedure        *EndProcedureToolConfig
	KnowledgeBase       *KnowledgeBaseToolConfig
	KnowledgeBaseRag    *KnowledgeBaseRagToolConfig
	LanguageDetection   *LanguageDetectionToolConfig
	PlayKeypadTouchTone *PlayDtmfToolConfig
	SkipTurn            *SkipTurnToolConfig
	StartProcedure      *StartProcedureToolConfig
	TransferToAgent     *TransferToAgentToolConfigInput
	TransferToNumber    *TransferToNumberToolConfigInput
	VoicemailDetection  *VoicemailDetectionToolConfig
}

func (s *SystemToolConfigInputParams) GetEndCall() *EndCallToolConfig {
	if s == nil {
		return nil
	}
	if s.EndCall == nil {
		return nil
	}
	return s.EndCall
}

func (s *SystemToolConfigInputParams) GetEndProcedure() *EndProcedureToolConfig {
	if s == nil {
		return nil
	}
	if s.EndProcedure == nil {
		return nil
	}
	return s.EndProcedure
}

func (s *SystemToolConfigInputParams) GetKnowledgeBase() *KnowledgeBaseToolConfig {
	if s == nil {
		return nil
	}
	if s.KnowledgeBase == nil {
		return nil
	}
	return s.KnowledgeBase
}

func (s *SystemToolConfigInputParams) GetKnowledgeBaseRag() *KnowledgeBaseRagToolConfig {
	if s == nil {
		return nil
	}
	if s.KnowledgeBaseRag == nil {
		return nil
	}
	return s.KnowledgeBaseRag
}

func (s *SystemToolConfigInputParams) GetLanguageDetection() *LanguageDetectionToolConfig {
	if s == nil {
		return nil
	}
	if s.LanguageDetection == nil {
		return nil
	}
	return s.LanguageDetection
}

func (s *SystemToolConfigInputParams) GetPlayKeypadTouchTone() *PlayDtmfToolConfig {
	if s == nil {
		return nil
	}
	if s.PlayKeypadTouchTone == nil {
		return nil
	}
	return s.PlayKeypadTouchTone
}

func (s *SystemToolConfigInputParams) GetSkipTurn() *SkipTurnToolConfig {
	if s == nil {
		return nil
	}
	if s.SkipTurn == nil {
		return nil
	}
	return s.SkipTurn
}

func (s *SystemToolConfigInputParams) GetStartProcedure() *StartProcedureToolConfig {
	if s == nil {
		return nil
	}
	if s.StartProcedure == nil {
		return nil
	}
	return s.StartProcedure
}

func (s *SystemToolConfigInputParams) GetTransferToAgent() *TransferToAgentToolConfigInput {
	if s == nil {
		return nil
	}
	if s.TransferToAgent == nil {
		return nil
	}
	return s.TransferToAgent
}

func (s *SystemToolConfigInputParams) GetTransferToNumber() *TransferToNumberToolConfigInput {
	if s == nil {
		return nil
	}
	if s.TransferToNumber == nil {
		return nil
	}
	return s.TransferToNumber
}

func (s *SystemToolConfigInputParams) GetVoicemailDetection() *VoicemailDetectionToolConfig {
	if s == nil {
		return nil
	}
	if s.VoicemailDetection == nil {
		return nil
	}
	return s.VoicemailDetection
}

func (s *SystemToolConfigInputParams) GetSystemToolType() string {
	if s == nil {
		return ""
	}
	return s.SystemToolType
}

func (s *SystemToolConfigInputParams) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		SystemToolType string `json:"system_tool_type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	s.SystemToolType = unmarshaler.SystemToolType
	if unmarshaler.SystemToolType == "" {
		return fmt.Errorf("%T did not include discriminant system_tool_type", s)
	}
	switch unmarshaler.SystemToolType {
	case "end_call":
		value := new(EndCallToolConfig)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.EndCall = value
	case "end_procedure":
		value := new(EndProcedureToolConfig)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.EndProcedure = value
	case "knowledge_base":
		value := new(KnowledgeBaseToolConfig)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.KnowledgeBase = value
	case "knowledge_base_rag":
		value := new(KnowledgeBaseRagToolConfig)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.KnowledgeBaseRag = value
	case "language_detection":
		value := new(LanguageDetectionToolConfig)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.LanguageDetection = value
	case "play_keypad_touch_tone":
		value := new(PlayDtmfToolConfig)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.PlayKeypadTouchTone = value
	case "skip_turn":
		value := new(SkipTurnToolConfig)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.SkipTurn = value
	case "start_procedure":
		value := new(StartProcedureToolConfig)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.StartProcedure = value
	case "transfer_to_agent":
		value := new(TransferToAgentToolConfigInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.TransferToAgent = value
	case "transfer_to_number":
		value := new(TransferToNumberToolConfigInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.TransferToNumber = value
	case "voicemail_detection":
		value := new(VoicemailDetectionToolConfig)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.VoicemailDetection = value
	}
	return nil
}

func (s SystemToolConfigInputParams) MarshalJSON() ([]byte, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	if s.EndCall != nil {
		return internal.MarshalJSONWithDiscriminant(s.EndCall, "system_tool_type", "end_call")
	}
	if s.EndProcedure != nil {
		return internal.MarshalJSONWithDiscriminant(s.EndProcedure, "system_tool_type", "end_procedure")
	}
	if s.KnowledgeBase != nil {
		return internal.MarshalJSONWithDiscriminant(s.KnowledgeBase, "system_tool_type", "knowledge_base")
	}
	if s.KnowledgeBaseRag != nil {
		return internal.MarshalJSONWithDiscriminant(s.KnowledgeBaseRag, "system_tool_type", "knowledge_base_rag")
	}
	if s.LanguageDetection != nil {
		return internal.MarshalJSONWithDiscriminant(s.LanguageDetection, "system_tool_type", "language_detection")
	}
	if s.PlayKeypadTouchTone != nil {
		return internal.MarshalJSONWithDiscriminant(s.PlayKeypadTouchTone, "system_tool_type", "play_keypad_touch_tone")
	}
	if s.SkipTurn != nil {
		return internal.MarshalJSONWithDiscriminant(s.SkipTurn, "system_tool_type", "skip_turn")
	}
	if s.StartProcedure != nil {
		return internal.MarshalJSONWithDiscriminant(s.StartProcedure, "system_tool_type", "start_procedure")
	}
	if s.TransferToAgent != nil {
		return internal.MarshalJSONWithDiscriminant(s.TransferToAgent, "system_tool_type", "transfer_to_agent")
	}
	if s.TransferToNumber != nil {
		return internal.MarshalJSONWithDiscriminant(s.TransferToNumber, "system_tool_type", "transfer_to_number")
	}
	if s.VoicemailDetection != nil {
		return internal.MarshalJSONWithDiscriminant(s.VoicemailDetection, "system_tool_type", "voicemail_detection")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", s)
}

type SystemToolConfigInputParamsVisitor interface {
	VisitEndCall(*EndCallToolConfig) error
	VisitEndProcedure(*EndProcedureToolConfig) error
	VisitKnowledgeBase(*KnowledgeBaseToolConfig) error
	VisitKnowledgeBaseRag(*KnowledgeBaseRagToolConfig) error
	VisitLanguageDetection(*LanguageDetectionToolConfig) error
	VisitPlayKeypadTouchTone(*PlayDtmfToolConfig) error
	VisitSkipTurn(*SkipTurnToolConfig) error
	VisitStartProcedure(*StartProcedureToolConfig) error
	VisitTransferToAgent(*TransferToAgentToolConfigInput) error
	VisitTransferToNumber(*TransferToNumberToolConfigInput) error
	VisitVoicemailDetection(*VoicemailDetectionToolConfig) error
}

func (s *SystemToolConfigInputParams) Accept(visitor SystemToolConfigInputParamsVisitor) error {
	if s.EndCall != nil {
		return visitor.VisitEndCall(s.EndCall)
	}
	if s.EndProcedure != nil {
		return visitor.VisitEndProcedure(s.EndProcedure)
	}
	if s.KnowledgeBase != nil {
		return visitor.VisitKnowledgeBase(s.KnowledgeBase)
	}
	if s.KnowledgeBaseRag != nil {
		return visitor.VisitKnowledgeBaseRag(s.KnowledgeBaseRag)
	}
	if s.LanguageDetection != nil {
		return visitor.VisitLanguageDetection(s.LanguageDetection)
	}
	if s.PlayKeypadTouchTone != nil {
		return visitor.VisitPlayKeypadTouchTone(s.PlayKeypadTouchTone)
	}
	if s.SkipTurn != nil {
		return visitor.VisitSkipTurn(s.SkipTurn)
	}
	if s.StartProcedure != nil {
		return visitor.VisitStartProcedure(s.StartProcedure)
	}
	if s.TransferToAgent != nil {
		return visitor.VisitTransferToAgent(s.TransferToAgent)
	}
	if s.TransferToNumber != nil {
		return visitor.VisitTransferToNumber(s.TransferToNumber)
	}
	if s.VoicemailDetection != nil {
		return visitor.VisitVoicemailDetection(s.VoicemailDetection)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", s)
}

func (s *SystemToolConfigInputParams) validate() error {
	if s == nil {
		return fmt.Errorf("type %T is nil", s)
	}
	var fields []string
	if s.EndCall != nil {
		fields = append(fields, "end_call")
	}
	if s.EndProcedure != nil {
		fields = append(fields, "end_procedure")
	}
	if s.KnowledgeBase != nil {
		fields = append(fields, "knowledge_base")
	}
	if s.KnowledgeBaseRag != nil {
		fields = append(fields, "knowledge_base_rag")
	}
	if s.LanguageDetection != nil {
		fields = append(fields, "language_detection")
	}
	if s.PlayKeypadTouchTone != nil {
		fields = append(fields, "play_keypad_touch_tone")
	}
	if s.SkipTurn != nil {
		fields = append(fields, "skip_turn")
	}
	if s.StartProcedure != nil {
		fields = append(fields, "start_procedure")
	}
	if s.TransferToAgent != nil {
		fields = append(fields, "transfer_to_agent")
	}
	if s.TransferToNumber != nil {
		fields = append(fields, "transfer_to_number")
	}
	if s.VoicemailDetection != nil {
		fields = append(fields, "voicemail_detection")
	}
	if len(fields) == 0 {
		if s.SystemToolType != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", s, s.SystemToolType)
		}
		return fmt.Errorf("type %T is empty", s)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", s, fields)
	}
	if s.SystemToolType != "" {
		field := fields[0]
		if s.SystemToolType != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				s,
				s.SystemToolType,
				s,
			)
		}
	}
	return nil
}
