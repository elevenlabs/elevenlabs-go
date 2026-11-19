package types

import (
	"fmt"
)

// Enum representing the possible sources for conversation initiation.
type ConversationInitiationSource string

const (
	ConversationInitiationSourceUnknown               ConversationInitiationSource = "unknown"
	ConversationInitiationSourceAndroidSdk            ConversationInitiationSource = "android_sdk"
	ConversationInitiationSourceNodeJsSdk             ConversationInitiationSource = "node_js_sdk"
	ConversationInitiationSourceReactNativeSdk        ConversationInitiationSource = "react_native_sdk"
	ConversationInitiationSourceReactSdk              ConversationInitiationSource = "react_sdk"
	ConversationInitiationSourceJsSdk                 ConversationInitiationSource = "js_sdk"
	ConversationInitiationSourcePythonSdk             ConversationInitiationSource = "python_sdk"
	ConversationInitiationSourceWidget                ConversationInitiationSource = "widget"
	ConversationInitiationSourceSipTrunk              ConversationInitiationSource = "sip_trunk"
	ConversationInitiationSourceTwilio                ConversationInitiationSource = "twilio"
	ConversationInitiationSourceExotel                ConversationInitiationSource = "exotel"
	ConversationInitiationSourceGenesys               ConversationInitiationSource = "genesys"
	ConversationInitiationSourceAvaya                 ConversationInitiationSource = "avaya"
	ConversationInitiationSourceAudiocodes            ConversationInitiationSource = "audiocodes"
	ConversationInitiationSourceSwiftSdk              ConversationInitiationSource = "swift_sdk"
	ConversationInitiationSourceWhatsapp              ConversationInitiationSource = "whatsapp"
	ConversationInitiationSourceTwilioSms             ConversationInitiationSource = "twilio_sms"
	ConversationInitiationSourceFlutterSdk            ConversationInitiationSource = "flutter_sdk"
	ConversationInitiationSourceZendeskIntegration    ConversationInitiationSource = "zendesk_integration"
	ConversationInitiationSourceSlackIntegration      ConversationInitiationSource = "slack_integration"
	ConversationInitiationSourceTelegramIntegration   ConversationInitiationSource = "telegram_integration"
	ConversationInitiationSourceIntercomIntegration   ConversationInitiationSource = "intercom_integration"
	ConversationInitiationSourceFreshdeskIntegration  ConversationInitiationSource = "freshdesk_integration"
	ConversationInitiationSourceSalesforceIntegration ConversationInitiationSource = "salesforce_integration"
	ConversationInitiationSourceTemplatePreview       ConversationInitiationSource = "template_preview"
	ConversationInitiationSourceGenesysBotConnector   ConversationInitiationSource = "genesys_bot_connector"
	ConversationInitiationSourceSubagentTool          ConversationInitiationSource = "subagent_tool"
	ConversationInitiationSourceAgentMessagesApi      ConversationInitiationSource = "agent_messages_api"
)

func NewConversationInitiationSourceFromString(s string) (ConversationInitiationSource, error) {
	switch s {
	case "unknown":
		return ConversationInitiationSourceUnknown, nil
	case "android_sdk":
		return ConversationInitiationSourceAndroidSdk, nil
	case "node_js_sdk":
		return ConversationInitiationSourceNodeJsSdk, nil
	case "react_native_sdk":
		return ConversationInitiationSourceReactNativeSdk, nil
	case "react_sdk":
		return ConversationInitiationSourceReactSdk, nil
	case "js_sdk":
		return ConversationInitiationSourceJsSdk, nil
	case "python_sdk":
		return ConversationInitiationSourcePythonSdk, nil
	case "widget":
		return ConversationInitiationSourceWidget, nil
	case "sip_trunk":
		return ConversationInitiationSourceSipTrunk, nil
	case "twilio":
		return ConversationInitiationSourceTwilio, nil
	case "exotel":
		return ConversationInitiationSourceExotel, nil
	case "genesys":
		return ConversationInitiationSourceGenesys, nil
	case "avaya":
		return ConversationInitiationSourceAvaya, nil
	case "audiocodes":
		return ConversationInitiationSourceAudiocodes, nil
	case "swift_sdk":
		return ConversationInitiationSourceSwiftSdk, nil
	case "whatsapp":
		return ConversationInitiationSourceWhatsapp, nil
	case "twilio_sms":
		return ConversationInitiationSourceTwilioSms, nil
	case "flutter_sdk":
		return ConversationInitiationSourceFlutterSdk, nil
	case "zendesk_integration":
		return ConversationInitiationSourceZendeskIntegration, nil
	case "slack_integration":
		return ConversationInitiationSourceSlackIntegration, nil
	case "telegram_integration":
		return ConversationInitiationSourceTelegramIntegration, nil
	case "intercom_integration":
		return ConversationInitiationSourceIntercomIntegration, nil
	case "freshdesk_integration":
		return ConversationInitiationSourceFreshdeskIntegration, nil
	case "salesforce_integration":
		return ConversationInitiationSourceSalesforceIntegration, nil
	case "template_preview":
		return ConversationInitiationSourceTemplatePreview, nil
	case "genesys_bot_connector":
		return ConversationInitiationSourceGenesysBotConnector, nil
	case "subagent_tool":
		return ConversationInitiationSourceSubagentTool, nil
	case "agent_messages_api":
		return ConversationInitiationSourceAgentMessagesApi, nil
	}
	var t ConversationInitiationSource
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConversationInitiationSource) Ptr() *ConversationInitiationSource {
	return &e
}
