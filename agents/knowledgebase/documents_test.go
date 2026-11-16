package knowledgebase

import (
	"testing"
)

func TestBodyCreateUrlDocumentV1ConvaiKnowledgeBaseUrlPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "url": "value",
  "name": "value",
  "parent_folder_id": "value",
  "enable_auto_sync": true,
  "auto_remove": true,
  "minimum_frequency_days": 123
}`
	assertRoundTrip(t, wire, new(BodyCreateUrlDocumentV1ConvaiKnowledgeBaseUrlPost))
}

func TestBodyCreateTextDocumentV1ConvaiKnowledgeBaseTextPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text": "value",
  "name": "value",
  "parent_folder_id": "value"
}`
	assertRoundTrip(t, wire, new(BodyCreateTextDocumentV1ConvaiKnowledgeBaseTextPost))
}

func TestBodyCreateFolderV1ConvaiKnowledgeBaseFolderPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "parent_folder_id": "value",
  "enable_auto_sync": true,
  "auto_remove": true,
  "minimum_frequency_days": 123
}`
	assertRoundTrip(t, wire, new(BodyCreateFolderV1ConvaiKnowledgeBaseFolderPost))
}

func TestGetDocumentsResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value",
  "metadata": {
    "created_at_unix_secs": 123,
    "last_updated_at_unix_secs": 123,
    "size_bytes": 123
  },
  "supported_usages": [
    "prompt"
  ],
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "folder_parent_id": "value",
  "chunking_override": {
    "max_chunk_size": 123,
    "overlap_percentage": 123,
    "custom_separators": [
      "value"
    ]
  },
  "folder_path": [
    {
      "id": "value",
      "name": "value"
    }
  ],
  "type": "url",
  "url": "value",
  "extracted_inner_html": "value",
  "content_format": "html",
  "auto_sync_info": {
    "minimum_frequency_days": 123,
    "auto_remove": true,
    "consec_failures": 123,
    "next_refresh_by": 123
  }
}`
	assertRoundTrip(t, wire, new(GetDocumentsResponse))
}

func TestUpdateDocumentsResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value",
  "metadata": {
    "created_at_unix_secs": 123,
    "last_updated_at_unix_secs": 123,
    "size_bytes": 123
  },
  "supported_usages": [
    "prompt"
  ],
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "folder_parent_id": "value",
  "chunking_override": {
    "max_chunk_size": 123,
    "overlap_percentage": 123,
    "custom_separators": [
      "value"
    ]
  },
  "folder_path": [
    {
      "id": "value",
      "name": "value"
    }
  ],
  "type": "url",
  "url": "value",
  "extracted_inner_html": "value",
  "content_format": "html",
  "auto_sync_info": {
    "minimum_frequency_days": 123,
    "auto_remove": true,
    "consec_failures": 123,
    "next_refresh_by": 123
  }
}`
	assertRoundTrip(t, wire, new(UpdateDocumentsResponse))
}

func TestBodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatchRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "content": "value",
  "enable_auto_sync": true,
  "auto_remove": true,
  "minimum_frequency_days": 123,
  "tighten_only": true
}`
	assertRoundTrip(t, wire, new(BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch))
}

func TestUpdateDocumentChunkingOverrideRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "chunking_override": {
    "max_chunk_size": 200,
    "overlap_percentage": 25,
    "custom_separators": [
      "value"
    ]
  }
}`
	assertRoundTrip(t, wire, new(UpdateDocumentChunkingOverrideRequestModel))
}

func TestBodyGetDependentAgentsForMultipleDocumentsV1ConvaiKnowledgeBaseDependentAgentsPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "document_ids": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(BodyGetDependentAgentsForMultipleDocumentsV1ConvaiKnowledgeBaseDependentAgentsPost))
}

func TestBodyMoveEntityToFolderV1ConvaiKnowledgeBaseDocumentIdMovePostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "move_to": "value"
}`
	assertRoundTrip(t, wire, new(BodyMoveEntityToFolderV1ConvaiKnowledgeBaseDocumentIdMovePost))
}

func TestBodyBulkMoveEntitiesToFolderV1ConvaiKnowledgeBaseBulkMovePostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "document_ids": [
    "value"
  ],
  "move_to": "value"
}`
	assertRoundTrip(t, wire, new(BodyBulkMoveEntitiesToFolderV1ConvaiKnowledgeBaseBulkMovePost))
}

func TestBulkDeleteDocumentsResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "success",
  "data": {
    "id": "value"
  }
}`
	assertRoundTrip(t, wire, new(BulkDeleteDocumentsResponse))
}

func TestBodyBulkDeleteKnowledgeBaseDocumentsV1ConvaiKnowledgeBaseBulkDeletePostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "document_ids": [
    "value"
  ],
  "force": true
}`
	assertRoundTrip(t, wire, new(BodyBulkDeleteKnowledgeBaseDocumentsV1ConvaiKnowledgeBaseBulkDeletePost))
}
