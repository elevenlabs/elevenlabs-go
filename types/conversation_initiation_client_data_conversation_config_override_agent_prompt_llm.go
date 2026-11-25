package types

import (
	"fmt"
)

// The LLM to query with the prompt and the chat history. If using data residency, the LLM must be supported in the data residency environment
type ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm string

const (
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4OMini                      ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-4o-mini"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4O                          ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-4o"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4                           ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-4"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4Turbo                      ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-4-turbo"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt41                          ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-4.1"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt41Mini                      ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-4.1-mini"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt41Nano                      ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-4.1-nano"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5                           ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-5"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt51                          ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-5.1"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt52                          ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-5.2"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt52ChatLatest                ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-5.2-chat-latest"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt54                          ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-5.4"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt54Mini                      ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-5.4-mini"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt54Nano                      ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-5.4-nano"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt55                          ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-5.5"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt56Sol                       ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-5.6-sol"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt56Terra                     ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-5.6-terra"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt56Luna                      ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-5.6-luna"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt6Astra                      ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-6-astra"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt6Sol                        ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-6-sol"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt6Luna                       ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-6-luna"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt61Sol                       ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-6.1-sol"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt61Terra                     ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-6.1-terra"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt62Astra                     ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-6.2-astra"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt62Sol                       ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-6.2-sol"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5Mini                       ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-5-mini"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5Nano                       ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-5-nano"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt35Turbo                     ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-3.5-turbo"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini15Pro                    ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-1.5-pro"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini15Flash                  ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-1.5-flash"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini20Flash                  ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-2.0-flash"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini20FlashLite              ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-2.0-flash-lite"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini25FlashLite              ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-2.5-flash-lite"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini25Flash                  ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-2.5-flash"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini3ProPreview              ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-3-pro-preview"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini3FlashPreview            ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-3-flash-preview"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini31ProPreview             ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-3.1-pro-preview"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini31FlashLitePreview       ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-3.1-flash-lite-preview"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini31FlashLite              ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-3.1-flash-lite"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini35Flash                  ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-3.5-flash"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini35FlashLite              ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-3.5-flash-lite"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini36Flash                  ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-3.6-flash"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini37Flash                  ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-3.7-flash"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini38Flash                  ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-3.8-flash"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini4Argon                   ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-4-argon"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini4Hydrogen                ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-4-hydrogen"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini4Neon                    ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-4-neon"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeSonnet45                 ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "claude-sonnet-4-5"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeOpus46                   ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "claude-opus-4-6"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeOpus47                   ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "claude-opus-4-7"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeOpus48                   ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "claude-opus-4-8"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeOpus5                    ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "claude-opus-5"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeOpus55                   ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "claude-opus-5-5"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeSonnet46                 ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "claude-sonnet-4-6"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeSonnet5                  ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "claude-sonnet-5"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeSonnet55                 ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "claude-sonnet-5-5"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeSonnet4                  ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "claude-sonnet-4"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeHaiku45                  ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "claude-haiku-4-5"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeHaiku55                  ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "claude-haiku-5-5"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude37Sonnet                 ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "claude-3-7-sonnet"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude35Sonnet                 ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "claude-3-5-sonnet"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude35SonnetV1               ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "claude-3-5-sonnet-v1"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude3Haiku                   ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "claude-3-haiku"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmCustomLlm                      ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "custom-llm"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmQwen34B                        ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "qwen3-4b"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmQwen330BA3B                    ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "qwen3-30b-a3b"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmQwen3635BA3B                   ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "qwen36-35b-a3b"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmQwen35397BA17B                 ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "qwen35-397b-a17b"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGptOss20B                      ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-oss-20b"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGptOss120B                     ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-oss-120b"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGlm45AirFp8                    ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "glm-45-air-fp8"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGlm52                          ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "glm-52"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmDeepseekV41Flash               ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "deepseek-v41-flash"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmNemotron35SuperVlNvfp4         ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "nemotron-35-super-vl-nvfp4"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini25FlashPreview092025     ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-2.5-flash-preview-09-2025"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini25FlashLitePreview092025 ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-2.5-flash-lite-preview-09-2025"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini25FlashPreview0520       ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-2.5-flash-preview-05-20"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini25FlashPreview0417       ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-2.5-flash-preview-04-17"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini25FlashLitePreview0617   ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-2.5-flash-lite-preview-06-17"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini20FlashLite001           ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-2.0-flash-lite-001"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini20Flash001               ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-2.0-flash-001"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini15Flash002               ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-1.5-flash-002"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini15Flash001               ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-1.5-flash-001"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini15Pro002                 ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-1.5-pro-002"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini15Pro001                 ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gemini-1.5-pro-001"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeSonnet4At20250514        ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "claude-sonnet-4@20250514"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeSonnet45At20250929       ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "claude-sonnet-4-5@20250929"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeHaiku45At20251001        ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "claude-haiku-4-5@20251001"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude37SonnetAt20250219       ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "claude-3-7-sonnet@20250219"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude35SonnetAt20240620       ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "claude-3-5-sonnet@20240620"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude35SonnetV2At20241022     ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "claude-3-5-sonnet-v2@20241022"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude3HaikuAt20240307         ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "claude-3-haiku@20240307"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt520250807                   ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-5-2025-08-07"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5120251113                  ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-5.1-2025-11-13"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5220251211                  ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-5.2-2025-12-11"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5420260305                  ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-5.4-2026-03-05"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt54Mini20260317              ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-5.4-mini-2026-03-17"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt54Nano20260317              ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-5.4-nano-2026-03-17"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5520260423                  ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-5.5-2026-04-23"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5Mini20250807               ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-5-mini-2025-08-07"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5Nano20250807               ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-5-nano-2025-08-07"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4120250414                  ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-4.1-2025-04-14"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt41Mini20250414              ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-4.1-mini-2025-04-14"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt41Nano20250414              ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-4.1-nano-2025-04-14"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4OMini20240718              ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-4o-mini-2024-07-18"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4O20241120                  ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-4o-2024-11-20"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4O20240806                  ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-4o-2024-08-06"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4O20240513                  ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-4o-2024-05-13"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt40613                       ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-4-0613"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt40314                       ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-4-0314"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4Turbo20240409              ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-4-turbo-2024-04-09"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt35Turbo0125                 ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-3.5-turbo-0125"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt35Turbo1106                 ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "gpt-3.5-turbo-1106"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmWattTool8B                     ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "watt-tool-8b"
	ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmWattTool70B                    ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm = "watt-tool-70b"
)

func NewConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmFromString(s string) (ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm, error) {
	switch s {
	case "gpt-4o-mini":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4OMini, nil
	case "gpt-4o":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4O, nil
	case "gpt-4":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4, nil
	case "gpt-4-turbo":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4Turbo, nil
	case "gpt-4.1":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt41, nil
	case "gpt-4.1-mini":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt41Mini, nil
	case "gpt-4.1-nano":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt41Nano, nil
	case "gpt-5":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5, nil
	case "gpt-5.1":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt51, nil
	case "gpt-5.2":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt52, nil
	case "gpt-5.2-chat-latest":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt52ChatLatest, nil
	case "gpt-5.4":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt54, nil
	case "gpt-5.4-mini":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt54Mini, nil
	case "gpt-5.4-nano":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt54Nano, nil
	case "gpt-5.5":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt55, nil
	case "gpt-5.6-sol":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt56Sol, nil
	case "gpt-5.6-terra":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt56Terra, nil
	case "gpt-5.6-luna":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt56Luna, nil
	case "gpt-6-astra":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt6Astra, nil
	case "gpt-6-sol":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt6Sol, nil
	case "gpt-6-luna":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt6Luna, nil
	case "gpt-6.1-sol":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt61Sol, nil
	case "gpt-6.1-terra":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt61Terra, nil
	case "gpt-6.2-astra":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt62Astra, nil
	case "gpt-6.2-sol":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt62Sol, nil
	case "gpt-5-mini":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5Mini, nil
	case "gpt-5-nano":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5Nano, nil
	case "gpt-3.5-turbo":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt35Turbo, nil
	case "gemini-1.5-pro":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini15Pro, nil
	case "gemini-1.5-flash":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini15Flash, nil
	case "gemini-2.0-flash":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini20Flash, nil
	case "gemini-2.0-flash-lite":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini20FlashLite, nil
	case "gemini-2.5-flash-lite":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini25FlashLite, nil
	case "gemini-2.5-flash":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini25Flash, nil
	case "gemini-3-pro-preview":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini3ProPreview, nil
	case "gemini-3-flash-preview":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini3FlashPreview, nil
	case "gemini-3.1-pro-preview":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini31ProPreview, nil
	case "gemini-3.1-flash-lite-preview":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini31FlashLitePreview, nil
	case "gemini-3.1-flash-lite":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini31FlashLite, nil
	case "gemini-3.5-flash":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini35Flash, nil
	case "gemini-3.5-flash-lite":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini35FlashLite, nil
	case "gemini-3.6-flash":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini36Flash, nil
	case "gemini-3.7-flash":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini37Flash, nil
	case "gemini-3.8-flash":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini38Flash, nil
	case "gemini-4-argon":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini4Argon, nil
	case "gemini-4-hydrogen":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini4Hydrogen, nil
	case "gemini-4-neon":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini4Neon, nil
	case "claude-sonnet-4-5":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeSonnet45, nil
	case "claude-opus-4-6":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeOpus46, nil
	case "claude-opus-4-7":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeOpus47, nil
	case "claude-opus-4-8":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeOpus48, nil
	case "claude-opus-5":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeOpus5, nil
	case "claude-opus-5-5":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeOpus55, nil
	case "claude-sonnet-4-6":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeSonnet46, nil
	case "claude-sonnet-5":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeSonnet5, nil
	case "claude-sonnet-5-5":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeSonnet55, nil
	case "claude-sonnet-4":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeSonnet4, nil
	case "claude-haiku-4-5":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeHaiku45, nil
	case "claude-haiku-5-5":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeHaiku55, nil
	case "claude-3-7-sonnet":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude37Sonnet, nil
	case "claude-3-5-sonnet":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude35Sonnet, nil
	case "claude-3-5-sonnet-v1":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude35SonnetV1, nil
	case "claude-3-haiku":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude3Haiku, nil
	case "custom-llm":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmCustomLlm, nil
	case "qwen3-4b":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmQwen34B, nil
	case "qwen3-30b-a3b":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmQwen330BA3B, nil
	case "qwen36-35b-a3b":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmQwen3635BA3B, nil
	case "qwen35-397b-a17b":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmQwen35397BA17B, nil
	case "gpt-oss-20b":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGptOss20B, nil
	case "gpt-oss-120b":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGptOss120B, nil
	case "glm-45-air-fp8":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGlm45AirFp8, nil
	case "glm-52":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGlm52, nil
	case "deepseek-v41-flash":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmDeepseekV41Flash, nil
	case "nemotron-35-super-vl-nvfp4":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmNemotron35SuperVlNvfp4, nil
	case "gemini-2.5-flash-preview-09-2025":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini25FlashPreview092025, nil
	case "gemini-2.5-flash-lite-preview-09-2025":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini25FlashLitePreview092025, nil
	case "gemini-2.5-flash-preview-05-20":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini25FlashPreview0520, nil
	case "gemini-2.5-flash-preview-04-17":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini25FlashPreview0417, nil
	case "gemini-2.5-flash-lite-preview-06-17":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini25FlashLitePreview0617, nil
	case "gemini-2.0-flash-lite-001":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini20FlashLite001, nil
	case "gemini-2.0-flash-001":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini20Flash001, nil
	case "gemini-1.5-flash-002":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini15Flash002, nil
	case "gemini-1.5-flash-001":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini15Flash001, nil
	case "gemini-1.5-pro-002":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini15Pro002, nil
	case "gemini-1.5-pro-001":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGemini15Pro001, nil
	case "claude-sonnet-4@20250514":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeSonnet4At20250514, nil
	case "claude-sonnet-4-5@20250929":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeSonnet45At20250929, nil
	case "claude-haiku-4-5@20251001":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaudeHaiku45At20251001, nil
	case "claude-3-7-sonnet@20250219":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude37SonnetAt20250219, nil
	case "claude-3-5-sonnet@20240620":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude35SonnetAt20240620, nil
	case "claude-3-5-sonnet-v2@20241022":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude35SonnetV2At20241022, nil
	case "claude-3-haiku@20240307":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmClaude3HaikuAt20240307, nil
	case "gpt-5-2025-08-07":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt520250807, nil
	case "gpt-5.1-2025-11-13":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5120251113, nil
	case "gpt-5.2-2025-12-11":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5220251211, nil
	case "gpt-5.4-2026-03-05":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5420260305, nil
	case "gpt-5.4-mini-2026-03-17":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt54Mini20260317, nil
	case "gpt-5.4-nano-2026-03-17":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt54Nano20260317, nil
	case "gpt-5.5-2026-04-23":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5520260423, nil
	case "gpt-5-mini-2025-08-07":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5Mini20250807, nil
	case "gpt-5-nano-2025-08-07":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt5Nano20250807, nil
	case "gpt-4.1-2025-04-14":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4120250414, nil
	case "gpt-4.1-mini-2025-04-14":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt41Mini20250414, nil
	case "gpt-4.1-nano-2025-04-14":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt41Nano20250414, nil
	case "gpt-4o-mini-2024-07-18":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4OMini20240718, nil
	case "gpt-4o-2024-11-20":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4O20241120, nil
	case "gpt-4o-2024-08-06":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4O20240806, nil
	case "gpt-4o-2024-05-13":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4O20240513, nil
	case "gpt-4-0613":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt40613, nil
	case "gpt-4-0314":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt40314, nil
	case "gpt-4-turbo-2024-04-09":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt4Turbo20240409, nil
	case "gpt-3.5-turbo-0125":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt35Turbo0125, nil
	case "gpt-3.5-turbo-1106":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmGpt35Turbo1106, nil
	case "watt-tool-8b":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmWattTool8B, nil
	case "watt-tool-70b":
		return ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlmWattTool70B, nil
	}
	var t ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm) Ptr() *ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm {
	return &e
}
