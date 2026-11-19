package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestLlmVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant Llm
		wire    string
	}
	cases := []enumCase{
		{variant: LlmGpt4OMini, wire: "\"gpt-4o-mini\""},
		{variant: LlmGpt4O, wire: "\"gpt-4o\""},
		{variant: LlmGpt4, wire: "\"gpt-4\""},
		{variant: LlmGpt4Turbo, wire: "\"gpt-4-turbo\""},
		{variant: LlmGpt41, wire: "\"gpt-4.1\""},
		{variant: LlmGpt41Mini, wire: "\"gpt-4.1-mini\""},
		{variant: LlmGpt41Nano, wire: "\"gpt-4.1-nano\""},
		{variant: LlmGpt5, wire: "\"gpt-5\""},
		{variant: LlmGpt51, wire: "\"gpt-5.1\""},
		{variant: LlmGpt52, wire: "\"gpt-5.2\""},
		{variant: LlmGpt52ChatLatest, wire: "\"gpt-5.2-chat-latest\""},
		{variant: LlmGpt54, wire: "\"gpt-5.4\""},
		{variant: LlmGpt54Mini, wire: "\"gpt-5.4-mini\""},
		{variant: LlmGpt54Nano, wire: "\"gpt-5.4-nano\""},
		{variant: LlmGpt55, wire: "\"gpt-5.5\""},
		{variant: LlmGpt56Sol, wire: "\"gpt-5.6-sol\""},
		{variant: LlmGpt56Terra, wire: "\"gpt-5.6-terra\""},
		{variant: LlmGpt56Luna, wire: "\"gpt-5.6-luna\""},
		{variant: LlmGpt6Astra, wire: "\"gpt-6-astra\""},
		{variant: LlmGpt6Sol, wire: "\"gpt-6-sol\""},
		{variant: LlmGpt6Luna, wire: "\"gpt-6-luna\""},
		{variant: LlmGpt61Sol, wire: "\"gpt-6.1-sol\""},
		{variant: LlmGpt61Terra, wire: "\"gpt-6.1-terra\""},
		{variant: LlmGpt5Mini, wire: "\"gpt-5-mini\""},
		{variant: LlmGpt5Nano, wire: "\"gpt-5-nano\""},
		{variant: LlmGpt35Turbo, wire: "\"gpt-3.5-turbo\""},
		{variant: LlmGemini15Pro, wire: "\"gemini-1.5-pro\""},
		{variant: LlmGemini15Flash, wire: "\"gemini-1.5-flash\""},
		{variant: LlmGemini20Flash, wire: "\"gemini-2.0-flash\""},
		{variant: LlmGemini20FlashLite, wire: "\"gemini-2.0-flash-lite\""},
		{variant: LlmGemini25FlashLite, wire: "\"gemini-2.5-flash-lite\""},
		{variant: LlmGemini25Flash, wire: "\"gemini-2.5-flash\""},
		{variant: LlmGemini3ProPreview, wire: "\"gemini-3-pro-preview\""},
		{variant: LlmGemini3FlashPreview, wire: "\"gemini-3-flash-preview\""},
		{variant: LlmGemini31ProPreview, wire: "\"gemini-3.1-pro-preview\""},
		{variant: LlmGemini31FlashLitePreview, wire: "\"gemini-3.1-flash-lite-preview\""},
		{variant: LlmGemini31FlashLite, wire: "\"gemini-3.1-flash-lite\""},
		{variant: LlmGemini35Flash, wire: "\"gemini-3.5-flash\""},
		{variant: LlmGemini35FlashLite, wire: "\"gemini-3.5-flash-lite\""},
		{variant: LlmGemini36Flash, wire: "\"gemini-3.6-flash\""},
		{variant: LlmGemini37Flash, wire: "\"gemini-3.7-flash\""},
		{variant: LlmGemini38Flash, wire: "\"gemini-3.8-flash\""},
		{variant: LlmGemini4Argon, wire: "\"gemini-4-argon\""},
		{variant: LlmGemini4Hydrogen, wire: "\"gemini-4-hydrogen\""},
		{variant: LlmClaudeSonnet45, wire: "\"claude-sonnet-4-5\""},
		{variant: LlmClaudeOpus46, wire: "\"claude-opus-4-6\""},
		{variant: LlmClaudeOpus47, wire: "\"claude-opus-4-7\""},
		{variant: LlmClaudeOpus48, wire: "\"claude-opus-4-8\""},
		{variant: LlmClaudeOpus5, wire: "\"claude-opus-5\""},
		{variant: LlmClaudeOpus55, wire: "\"claude-opus-5-5\""},
		{variant: LlmClaudeSonnet46, wire: "\"claude-sonnet-4-6\""},
		{variant: LlmClaudeSonnet5, wire: "\"claude-sonnet-5\""},
		{variant: LlmClaudeSonnet55, wire: "\"claude-sonnet-5-5\""},
		{variant: LlmClaudeSonnet4, wire: "\"claude-sonnet-4\""},
		{variant: LlmClaudeHaiku45, wire: "\"claude-haiku-4-5\""},
		{variant: LlmClaudeHaiku55, wire: "\"claude-haiku-5-5\""},
		{variant: LlmClaude37Sonnet, wire: "\"claude-3-7-sonnet\""},
		{variant: LlmClaude35Sonnet, wire: "\"claude-3-5-sonnet\""},
		{variant: LlmClaude35SonnetV1, wire: "\"claude-3-5-sonnet-v1\""},
		{variant: LlmClaude3Haiku, wire: "\"claude-3-haiku\""},
		{variant: LlmCustomLlm, wire: "\"custom-llm\""},
		{variant: LlmQwen34B, wire: "\"qwen3-4b\""},
		{variant: LlmQwen330BA3B, wire: "\"qwen3-30b-a3b\""},
		{variant: LlmQwen3635BA3B, wire: "\"qwen36-35b-a3b\""},
		{variant: LlmQwen35397BA17B, wire: "\"qwen35-397b-a17b\""},
		{variant: LlmGptOss20B, wire: "\"gpt-oss-20b\""},
		{variant: LlmGptOss120B, wire: "\"gpt-oss-120b\""},
		{variant: LlmGlm45AirFp8, wire: "\"glm-45-air-fp8\""},
		{variant: LlmGlm52, wire: "\"glm-52\""},
		{variant: LlmDeepseekV41Flash, wire: "\"deepseek-v41-flash\""},
		{variant: LlmNemotron35SuperVlNvfp4, wire: "\"nemotron-35-super-vl-nvfp4\""},
		{variant: LlmGemini25FlashPreview092025, wire: "\"gemini-2.5-flash-preview-09-2025\""},
		{variant: LlmGemini25FlashLitePreview092025, wire: "\"gemini-2.5-flash-lite-preview-09-2025\""},
		{variant: LlmGemini25FlashPreview0520, wire: "\"gemini-2.5-flash-preview-05-20\""},
		{variant: LlmGemini25FlashPreview0417, wire: "\"gemini-2.5-flash-preview-04-17\""},
		{variant: LlmGemini25FlashLitePreview0617, wire: "\"gemini-2.5-flash-lite-preview-06-17\""},
		{variant: LlmGemini20FlashLite001, wire: "\"gemini-2.0-flash-lite-001\""},
		{variant: LlmGemini20Flash001, wire: "\"gemini-2.0-flash-001\""},
		{variant: LlmGemini15Flash002, wire: "\"gemini-1.5-flash-002\""},
		{variant: LlmGemini15Flash001, wire: "\"gemini-1.5-flash-001\""},
		{variant: LlmGemini15Pro002, wire: "\"gemini-1.5-pro-002\""},
		{variant: LlmGemini15Pro001, wire: "\"gemini-1.5-pro-001\""},
		{variant: LlmClaudeSonnet4At20250514, wire: "\"claude-sonnet-4@20250514\""},
		{variant: LlmClaudeSonnet45At20250929, wire: "\"claude-sonnet-4-5@20250929\""},
		{variant: LlmClaudeHaiku45At20251001, wire: "\"claude-haiku-4-5@20251001\""},
		{variant: LlmClaude37SonnetAt20250219, wire: "\"claude-3-7-sonnet@20250219\""},
		{variant: LlmClaude35SonnetAt20240620, wire: "\"claude-3-5-sonnet@20240620\""},
		{variant: LlmClaude35SonnetV2At20241022, wire: "\"claude-3-5-sonnet-v2@20241022\""},
		{variant: LlmClaude3HaikuAt20240307, wire: "\"claude-3-haiku@20240307\""},
		{variant: LlmGpt520250807, wire: "\"gpt-5-2025-08-07\""},
		{variant: LlmGpt5120251113, wire: "\"gpt-5.1-2025-11-13\""},
		{variant: LlmGpt5220251211, wire: "\"gpt-5.2-2025-12-11\""},
		{variant: LlmGpt5420260305, wire: "\"gpt-5.4-2026-03-05\""},
		{variant: LlmGpt54Mini20260317, wire: "\"gpt-5.4-mini-2026-03-17\""},
		{variant: LlmGpt54Nano20260317, wire: "\"gpt-5.4-nano-2026-03-17\""},
		{variant: LlmGpt5520260423, wire: "\"gpt-5.5-2026-04-23\""},
		{variant: LlmGpt5Mini20250807, wire: "\"gpt-5-mini-2025-08-07\""},
		{variant: LlmGpt5Nano20250807, wire: "\"gpt-5-nano-2025-08-07\""},
		{variant: LlmGpt4120250414, wire: "\"gpt-4.1-2025-04-14\""},
		{variant: LlmGpt41Mini20250414, wire: "\"gpt-4.1-mini-2025-04-14\""},
		{variant: LlmGpt41Nano20250414, wire: "\"gpt-4.1-nano-2025-04-14\""},
		{variant: LlmGpt4OMini20240718, wire: "\"gpt-4o-mini-2024-07-18\""},
		{variant: LlmGpt4O20241120, wire: "\"gpt-4o-2024-11-20\""},
		{variant: LlmGpt4O20240806, wire: "\"gpt-4o-2024-08-06\""},
		{variant: LlmGpt4O20240513, wire: "\"gpt-4o-2024-05-13\""},
		{variant: LlmGpt40613, wire: "\"gpt-4-0613\""},
		{variant: LlmGpt40314, wire: "\"gpt-4-0314\""},
		{variant: LlmGpt4Turbo20240409, wire: "\"gpt-4-turbo-2024-04-09\""},
		{variant: LlmGpt35Turbo0125, wire: "\"gpt-3.5-turbo-0125\""},
		{variant: LlmGpt35Turbo1106, wire: "\"gpt-3.5-turbo-1106\""},
		{variant: LlmWattTool8B, wire: "\"watt-tool-8b\""},
		{variant: LlmWattTool70B, wire: "\"watt-tool-70b\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded Llm
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

func TestLlmVariantsRoundTrip2(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant Llm
		wire    string
	}
	cases := []enumCase{
		{variant: LlmGpt4OMini, wire: "\"gpt-4o-mini\""},
		{variant: LlmGpt4O, wire: "\"gpt-4o\""},
		{variant: LlmGpt4, wire: "\"gpt-4\""},
		{variant: LlmGpt4Turbo, wire: "\"gpt-4-turbo\""},
		{variant: LlmGpt41, wire: "\"gpt-4.1\""},
		{variant: LlmGpt41Mini, wire: "\"gpt-4.1-mini\""},
		{variant: LlmGpt41Nano, wire: "\"gpt-4.1-nano\""},
		{variant: LlmGpt5, wire: "\"gpt-5\""},
		{variant: LlmGpt51, wire: "\"gpt-5.1\""},
		{variant: LlmGpt52, wire: "\"gpt-5.2\""},
		{variant: LlmGpt52ChatLatest, wire: "\"gpt-5.2-chat-latest\""},
		{variant: LlmGpt54, wire: "\"gpt-5.4\""},
		{variant: LlmGpt54Mini, wire: "\"gpt-5.4-mini\""},
		{variant: LlmGpt54Nano, wire: "\"gpt-5.4-nano\""},
		{variant: LlmGpt55, wire: "\"gpt-5.5\""},
		{variant: LlmGpt56Sol, wire: "\"gpt-5.6-sol\""},
		{variant: LlmGpt56Terra, wire: "\"gpt-5.6-terra\""},
		{variant: LlmGpt56Luna, wire: "\"gpt-5.6-luna\""},
		{variant: LlmGpt6Astra, wire: "\"gpt-6-astra\""},
		{variant: LlmGpt6Sol, wire: "\"gpt-6-sol\""},
		{variant: LlmGpt6Luna, wire: "\"gpt-6-luna\""},
		{variant: LlmGpt61Sol, wire: "\"gpt-6.1-sol\""},
		{variant: LlmGpt61Terra, wire: "\"gpt-6.1-terra\""},
		{variant: LlmGpt5Mini, wire: "\"gpt-5-mini\""},
		{variant: LlmGpt5Nano, wire: "\"gpt-5-nano\""},
		{variant: LlmGpt35Turbo, wire: "\"gpt-3.5-turbo\""},
		{variant: LlmGemini15Pro, wire: "\"gemini-1.5-pro\""},
		{variant: LlmGemini15Flash, wire: "\"gemini-1.5-flash\""},
		{variant: LlmGemini20Flash, wire: "\"gemini-2.0-flash\""},
		{variant: LlmGemini20FlashLite, wire: "\"gemini-2.0-flash-lite\""},
		{variant: LlmGemini25FlashLite, wire: "\"gemini-2.5-flash-lite\""},
		{variant: LlmGemini25Flash, wire: "\"gemini-2.5-flash\""},
		{variant: LlmGemini3ProPreview, wire: "\"gemini-3-pro-preview\""},
		{variant: LlmGemini3FlashPreview, wire: "\"gemini-3-flash-preview\""},
		{variant: LlmGemini31ProPreview, wire: "\"gemini-3.1-pro-preview\""},
		{variant: LlmGemini31FlashLitePreview, wire: "\"gemini-3.1-flash-lite-preview\""},
		{variant: LlmGemini31FlashLite, wire: "\"gemini-3.1-flash-lite\""},
		{variant: LlmGemini35Flash, wire: "\"gemini-3.5-flash\""},
		{variant: LlmGemini35FlashLite, wire: "\"gemini-3.5-flash-lite\""},
		{variant: LlmGemini36Flash, wire: "\"gemini-3.6-flash\""},
		{variant: LlmGemini37Flash, wire: "\"gemini-3.7-flash\""},
		{variant: LlmGemini38Flash, wire: "\"gemini-3.8-flash\""},
		{variant: LlmGemini4Argon, wire: "\"gemini-4-argon\""},
		{variant: LlmGemini4Hydrogen, wire: "\"gemini-4-hydrogen\""},
		{variant: LlmClaudeSonnet45, wire: "\"claude-sonnet-4-5\""},
		{variant: LlmClaudeOpus46, wire: "\"claude-opus-4-6\""},
		{variant: LlmClaudeOpus47, wire: "\"claude-opus-4-7\""},
		{variant: LlmClaudeOpus48, wire: "\"claude-opus-4-8\""},
		{variant: LlmClaudeOpus5, wire: "\"claude-opus-5\""},
		{variant: LlmClaudeOpus55, wire: "\"claude-opus-5-5\""},
		{variant: LlmClaudeSonnet46, wire: "\"claude-sonnet-4-6\""},
		{variant: LlmClaudeSonnet5, wire: "\"claude-sonnet-5\""},
		{variant: LlmClaudeSonnet55, wire: "\"claude-sonnet-5-5\""},
		{variant: LlmClaudeSonnet4, wire: "\"claude-sonnet-4\""},
		{variant: LlmClaudeHaiku45, wire: "\"claude-haiku-4-5\""},
		{variant: LlmClaudeHaiku55, wire: "\"claude-haiku-5-5\""},
		{variant: LlmClaude37Sonnet, wire: "\"claude-3-7-sonnet\""},
		{variant: LlmClaude35Sonnet, wire: "\"claude-3-5-sonnet\""},
		{variant: LlmClaude35SonnetV1, wire: "\"claude-3-5-sonnet-v1\""},
		{variant: LlmClaude3Haiku, wire: "\"claude-3-haiku\""},
		{variant: LlmCustomLlm, wire: "\"custom-llm\""},
		{variant: LlmQwen34B, wire: "\"qwen3-4b\""},
		{variant: LlmQwen330BA3B, wire: "\"qwen3-30b-a3b\""},
		{variant: LlmQwen3635BA3B, wire: "\"qwen36-35b-a3b\""},
		{variant: LlmQwen35397BA17B, wire: "\"qwen35-397b-a17b\""},
		{variant: LlmGptOss20B, wire: "\"gpt-oss-20b\""},
		{variant: LlmGptOss120B, wire: "\"gpt-oss-120b\""},
		{variant: LlmGlm45AirFp8, wire: "\"glm-45-air-fp8\""},
		{variant: LlmGlm52, wire: "\"glm-52\""},
		{variant: LlmDeepseekV41Flash, wire: "\"deepseek-v41-flash\""},
		{variant: LlmNemotron35SuperVlNvfp4, wire: "\"nemotron-35-super-vl-nvfp4\""},
		{variant: LlmGemini25FlashPreview092025, wire: "\"gemini-2.5-flash-preview-09-2025\""},
		{variant: LlmGemini25FlashLitePreview092025, wire: "\"gemini-2.5-flash-lite-preview-09-2025\""},
		{variant: LlmGemini25FlashPreview0520, wire: "\"gemini-2.5-flash-preview-05-20\""},
		{variant: LlmGemini25FlashPreview0417, wire: "\"gemini-2.5-flash-preview-04-17\""},
		{variant: LlmGemini25FlashLitePreview0617, wire: "\"gemini-2.5-flash-lite-preview-06-17\""},
		{variant: LlmGemini20FlashLite001, wire: "\"gemini-2.0-flash-lite-001\""},
		{variant: LlmGemini20Flash001, wire: "\"gemini-2.0-flash-001\""},
		{variant: LlmGemini15Flash002, wire: "\"gemini-1.5-flash-002\""},
		{variant: LlmGemini15Flash001, wire: "\"gemini-1.5-flash-001\""},
		{variant: LlmGemini15Pro002, wire: "\"gemini-1.5-pro-002\""},
		{variant: LlmGemini15Pro001, wire: "\"gemini-1.5-pro-001\""},
		{variant: LlmClaudeSonnet4At20250514, wire: "\"claude-sonnet-4@20250514\""},
		{variant: LlmClaudeSonnet45At20250929, wire: "\"claude-sonnet-4-5@20250929\""},
		{variant: LlmClaudeHaiku45At20251001, wire: "\"claude-haiku-4-5@20251001\""},
		{variant: LlmClaude37SonnetAt20250219, wire: "\"claude-3-7-sonnet@20250219\""},
		{variant: LlmClaude35SonnetAt20240620, wire: "\"claude-3-5-sonnet@20240620\""},
		{variant: LlmClaude35SonnetV2At20241022, wire: "\"claude-3-5-sonnet-v2@20241022\""},
		{variant: LlmClaude3HaikuAt20240307, wire: "\"claude-3-haiku@20240307\""},
		{variant: LlmGpt520250807, wire: "\"gpt-5-2025-08-07\""},
		{variant: LlmGpt5120251113, wire: "\"gpt-5.1-2025-11-13\""},
		{variant: LlmGpt5220251211, wire: "\"gpt-5.2-2025-12-11\""},
		{variant: LlmGpt5420260305, wire: "\"gpt-5.4-2026-03-05\""},
		{variant: LlmGpt54Mini20260317, wire: "\"gpt-5.4-mini-2026-03-17\""},
		{variant: LlmGpt54Nano20260317, wire: "\"gpt-5.4-nano-2026-03-17\""},
		{variant: LlmGpt5520260423, wire: "\"gpt-5.5-2026-04-23\""},
		{variant: LlmGpt5Mini20250807, wire: "\"gpt-5-mini-2025-08-07\""},
		{variant: LlmGpt5Nano20250807, wire: "\"gpt-5-nano-2025-08-07\""},
		{variant: LlmGpt4120250414, wire: "\"gpt-4.1-2025-04-14\""},
		{variant: LlmGpt41Mini20250414, wire: "\"gpt-4.1-mini-2025-04-14\""},
		{variant: LlmGpt41Nano20250414, wire: "\"gpt-4.1-nano-2025-04-14\""},
		{variant: LlmGpt4OMini20240718, wire: "\"gpt-4o-mini-2024-07-18\""},
		{variant: LlmGpt4O20241120, wire: "\"gpt-4o-2024-11-20\""},
		{variant: LlmGpt4O20240806, wire: "\"gpt-4o-2024-08-06\""},
		{variant: LlmGpt4O20240513, wire: "\"gpt-4o-2024-05-13\""},
		{variant: LlmGpt40613, wire: "\"gpt-4-0613\""},
		{variant: LlmGpt40314, wire: "\"gpt-4-0314\""},
		{variant: LlmGpt4Turbo20240409, wire: "\"gpt-4-turbo-2024-04-09\""},
		{variant: LlmGpt35Turbo0125, wire: "\"gpt-3.5-turbo-0125\""},
		{variant: LlmGpt35Turbo1106, wire: "\"gpt-3.5-turbo-1106\""},
		{variant: LlmWattTool8B, wire: "\"watt-tool-8b\""},
		{variant: LlmWattTool70B, wire: "\"watt-tool-70b\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded Llm
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

func TestLlmVariantsRoundTrip3(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant Llm
		wire    string
	}
	cases := []enumCase{
		{variant: LlmGpt4OMini, wire: "\"gpt-4o-mini\""},
		{variant: LlmGpt4O, wire: "\"gpt-4o\""},
		{variant: LlmGpt4, wire: "\"gpt-4\""},
		{variant: LlmGpt4Turbo, wire: "\"gpt-4-turbo\""},
		{variant: LlmGpt41, wire: "\"gpt-4.1\""},
		{variant: LlmGpt41Mini, wire: "\"gpt-4.1-mini\""},
		{variant: LlmGpt41Nano, wire: "\"gpt-4.1-nano\""},
		{variant: LlmGpt5, wire: "\"gpt-5\""},
		{variant: LlmGpt51, wire: "\"gpt-5.1\""},
		{variant: LlmGpt52, wire: "\"gpt-5.2\""},
		{variant: LlmGpt52ChatLatest, wire: "\"gpt-5.2-chat-latest\""},
		{variant: LlmGpt54, wire: "\"gpt-5.4\""},
		{variant: LlmGpt54Mini, wire: "\"gpt-5.4-mini\""},
		{variant: LlmGpt54Nano, wire: "\"gpt-5.4-nano\""},
		{variant: LlmGpt55, wire: "\"gpt-5.5\""},
		{variant: LlmGpt56Sol, wire: "\"gpt-5.6-sol\""},
		{variant: LlmGpt56Terra, wire: "\"gpt-5.6-terra\""},
		{variant: LlmGpt56Luna, wire: "\"gpt-5.6-luna\""},
		{variant: LlmGpt6Astra, wire: "\"gpt-6-astra\""},
		{variant: LlmGpt6Sol, wire: "\"gpt-6-sol\""},
		{variant: LlmGpt6Luna, wire: "\"gpt-6-luna\""},
		{variant: LlmGpt61Sol, wire: "\"gpt-6.1-sol\""},
		{variant: LlmGpt61Terra, wire: "\"gpt-6.1-terra\""},
		{variant: LlmGpt5Mini, wire: "\"gpt-5-mini\""},
		{variant: LlmGpt5Nano, wire: "\"gpt-5-nano\""},
		{variant: LlmGpt35Turbo, wire: "\"gpt-3.5-turbo\""},
		{variant: LlmGemini15Pro, wire: "\"gemini-1.5-pro\""},
		{variant: LlmGemini15Flash, wire: "\"gemini-1.5-flash\""},
		{variant: LlmGemini20Flash, wire: "\"gemini-2.0-flash\""},
		{variant: LlmGemini20FlashLite, wire: "\"gemini-2.0-flash-lite\""},
		{variant: LlmGemini25FlashLite, wire: "\"gemini-2.5-flash-lite\""},
		{variant: LlmGemini25Flash, wire: "\"gemini-2.5-flash\""},
		{variant: LlmGemini3ProPreview, wire: "\"gemini-3-pro-preview\""},
		{variant: LlmGemini3FlashPreview, wire: "\"gemini-3-flash-preview\""},
		{variant: LlmGemini31ProPreview, wire: "\"gemini-3.1-pro-preview\""},
		{variant: LlmGemini31FlashLitePreview, wire: "\"gemini-3.1-flash-lite-preview\""},
		{variant: LlmGemini31FlashLite, wire: "\"gemini-3.1-flash-lite\""},
		{variant: LlmGemini35Flash, wire: "\"gemini-3.5-flash\""},
		{variant: LlmGemini35FlashLite, wire: "\"gemini-3.5-flash-lite\""},
		{variant: LlmGemini36Flash, wire: "\"gemini-3.6-flash\""},
		{variant: LlmGemini37Flash, wire: "\"gemini-3.7-flash\""},
		{variant: LlmGemini38Flash, wire: "\"gemini-3.8-flash\""},
		{variant: LlmGemini4Argon, wire: "\"gemini-4-argon\""},
		{variant: LlmGemini4Hydrogen, wire: "\"gemini-4-hydrogen\""},
		{variant: LlmClaudeSonnet45, wire: "\"claude-sonnet-4-5\""},
		{variant: LlmClaudeOpus46, wire: "\"claude-opus-4-6\""},
		{variant: LlmClaudeOpus47, wire: "\"claude-opus-4-7\""},
		{variant: LlmClaudeOpus48, wire: "\"claude-opus-4-8\""},
		{variant: LlmClaudeOpus5, wire: "\"claude-opus-5\""},
		{variant: LlmClaudeOpus55, wire: "\"claude-opus-5-5\""},
		{variant: LlmClaudeSonnet46, wire: "\"claude-sonnet-4-6\""},
		{variant: LlmClaudeSonnet5, wire: "\"claude-sonnet-5\""},
		{variant: LlmClaudeSonnet55, wire: "\"claude-sonnet-5-5\""},
		{variant: LlmClaudeSonnet4, wire: "\"claude-sonnet-4\""},
		{variant: LlmClaudeHaiku45, wire: "\"claude-haiku-4-5\""},
		{variant: LlmClaudeHaiku55, wire: "\"claude-haiku-5-5\""},
		{variant: LlmClaude37Sonnet, wire: "\"claude-3-7-sonnet\""},
		{variant: LlmClaude35Sonnet, wire: "\"claude-3-5-sonnet\""},
		{variant: LlmClaude35SonnetV1, wire: "\"claude-3-5-sonnet-v1\""},
		{variant: LlmClaude3Haiku, wire: "\"claude-3-haiku\""},
		{variant: LlmCustomLlm, wire: "\"custom-llm\""},
		{variant: LlmQwen34B, wire: "\"qwen3-4b\""},
		{variant: LlmQwen330BA3B, wire: "\"qwen3-30b-a3b\""},
		{variant: LlmQwen3635BA3B, wire: "\"qwen36-35b-a3b\""},
		{variant: LlmQwen35397BA17B, wire: "\"qwen35-397b-a17b\""},
		{variant: LlmGptOss20B, wire: "\"gpt-oss-20b\""},
		{variant: LlmGptOss120B, wire: "\"gpt-oss-120b\""},
		{variant: LlmGlm45AirFp8, wire: "\"glm-45-air-fp8\""},
		{variant: LlmGlm52, wire: "\"glm-52\""},
		{variant: LlmDeepseekV41Flash, wire: "\"deepseek-v41-flash\""},
		{variant: LlmNemotron35SuperVlNvfp4, wire: "\"nemotron-35-super-vl-nvfp4\""},
		{variant: LlmGemini25FlashPreview092025, wire: "\"gemini-2.5-flash-preview-09-2025\""},
		{variant: LlmGemini25FlashLitePreview092025, wire: "\"gemini-2.5-flash-lite-preview-09-2025\""},
		{variant: LlmGemini25FlashPreview0520, wire: "\"gemini-2.5-flash-preview-05-20\""},
		{variant: LlmGemini25FlashPreview0417, wire: "\"gemini-2.5-flash-preview-04-17\""},
		{variant: LlmGemini25FlashLitePreview0617, wire: "\"gemini-2.5-flash-lite-preview-06-17\""},
		{variant: LlmGemini20FlashLite001, wire: "\"gemini-2.0-flash-lite-001\""},
		{variant: LlmGemini20Flash001, wire: "\"gemini-2.0-flash-001\""},
		{variant: LlmGemini15Flash002, wire: "\"gemini-1.5-flash-002\""},
		{variant: LlmGemini15Flash001, wire: "\"gemini-1.5-flash-001\""},
		{variant: LlmGemini15Pro002, wire: "\"gemini-1.5-pro-002\""},
		{variant: LlmGemini15Pro001, wire: "\"gemini-1.5-pro-001\""},
		{variant: LlmClaudeSonnet4At20250514, wire: "\"claude-sonnet-4@20250514\""},
		{variant: LlmClaudeSonnet45At20250929, wire: "\"claude-sonnet-4-5@20250929\""},
		{variant: LlmClaudeHaiku45At20251001, wire: "\"claude-haiku-4-5@20251001\""},
		{variant: LlmClaude37SonnetAt20250219, wire: "\"claude-3-7-sonnet@20250219\""},
		{variant: LlmClaude35SonnetAt20240620, wire: "\"claude-3-5-sonnet@20240620\""},
		{variant: LlmClaude35SonnetV2At20241022, wire: "\"claude-3-5-sonnet-v2@20241022\""},
		{variant: LlmClaude3HaikuAt20240307, wire: "\"claude-3-haiku@20240307\""},
		{variant: LlmGpt520250807, wire: "\"gpt-5-2025-08-07\""},
		{variant: LlmGpt5120251113, wire: "\"gpt-5.1-2025-11-13\""},
		{variant: LlmGpt5220251211, wire: "\"gpt-5.2-2025-12-11\""},
		{variant: LlmGpt5420260305, wire: "\"gpt-5.4-2026-03-05\""},
		{variant: LlmGpt54Mini20260317, wire: "\"gpt-5.4-mini-2026-03-17\""},
		{variant: LlmGpt54Nano20260317, wire: "\"gpt-5.4-nano-2026-03-17\""},
		{variant: LlmGpt5520260423, wire: "\"gpt-5.5-2026-04-23\""},
		{variant: LlmGpt5Mini20250807, wire: "\"gpt-5-mini-2025-08-07\""},
		{variant: LlmGpt5Nano20250807, wire: "\"gpt-5-nano-2025-08-07\""},
		{variant: LlmGpt4120250414, wire: "\"gpt-4.1-2025-04-14\""},
		{variant: LlmGpt41Mini20250414, wire: "\"gpt-4.1-mini-2025-04-14\""},
		{variant: LlmGpt41Nano20250414, wire: "\"gpt-4.1-nano-2025-04-14\""},
		{variant: LlmGpt4OMini20240718, wire: "\"gpt-4o-mini-2024-07-18\""},
		{variant: LlmGpt4O20241120, wire: "\"gpt-4o-2024-11-20\""},
		{variant: LlmGpt4O20240806, wire: "\"gpt-4o-2024-08-06\""},
		{variant: LlmGpt4O20240513, wire: "\"gpt-4o-2024-05-13\""},
		{variant: LlmGpt40613, wire: "\"gpt-4-0613\""},
		{variant: LlmGpt40314, wire: "\"gpt-4-0314\""},
		{variant: LlmGpt4Turbo20240409, wire: "\"gpt-4-turbo-2024-04-09\""},
		{variant: LlmGpt35Turbo0125, wire: "\"gpt-3.5-turbo-0125\""},
		{variant: LlmGpt35Turbo1106, wire: "\"gpt-3.5-turbo-1106\""},
		{variant: LlmWattTool8B, wire: "\"watt-tool-8b\""},
		{variant: LlmWattTool70B, wire: "\"watt-tool-70b\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded Llm
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

func TestLlmVariantsRoundTrip4(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant Llm
		wire    string
	}
	cases := []enumCase{
		{variant: LlmGpt4OMini, wire: "\"gpt-4o-mini\""},
		{variant: LlmGpt4O, wire: "\"gpt-4o\""},
		{variant: LlmGpt4, wire: "\"gpt-4\""},
		{variant: LlmGpt4Turbo, wire: "\"gpt-4-turbo\""},
		{variant: LlmGpt41, wire: "\"gpt-4.1\""},
		{variant: LlmGpt41Mini, wire: "\"gpt-4.1-mini\""},
		{variant: LlmGpt41Nano, wire: "\"gpt-4.1-nano\""},
		{variant: LlmGpt5, wire: "\"gpt-5\""},
		{variant: LlmGpt51, wire: "\"gpt-5.1\""},
		{variant: LlmGpt52, wire: "\"gpt-5.2\""},
		{variant: LlmGpt52ChatLatest, wire: "\"gpt-5.2-chat-latest\""},
		{variant: LlmGpt54, wire: "\"gpt-5.4\""},
		{variant: LlmGpt54Mini, wire: "\"gpt-5.4-mini\""},
		{variant: LlmGpt54Nano, wire: "\"gpt-5.4-nano\""},
		{variant: LlmGpt55, wire: "\"gpt-5.5\""},
		{variant: LlmGpt56Sol, wire: "\"gpt-5.6-sol\""},
		{variant: LlmGpt56Terra, wire: "\"gpt-5.6-terra\""},
		{variant: LlmGpt56Luna, wire: "\"gpt-5.6-luna\""},
		{variant: LlmGpt6Astra, wire: "\"gpt-6-astra\""},
		{variant: LlmGpt6Sol, wire: "\"gpt-6-sol\""},
		{variant: LlmGpt6Luna, wire: "\"gpt-6-luna\""},
		{variant: LlmGpt61Sol, wire: "\"gpt-6.1-sol\""},
		{variant: LlmGpt61Terra, wire: "\"gpt-6.1-terra\""},
		{variant: LlmGpt5Mini, wire: "\"gpt-5-mini\""},
		{variant: LlmGpt5Nano, wire: "\"gpt-5-nano\""},
		{variant: LlmGpt35Turbo, wire: "\"gpt-3.5-turbo\""},
		{variant: LlmGemini15Pro, wire: "\"gemini-1.5-pro\""},
		{variant: LlmGemini15Flash, wire: "\"gemini-1.5-flash\""},
		{variant: LlmGemini20Flash, wire: "\"gemini-2.0-flash\""},
		{variant: LlmGemini20FlashLite, wire: "\"gemini-2.0-flash-lite\""},
		{variant: LlmGemini25FlashLite, wire: "\"gemini-2.5-flash-lite\""},
		{variant: LlmGemini25Flash, wire: "\"gemini-2.5-flash\""},
		{variant: LlmGemini3ProPreview, wire: "\"gemini-3-pro-preview\""},
		{variant: LlmGemini3FlashPreview, wire: "\"gemini-3-flash-preview\""},
		{variant: LlmGemini31ProPreview, wire: "\"gemini-3.1-pro-preview\""},
		{variant: LlmGemini31FlashLitePreview, wire: "\"gemini-3.1-flash-lite-preview\""},
		{variant: LlmGemini31FlashLite, wire: "\"gemini-3.1-flash-lite\""},
		{variant: LlmGemini35Flash, wire: "\"gemini-3.5-flash\""},
		{variant: LlmGemini35FlashLite, wire: "\"gemini-3.5-flash-lite\""},
		{variant: LlmGemini36Flash, wire: "\"gemini-3.6-flash\""},
		{variant: LlmGemini37Flash, wire: "\"gemini-3.7-flash\""},
		{variant: LlmGemini38Flash, wire: "\"gemini-3.8-flash\""},
		{variant: LlmGemini4Argon, wire: "\"gemini-4-argon\""},
		{variant: LlmGemini4Hydrogen, wire: "\"gemini-4-hydrogen\""},
		{variant: LlmClaudeSonnet45, wire: "\"claude-sonnet-4-5\""},
		{variant: LlmClaudeOpus46, wire: "\"claude-opus-4-6\""},
		{variant: LlmClaudeOpus47, wire: "\"claude-opus-4-7\""},
		{variant: LlmClaudeOpus48, wire: "\"claude-opus-4-8\""},
		{variant: LlmClaudeOpus5, wire: "\"claude-opus-5\""},
		{variant: LlmClaudeOpus55, wire: "\"claude-opus-5-5\""},
		{variant: LlmClaudeSonnet46, wire: "\"claude-sonnet-4-6\""},
		{variant: LlmClaudeSonnet5, wire: "\"claude-sonnet-5\""},
		{variant: LlmClaudeSonnet55, wire: "\"claude-sonnet-5-5\""},
		{variant: LlmClaudeSonnet4, wire: "\"claude-sonnet-4\""},
		{variant: LlmClaudeHaiku45, wire: "\"claude-haiku-4-5\""},
		{variant: LlmClaudeHaiku55, wire: "\"claude-haiku-5-5\""},
		{variant: LlmClaude37Sonnet, wire: "\"claude-3-7-sonnet\""},
		{variant: LlmClaude35Sonnet, wire: "\"claude-3-5-sonnet\""},
		{variant: LlmClaude35SonnetV1, wire: "\"claude-3-5-sonnet-v1\""},
		{variant: LlmClaude3Haiku, wire: "\"claude-3-haiku\""},
		{variant: LlmCustomLlm, wire: "\"custom-llm\""},
		{variant: LlmQwen34B, wire: "\"qwen3-4b\""},
		{variant: LlmQwen330BA3B, wire: "\"qwen3-30b-a3b\""},
		{variant: LlmQwen3635BA3B, wire: "\"qwen36-35b-a3b\""},
		{variant: LlmQwen35397BA17B, wire: "\"qwen35-397b-a17b\""},
		{variant: LlmGptOss20B, wire: "\"gpt-oss-20b\""},
		{variant: LlmGptOss120B, wire: "\"gpt-oss-120b\""},
		{variant: LlmGlm45AirFp8, wire: "\"glm-45-air-fp8\""},
		{variant: LlmGlm52, wire: "\"glm-52\""},
		{variant: LlmDeepseekV41Flash, wire: "\"deepseek-v41-flash\""},
		{variant: LlmNemotron35SuperVlNvfp4, wire: "\"nemotron-35-super-vl-nvfp4\""},
		{variant: LlmGemini25FlashPreview092025, wire: "\"gemini-2.5-flash-preview-09-2025\""},
		{variant: LlmGemini25FlashLitePreview092025, wire: "\"gemini-2.5-flash-lite-preview-09-2025\""},
		{variant: LlmGemini25FlashPreview0520, wire: "\"gemini-2.5-flash-preview-05-20\""},
		{variant: LlmGemini25FlashPreview0417, wire: "\"gemini-2.5-flash-preview-04-17\""},
		{variant: LlmGemini25FlashLitePreview0617, wire: "\"gemini-2.5-flash-lite-preview-06-17\""},
		{variant: LlmGemini20FlashLite001, wire: "\"gemini-2.0-flash-lite-001\""},
		{variant: LlmGemini20Flash001, wire: "\"gemini-2.0-flash-001\""},
		{variant: LlmGemini15Flash002, wire: "\"gemini-1.5-flash-002\""},
		{variant: LlmGemini15Flash001, wire: "\"gemini-1.5-flash-001\""},
		{variant: LlmGemini15Pro002, wire: "\"gemini-1.5-pro-002\""},
		{variant: LlmGemini15Pro001, wire: "\"gemini-1.5-pro-001\""},
		{variant: LlmClaudeSonnet4At20250514, wire: "\"claude-sonnet-4@20250514\""},
		{variant: LlmClaudeSonnet45At20250929, wire: "\"claude-sonnet-4-5@20250929\""},
		{variant: LlmClaudeHaiku45At20251001, wire: "\"claude-haiku-4-5@20251001\""},
		{variant: LlmClaude37SonnetAt20250219, wire: "\"claude-3-7-sonnet@20250219\""},
		{variant: LlmClaude35SonnetAt20240620, wire: "\"claude-3-5-sonnet@20240620\""},
		{variant: LlmClaude35SonnetV2At20241022, wire: "\"claude-3-5-sonnet-v2@20241022\""},
		{variant: LlmClaude3HaikuAt20240307, wire: "\"claude-3-haiku@20240307\""},
		{variant: LlmGpt520250807, wire: "\"gpt-5-2025-08-07\""},
		{variant: LlmGpt5120251113, wire: "\"gpt-5.1-2025-11-13\""},
		{variant: LlmGpt5220251211, wire: "\"gpt-5.2-2025-12-11\""},
		{variant: LlmGpt5420260305, wire: "\"gpt-5.4-2026-03-05\""},
		{variant: LlmGpt54Mini20260317, wire: "\"gpt-5.4-mini-2026-03-17\""},
		{variant: LlmGpt54Nano20260317, wire: "\"gpt-5.4-nano-2026-03-17\""},
		{variant: LlmGpt5520260423, wire: "\"gpt-5.5-2026-04-23\""},
		{variant: LlmGpt5Mini20250807, wire: "\"gpt-5-mini-2025-08-07\""},
		{variant: LlmGpt5Nano20250807, wire: "\"gpt-5-nano-2025-08-07\""},
		{variant: LlmGpt4120250414, wire: "\"gpt-4.1-2025-04-14\""},
		{variant: LlmGpt41Mini20250414, wire: "\"gpt-4.1-mini-2025-04-14\""},
		{variant: LlmGpt41Nano20250414, wire: "\"gpt-4.1-nano-2025-04-14\""},
		{variant: LlmGpt4OMini20240718, wire: "\"gpt-4o-mini-2024-07-18\""},
		{variant: LlmGpt4O20241120, wire: "\"gpt-4o-2024-11-20\""},
		{variant: LlmGpt4O20240806, wire: "\"gpt-4o-2024-08-06\""},
		{variant: LlmGpt4O20240513, wire: "\"gpt-4o-2024-05-13\""},
		{variant: LlmGpt40613, wire: "\"gpt-4-0613\""},
		{variant: LlmGpt40314, wire: "\"gpt-4-0314\""},
		{variant: LlmGpt4Turbo20240409, wire: "\"gpt-4-turbo-2024-04-09\""},
		{variant: LlmGpt35Turbo0125, wire: "\"gpt-3.5-turbo-0125\""},
		{variant: LlmGpt35Turbo1106, wire: "\"gpt-3.5-turbo-1106\""},
		{variant: LlmWattTool8B, wire: "\"watt-tool-8b\""},
		{variant: LlmWattTool70B, wire: "\"watt-tool-70b\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded Llm
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

func TestLlmVariantsRoundTrip5(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant Llm
		wire    string
	}
	cases := []enumCase{
		{variant: LlmGpt4OMini, wire: "\"gpt-4o-mini\""},
		{variant: LlmGpt4O, wire: "\"gpt-4o\""},
		{variant: LlmGpt4, wire: "\"gpt-4\""},
		{variant: LlmGpt4Turbo, wire: "\"gpt-4-turbo\""},
		{variant: LlmGpt41, wire: "\"gpt-4.1\""},
		{variant: LlmGpt41Mini, wire: "\"gpt-4.1-mini\""},
		{variant: LlmGpt41Nano, wire: "\"gpt-4.1-nano\""},
		{variant: LlmGpt5, wire: "\"gpt-5\""},
		{variant: LlmGpt51, wire: "\"gpt-5.1\""},
		{variant: LlmGpt52, wire: "\"gpt-5.2\""},
		{variant: LlmGpt52ChatLatest, wire: "\"gpt-5.2-chat-latest\""},
		{variant: LlmGpt54, wire: "\"gpt-5.4\""},
		{variant: LlmGpt54Mini, wire: "\"gpt-5.4-mini\""},
		{variant: LlmGpt54Nano, wire: "\"gpt-5.4-nano\""},
		{variant: LlmGpt55, wire: "\"gpt-5.5\""},
		{variant: LlmGpt56Sol, wire: "\"gpt-5.6-sol\""},
		{variant: LlmGpt56Terra, wire: "\"gpt-5.6-terra\""},
		{variant: LlmGpt56Luna, wire: "\"gpt-5.6-luna\""},
		{variant: LlmGpt6Astra, wire: "\"gpt-6-astra\""},
		{variant: LlmGpt6Sol, wire: "\"gpt-6-sol\""},
		{variant: LlmGpt6Luna, wire: "\"gpt-6-luna\""},
		{variant: LlmGpt61Sol, wire: "\"gpt-6.1-sol\""},
		{variant: LlmGpt61Terra, wire: "\"gpt-6.1-terra\""},
		{variant: LlmGpt5Mini, wire: "\"gpt-5-mini\""},
		{variant: LlmGpt5Nano, wire: "\"gpt-5-nano\""},
		{variant: LlmGpt35Turbo, wire: "\"gpt-3.5-turbo\""},
		{variant: LlmGemini15Pro, wire: "\"gemini-1.5-pro\""},
		{variant: LlmGemini15Flash, wire: "\"gemini-1.5-flash\""},
		{variant: LlmGemini20Flash, wire: "\"gemini-2.0-flash\""},
		{variant: LlmGemini20FlashLite, wire: "\"gemini-2.0-flash-lite\""},
		{variant: LlmGemini25FlashLite, wire: "\"gemini-2.5-flash-lite\""},
		{variant: LlmGemini25Flash, wire: "\"gemini-2.5-flash\""},
		{variant: LlmGemini3ProPreview, wire: "\"gemini-3-pro-preview\""},
		{variant: LlmGemini3FlashPreview, wire: "\"gemini-3-flash-preview\""},
		{variant: LlmGemini31ProPreview, wire: "\"gemini-3.1-pro-preview\""},
		{variant: LlmGemini31FlashLitePreview, wire: "\"gemini-3.1-flash-lite-preview\""},
		{variant: LlmGemini31FlashLite, wire: "\"gemini-3.1-flash-lite\""},
		{variant: LlmGemini35Flash, wire: "\"gemini-3.5-flash\""},
		{variant: LlmGemini35FlashLite, wire: "\"gemini-3.5-flash-lite\""},
		{variant: LlmGemini36Flash, wire: "\"gemini-3.6-flash\""},
		{variant: LlmGemini37Flash, wire: "\"gemini-3.7-flash\""},
		{variant: LlmGemini38Flash, wire: "\"gemini-3.8-flash\""},
		{variant: LlmGemini4Argon, wire: "\"gemini-4-argon\""},
		{variant: LlmGemini4Hydrogen, wire: "\"gemini-4-hydrogen\""},
		{variant: LlmClaudeSonnet45, wire: "\"claude-sonnet-4-5\""},
		{variant: LlmClaudeOpus46, wire: "\"claude-opus-4-6\""},
		{variant: LlmClaudeOpus47, wire: "\"claude-opus-4-7\""},
		{variant: LlmClaudeOpus48, wire: "\"claude-opus-4-8\""},
		{variant: LlmClaudeOpus5, wire: "\"claude-opus-5\""},
		{variant: LlmClaudeOpus55, wire: "\"claude-opus-5-5\""},
		{variant: LlmClaudeSonnet46, wire: "\"claude-sonnet-4-6\""},
		{variant: LlmClaudeSonnet5, wire: "\"claude-sonnet-5\""},
		{variant: LlmClaudeSonnet55, wire: "\"claude-sonnet-5-5\""},
		{variant: LlmClaudeSonnet4, wire: "\"claude-sonnet-4\""},
		{variant: LlmClaudeHaiku45, wire: "\"claude-haiku-4-5\""},
		{variant: LlmClaudeHaiku55, wire: "\"claude-haiku-5-5\""},
		{variant: LlmClaude37Sonnet, wire: "\"claude-3-7-sonnet\""},
		{variant: LlmClaude35Sonnet, wire: "\"claude-3-5-sonnet\""},
		{variant: LlmClaude35SonnetV1, wire: "\"claude-3-5-sonnet-v1\""},
		{variant: LlmClaude3Haiku, wire: "\"claude-3-haiku\""},
		{variant: LlmCustomLlm, wire: "\"custom-llm\""},
		{variant: LlmQwen34B, wire: "\"qwen3-4b\""},
		{variant: LlmQwen330BA3B, wire: "\"qwen3-30b-a3b\""},
		{variant: LlmQwen3635BA3B, wire: "\"qwen36-35b-a3b\""},
		{variant: LlmQwen35397BA17B, wire: "\"qwen35-397b-a17b\""},
		{variant: LlmGptOss20B, wire: "\"gpt-oss-20b\""},
		{variant: LlmGptOss120B, wire: "\"gpt-oss-120b\""},
		{variant: LlmGlm45AirFp8, wire: "\"glm-45-air-fp8\""},
		{variant: LlmGlm52, wire: "\"glm-52\""},
		{variant: LlmDeepseekV41Flash, wire: "\"deepseek-v41-flash\""},
		{variant: LlmNemotron35SuperVlNvfp4, wire: "\"nemotron-35-super-vl-nvfp4\""},
		{variant: LlmGemini25FlashPreview092025, wire: "\"gemini-2.5-flash-preview-09-2025\""},
		{variant: LlmGemini25FlashLitePreview092025, wire: "\"gemini-2.5-flash-lite-preview-09-2025\""},
		{variant: LlmGemini25FlashPreview0520, wire: "\"gemini-2.5-flash-preview-05-20\""},
		{variant: LlmGemini25FlashPreview0417, wire: "\"gemini-2.5-flash-preview-04-17\""},
		{variant: LlmGemini25FlashLitePreview0617, wire: "\"gemini-2.5-flash-lite-preview-06-17\""},
		{variant: LlmGemini20FlashLite001, wire: "\"gemini-2.0-flash-lite-001\""},
		{variant: LlmGemini20Flash001, wire: "\"gemini-2.0-flash-001\""},
		{variant: LlmGemini15Flash002, wire: "\"gemini-1.5-flash-002\""},
		{variant: LlmGemini15Flash001, wire: "\"gemini-1.5-flash-001\""},
		{variant: LlmGemini15Pro002, wire: "\"gemini-1.5-pro-002\""},
		{variant: LlmGemini15Pro001, wire: "\"gemini-1.5-pro-001\""},
		{variant: LlmClaudeSonnet4At20250514, wire: "\"claude-sonnet-4@20250514\""},
		{variant: LlmClaudeSonnet45At20250929, wire: "\"claude-sonnet-4-5@20250929\""},
		{variant: LlmClaudeHaiku45At20251001, wire: "\"claude-haiku-4-5@20251001\""},
		{variant: LlmClaude37SonnetAt20250219, wire: "\"claude-3-7-sonnet@20250219\""},
		{variant: LlmClaude35SonnetAt20240620, wire: "\"claude-3-5-sonnet@20240620\""},
		{variant: LlmClaude35SonnetV2At20241022, wire: "\"claude-3-5-sonnet-v2@20241022\""},
		{variant: LlmClaude3HaikuAt20240307, wire: "\"claude-3-haiku@20240307\""},
		{variant: LlmGpt520250807, wire: "\"gpt-5-2025-08-07\""},
		{variant: LlmGpt5120251113, wire: "\"gpt-5.1-2025-11-13\""},
		{variant: LlmGpt5220251211, wire: "\"gpt-5.2-2025-12-11\""},
		{variant: LlmGpt5420260305, wire: "\"gpt-5.4-2026-03-05\""},
		{variant: LlmGpt54Mini20260317, wire: "\"gpt-5.4-mini-2026-03-17\""},
		{variant: LlmGpt54Nano20260317, wire: "\"gpt-5.4-nano-2026-03-17\""},
		{variant: LlmGpt5520260423, wire: "\"gpt-5.5-2026-04-23\""},
		{variant: LlmGpt5Mini20250807, wire: "\"gpt-5-mini-2025-08-07\""},
		{variant: LlmGpt5Nano20250807, wire: "\"gpt-5-nano-2025-08-07\""},
		{variant: LlmGpt4120250414, wire: "\"gpt-4.1-2025-04-14\""},
		{variant: LlmGpt41Mini20250414, wire: "\"gpt-4.1-mini-2025-04-14\""},
		{variant: LlmGpt41Nano20250414, wire: "\"gpt-4.1-nano-2025-04-14\""},
		{variant: LlmGpt4OMini20240718, wire: "\"gpt-4o-mini-2024-07-18\""},
		{variant: LlmGpt4O20241120, wire: "\"gpt-4o-2024-11-20\""},
		{variant: LlmGpt4O20240806, wire: "\"gpt-4o-2024-08-06\""},
		{variant: LlmGpt4O20240513, wire: "\"gpt-4o-2024-05-13\""},
		{variant: LlmGpt40613, wire: "\"gpt-4-0613\""},
		{variant: LlmGpt40314, wire: "\"gpt-4-0314\""},
		{variant: LlmGpt4Turbo20240409, wire: "\"gpt-4-turbo-2024-04-09\""},
		{variant: LlmGpt35Turbo0125, wire: "\"gpt-3.5-turbo-0125\""},
		{variant: LlmGpt35Turbo1106, wire: "\"gpt-3.5-turbo-1106\""},
		{variant: LlmWattTool8B, wire: "\"watt-tool-8b\""},
		{variant: LlmWattTool70B, wire: "\"watt-tool-70b\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded Llm
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

func TestLlmVariantsRoundTrip6(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant Llm
		wire    string
	}
	cases := []enumCase{
		{variant: LlmGpt4OMini, wire: "\"gpt-4o-mini\""},
		{variant: LlmGpt4O, wire: "\"gpt-4o\""},
		{variant: LlmGpt4, wire: "\"gpt-4\""},
		{variant: LlmGpt4Turbo, wire: "\"gpt-4-turbo\""},
		{variant: LlmGpt41, wire: "\"gpt-4.1\""},
		{variant: LlmGpt41Mini, wire: "\"gpt-4.1-mini\""},
		{variant: LlmGpt41Nano, wire: "\"gpt-4.1-nano\""},
		{variant: LlmGpt5, wire: "\"gpt-5\""},
		{variant: LlmGpt51, wire: "\"gpt-5.1\""},
		{variant: LlmGpt52, wire: "\"gpt-5.2\""},
		{variant: LlmGpt52ChatLatest, wire: "\"gpt-5.2-chat-latest\""},
		{variant: LlmGpt54, wire: "\"gpt-5.4\""},
		{variant: LlmGpt54Mini, wire: "\"gpt-5.4-mini\""},
		{variant: LlmGpt54Nano, wire: "\"gpt-5.4-nano\""},
		{variant: LlmGpt55, wire: "\"gpt-5.5\""},
		{variant: LlmGpt56Sol, wire: "\"gpt-5.6-sol\""},
		{variant: LlmGpt56Terra, wire: "\"gpt-5.6-terra\""},
		{variant: LlmGpt56Luna, wire: "\"gpt-5.6-luna\""},
		{variant: LlmGpt6Astra, wire: "\"gpt-6-astra\""},
		{variant: LlmGpt6Sol, wire: "\"gpt-6-sol\""},
		{variant: LlmGpt6Luna, wire: "\"gpt-6-luna\""},
		{variant: LlmGpt61Sol, wire: "\"gpt-6.1-sol\""},
		{variant: LlmGpt61Terra, wire: "\"gpt-6.1-terra\""},
		{variant: LlmGpt5Mini, wire: "\"gpt-5-mini\""},
		{variant: LlmGpt5Nano, wire: "\"gpt-5-nano\""},
		{variant: LlmGpt35Turbo, wire: "\"gpt-3.5-turbo\""},
		{variant: LlmGemini15Pro, wire: "\"gemini-1.5-pro\""},
		{variant: LlmGemini15Flash, wire: "\"gemini-1.5-flash\""},
		{variant: LlmGemini20Flash, wire: "\"gemini-2.0-flash\""},
		{variant: LlmGemini20FlashLite, wire: "\"gemini-2.0-flash-lite\""},
		{variant: LlmGemini25FlashLite, wire: "\"gemini-2.5-flash-lite\""},
		{variant: LlmGemini25Flash, wire: "\"gemini-2.5-flash\""},
		{variant: LlmGemini3ProPreview, wire: "\"gemini-3-pro-preview\""},
		{variant: LlmGemini3FlashPreview, wire: "\"gemini-3-flash-preview\""},
		{variant: LlmGemini31ProPreview, wire: "\"gemini-3.1-pro-preview\""},
		{variant: LlmGemini31FlashLitePreview, wire: "\"gemini-3.1-flash-lite-preview\""},
		{variant: LlmGemini31FlashLite, wire: "\"gemini-3.1-flash-lite\""},
		{variant: LlmGemini35Flash, wire: "\"gemini-3.5-flash\""},
		{variant: LlmGemini35FlashLite, wire: "\"gemini-3.5-flash-lite\""},
		{variant: LlmGemini36Flash, wire: "\"gemini-3.6-flash\""},
		{variant: LlmGemini37Flash, wire: "\"gemini-3.7-flash\""},
		{variant: LlmGemini38Flash, wire: "\"gemini-3.8-flash\""},
		{variant: LlmGemini4Argon, wire: "\"gemini-4-argon\""},
		{variant: LlmGemini4Hydrogen, wire: "\"gemini-4-hydrogen\""},
		{variant: LlmClaudeSonnet45, wire: "\"claude-sonnet-4-5\""},
		{variant: LlmClaudeOpus46, wire: "\"claude-opus-4-6\""},
		{variant: LlmClaudeOpus47, wire: "\"claude-opus-4-7\""},
		{variant: LlmClaudeOpus48, wire: "\"claude-opus-4-8\""},
		{variant: LlmClaudeOpus5, wire: "\"claude-opus-5\""},
		{variant: LlmClaudeOpus55, wire: "\"claude-opus-5-5\""},
		{variant: LlmClaudeSonnet46, wire: "\"claude-sonnet-4-6\""},
		{variant: LlmClaudeSonnet5, wire: "\"claude-sonnet-5\""},
		{variant: LlmClaudeSonnet55, wire: "\"claude-sonnet-5-5\""},
		{variant: LlmClaudeSonnet4, wire: "\"claude-sonnet-4\""},
		{variant: LlmClaudeHaiku45, wire: "\"claude-haiku-4-5\""},
		{variant: LlmClaudeHaiku55, wire: "\"claude-haiku-5-5\""},
		{variant: LlmClaude37Sonnet, wire: "\"claude-3-7-sonnet\""},
		{variant: LlmClaude35Sonnet, wire: "\"claude-3-5-sonnet\""},
		{variant: LlmClaude35SonnetV1, wire: "\"claude-3-5-sonnet-v1\""},
		{variant: LlmClaude3Haiku, wire: "\"claude-3-haiku\""},
		{variant: LlmCustomLlm, wire: "\"custom-llm\""},
		{variant: LlmQwen34B, wire: "\"qwen3-4b\""},
		{variant: LlmQwen330BA3B, wire: "\"qwen3-30b-a3b\""},
		{variant: LlmQwen3635BA3B, wire: "\"qwen36-35b-a3b\""},
		{variant: LlmQwen35397BA17B, wire: "\"qwen35-397b-a17b\""},
		{variant: LlmGptOss20B, wire: "\"gpt-oss-20b\""},
		{variant: LlmGptOss120B, wire: "\"gpt-oss-120b\""},
		{variant: LlmGlm45AirFp8, wire: "\"glm-45-air-fp8\""},
		{variant: LlmGlm52, wire: "\"glm-52\""},
		{variant: LlmDeepseekV41Flash, wire: "\"deepseek-v41-flash\""},
		{variant: LlmNemotron35SuperVlNvfp4, wire: "\"nemotron-35-super-vl-nvfp4\""},
		{variant: LlmGemini25FlashPreview092025, wire: "\"gemini-2.5-flash-preview-09-2025\""},
		{variant: LlmGemini25FlashLitePreview092025, wire: "\"gemini-2.5-flash-lite-preview-09-2025\""},
		{variant: LlmGemini25FlashPreview0520, wire: "\"gemini-2.5-flash-preview-05-20\""},
		{variant: LlmGemini25FlashPreview0417, wire: "\"gemini-2.5-flash-preview-04-17\""},
		{variant: LlmGemini25FlashLitePreview0617, wire: "\"gemini-2.5-flash-lite-preview-06-17\""},
		{variant: LlmGemini20FlashLite001, wire: "\"gemini-2.0-flash-lite-001\""},
		{variant: LlmGemini20Flash001, wire: "\"gemini-2.0-flash-001\""},
		{variant: LlmGemini15Flash002, wire: "\"gemini-1.5-flash-002\""},
		{variant: LlmGemini15Flash001, wire: "\"gemini-1.5-flash-001\""},
		{variant: LlmGemini15Pro002, wire: "\"gemini-1.5-pro-002\""},
		{variant: LlmGemini15Pro001, wire: "\"gemini-1.5-pro-001\""},
		{variant: LlmClaudeSonnet4At20250514, wire: "\"claude-sonnet-4@20250514\""},
		{variant: LlmClaudeSonnet45At20250929, wire: "\"claude-sonnet-4-5@20250929\""},
		{variant: LlmClaudeHaiku45At20251001, wire: "\"claude-haiku-4-5@20251001\""},
		{variant: LlmClaude37SonnetAt20250219, wire: "\"claude-3-7-sonnet@20250219\""},
		{variant: LlmClaude35SonnetAt20240620, wire: "\"claude-3-5-sonnet@20240620\""},
		{variant: LlmClaude35SonnetV2At20241022, wire: "\"claude-3-5-sonnet-v2@20241022\""},
		{variant: LlmClaude3HaikuAt20240307, wire: "\"claude-3-haiku@20240307\""},
		{variant: LlmGpt520250807, wire: "\"gpt-5-2025-08-07\""},
		{variant: LlmGpt5120251113, wire: "\"gpt-5.1-2025-11-13\""},
		{variant: LlmGpt5220251211, wire: "\"gpt-5.2-2025-12-11\""},
		{variant: LlmGpt5420260305, wire: "\"gpt-5.4-2026-03-05\""},
		{variant: LlmGpt54Mini20260317, wire: "\"gpt-5.4-mini-2026-03-17\""},
		{variant: LlmGpt54Nano20260317, wire: "\"gpt-5.4-nano-2026-03-17\""},
		{variant: LlmGpt5520260423, wire: "\"gpt-5.5-2026-04-23\""},
		{variant: LlmGpt5Mini20250807, wire: "\"gpt-5-mini-2025-08-07\""},
		{variant: LlmGpt5Nano20250807, wire: "\"gpt-5-nano-2025-08-07\""},
		{variant: LlmGpt4120250414, wire: "\"gpt-4.1-2025-04-14\""},
		{variant: LlmGpt41Mini20250414, wire: "\"gpt-4.1-mini-2025-04-14\""},
		{variant: LlmGpt41Nano20250414, wire: "\"gpt-4.1-nano-2025-04-14\""},
		{variant: LlmGpt4OMini20240718, wire: "\"gpt-4o-mini-2024-07-18\""},
		{variant: LlmGpt4O20241120, wire: "\"gpt-4o-2024-11-20\""},
		{variant: LlmGpt4O20240806, wire: "\"gpt-4o-2024-08-06\""},
		{variant: LlmGpt4O20240513, wire: "\"gpt-4o-2024-05-13\""},
		{variant: LlmGpt40613, wire: "\"gpt-4-0613\""},
		{variant: LlmGpt40314, wire: "\"gpt-4-0314\""},
		{variant: LlmGpt4Turbo20240409, wire: "\"gpt-4-turbo-2024-04-09\""},
		{variant: LlmGpt35Turbo0125, wire: "\"gpt-3.5-turbo-0125\""},
		{variant: LlmGpt35Turbo1106, wire: "\"gpt-3.5-turbo-1106\""},
		{variant: LlmWattTool8B, wire: "\"watt-tool-8b\""},
		{variant: LlmWattTool70B, wire: "\"watt-tool-70b\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded Llm
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

func TestLlmVariantsRoundTrip7(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant Llm
		wire    string
	}
	cases := []enumCase{
		{variant: LlmGpt4OMini, wire: "\"gpt-4o-mini\""},
		{variant: LlmGpt4O, wire: "\"gpt-4o\""},
		{variant: LlmGpt4, wire: "\"gpt-4\""},
		{variant: LlmGpt4Turbo, wire: "\"gpt-4-turbo\""},
		{variant: LlmGpt41, wire: "\"gpt-4.1\""},
		{variant: LlmGpt41Mini, wire: "\"gpt-4.1-mini\""},
		{variant: LlmGpt41Nano, wire: "\"gpt-4.1-nano\""},
		{variant: LlmGpt5, wire: "\"gpt-5\""},
		{variant: LlmGpt51, wire: "\"gpt-5.1\""},
		{variant: LlmGpt52, wire: "\"gpt-5.2\""},
		{variant: LlmGpt52ChatLatest, wire: "\"gpt-5.2-chat-latest\""},
		{variant: LlmGpt54, wire: "\"gpt-5.4\""},
		{variant: LlmGpt54Mini, wire: "\"gpt-5.4-mini\""},
		{variant: LlmGpt54Nano, wire: "\"gpt-5.4-nano\""},
		{variant: LlmGpt55, wire: "\"gpt-5.5\""},
		{variant: LlmGpt56Sol, wire: "\"gpt-5.6-sol\""},
		{variant: LlmGpt56Terra, wire: "\"gpt-5.6-terra\""},
		{variant: LlmGpt56Luna, wire: "\"gpt-5.6-luna\""},
		{variant: LlmGpt6Astra, wire: "\"gpt-6-astra\""},
		{variant: LlmGpt6Sol, wire: "\"gpt-6-sol\""},
		{variant: LlmGpt6Luna, wire: "\"gpt-6-luna\""},
		{variant: LlmGpt61Sol, wire: "\"gpt-6.1-sol\""},
		{variant: LlmGpt61Terra, wire: "\"gpt-6.1-terra\""},
		{variant: LlmGpt5Mini, wire: "\"gpt-5-mini\""},
		{variant: LlmGpt5Nano, wire: "\"gpt-5-nano\""},
		{variant: LlmGpt35Turbo, wire: "\"gpt-3.5-turbo\""},
		{variant: LlmGemini15Pro, wire: "\"gemini-1.5-pro\""},
		{variant: LlmGemini15Flash, wire: "\"gemini-1.5-flash\""},
		{variant: LlmGemini20Flash, wire: "\"gemini-2.0-flash\""},
		{variant: LlmGemini20FlashLite, wire: "\"gemini-2.0-flash-lite\""},
		{variant: LlmGemini25FlashLite, wire: "\"gemini-2.5-flash-lite\""},
		{variant: LlmGemini25Flash, wire: "\"gemini-2.5-flash\""},
		{variant: LlmGemini3ProPreview, wire: "\"gemini-3-pro-preview\""},
		{variant: LlmGemini3FlashPreview, wire: "\"gemini-3-flash-preview\""},
		{variant: LlmGemini31ProPreview, wire: "\"gemini-3.1-pro-preview\""},
		{variant: LlmGemini31FlashLitePreview, wire: "\"gemini-3.1-flash-lite-preview\""},
		{variant: LlmGemini31FlashLite, wire: "\"gemini-3.1-flash-lite\""},
		{variant: LlmGemini35Flash, wire: "\"gemini-3.5-flash\""},
		{variant: LlmGemini35FlashLite, wire: "\"gemini-3.5-flash-lite\""},
		{variant: LlmGemini36Flash, wire: "\"gemini-3.6-flash\""},
		{variant: LlmGemini37Flash, wire: "\"gemini-3.7-flash\""},
		{variant: LlmGemini38Flash, wire: "\"gemini-3.8-flash\""},
		{variant: LlmGemini4Argon, wire: "\"gemini-4-argon\""},
		{variant: LlmGemini4Hydrogen, wire: "\"gemini-4-hydrogen\""},
		{variant: LlmClaudeSonnet45, wire: "\"claude-sonnet-4-5\""},
		{variant: LlmClaudeOpus46, wire: "\"claude-opus-4-6\""},
		{variant: LlmClaudeOpus47, wire: "\"claude-opus-4-7\""},
		{variant: LlmClaudeOpus48, wire: "\"claude-opus-4-8\""},
		{variant: LlmClaudeOpus5, wire: "\"claude-opus-5\""},
		{variant: LlmClaudeOpus55, wire: "\"claude-opus-5-5\""},
		{variant: LlmClaudeSonnet46, wire: "\"claude-sonnet-4-6\""},
		{variant: LlmClaudeSonnet5, wire: "\"claude-sonnet-5\""},
		{variant: LlmClaudeSonnet55, wire: "\"claude-sonnet-5-5\""},
		{variant: LlmClaudeSonnet4, wire: "\"claude-sonnet-4\""},
		{variant: LlmClaudeHaiku45, wire: "\"claude-haiku-4-5\""},
		{variant: LlmClaudeHaiku55, wire: "\"claude-haiku-5-5\""},
		{variant: LlmClaude37Sonnet, wire: "\"claude-3-7-sonnet\""},
		{variant: LlmClaude35Sonnet, wire: "\"claude-3-5-sonnet\""},
		{variant: LlmClaude35SonnetV1, wire: "\"claude-3-5-sonnet-v1\""},
		{variant: LlmClaude3Haiku, wire: "\"claude-3-haiku\""},
		{variant: LlmCustomLlm, wire: "\"custom-llm\""},
		{variant: LlmQwen34B, wire: "\"qwen3-4b\""},
		{variant: LlmQwen330BA3B, wire: "\"qwen3-30b-a3b\""},
		{variant: LlmQwen3635BA3B, wire: "\"qwen36-35b-a3b\""},
		{variant: LlmQwen35397BA17B, wire: "\"qwen35-397b-a17b\""},
		{variant: LlmGptOss20B, wire: "\"gpt-oss-20b\""},
		{variant: LlmGptOss120B, wire: "\"gpt-oss-120b\""},
		{variant: LlmGlm45AirFp8, wire: "\"glm-45-air-fp8\""},
		{variant: LlmGlm52, wire: "\"glm-52\""},
		{variant: LlmDeepseekV41Flash, wire: "\"deepseek-v41-flash\""},
		{variant: LlmNemotron35SuperVlNvfp4, wire: "\"nemotron-35-super-vl-nvfp4\""},
		{variant: LlmGemini25FlashPreview092025, wire: "\"gemini-2.5-flash-preview-09-2025\""},
		{variant: LlmGemini25FlashLitePreview092025, wire: "\"gemini-2.5-flash-lite-preview-09-2025\""},
		{variant: LlmGemini25FlashPreview0520, wire: "\"gemini-2.5-flash-preview-05-20\""},
		{variant: LlmGemini25FlashPreview0417, wire: "\"gemini-2.5-flash-preview-04-17\""},
		{variant: LlmGemini25FlashLitePreview0617, wire: "\"gemini-2.5-flash-lite-preview-06-17\""},
		{variant: LlmGemini20FlashLite001, wire: "\"gemini-2.0-flash-lite-001\""},
		{variant: LlmGemini20Flash001, wire: "\"gemini-2.0-flash-001\""},
		{variant: LlmGemini15Flash002, wire: "\"gemini-1.5-flash-002\""},
		{variant: LlmGemini15Flash001, wire: "\"gemini-1.5-flash-001\""},
		{variant: LlmGemini15Pro002, wire: "\"gemini-1.5-pro-002\""},
		{variant: LlmGemini15Pro001, wire: "\"gemini-1.5-pro-001\""},
		{variant: LlmClaudeSonnet4At20250514, wire: "\"claude-sonnet-4@20250514\""},
		{variant: LlmClaudeSonnet45At20250929, wire: "\"claude-sonnet-4-5@20250929\""},
		{variant: LlmClaudeHaiku45At20251001, wire: "\"claude-haiku-4-5@20251001\""},
		{variant: LlmClaude37SonnetAt20250219, wire: "\"claude-3-7-sonnet@20250219\""},
		{variant: LlmClaude35SonnetAt20240620, wire: "\"claude-3-5-sonnet@20240620\""},
		{variant: LlmClaude35SonnetV2At20241022, wire: "\"claude-3-5-sonnet-v2@20241022\""},
		{variant: LlmClaude3HaikuAt20240307, wire: "\"claude-3-haiku@20240307\""},
		{variant: LlmGpt520250807, wire: "\"gpt-5-2025-08-07\""},
		{variant: LlmGpt5120251113, wire: "\"gpt-5.1-2025-11-13\""},
		{variant: LlmGpt5220251211, wire: "\"gpt-5.2-2025-12-11\""},
		{variant: LlmGpt5420260305, wire: "\"gpt-5.4-2026-03-05\""},
		{variant: LlmGpt54Mini20260317, wire: "\"gpt-5.4-mini-2026-03-17\""},
		{variant: LlmGpt54Nano20260317, wire: "\"gpt-5.4-nano-2026-03-17\""},
		{variant: LlmGpt5520260423, wire: "\"gpt-5.5-2026-04-23\""},
		{variant: LlmGpt5Mini20250807, wire: "\"gpt-5-mini-2025-08-07\""},
		{variant: LlmGpt5Nano20250807, wire: "\"gpt-5-nano-2025-08-07\""},
		{variant: LlmGpt4120250414, wire: "\"gpt-4.1-2025-04-14\""},
		{variant: LlmGpt41Mini20250414, wire: "\"gpt-4.1-mini-2025-04-14\""},
		{variant: LlmGpt41Nano20250414, wire: "\"gpt-4.1-nano-2025-04-14\""},
		{variant: LlmGpt4OMini20240718, wire: "\"gpt-4o-mini-2024-07-18\""},
		{variant: LlmGpt4O20241120, wire: "\"gpt-4o-2024-11-20\""},
		{variant: LlmGpt4O20240806, wire: "\"gpt-4o-2024-08-06\""},
		{variant: LlmGpt4O20240513, wire: "\"gpt-4o-2024-05-13\""},
		{variant: LlmGpt40613, wire: "\"gpt-4-0613\""},
		{variant: LlmGpt40314, wire: "\"gpt-4-0314\""},
		{variant: LlmGpt4Turbo20240409, wire: "\"gpt-4-turbo-2024-04-09\""},
		{variant: LlmGpt35Turbo0125, wire: "\"gpt-3.5-turbo-0125\""},
		{variant: LlmGpt35Turbo1106, wire: "\"gpt-3.5-turbo-1106\""},
		{variant: LlmWattTool8B, wire: "\"watt-tool-8b\""},
		{variant: LlmWattTool70B, wire: "\"watt-tool-70b\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded Llm
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

func TestLlmVariantsRoundTrip8(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant Llm
		wire    string
	}
	cases := []enumCase{
		{variant: LlmGpt4OMini, wire: "\"gpt-4o-mini\""},
		{variant: LlmGpt4O, wire: "\"gpt-4o\""},
		{variant: LlmGpt4, wire: "\"gpt-4\""},
		{variant: LlmGpt4Turbo, wire: "\"gpt-4-turbo\""},
		{variant: LlmGpt41, wire: "\"gpt-4.1\""},
		{variant: LlmGpt41Mini, wire: "\"gpt-4.1-mini\""},
		{variant: LlmGpt41Nano, wire: "\"gpt-4.1-nano\""},
		{variant: LlmGpt5, wire: "\"gpt-5\""},
		{variant: LlmGpt51, wire: "\"gpt-5.1\""},
		{variant: LlmGpt52, wire: "\"gpt-5.2\""},
		{variant: LlmGpt52ChatLatest, wire: "\"gpt-5.2-chat-latest\""},
		{variant: LlmGpt54, wire: "\"gpt-5.4\""},
		{variant: LlmGpt54Mini, wire: "\"gpt-5.4-mini\""},
		{variant: LlmGpt54Nano, wire: "\"gpt-5.4-nano\""},
		{variant: LlmGpt55, wire: "\"gpt-5.5\""},
		{variant: LlmGpt56Sol, wire: "\"gpt-5.6-sol\""},
		{variant: LlmGpt56Terra, wire: "\"gpt-5.6-terra\""},
		{variant: LlmGpt56Luna, wire: "\"gpt-5.6-luna\""},
		{variant: LlmGpt6Astra, wire: "\"gpt-6-astra\""},
		{variant: LlmGpt6Sol, wire: "\"gpt-6-sol\""},
		{variant: LlmGpt6Luna, wire: "\"gpt-6-luna\""},
		{variant: LlmGpt61Sol, wire: "\"gpt-6.1-sol\""},
		{variant: LlmGpt61Terra, wire: "\"gpt-6.1-terra\""},
		{variant: LlmGpt5Mini, wire: "\"gpt-5-mini\""},
		{variant: LlmGpt5Nano, wire: "\"gpt-5-nano\""},
		{variant: LlmGpt35Turbo, wire: "\"gpt-3.5-turbo\""},
		{variant: LlmGemini15Pro, wire: "\"gemini-1.5-pro\""},
		{variant: LlmGemini15Flash, wire: "\"gemini-1.5-flash\""},
		{variant: LlmGemini20Flash, wire: "\"gemini-2.0-flash\""},
		{variant: LlmGemini20FlashLite, wire: "\"gemini-2.0-flash-lite\""},
		{variant: LlmGemini25FlashLite, wire: "\"gemini-2.5-flash-lite\""},
		{variant: LlmGemini25Flash, wire: "\"gemini-2.5-flash\""},
		{variant: LlmGemini3ProPreview, wire: "\"gemini-3-pro-preview\""},
		{variant: LlmGemini3FlashPreview, wire: "\"gemini-3-flash-preview\""},
		{variant: LlmGemini31ProPreview, wire: "\"gemini-3.1-pro-preview\""},
		{variant: LlmGemini31FlashLitePreview, wire: "\"gemini-3.1-flash-lite-preview\""},
		{variant: LlmGemini31FlashLite, wire: "\"gemini-3.1-flash-lite\""},
		{variant: LlmGemini35Flash, wire: "\"gemini-3.5-flash\""},
		{variant: LlmGemini35FlashLite, wire: "\"gemini-3.5-flash-lite\""},
		{variant: LlmGemini36Flash, wire: "\"gemini-3.6-flash\""},
		{variant: LlmGemini37Flash, wire: "\"gemini-3.7-flash\""},
		{variant: LlmGemini38Flash, wire: "\"gemini-3.8-flash\""},
		{variant: LlmGemini4Argon, wire: "\"gemini-4-argon\""},
		{variant: LlmGemini4Hydrogen, wire: "\"gemini-4-hydrogen\""},
		{variant: LlmClaudeSonnet45, wire: "\"claude-sonnet-4-5\""},
		{variant: LlmClaudeOpus46, wire: "\"claude-opus-4-6\""},
		{variant: LlmClaudeOpus47, wire: "\"claude-opus-4-7\""},
		{variant: LlmClaudeOpus48, wire: "\"claude-opus-4-8\""},
		{variant: LlmClaudeOpus5, wire: "\"claude-opus-5\""},
		{variant: LlmClaudeOpus55, wire: "\"claude-opus-5-5\""},
		{variant: LlmClaudeSonnet46, wire: "\"claude-sonnet-4-6\""},
		{variant: LlmClaudeSonnet5, wire: "\"claude-sonnet-5\""},
		{variant: LlmClaudeSonnet55, wire: "\"claude-sonnet-5-5\""},
		{variant: LlmClaudeSonnet4, wire: "\"claude-sonnet-4\""},
		{variant: LlmClaudeHaiku45, wire: "\"claude-haiku-4-5\""},
		{variant: LlmClaudeHaiku55, wire: "\"claude-haiku-5-5\""},
		{variant: LlmClaude37Sonnet, wire: "\"claude-3-7-sonnet\""},
		{variant: LlmClaude35Sonnet, wire: "\"claude-3-5-sonnet\""},
		{variant: LlmClaude35SonnetV1, wire: "\"claude-3-5-sonnet-v1\""},
		{variant: LlmClaude3Haiku, wire: "\"claude-3-haiku\""},
		{variant: LlmCustomLlm, wire: "\"custom-llm\""},
		{variant: LlmQwen34B, wire: "\"qwen3-4b\""},
		{variant: LlmQwen330BA3B, wire: "\"qwen3-30b-a3b\""},
		{variant: LlmQwen3635BA3B, wire: "\"qwen36-35b-a3b\""},
		{variant: LlmQwen35397BA17B, wire: "\"qwen35-397b-a17b\""},
		{variant: LlmGptOss20B, wire: "\"gpt-oss-20b\""},
		{variant: LlmGptOss120B, wire: "\"gpt-oss-120b\""},
		{variant: LlmGlm45AirFp8, wire: "\"glm-45-air-fp8\""},
		{variant: LlmGlm52, wire: "\"glm-52\""},
		{variant: LlmDeepseekV41Flash, wire: "\"deepseek-v41-flash\""},
		{variant: LlmNemotron35SuperVlNvfp4, wire: "\"nemotron-35-super-vl-nvfp4\""},
		{variant: LlmGemini25FlashPreview092025, wire: "\"gemini-2.5-flash-preview-09-2025\""},
		{variant: LlmGemini25FlashLitePreview092025, wire: "\"gemini-2.5-flash-lite-preview-09-2025\""},
		{variant: LlmGemini25FlashPreview0520, wire: "\"gemini-2.5-flash-preview-05-20\""},
		{variant: LlmGemini25FlashPreview0417, wire: "\"gemini-2.5-flash-preview-04-17\""},
		{variant: LlmGemini25FlashLitePreview0617, wire: "\"gemini-2.5-flash-lite-preview-06-17\""},
		{variant: LlmGemini20FlashLite001, wire: "\"gemini-2.0-flash-lite-001\""},
		{variant: LlmGemini20Flash001, wire: "\"gemini-2.0-flash-001\""},
		{variant: LlmGemini15Flash002, wire: "\"gemini-1.5-flash-002\""},
		{variant: LlmGemini15Flash001, wire: "\"gemini-1.5-flash-001\""},
		{variant: LlmGemini15Pro002, wire: "\"gemini-1.5-pro-002\""},
		{variant: LlmGemini15Pro001, wire: "\"gemini-1.5-pro-001\""},
		{variant: LlmClaudeSonnet4At20250514, wire: "\"claude-sonnet-4@20250514\""},
		{variant: LlmClaudeSonnet45At20250929, wire: "\"claude-sonnet-4-5@20250929\""},
		{variant: LlmClaudeHaiku45At20251001, wire: "\"claude-haiku-4-5@20251001\""},
		{variant: LlmClaude37SonnetAt20250219, wire: "\"claude-3-7-sonnet@20250219\""},
		{variant: LlmClaude35SonnetAt20240620, wire: "\"claude-3-5-sonnet@20240620\""},
		{variant: LlmClaude35SonnetV2At20241022, wire: "\"claude-3-5-sonnet-v2@20241022\""},
		{variant: LlmClaude3HaikuAt20240307, wire: "\"claude-3-haiku@20240307\""},
		{variant: LlmGpt520250807, wire: "\"gpt-5-2025-08-07\""},
		{variant: LlmGpt5120251113, wire: "\"gpt-5.1-2025-11-13\""},
		{variant: LlmGpt5220251211, wire: "\"gpt-5.2-2025-12-11\""},
		{variant: LlmGpt5420260305, wire: "\"gpt-5.4-2026-03-05\""},
		{variant: LlmGpt54Mini20260317, wire: "\"gpt-5.4-mini-2026-03-17\""},
		{variant: LlmGpt54Nano20260317, wire: "\"gpt-5.4-nano-2026-03-17\""},
		{variant: LlmGpt5520260423, wire: "\"gpt-5.5-2026-04-23\""},
		{variant: LlmGpt5Mini20250807, wire: "\"gpt-5-mini-2025-08-07\""},
		{variant: LlmGpt5Nano20250807, wire: "\"gpt-5-nano-2025-08-07\""},
		{variant: LlmGpt4120250414, wire: "\"gpt-4.1-2025-04-14\""},
		{variant: LlmGpt41Mini20250414, wire: "\"gpt-4.1-mini-2025-04-14\""},
		{variant: LlmGpt41Nano20250414, wire: "\"gpt-4.1-nano-2025-04-14\""},
		{variant: LlmGpt4OMini20240718, wire: "\"gpt-4o-mini-2024-07-18\""},
		{variant: LlmGpt4O20241120, wire: "\"gpt-4o-2024-11-20\""},
		{variant: LlmGpt4O20240806, wire: "\"gpt-4o-2024-08-06\""},
		{variant: LlmGpt4O20240513, wire: "\"gpt-4o-2024-05-13\""},
		{variant: LlmGpt40613, wire: "\"gpt-4-0613\""},
		{variant: LlmGpt40314, wire: "\"gpt-4-0314\""},
		{variant: LlmGpt4Turbo20240409, wire: "\"gpt-4-turbo-2024-04-09\""},
		{variant: LlmGpt35Turbo0125, wire: "\"gpt-3.5-turbo-0125\""},
		{variant: LlmGpt35Turbo1106, wire: "\"gpt-3.5-turbo-1106\""},
		{variant: LlmWattTool8B, wire: "\"watt-tool-8b\""},
		{variant: LlmWattTool70B, wire: "\"watt-tool-70b\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded Llm
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

func TestLlmVariantsRoundTrip9(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant Llm
		wire    string
	}
	cases := []enumCase{
		{variant: LlmGpt4OMini, wire: "\"gpt-4o-mini\""},
		{variant: LlmGpt4O, wire: "\"gpt-4o\""},
		{variant: LlmGpt4, wire: "\"gpt-4\""},
		{variant: LlmGpt4Turbo, wire: "\"gpt-4-turbo\""},
		{variant: LlmGpt41, wire: "\"gpt-4.1\""},
		{variant: LlmGpt41Mini, wire: "\"gpt-4.1-mini\""},
		{variant: LlmGpt41Nano, wire: "\"gpt-4.1-nano\""},
		{variant: LlmGpt5, wire: "\"gpt-5\""},
		{variant: LlmGpt51, wire: "\"gpt-5.1\""},
		{variant: LlmGpt52, wire: "\"gpt-5.2\""},
		{variant: LlmGpt52ChatLatest, wire: "\"gpt-5.2-chat-latest\""},
		{variant: LlmGpt54, wire: "\"gpt-5.4\""},
		{variant: LlmGpt54Mini, wire: "\"gpt-5.4-mini\""},
		{variant: LlmGpt54Nano, wire: "\"gpt-5.4-nano\""},
		{variant: LlmGpt55, wire: "\"gpt-5.5\""},
		{variant: LlmGpt56Sol, wire: "\"gpt-5.6-sol\""},
		{variant: LlmGpt56Terra, wire: "\"gpt-5.6-terra\""},
		{variant: LlmGpt56Luna, wire: "\"gpt-5.6-luna\""},
		{variant: LlmGpt6Astra, wire: "\"gpt-6-astra\""},
		{variant: LlmGpt6Sol, wire: "\"gpt-6-sol\""},
		{variant: LlmGpt6Luna, wire: "\"gpt-6-luna\""},
		{variant: LlmGpt61Sol, wire: "\"gpt-6.1-sol\""},
		{variant: LlmGpt61Terra, wire: "\"gpt-6.1-terra\""},
		{variant: LlmGpt5Mini, wire: "\"gpt-5-mini\""},
		{variant: LlmGpt5Nano, wire: "\"gpt-5-nano\""},
		{variant: LlmGpt35Turbo, wire: "\"gpt-3.5-turbo\""},
		{variant: LlmGemini15Pro, wire: "\"gemini-1.5-pro\""},
		{variant: LlmGemini15Flash, wire: "\"gemini-1.5-flash\""},
		{variant: LlmGemini20Flash, wire: "\"gemini-2.0-flash\""},
		{variant: LlmGemini20FlashLite, wire: "\"gemini-2.0-flash-lite\""},
		{variant: LlmGemini25FlashLite, wire: "\"gemini-2.5-flash-lite\""},
		{variant: LlmGemini25Flash, wire: "\"gemini-2.5-flash\""},
		{variant: LlmGemini3ProPreview, wire: "\"gemini-3-pro-preview\""},
		{variant: LlmGemini3FlashPreview, wire: "\"gemini-3-flash-preview\""},
		{variant: LlmGemini31ProPreview, wire: "\"gemini-3.1-pro-preview\""},
		{variant: LlmGemini31FlashLitePreview, wire: "\"gemini-3.1-flash-lite-preview\""},
		{variant: LlmGemini31FlashLite, wire: "\"gemini-3.1-flash-lite\""},
		{variant: LlmGemini35Flash, wire: "\"gemini-3.5-flash\""},
		{variant: LlmGemini35FlashLite, wire: "\"gemini-3.5-flash-lite\""},
		{variant: LlmGemini36Flash, wire: "\"gemini-3.6-flash\""},
		{variant: LlmGemini37Flash, wire: "\"gemini-3.7-flash\""},
		{variant: LlmGemini38Flash, wire: "\"gemini-3.8-flash\""},
		{variant: LlmGemini4Argon, wire: "\"gemini-4-argon\""},
		{variant: LlmGemini4Hydrogen, wire: "\"gemini-4-hydrogen\""},
		{variant: LlmClaudeSonnet45, wire: "\"claude-sonnet-4-5\""},
		{variant: LlmClaudeOpus46, wire: "\"claude-opus-4-6\""},
		{variant: LlmClaudeOpus47, wire: "\"claude-opus-4-7\""},
		{variant: LlmClaudeOpus48, wire: "\"claude-opus-4-8\""},
		{variant: LlmClaudeOpus5, wire: "\"claude-opus-5\""},
		{variant: LlmClaudeOpus55, wire: "\"claude-opus-5-5\""},
		{variant: LlmClaudeSonnet46, wire: "\"claude-sonnet-4-6\""},
		{variant: LlmClaudeSonnet5, wire: "\"claude-sonnet-5\""},
		{variant: LlmClaudeSonnet55, wire: "\"claude-sonnet-5-5\""},
		{variant: LlmClaudeSonnet4, wire: "\"claude-sonnet-4\""},
		{variant: LlmClaudeHaiku45, wire: "\"claude-haiku-4-5\""},
		{variant: LlmClaudeHaiku55, wire: "\"claude-haiku-5-5\""},
		{variant: LlmClaude37Sonnet, wire: "\"claude-3-7-sonnet\""},
		{variant: LlmClaude35Sonnet, wire: "\"claude-3-5-sonnet\""},
		{variant: LlmClaude35SonnetV1, wire: "\"claude-3-5-sonnet-v1\""},
		{variant: LlmClaude3Haiku, wire: "\"claude-3-haiku\""},
		{variant: LlmCustomLlm, wire: "\"custom-llm\""},
		{variant: LlmQwen34B, wire: "\"qwen3-4b\""},
		{variant: LlmQwen330BA3B, wire: "\"qwen3-30b-a3b\""},
		{variant: LlmQwen3635BA3B, wire: "\"qwen36-35b-a3b\""},
		{variant: LlmQwen35397BA17B, wire: "\"qwen35-397b-a17b\""},
		{variant: LlmGptOss20B, wire: "\"gpt-oss-20b\""},
		{variant: LlmGptOss120B, wire: "\"gpt-oss-120b\""},
		{variant: LlmGlm45AirFp8, wire: "\"glm-45-air-fp8\""},
		{variant: LlmGlm52, wire: "\"glm-52\""},
		{variant: LlmDeepseekV41Flash, wire: "\"deepseek-v41-flash\""},
		{variant: LlmNemotron35SuperVlNvfp4, wire: "\"nemotron-35-super-vl-nvfp4\""},
		{variant: LlmGemini25FlashPreview092025, wire: "\"gemini-2.5-flash-preview-09-2025\""},
		{variant: LlmGemini25FlashLitePreview092025, wire: "\"gemini-2.5-flash-lite-preview-09-2025\""},
		{variant: LlmGemini25FlashPreview0520, wire: "\"gemini-2.5-flash-preview-05-20\""},
		{variant: LlmGemini25FlashPreview0417, wire: "\"gemini-2.5-flash-preview-04-17\""},
		{variant: LlmGemini25FlashLitePreview0617, wire: "\"gemini-2.5-flash-lite-preview-06-17\""},
		{variant: LlmGemini20FlashLite001, wire: "\"gemini-2.0-flash-lite-001\""},
		{variant: LlmGemini20Flash001, wire: "\"gemini-2.0-flash-001\""},
		{variant: LlmGemini15Flash002, wire: "\"gemini-1.5-flash-002\""},
		{variant: LlmGemini15Flash001, wire: "\"gemini-1.5-flash-001\""},
		{variant: LlmGemini15Pro002, wire: "\"gemini-1.5-pro-002\""},
		{variant: LlmGemini15Pro001, wire: "\"gemini-1.5-pro-001\""},
		{variant: LlmClaudeSonnet4At20250514, wire: "\"claude-sonnet-4@20250514\""},
		{variant: LlmClaudeSonnet45At20250929, wire: "\"claude-sonnet-4-5@20250929\""},
		{variant: LlmClaudeHaiku45At20251001, wire: "\"claude-haiku-4-5@20251001\""},
		{variant: LlmClaude37SonnetAt20250219, wire: "\"claude-3-7-sonnet@20250219\""},
		{variant: LlmClaude35SonnetAt20240620, wire: "\"claude-3-5-sonnet@20240620\""},
		{variant: LlmClaude35SonnetV2At20241022, wire: "\"claude-3-5-sonnet-v2@20241022\""},
		{variant: LlmClaude3HaikuAt20240307, wire: "\"claude-3-haiku@20240307\""},
		{variant: LlmGpt520250807, wire: "\"gpt-5-2025-08-07\""},
		{variant: LlmGpt5120251113, wire: "\"gpt-5.1-2025-11-13\""},
		{variant: LlmGpt5220251211, wire: "\"gpt-5.2-2025-12-11\""},
		{variant: LlmGpt5420260305, wire: "\"gpt-5.4-2026-03-05\""},
		{variant: LlmGpt54Mini20260317, wire: "\"gpt-5.4-mini-2026-03-17\""},
		{variant: LlmGpt54Nano20260317, wire: "\"gpt-5.4-nano-2026-03-17\""},
		{variant: LlmGpt5520260423, wire: "\"gpt-5.5-2026-04-23\""},
		{variant: LlmGpt5Mini20250807, wire: "\"gpt-5-mini-2025-08-07\""},
		{variant: LlmGpt5Nano20250807, wire: "\"gpt-5-nano-2025-08-07\""},
		{variant: LlmGpt4120250414, wire: "\"gpt-4.1-2025-04-14\""},
		{variant: LlmGpt41Mini20250414, wire: "\"gpt-4.1-mini-2025-04-14\""},
		{variant: LlmGpt41Nano20250414, wire: "\"gpt-4.1-nano-2025-04-14\""},
		{variant: LlmGpt4OMini20240718, wire: "\"gpt-4o-mini-2024-07-18\""},
		{variant: LlmGpt4O20241120, wire: "\"gpt-4o-2024-11-20\""},
		{variant: LlmGpt4O20240806, wire: "\"gpt-4o-2024-08-06\""},
		{variant: LlmGpt4O20240513, wire: "\"gpt-4o-2024-05-13\""},
		{variant: LlmGpt40613, wire: "\"gpt-4-0613\""},
		{variant: LlmGpt40314, wire: "\"gpt-4-0314\""},
		{variant: LlmGpt4Turbo20240409, wire: "\"gpt-4-turbo-2024-04-09\""},
		{variant: LlmGpt35Turbo0125, wire: "\"gpt-3.5-turbo-0125\""},
		{variant: LlmGpt35Turbo1106, wire: "\"gpt-3.5-turbo-1106\""},
		{variant: LlmWattTool8B, wire: "\"watt-tool-8b\""},
		{variant: LlmWattTool70B, wire: "\"watt-tool-70b\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded Llm
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
