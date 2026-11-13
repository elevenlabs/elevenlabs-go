package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestConversationInitiationSourceVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ConversationInitiationSource
		wire    string
	}
	cases := []enumCase{
		{variant: ConversationInitiationSourceUnknown, wire: "\"unknown\""},
		{variant: ConversationInitiationSourceAndroidSdk, wire: "\"android_sdk\""},
		{variant: ConversationInitiationSourceNodeJsSdk, wire: "\"node_js_sdk\""},
		{variant: ConversationInitiationSourceReactNativeSdk, wire: "\"react_native_sdk\""},
		{variant: ConversationInitiationSourceReactSdk, wire: "\"react_sdk\""},
		{variant: ConversationInitiationSourceJsSdk, wire: "\"js_sdk\""},
		{variant: ConversationInitiationSourcePythonSdk, wire: "\"python_sdk\""},
		{variant: ConversationInitiationSourceWidget, wire: "\"widget\""},
		{variant: ConversationInitiationSourceSipTrunk, wire: "\"sip_trunk\""},
		{variant: ConversationInitiationSourceTwilio, wire: "\"twilio\""},
		{variant: ConversationInitiationSourceExotel, wire: "\"exotel\""},
		{variant: ConversationInitiationSourceGenesys, wire: "\"genesys\""},
		{variant: ConversationInitiationSourceAvaya, wire: "\"avaya\""},
		{variant: ConversationInitiationSourceAudiocodes, wire: "\"audiocodes\""},
		{variant: ConversationInitiationSourceSwiftSdk, wire: "\"swift_sdk\""},
		{variant: ConversationInitiationSourceWhatsapp, wire: "\"whatsapp\""},
		{variant: ConversationInitiationSourceTwilioSms, wire: "\"twilio_sms\""},
		{variant: ConversationInitiationSourceFlutterSdk, wire: "\"flutter_sdk\""},
		{variant: ConversationInitiationSourceZendeskIntegration, wire: "\"zendesk_integration\""},
		{variant: ConversationInitiationSourceSlackIntegration, wire: "\"slack_integration\""},
		{variant: ConversationInitiationSourceTelegramIntegration, wire: "\"telegram_integration\""},
		{variant: ConversationInitiationSourceIntercomIntegration, wire: "\"intercom_integration\""},
		{variant: ConversationInitiationSourceFreshdeskIntegration, wire: "\"freshdesk_integration\""},
		{variant: ConversationInitiationSourceSalesforceIntegration, wire: "\"salesforce_integration\""},
		{variant: ConversationInitiationSourceTemplatePreview, wire: "\"template_preview\""},
		{variant: ConversationInitiationSourceGenesysBotConnector, wire: "\"genesys_bot_connector\""},
		{variant: ConversationInitiationSourceSubagentTool, wire: "\"subagent_tool\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ConversationInitiationSource
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

func TestConversationInitiationSourceVariantsRoundTrip2(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ConversationInitiationSource
		wire    string
	}
	cases := []enumCase{
		{variant: ConversationInitiationSourceUnknown, wire: "\"unknown\""},
		{variant: ConversationInitiationSourceAndroidSdk, wire: "\"android_sdk\""},
		{variant: ConversationInitiationSourceNodeJsSdk, wire: "\"node_js_sdk\""},
		{variant: ConversationInitiationSourceReactNativeSdk, wire: "\"react_native_sdk\""},
		{variant: ConversationInitiationSourceReactSdk, wire: "\"react_sdk\""},
		{variant: ConversationInitiationSourceJsSdk, wire: "\"js_sdk\""},
		{variant: ConversationInitiationSourcePythonSdk, wire: "\"python_sdk\""},
		{variant: ConversationInitiationSourceWidget, wire: "\"widget\""},
		{variant: ConversationInitiationSourceSipTrunk, wire: "\"sip_trunk\""},
		{variant: ConversationInitiationSourceTwilio, wire: "\"twilio\""},
		{variant: ConversationInitiationSourceExotel, wire: "\"exotel\""},
		{variant: ConversationInitiationSourceGenesys, wire: "\"genesys\""},
		{variant: ConversationInitiationSourceAvaya, wire: "\"avaya\""},
		{variant: ConversationInitiationSourceAudiocodes, wire: "\"audiocodes\""},
		{variant: ConversationInitiationSourceSwiftSdk, wire: "\"swift_sdk\""},
		{variant: ConversationInitiationSourceWhatsapp, wire: "\"whatsapp\""},
		{variant: ConversationInitiationSourceTwilioSms, wire: "\"twilio_sms\""},
		{variant: ConversationInitiationSourceFlutterSdk, wire: "\"flutter_sdk\""},
		{variant: ConversationInitiationSourceZendeskIntegration, wire: "\"zendesk_integration\""},
		{variant: ConversationInitiationSourceSlackIntegration, wire: "\"slack_integration\""},
		{variant: ConversationInitiationSourceTelegramIntegration, wire: "\"telegram_integration\""},
		{variant: ConversationInitiationSourceIntercomIntegration, wire: "\"intercom_integration\""},
		{variant: ConversationInitiationSourceFreshdeskIntegration, wire: "\"freshdesk_integration\""},
		{variant: ConversationInitiationSourceSalesforceIntegration, wire: "\"salesforce_integration\""},
		{variant: ConversationInitiationSourceTemplatePreview, wire: "\"template_preview\""},
		{variant: ConversationInitiationSourceGenesysBotConnector, wire: "\"genesys_bot_connector\""},
		{variant: ConversationInitiationSourceSubagentTool, wire: "\"subagent_tool\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ConversationInitiationSource
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

func TestConversationInitiationSourceVariantsRoundTrip3(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ConversationInitiationSource
		wire    string
	}
	cases := []enumCase{
		{variant: ConversationInitiationSourceUnknown, wire: "\"unknown\""},
		{variant: ConversationInitiationSourceAndroidSdk, wire: "\"android_sdk\""},
		{variant: ConversationInitiationSourceNodeJsSdk, wire: "\"node_js_sdk\""},
		{variant: ConversationInitiationSourceReactNativeSdk, wire: "\"react_native_sdk\""},
		{variant: ConversationInitiationSourceReactSdk, wire: "\"react_sdk\""},
		{variant: ConversationInitiationSourceJsSdk, wire: "\"js_sdk\""},
		{variant: ConversationInitiationSourcePythonSdk, wire: "\"python_sdk\""},
		{variant: ConversationInitiationSourceWidget, wire: "\"widget\""},
		{variant: ConversationInitiationSourceSipTrunk, wire: "\"sip_trunk\""},
		{variant: ConversationInitiationSourceTwilio, wire: "\"twilio\""},
		{variant: ConversationInitiationSourceExotel, wire: "\"exotel\""},
		{variant: ConversationInitiationSourceGenesys, wire: "\"genesys\""},
		{variant: ConversationInitiationSourceAvaya, wire: "\"avaya\""},
		{variant: ConversationInitiationSourceAudiocodes, wire: "\"audiocodes\""},
		{variant: ConversationInitiationSourceSwiftSdk, wire: "\"swift_sdk\""},
		{variant: ConversationInitiationSourceWhatsapp, wire: "\"whatsapp\""},
		{variant: ConversationInitiationSourceTwilioSms, wire: "\"twilio_sms\""},
		{variant: ConversationInitiationSourceFlutterSdk, wire: "\"flutter_sdk\""},
		{variant: ConversationInitiationSourceZendeskIntegration, wire: "\"zendesk_integration\""},
		{variant: ConversationInitiationSourceSlackIntegration, wire: "\"slack_integration\""},
		{variant: ConversationInitiationSourceTelegramIntegration, wire: "\"telegram_integration\""},
		{variant: ConversationInitiationSourceIntercomIntegration, wire: "\"intercom_integration\""},
		{variant: ConversationInitiationSourceFreshdeskIntegration, wire: "\"freshdesk_integration\""},
		{variant: ConversationInitiationSourceSalesforceIntegration, wire: "\"salesforce_integration\""},
		{variant: ConversationInitiationSourceTemplatePreview, wire: "\"template_preview\""},
		{variant: ConversationInitiationSourceGenesysBotConnector, wire: "\"genesys_bot_connector\""},
		{variant: ConversationInitiationSourceSubagentTool, wire: "\"subagent_tool\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ConversationInitiationSource
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

func TestConversationInitiationSourceVariantsRoundTrip4(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ConversationInitiationSource
		wire    string
	}
	cases := []enumCase{
		{variant: ConversationInitiationSourceUnknown, wire: "\"unknown\""},
		{variant: ConversationInitiationSourceAndroidSdk, wire: "\"android_sdk\""},
		{variant: ConversationInitiationSourceNodeJsSdk, wire: "\"node_js_sdk\""},
		{variant: ConversationInitiationSourceReactNativeSdk, wire: "\"react_native_sdk\""},
		{variant: ConversationInitiationSourceReactSdk, wire: "\"react_sdk\""},
		{variant: ConversationInitiationSourceJsSdk, wire: "\"js_sdk\""},
		{variant: ConversationInitiationSourcePythonSdk, wire: "\"python_sdk\""},
		{variant: ConversationInitiationSourceWidget, wire: "\"widget\""},
		{variant: ConversationInitiationSourceSipTrunk, wire: "\"sip_trunk\""},
		{variant: ConversationInitiationSourceTwilio, wire: "\"twilio\""},
		{variant: ConversationInitiationSourceExotel, wire: "\"exotel\""},
		{variant: ConversationInitiationSourceGenesys, wire: "\"genesys\""},
		{variant: ConversationInitiationSourceAvaya, wire: "\"avaya\""},
		{variant: ConversationInitiationSourceAudiocodes, wire: "\"audiocodes\""},
		{variant: ConversationInitiationSourceSwiftSdk, wire: "\"swift_sdk\""},
		{variant: ConversationInitiationSourceWhatsapp, wire: "\"whatsapp\""},
		{variant: ConversationInitiationSourceTwilioSms, wire: "\"twilio_sms\""},
		{variant: ConversationInitiationSourceFlutterSdk, wire: "\"flutter_sdk\""},
		{variant: ConversationInitiationSourceZendeskIntegration, wire: "\"zendesk_integration\""},
		{variant: ConversationInitiationSourceSlackIntegration, wire: "\"slack_integration\""},
		{variant: ConversationInitiationSourceTelegramIntegration, wire: "\"telegram_integration\""},
		{variant: ConversationInitiationSourceIntercomIntegration, wire: "\"intercom_integration\""},
		{variant: ConversationInitiationSourceFreshdeskIntegration, wire: "\"freshdesk_integration\""},
		{variant: ConversationInitiationSourceSalesforceIntegration, wire: "\"salesforce_integration\""},
		{variant: ConversationInitiationSourceTemplatePreview, wire: "\"template_preview\""},
		{variant: ConversationInitiationSourceGenesysBotConnector, wire: "\"genesys_bot_connector\""},
		{variant: ConversationInitiationSourceSubagentTool, wire: "\"subagent_tool\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ConversationInitiationSource
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
