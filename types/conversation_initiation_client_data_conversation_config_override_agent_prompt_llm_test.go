package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm
		wire    string
	}
	cases := []enumCase{
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4OMini, wire: "\"gpt-4o-mini\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4O, wire: "\"gpt-4o\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4, wire: "\"gpt-4\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4Turbo, wire: "\"gpt-4-turbo\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt41, wire: "\"gpt-4.1\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt41Mini, wire: "\"gpt-4.1-mini\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt41Nano, wire: "\"gpt-4.1-nano\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5, wire: "\"gpt-5\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt51, wire: "\"gpt-5.1\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt52, wire: "\"gpt-5.2\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt52ChatLatest, wire: "\"gpt-5.2-chat-latest\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt54, wire: "\"gpt-5.4\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt54Mini, wire: "\"gpt-5.4-mini\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt54Nano, wire: "\"gpt-5.4-nano\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt55, wire: "\"gpt-5.5\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt56Sol, wire: "\"gpt-5.6-sol\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt56Terra, wire: "\"gpt-5.6-terra\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt56Luna, wire: "\"gpt-5.6-luna\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt6Astra, wire: "\"gpt-6-astra\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt6Sol, wire: "\"gpt-6-sol\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt6Luna, wire: "\"gpt-6-luna\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt61Sol, wire: "\"gpt-6.1-sol\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt61Terra, wire: "\"gpt-6.1-terra\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt62Astra, wire: "\"gpt-6.2-astra\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt62Sol, wire: "\"gpt-6.2-sol\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5Mini, wire: "\"gpt-5-mini\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5Nano, wire: "\"gpt-5-nano\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt35Turbo, wire: "\"gpt-3.5-turbo\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini15Pro, wire: "\"gemini-1.5-pro\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini15Flash, wire: "\"gemini-1.5-flash\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini20Flash, wire: "\"gemini-2.0-flash\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini20FlashLite, wire: "\"gemini-2.0-flash-lite\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini25FlashLite, wire: "\"gemini-2.5-flash-lite\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini25Flash, wire: "\"gemini-2.5-flash\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini3ProPreview, wire: "\"gemini-3-pro-preview\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini3FlashPreview, wire: "\"gemini-3-flash-preview\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini31ProPreview, wire: "\"gemini-3.1-pro-preview\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini31FlashLitePreview, wire: "\"gemini-3.1-flash-lite-preview\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini31FlashLite, wire: "\"gemini-3.1-flash-lite\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini35Flash, wire: "\"gemini-3.5-flash\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini35FlashLite, wire: "\"gemini-3.5-flash-lite\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini36Flash, wire: "\"gemini-3.6-flash\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini37Flash, wire: "\"gemini-3.7-flash\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini38Flash, wire: "\"gemini-3.8-flash\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini4Argon, wire: "\"gemini-4-argon\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini4Hydrogen, wire: "\"gemini-4-hydrogen\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini4Neon, wire: "\"gemini-4-neon\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeSonnet45, wire: "\"claude-sonnet-4-5\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeOpus46, wire: "\"claude-opus-4-6\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeOpus47, wire: "\"claude-opus-4-7\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeOpus48, wire: "\"claude-opus-4-8\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeOpus5, wire: "\"claude-opus-5\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeOpus55, wire: "\"claude-opus-5-5\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeSonnet46, wire: "\"claude-sonnet-4-6\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeSonnet5, wire: "\"claude-sonnet-5\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeSonnet55, wire: "\"claude-sonnet-5-5\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeSonnet4, wire: "\"claude-sonnet-4\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeHaiku45, wire: "\"claude-haiku-4-5\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeHaiku55, wire: "\"claude-haiku-5-5\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude37Sonnet, wire: "\"claude-3-7-sonnet\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude35Sonnet, wire: "\"claude-3-5-sonnet\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude35SonnetV1, wire: "\"claude-3-5-sonnet-v1\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude3Haiku, wire: "\"claude-3-haiku\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmCustomLlm, wire: "\"custom-llm\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmQwen34B, wire: "\"qwen3-4b\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmQwen330BA3B, wire: "\"qwen3-30b-a3b\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmQwen3635BA3B, wire: "\"qwen36-35b-a3b\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmQwen35397BA17B, wire: "\"qwen35-397b-a17b\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGptOss20B, wire: "\"gpt-oss-20b\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGptOss120B, wire: "\"gpt-oss-120b\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGlm45AirFp8, wire: "\"glm-45-air-fp8\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGlm52, wire: "\"glm-52\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmDeepseekV41Flash, wire: "\"deepseek-v41-flash\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmNemotron35SuperVlNvfp4, wire: "\"nemotron-35-super-vl-nvfp4\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini25FlashPreview092025, wire: "\"gemini-2.5-flash-preview-09-2025\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini25FlashLitePreview092025, wire: "\"gemini-2.5-flash-lite-preview-09-2025\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini25FlashPreview0520, wire: "\"gemini-2.5-flash-preview-05-20\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini25FlashPreview0417, wire: "\"gemini-2.5-flash-preview-04-17\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini25FlashLitePreview0617, wire: "\"gemini-2.5-flash-lite-preview-06-17\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini20FlashLite001, wire: "\"gemini-2.0-flash-lite-001\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini20Flash001, wire: "\"gemini-2.0-flash-001\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini15Flash002, wire: "\"gemini-1.5-flash-002\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini15Flash001, wire: "\"gemini-1.5-flash-001\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini15Pro002, wire: "\"gemini-1.5-pro-002\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini15Pro001, wire: "\"gemini-1.5-pro-001\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeSonnet4At20250514, wire: "\"claude-sonnet-4@20250514\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeSonnet45At20250929, wire: "\"claude-sonnet-4-5@20250929\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeHaiku45At20251001, wire: "\"claude-haiku-4-5@20251001\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude37SonnetAt20250219, wire: "\"claude-3-7-sonnet@20250219\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude35SonnetAt20240620, wire: "\"claude-3-5-sonnet@20240620\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude35SonnetV2At20241022, wire: "\"claude-3-5-sonnet-v2@20241022\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude3HaikuAt20240307, wire: "\"claude-3-haiku@20240307\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt520250807, wire: "\"gpt-5-2025-08-07\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5120251113, wire: "\"gpt-5.1-2025-11-13\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5220251211, wire: "\"gpt-5.2-2025-12-11\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5420260305, wire: "\"gpt-5.4-2026-03-05\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt54Mini20260317, wire: "\"gpt-5.4-mini-2026-03-17\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt54Nano20260317, wire: "\"gpt-5.4-nano-2026-03-17\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5520260423, wire: "\"gpt-5.5-2026-04-23\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5Mini20250807, wire: "\"gpt-5-mini-2025-08-07\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5Nano20250807, wire: "\"gpt-5-nano-2025-08-07\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4120250414, wire: "\"gpt-4.1-2025-04-14\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt41Mini20250414, wire: "\"gpt-4.1-mini-2025-04-14\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt41Nano20250414, wire: "\"gpt-4.1-nano-2025-04-14\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4OMini20240718, wire: "\"gpt-4o-mini-2024-07-18\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4O20241120, wire: "\"gpt-4o-2024-11-20\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4O20240806, wire: "\"gpt-4o-2024-08-06\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4O20240513, wire: "\"gpt-4o-2024-05-13\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt40613, wire: "\"gpt-4-0613\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt40314, wire: "\"gpt-4-0314\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4Turbo20240409, wire: "\"gpt-4-turbo-2024-04-09\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt35Turbo0125, wire: "\"gpt-3.5-turbo-0125\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt35Turbo1106, wire: "\"gpt-3.5-turbo-1106\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmWattTool8B, wire: "\"watt-tool-8b\""},
		{variant: ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmWattTool70B, wire: "\"watt-tool-70b\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm
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
