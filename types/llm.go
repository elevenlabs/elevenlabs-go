package types

import (
	"fmt"
)

type Llm string

const (
	LlmGpt4OMini                      Llm = "gpt-4o-mini"
	LlmGpt4O                          Llm = "gpt-4o"
	LlmGpt4                           Llm = "gpt-4"
	LlmGpt4Turbo                      Llm = "gpt-4-turbo"
	LlmGpt41                          Llm = "gpt-4.1"
	LlmGpt41Mini                      Llm = "gpt-4.1-mini"
	LlmGpt41Nano                      Llm = "gpt-4.1-nano"
	LlmGpt5                           Llm = "gpt-5"
	LlmGpt51                          Llm = "gpt-5.1"
	LlmGpt52                          Llm = "gpt-5.2"
	LlmGpt52ChatLatest                Llm = "gpt-5.2-chat-latest"
	LlmGpt54                          Llm = "gpt-5.4"
	LlmGpt54Mini                      Llm = "gpt-5.4-mini"
	LlmGpt54Nano                      Llm = "gpt-5.4-nano"
	LlmGpt55                          Llm = "gpt-5.5"
	LlmGpt56Sol                       Llm = "gpt-5.6-sol"
	LlmGpt56Terra                     Llm = "gpt-5.6-terra"
	LlmGpt56Luna                      Llm = "gpt-5.6-luna"
	LlmGpt6Astra                      Llm = "gpt-6-astra"
	LlmGpt6Sol                        Llm = "gpt-6-sol"
	LlmGpt6Luna                       Llm = "gpt-6-luna"
	LlmGpt61Sol                       Llm = "gpt-6.1-sol"
	LlmGpt61Terra                     Llm = "gpt-6.1-terra"
	LlmGpt5Mini                       Llm = "gpt-5-mini"
	LlmGpt5Nano                       Llm = "gpt-5-nano"
	LlmGpt35Turbo                     Llm = "gpt-3.5-turbo"
	LlmGemini15Pro                    Llm = "gemini-1.5-pro"
	LlmGemini15Flash                  Llm = "gemini-1.5-flash"
	LlmGemini20Flash                  Llm = "gemini-2.0-flash"
	LlmGemini20FlashLite              Llm = "gemini-2.0-flash-lite"
	LlmGemini25FlashLite              Llm = "gemini-2.5-flash-lite"
	LlmGemini25Flash                  Llm = "gemini-2.5-flash"
	LlmGemini3ProPreview              Llm = "gemini-3-pro-preview"
	LlmGemini3FlashPreview            Llm = "gemini-3-flash-preview"
	LlmGemini31ProPreview             Llm = "gemini-3.1-pro-preview"
	LlmGemini31FlashLitePreview       Llm = "gemini-3.1-flash-lite-preview"
	LlmGemini31FlashLite              Llm = "gemini-3.1-flash-lite"
	LlmGemini35Flash                  Llm = "gemini-3.5-flash"
	LlmGemini35FlashLite              Llm = "gemini-3.5-flash-lite"
	LlmGemini36Flash                  Llm = "gemini-3.6-flash"
	LlmGemini37Flash                  Llm = "gemini-3.7-flash"
	LlmGemini38Flash                  Llm = "gemini-3.8-flash"
	LlmGemini4Argon                   Llm = "gemini-4-argon"
	LlmGemini4Hydrogen                Llm = "gemini-4-hydrogen"
	LlmClaudeSonnet45                 Llm = "claude-sonnet-4-5"
	LlmClaudeOpus46                   Llm = "claude-opus-4-6"
	LlmClaudeOpus47                   Llm = "claude-opus-4-7"
	LlmClaudeOpus48                   Llm = "claude-opus-4-8"
	LlmClaudeOpus5                    Llm = "claude-opus-5"
	LlmClaudeOpus55                   Llm = "claude-opus-5-5"
	LlmClaudeSonnet46                 Llm = "claude-sonnet-4-6"
	LlmClaudeSonnet5                  Llm = "claude-sonnet-5"
	LlmClaudeSonnet55                 Llm = "claude-sonnet-5-5"
	LlmClaudeSonnet4                  Llm = "claude-sonnet-4"
	LlmClaudeHaiku45                  Llm = "claude-haiku-4-5"
	LlmClaudeHaiku55                  Llm = "claude-haiku-5-5"
	LlmClaude37Sonnet                 Llm = "claude-3-7-sonnet"
	LlmClaude35Sonnet                 Llm = "claude-3-5-sonnet"
	LlmClaude35SonnetV1               Llm = "claude-3-5-sonnet-v1"
	LlmClaude3Haiku                   Llm = "claude-3-haiku"
	LlmCustomLlm                      Llm = "custom-llm"
	LlmQwen34B                        Llm = "qwen3-4b"
	LlmQwen330BA3B                    Llm = "qwen3-30b-a3b"
	LlmQwen3635BA3B                   Llm = "qwen36-35b-a3b"
	LlmQwen35397BA17B                 Llm = "qwen35-397b-a17b"
	LlmGptOss20B                      Llm = "gpt-oss-20b"
	LlmGptOss120B                     Llm = "gpt-oss-120b"
	LlmGlm45AirFp8                    Llm = "glm-45-air-fp8"
	LlmGlm52                          Llm = "glm-52"
	LlmDeepseekV41Flash               Llm = "deepseek-v41-flash"
	LlmNemotron35SuperVlNvfp4         Llm = "nemotron-35-super-vl-nvfp4"
	LlmGemini25FlashPreview092025     Llm = "gemini-2.5-flash-preview-09-2025"
	LlmGemini25FlashLitePreview092025 Llm = "gemini-2.5-flash-lite-preview-09-2025"
	LlmGemini25FlashPreview0520       Llm = "gemini-2.5-flash-preview-05-20"
	LlmGemini25FlashPreview0417       Llm = "gemini-2.5-flash-preview-04-17"
	LlmGemini25FlashLitePreview0617   Llm = "gemini-2.5-flash-lite-preview-06-17"
	LlmGemini20FlashLite001           Llm = "gemini-2.0-flash-lite-001"
	LlmGemini20Flash001               Llm = "gemini-2.0-flash-001"
	LlmGemini15Flash002               Llm = "gemini-1.5-flash-002"
	LlmGemini15Flash001               Llm = "gemini-1.5-flash-001"
	LlmGemini15Pro002                 Llm = "gemini-1.5-pro-002"
	LlmGemini15Pro001                 Llm = "gemini-1.5-pro-001"
	LlmClaudeSonnet4At20250514        Llm = "claude-sonnet-4@20250514"
	LlmClaudeSonnet45At20250929       Llm = "claude-sonnet-4-5@20250929"
	LlmClaudeHaiku45At20251001        Llm = "claude-haiku-4-5@20251001"
	LlmClaude37SonnetAt20250219       Llm = "claude-3-7-sonnet@20250219"
	LlmClaude35SonnetAt20240620       Llm = "claude-3-5-sonnet@20240620"
	LlmClaude35SonnetV2At20241022     Llm = "claude-3-5-sonnet-v2@20241022"
	LlmClaude3HaikuAt20240307         Llm = "claude-3-haiku@20240307"
	LlmGpt520250807                   Llm = "gpt-5-2025-08-07"
	LlmGpt5120251113                  Llm = "gpt-5.1-2025-11-13"
	LlmGpt5220251211                  Llm = "gpt-5.2-2025-12-11"
	LlmGpt5420260305                  Llm = "gpt-5.4-2026-03-05"
	LlmGpt54Mini20260317              Llm = "gpt-5.4-mini-2026-03-17"
	LlmGpt54Nano20260317              Llm = "gpt-5.4-nano-2026-03-17"
	LlmGpt5520260423                  Llm = "gpt-5.5-2026-04-23"
	LlmGpt5Mini20250807               Llm = "gpt-5-mini-2025-08-07"
	LlmGpt5Nano20250807               Llm = "gpt-5-nano-2025-08-07"
	LlmGpt4120250414                  Llm = "gpt-4.1-2025-04-14"
	LlmGpt41Mini20250414              Llm = "gpt-4.1-mini-2025-04-14"
	LlmGpt41Nano20250414              Llm = "gpt-4.1-nano-2025-04-14"
	LlmGpt4OMini20240718              Llm = "gpt-4o-mini-2024-07-18"
	LlmGpt4O20241120                  Llm = "gpt-4o-2024-11-20"
	LlmGpt4O20240806                  Llm = "gpt-4o-2024-08-06"
	LlmGpt4O20240513                  Llm = "gpt-4o-2024-05-13"
	LlmGpt40613                       Llm = "gpt-4-0613"
	LlmGpt40314                       Llm = "gpt-4-0314"
	LlmGpt4Turbo20240409              Llm = "gpt-4-turbo-2024-04-09"
	LlmGpt35Turbo0125                 Llm = "gpt-3.5-turbo-0125"
	LlmGpt35Turbo1106                 Llm = "gpt-3.5-turbo-1106"
	LlmWattTool8B                     Llm = "watt-tool-8b"
	LlmWattTool70B                    Llm = "watt-tool-70b"
)

func NewLlmFromString(s string) (Llm, error) {
	switch s {
	case "gpt-4o-mini":
		return LlmGpt4OMini, nil
	case "gpt-4o":
		return LlmGpt4O, nil
	case "gpt-4":
		return LlmGpt4, nil
	case "gpt-4-turbo":
		return LlmGpt4Turbo, nil
	case "gpt-4.1":
		return LlmGpt41, nil
	case "gpt-4.1-mini":
		return LlmGpt41Mini, nil
	case "gpt-4.1-nano":
		return LlmGpt41Nano, nil
	case "gpt-5":
		return LlmGpt5, nil
	case "gpt-5.1":
		return LlmGpt51, nil
	case "gpt-5.2":
		return LlmGpt52, nil
	case "gpt-5.2-chat-latest":
		return LlmGpt52ChatLatest, nil
	case "gpt-5.4":
		return LlmGpt54, nil
	case "gpt-5.4-mini":
		return LlmGpt54Mini, nil
	case "gpt-5.4-nano":
		return LlmGpt54Nano, nil
	case "gpt-5.5":
		return LlmGpt55, nil
	case "gpt-5.6-sol":
		return LlmGpt56Sol, nil
	case "gpt-5.6-terra":
		return LlmGpt56Terra, nil
	case "gpt-5.6-luna":
		return LlmGpt56Luna, nil
	case "gpt-6-astra":
		return LlmGpt6Astra, nil
	case "gpt-6-sol":
		return LlmGpt6Sol, nil
	case "gpt-6-luna":
		return LlmGpt6Luna, nil
	case "gpt-6.1-sol":
		return LlmGpt61Sol, nil
	case "gpt-6.1-terra":
		return LlmGpt61Terra, nil
	case "gpt-5-mini":
		return LlmGpt5Mini, nil
	case "gpt-5-nano":
		return LlmGpt5Nano, nil
	case "gpt-3.5-turbo":
		return LlmGpt35Turbo, nil
	case "gemini-1.5-pro":
		return LlmGemini15Pro, nil
	case "gemini-1.5-flash":
		return LlmGemini15Flash, nil
	case "gemini-2.0-flash":
		return LlmGemini20Flash, nil
	case "gemini-2.0-flash-lite":
		return LlmGemini20FlashLite, nil
	case "gemini-2.5-flash-lite":
		return LlmGemini25FlashLite, nil
	case "gemini-2.5-flash":
		return LlmGemini25Flash, nil
	case "gemini-3-pro-preview":
		return LlmGemini3ProPreview, nil
	case "gemini-3-flash-preview":
		return LlmGemini3FlashPreview, nil
	case "gemini-3.1-pro-preview":
		return LlmGemini31ProPreview, nil
	case "gemini-3.1-flash-lite-preview":
		return LlmGemini31FlashLitePreview, nil
	case "gemini-3.1-flash-lite":
		return LlmGemini31FlashLite, nil
	case "gemini-3.5-flash":
		return LlmGemini35Flash, nil
	case "gemini-3.5-flash-lite":
		return LlmGemini35FlashLite, nil
	case "gemini-3.6-flash":
		return LlmGemini36Flash, nil
	case "gemini-3.7-flash":
		return LlmGemini37Flash, nil
	case "gemini-3.8-flash":
		return LlmGemini38Flash, nil
	case "gemini-4-argon":
		return LlmGemini4Argon, nil
	case "gemini-4-hydrogen":
		return LlmGemini4Hydrogen, nil
	case "claude-sonnet-4-5":
		return LlmClaudeSonnet45, nil
	case "claude-opus-4-6":
		return LlmClaudeOpus46, nil
	case "claude-opus-4-7":
		return LlmClaudeOpus47, nil
	case "claude-opus-4-8":
		return LlmClaudeOpus48, nil
	case "claude-opus-5":
		return LlmClaudeOpus5, nil
	case "claude-opus-5-5":
		return LlmClaudeOpus55, nil
	case "claude-sonnet-4-6":
		return LlmClaudeSonnet46, nil
	case "claude-sonnet-5":
		return LlmClaudeSonnet5, nil
	case "claude-sonnet-5-5":
		return LlmClaudeSonnet55, nil
	case "claude-sonnet-4":
		return LlmClaudeSonnet4, nil
	case "claude-haiku-4-5":
		return LlmClaudeHaiku45, nil
	case "claude-haiku-5-5":
		return LlmClaudeHaiku55, nil
	case "claude-3-7-sonnet":
		return LlmClaude37Sonnet, nil
	case "claude-3-5-sonnet":
		return LlmClaude35Sonnet, nil
	case "claude-3-5-sonnet-v1":
		return LlmClaude35SonnetV1, nil
	case "claude-3-haiku":
		return LlmClaude3Haiku, nil
	case "custom-llm":
		return LlmCustomLlm, nil
	case "qwen3-4b":
		return LlmQwen34B, nil
	case "qwen3-30b-a3b":
		return LlmQwen330BA3B, nil
	case "qwen36-35b-a3b":
		return LlmQwen3635BA3B, nil
	case "qwen35-397b-a17b":
		return LlmQwen35397BA17B, nil
	case "gpt-oss-20b":
		return LlmGptOss20B, nil
	case "gpt-oss-120b":
		return LlmGptOss120B, nil
	case "glm-45-air-fp8":
		return LlmGlm45AirFp8, nil
	case "glm-52":
		return LlmGlm52, nil
	case "deepseek-v41-flash":
		return LlmDeepseekV41Flash, nil
	case "nemotron-35-super-vl-nvfp4":
		return LlmNemotron35SuperVlNvfp4, nil
	case "gemini-2.5-flash-preview-09-2025":
		return LlmGemini25FlashPreview092025, nil
	case "gemini-2.5-flash-lite-preview-09-2025":
		return LlmGemini25FlashLitePreview092025, nil
	case "gemini-2.5-flash-preview-05-20":
		return LlmGemini25FlashPreview0520, nil
	case "gemini-2.5-flash-preview-04-17":
		return LlmGemini25FlashPreview0417, nil
	case "gemini-2.5-flash-lite-preview-06-17":
		return LlmGemini25FlashLitePreview0617, nil
	case "gemini-2.0-flash-lite-001":
		return LlmGemini20FlashLite001, nil
	case "gemini-2.0-flash-001":
		return LlmGemini20Flash001, nil
	case "gemini-1.5-flash-002":
		return LlmGemini15Flash002, nil
	case "gemini-1.5-flash-001":
		return LlmGemini15Flash001, nil
	case "gemini-1.5-pro-002":
		return LlmGemini15Pro002, nil
	case "gemini-1.5-pro-001":
		return LlmGemini15Pro001, nil
	case "claude-sonnet-4@20250514":
		return LlmClaudeSonnet4At20250514, nil
	case "claude-sonnet-4-5@20250929":
		return LlmClaudeSonnet45At20250929, nil
	case "claude-haiku-4-5@20251001":
		return LlmClaudeHaiku45At20251001, nil
	case "claude-3-7-sonnet@20250219":
		return LlmClaude37SonnetAt20250219, nil
	case "claude-3-5-sonnet@20240620":
		return LlmClaude35SonnetAt20240620, nil
	case "claude-3-5-sonnet-v2@20241022":
		return LlmClaude35SonnetV2At20241022, nil
	case "claude-3-haiku@20240307":
		return LlmClaude3HaikuAt20240307, nil
	case "gpt-5-2025-08-07":
		return LlmGpt520250807, nil
	case "gpt-5.1-2025-11-13":
		return LlmGpt5120251113, nil
	case "gpt-5.2-2025-12-11":
		return LlmGpt5220251211, nil
	case "gpt-5.4-2026-03-05":
		return LlmGpt5420260305, nil
	case "gpt-5.4-mini-2026-03-17":
		return LlmGpt54Mini20260317, nil
	case "gpt-5.4-nano-2026-03-17":
		return LlmGpt54Nano20260317, nil
	case "gpt-5.5-2026-04-23":
		return LlmGpt5520260423, nil
	case "gpt-5-mini-2025-08-07":
		return LlmGpt5Mini20250807, nil
	case "gpt-5-nano-2025-08-07":
		return LlmGpt5Nano20250807, nil
	case "gpt-4.1-2025-04-14":
		return LlmGpt4120250414, nil
	case "gpt-4.1-mini-2025-04-14":
		return LlmGpt41Mini20250414, nil
	case "gpt-4.1-nano-2025-04-14":
		return LlmGpt41Nano20250414, nil
	case "gpt-4o-mini-2024-07-18":
		return LlmGpt4OMini20240718, nil
	case "gpt-4o-2024-11-20":
		return LlmGpt4O20241120, nil
	case "gpt-4o-2024-08-06":
		return LlmGpt4O20240806, nil
	case "gpt-4o-2024-05-13":
		return LlmGpt4O20240513, nil
	case "gpt-4-0613":
		return LlmGpt40613, nil
	case "gpt-4-0314":
		return LlmGpt40314, nil
	case "gpt-4-turbo-2024-04-09":
		return LlmGpt4Turbo20240409, nil
	case "gpt-3.5-turbo-0125":
		return LlmGpt35Turbo0125, nil
	case "gpt-3.5-turbo-1106":
		return LlmGpt35Turbo1106, nil
	case "watt-tool-8b":
		return LlmWattTool8B, nil
	case "watt-tool-70b":
		return LlmWattTool70B, nil
	}
	var t Llm
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e Llm) Ptr() *Llm {
	return &e
}
