package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAuditActivityNameVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant AuditActivityName
		wire    string
	}
	cases := []enumCase{
		{variant: AuditActivityNameAccountSecurityChange, wire: "\"Account Security Change\""},
		{variant: AuditActivityNameAgentsUserMemoryEntityDeleted, wire: "\"Agents User Memory Entity Deleted\""},
		{variant: AuditActivityNameAgentsUserMemoryErased, wire: "\"Agents User Memory Erased\""},
		{variant: AuditActivityNameAgentsUserMemoryExported, wire: "\"Agents User Memory Exported\""},
		{variant: AuditActivityNameAgentsUserMemoryFactDeleted, wire: "\"Agents User Memory Fact Deleted\""},
		{variant: AuditActivityNameAgentsUserMemoryViewed, wire: "\"Agents User Memory Viewed\""},
		{variant: AuditActivityNameAgentsUserPropertiesReordered, wire: "\"Agents User Properties Reordered\""},
		{variant: AuditActivityNameAgentsUserPropertyArchived, wire: "\"Agents User Property Archived\""},
		{variant: AuditActivityNameAgentsUserPropertyCreated, wire: "\"Agents User Property Created\""},
		{variant: AuditActivityNameAgentsUserPropertyUnarchived, wire: "\"Agents User Property Unarchived\""},
		{variant: AuditActivityNameAgentsUserPropertyUpdated, wire: "\"Agents User Property Updated\""},
		{variant: AuditActivityNameApiKeyCreation, wire: "\"API Key Creation\""},
		{variant: AuditActivityNameApiKeyDeletion, wire: "\"API Key Deletion\""},
		{variant: AuditActivityNameApiKeyManagement, wire: "\"API Key Management\""},
		{variant: AuditActivityNameApiKeyUpdate, wire: "\"API Key Update\""},
		{variant: AuditActivityNameArchitectMemberSpendLimitsUpdate, wire: "\"Architect Member Spend Limits Update\""},
		{variant: AuditActivityNameAuthConnectionCreation, wire: "\"Auth Connection Creation\""},
		{variant: AuditActivityNameAuthConnectionDeletion, wire: "\"Auth Connection Deletion\""},
		{variant: AuditActivityNameAuthConnectionManagement, wire: "\"Auth Connection Management\""},
		{variant: AuditActivityNameAuthConnectionUpdate, wire: "\"Auth Connection Update\""},
		{variant: AuditActivityNameBillingManagement, wire: "\"Billing Management\""},
		{variant: AuditActivityNameChildWorkspaceResourceOverride, wire: "\"Child Workspace Resource Override\""},
		{variant: AuditActivityNameContentGenerationTermsUpdate, wire: "\"Content Generation Terms Update\""},
		{variant: AuditActivityNameCreateWorkspaceCustomRole, wire: "\"Create workspace custom role\""},
		{variant: AuditActivityNameCustomAuthTokenCreated, wire: "\"Custom Auth Token Created\""},
		{variant: AuditActivityNameDataExportRequest, wire: "\"Data Export Request\""},
		{variant: AuditActivityNameDeleteScribeMeeting, wire: "\"Delete Scribe meeting\""},
		{variant: AuditActivityNameDeleteWorkspaceCustomRole, wire: "\"Delete workspace custom role\""},
		{variant: AuditActivityNameDomainVerificationAddition, wire: "\"Domain Verification Addition\""},
		{variant: AuditActivityNameDomainVerificationManagement, wire: "\"Domain Verification Management\""},
		{variant: AuditActivityNameDomainVerificationRemoval, wire: "\"Domain Verification Removal\""},
		{variant: AuditActivityNameDomainVerificationStatusUpdate, wire: "\"Domain Verification Status Update\""},
		{variant: AuditActivityNameEncryptionKeyDeleted, wire: "\"Encryption Key Deleted\""},
		{variant: AuditActivityNameEncryptionKeyPaused, wire: "\"Encryption Key Paused\""},
		{variant: AuditActivityNameEncryptionKeyRegistered, wire: "\"Encryption Key Registered\""},
		{variant: AuditActivityNameEncryptionKeyUnpaused, wire: "\"Encryption Key Unpaused\""},
		{variant: AuditActivityNameEnterpriseCopyingToggle, wire: "\"Enterprise Copying Toggle\""},
		{variant: AuditActivityNameGiftCreditCodeApplied, wire: "\"Gift Credit Code Applied\""},
		{variant: AuditActivityNameGroupCreation, wire: "\"Group Creation\""},
		{variant: AuditActivityNameGroupDeletion, wire: "\"Group Deletion\""},
		{variant: AuditActivityNameGroupManagement, wire: "\"Group Management\""},
		{variant: AuditActivityNameGroupMemberAddition, wire: "\"Group Member Addition\""},
		{variant: AuditActivityNameGroupMemberRemoval, wire: "\"Group Member Removal\""},
		{variant: AuditActivityNameGroupToBillingGroupConversion, wire: "\"Group to Billing Group Conversion\""},
		{variant: AuditActivityNameGroupUpdate, wire: "\"Group Update\""},
		{variant: AuditActivityNameIdentityProviderUpdate, wire: "\"Identity Provider Update\""},
		{variant: AuditActivityNameInfluencerCodeApplied, wire: "\"Influencer Code Applied\""},
		{variant: AuditActivityNameInAppTermAcceptance, wire: "\"In-App Term Acceptance\""},
		{variant: AuditActivityNameLogon, wire: "\"Logon\""},
		{variant: AuditActivityNameMfaFactorEnrollment, wire: "\"MFA Factor Enrollment\""},
		{variant: AuditActivityNameMfaFactorEnrollmentRevertedUnverified, wire: "\"MFA Factor Enrollment Reverted (unverified)\""},
		{variant: AuditActivityNameMfaFactorRemoval, wire: "\"MFA Factor Removal\""},
		{variant: AuditActivityNameMfaRemovalViaBackupCode, wire: "\"MFA Removal via Backup Code\""},
		{variant: AuditActivityNameModelApprovalsUpdate, wire: "\"Model Approvals Update\""},
		{variant: AuditActivityNameNotificationPreferenceCreation, wire: "\"Notification Preference Creation\""},
		{variant: AuditActivityNameNotificationPreferenceDeletion, wire: "\"Notification Preference Deletion\""},
		{variant: AuditActivityNameNotificationPreferenceManagement, wire: "\"Notification Preference Management\""},
		{variant: AuditActivityNameNotificationPreferenceUpdate, wire: "\"Notification Preference Update\""},
		{variant: AuditActivityNamePasswordResetEmailRequest, wire: "\"Password Reset Email Request\""},
		{variant: AuditActivityNamePasswordResetLinkGeneration, wire: "\"Password Reset Link Generation\""},
		{variant: AuditActivityNameReportingWorkspaceSsoInheritanceUpdate, wire: "\"Reporting Workspace SSO Inheritance Update\""},
		{variant: AuditActivityNameResourceAccessGranted, wire: "\"Resource Access Granted\""},
		{variant: AuditActivityNameResourceAccessManagement, wire: "\"Resource Access Management\""},
		{variant: AuditActivityNameResourceAccessRevoked, wire: "\"Resource Access Revoked\""},
		{variant: AuditActivityNameResourceCopiedToWorkspace, wire: "\"Resource Copied to Workspace\""},
		{variant: AuditActivityNameResourceLocatorCreation, wire: "\"Resource Locator Creation\""},
		{variant: AuditActivityNameResourceLocatorDeletion, wire: "\"Resource Locator Deletion\""},
		{variant: AuditActivityNameResourceLocatorManagement, wire: "\"Resource Locator Management\""},
		{variant: AuditActivityNameResourceLocatorUpdate, wire: "\"Resource Locator Update\""},
		{variant: AuditActivityNameResourceOwnershipTransferred, wire: "\"Resource Ownership Transferred\""},
		{variant: AuditActivityNameScimGroupCreation, wire: "\"SCIM Group Creation\""},
		{variant: AuditActivityNameScimGroupDeletion, wire: "\"SCIM Group Deletion\""},
		{variant: AuditActivityNameScimGroupManagement, wire: "\"SCIM Group Management\""},
		{variant: AuditActivityNameScimGroupUpdate, wire: "\"SCIM Group Update\""},
		{variant: AuditActivityNameScimProvisioningSeatTypeUpdated, wire: "\"SCIM Provisioning Seat Type Updated\""},
		{variant: AuditActivityNameScimTokenDeleted, wire: "\"SCIM Token Deleted\""},
		{variant: AuditActivityNameScimTokenGenerated, wire: "\"SCIM Token Generated\""},
		{variant: AuditActivityNameScimTokenManagement, wire: "\"SCIM Token Management\""},
		{variant: AuditActivityNameScimTokenRotated, wire: "\"SCIM Token Rotated\""},
		{variant: AuditActivityNameScimUserCreation, wire: "\"SCIM User Creation\""},
		{variant: AuditActivityNameScimUserDeletion, wire: "\"SCIM User Deletion\""},
		{variant: AuditActivityNameScimUserManagement, wire: "\"SCIM User Management\""},
		{variant: AuditActivityNameScimUserUpdate, wire: "\"SCIM User Update\""},
		{variant: AuditActivityNameSensitiveActionVerification, wire: "\"Sensitive Action Verification\""},
		{variant: AuditActivityNameServiceAccountCreation, wire: "\"Service Account Creation\""},
		{variant: AuditActivityNameServiceAccountDeletion, wire: "\"Service Account Deletion\""},
		{variant: AuditActivityNameServiceAccountManagement, wire: "\"Service Account Management\""},
		{variant: AuditActivityNameServiceAccountUpdate, wire: "\"Service Account Update\""},
		{variant: AuditActivityNameSharedSsoSiblingWorkspacesUpdate, wire: "\"Shared SSO Sibling Workspaces Update\""},
		{variant: AuditActivityNameSignInRequestApproval, wire: "\"Sign-In Request Approval\""},
		{variant: AuditActivityNameSignInRequestDecline, wire: "\"Sign-In Request Decline\""},
		{variant: AuditActivityNameSmbFreeTrialActivation, wire: "\"SMB Free Trial Activation\""},
		{variant: AuditActivityNameSmbResellerSubscriptionLifecycle, wire: "\"SMB Reseller Subscription Lifecycle\""},
		{variant: AuditActivityNameSmbResellerTierAssignment, wire: "\"SMB Reseller Tier Assignment\""},
		{variant: AuditActivityNameSmbResellerTierChange, wire: "\"SMB Reseller Tier Change\""},
		{variant: AuditActivityNameSsoManagement, wire: "\"SSO Management\""},
		{variant: AuditActivityNameSsoProviderConfiguration, wire: "\"SSO Provider Configuration\""},
		{variant: AuditActivityNameSsoProviderRemoval, wire: "\"SSO Provider Removal\""},
		{variant: AuditActivityNameUpdateWorkspaceCustomRole, wire: "\"Update workspace custom role\""},
		{variant: AuditActivityNameUserApiKeyCreation, wire: "\"User API Key Creation\""},
		{variant: AuditActivityNameUserApiKeyDeletion, wire: "\"User API Key Deletion\""},
		{variant: AuditActivityNameUserApiKeyManagement, wire: "\"User API Key Management\""},
		{variant: AuditActivityNameUserApiKeyUpdate, wire: "\"User API Key Update\""},
		{variant: AuditActivityNameUserAutoProvisioningDisabled, wire: "\"User Auto-Provisioning Disabled\""},
		{variant: AuditActivityNameUserAutoProvisioningEnabled, wire: "\"User Auto-Provisioning Enabled\""},
		{variant: AuditActivityNameUserCreation, wire: "\"User Creation\""},
		{variant: AuditActivityNameUserMetadataUpdate, wire: "\"User Metadata Update\""},
		{variant: AuditActivityNameUserSelfDeletion, wire: "\"User Self-Deletion\""},
		{variant: AuditActivityNameUserSharingPreferencesUpdate, wire: "\"User Sharing Preferences Update\""},
		{variant: AuditActivityNameUserSharingPreferencesUpdateAdmin, wire: "\"User Sharing Preferences Update (Admin)\""},
		{variant: AuditActivityNameWebhookCreation, wire: "\"Webhook Creation\""},
		{variant: AuditActivityNameWebhookDeletion, wire: "\"Webhook Deletion\""},
		{variant: AuditActivityNameWebhookManagement, wire: "\"Webhook Management\""},
		{variant: AuditActivityNameWebhookUpdate, wire: "\"Webhook Update\""},
		{variant: AuditActivityNameWorkspaceCreation, wire: "\"Workspace Creation\""},
		{variant: AuditActivityNameWorkspaceDeletion, wire: "\"Workspace Deletion\""},
		{variant: AuditActivityNameWorkspaceDeletionCancellation, wire: "\"Workspace Deletion Cancellation\""},
		{variant: AuditActivityNameWorkspaceInvitationAccepted, wire: "\"Workspace Invitation Accepted\""},
		{variant: AuditActivityNameWorkspaceInvitationDeclined, wire: "\"Workspace Invitation Declined\""},
		{variant: AuditActivityNameWorkspaceInvitationManagement, wire: "\"Workspace Invitation Management\""},
		{variant: AuditActivityNameWorkspaceInvitationReminderSent, wire: "\"Workspace Invitation Reminder Sent\""},
		{variant: AuditActivityNameWorkspaceInvitationRevoked, wire: "\"Workspace Invitation Revoked\""},
		{variant: AuditActivityNameWorkspaceInvitationSent, wire: "\"Workspace Invitation Sent\""},
		{variant: AuditActivityNameWorkspaceMemberRemoval, wire: "\"Workspace Member Removal\""},
		{variant: AuditActivityNameWorkspaceMemberResourceTransfer, wire: "\"Workspace Member Resource Transfer\""},
		{variant: AuditActivityNameWorkspaceMemberUpdate, wire: "\"Workspace Member Update\""},
		{variant: AuditActivityNameWorkspaceMemberUsageLimitUpdate, wire: "\"Workspace Member Usage Limit Update\""},
		{variant: AuditActivityNameWorkspaceMetadataUpdate, wire: "\"Workspace Metadata Update\""},
		{variant: AuditActivityNameWorkspaceOwnershipTransfer, wire: "\"Workspace Ownership Transfer\""},
		{variant: AuditActivityNameWorkspaceSignIn, wire: "\"Workspace Sign-In\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded AuditActivityName
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
