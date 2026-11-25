# Reference

## History

### `List`

`GET /v1/history`

Get generated items

**Parameters:**

- `pageSize`
- `startAfterHistoryItemId`
- `voiceId`
- `modelId`
- `dateBeforeUnix`
- `dateAfterUnix`
- `sortDirection`
- `search`
- `source`
- `sources`

### `Get`

`GET /v1/history/{history_item_id}`

Get history item

**Parameters:**

- `historyItemId` (required)

### `Delete`

`DELETE /v1/history/{history_item_id}`

Delete history item

**Parameters:**

- `historyItemId` (required)

### `GetAudio`

`GET /v1/history/{history_item_id}/audio`

Get audio from history item

**Parameters:**

- `historyItemId` (required)

### `Download`

`POST /v1/history/download`

Download history items

## TextToSoundEffects

### `Convert`

`POST /v1/sound-generation`

Create sound effect

**Parameters:**

- `outputFormat`
- `enableLogging`

## AudioIsolation

### `Convert`

`POST /v1/audio-isolation`

Audio isolation

### `List`

`GET /v1/audio-isolation/history`

Get Audio Isolation History

**Parameters:**

- `pageSize`
- `page`
- `search`

### `Delete`

`DELETE /v1/audio-isolation/history/{history_item_id}`

Delete Audio Isolation History Item

**Parameters:**

- `historyItemId` (required)

### `Stream`

`POST /v1/audio-isolation/stream`

Audio isolation stream

## AudioEffects

### `GetCatalog`

`GET /v1/audio-effects/catalog`

Get Audio Effects Catalog

## Samples

### `Delete`

`DELETE /v1/voices/{voice_id}/samples/{sample_id}`

Delete voice sample

**Parameters:**

- `voiceId` (required)
- `sampleId` (required)

### `Create`

`POST /v1/voices/pvc/{voice_id}/samples`

Add Samples To Pvc Voice

**Parameters:**

- `voiceId` (required)

### `Update`

`POST /v1/voices/pvc/{voice_id}/samples/{sample_id}`

Update Pvc Voice Sample

**Parameters:**

- `voiceId` (required)
- `sampleId` (required)

### `Delete1`

`DELETE /v1/voices/pvc/{voice_id}/samples/{sample_id}`

Delete Pvc Voice Sample

**Parameters:**

- `voiceId` (required)
- `sampleId` (required)

## Audio

### `Get`

`GET /v1/voices/{voice_id}/samples/{sample_id}/audio`

Get audio from sample

**Parameters:**

- `voiceId` (required)
- `sampleId` (required)

### `Get1`

`GET /v1/dubbing/{dubbing_id}/audio/{language_code}`

Get dubbed audio

**Parameters:**

- `dubbingId` (required)
- `languageCode` (required)

### `Get2`

`GET /v1/convai/conversations/{conversation_id}/audio`

Get Conversation Audio

**Parameters:**

- `conversationId` (required)

### `Get3`

`GET /v1/voices/pvc/{voice_id}/samples/{sample_id}/audio`

Retrieve Voice Sample Audio

**Parameters:**

- `voiceId` (required)
- `sampleId` (required)
- `removeBackgroundNoise`

### `Get4`

`GET /v1/voices/pvc/{voice_id}/samples/{sample_id}/speakers/{speaker_id}/audio`

Retrieve Separated Speaker Audio

**Parameters:**

- `voiceId` (required)
- `sampleId` (required)
- `speakerId` (required)

## TextToSpeech

### `Convert`

`POST /v1/text-to-speech/{voice_id}`

Create speech

**Parameters:**

- `voiceId` (required)
- `enableLogging`
- `optimizeStreamingLatency`
- `outputFormat`

### `ConvertWithTimestamps`

`POST /v1/text-to-speech/{voice_id}/with-timestamps`

Create speech with timing

**Parameters:**

- `voiceId` (required)
- `enableLogging`
- `optimizeStreamingLatency`
- `outputFormat`

### `Stream`

`POST /v1/text-to-speech/{voice_id}/stream`

Stream speech

**Parameters:**

- `voiceId` (required)
- `enableLogging`
- `optimizeStreamingLatency`
- `outputFormat`

### `StreamWithTimestamps`

`POST /v1/text-to-speech/{voice_id}/stream/with-timestamps`

Stream speech with timing

**Parameters:**

- `voiceId` (required)
- `enableLogging`
- `optimizeStreamingLatency`
- `outputFormat`

### `List`

`GET /v1/flows/text-to-speech`

List Speech Generations

**Parameters:**

- `cursor`
- `pageSize`
- `status`
- `modelId`

### `Create`

`POST /v1/flows/text-to-speech`

Create Speech Generation

### `Get`

`GET /v1/flows/text-to-speech/{generation_id}`

Get Speech Generation

**Parameters:**

- `generationId` (required)

## TextToDialogue

### `Convert`

`POST /v1/text-to-dialogue`

Create dialogue

**Parameters:**

- `outputFormat`
- `enableLogging`

### `Stream`

`POST /v1/text-to-dialogue/stream`

Stream dialogue

**Parameters:**

- `outputFormat`
- `enableLogging`

### `StreamWithTimestamps`

`POST /v1/text-to-dialogue/stream/with-timestamps`

Text To Dialogue Streaming With Timestamps

**Parameters:**

- `outputFormat`
- `enableLogging`

### `FullWithTimestamps`

`POST /v1/text-to-dialogue/with-timestamps`

Text To Dialogue With Timestamps

**Parameters:**

- `outputFormat`
- `enableLogging`

## SpeechToSpeech

### `Convert`

`POST /v1/voice-changer/{voice_id}`

Voice changer

**Parameters:**

- `voiceId` (required)
- `enableLogging`
- `optimizeStreamingLatency`
- `outputFormat`

### `Stream`

`POST /v1/voice-changer/{voice_id}/stream`

Voice changer stream

**Parameters:**

- `voiceId` (required)
- `enableLogging`
- `optimizeStreamingLatency`
- `outputFormat`

## TextToVoice

### `Create`

`POST /v1/text-to-voice`

Create A New Voice From Voice Preview

### `Design`

`POST /v1/text-to-voice/design`

Design A Voice.

**Parameters:**

- `outputFormat`

### `Remix`

`POST /v1/text-to-voice/{voice_id}/remix`

Remix A Voice.

**Parameters:**

- `voiceId` (required)
- `outputFormat`

## Preview

### `Stream`

`GET /v1/text-to-voice/{generated_voice_id}/stream`

Text To Voice Preview Streaming

**Parameters:**

- `generatedVoiceId` (required)

## User

### `Get`

`GET /v1/user`

Get user

## Subscription

### `Get`

`GET /v1/user/subscription`

Get user subscription

## Collections

### `List`

`GET /v1/voices/collections`

List Voice Collections

### `Create`

`POST /v1/voices/collections`

Create Voice Collection

### `Get`

`GET /v1/voices/collections/{collection_id}`

Get Voice Collection

**Parameters:**

- `collectionId` (required)

### `Update`

`PATCH /v1/voices/collections/{collection_id}`

Update Voice Collection

**Parameters:**

- `collectionId` (required)

### `Delete`

`DELETE /v1/voices/collections/{collection_id}`

Delete Voice Collection

**Parameters:**

- `collectionId` (required)

### `AddVoice`

`POST /v1/voices/collections/{collection_id}/voice/{voice_id}`

Add Voice To Collection

**Parameters:**

- `collectionId` (required)
- `voiceId` (required)

### `RemoveVoice`

`DELETE /v1/voices/collections/{collection_id}/voice/{voice_id}`

Remove Voice From Collection

**Parameters:**

- `collectionId` (required)
- `voiceId` (required)

## Settings

### `GetDefault`

`GET /v1/voices/settings/default`

Get default voice settings

### `Get`

`GET /v1/voices/{voice_id}/settings`

Get voice settings

**Parameters:**

- `voiceId` (required)

### `Update`

`POST /v1/voices/{voice_id}/settings/edit`

Edit voice settings

**Parameters:**

- `voiceId` (required)

### `Get1`

`GET /v1/convai/settings`

Get Convai Settings

### `Update1`

`PATCH /v1/convai/settings`

Update Convai Settings

### `Get2`

`GET /v1/convai/settings/dashboard`

Get Convai Dashboard Settings

### `Update2`

`PATCH /v1/convai/settings/dashboard`

Update Convai Dashboard Settings

## Accents

### `Get`

`GET /v1/voices/accents`

Get Voice Accents

**Parameters:**

- `language`
- `modelId`

## Voices

### `Get`

`GET /v1/voices/{voice_id}`

Get voice

**Parameters:**

- `voiceId` (required)
- `withSettings`

### `Delete`

`DELETE /v1/voices/{voice_id}`

Delete voice

**Parameters:**

- `voiceId` (required)

### `Update`

`POST /v1/voices/{voice_id}/edit`

Edit voice

**Parameters:**

- `voiceId` (required)

### `Search`

`GET /v2/voices`

List voices

**Parameters:**

- `nextPageToken`
- `pageSize`
- `search`
- `sort`
- `sortDirection`
- `voiceType`
- `category`
- `fineTuningState`
- `collectionId`
- `gender`
- `age`
- `language`
- `accent`
- `useCases`
- `minNoticePeriodDays`
- `includeCustomRates`
- `includeLiveModerated`
- `highQuality`
- `includeTotalCount`
- `voiceIds`

### `ReplicateToIsolatedEnvironment`

`POST /v1/voices/{voice_id}/replicate-to-isolated-environment`

Replicate Voice To Isolated Environment

**Parameters:**

- `voiceId` (required)

### `Share`

`POST /v1/voices/add/{public_user_id}/{voice_id}`

Add shared voice

**Parameters:**

- `publicUserId` (required)
- `voiceId` (required)

### `FindSimilarVoices`

`POST /v1/similar-voices`

List similar voices

**Parameters:**

- `category`
- `gender`
- `age`
- `useCases`
- `minNoticePeriodDays`
- `includeCustomRates`
- `includeLiveModerated`
- `ownerId`

### `GetShared`

`GET /v1/shared-voices`

Get shared voices

**Parameters:**

- `pageSize`
- `category`
- `gender`
- `age`
- `accent`
- `language`
- `locale`
- `search`
- `useCases`
- `descriptives`
- `featured`
- `minNoticePeriodDays`
- `includeCustomRates`
- `includeLiveModerated`
- `readerAppEnabled`
- `ownerId`
- `sort`
- `page`

## Ivc

### `Create`

`POST /v1/voices/add`

Create voice clone

## Studio

### `CreatePodcast`

`POST /v1/studio/podcasts`

Create Podcast

## PronunciationDictionaries

### `Create`

`POST /v1/studio/projects/{project_id}/pronunciation-dictionaries`

Create Pronunciation Dictionaries

**Parameters:**

- `projectId` (required)

### `CreateFromFile`

`POST /v1/pronunciation-dictionaries/add-from-file`

Create a pronunciation dictionary from a file

### `CreateFromRules`

`POST /v1/pronunciation-dictionaries/add-from-rules`

Create a pronunciation dictionary from rules

### `Get`

`GET /v1/pronunciation-dictionaries/{pronunciation_dictionary_id}`

Get pronunciation dictionary

**Parameters:**

- `pronunciationDictionaryId` (required)

### `Update`

`PATCH /v1/pronunciation-dictionaries/{pronunciation_dictionary_id}`

Update Pronunciation Dictionary

**Parameters:**

- `pronunciationDictionaryId` (required)

### `Download`

`GET /v1/pronunciation-dictionaries/{dictionary_id}/{version_id}/download`

Get pronunciation dictionary by version

**Parameters:**

- `dictionaryId` (required)
- `versionId` (required)

### `List`

`GET /v1/pronunciation-dictionaries`

List pronunciation dictionaries

**Parameters:**

- `cursor`
- `pageSize`
- `sort`
- `sortDirection`
- `includeArchived`

## Projects

### `List`

`GET /v1/studio/projects`

List Studio Projects

### `Create`

`POST /v1/studio/projects`

Create Studio Project

### `Get`

`GET /v1/studio/projects/{project_id}`

Get Studio Project

**Parameters:**

- `projectId` (required)
- `shareId`

### `Update`

`POST /v1/studio/projects/{project_id}`

Update Studio Project

**Parameters:**

- `projectId` (required)

### `Delete`

`DELETE /v1/studio/projects/{project_id}`

Delete Studio Project

**Parameters:**

- `projectId` (required)

### `Convert`

`POST /v1/studio/projects/{project_id}/convert`

Convert Studio Project

**Parameters:**

- `projectId` (required)

### `GetMutedTracks`

`GET /v1/studio/projects/{project_id}/muted-tracks`

Get Project Muted Tracks

**Parameters:**

- `projectId` (required)

## Content

### `Update`

`POST /v1/studio/projects/{project_id}/content`

Update Studio Project Content

**Parameters:**

- `projectId` (required)

## Snapshots

### `List`

`GET /v1/studio/projects/{project_id}/snapshots`

List Studio Project Snapshots

**Parameters:**

- `projectId` (required)

### `Get`

`GET /v1/studio/projects/{project_id}/snapshots/{project_snapshot_id}`

Get Project Snapshot

**Parameters:**

- `projectId` (required)
- `projectSnapshotId` (required)

### `Stream`

`POST /v1/studio/projects/{project_id}/snapshots/{project_snapshot_id}/stream`

Stream Studio Project Audio

**Parameters:**

- `projectId` (required)
- `projectSnapshotId` (required)

### `StreamArchive`

`POST /v1/studio/projects/{project_id}/snapshots/{project_snapshot_id}/archive`

Stream Archive With Studio Project Audio

**Parameters:**

- `projectId` (required)
- `projectSnapshotId` (required)

### `List1`

`GET /v1/studio/projects/{project_id}/chapters/{chapter_id}/snapshots`

List Chapter Snapshots

**Parameters:**

- `projectId` (required)
- `chapterId` (required)

### `Get1`

`GET /v1/studio/projects/{project_id}/chapters/{chapter_id}/snapshots/{chapter_snapshot_id}`

Get Chapter Snapshot

**Parameters:**

- `projectId` (required)
- `chapterId` (required)
- `chapterSnapshotId` (required)

### `Stream1`

`POST /v1/studio/projects/{project_id}/chapters/{chapter_id}/snapshots/{chapter_snapshot_id}/stream`

Stream Chapter Audio

**Parameters:**

- `projectId` (required)
- `chapterId` (required)
- `chapterSnapshotId` (required)

## Chapters

### `List`

`GET /v1/studio/projects/{project_id}/chapters`

List Chapters

**Parameters:**

- `projectId` (required)

### `Create`

`POST /v1/studio/projects/{project_id}/chapters`

Create Chapter

**Parameters:**

- `projectId` (required)

### `Get`

`GET /v1/studio/projects/{project_id}/chapters/{chapter_id}`

Get Chapter

**Parameters:**

- `projectId` (required)
- `chapterId` (required)

### `Update`

`POST /v1/studio/projects/{project_id}/chapters/{chapter_id}`

Update Chapter

**Parameters:**

- `projectId` (required)
- `chapterId` (required)

### `Delete`

`DELETE /v1/studio/projects/{project_id}/chapters/{chapter_id}`

Delete Chapter

**Parameters:**

- `projectId` (required)
- `chapterId` (required)

### `Convert`

`POST /v1/studio/projects/{project_id}/chapters/{chapter_id}/convert`

Convert Chapter

**Parameters:**

- `projectId` (required)
- `chapterId` (required)

## AudioNative

### `Create`

`POST /v1/audio-native`

Create audio native project

### `GetSettings`

`GET /v1/audio-native/{project_id}/settings`

Get Audio Native Project Settings

**Parameters:**

- `projectId` (required)

### `Update`

`POST /v1/audio-native/{project_id}/content`

Update audio native project

**Parameters:**

- `projectId` (required)

### `UpdateContentFromUrl`

`POST /v1/audio-native/content`

Update Audio-Native Content From Url

## Project

### `List`

`GET /v1/dubbing/project`

List Dubbing Projects

**Parameters:**

- `cursor`
- `pageSize`
- `status`
- `sortDirection`

### `Create`

`POST /v1/dubbing/project`

Create Dubbing Project

### `Get`

`GET /v1/dubbing/project/{project_id}`

Get Dubbing Project

**Parameters:**

- `projectId` (required)

### `Delete`

`DELETE /v1/dubbing/project/{project_id}`

Delete Dubbing Project

**Parameters:**

- `projectId` (required)

## Language

### `List`

`GET /v1/dubbing/project/{project_id}/language`

List Dubbing Language Targets

**Parameters:**

- `projectId` (required)
- `cursor`
- `pageSize`
- `status`

### `Create`

`POST /v1/dubbing/project/{project_id}/language`

Create Dubbing Language Target

**Parameters:**

- `projectId` (required)

### `Get`

`GET /v1/dubbing/project/{project_id}/language/{language_id}`

Get Dubbing Language Target

**Parameters:**

- `projectId` (required)
- `languageId` (required)

### `Delete`

`DELETE /v1/dubbing/project/{project_id}/language/{language_id}`

Delete Dubbing Language Target

**Parameters:**

- `projectId` (required)
- `languageId` (required)

## Transcript

### `Get`

`GET /v1/dubbing/project/{project_id}/transcript`

Get Dubbing Transcript

**Parameters:**

- `projectId` (required)

### `UpdateSegment`

`PATCH /v1/dubbing/project/{project_id}/transcript/segment/{segment_id}`

Update Dubbing Transcript Segment

**Parameters:**

- `projectId` (required)
- `segmentId` (required)

### `DeleteSegment`

`DELETE /v1/dubbing/project/{project_id}/transcript/segment/{segment_id}`

Delete Dubbing Transcript Segment

**Parameters:**

- `projectId` (required)
- `segmentId` (required)

### `UpdateSegments`

`PATCH /v1/dubbing/project/{project_id}/transcript/segments`

Update Dubbing Transcript Segments

**Parameters:**

- `projectId` (required)

### `CreateSegment`

`POST /v1/dubbing/project/{project_id}/transcript/segment`

Add Dubbing Transcript Segment

**Parameters:**

- `projectId` (required)

### `Get1`

`GET /v1/dubbing/project/{project_id}/language/{language_id}/transcript`

Get Dubbing Target Transcript

**Parameters:**

- `projectId` (required)
- `languageId` (required)

### `UpdateSegment1`

`PATCH /v1/dubbing/project/{project_id}/language/{language_id}/transcript/segment/{segment_id}`

Update Dubbing Target Transcript Segment

**Parameters:**

- `projectId` (required)
- `languageId` (required)
- `segmentId` (required)

### `UpdateSegments1`

`PATCH /v1/dubbing/project/{project_id}/language/{language_id}/transcript/segments`

Update Dubbing Target Transcript Segments

**Parameters:**

- `projectId` (required)
- `languageId` (required)

### `Regenerate`

`POST /v1/dubbing/project/{project_id}/language/{language_id}/transcript/regenerate`

Regenerate Dubbing Target

**Parameters:**

- `projectId` (required)
- `languageId` (required)

## Dubbing

### `List`

`GET /v1/dubbing`

List Dubs

**Parameters:**

- `cursor`
- `pageSize`
- `dubbingStatus`
- `dubbingStatuses`
- `dubbingModels`
- `targetLanguageCodes`
- `creationSources`
- `filterByCreator`
- `orderBy`
- `orderDirection`

### `Create`

`POST /v1/dubbing`

Dub a video or audio file

### `Get`

`GET /v1/dubbing/{dubbing_id}`

Get dubbing

**Parameters:**

- `dubbingId` (required)

### `Delete`

`DELETE /v1/dubbing/{dubbing_id}`

Delete dubbing

**Parameters:**

- `dubbingId` (required)

## Transcripts

### `Get`

`GET /v1/dubbing/{dubbing_id}/transcripts/{language_code}/format/{format_type}`

Retrieve A Transcript

**Parameters:**

- `dubbingId` (required)
- `languageCode` (required)
- `formatType` (required)

### `Get1`

`GET /v1/speech-to-text/transcripts/{transcription_id}`

Get Transcript By Id

**Parameters:**

- `transcriptionId` (required)

### `Delete`

`DELETE /v1/speech-to-text/transcripts/{transcription_id}`

Delete Transcript By Id

**Parameters:**

- `transcriptionId` (required)

## Models

### `List`

`GET /v1/models`

List models

## ContentProvenance

### `Watermark`

`POST /v1/content-provenance/watermark`

Check Watermark

## Rules

### `Set`

`POST /v1/pronunciation-dictionaries/{pronunciation_dictionary_id}/set-rules`

Set pronunciation dictionary rules

**Parameters:**

- `pronunciationDictionaryId` (required)

### `Add`

`POST /v1/pronunciation-dictionaries/{pronunciation_dictionary_id}/add-rules`

Add pronunciation dictionary rules

**Parameters:**

- `pronunciationDictionaryId` (required)

### `Remove`

`POST /v1/pronunciation-dictionaries/{pronunciation_dictionary_id}/remove-rules`

Remove pronunciation dictionary rules

**Parameters:**

- `pronunciationDictionaryId` (required)

### `Create`

`POST /v1/reception/customer-facing-agents/{agent_id}/rules`

Add Rule

**Parameters:**

- `agentId` (required)

### `Update`

`PUT /v1/reception/customer-facing-agents/{agent_id}/rules/{rule_id}`

Update Rule

**Parameters:**

- `agentId` (required)
- `ruleId` (required)

### `Delete`

`DELETE /v1/reception/customer-facing-agents/{agent_id}/rules/{rule_id}`

Delete Rule

**Parameters:**

- `agentId` (required)
- `ruleId` (required)

## ApiKeys

### `Disable`

`POST /v1/workspaces/api-keys/disable`

Disable API key

**Parameters:**

- `apiKeyName` (required)

### `List`

`GET /v1/service-accounts/{service_account_user_id}/api-keys`

Get API keys

**Parameters:**

- `serviceAccountUserId` (required)

### `Create`

`POST /v1/service-accounts/{service_account_user_id}/api-keys`

Create API key

**Parameters:**

- `serviceAccountUserId` (required)

### `Update`

`PATCH /v1/service-accounts/{service_account_user_id}/api-keys/{api_key_id}`

Update API key

**Parameters:**

- `serviceAccountUserId` (required)
- `apiKeyId` (required)

### `Delete`

`DELETE /v1/service-accounts/{service_account_user_id}/api-keys/{api_key_id}`

Delete API key

**Parameters:**

- `serviceAccountUserId` (required)
- `apiKeyId` (required)

### `Disable1`

`POST /v1/service-accounts/{service_account_user_id}/api-keys/{api_key_id}/disable`

Disable Service Account Api Key

**Parameters:**

- `serviceAccountUserId` (required)
- `apiKeyId` (required)

## Workspace

### `SetThirdPartyDisablingPolicy`

`POST /v1/workspaces/api-keys/third-party-disabling`

Set Workspace Third-Party Disabling Policy

## AuditLogs

### `List`

`GET /v1/workspace/audit-logs`

Get Workspace Audit Logs

**Parameters:**

- `limit`
- `cursor`
- `timeFromUnixMs`
- `timeToUnixMs`
- `actorUid`
- `className`
- `activityName`

## AuthConnections

### `List`

`GET /v1/workspace/auth-connections`

Get Workspace Auth Connections

### `Create`

`POST /v1/workspace/auth-connections`

Create Workspace Auth Connection

### `Update`

`PATCH /v1/workspace/auth-connections/{auth_connection_id}`

Update Workspace Auth Connection

**Parameters:**

- `authConnectionId` (required)

### `Delete`

`DELETE /v1/workspace/auth-connections/{auth_connection_id}`

Delete Workspace Auth Connection

**Parameters:**

- `authConnectionId` (required)

## ServiceAccounts

### `List`

`GET /v1/service-accounts`

Get service accounts

### `Create`

`POST /v1/service-accounts`

Create Service Account

## Groups

### `List`

`GET /v1/workspace/groups`

List workspace groups

### `Search`

`GET /v1/workspace/groups/search`

Search user group

**Parameters:**

- `name` (required)

## Members

### `Remove`

`POST /v1/workspace/groups/{group_id}/members/remove`

Remove member from user group

**Parameters:**

- `groupId` (required)

### `Add`

`POST /v1/workspace/groups/{group_id}/members`

Add member to user group

**Parameters:**

- `groupId` (required)

### `List`

`GET /v1/workspace/members`

Get Workspace Members

### `Update`

`POST /v1/workspace/members`

Update member

### `UpdateUsageLimit`

`PUT /v1/workspace/members/{user_id}/usage-limit`

Update Member Usage Limit

**Parameters:**

- `userId` (required)

### `List1`

`GET /v1/reception/team/members`

Get Team Members

### `UpdateRole`

`PATCH /v1/reception/team/members/{user_id}`

Update Team Member Role

**Parameters:**

- `userId` (required)

### `Remove1`

`DELETE /v1/reception/team/members/{user_id}`

Remove Team Member

**Parameters:**

- `userId` (required)

## Invites

### `Create`

`POST /v1/workspace/invites/add`

Invite user

### `CreateBatch`

`POST /v1/workspace/invites/add-bulk`

Invite Multiple Users

### `Delete`

`DELETE /v1/workspace/invites`

Delete invite

### `List`

`GET /v1/reception/team/invites`

Get Pending Team Invites

### `Create1`

`POST /v1/reception/team/invites`

Invite Team Member

### `Cancel`

`DELETE /v1/reception/team/invites`

Cancel Team Invite

## Resources

### `Get`

`GET /v1/workspace/resources/{resource_id}`

Get Resource

**Parameters:**

- `resourceId` (required)
- `resourceType` (required)

### `Share`

`POST /v1/workspace/resources/{resource_id}/share`

Share Workspace Resource

**Parameters:**

- `resourceId` (required)

### `Unshare`

`POST /v1/workspace/resources/{resource_id}/unshare`

Unshare Workspace Resource

**Parameters:**

- `resourceId` (required)

## Webhooks

### `List`

`GET /v1/workspace/webhooks`

List Workspace Webhooks

**Parameters:**

- `includeUsages`

### `Create`

`POST /v1/workspace/webhooks`

Create Workspace Webhook

### `Update`

`PATCH /v1/workspace/webhooks/{webhook_id}`

Update Workspace Webhook

**Parameters:**

- `webhookId` (required)

### `Delete`

`DELETE /v1/workspace/webhooks/{webhook_id}`

Delete Workspace Webhook

**Parameters:**

- `webhookId` (required)

### `Test`

`POST /v1/workspace/webhooks/{webhook_id}/test`

Send Test Webhook Event

**Parameters:**

- `webhookId` (required)

### `Redeliver`

`POST /v1/workspace/webhooks/{webhook_id}/deliveries/{delivery_id}/redeliver`

Redeliver Webhook Delivery

**Parameters:**

- `webhookId` (required)
- `deliveryId` (required)

## Music

### `VideoToMusic`

`POST /v1/music/video-to-music`

Video To Music

**Parameters:**

- `outputFormat`
- `enableLogging`

### `Compose`

`POST /v1/music`

Compose Music

**Parameters:**

- `outputFormat`
- `enableLogging`

### `ComposeDetailed`

`POST /v1/music/detailed`

Compose Music With A Detailed Response

**Parameters:**

- `outputFormat`
- `enableLogging`

### `ComposeDetailedStream`

`POST /v1/music/detailed/stream`

Stream Composed Music With A Detailed Response

**Parameters:**

- `outputFormat`
- `enableLogging`

### `Stream`

`POST /v1/music/stream`

Stream Composed Music

**Parameters:**

- `outputFormat`
- `enableLogging`

### `Upload`

`POST /v1/music/upload`

Upload Music

### `SeparateStems`

`POST /v1/music/stem-separation`

Stem Separation

**Parameters:**

- `outputFormat`
- `enableLogging`

## SpeechToText

### `Convert`

`POST /v1/speech-to-text`

Create transcript

**Parameters:**

- `token`
- `enableLogging`

## SingleUse

### `Create`

`POST /v1/single-use-token/{token_type}`

Create Single Use Token

**Parameters:**

- `tokenType` (required)
- `maxSessionDurationSeconds`

## ForcedAlignment

### `Create`

`POST /v1/forced-alignment`

Create Forced Alignment

## Conversations

### `GetSignedUrl`

`GET /v1/convai/conversation/get-signed-url`

Get Signed Url

**Parameters:**

- `agentId` (required)
- `includeConversationId`
- `branchId`
- `versionId`
- `environment`
- `debugEventsRequest`

### `GetWebrtcToken`

`GET /v1/convai/conversation/token`

Get a webrtc token to start a conversation with an agent that requires authorization

**Parameters:**

- `agentId` (required)
- `participantName`
- `branchId`
- `versionId`
- `environment`
- `debugEventsRequest`

### `List`

`GET /v1/convai/conversations`

List conversations

**Parameters:**

- `cursor`
- `agentId`
- `visitedAgentIds`
- `visitedAgentBranchIds`
- `triggeredProcedureIds`
- `callSuccessful`
- `callStartBeforeUnix`
- `callStartAfterUnix`
- `callDurationMinSecs`
- `callDurationMaxSecs`
- `ratingMax`
- `ratingMin`
- `hasFeedbackComment`
- `userId`
- `evaluationParams`
- `dataCollectionParams`
- `dynamicVariableParams`
- `dataCollectionFilter`
- `dynamicVariableFilter`
- `dataCollectionIds`
- `evaluationCriteriaIds`
- `toolNames`
- `toolNamesSuccessful`
- `toolNamesErrored`
- `toolIds`
- `toolIdsSuccessful`
- `toolIdsErrored`
- `includeInvalidToolCalls`
- `mainLanguages`
- `phoneNumberIds`
- `whatsappPhoneNumberIds`
- `pageSize`
- `summaryMode`
- `search`
- `conversationInitiationSource`
- `textOnly`
- `conversationProductType`
- `branchId`
- `versionId`
- `parentConversationId`
- `topicIds`
- `excludeStatuses`
- `tagIds`
- `workflowNodeEnteredId`
- `terminationReasons`
- `guardrailTypes`
- `customGuardrailNames`
- `sortDirection`

### `Create`

`POST /v1/convai/conversations`

Create Conversation

### `Resolve`

`GET /v1/convai/conversations/resolve`

Resolve Conversation Reference

**Parameters:**

- `agentId` (required)
- `reference` (required)

### `Get`

`GET /v1/convai/conversations/{conversation_id}`

Get Conversation Details

**Parameters:**

- `conversationId` (required)
- `format`

### `Delete`

`DELETE /v1/convai/conversations/{conversation_id}`

Delete Conversation

**Parameters:**

- `conversationId` (required)

### `GetSummary`

`GET /v1/convai/conversations/{conversation_id}/summary`

Get Conversation Summary

**Parameters:**

- `conversationId` (required)
- `maxMessages`

### `GetSipMessages`

`GET /v1/convai/conversations/{conversation_id}/sip-messages`

Get Sip Messages For A Conversation

**Parameters:**

- `conversationId` (required)
- `pageSize`
- `cursor`

### `SendMessage`

`POST /v1/convai/conversations/{conversation_id}/messages`

Send Conversation Message

**Parameters:**

- `conversationId` (required)

### `SendMessageStream`

`POST /v1/convai/conversations/{conversation_id}/messages/stream`

Send Conversation Message Stream

**Parameters:**

- `conversationId` (required)

### `GetTurn`

`GET /v1/convai/conversations/{conversation_id}/turns/{turn_id}`

Get Conversation Turn

**Parameters:**

- `conversationId` (required)
- `turnId` (required)

### `StreamTurnEvents`

`GET /v1/convai/conversations/{conversation_id}/turns/{turn_id}/events`

Stream Conversation Turn Events

**Parameters:**

- `conversationId` (required)
- `turnId` (required)

## Twilio

### `OutboundCall`

`POST /v1/convai/twilio/outbound-call`

Handle An Outbound Call Via Twilio

### `RegisterCall`

`POST /v1/convai/twilio/register-call`

Register A Twilio Call And Return Twiml

## Exotel

### `OutboundCall`

`POST /v1/convai/exotel/outbound-call`

Handle An Outbound Call Via Exotel

## Whatsapp

### `OutboundCall`

`POST /v1/convai/whatsapp/outbound-call`

Make An Outbound Call Via Whatsapp

### `OutboundMessage`

`POST /v1/convai/whatsapp/outbound-message`

Send An Outbound Message Via Whatsapp

## Agents

### `Create`

`POST /v1/convai/agents/create`

Create agent

**Parameters:**

- `enableVersioning`

### `Get`

`GET /v1/convai/agents/{agent_id}`

Get agent

**Parameters:**

- `agentId` (required)
- `versionId`
- `branchId`
- `includeProcedureGeneratedWorkflowElements`

### `Update`

`PATCH /v1/convai/agents/{agent_id}`

Update agent

**Parameters:**

- `agentId` (required)
- `enableVersioningIfNotEnabled`
- `branchId`
- `includeProcedureGeneratedWorkflowElements`

### `Delete`

`DELETE /v1/convai/agents/{agent_id}`

Delete agent

**Parameters:**

- `agentId` (required)
- `force`

### `List`

`GET /v1/convai/agents`

List Agents

**Parameters:**

- `pageSize`
- `search`
- `archived`
- `showOnlyOwnedAgents`
- `createdByUserId`
- `tags`
- `sortDirection`
- `sortBy`
- `cursor`

### `Duplicate`

`POST /v1/convai/agents/{agent_id}/duplicate`

Duplicate Agent

**Parameters:**

- `agentId` (required)

### `RunTests`

`POST /v1/convai/agents/{agent_id}/run-tests`

Run Tests On The Agent

**Parameters:**

- `agentId` (required)

### `RagIndexOverview`

`GET /v1/convai/knowledge-base/rag-index`

Get Rag Index Overview.

### `GetDocumentRagIndexes`

`GET /v1/convai/knowledge-base/{documentation_id}/rag-index`

Get Rag Indexes Of The Specified Knowledgebase Document.

**Parameters:**

- `documentationId` (required)

### `DeleteDocumentRagIndex`

`DELETE /v1/convai/knowledge-base/{documentation_id}/rag-index/{rag_index_id}`

Delete Rag Index.

**Parameters:**

- `documentationId` (required)
- `ragIndexId` (required)

## Summaries

### `Get`

`GET /v1/convai/agents/summaries`

Get Agent Summaries

**Parameters:**

- `agentIds` (required)

### `Get1`

`GET /v1/convai/knowledge-base/summaries`

Get Knowledge Base Summaries By Ids

**Parameters:**

- `documentIds` (required)

## Subagents

### `GetDependents`

`GET /v1/convai/agents/{agent_id}/subagent-dependents`

Get Agent Subagent Dependents

**Parameters:**

- `agentId` (required)
- `pageSize`
- `cursor`

## Widget

### `Get`

`GET /v1/convai/agents/{agent_id}/widget`

Get Agent Widget Config

**Parameters:**

- `agentId` (required)
- `conversationSignature`

## Link

### `Get`

`GET /v1/convai/agents/{agent_id}/link`

Get Shareable Agent Link

**Parameters:**

- `agentId` (required)

## Avatar

### `Create`

`POST /v1/convai/agents/{agent_id}/avatar`

Post Agent Avatar

**Parameters:**

- `agentId` (required)

## HoldAudio

### `Create`

`POST /v1/convai/agents/{agent_id}/hold-audio`

Post Agent Hold Audio

**Parameters:**

- `agentId` (required)

### `Delete`

`DELETE /v1/convai/agents/{agent_id}/hold-audio`

Delete Agent Hold Audio

**Parameters:**

- `agentId` (required)

## KnowledgeBase

### `Size`

`GET /v1/convai/agent/{agent_id}/knowledge-base/size`

Returns The Size Of The Agent'S Knowledge Base

**Parameters:**

- `agentId` (required)

### `List`

`GET /v1/convai/knowledge-base`

Get Knowledge Base List

**Parameters:**

- `pageSize`
- `search`
- `showOnlyOwnedDocuments`
- `createdByUserId`
- `types`
- `parentFolderId`
- `ancestorFolderId`
- `foldersFirst`
- `sortDirection`
- `sortBy`
- `cursor`

### `GetOrCreateRagIndexes`

`POST /v1/convai/knowledge-base/rag-index`

Compute Rag Indexes In Batch

### `Search`

`GET /v1/convai/knowledge-base/search`

Search Knowledge Base Content

**Parameters:**

- `query` (required)
- `pageSize`
- `types`
- `cursor`

## LlmUsage

### `CalculateForAgent`

`POST /v1/convai/agent/{agent_id}/llm-usage/calculate`

Calculate Expected Llm Usage For An Agent

**Parameters:**

- `agentId` (required)

### `Calculate`

`POST /v1/convai/llm-usage/calculate`

Calculate Expected Llm Usage

## Tests

### `Create`

`POST /v1/convai/agent-testing/create`

Create Agent Response Test

### `Move`

`POST /v1/convai/agent-testing/bulk-move`

Bulk Move Tests To Folder

### `Get`

`GET /v1/convai/agent-testing/{test_id}`

Get Agent Response Test By Id

**Parameters:**

- `testId` (required)

### `Update`

`PUT /v1/convai/agent-testing/{test_id}`

Update Agent Response Test

**Parameters:**

- `testId` (required)

### `Delete`

`DELETE /v1/convai/agent-testing/{test_id}`

Delete Agent Response Test

**Parameters:**

- `testId` (required)

### `Summaries`

`POST /v1/convai/agent-testing/summaries`

Get Agent Response Test Summaries By Ids

### `List`

`GET /v1/convai/agent-testing`

List Agent Response Tests

**Parameters:**

- `cursor`
- `pageSize`
- `search`
- `parentFolderId`
- `types`
- `includeFolders`
- `sortMode`
- `sharingMode`
- `createdByUserId`

## Folders

### `Create`

`POST /v1/convai/agent-testing/folders`

Create Agent Test Folder

### `Get`

`GET /v1/convai/agent-testing/folders/{folder_id}`

Get Agent Test Folder By Id

**Parameters:**

- `folderId` (required)

### `Update`

`PATCH /v1/convai/agent-testing/folders/{folder_id}`

Update Agent Test Folder

**Parameters:**

- `folderId` (required)

### `Delete`

`DELETE /v1/convai/agent-testing/folders/{folder_id}`

Delete Agent Test Folder

**Parameters:**

- `folderId` (required)
- `force`

## Invocations

### `List`

`GET /v1/convai/test-invocations`

List Test Invocations

**Parameters:**

- `agentId`
- `branchId`
- `createdByUserId`
- `testId`
- `startTime`
- `endTime`
- `versionId`
- `pageSize`
- `search`
- `cursor`

### `Get`

`GET /v1/convai/test-invocations/{test_invocation_id}`

Get Test Invocation

**Parameters:**

- `testInvocationId` (required)

### `Update`

`PATCH /v1/convai/test-invocations/{test_invocation_id}`

Update Test Invocation

**Parameters:**

- `testInvocationId` (required)

### `Cancel`

`POST /v1/convai/test-invocations/{test_invocation_id}/cancel`

Cancel Test Invocation

**Parameters:**

- `testInvocationId` (required)

### `Resubmit`

`POST /v1/convai/test-invocations/{test_invocation_id}/resubmit`

Resubmit Tests

**Parameters:**

- `testInvocationId` (required)

## Users

### `List`

`GET /v1/convai/users`

Get Conversation Users

**Parameters:**

- `agentId`
- `branchId`
- `callStartBeforeUnix`
- `callStartAfterUnix`
- `search`
- `pageSize`
- `sortBy`
- `sortDirection`
- `cursor`

## Files

### `Get`

`GET /v1/convai/conversations/{conversation_id}/files/{file_id}`

Get Conversation File

**Parameters:**

- `conversationId` (required)
- `fileId` (required)

### `Delete`

`DELETE /v1/convai/conversations/{conversation_id}/files/{file_id}`

Delete File Upload

**Parameters:**

- `fileId` (required)
- `conversationId` (required)

### `Create`

`POST /v1/convai/conversations/{conversation_id}/files`

Upload File

**Parameters:**

- `conversationId` (required)

## Feedback

### `Create`

`POST /v1/convai/conversations/{conversation_id}/feedback`

Send Conversation Feedback

**Parameters:**

- `conversationId` (required)

## Messages

### `TextSearch`

`GET /v1/convai/conversations/messages/text-search`

Text Search Conversation Messages

**Parameters:**

- `textQuery` (required)
- `agentId`
- `visitedAgentIds`
- `visitedAgentBranchIds`
- `triggeredProcedureIds`
- `callSuccessful`
- `callStartBeforeUnix`
- `callStartAfterUnix`
- `callDurationMinSecs`
- `callDurationMaxSecs`
- `ratingMax`
- `ratingMin`
- `hasFeedbackComment`
- `userId`
- `evaluationParams`
- `dataCollectionParams`
- `dynamicVariableParams`
- `toolNames`
- `toolNamesSuccessful`
- `toolNamesErrored`
- `toolIds`
- `toolIdsSuccessful`
- `toolIdsErrored`
- `includeInvalidToolCalls`
- `mainLanguages`
- `phoneNumberIds`
- `whatsappPhoneNumberIds`
- `excludeStatuses`
- `terminationReasons`
- `pageSize`
- `summaryMode`
- `conversationInitiationSource`
- `textOnly`
- `conversationProductType`
- `branchId`
- `versionId`
- `topicIds`
- `sortBy`
- `cursor`

### `Search`

`GET /v1/convai/conversations/messages/smart-search`

Smart Search Conversation Messages

**Parameters:**

- `textQuery` (required)
- `agentId`
- `branchId`
- `pageSize`
- `cursor`

## Tags

### `Assign`

`POST /v1/convai/conversations/{conversation_id}/tags`

Assign Conversation Tags

**Parameters:**

- `conversationId` (required)

### `Unassign`

`DELETE /v1/convai/conversations/{conversation_id}/tags/{tag_id}`

Unassign Conversation Tag

**Parameters:**

- `conversationId` (required)
- `tagId` (required)

### `List`

`GET /v1/convai/tags`

List Conversation Tags

**Parameters:**

- `pageSize`
- `cursor`

### `Create`

`POST /v1/convai/tags`

Create Conversation Tag

### `Get`

`GET /v1/convai/tags/{tag_id}`

Get Conversation Tag

**Parameters:**

- `tagId` (required)

### `Update`

`PATCH /v1/convai/tags/{tag_id}`

Update Conversation Tag

**Parameters:**

- `tagId` (required)

### `Delete`

`DELETE /v1/convai/tags/{tag_id}`

Delete Conversation Tag

**Parameters:**

- `tagId` (required)

## TriageTickets

### `List`

`GET /v1/convai/agents/{agent_id}/triage-tickets`

List Agent Conversation Tickets

**Parameters:**

- `agentId` (required)
- `pageSize`
- `conversationId`
- `status`
- `sources`
- `priorities`
- `sortBy`
- `ownerUserId`
- `assigneeUserId`
- `label`
- `mergedIntoTicketId`
- `search`
- `cursor`

### `CreateManual`

`POST /v1/convai/agents/{agent_id}/triage-tickets`

Create Manual Agent Ticket

**Parameters:**

- `agentId` (required)

### `ListForWorkspace`

`GET /v1/convai/triage-tickets`

List Workspace Conversation Tickets

**Parameters:**

- `pageSize`
- `agentId`
- `status`
- `sources`
- `priorities`
- `sortBy`
- `ownerUserId`
- `assigneeUserId`
- `label`
- `mergedIntoTicketId`
- `search`
- `cursor`

### `Create`

`POST /v1/convai/triage-tickets`

Create Agent Conversation Ticket

### `ListAssignableUsers`

`GET /v1/convai/agents/{agent_id}/triage-tickets/assignable-users`

Get Agent Conversation Ticket Assignable Users

**Parameters:**

- `agentId` (required)

### `Get`

`GET /v1/convai/triage-tickets/{agentqa_ticket_id}`

Get Agent Conversation Ticket

**Parameters:**

- `agentqaTicketId` (required)

### `Update`

`PATCH /v1/convai/triage-tickets/{agentqa_ticket_id}`

Update Agent Conversation Ticket

**Parameters:**

- `agentqaTicketId` (required)

### `Delete`

`DELETE /v1/convai/triage-tickets/{agentqa_ticket_id}`

Delete Agent Conversation Ticket

**Parameters:**

- `agentqaTicketId` (required)

### `AddComment`

`POST /v1/convai/triage-tickets/{agentqa_ticket_id}/comments`

Add Comment To Agent Conversation Ticket

**Parameters:**

- `agentqaTicketId` (required)

### `Merge`

`POST /v1/convai/triage-tickets/{agentqa_ticket_id}/merge`

Merge Agent Conversation Tickets

**Parameters:**

- `agentqaTicketId` (required)

### `AddTurnComment`

`POST /v1/convai/triage-tickets/{agentqa_ticket_id}/turn-comments`

Add Turn Comment To Agent Conversation Ticket

**Parameters:**

- `agentqaTicketId` (required)

## PhoneNumbers

### `List`

`GET /v1/convai/phone-numbers`

List Phone Numbers

**Parameters:**

- `provider`
- `agentId`
- `branchId`

### `Create`

`POST /v1/convai/phone-numbers`

Import Phone Number

### `Get`

`GET /v1/convai/phone-numbers/{phone_number_id}`

Get Phone Number

**Parameters:**

- `phoneNumberId` (required)

### `Update`

`PATCH /v1/convai/phone-numbers/{phone_number_id}`

Update Phone Number

**Parameters:**

- `phoneNumberId` (required)

### `Delete`

`DELETE /v1/convai/phone-numbers/{phone_number_id}`

Delete Phone Number

**Parameters:**

- `phoneNumberId` (required)

### `ListV2`

`GET /v1/convai/v2/phone-numbers`

List Phone Numbers Page

**Parameters:**

- `pageSize`
- `search`
- `label`
- `phoneNumber`
- `provider`
- `supportsOutbound`
- `agentId`
- `branchId`
- `sortBy`
- `sortDirection`
- `cursor`

### `GetSipMessages`

`GET /v1/convai/phone-numbers/{phone_number_id}/sip-messages`

Get Sip Messages For A Phone Number

**Parameters:**

- `phoneNumberId` (required)
- `pageSize`
- `cursor`

### `List1`

`GET /v1/reception/phone-numbers/assignments`

List Phone Numbers

### `Create1`

`POST /v1/reception/phone-numbers/assignments`

Assign Phone Number

### `Update1`

`PATCH /v1/reception/phone-numbers/assignments/{phone_assignment_id}`

Move Phone Number

**Parameters:**

- `phoneAssignmentId` (required)

### `Delete1`

`DELETE /v1/reception/phone-numbers/assignments/{phone_assignment_id}`

Release Phone Number

**Parameters:**

- `phoneAssignmentId` (required)

## Llm

### `List`

`GET /v1/convai/llm/list`

List Available Llms

## LiveCount

### `Get`

`GET /v1/convai/analytics/live-count`

Get Live Count

**Parameters:**

- `agentId`
- `agentIds`

## Documents

### `CreateFromUrl`

`POST /v1/convai/knowledge-base/url`

Create Url Document

### `CreateFromFile`

`POST /v1/convai/knowledge-base/file`

Create File Document

### `CreateFromText`

`POST /v1/convai/knowledge-base/text`

Create Text Document

### `CreateFolder`

`POST /v1/convai/knowledge-base/folder`

Create Folder

### `Get`

`GET /v1/convai/knowledge-base/{documentation_id}`

Get Documentation From Knowledge Base

**Parameters:**

- `documentationId` (required)
- `agentId`

### `Update`

`PATCH /v1/convai/knowledge-base/{documentation_id}`

Update Document

**Parameters:**

- `documentationId` (required)

### `Delete`

`DELETE /v1/convai/knowledge-base/{documentation_id}`

Delete Knowledge Base Document Or Folder

**Parameters:**

- `documentationId` (required)
- `force`

### `GetChunkingOverride`

`GET /v1/convai/knowledge-base/{documentation_id}/chunking-override`

Get Chunking Override

**Parameters:**

- `documentationId` (required)

### `UpdateChunkingOverride`

`PATCH /v1/convai/knowledge-base/{documentation_id}/chunking-override`

Update Chunking Override

**Parameters:**

- `documentationId` (required)

### `GetAgents`

`GET /v1/convai/knowledge-base/{documentation_id}/dependent-agents`

Get Dependent Agents List

**Parameters:**

- `documentationId` (required)
- `dependentType`
- `pageSize`
- `cursor`

### `GetBulkAgents`

`POST /v1/convai/knowledge-base/dependent-agents`

Get Dependent Agents For Multiple Documents

**Parameters:**

- `dependentType`
- `pageSize`
- `cursor`

### `GetContent`

`GET /v1/convai/knowledge-base/{documentation_id}/content`

Get Document Content

**Parameters:**

- `documentationId` (required)

### `GetSourceFileUrl`

`GET /v1/convai/knowledge-base/{documentation_id}/source-file-url`

Get Document Source File Url

**Parameters:**

- `documentationId` (required)

### `Move`

`POST /v1/convai/knowledge-base/{document_id}/move`

Move Entity To Folder

**Parameters:**

- `documentId` (required)

### `BulkMove`

`POST /v1/convai/knowledge-base/bulk-move`

Bulk Move Entities To Folder

### `BulkDelete`

`POST /v1/convai/knowledge-base/bulk-delete`

Bulk Delete Knowledge Base Documents

## CrawlJobs

### `List`

`GET /v1/convai/knowledge-base/crawl`

List Ongoing And Recent Crawl Jobs Created By A User

**Parameters:**

- `includeJobIds`
- `pageSize`
- `cursor`

### `Create`

`POST /v1/convai/knowledge-base/crawl`

Create Crawl Job

### `Get`

`GET /v1/convai/knowledge-base/crawl/{crawl_job_id}`

Get Crawl Job Details

**Parameters:**

- `crawlJobId` (required)

### `Cancel`

`POST /v1/convai/knowledge-base/crawl/{crawl_job_id}/cancel`

Cancel Crawl Job

**Parameters:**

- `crawlJobId` (required)

## Document

### `UpdateFile`

`PATCH /v1/convai/knowledge-base/{documentation_id}/update-file`

Update File Document

**Parameters:**

- `documentationId` (required)

### `Refresh`

`POST /v1/convai/knowledge-base/{documentation_id}/refresh`

Refresh Url Document Content

**Parameters:**

- `documentationId` (required)

### `ComputeRagIndex`

`POST /v1/convai/knowledge-base/{documentation_id}/rag-index`

Compute Rag Index.

**Parameters:**

- `documentationId` (required)

## Chunk

### `Get`

`GET /v1/convai/knowledge-base/{documentation_id}/chunk/{chunk_id}`

Get Documentation Chunk From Knowledge Base

**Parameters:**

- `documentationId` (required)
- `chunkId` (required)
- `embeddingModel`

## Chunks

### `List`

`GET /v1/convai/knowledge-base/{documentation_id}/chunks`

Get All Rag Chunks For A Document

**Parameters:**

- `documentationId` (required)
- `embeddingModel` (required)
- `pageSize`
- `cursor`

## Topics

### `Get`

`GET /v1/convai/agents/{agent_id}/topics`

Get Agent Conversation Topics

**Parameters:**

- `agentId` (required)
- `pageSize`
- `sortBy`
- `sortDirection`
- `fromUnixSecs`
- `toUnixSecs`
- `includeEvaluationCriteria`
- `cursor`

## Tools

### `List`

`GET /v1/convai/tools`

Get Tools

**Parameters:**

- `search`
- `pageSize`
- `showOnlyOwnedDocuments`
- `createdByUserId`
- `types`
- `sortDirection`
- `sortBy`
- `usedByAgentId`
- `cursor`

### `Create`

`POST /v1/convai/tools`

Add Tool

### `Get`

`GET /v1/convai/tools/{tool_id}`

Get Tool

**Parameters:**

- `toolId` (required)
- `environment`

### `Update`

`PATCH /v1/convai/tools/{tool_id}`

Update Tool

**Parameters:**

- `toolId` (required)

### `Delete`

`DELETE /v1/convai/tools/{tool_id}`

Delete Tool

**Parameters:**

- `toolId` (required)
- `force`

### `GetDependentAgents`

`GET /v1/convai/tools/{tool_id}/dependent-agents`

Get Dependent Agents List

**Parameters:**

- `toolId` (required)
- `cursor`
- `pageSize`

### `List1`

`GET /v1/convai/mcp-servers/{mcp_server_id}/tools`

List Mcp Server Tools

**Parameters:**

- `mcpServerId` (required)
- `environment`

## Executions

### `Get`

`GET /v1/convai/tools/{tool_id}/executions`

Get Tool Executions

**Parameters:**

- `toolId` (required)
- `cursor`
- `pageSize`
- `isError`
- `agentId`
- `branchId`
- `startTime`
- `endTime`

## AnalysisItems

### `List`

`GET /v1/convai/analysis-items`

List Analysis Items

**Parameters:**

- `search`
- `pageSize`
- `kind`
- `analysisItemIds`
- `createdByUserId`
- `sortDirection`
- `cursor`

### `Create`

`POST /v1/convai/analysis-items`

Create Analysis Item

### `UpdateAttachments`

`PATCH /v1/convai/analysis-items/{analysis_item_id}/attachments`

Update Analysis Item Attachments

**Parameters:**

- `analysisItemId` (required)

### `Get`

`GET /v1/convai/analysis-items/{analysis_item_id}`

Get Analysis Item

**Parameters:**

- `analysisItemId` (required)

### `Update`

`PATCH /v1/convai/analysis-items/{analysis_item_id}`

Update Analysis Item

**Parameters:**

- `analysisItemId` (required)

### `Delete`

`DELETE /v1/convai/analysis-items/{analysis_item_id}`

Delete Analysis Item

**Parameters:**

- `analysisItemId` (required)
- `force`

### `ListVersions`

`GET /v1/convai/analysis-items/{analysis_item_id}/versions`

List Analysis Item Versions

**Parameters:**

- `analysisItemId` (required)
- `pageSize`
- `cursor`

### `Copy`

`POST /v1/convai/analysis-items/{analysis_item_id}/copy`

Copy Analysis Item

**Parameters:**

- `analysisItemId` (required)

## DependentAgents

### `Get`

`GET /v1/convai/analysis-items/{analysis_item_id}/dependent-agents`

Get Analysis Item Dependent Agents

**Parameters:**

- `analysisItemId` (required)
- `pageSize`
- `cursor`

## Secrets

### `List`

`GET /v1/convai/secrets`

Get Convai Workspace Secrets

**Parameters:**

- `pageSize`
- `dependencyLimit`
- `search`
- `cursor`

### `Create`

`POST /v1/convai/secrets`

Create Convai Workspace Secret

### `Get`

`GET /v1/convai/secrets/{secret_id}`

Get Convai Workspace Secret

**Parameters:**

- `secretId` (required)

### `Update`

`PATCH /v1/convai/secrets/{secret_id}`

Update Convai Workspace Secret

**Parameters:**

- `secretId` (required)

### `Delete`

`DELETE /v1/convai/secrets/{secret_id}`

Delete Convai Workspace Secret

**Parameters:**

- `secretId` (required)

### `GetDependencies`

`GET /v1/convai/secrets/{secret_id}/dependencies/{resource_type}`

Get Secret Dependencies By Type

**Parameters:**

- `secretId` (required)
- `resourceType` (required)
- `pageSize`
- `cursor`

## BatchCalls

### `Create`

`POST /v1/convai/batch-calling/submit`

Submit A Batch Call Request.

### `List`

`GET /v1/convai/batch-calling/workspace`

Get All Batch Calls For A Workspace.

**Parameters:**

- `limit`
- `lastDoc`
- `agentId`

### `Get`

`GET /v1/convai/batch-calling/{batch_id}`

Get A Batch Call By Id.

**Parameters:**

- `batchId` (required)

### `Delete`

`DELETE /v1/convai/batch-calling/{batch_id}`

Delete A Batch Call.

**Parameters:**

- `batchId` (required)

### `Cancel`

`POST /v1/convai/batch-calling/{batch_id}/cancel`

Cancel A Batch Call.

**Parameters:**

- `batchId` (required)

### `Retry`

`POST /v1/convai/batch-calling/{batch_id}/retry`

Retry A Batch Call.

**Parameters:**

- `batchId` (required)

### `Export`

`GET /v1/convai/batch-calling/{batch_id}/export`

Export Batch Call Results

**Parameters:**

- `batchId` (required)
- `limit`

## SipTrunk

### `OutboundCall`

`POST /v1/convai/sip-trunk/outbound-call`

Handle An Outbound Call Via Sip Trunk

## McpServers

### `List`

`GET /v1/convai/mcp-servers`

List Mcp Servers

### `Create`

`POST /v1/convai/mcp-servers`

Create Mcp Server

### `Get`

`GET /v1/convai/mcp-servers/{mcp_server_id}`

Get Mcp Server

**Parameters:**

- `mcpServerId` (required)

### `Update`

`PATCH /v1/convai/mcp-servers/{mcp_server_id}`

Update Mcp Server Configuration

**Parameters:**

- `mcpServerId` (required)

### `Delete`

`DELETE /v1/convai/mcp-servers/{mcp_server_id}`

Delete Mcp Server

**Parameters:**

- `mcpServerId` (required)

## ToolApprovals

### `Create`

`POST /v1/convai/mcp-servers/{mcp_server_id}/tool-approvals`

Create Mcp Server Tool Approval

**Parameters:**

- `mcpServerId` (required)

### `Delete`

`DELETE /v1/convai/mcp-servers/{mcp_server_id}/tool-approvals/{tool_name}`

Delete Mcp Server Tool Approval

**Parameters:**

- `mcpServerId` (required)
- `toolName` (required)

## ToolConfigs

### `Create`

`POST /v1/convai/mcp-servers/{mcp_server_id}/tool-configs`

Create Mcp Tool Configuration Override

**Parameters:**

- `mcpServerId` (required)
- `environment`

### `Get`

`GET /v1/convai/mcp-servers/{mcp_server_id}/tool-configs/{tool_name}`

Get Mcp Tool Configuration Override

**Parameters:**

- `mcpServerId` (required)
- `toolName` (required)

### `Update`

`PATCH /v1/convai/mcp-servers/{mcp_server_id}/tool-configs/{tool_name}`

Update Mcp Tool Configuration Override

**Parameters:**

- `mcpServerId` (required)
- `toolName` (required)
- `environment`

### `Delete`

`DELETE /v1/convai/mcp-servers/{mcp_server_id}/tool-configs/{tool_name}`

Delete Mcp Tool Configuration Override

**Parameters:**

- `mcpServerId` (required)
- `toolName` (required)

## WhatsappAccounts

### `Get`

`GET /v1/convai/whatsapp-accounts/{phone_number_id}`

Get Whatsapp Account

**Parameters:**

- `phoneNumberId` (required)

### `Update`

`PATCH /v1/convai/whatsapp-accounts/{phone_number_id}`

Update Whatsapp Account

**Parameters:**

- `phoneNumberId` (required)

### `Delete`

`DELETE /v1/convai/whatsapp-accounts/{phone_number_id}`

Delete Whatsapp Account

**Parameters:**

- `phoneNumberId` (required)

### `List`

`GET /v1/convai/whatsapp-accounts`

List Whatsapp Accounts

**Parameters:**

- `agentId`

## Branches

### `List`

`GET /v1/convai/agents/{agent_id}/branches`

List Agent Branches

**Parameters:**

- `agentId` (required)
- `includeArchived`
- `limit`
- `includeCommitStatus`

### `Create`

`POST /v1/convai/agents/{agent_id}/branches`

Create A New Branch

**Parameters:**

- `agentId` (required)

### `Get`

`GET /v1/convai/agents/{agent_id}/branches/{branch_id}`

Get Agent Branch

**Parameters:**

- `agentId` (required)
- `branchId` (required)

### `Update`

`PATCH /v1/convai/agents/{agent_id}/branches/{branch_id}`

Update Agent Branch

**Parameters:**

- `agentId` (required)
- `branchId` (required)

### `PreviewMerge`

`GET /v1/convai/agents/{agent_id}/branches/{source_branch_id}/merge-preview`

Preview Merged Configuration

**Parameters:**

- `agentId` (required)
- `sourceBranchId` (required)
- `targetBranchId` (required)
- `force`
- `includeProcedureGeneratedWorkflowElements`

### `Merge`

`POST /v1/convai/agents/{agent_id}/branches/{source_branch_id}/merge`

Merge A Branch Into A Target Branch

**Parameters:**

- `agentId` (required)
- `sourceBranchId` (required)
- `targetBranchId` (required)

### `PreviewRebase`

`GET /v1/convai/agents/{agent_id}/branches/{branch_id}/rebase-preview`

Preview Rebased Configuration

**Parameters:**

- `agentId` (required)
- `branchId` (required)
- `sourceBranchId`
- `includeProcedureGeneratedWorkflowElements`

### `Rebase`

`POST /v1/convai/agents/{agent_id}/branches/{branch_id}/rebase`

Rebase A Branch Onto Main

**Parameters:**

- `agentId` (required)
- `branchId` (required)
- `sourceBranchId`

## Versions

### `Get`

`GET /v1/convai/agents/{agent_id}/versions/{version_id}`

Get Agent Version Metadata

**Parameters:**

- `agentId` (required)
- `versionId` (required)

## Deployments

### `List`

`GET /v1/convai/agents/{agent_id}/deployments`

List Agent Deployments

**Parameters:**

- `agentId` (required)
- `page`
- `pageSize`

### `Create`

`POST /v1/convai/agents/{agent_id}/deployments`

Create Or Update Deployments

**Parameters:**

- `agentId` (required)

## Drafts

### `Create`

`POST /v1/convai/agents/{agent_id}/drafts`

Create Agent Draft

**Parameters:**

- `agentId` (required)
- `branchId` (required)

### `Delete`

`DELETE /v1/convai/agents/{agent_id}/drafts`

Delete Agent Draft

**Parameters:**

- `agentId` (required)
- `branchId` (required)

### `Get`

`GET /v1/convai/agents/{agent_id}/branches/{branch_id}/procedures/{procedure_id}/draft`

Get Procedure Draft

**Parameters:**

- `agentId` (required)
- `branchId` (required)
- `procedureId` (required)

### `Update`

`PATCH /v1/convai/agents/{agent_id}/branches/{branch_id}/procedures/{procedure_id}/draft`

Update Procedure Draft

**Parameters:**

- `agentId` (required)
- `branchId` (required)
- `procedureId` (required)

### `Delete1`

`DELETE /v1/convai/agents/{agent_id}/branches/{branch_id}/procedures/{procedure_id}/draft`

Delete Procedure Draft

**Parameters:**

- `agentId` (required)
- `branchId` (required)
- `procedureId` (required)

## MergeProposals

### `List`

`GET /v1/convai/agents/{agent_id}/merge-proposals`

List Proposals

**Parameters:**

- `agentId` (required)
- `status`
- `sourceBranchId`
- `targetBranchId`
- `search`
- `pageSize`
- `cursor`

### `Create`

`POST /v1/convai/agents/{agent_id}/merge-proposals`

Create A Merge Proposal

**Parameters:**

- `agentId` (required)

### `Get`

`GET /v1/convai/agents/{agent_id}/merge-proposals/{merge_proposal_id}`

Get A Merge Proposal

**Parameters:**

- `agentId` (required)
- `mergeProposalId` (required)

### `Update`

`PATCH /v1/convai/agents/{agent_id}/merge-proposals/{merge_proposal_id}`

Update A Merge Proposal

**Parameters:**

- `agentId` (required)
- `mergeProposalId` (required)

### `SubmitReview`

`POST /v1/convai/agents/{agent_id}/merge-proposals/{merge_proposal_id}/reviews`

Review A Merge Proposal

**Parameters:**

- `agentId` (required)
- `mergeProposalId` (required)

### `AddComment`

`POST /v1/convai/agents/{agent_id}/merge-proposals/{merge_proposal_id}/comments`

Comment On A Merge Proposal

**Parameters:**

- `agentId` (required)
- `mergeProposalId` (required)

### `Merge`

`POST /v1/convai/agents/{agent_id}/merge-proposals/{merge_proposal_id}/merge`

Merge A Merge Proposal

**Parameters:**

- `agentId` (required)
- `mergeProposalId` (required)

## Config

### `Get`

`GET /v1/convai/outbound-campaigns/config`

Get Outbound Campaign Configuration

### `Update`

`PUT /v1/convai/outbound-campaigns/config`

Replace Outbound Campaign Configuration

## OutboundCampaigns

### `List`

`GET /v1/convai/outbound-campaigns/campaigns`

List Outbound Campaigns

**Parameters:**

- `pageSize`
- `status`
- `agentId`
- `channel`
- `mode`
- `search`
- `cursor`

### `PauseAll`

`POST /v1/convai/outbound-campaigns/campaigns/pause-all`

Pause All Outbound Campaigns

### `StopAll`

`POST /v1/convai/outbound-campaigns/campaigns/stop-all`

Stop All Outbound Campaigns

### `Create`

`POST /v1/convai/outbound-campaigns/campaigns/submit`

Create An Outbound Campaign

### `Get`

`GET /v1/convai/outbound-campaigns/campaigns/{campaign_id}`

Get An Outbound Campaign

**Parameters:**

- `campaignId` (required)

### `GetProgress`

`GET /v1/convai/outbound-campaigns/campaigns/{campaign_id}/progress`

Get Outbound Campaign Progress

**Parameters:**

- `campaignId` (required)

### `GetFollowUpRequests`

`GET /v1/convai/outbound-campaigns/campaigns/{campaign_id}/follow-up-requests`

Get Outbound Campaign Follow-Up Requests

**Parameters:**

- `campaignId` (required)

### `Pause`

`POST /v1/convai/outbound-campaigns/campaigns/{campaign_id}/pause`

Pause An Outbound Campaign

**Parameters:**

- `campaignId` (required)

### `Resume`

`POST /v1/convai/outbound-campaigns/campaigns/{campaign_id}/resume`

Resume An Outbound Campaign

**Parameters:**

- `campaignId` (required)

### `Stop`

`POST /v1/convai/outbound-campaigns/campaigns/{campaign_id}/stop`

Stop An Outbound Campaign

**Parameters:**

- `campaignId` (required)

## Leads

### `List`

`GET /v1/convai/outbound-campaigns/campaigns/{campaign_id}/leads`

List Outbound Campaign Leads

**Parameters:**

- `campaignId` (required)
- `pageSize`
- `search`
- `status`
- `disposition`
- `cursor`

### `Get`

`GET /v1/convai/outbound-campaigns/campaigns/{campaign_id}/leads/{lead_id}`

Get An Outbound Campaign Lead

**Parameters:**

- `campaignId` (required)
- `leadId` (required)

### `Delete`

`DELETE /v1/convai/outbound-campaigns/campaigns/{campaign_id}/leads/{lead_id}`

Remove An Outbound Campaign Lead

**Parameters:**

- `campaignId` (required)
- `leadId` (required)

## Calls

### `List`

`GET /v1/convai/outbound-campaigns/campaigns/{campaign_id}/calls`

List Outbound Campaign Calls

**Parameters:**

- `campaignId` (required)
- `pageSize`
- `leadId`
- `status`
- `callResult`
- `stage`
- `cursor`

### `Export`

`GET /v1/convai/outbound-campaigns/campaigns/{campaign_id}/export`

Export Outbound Campaign Calls

**Parameters:**

- `campaignId` (required)
- `limit`

## Attention

### `List`

`GET /v1/convai/outbound-campaigns/attention`

List Outbound Campaign Attention Items

**Parameters:**

- `campaignId`
- `includeResolved`
- `pageSize`
- `cursor`

### `Retry`

`POST /v1/convai/outbound-campaigns/attention/{attention_id}/retry`

Retry An Outbound Campaign Attention Item

**Parameters:**

- `attentionId` (required)

### `Resolve`

`POST /v1/convai/outbound-campaigns/attention/{attention_id}/resolve`

Resolve An Outbound Campaign Attention Item

**Parameters:**

- `attentionId` (required)

## SpeechEngine

### `List`

`GET /v1/speech-engine`

List Speech Engines

**Parameters:**

- `pageSize`
- `search`
- `sortDirection`
- `sortBy`
- `cursor`

### `Create`

`POST /v1/speech-engine`

Create Speech Engine

### `Get`

`GET /v1/speech-engine/{speech_engine_id}`

Get Speech Engine

**Parameters:**

- `speechEngineId` (required)

### `Update`

`PATCH /v1/speech-engine/{speech_engine_id}`

Update Speech Engine

**Parameters:**

- `speechEngineId` (required)

### `Delete`

`DELETE /v1/speech-engine/{speech_engine_id}`

Delete Speech Engine

**Parameters:**

- `speechEngineId` (required)

### `Duplicate`

`POST /v1/speech-engine/{speech_engine_id}/duplicate`

Duplicate Speech Engine

**Parameters:**

- `speechEngineId` (required)

## Procedures

### `List`

`GET /v1/convai/agents/{agent_id}/branches/{branch_id}/procedures`

List Procedures

**Parameters:**

- `agentId` (required)
- `branchId` (required)
- `agentVersionId`

### `Create`

`POST /v1/convai/agents/{agent_id}/branches/{branch_id}/procedures`

Create Procedure

**Parameters:**

- `agentId` (required)
- `branchId` (required)

### `Compile`

`POST /v1/convai/agents/{agent_id}/branches/{branch_id}/procedures/compile`

Legacy: Compile Procedures

**Parameters:**

- `agentId` (required)
- `branchId` (required)

### `Get`

`GET /v1/convai/agents/{agent_id}/branches/{branch_id}/procedures/{procedure_id}`

Get Procedure

**Parameters:**

- `agentId` (required)
- `branchId` (required)
- `procedureId` (required)
- `versionId`
- `agentVersionId`

### `Remove`

`DELETE /v1/convai/agents/{agent_id}/branches/{branch_id}/procedures/{procedure_id}`

Remove Procedure

**Parameters:**

- `agentId` (required)
- `branchId` (required)
- `procedureId` (required)

### `List1`

`GET /v1/reception/customer-facing-agents/{agent_id}/procedures`

List Procedures

**Parameters:**

- `agentId` (required)

### `Create1`

`POST /v1/reception/customer-facing-agents/{agent_id}/procedures`

Add Procedure

**Parameters:**

- `agentId` (required)

### `Update`

`PUT /v1/reception/customer-facing-agents/{agent_id}/procedures/{procedure_id}`

Update Procedure

**Parameters:**

- `agentId` (required)
- `procedureId` (required)

### `Delete`

`DELETE /v1/reception/customer-facing-agents/{agent_id}/procedures/{procedure_id}`

Delete Procedure

**Parameters:**

- `agentId` (required)
- `procedureId` (required)

## Analysis

### `Run`

`POST /v1/convai/conversations/{conversation_id}/analysis/run`

Run Conversation Analysis

**Parameters:**

- `conversationId` (required)

### `RunEvaluation`

`POST /v1/convai/conversations/{conversation_id}/analysis/evaluations/run`

Run Conversation Evaluation

**Parameters:**

- `conversationId` (required)

## EnvironmentVariables

### `List`

`GET /v1/convai/environment-variables`

List Environment Variables

**Parameters:**

- `cursor`
- `pageSize`
- `label`
- `environment`
- `type_`

### `Create`

`POST /v1/convai/environment-variables`

Create Environment Variable

### `Get`

`GET /v1/convai/environment-variables/{env_var_id}`

Get Environment Variable

**Parameters:**

- `envVarId` (required)

### `Update`

`PATCH /v1/convai/environment-variables/{env_var_id}`

Update Environment Variable

**Parameters:**

- `envVarId` (required)

### `Delete`

`DELETE /v1/convai/environment-variables/{env_var_id}`

Delete Environment Variable

**Parameters:**

- `envVarId` (required)

## Environments

### `List`

`GET /v1/convai/environments`

List Environments

### `Create`

`POST /v1/convai/environments`

Create Environment

### `Delete`

`DELETE /v1/convai/environments/{environment}`

Delete Environment

**Parameters:**

- `environment` (required)

## Realtime

### `ListConversationGroups`

`GET /v1/convai/realtime/conversation-groups`

List Realtime Conversation Groups

**Parameters:**

- `pageSize`
- `search`
- `cursor`

### `CreateConversationGroup`

`POST /v1/convai/realtime/conversation-groups`

Create Realtime Conversation Group

### `GetConversationGroup`

`GET /v1/convai/realtime/conversation-groups/{conversation_group_id}`

Get Realtime Conversation Group

**Parameters:**

- `conversationGroupId` (required)

### `UpdateConversationGroup`

`PATCH /v1/convai/realtime/conversation-groups/{conversation_group_id}`

Update Realtime Conversation Group

**Parameters:**

- `conversationGroupId` (required)

### `DeleteConversationGroup`

`DELETE /v1/convai/realtime/conversation-groups/{conversation_group_id}`

Delete Realtime Conversation Group

**Parameters:**

- `conversationGroupId` (required)

### `SetConversationGroup`

`PUT /v1/convai/realtime/conversations/{conversation_id}/conversation-group`

Set Realtime Conversation Group

**Parameters:**

- `conversationId` (required)

### `ClearConversationGroup`

`DELETE /v1/convai/realtime/conversations/{conversation_id}/conversation-group`

Clear Realtime Conversation Group

**Parameters:**

- `conversationId` (required)

### `CreateClientSecret`

`POST /v1/convai/realtime/client_secrets`

Create Realtime Client Secret

## CompositionPlan

### `Create`

`POST /v1/music/plan`

Generate Composition Plan

**Parameters:**

- `enableLogging`

## Finetunes

### `List`

`GET /v1/music/finetunes`

Get Music Finetunes

**Parameters:**

- `cursor`
- `pageSize`
- `visibility`
- `createdBy`
- `sort`
- `sortDirection`

### `Create`

`POST /v1/music/finetunes`

Create Music Finetune

### `Get`

`GET /v1/music/finetunes/{finetune_id}`

Get Music Finetune

**Parameters:**

- `finetuneId` (required)

### `Update`

`PATCH /v1/music/finetunes/{finetune_id}`

Update Music Finetune

**Parameters:**

- `finetuneId` (required)

### `Delete`

`DELETE /v1/music/finetunes/{finetune_id}`

Delete Music Finetune

**Parameters:**

- `finetuneId` (required)

## Orders

### `List`

`GET /v1/productions/orders`

List Orders

**Parameters:**

- `pageSize`
- `offset`
- `status`
- `startDate`
- `endDate`

### `Create`

`POST /v1/productions/orders`

Create Order

### `Get`

`GET /v1/productions/orders/{order_id}`

Get Order

**Parameters:**

- `orderId` (required)

### `Update`

`PATCH /v1/productions/orders/{order_id}`

Update Order

**Parameters:**

- `orderId` (required)

### `Submit`

`POST /v1/productions/orders/{order_id}/submit`

Submit Order

**Parameters:**

- `orderId` (required)

## Media

### `Register`

`POST /v1/productions/orders/{order_id}/media`

Register Media

**Parameters:**

- `orderId` (required)

### `Get`

`GET /v1/productions/orders/{order_id}/media/{media_id}`

Get Media Info

**Parameters:**

- `orderId` (required)
- `mediaId` (required)

## Items

### `Upsert`

`POST /v1/productions/orders/{order_id}/items`

Upsert Order Item

**Parameters:**

- `orderId` (required)

### `Remove`

`DELETE /v1/productions/orders/{order_id}/items/{item_id}`

Remove Order Item

**Parameters:**

- `orderId` (required)
- `itemId` (required)

## Deliverables

### `List`

`GET /v1/productions/orders/{order_id}/deliverables`

Get Order Deliverables

**Parameters:**

- `orderId` (required)

## Languages

### `List`

`GET /v1/productions/orders/languages/{order_item_kind}`

Get Available Languages

**Parameters:**

- `orderItemKind` (required)

## Pvc

### `Create`

`POST /v1/voices/pvc`

Create PVC voice

### `Update`

`POST /v1/voices/pvc/{voice_id}`

Edit Pvc Voice

**Parameters:**

- `voiceId` (required)

### `Train`

`POST /v1/voices/pvc/{voice_id}/train`

Run Pvc Training

**Parameters:**

- `voiceId` (required)

## Waveform

### `Get`

`GET /v1/voices/pvc/{voice_id}/samples/{sample_id}/waveform`

Retrieve Voice Sample Visual Waveform

**Parameters:**

- `voiceId` (required)
- `sampleId` (required)

## Speakers

### `Get`

`GET /v1/voices/pvc/{voice_id}/samples/{sample_id}/speakers`

Retrieve Speaker Separation Status

**Parameters:**

- `voiceId` (required)
- `sampleId` (required)

### `Separate`

`POST /v1/voices/pvc/{voice_id}/samples/{sample_id}/separate-speakers`

Start Speaker Separation

**Parameters:**

- `voiceId` (required)
- `sampleId` (required)

## Captcha

### `Get`

`GET /v1/voices/pvc/{voice_id}/captcha`

Get Pvc Voice Captcha

**Parameters:**

- `voiceId` (required)

### `Verify`

`POST /v1/voices/pvc/{voice_id}/captcha`

Verify Pvc Voice Captcha

**Parameters:**

- `voiceId` (required)

## Verification

### `Request`

`POST /v1/voices/pvc/{voice_id}/verification`

Request Manual Verification

**Parameters:**

- `voiceId` (required)

## Video

### `List`

`GET /v1/flows/video`

List Video Generations

**Parameters:**

- `cursor`
- `pageSize`
- `status`
- `modelId`

### `Create`

`POST /v1/flows/video`

Create Video Generation

### `Get`

`GET /v1/flows/video/{generation_id}`

Get Video Generation

**Parameters:**

- `generationId` (required)

## Image

### `List`

`GET /v1/flows/image`

List Image Generations

**Parameters:**

- `cursor`
- `pageSize`
- `status`
- `modelId`

### `Create`

`POST /v1/flows/image`

Create Image Generation

### `Get`

`GET /v1/flows/image/{generation_id}`

Get Image Generation

**Parameters:**

- `generationId` (required)

## Runs

### `List`

`GET /v1/flows/templates/{template_id}/runs`

List Template Runs

**Parameters:**

- `templateId` (required)
- `cursor`
- `pageSize`
- `versionId`

### `Create`

`POST /v1/flows/templates/{template_id}/runs`

Create Template Run

**Parameters:**

- `templateId` (required)

### `Get`

`GET /v1/flows/templates/{template_id}/runs/{run_id}`

Get Template Run

**Parameters:**

- `templateId` (required)
- `runId` (required)

## Templates

### `List`

`GET /v1/flows/templates`

List Templates

**Parameters:**

- `cursor`
- `pageSize`
- `versionsPerTemplate`
- `search`

### `Get`

`GET /v1/flows/templates/{template_id}`

Get Template

**Parameters:**

- `templateId` (required)
- `versionsPerTemplate`

## Assets

### `List`

`GET /v1/assets`

List Assets

**Parameters:**

- `pageSize`
- `cursor`
- `search`

### `Create`

`POST /v1/assets`

Upload Asset

### `Get`

`GET /v1/assets/{asset_id}`

Get Asset

**Parameters:**

- `assetId` (required)

### `Delete`

`DELETE /v1/assets/{asset_id}`

Delete Asset

**Parameters:**

- `assetId` (required)

### `List1`

`GET /v1/reception/assets`

Get All Assets

### `Create1`

`POST /v1/reception/assets`

Create Asset

### `Get1`

`GET /v1/reception/assets/{asset_id}`

Get Asset

**Parameters:**

- `assetId` (required)

### `Update`

`PUT /v1/reception/assets/{asset_id}`

Update Asset

**Parameters:**

- `assetId` (required)

### `Delete1`

`DELETE /v1/reception/assets/{asset_id}`

Delete Asset

**Parameters:**

- `assetId` (required)

## Usage

### `GetUsageByProductOverTime`

`POST /v1/workspace/analytics/query/usage-by-product-over-time`

Get Workspace Usage

### `GetUsage`

`POST /v1/workspace/analytics/usage`

Get Workspace Cost Usage

## Requests

### `Get`

`POST /v1/workspace/analytics/requests`

List Api Requests

## Services

### `List`

`GET /v1/reception/services`

Get All Services

**Parameters:**

- `staffId`
- `assetId`
- `isActive`
- `kind`

### `Create`

`POST /v1/reception/services`

Create Service

### `Get`

`GET /v1/reception/services/{service_id}`

Get Service

**Parameters:**

- `serviceId` (required)

### `Update`

`PUT /v1/reception/services/{service_id}`

Update Service

**Parameters:**

- `serviceId` (required)

### `Delete`

`DELETE /v1/reception/services/{service_id}`

Delete Service

**Parameters:**

- `serviceId` (required)

### `ConfirmUpdate`

`POST /v1/reception/services/{service_id}/confirm-update`

Confirm Service Update

**Parameters:**

- `serviceId` (required)
- `reprice`

## Products

### `List`

`GET /v1/reception/products`

Get All Products

**Parameters:**

- `group`
- `isActive`

### `Create`

`POST /v1/reception/products`

Create Product

### `Get`

`GET /v1/reception/products/{product_id}`

Get Product

**Parameters:**

- `productId` (required)

### `Update`

`PUT /v1/reception/products/{product_id}`

Update Product

**Parameters:**

- `productId` (required)

### `Delete`

`DELETE /v1/reception/products/{product_id}`

Delete Product

**Parameters:**

- `productId` (required)

## Stock

### `Set`

`PATCH /v1/reception/products/{product_id}/stock`

Update Product Stock

**Parameters:**

- `productId` (required)

### `Adjust`

`POST /v1/reception/products/{product_id}/stock/adjust`

Adjust Product Stock

**Parameters:**

- `productId` (required)

## Staff

### `List`

`GET /v1/reception/staff`

Get All Staff

### `Create`

`POST /v1/reception/staff`

Create Staff Member

### `Get`

`GET /v1/reception/staff/{staff_id}`

Get Staff Member

**Parameters:**

- `staffId` (required)

### `Update`

`PUT /v1/reception/staff/{staff_id}`

Update Staff Member

**Parameters:**

- `staffId` (required)

### `Delete`

`DELETE /v1/reception/staff/{staff_id}`

Delete Staff Member

**Parameters:**

- `staffId` (required)

## Receptionists

### `List`

`GET /v1/reception/customer-facing-agents`

List Receptionists

### `Update`

`PATCH /v1/reception/customer-facing-agents/{agent_id}`

Update Receptionist

**Parameters:**

- `agentId` (required)

### `Activate`

`POST /v1/reception/customer-facing-agents/{agent_id}/activate`

Activate Receptionist

**Parameters:**

- `agentId` (required)

## TransferRules

### `Create`

`POST /v1/reception/customer-facing-agents/{agent_id}/transfer-rules`

Add Transfer Rule

**Parameters:**

- `agentId` (required)

### `Update`

`PUT /v1/reception/customer-facing-agents/{agent_id}/transfer-rules/{rule_id}`

Update Transfer Rule

**Parameters:**

- `agentId` (required)
- `ruleId` (required)

### `Delete`

`DELETE /v1/reception/customer-facing-agents/{agent_id}/transfer-rules/{rule_id}`

Delete Transfer Rule

**Parameters:**

- `agentId` (required)
- `ruleId` (required)

## Holidays

### `List`

`GET /v1/reception/holidays`

Get All Holidays

**Parameters:**

- `startUnix`
- `endUnix`
- `timeOffType`
- `resourceId`
- `locationId`

### `Create`

`POST /v1/reception/holidays`

Create Holiday

### `Get`

`GET /v1/reception/holidays/{holiday_id}`

Get Holiday

**Parameters:**

- `holidayId` (required)

### `Update`

`PUT /v1/reception/holidays/{holiday_id}`

Update Holiday

**Parameters:**

- `holidayId` (required)

### `Delete`

`DELETE /v1/reception/holidays/{holiday_id}`

Delete Holiday

**Parameters:**

- `holidayId` (required)

## Locations

### `List`

`GET /v1/reception/locations`

List Locations

### `Create`

`POST /v1/reception/locations`

Create Location

### `Get`

`GET /v1/reception/locations/{location_id}`

Get Location

**Parameters:**

- `locationId` (required)

### `Update`

`PUT /v1/reception/locations/{location_id}`

Update Location

**Parameters:**

- `locationId` (required)

### `Delete`

`DELETE /v1/reception/locations/{location_id}`

Delete Location

**Parameters:**

- `locationId` (required)

## BusinessProfile

### `Get`

`GET /v1/reception/agent-info`

Get Business Profile

### `Update`

`PUT /v1/reception/agent-info`

Update Business Profile

## KnowledgeSources

### `List`

`GET /v1/reception/knowledge-sources`

List Knowledge Sources

### `Upload`

`POST /v1/reception/knowledge-sources`

Upload Knowledge Source

### `Get`

`GET /v1/reception/knowledge-sources/{source_id}`

Get Knowledge Source

**Parameters:**

- `sourceId` (required)

### `Delete`

`DELETE /v1/reception/knowledge-sources/{source_id}`

Delete Knowledge Source

**Parameters:**

- `sourceId` (required)

### `AddUrl`

`POST /v1/reception/knowledge-sources/urls`

Add Url Source

## Workspaces

### `List`

`GET /v1/reception/reseller/workspaces`

List Reseller Workspaces

**Parameters:**

- `cursor`
- `pageSize`

### `Create`

`POST /v1/reception/reseller/workspaces`

Create Reseller Workspace

### `CompleteOnboarding`

`POST /v1/reception/reseller/workspaces/{workspace_id}/complete-onboarding`

Complete Reseller Workspace Onboarding

**Parameters:**

- `workspaceId` (required)

## Subscriptions

### `Assign`

`POST /v1/reception/reseller/workspaces/{workspace_id}/subscription`

Assign Reseller Tier

**Parameters:**

- `workspaceId` (required)

### `Change`

`PATCH /v1/reception/reseller/workspaces/{workspace_id}/subscription`

Change Reseller Tier

**Parameters:**

- `workspaceId` (required)

### `Unassign`

`DELETE /v1/reception/reseller/workspaces/{workspace_id}/subscription`

Unassign Reseller Tier

**Parameters:**

- `workspaceId` (required)

### `Restore`

`POST /v1/reception/reseller/workspaces/{workspace_id}/subscription/restore`

Restore Reseller Subscription

**Parameters:**

- `workspaceId` (required)

### `CancelPendingDowngrade`

`POST /v1/reception/reseller/workspaces/{workspace_id}/subscription/cancel-downgrade`

Cancel Reseller Pending Downgrade

**Parameters:**

- `workspaceId` (required)

## Tiers

### `List`

`GET /v1/reception/reseller/tier-ladder`

List Reseller Tiers

## Billing

### `ListInvoices`

`GET /v1/reception/reseller/billing/invoices`

List Reseller Invoices

**Parameters:**

- `cursor`
- `pageSize`

### `GetInvoice`

`GET /v1/reception/reseller/billing/invoices/{invoice_id}`

Get Reseller Invoice

**Parameters:**

- `invoiceId` (required)

### `ListUpcomingPayments`

`GET /v1/reception/reseller/billing/upcoming`

List Reseller Upcoming Payments
