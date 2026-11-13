package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodySpeechToTextV1SpeechToTextPost struct {
	// When enable_logging is set to false zero retention mode will be used for the request. This will mean log and transcript storage features are unavailable for this request. Zero retention mode may only be used by enterprise customers.
	EnableLogging *bool `json:"-" url:"enable_logging,omitempty" explode:"true" serializationStyle:"form" queryParam:"enable_logging"`
	// A single-use authentication token created via POST /v1/single-use-token/batch_scribe. This token can only be used once and expires after 15 minutes. Alternative to API key or bearer token authentication for frontend clients.
	Token *string `json:"-" url:"token,omitempty" explode:"true" serializationStyle:"form" queryParam:"token"`
	// The ID of the model to use for transcription.
	ModelId string `json:"model_id" xml:"model_id" required:"true"`
	// The file to transcribe (100ms minimum audio length). All major audio and video formats are supported. Exactly one of the file or cloud_storage_url parameters must be provided. The file size must be less than 5.0GB.
	File *string `json:"file,omitempty" xml:"file,omitempty"`
	// An ISO-639-1 or ISO-639-3 language_code corresponding to the language of the audio file. Can sometimes improve transcription performance if known beforehand. Defaults to null, in this case the language is predicted automatically.
	LanguageCode *string `json:"language_code,omitempty" xml:"language_code,omitempty"`
	// Additional ISO-639-1 or ISO-639-3 language codes that may be present in the audio, alongside language_code. Improves transcription accuracy for multilingual audio. Supported by scribe_v3.
	SecondaryLanguageCodes []string `json:"secondary_language_codes,omitempty" xml:"secondary_language_codes,omitempty"`
	// How numbers are formatted in the transcript: 'digits' (e.g. '42') or 'words' (e.g. 'forty-two'). Supported by scribe_v3.
	NumbersFormat *ConvertSpeechToTextRequestNumbersFormat `json:"numbers_format,omitempty" xml:"numbers_format,omitempty"`
	// Natural-language instruction applied to the final transcript (max 2000 characters). The edited text is returned in 'edited_transcript' alongside the original transcript. Cannot be combined with entity_detection, entity_redaction or use_multi_channel. Usage of this parameter will incur an additional 30% surcharge on the base transcription cost, billed for at least 10 seconds of audio.
	TranscriptEdit *string `json:"transcript_edit,omitempty" xml:"transcript_edit,omitempty"`
	// Whether to tag audio events like (laughter), (footsteps), etc. in the transcription.
	TagAudioEvents *bool `json:"tag_audio_events,omitempty" xml:"tag_audio_events,omitempty"`
	// The maximum amount of speakers talking in the uploaded file. Can help with predicting who speaks when. The maximum amount of speakers that can be predicted is 32. Defaults to null, in this case the amount of speakers is set to the maximum value the model supports.
	NumSpeakers *int `json:"num_speakers,omitempty" xml:"num_speakers,omitempty" min:"1" max:"32"`
	// The granularity of the timestamps in the transcription. 'word' provides word-level timestamps and 'character' provides character-level timestamps per word.
	TimestampsGranularity *ConvertSpeechToTextRequestTimestampsGranularity `json:"timestamps_granularity,omitempty" xml:"timestamps_granularity,omitempty"`
	// Whether to annotate which speaker is currently talking in the uploaded file.
	Diarize *bool `json:"diarize,omitempty" xml:"diarize,omitempty"`
	// Diarization threshold to apply during speaker diarization. A higher value means there will be a lower chance of one speaker being diarized as two different speakers but also a higher chance of two different speakers being diarized as one speaker (less total speakers predicted). A low value means there will be a higher chance of one speaker being diarized as two different speakers but also a lower chance of two different speakers being diarized as one speaker (more total speakers predicted). Can only be set when diarize=True and num_speakers=None. Defaults to None, in which case we will choose a threshold based on the model_id (0.22 usually).
	DiarizationThreshold *float64           `json:"diarization_threshold,omitempty" xml:"diarization_threshold,omitempty" min:"0.1" max:"0.4"`
	AdditionalFormats    *AdditionalFormats `json:"additional_formats,omitempty" xml:"additional_formats,omitempty" maxItems:"10"`
	// The format of input audio. Options are 'pcm_s16le_16' or 'other' For `pcm_s16le_16`, the input audio must be 16-bit PCM at a 16kHz sample rate, single channel (mono), and little-endian byte order. Latency will be lower than with passing an encoded waveform.
	FileFormat *ConvertSpeechToTextRequestFileFormat `json:"file_format,omitempty" xml:"file_format,omitempty"`
	// The URL of an audio or video file to transcribe. Supports hosted video or audio files, YouTube video URLs, TikTok video URLs, and other video hosting services.
	SourceUrl *string `json:"source_url,omitempty" xml:"source_url,omitempty"`
	// Whether to send the transcription result to configured speech-to-text webhooks.  If set the request will return early without the transcription, which will be delivered later via webhook.
	Webhook *bool `json:"webhook,omitempty" xml:"webhook,omitempty"`
	// Optional specific webhook ID to send the transcription result to. Only valid when webhook is set to true. If not provided, transcription will be sent to all configured speech-to-text webhooks.
	WebhookId *string `json:"webhook_id,omitempty" xml:"webhook_id,omitempty"`
	// Controls the randomness of the transcription output. Accepts values between 0.0 and 2.0, where higher values result in more diverse and less deterministic results. If omitted, we will use a temperature based on the model you selected which is usually 0.
	Temperature *float64 `json:"temperature,omitempty" xml:"temperature,omitempty" min:"0" max:"2"`
	// If specified, our system will make a best effort to sample deterministically, such that repeated requests with the same seed and parameters should return the same result. Determinism is not guaranteed. Must be an integer between 0 and 2147483647.
	Seed *int `json:"seed,omitempty" xml:"seed,omitempty" min:"0" max:"2147483647"`
	// Whether the audio file contains multiple channels where each channel contains a single speaker. When enabled, each channel is transcribed independently. By default a separate transcript is returned per channel; set multichannel_output_style='combined' to instead receive a single transcript with all channels merged and sorted by time. Each word in the response includes a 'channel_index' field indicating which channel it was spoken on. A maximum of 5 channels is supported. Each channel is billed independently at the full audio duration, so cost scales linearly with the number of channels.
	UseMultiChannel *bool `json:"use_multi_channel,omitempty" xml:"use_multi_channel,omitempty"`
	// Controls the response shape when use_multi_channel is enabled. 'separate' (default) returns one transcript per channel under 'transcripts'. 'combined' merges all channels into a single transcript whose words are sorted by start time, each carrying a 'channel_index' - matching the single-channel response shape. 'combined' requires timestamps (timestamps_granularity must not be 'none') and does not support entity detection or redaction.
	MultichannelOutputStyle *ConvertSpeechToTextRequestMultichannelOutputStyle `json:"multichannel_output_style,omitempty" xml:"multichannel_output_style,omitempty"`
	// Optional metadata to be included in the webhook response. This should be a JSON string representing an object with a maximum depth of 2 levels and maximum size of 16KB. Useful for tracking internal IDs, job references, or other contextual information.
	WebhookMetadata *BodySpeechToTextV1SpeechToTextPostWebhookMetadata `json:"webhook_metadata,omitempty" xml:"webhook_metadata,omitempty"`
	// Detect entities in the transcript. Can be 'all' to detect all entities, a single entity type or category string, or a list of entity types/categories. Categories include 'pii', 'phi', 'pci', 'other', 'offensive_language'. When enabled, detected entities will be returned in the 'entities' field with their text, type, and character positions.
	EntityDetection *BodySpeechToTextV1SpeechToTextPostEntityDetection `json:"entity_detection,omitempty" xml:"entity_detection,omitempty"`
	// If true, the transcription will not have any filler words, false starts and non-speech sounds. Supported by scribe_v2, scribe_v2_medical and scribe_v3.
	NoVerbatim *bool `json:"no_verbatim,omitempty" xml:"no_verbatim,omitempty"`
	// Whether to use the speaker library for identifying known speakers during diarization. When enabled and diarize is true, detected speakers will be matched against registered speakers in the workspace's speaker library.
	UseSpeakerLibrary *bool `json:"use_speaker_library,omitempty" xml:"use_speaker_library,omitempty"`
	// Whether to detect speaker roles (agent vs customer). Requires diarize=true. Cannot be used with use_multi_channel=true. When enabled, speaker_id values will be 'agent' and 'customer' instead of 'speaker_0', 'speaker_1', etc. Usage incurs an additional 10% surcharge on base transcription cost.
	DetectSpeakerRoles *bool `json:"detect_speaker_roles,omitempty" xml:"detect_speaker_roles,omitempty"`
	// Redact entities from the transcript text. Accepts the same format as entity_detection: 'all', a category ('pii', 'phi'), or specific entity types. Must be a subset of entity_detection. When redaction is enabled, the entities field will not be returned.
	EntityRedaction *BodySpeechToTextV1SpeechToTextPostEntityRedaction `json:"entity_redaction,omitempty" xml:"entity_redaction,omitempty"`
	// How to format redacted entities. 'redacted' replaces with {REDACTED}, 'entity_type' replaces with {ENTITY_TYPE}, 'enumerated_entity_type' replaces with {ENTITY_TYPE_N} where N enumerates each occurrence. Only used when entity_redaction is set.
	EntityRedactionMode *string `json:"entity_redaction_mode,omitempty" xml:"entity_redaction_mode,omitempty"`
	// A list of keyterms to bias the transcription towards.           The keyterms are words or phrases you want the model to recognise more accurately.           The number of keyterms cannot exceed 1000.           The length of each keyterm must be at most 50 characters.           Keyterms can contain at most 5 words (after normalisation).           For example ["hello", "world", "technical term"].           The following characters are not supported: `<`, `>`, `{`, `}`, `[`, `]`, `\`.           Usage of this parameter will incur an additional 20% surcharge on the base transcription cost.           When more than 100 keyterms are provided, a minimum billable duration of 20 seconds applies per request.
	Keyterms []string `json:"keyterms,omitempty" xml:"keyterms,omitempty"`
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetEnableLogging() (value *bool) {
	if b == nil {
		return
	}
	return b.EnableLogging
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetToken() (value *string) {
	if b == nil {
		return
	}
	return b.Token
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetModelId() (value string) {
	if b == nil {
		return
	}
	return b.ModelId
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetFile() (value *string) {
	if b == nil {
		return
	}
	return b.File
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetLanguageCode() (value *string) {
	if b == nil {
		return
	}
	return b.LanguageCode
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetSecondaryLanguageCodes() (value []string) {
	if b == nil {
		return
	}
	return b.SecondaryLanguageCodes
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetNumbersFormat() (value *ConvertSpeechToTextRequestNumbersFormat) {
	if b == nil {
		return
	}
	return b.NumbersFormat
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetTranscriptEdit() (value *string) {
	if b == nil {
		return
	}
	return b.TranscriptEdit
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetTagAudioEvents() (value *bool) {
	if b == nil {
		return
	}
	return b.TagAudioEvents
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetNumSpeakers() (value *int) {
	if b == nil {
		return
	}
	return b.NumSpeakers
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetTimestampsGranularity() (value *ConvertSpeechToTextRequestTimestampsGranularity) {
	if b == nil {
		return
	}
	return b.TimestampsGranularity
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetDiarize() (value *bool) {
	if b == nil {
		return
	}
	return b.Diarize
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetDiarizationThreshold() (value *float64) {
	if b == nil {
		return
	}
	return b.DiarizationThreshold
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetAdditionalFormats() (value *AdditionalFormats) {
	if b == nil {
		return
	}
	return b.AdditionalFormats
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetFileFormat() (value *ConvertSpeechToTextRequestFileFormat) {
	if b == nil {
		return
	}
	return b.FileFormat
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetSourceUrl() (value *string) {
	if b == nil {
		return
	}
	return b.SourceUrl
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetWebhook() (value *bool) {
	if b == nil {
		return
	}
	return b.Webhook
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetWebhookId() (value *string) {
	if b == nil {
		return
	}
	return b.WebhookId
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetTemperature() (value *float64) {
	if b == nil {
		return
	}
	return b.Temperature
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetSeed() (value *int) {
	if b == nil {
		return
	}
	return b.Seed
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetUseMultiChannel() (value *bool) {
	if b == nil {
		return
	}
	return b.UseMultiChannel
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetMultichannelOutputStyle() (value *ConvertSpeechToTextRequestMultichannelOutputStyle) {
	if b == nil {
		return
	}
	return b.MultichannelOutputStyle
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetWebhookMetadata() (value *BodySpeechToTextV1SpeechToTextPostWebhookMetadata) {
	if b == nil {
		return
	}
	return b.WebhookMetadata
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetEntityDetection() (value *BodySpeechToTextV1SpeechToTextPostEntityDetection) {
	if b == nil {
		return
	}
	return b.EntityDetection
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetNoVerbatim() (value *bool) {
	if b == nil {
		return
	}
	return b.NoVerbatim
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetUseSpeakerLibrary() (value *bool) {
	if b == nil {
		return
	}
	return b.UseSpeakerLibrary
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetDetectSpeakerRoles() (value *bool) {
	if b == nil {
		return
	}
	return b.DetectSpeakerRoles
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetEntityRedaction() (value *BodySpeechToTextV1SpeechToTextPostEntityRedaction) {
	if b == nil {
		return
	}
	return b.EntityRedaction
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetEntityRedactionMode() (value *string) {
	if b == nil {
		return
	}
	return b.EntityRedactionMode
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetKeyterms() (value []string) {
	if b == nil {
		return
	}
	return b.Keyterms
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetEnableLogging(enableLogging *bool) {
	b.EnableLogging = enableLogging
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetToken(token *string) {
	b.Token = token
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetModelId(modelId string) {
	b.ModelId = modelId
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetFile(file *string) {
	b.File = file
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetLanguageCode(languageCode *string) {
	b.LanguageCode = languageCode
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetSecondaryLanguageCodes(secondaryLanguageCodes []string) {
	b.SecondaryLanguageCodes = secondaryLanguageCodes
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetNumbersFormat(numbersFormat *ConvertSpeechToTextRequestNumbersFormat) {
	b.NumbersFormat = numbersFormat
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetTranscriptEdit(transcriptEdit *string) {
	b.TranscriptEdit = transcriptEdit
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetTagAudioEvents(tagAudioEvents *bool) {
	b.TagAudioEvents = tagAudioEvents
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetNumSpeakers(numSpeakers *int) {
	b.NumSpeakers = numSpeakers
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetTimestampsGranularity(timestampsGranularity *ConvertSpeechToTextRequestTimestampsGranularity) {
	b.TimestampsGranularity = timestampsGranularity
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetDiarize(diarize *bool) {
	b.Diarize = diarize
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetDiarizationThreshold(diarizationThreshold *float64) {
	b.DiarizationThreshold = diarizationThreshold
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetAdditionalFormats(additionalFormats *AdditionalFormats) {
	b.AdditionalFormats = additionalFormats
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetFileFormat(fileFormat *ConvertSpeechToTextRequestFileFormat) {
	b.FileFormat = fileFormat
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetSourceUrl(sourceUrl *string) {
	b.SourceUrl = sourceUrl
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetWebhook(webhook *bool) {
	b.Webhook = webhook
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetWebhookId(webhookId *string) {
	b.WebhookId = webhookId
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetTemperature(temperature *float64) {
	b.Temperature = temperature
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetSeed(seed *int) {
	b.Seed = seed
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetUseMultiChannel(useMultiChannel *bool) {
	b.UseMultiChannel = useMultiChannel
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetMultichannelOutputStyle(multichannelOutputStyle *ConvertSpeechToTextRequestMultichannelOutputStyle) {
	b.MultichannelOutputStyle = multichannelOutputStyle
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetWebhookMetadata(webhookMetadata *BodySpeechToTextV1SpeechToTextPostWebhookMetadata) {
	b.WebhookMetadata = webhookMetadata
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetEntityDetection(entityDetection *BodySpeechToTextV1SpeechToTextPostEntityDetection) {
	b.EntityDetection = entityDetection
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetNoVerbatim(noVerbatim *bool) {
	b.NoVerbatim = noVerbatim
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetUseSpeakerLibrary(useSpeakerLibrary *bool) {
	b.UseSpeakerLibrary = useSpeakerLibrary
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetDetectSpeakerRoles(detectSpeakerRoles *bool) {
	b.DetectSpeakerRoles = detectSpeakerRoles
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetEntityRedaction(entityRedaction *BodySpeechToTextV1SpeechToTextPostEntityRedaction) {
	b.EntityRedaction = entityRedaction
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetEntityRedactionMode(entityRedactionMode *string) {
	b.EntityRedactionMode = entityRedactionMode
}

func (b *BodySpeechToTextV1SpeechToTextPost) SetKeyterms(keyterms []string) {
	b.Keyterms = keyterms
}

func (b *BodySpeechToTextV1SpeechToTextPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodySpeechToTextV1SpeechToTextPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodySpeechToTextV1SpeechToTextPost to string"
	}
	return string(jsonData)
}

func (b *BodySpeechToTextV1SpeechToTextPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodySpeechToTextV1SpeechToTextPost) MarshalJSON() ([]byte, error) {
	type embed BodySpeechToTextV1SpeechToTextPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
