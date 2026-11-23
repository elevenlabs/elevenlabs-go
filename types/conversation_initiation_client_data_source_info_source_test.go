package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestConversationInitiationClientDataSourceInfoSourceVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ConversationInitiationClientDataSourceInfoSource
		wire    string
	}
	cases := []enumCase{
		{variant: ConversationInitiationClientDataSourceInfoSourceUnknown, wire: "\"unknown\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceAndroidSdk, wire: "\"android_sdk\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceNodeJsSdk, wire: "\"node_js_sdk\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceReactNativeSdk, wire: "\"react_native_sdk\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceReactSdk, wire: "\"react_sdk\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceJsSdk, wire: "\"js_sdk\""},
		{variant: ConversationInitiationClientDataSourceInfoSourcePythonSdk, wire: "\"python_sdk\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceWidget, wire: "\"widget\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceSipTrunk, wire: "\"sip_trunk\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceTwilio, wire: "\"twilio\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceExotel, wire: "\"exotel\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceGenesys, wire: "\"genesys\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceAvaya, wire: "\"avaya\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceAudiocodes, wire: "\"audiocodes\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceSwiftSdk, wire: "\"swift_sdk\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceWhatsapp, wire: "\"whatsapp\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceTwilioSms, wire: "\"twilio_sms\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceFlutterSdk, wire: "\"flutter_sdk\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceZendeskIntegration, wire: "\"zendesk_integration\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceSlackIntegration, wire: "\"slack_integration\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceTelegramIntegration, wire: "\"telegram_integration\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceIntercomIntegration, wire: "\"intercom_integration\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceFreshdeskIntegration, wire: "\"freshdesk_integration\""},
		{variant: ConversationInitiationClientDataSourceInfoSourcePlainIntegration, wire: "\"plain_integration\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceSalesforceIntegration, wire: "\"salesforce_integration\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceTemplatePreview, wire: "\"template_preview\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceGenesysBotConnector, wire: "\"genesys_bot_connector\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceSubagentTool, wire: "\"subagent_tool\""},
		{variant: ConversationInitiationClientDataSourceInfoSourceAgentMessagesApi, wire: "\"agent_messages_api\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ConversationInitiationClientDataSourceInfoSource
			if err := json.Unmarshal([]byte(tc.wire), &decoded); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if decoded != tc.variant {
				t.Fatalf("decoded mismatch: got %#v want %#v", decoded, tc.variant)
			}
			var want any
			var got any
			if err := json.Unmarshal([]byte(tc.wire), &want); err != nil {
				t.Fatalf("unmarshal want: %v", err)
			}
			if err := json.Unmarshal(encoded, &got); err != nil {
				t.Fatalf("unmarshal got: %v", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("round-trip mismatch\n got: %#v\nwant: %#v", got, want)
			}
		})
	}
}
