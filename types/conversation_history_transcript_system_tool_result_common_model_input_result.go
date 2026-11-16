package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ConversationHistoryTranscriptSystemToolResultCommonModelInputResult struct {
	ResultType                           string
	Dummy                                *DummyToolResultModel
	EndCallSuccess                       *EndCallToolResultModel
	EndProcedureError                    *EndProcedureToolResultErrorModel
	EndProcedureSuccess                  *EndProcedureToolResultSuccessModel
	KnowledgeBaseRagSuccess              *KnowledgeBaseRagToolResultModel
	KnowledgeBaseSuccess                 *KnowledgeBaseToolResultModel
	LanguageDetectionSuccess             *LanguageDetectionToolResultModel
	PlayDtmfError                        *PlayDtmfResultErrorModel
	PlayDtmfSuccess                      *PlayDtmfResultSuccessModel
	SkipTurnSuccess                      *SkipTurnToolResponseModel
	StartProcedureError                  *StartProcedureToolResultErrorModel
	StartProcedureSuccess                *StartProcedureToolResultSuccessModel
	TestingToolResult                    *TestToolResultModel
	TransferToAgentError                 *TransferToAgentToolResultErrorModel
	TransferToAgentSuccess               *TransferToAgentToolResultSuccessModelInput
	TransferToNumberAmazonConnectSuccess *TransferToNumberResultAmazonConnectSuccessModel
	TransferToNumberError                *TransferToNumberResultErrorModel
	TransferToNumberExotelSuccess        *TransferToNumberResultExotelSuccessModel
	TransferToNumberSipSuccess           *TransferToNumberResultSipSuccessModel
	TransferToNumberTwilioSuccess        *TransferToNumberResultTwilioSuccessModel
	VoicemailDetectionSuccess            *VoiceMailDetectionResultSuccessModel
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) GetDummy() *DummyToolResultModel {
	if c == nil {
		return nil
	}
	if c.Dummy == nil {
		return nil
	}
	return c.Dummy
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) GetEndCallSuccess() *EndCallToolResultModel {
	if c == nil {
		return nil
	}
	if c.EndCallSuccess == nil {
		return nil
	}
	return c.EndCallSuccess
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) GetEndProcedureError() *EndProcedureToolResultErrorModel {
	if c == nil {
		return nil
	}
	if c.EndProcedureError == nil {
		return nil
	}
	return c.EndProcedureError
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) GetEndProcedureSuccess() *EndProcedureToolResultSuccessModel {
	if c == nil {
		return nil
	}
	if c.EndProcedureSuccess == nil {
		return nil
	}
	return c.EndProcedureSuccess
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) GetKnowledgeBaseRagSuccess() *KnowledgeBaseRagToolResultModel {
	if c == nil {
		return nil
	}
	if c.KnowledgeBaseRagSuccess == nil {
		return nil
	}
	return c.KnowledgeBaseRagSuccess
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) GetKnowledgeBaseSuccess() *KnowledgeBaseToolResultModel {
	if c == nil {
		return nil
	}
	if c.KnowledgeBaseSuccess == nil {
		return nil
	}
	return c.KnowledgeBaseSuccess
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) GetLanguageDetectionSuccess() *LanguageDetectionToolResultModel {
	if c == nil {
		return nil
	}
	if c.LanguageDetectionSuccess == nil {
		return nil
	}
	return c.LanguageDetectionSuccess
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) GetPlayDtmfError() *PlayDtmfResultErrorModel {
	if c == nil {
		return nil
	}
	if c.PlayDtmfError == nil {
		return nil
	}
	return c.PlayDtmfError
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) GetPlayDtmfSuccess() *PlayDtmfResultSuccessModel {
	if c == nil {
		return nil
	}
	if c.PlayDtmfSuccess == nil {
		return nil
	}
	return c.PlayDtmfSuccess
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) GetSkipTurnSuccess() *SkipTurnToolResponseModel {
	if c == nil {
		return nil
	}
	if c.SkipTurnSuccess == nil {
		return nil
	}
	return c.SkipTurnSuccess
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) GetStartProcedureError() *StartProcedureToolResultErrorModel {
	if c == nil {
		return nil
	}
	if c.StartProcedureError == nil {
		return nil
	}
	return c.StartProcedureError
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) GetStartProcedureSuccess() *StartProcedureToolResultSuccessModel {
	if c == nil {
		return nil
	}
	if c.StartProcedureSuccess == nil {
		return nil
	}
	return c.StartProcedureSuccess
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) GetTestingToolResult() *TestToolResultModel {
	if c == nil {
		return nil
	}
	if c.TestingToolResult == nil {
		return nil
	}
	return c.TestingToolResult
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) GetTransferToAgentError() *TransferToAgentToolResultErrorModel {
	if c == nil {
		return nil
	}
	if c.TransferToAgentError == nil {
		return nil
	}
	return c.TransferToAgentError
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) GetTransferToAgentSuccess() *TransferToAgentToolResultSuccessModelInput {
	if c == nil {
		return nil
	}
	if c.TransferToAgentSuccess == nil {
		return nil
	}
	return c.TransferToAgentSuccess
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) GetTransferToNumberAmazonConnectSuccess() *TransferToNumberResultAmazonConnectSuccessModel {
	if c == nil {
		return nil
	}
	if c.TransferToNumberAmazonConnectSuccess == nil {
		return nil
	}
	return c.TransferToNumberAmazonConnectSuccess
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) GetTransferToNumberError() *TransferToNumberResultErrorModel {
	if c == nil {
		return nil
	}
	if c.TransferToNumberError == nil {
		return nil
	}
	return c.TransferToNumberError
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) GetTransferToNumberExotelSuccess() *TransferToNumberResultExotelSuccessModel {
	if c == nil {
		return nil
	}
	if c.TransferToNumberExotelSuccess == nil {
		return nil
	}
	return c.TransferToNumberExotelSuccess
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) GetTransferToNumberSipSuccess() *TransferToNumberResultSipSuccessModel {
	if c == nil {
		return nil
	}
	if c.TransferToNumberSipSuccess == nil {
		return nil
	}
	return c.TransferToNumberSipSuccess
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) GetTransferToNumberTwilioSuccess() *TransferToNumberResultTwilioSuccessModel {
	if c == nil {
		return nil
	}
	if c.TransferToNumberTwilioSuccess == nil {
		return nil
	}
	return c.TransferToNumberTwilioSuccess
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) GetVoicemailDetectionSuccess() *VoiceMailDetectionResultSuccessModel {
	if c == nil {
		return nil
	}
	if c.VoicemailDetectionSuccess == nil {
		return nil
	}
	return c.VoicemailDetectionSuccess
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) GetResultType() string {
	if c == nil {
		return ""
	}
	return c.ResultType
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		ResultType string `json:"result_type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	c.ResultType = unmarshaler.ResultType
	if unmarshaler.ResultType == "" {
		return fmt.Errorf("%T did not include discriminant result_type", c)
	}
	switch unmarshaler.ResultType {
	case "dummy":
		value := new(DummyToolResultModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Dummy = value
	case "end_call_success":
		value := new(EndCallToolResultModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.EndCallSuccess = value
	case "end_procedure_error":
		value := new(EndProcedureToolResultErrorModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.EndProcedureError = value
	case "end_procedure_success":
		value := new(EndProcedureToolResultSuccessModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.EndProcedureSuccess = value
	case "knowledge_base_rag_success":
		value := new(KnowledgeBaseRagToolResultModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.KnowledgeBaseRagSuccess = value
	case "knowledge_base_success":
		value := new(KnowledgeBaseToolResultModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.KnowledgeBaseSuccess = value
	case "language_detection_success":
		value := new(LanguageDetectionToolResultModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.LanguageDetectionSuccess = value
	case "play_dtmf_error":
		value := new(PlayDtmfResultErrorModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.PlayDtmfError = value
	case "play_dtmf_success":
		value := new(PlayDtmfResultSuccessModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.PlayDtmfSuccess = value
	case "skip_turn_success":
		value := new(SkipTurnToolResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.SkipTurnSuccess = value
	case "start_procedure_error":
		value := new(StartProcedureToolResultErrorModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.StartProcedureError = value
	case "start_procedure_success":
		value := new(StartProcedureToolResultSuccessModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.StartProcedureSuccess = value
	case "testing_tool_result":
		value := new(TestToolResultModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.TestingToolResult = value
	case "transfer_to_agent_error":
		value := new(TransferToAgentToolResultErrorModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.TransferToAgentError = value
	case "transfer_to_agent_success":
		value := new(TransferToAgentToolResultSuccessModelInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.TransferToAgentSuccess = value
	case "transfer_to_number_amazon_connect_success":
		value := new(TransferToNumberResultAmazonConnectSuccessModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.TransferToNumberAmazonConnectSuccess = value
	case "transfer_to_number_error":
		value := new(TransferToNumberResultErrorModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.TransferToNumberError = value
	case "transfer_to_number_exotel_success":
		value := new(TransferToNumberResultExotelSuccessModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.TransferToNumberExotelSuccess = value
	case "transfer_to_number_sip_success":
		value := new(TransferToNumberResultSipSuccessModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.TransferToNumberSipSuccess = value
	case "transfer_to_number_twilio_success":
		value := new(TransferToNumberResultTwilioSuccessModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.TransferToNumberTwilioSuccess = value
	case "voicemail_detection_success":
		value := new(VoiceMailDetectionResultSuccessModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.VoicemailDetectionSuccess = value
	}
	return nil
}

func (c ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) MarshalJSON() ([]byte, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}
	if c.Dummy != nil {
		return internal.MarshalJSONWithDiscriminant(c.Dummy, "result_type", "dummy")
	}
	if c.EndCallSuccess != nil {
		return internal.MarshalJSONWithDiscriminant(c.EndCallSuccess, "result_type", "end_call_success")
	}
	if c.EndProcedureError != nil {
		return internal.MarshalJSONWithDiscriminant(c.EndProcedureError, "result_type", "end_procedure_error")
	}
	if c.EndProcedureSuccess != nil {
		return internal.MarshalJSONWithDiscriminant(c.EndProcedureSuccess, "result_type", "end_procedure_success")
	}
	if c.KnowledgeBaseRagSuccess != nil {
		return internal.MarshalJSONWithDiscriminant(c.KnowledgeBaseRagSuccess, "result_type", "knowledge_base_rag_success")
	}
	if c.KnowledgeBaseSuccess != nil {
		return internal.MarshalJSONWithDiscriminant(c.KnowledgeBaseSuccess, "result_type", "knowledge_base_success")
	}
	if c.LanguageDetectionSuccess != nil {
		return internal.MarshalJSONWithDiscriminant(c.LanguageDetectionSuccess, "result_type", "language_detection_success")
	}
	if c.PlayDtmfError != nil {
		return internal.MarshalJSONWithDiscriminant(c.PlayDtmfError, "result_type", "play_dtmf_error")
	}
	if c.PlayDtmfSuccess != nil {
		return internal.MarshalJSONWithDiscriminant(c.PlayDtmfSuccess, "result_type", "play_dtmf_success")
	}
	if c.SkipTurnSuccess != nil {
		return internal.MarshalJSONWithDiscriminant(c.SkipTurnSuccess, "result_type", "skip_turn_success")
	}
	if c.StartProcedureError != nil {
		return internal.MarshalJSONWithDiscriminant(c.StartProcedureError, "result_type", "start_procedure_error")
	}
	if c.StartProcedureSuccess != nil {
		return internal.MarshalJSONWithDiscriminant(c.StartProcedureSuccess, "result_type", "start_procedure_success")
	}
	if c.TestingToolResult != nil {
		return internal.MarshalJSONWithDiscriminant(c.TestingToolResult, "result_type", "testing_tool_result")
	}
	if c.TransferToAgentError != nil {
		return internal.MarshalJSONWithDiscriminant(c.TransferToAgentError, "result_type", "transfer_to_agent_error")
	}
	if c.TransferToAgentSuccess != nil {
		return internal.MarshalJSONWithDiscriminant(c.TransferToAgentSuccess, "result_type", "transfer_to_agent_success")
	}
	if c.TransferToNumberAmazonConnectSuccess != nil {
		return internal.MarshalJSONWithDiscriminant(c.TransferToNumberAmazonConnectSuccess, "result_type", "transfer_to_number_amazon_connect_success")
	}
	if c.TransferToNumberError != nil {
		return internal.MarshalJSONWithDiscriminant(c.TransferToNumberError, "result_type", "transfer_to_number_error")
	}
	if c.TransferToNumberExotelSuccess != nil {
		return internal.MarshalJSONWithDiscriminant(c.TransferToNumberExotelSuccess, "result_type", "transfer_to_number_exotel_success")
	}
	if c.TransferToNumberSipSuccess != nil {
		return internal.MarshalJSONWithDiscriminant(c.TransferToNumberSipSuccess, "result_type", "transfer_to_number_sip_success")
	}
	if c.TransferToNumberTwilioSuccess != nil {
		return internal.MarshalJSONWithDiscriminant(c.TransferToNumberTwilioSuccess, "result_type", "transfer_to_number_twilio_success")
	}
	if c.VoicemailDetectionSuccess != nil {
		return internal.MarshalJSONWithDiscriminant(c.VoicemailDetectionSuccess, "result_type", "voicemail_detection_success")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", c)
}

type ConversationHistoryTranscriptSystemToolResultCommonModelInputResultVisitor interface {
	VisitDummy(*DummyToolResultModel) error
	VisitEndCallSuccess(*EndCallToolResultModel) error
	VisitEndProcedureError(*EndProcedureToolResultErrorModel) error
	VisitEndProcedureSuccess(*EndProcedureToolResultSuccessModel) error
	VisitKnowledgeBaseRagSuccess(*KnowledgeBaseRagToolResultModel) error
	VisitKnowledgeBaseSuccess(*KnowledgeBaseToolResultModel) error
	VisitLanguageDetectionSuccess(*LanguageDetectionToolResultModel) error
	VisitPlayDtmfError(*PlayDtmfResultErrorModel) error
	VisitPlayDtmfSuccess(*PlayDtmfResultSuccessModel) error
	VisitSkipTurnSuccess(*SkipTurnToolResponseModel) error
	VisitStartProcedureError(*StartProcedureToolResultErrorModel) error
	VisitStartProcedureSuccess(*StartProcedureToolResultSuccessModel) error
	VisitTestingToolResult(*TestToolResultModel) error
	VisitTransferToAgentError(*TransferToAgentToolResultErrorModel) error
	VisitTransferToAgentSuccess(*TransferToAgentToolResultSuccessModelInput) error
	VisitTransferToNumberAmazonConnectSuccess(*TransferToNumberResultAmazonConnectSuccessModel) error
	VisitTransferToNumberError(*TransferToNumberResultErrorModel) error
	VisitTransferToNumberExotelSuccess(*TransferToNumberResultExotelSuccessModel) error
	VisitTransferToNumberSipSuccess(*TransferToNumberResultSipSuccessModel) error
	VisitTransferToNumberTwilioSuccess(*TransferToNumberResultTwilioSuccessModel) error
	VisitVoicemailDetectionSuccess(*VoiceMailDetectionResultSuccessModel) error
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) Accept(visitor ConversationHistoryTranscriptSystemToolResultCommonModelInputResultVisitor) error {
	if c.Dummy != nil {
		return visitor.VisitDummy(c.Dummy)
	}
	if c.EndCallSuccess != nil {
		return visitor.VisitEndCallSuccess(c.EndCallSuccess)
	}
	if c.EndProcedureError != nil {
		return visitor.VisitEndProcedureError(c.EndProcedureError)
	}
	if c.EndProcedureSuccess != nil {
		return visitor.VisitEndProcedureSuccess(c.EndProcedureSuccess)
	}
	if c.KnowledgeBaseRagSuccess != nil {
		return visitor.VisitKnowledgeBaseRagSuccess(c.KnowledgeBaseRagSuccess)
	}
	if c.KnowledgeBaseSuccess != nil {
		return visitor.VisitKnowledgeBaseSuccess(c.KnowledgeBaseSuccess)
	}
	if c.LanguageDetectionSuccess != nil {
		return visitor.VisitLanguageDetectionSuccess(c.LanguageDetectionSuccess)
	}
	if c.PlayDtmfError != nil {
		return visitor.VisitPlayDtmfError(c.PlayDtmfError)
	}
	if c.PlayDtmfSuccess != nil {
		return visitor.VisitPlayDtmfSuccess(c.PlayDtmfSuccess)
	}
	if c.SkipTurnSuccess != nil {
		return visitor.VisitSkipTurnSuccess(c.SkipTurnSuccess)
	}
	if c.StartProcedureError != nil {
		return visitor.VisitStartProcedureError(c.StartProcedureError)
	}
	if c.StartProcedureSuccess != nil {
		return visitor.VisitStartProcedureSuccess(c.StartProcedureSuccess)
	}
	if c.TestingToolResult != nil {
		return visitor.VisitTestingToolResult(c.TestingToolResult)
	}
	if c.TransferToAgentError != nil {
		return visitor.VisitTransferToAgentError(c.TransferToAgentError)
	}
	if c.TransferToAgentSuccess != nil {
		return visitor.VisitTransferToAgentSuccess(c.TransferToAgentSuccess)
	}
	if c.TransferToNumberAmazonConnectSuccess != nil {
		return visitor.VisitTransferToNumberAmazonConnectSuccess(c.TransferToNumberAmazonConnectSuccess)
	}
	if c.TransferToNumberError != nil {
		return visitor.VisitTransferToNumberError(c.TransferToNumberError)
	}
	if c.TransferToNumberExotelSuccess != nil {
		return visitor.VisitTransferToNumberExotelSuccess(c.TransferToNumberExotelSuccess)
	}
	if c.TransferToNumberSipSuccess != nil {
		return visitor.VisitTransferToNumberSipSuccess(c.TransferToNumberSipSuccess)
	}
	if c.TransferToNumberTwilioSuccess != nil {
		return visitor.VisitTransferToNumberTwilioSuccess(c.TransferToNumberTwilioSuccess)
	}
	if c.VoicemailDetectionSuccess != nil {
		return visitor.VisitVoicemailDetectionSuccess(c.VoicemailDetectionSuccess)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", c)
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) validate() error {
	if c == nil {
		return fmt.Errorf("type %T is nil", c)
	}
	var fields []string
	if c.Dummy != nil {
		fields = append(fields, "dummy")
	}
	if c.EndCallSuccess != nil {
		fields = append(fields, "end_call_success")
	}
	if c.EndProcedureError != nil {
		fields = append(fields, "end_procedure_error")
	}
	if c.EndProcedureSuccess != nil {
		fields = append(fields, "end_procedure_success")
	}
	if c.KnowledgeBaseRagSuccess != nil {
		fields = append(fields, "knowledge_base_rag_success")
	}
	if c.KnowledgeBaseSuccess != nil {
		fields = append(fields, "knowledge_base_success")
	}
	if c.LanguageDetectionSuccess != nil {
		fields = append(fields, "language_detection_success")
	}
	if c.PlayDtmfError != nil {
		fields = append(fields, "play_dtmf_error")
	}
	if c.PlayDtmfSuccess != nil {
		fields = append(fields, "play_dtmf_success")
	}
	if c.SkipTurnSuccess != nil {
		fields = append(fields, "skip_turn_success")
	}
	if c.StartProcedureError != nil {
		fields = append(fields, "start_procedure_error")
	}
	if c.StartProcedureSuccess != nil {
		fields = append(fields, "start_procedure_success")
	}
	if c.TestingToolResult != nil {
		fields = append(fields, "testing_tool_result")
	}
	if c.TransferToAgentError != nil {
		fields = append(fields, "transfer_to_agent_error")
	}
	if c.TransferToAgentSuccess != nil {
		fields = append(fields, "transfer_to_agent_success")
	}
	if c.TransferToNumberAmazonConnectSuccess != nil {
		fields = append(fields, "transfer_to_number_amazon_connect_success")
	}
	if c.TransferToNumberError != nil {
		fields = append(fields, "transfer_to_number_error")
	}
	if c.TransferToNumberExotelSuccess != nil {
		fields = append(fields, "transfer_to_number_exotel_success")
	}
	if c.TransferToNumberSipSuccess != nil {
		fields = append(fields, "transfer_to_number_sip_success")
	}
	if c.TransferToNumberTwilioSuccess != nil {
		fields = append(fields, "transfer_to_number_twilio_success")
	}
	if c.VoicemailDetectionSuccess != nil {
		fields = append(fields, "voicemail_detection_success")
	}
	if len(fields) == 0 {
		if c.ResultType != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", c, c.ResultType)
		}
		return fmt.Errorf("type %T is empty", c)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", c, fields)
	}
	if c.ResultType != "" {
		field := fields[0]
		if c.ResultType != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				c,
				c.ResultType,
				c,
			)
		}
	}
	return nil
}
