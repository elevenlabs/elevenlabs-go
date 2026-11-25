package types

import (
	"fmt"
)

// Source of the conversation initiation
type ConversationInitiationClientDataSourceInfoSource string

const (
	ConversationInitiationClientDataSourceInfoSourceUnknown               ConversationInitiationClientDataSourceInfoSource = "unknown"
	ConversationInitiationClientDataSourceInfoSourceAndroidSdk            ConversationInitiationClientDataSourceInfoSource = "android_sdk"
	ConversationInitiationClientDataSourceInfoSourceNodeJsSdk             ConversationInitiationClientDataSourceInfoSource = "node_js_sdk"
	ConversationInitiationClientDataSourceInfoSourceReactNativeSdk        ConversationInitiationClientDataSourceInfoSource = "react_native_sdk"
	ConversationInitiationClientDataSourceInfoSourceReactSdk              ConversationInitiationClientDataSourceInfoSource = "react_sdk"
	ConversationInitiationClientDataSourceInfoSourceJsSdk                 ConversationInitiationClientDataSourceInfoSource = "js_sdk"
	ConversationInitiationClientDataSourceInfoSourcePythonSdk             ConversationInitiationClientDataSourceInfoSource = "python_sdk"
	ConversationInitiationClientDataSourceInfoSourceWidget                ConversationInitiationClientDataSourceInfoSource = "widget"
	ConversationInitiationClientDataSourceInfoSourceSipTrunk              ConversationInitiationClientDataSourceInfoSource = "sip_trunk"
	ConversationInitiationClientDataSourceInfoSourceTwilio                ConversationInitiationClientDataSourceInfoSource = "twilio"
	ConversationInitiationClientDataSourceInfoSourceExotel                ConversationInitiationClientDataSourceInfoSource = "exotel"
	ConversationInitiationClientDataSourceInfoSourceGenesys               ConversationInitiationClientDataSourceInfoSource = "genesys"
	ConversationInitiationClientDataSourceInfoSourceAvaya                 ConversationInitiationClientDataSourceInfoSource = "avaya"
	ConversationInitiationClientDataSourceInfoSourceAudiocodes            ConversationInitiationClientDataSourceInfoSource = "audiocodes"
	ConversationInitiationClientDataSourceInfoSourceSwiftSdk              ConversationInitiationClientDataSourceInfoSource = "swift_sdk"
	ConversationInitiationClientDataSourceInfoSourceWhatsapp              ConversationInitiationClientDataSourceInfoSource = "whatsapp"
	ConversationInitiationClientDataSourceInfoSourceTwilioSms             ConversationInitiationClientDataSourceInfoSource = "twilio_sms"
	ConversationInitiationClientDataSourceInfoSourceFlutterSdk            ConversationInitiationClientDataSourceInfoSource = "flutter_sdk"
	ConversationInitiationClientDataSourceInfoSourceUnitySdk              ConversationInitiationClientDataSourceInfoSource = "unity_sdk"
	ConversationInitiationClientDataSourceInfoSourceZendeskIntegration    ConversationInitiationClientDataSourceInfoSource = "zendesk_integration"
	ConversationInitiationClientDataSourceInfoSourceSlackIntegration      ConversationInitiationClientDataSourceInfoSource = "slack_integration"
	ConversationInitiationClientDataSourceInfoSourceTelegramIntegration   ConversationInitiationClientDataSourceInfoSource = "telegram_integration"
	ConversationInitiationClientDataSourceInfoSourceIntercomIntegration   ConversationInitiationClientDataSourceInfoSource = "intercom_integration"
	ConversationInitiationClientDataSourceInfoSourceFreshdeskIntegration  ConversationInitiationClientDataSourceInfoSource = "freshdesk_integration"
	ConversationInitiationClientDataSourceInfoSourcePlainIntegration      ConversationInitiationClientDataSourceInfoSource = "plain_integration"
	ConversationInitiationClientDataSourceInfoSourceSalesforceIntegration ConversationInitiationClientDataSourceInfoSource = "salesforce_integration"
	ConversationInitiationClientDataSourceInfoSourceTemplatePreview       ConversationInitiationClientDataSourceInfoSource = "template_preview"
	ConversationInitiationClientDataSourceInfoSourceGenesysBotConnector   ConversationInitiationClientDataSourceInfoSource = "genesys_bot_connector"
	ConversationInitiationClientDataSourceInfoSourceSubagentTool          ConversationInitiationClientDataSourceInfoSource = "subagent_tool"
	ConversationInitiationClientDataSourceInfoSourceAgentMessagesApi      ConversationInitiationClientDataSourceInfoSource = "agent_messages_api"
)

func NewConversationInitiationClientDataSourceInfoSourceFromString(s string) (ConversationInitiationClientDataSourceInfoSource, error) {
	switch s {
	case "unknown":
		return ConversationInitiationClientDataSourceInfoSourceUnknown, nil
	case "android_sdk":
		return ConversationInitiationClientDataSourceInfoSourceAndroidSdk, nil
	case "node_js_sdk":
		return ConversationInitiationClientDataSourceInfoSourceNodeJsSdk, nil
	case "react_native_sdk":
		return ConversationInitiationClientDataSourceInfoSourceReactNativeSdk, nil
	case "react_sdk":
		return ConversationInitiationClientDataSourceInfoSourceReactSdk, nil
	case "js_sdk":
		return ConversationInitiationClientDataSourceInfoSourceJsSdk, nil
	case "python_sdk":
		return ConversationInitiationClientDataSourceInfoSourcePythonSdk, nil
	case "widget":
		return ConversationInitiationClientDataSourceInfoSourceWidget, nil
	case "sip_trunk":
		return ConversationInitiationClientDataSourceInfoSourceSipTrunk, nil
	case "twilio":
		return ConversationInitiationClientDataSourceInfoSourceTwilio, nil
	case "exotel":
		return ConversationInitiationClientDataSourceInfoSourceExotel, nil
	case "genesys":
		return ConversationInitiationClientDataSourceInfoSourceGenesys, nil
	case "avaya":
		return ConversationInitiationClientDataSourceInfoSourceAvaya, nil
	case "audiocodes":
		return ConversationInitiationClientDataSourceInfoSourceAudiocodes, nil
	case "swift_sdk":
		return ConversationInitiationClientDataSourceInfoSourceSwiftSdk, nil
	case "whatsapp":
		return ConversationInitiationClientDataSourceInfoSourceWhatsapp, nil
	case "twilio_sms":
		return ConversationInitiationClientDataSourceInfoSourceTwilioSms, nil
	case "flutter_sdk":
		return ConversationInitiationClientDataSourceInfoSourceFlutterSdk, nil
	case "unity_sdk":
		return ConversationInitiationClientDataSourceInfoSourceUnitySdk, nil
	case "zendesk_integration":
		return ConversationInitiationClientDataSourceInfoSourceZendeskIntegration, nil
	case "slack_integration":
		return ConversationInitiationClientDataSourceInfoSourceSlackIntegration, nil
	case "telegram_integration":
		return ConversationInitiationClientDataSourceInfoSourceTelegramIntegration, nil
	case "intercom_integration":
		return ConversationInitiationClientDataSourceInfoSourceIntercomIntegration, nil
	case "freshdesk_integration":
		return ConversationInitiationClientDataSourceInfoSourceFreshdeskIntegration, nil
	case "plain_integration":
		return ConversationInitiationClientDataSourceInfoSourcePlainIntegration, nil
	case "salesforce_integration":
		return ConversationInitiationClientDataSourceInfoSourceSalesforceIntegration, nil
	case "template_preview":
		return ConversationInitiationClientDataSourceInfoSourceTemplatePreview, nil
	case "genesys_bot_connector":
		return ConversationInitiationClientDataSourceInfoSourceGenesysBotConnector, nil
	case "subagent_tool":
		return ConversationInitiationClientDataSourceInfoSourceSubagentTool, nil
	case "agent_messages_api":
		return ConversationInitiationClientDataSourceInfoSourceAgentMessagesApi, nil
	}
	var t ConversationInitiationClientDataSourceInfoSource
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConversationInitiationClientDataSourceInfoSource) Ptr() *ConversationInitiationClientDataSourceInfoSource {
	return &e
}
