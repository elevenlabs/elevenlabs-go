package types

import (
	"fmt"
)

// Audit log `activity_name` values that can be stored in a customer workspace's audit log.
type AuditActivityName string

const (
	AuditActivityNameAccountSecurityChange                          AuditActivityName = "Account Security Change"
	AuditActivityNameAgentsUserMemoryEntityDeleted                  AuditActivityName = "Agents User Memory Entity Deleted"
	AuditActivityNameAgentsUserMemoryErased                         AuditActivityName = "Agents User Memory Erased"
	AuditActivityNameAgentsUserMemoryExported                       AuditActivityName = "Agents User Memory Exported"
	AuditActivityNameAgentsUserMemoryFactDeleted                    AuditActivityName = "Agents User Memory Fact Deleted"
	AuditActivityNameAgentsUserMemoryViewed                         AuditActivityName = "Agents User Memory Viewed"
	AuditActivityNameAgentsUserPropertiesReordered                  AuditActivityName = "Agents User Properties Reordered"
	AuditActivityNameAgentsUserPropertyArchived                     AuditActivityName = "Agents User Property Archived"
	AuditActivityNameAgentsUserPropertyCreated                      AuditActivityName = "Agents User Property Created"
	AuditActivityNameAgentsUserPropertyUnarchived                   AuditActivityName = "Agents User Property Unarchived"
	AuditActivityNameAgentsUserPropertyUpdated                      AuditActivityName = "Agents User Property Updated"
	AuditActivityNameApiKeyCreation                                 AuditActivityName = "API Key Creation"
	AuditActivityNameApiKeyDeletion                                 AuditActivityName = "API Key Deletion"
	AuditActivityNameApiKeyManagement                               AuditActivityName = "API Key Management"
	AuditActivityNameApiKeyUpdate                                   AuditActivityName = "API Key Update"
	AuditActivityNameArchitectMemberSpendLimitsUpdate               AuditActivityName = "Architect Member Spend Limits Update"
	AuditActivityNameAuthConnectionCreation                         AuditActivityName = "Auth Connection Creation"
	AuditActivityNameAuthConnectionDeletion                         AuditActivityName = "Auth Connection Deletion"
	AuditActivityNameAuthConnectionManagement                       AuditActivityName = "Auth Connection Management"
	AuditActivityNameAuthConnectionUpdate                           AuditActivityName = "Auth Connection Update"
	AuditActivityNameBillingManagement                              AuditActivityName = "Billing Management"
	AuditActivityNameChildWorkspaceResourceOverride                 AuditActivityName = "Child Workspace Resource Override"
	AuditActivityNameContentGenerationTermsUpdate                   AuditActivityName = "Content Generation Terms Update"
	AuditActivityNameCreateWorkspaceCustomRole                      AuditActivityName = "Create workspace custom role"
	AuditActivityNameCustomAuthTokenCreated                         AuditActivityName = "Custom Auth Token Created"
	AuditActivityNameDataExportRequest                              AuditActivityName = "Data Export Request"
	AuditActivityNameDeleteScribeMeeting                            AuditActivityName = "Delete Scribe meeting"
	AuditActivityNameDeleteWorkspaceCustomRole                      AuditActivityName = "Delete workspace custom role"
	AuditActivityNameDomainVerificationAddition                     AuditActivityName = "Domain Verification Addition"
	AuditActivityNameDomainVerificationManagement                   AuditActivityName = "Domain Verification Management"
	AuditActivityNameDomainVerificationRemoval                      AuditActivityName = "Domain Verification Removal"
	AuditActivityNameDomainVerificationStatusUpdate                 AuditActivityName = "Domain Verification Status Update"
	AuditActivityNameEncryptionKeyDeleted                           AuditActivityName = "Encryption Key Deleted"
	AuditActivityNameEncryptionKeyPaused                            AuditActivityName = "Encryption Key Paused"
	AuditActivityNameEncryptionKeyRegistered                        AuditActivityName = "Encryption Key Registered"
	AuditActivityNameEncryptionKeyUnpaused                          AuditActivityName = "Encryption Key Unpaused"
	AuditActivityNameEnterpriseCopyingToggle                        AuditActivityName = "Enterprise Copying Toggle"
	AuditActivityNameGiftCreditCodeApplied                          AuditActivityName = "Gift Credit Code Applied"
	AuditActivityNameGroupCreation                                  AuditActivityName = "Group Creation"
	AuditActivityNameGroupDeletion                                  AuditActivityName = "Group Deletion"
	AuditActivityNameGroupManagement                                AuditActivityName = "Group Management"
	AuditActivityNameGroupMemberAddition                            AuditActivityName = "Group Member Addition"
	AuditActivityNameGroupMemberRemoval                             AuditActivityName = "Group Member Removal"
	AuditActivityNameGroupToBillingGroupConversion                  AuditActivityName = "Group to Billing Group Conversion"
	AuditActivityNameGroupUpdate                                    AuditActivityName = "Group Update"
	AuditActivityNameIdentityProviderUpdate                         AuditActivityName = "Identity Provider Update"
	AuditActivityNameInfluencerCodeApplied                          AuditActivityName = "Influencer Code Applied"
	AuditActivityNameInAppTermAcceptance                            AuditActivityName = "In-App Term Acceptance"
	AuditActivityNameLogon                                          AuditActivityName = "Logon"
	AuditActivityNameMfaFactorEnrollment                            AuditActivityName = "MFA Factor Enrollment"
	AuditActivityNameMfaFactorEnrollmentRevertedUnverified          AuditActivityName = "MFA Factor Enrollment Reverted (unverified)"
	AuditActivityNameMfaFactorRemoval                               AuditActivityName = "MFA Factor Removal"
	AuditActivityNameMfaRemovalViaBackupCode                        AuditActivityName = "MFA Removal via Backup Code"
	AuditActivityNameModelApprovalsUpdate                           AuditActivityName = "Model Approvals Update"
	AuditActivityNameNotificationPreferenceCreation                 AuditActivityName = "Notification Preference Creation"
	AuditActivityNameNotificationPreferenceDeletion                 AuditActivityName = "Notification Preference Deletion"
	AuditActivityNameNotificationPreferenceManagement               AuditActivityName = "Notification Preference Management"
	AuditActivityNameNotificationPreferenceUpdate                   AuditActivityName = "Notification Preference Update"
	AuditActivityNameOAuthAppAuthorizationGrant                     AuditActivityName = "OAuth App Authorization Grant"
	AuditActivityNameOAuthAppAuthorizationManagement                AuditActivityName = "OAuth App Authorization Management"
	AuditActivityNameOAuthAppAuthorizationRevocation                AuditActivityName = "OAuth App Authorization Revocation"
	AuditActivityNamePasswordResetEmailRequest                      AuditActivityName = "Password Reset Email Request"
	AuditActivityNamePasswordResetLinkGeneration                    AuditActivityName = "Password Reset Link Generation"
	AuditActivityNameReportingWorkspaceSsoInheritanceUpdate         AuditActivityName = "Reporting Workspace SSO Inheritance Update"
	AuditActivityNameResourceAccessGranted                          AuditActivityName = "Resource Access Granted"
	AuditActivityNameResourceAccessManagement                       AuditActivityName = "Resource Access Management"
	AuditActivityNameResourceAccessRevoked                          AuditActivityName = "Resource Access Revoked"
	AuditActivityNameResourceCopiedToWorkspace                      AuditActivityName = "Resource Copied to Workspace"
	AuditActivityNameResourceLocatorCreation                        AuditActivityName = "Resource Locator Creation"
	AuditActivityNameResourceLocatorDeletion                        AuditActivityName = "Resource Locator Deletion"
	AuditActivityNameResourceLocatorManagement                      AuditActivityName = "Resource Locator Management"
	AuditActivityNameResourceLocatorUpdate                          AuditActivityName = "Resource Locator Update"
	AuditActivityNameResourceOwnershipTransferred                   AuditActivityName = "Resource Ownership Transferred"
	AuditActivityNameScimGroupCreation                              AuditActivityName = "SCIM Group Creation"
	AuditActivityNameScimGroupDeletion                              AuditActivityName = "SCIM Group Deletion"
	AuditActivityNameScimGroupManagement                            AuditActivityName = "SCIM Group Management"
	AuditActivityNameScimGroupUpdate                                AuditActivityName = "SCIM Group Update"
	AuditActivityNameScimProvisioningSeatTypeUpdated                AuditActivityName = "SCIM Provisioning Seat Type Updated"
	AuditActivityNameScimTokenDeleted                               AuditActivityName = "SCIM Token Deleted"
	AuditActivityNameScimTokenGenerated                             AuditActivityName = "SCIM Token Generated"
	AuditActivityNameScimTokenManagement                            AuditActivityName = "SCIM Token Management"
	AuditActivityNameScimTokenRotated                               AuditActivityName = "SCIM Token Rotated"
	AuditActivityNameScimUserCreation                               AuditActivityName = "SCIM User Creation"
	AuditActivityNameScimUserDeletion                               AuditActivityName = "SCIM User Deletion"
	AuditActivityNameScimUserManagement                             AuditActivityName = "SCIM User Management"
	AuditActivityNameScimUserUpdate                                 AuditActivityName = "SCIM User Update"
	AuditActivityNameSensitiveActionVerification                    AuditActivityName = "Sensitive Action Verification"
	AuditActivityNameServiceAccountCreation                         AuditActivityName = "Service Account Creation"
	AuditActivityNameServiceAccountDeletion                         AuditActivityName = "Service Account Deletion"
	AuditActivityNameServiceAccountManagement                       AuditActivityName = "Service Account Management"
	AuditActivityNameServiceAccountUpdate                           AuditActivityName = "Service Account Update"
	AuditActivityNameSharedSsoSiblingWorkspacesUpdate               AuditActivityName = "Shared SSO Sibling Workspaces Update"
	AuditActivityNameSignInRequestApproval                          AuditActivityName = "Sign-In Request Approval"
	AuditActivityNameSignInRequestDecline                           AuditActivityName = "Sign-In Request Decline"
	AuditActivityNameSmbFreeTrialActivation                         AuditActivityName = "SMB Free Trial Activation"
	AuditActivityNameSmbResellerSubscriptionLifecycle               AuditActivityName = "SMB Reseller Subscription Lifecycle"
	AuditActivityNameSmbResellerTierAssignment                      AuditActivityName = "SMB Reseller Tier Assignment"
	AuditActivityNameSmbResellerTierChange                          AuditActivityName = "SMB Reseller Tier Change"
	AuditActivityNameSsoManagement                                  AuditActivityName = "SSO Management"
	AuditActivityNameSsoProviderConfiguration                       AuditActivityName = "SSO Provider Configuration"
	AuditActivityNameSsoProviderRemoval                             AuditActivityName = "SSO Provider Removal"
	AuditActivityNameUpdateWorkspaceCustomRole                      AuditActivityName = "Update workspace custom role"
	AuditActivityNameUserApiKeyCreation                             AuditActivityName = "User API Key Creation"
	AuditActivityNameUserApiKeyDeletion                             AuditActivityName = "User API Key Deletion"
	AuditActivityNameUserApiKeyManagement                           AuditActivityName = "User API Key Management"
	AuditActivityNameUserApiKeyUpdate                               AuditActivityName = "User API Key Update"
	AuditActivityNameUserAutoProvisioningDisabled                   AuditActivityName = "User Auto-Provisioning Disabled"
	AuditActivityNameUserAutoProvisioningEnabled                    AuditActivityName = "User Auto-Provisioning Enabled"
	AuditActivityNameUserCreation                                   AuditActivityName = "User Creation"
	AuditActivityNameUserMetadataUpdate                             AuditActivityName = "User Metadata Update"
	AuditActivityNameUserSelfDeletion                               AuditActivityName = "User Self-Deletion"
	AuditActivityNameUserSharingPreferencesUpdate                   AuditActivityName = "User Sharing Preferences Update"
	AuditActivityNameUserSharingPreferencesUpdateAdmin              AuditActivityName = "User Sharing Preferences Update (Admin)"
	AuditActivityNameWebhookCreation                                AuditActivityName = "Webhook Creation"
	AuditActivityNameWebhookDeletion                                AuditActivityName = "Webhook Deletion"
	AuditActivityNameWebhookManagement                              AuditActivityName = "Webhook Management"
	AuditActivityNameWebhookUpdate                                  AuditActivityName = "Webhook Update"
	AuditActivityNameWorkloadIdentityFederationAssignmentCreation   AuditActivityName = "Workload Identity Federation Assignment Creation"
	AuditActivityNameWorkloadIdentityFederationAssignmentDeletion   AuditActivityName = "Workload Identity Federation Assignment Deletion"
	AuditActivityNameWorkloadIdentityFederationAssignmentManagement AuditActivityName = "Workload Identity Federation Assignment Management"
	AuditActivityNameWorkloadIdentityFederationPoolCreation         AuditActivityName = "Workload Identity Federation Pool Creation"
	AuditActivityNameWorkloadIdentityFederationPoolDeletion         AuditActivityName = "Workload Identity Federation Pool Deletion"
	AuditActivityNameWorkloadIdentityFederationPoolManagement       AuditActivityName = "Workload Identity Federation Pool Management"
	AuditActivityNameWorkloadIdentityFederationPoolUpdate           AuditActivityName = "Workload Identity Federation Pool Update"
	AuditActivityNameWorkloadIdentityFederationTokenExchange        AuditActivityName = "Workload Identity Federation Token Exchange"
	AuditActivityNameWorkspaceCreation                              AuditActivityName = "Workspace Creation"
	AuditActivityNameWorkspaceDeletion                              AuditActivityName = "Workspace Deletion"
	AuditActivityNameWorkspaceDeletionCancellation                  AuditActivityName = "Workspace Deletion Cancellation"
	AuditActivityNameWorkspaceInvitationAccepted                    AuditActivityName = "Workspace Invitation Accepted"
	AuditActivityNameWorkspaceInvitationDeclined                    AuditActivityName = "Workspace Invitation Declined"
	AuditActivityNameWorkspaceInvitationManagement                  AuditActivityName = "Workspace Invitation Management"
	AuditActivityNameWorkspaceInvitationReminderSent                AuditActivityName = "Workspace Invitation Reminder Sent"
	AuditActivityNameWorkspaceInvitationRevoked                     AuditActivityName = "Workspace Invitation Revoked"
	AuditActivityNameWorkspaceInvitationSent                        AuditActivityName = "Workspace Invitation Sent"
	AuditActivityNameWorkspaceMemberRemoval                         AuditActivityName = "Workspace Member Removal"
	AuditActivityNameWorkspaceMemberResourceTransfer                AuditActivityName = "Workspace Member Resource Transfer"
	AuditActivityNameWorkspaceMemberUpdate                          AuditActivityName = "Workspace Member Update"
	AuditActivityNameWorkspaceMemberUsageLimitUpdate                AuditActivityName = "Workspace Member Usage Limit Update"
	AuditActivityNameWorkspaceMetadataUpdate                        AuditActivityName = "Workspace Metadata Update"
	AuditActivityNameWorkspaceOwnershipTransfer                     AuditActivityName = "Workspace Ownership Transfer"
	AuditActivityNameWorkspaceSignIn                                AuditActivityName = "Workspace Sign-In"
)

func NewAuditActivityNameFromString(s string) (AuditActivityName, error) {
	switch s {
	case "Account Security Change":
		return AuditActivityNameAccountSecurityChange, nil
	case "Agents User Memory Entity Deleted":
		return AuditActivityNameAgentsUserMemoryEntityDeleted, nil
	case "Agents User Memory Erased":
		return AuditActivityNameAgentsUserMemoryErased, nil
	case "Agents User Memory Exported":
		return AuditActivityNameAgentsUserMemoryExported, nil
	case "Agents User Memory Fact Deleted":
		return AuditActivityNameAgentsUserMemoryFactDeleted, nil
	case "Agents User Memory Viewed":
		return AuditActivityNameAgentsUserMemoryViewed, nil
	case "Agents User Properties Reordered":
		return AuditActivityNameAgentsUserPropertiesReordered, nil
	case "Agents User Property Archived":
		return AuditActivityNameAgentsUserPropertyArchived, nil
	case "Agents User Property Created":
		return AuditActivityNameAgentsUserPropertyCreated, nil
	case "Agents User Property Unarchived":
		return AuditActivityNameAgentsUserPropertyUnarchived, nil
	case "Agents User Property Updated":
		return AuditActivityNameAgentsUserPropertyUpdated, nil
	case "API Key Creation":
		return AuditActivityNameApiKeyCreation, nil
	case "API Key Deletion":
		return AuditActivityNameApiKeyDeletion, nil
	case "API Key Management":
		return AuditActivityNameApiKeyManagement, nil
	case "API Key Update":
		return AuditActivityNameApiKeyUpdate, nil
	case "Architect Member Spend Limits Update":
		return AuditActivityNameArchitectMemberSpendLimitsUpdate, nil
	case "Auth Connection Creation":
		return AuditActivityNameAuthConnectionCreation, nil
	case "Auth Connection Deletion":
		return AuditActivityNameAuthConnectionDeletion, nil
	case "Auth Connection Management":
		return AuditActivityNameAuthConnectionManagement, nil
	case "Auth Connection Update":
		return AuditActivityNameAuthConnectionUpdate, nil
	case "Billing Management":
		return AuditActivityNameBillingManagement, nil
	case "Child Workspace Resource Override":
		return AuditActivityNameChildWorkspaceResourceOverride, nil
	case "Content Generation Terms Update":
		return AuditActivityNameContentGenerationTermsUpdate, nil
	case "Create workspace custom role":
		return AuditActivityNameCreateWorkspaceCustomRole, nil
	case "Custom Auth Token Created":
		return AuditActivityNameCustomAuthTokenCreated, nil
	case "Data Export Request":
		return AuditActivityNameDataExportRequest, nil
	case "Delete Scribe meeting":
		return AuditActivityNameDeleteScribeMeeting, nil
	case "Delete workspace custom role":
		return AuditActivityNameDeleteWorkspaceCustomRole, nil
	case "Domain Verification Addition":
		return AuditActivityNameDomainVerificationAddition, nil
	case "Domain Verification Management":
		return AuditActivityNameDomainVerificationManagement, nil
	case "Domain Verification Removal":
		return AuditActivityNameDomainVerificationRemoval, nil
	case "Domain Verification Status Update":
		return AuditActivityNameDomainVerificationStatusUpdate, nil
	case "Encryption Key Deleted":
		return AuditActivityNameEncryptionKeyDeleted, nil
	case "Encryption Key Paused":
		return AuditActivityNameEncryptionKeyPaused, nil
	case "Encryption Key Registered":
		return AuditActivityNameEncryptionKeyRegistered, nil
	case "Encryption Key Unpaused":
		return AuditActivityNameEncryptionKeyUnpaused, nil
	case "Enterprise Copying Toggle":
		return AuditActivityNameEnterpriseCopyingToggle, nil
	case "Gift Credit Code Applied":
		return AuditActivityNameGiftCreditCodeApplied, nil
	case "Group Creation":
		return AuditActivityNameGroupCreation, nil
	case "Group Deletion":
		return AuditActivityNameGroupDeletion, nil
	case "Group Management":
		return AuditActivityNameGroupManagement, nil
	case "Group Member Addition":
		return AuditActivityNameGroupMemberAddition, nil
	case "Group Member Removal":
		return AuditActivityNameGroupMemberRemoval, nil
	case "Group to Billing Group Conversion":
		return AuditActivityNameGroupToBillingGroupConversion, nil
	case "Group Update":
		return AuditActivityNameGroupUpdate, nil
	case "Identity Provider Update":
		return AuditActivityNameIdentityProviderUpdate, nil
	case "Influencer Code Applied":
		return AuditActivityNameInfluencerCodeApplied, nil
	case "In-App Term Acceptance":
		return AuditActivityNameInAppTermAcceptance, nil
	case "Logon":
		return AuditActivityNameLogon, nil
	case "MFA Factor Enrollment":
		return AuditActivityNameMfaFactorEnrollment, nil
	case "MFA Factor Enrollment Reverted (unverified)":
		return AuditActivityNameMfaFactorEnrollmentRevertedUnverified, nil
	case "MFA Factor Removal":
		return AuditActivityNameMfaFactorRemoval, nil
	case "MFA Removal via Backup Code":
		return AuditActivityNameMfaRemovalViaBackupCode, nil
	case "Model Approvals Update":
		return AuditActivityNameModelApprovalsUpdate, nil
	case "Notification Preference Creation":
		return AuditActivityNameNotificationPreferenceCreation, nil
	case "Notification Preference Deletion":
		return AuditActivityNameNotificationPreferenceDeletion, nil
	case "Notification Preference Management":
		return AuditActivityNameNotificationPreferenceManagement, nil
	case "Notification Preference Update":
		return AuditActivityNameNotificationPreferenceUpdate, nil
	case "OAuth App Authorization Grant":
		return AuditActivityNameOAuthAppAuthorizationGrant, nil
	case "OAuth App Authorization Management":
		return AuditActivityNameOAuthAppAuthorizationManagement, nil
	case "OAuth App Authorization Revocation":
		return AuditActivityNameOAuthAppAuthorizationRevocation, nil
	case "Password Reset Email Request":
		return AuditActivityNamePasswordResetEmailRequest, nil
	case "Password Reset Link Generation":
		return AuditActivityNamePasswordResetLinkGeneration, nil
	case "Reporting Workspace SSO Inheritance Update":
		return AuditActivityNameReportingWorkspaceSsoInheritanceUpdate, nil
	case "Resource Access Granted":
		return AuditActivityNameResourceAccessGranted, nil
	case "Resource Access Management":
		return AuditActivityNameResourceAccessManagement, nil
	case "Resource Access Revoked":
		return AuditActivityNameResourceAccessRevoked, nil
	case "Resource Copied to Workspace":
		return AuditActivityNameResourceCopiedToWorkspace, nil
	case "Resource Locator Creation":
		return AuditActivityNameResourceLocatorCreation, nil
	case "Resource Locator Deletion":
		return AuditActivityNameResourceLocatorDeletion, nil
	case "Resource Locator Management":
		return AuditActivityNameResourceLocatorManagement, nil
	case "Resource Locator Update":
		return AuditActivityNameResourceLocatorUpdate, nil
	case "Resource Ownership Transferred":
		return AuditActivityNameResourceOwnershipTransferred, nil
	case "SCIM Group Creation":
		return AuditActivityNameScimGroupCreation, nil
	case "SCIM Group Deletion":
		return AuditActivityNameScimGroupDeletion, nil
	case "SCIM Group Management":
		return AuditActivityNameScimGroupManagement, nil
	case "SCIM Group Update":
		return AuditActivityNameScimGroupUpdate, nil
	case "SCIM Provisioning Seat Type Updated":
		return AuditActivityNameScimProvisioningSeatTypeUpdated, nil
	case "SCIM Token Deleted":
		return AuditActivityNameScimTokenDeleted, nil
	case "SCIM Token Generated":
		return AuditActivityNameScimTokenGenerated, nil
	case "SCIM Token Management":
		return AuditActivityNameScimTokenManagement, nil
	case "SCIM Token Rotated":
		return AuditActivityNameScimTokenRotated, nil
	case "SCIM User Creation":
		return AuditActivityNameScimUserCreation, nil
	case "SCIM User Deletion":
		return AuditActivityNameScimUserDeletion, nil
	case "SCIM User Management":
		return AuditActivityNameScimUserManagement, nil
	case "SCIM User Update":
		return AuditActivityNameScimUserUpdate, nil
	case "Sensitive Action Verification":
		return AuditActivityNameSensitiveActionVerification, nil
	case "Service Account Creation":
		return AuditActivityNameServiceAccountCreation, nil
	case "Service Account Deletion":
		return AuditActivityNameServiceAccountDeletion, nil
	case "Service Account Management":
		return AuditActivityNameServiceAccountManagement, nil
	case "Service Account Update":
		return AuditActivityNameServiceAccountUpdate, nil
	case "Shared SSO Sibling Workspaces Update":
		return AuditActivityNameSharedSsoSiblingWorkspacesUpdate, nil
	case "Sign-In Request Approval":
		return AuditActivityNameSignInRequestApproval, nil
	case "Sign-In Request Decline":
		return AuditActivityNameSignInRequestDecline, nil
	case "SMB Free Trial Activation":
		return AuditActivityNameSmbFreeTrialActivation, nil
	case "SMB Reseller Subscription Lifecycle":
		return AuditActivityNameSmbResellerSubscriptionLifecycle, nil
	case "SMB Reseller Tier Assignment":
		return AuditActivityNameSmbResellerTierAssignment, nil
	case "SMB Reseller Tier Change":
		return AuditActivityNameSmbResellerTierChange, nil
	case "SSO Management":
		return AuditActivityNameSsoManagement, nil
	case "SSO Provider Configuration":
		return AuditActivityNameSsoProviderConfiguration, nil
	case "SSO Provider Removal":
		return AuditActivityNameSsoProviderRemoval, nil
	case "Update workspace custom role":
		return AuditActivityNameUpdateWorkspaceCustomRole, nil
	case "User API Key Creation":
		return AuditActivityNameUserApiKeyCreation, nil
	case "User API Key Deletion":
		return AuditActivityNameUserApiKeyDeletion, nil
	case "User API Key Management":
		return AuditActivityNameUserApiKeyManagement, nil
	case "User API Key Update":
		return AuditActivityNameUserApiKeyUpdate, nil
	case "User Auto-Provisioning Disabled":
		return AuditActivityNameUserAutoProvisioningDisabled, nil
	case "User Auto-Provisioning Enabled":
		return AuditActivityNameUserAutoProvisioningEnabled, nil
	case "User Creation":
		return AuditActivityNameUserCreation, nil
	case "User Metadata Update":
		return AuditActivityNameUserMetadataUpdate, nil
	case "User Self-Deletion":
		return AuditActivityNameUserSelfDeletion, nil
	case "User Sharing Preferences Update":
		return AuditActivityNameUserSharingPreferencesUpdate, nil
	case "User Sharing Preferences Update (Admin)":
		return AuditActivityNameUserSharingPreferencesUpdateAdmin, nil
	case "Webhook Creation":
		return AuditActivityNameWebhookCreation, nil
	case "Webhook Deletion":
		return AuditActivityNameWebhookDeletion, nil
	case "Webhook Management":
		return AuditActivityNameWebhookManagement, nil
	case "Webhook Update":
		return AuditActivityNameWebhookUpdate, nil
	case "Workload Identity Federation Assignment Creation":
		return AuditActivityNameWorkloadIdentityFederationAssignmentCreation, nil
	case "Workload Identity Federation Assignment Deletion":
		return AuditActivityNameWorkloadIdentityFederationAssignmentDeletion, nil
	case "Workload Identity Federation Assignment Management":
		return AuditActivityNameWorkloadIdentityFederationAssignmentManagement, nil
	case "Workload Identity Federation Pool Creation":
		return AuditActivityNameWorkloadIdentityFederationPoolCreation, nil
	case "Workload Identity Federation Pool Deletion":
		return AuditActivityNameWorkloadIdentityFederationPoolDeletion, nil
	case "Workload Identity Federation Pool Management":
		return AuditActivityNameWorkloadIdentityFederationPoolManagement, nil
	case "Workload Identity Federation Pool Update":
		return AuditActivityNameWorkloadIdentityFederationPoolUpdate, nil
	case "Workload Identity Federation Token Exchange":
		return AuditActivityNameWorkloadIdentityFederationTokenExchange, nil
	case "Workspace Creation":
		return AuditActivityNameWorkspaceCreation, nil
	case "Workspace Deletion":
		return AuditActivityNameWorkspaceDeletion, nil
	case "Workspace Deletion Cancellation":
		return AuditActivityNameWorkspaceDeletionCancellation, nil
	case "Workspace Invitation Accepted":
		return AuditActivityNameWorkspaceInvitationAccepted, nil
	case "Workspace Invitation Declined":
		return AuditActivityNameWorkspaceInvitationDeclined, nil
	case "Workspace Invitation Management":
		return AuditActivityNameWorkspaceInvitationManagement, nil
	case "Workspace Invitation Reminder Sent":
		return AuditActivityNameWorkspaceInvitationReminderSent, nil
	case "Workspace Invitation Revoked":
		return AuditActivityNameWorkspaceInvitationRevoked, nil
	case "Workspace Invitation Sent":
		return AuditActivityNameWorkspaceInvitationSent, nil
	case "Workspace Member Removal":
		return AuditActivityNameWorkspaceMemberRemoval, nil
	case "Workspace Member Resource Transfer":
		return AuditActivityNameWorkspaceMemberResourceTransfer, nil
	case "Workspace Member Update":
		return AuditActivityNameWorkspaceMemberUpdate, nil
	case "Workspace Member Usage Limit Update":
		return AuditActivityNameWorkspaceMemberUsageLimitUpdate, nil
	case "Workspace Metadata Update":
		return AuditActivityNameWorkspaceMetadataUpdate, nil
	case "Workspace Ownership Transfer":
		return AuditActivityNameWorkspaceOwnershipTransfer, nil
	case "Workspace Sign-In":
		return AuditActivityNameWorkspaceSignIn, nil
	}
	var t AuditActivityName
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AuditActivityName) Ptr() *AuditActivityName {
	return &e
}
