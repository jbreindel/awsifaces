// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/auditmanager"
)

// AuditmanagerClient ...
type AuditmanagerClient interface {
	Options() auditmanager.Options
	AssociateAssessmentReportEvidenceFolder(ctx context.Context, params *auditmanager.AssociateAssessmentReportEvidenceFolderInput, optFns ...func(*auditmanager.Options)) (*auditmanager.AssociateAssessmentReportEvidenceFolderOutput, error)
	BatchAssociateAssessmentReportEvidence(ctx context.Context, params *auditmanager.BatchAssociateAssessmentReportEvidenceInput, optFns ...func(*auditmanager.Options)) (*auditmanager.BatchAssociateAssessmentReportEvidenceOutput, error)
	BatchCreateDelegationByAssessment(ctx context.Context, params *auditmanager.BatchCreateDelegationByAssessmentInput, optFns ...func(*auditmanager.Options)) (*auditmanager.BatchCreateDelegationByAssessmentOutput, error)
	BatchDeleteDelegationByAssessment(ctx context.Context, params *auditmanager.BatchDeleteDelegationByAssessmentInput, optFns ...func(*auditmanager.Options)) (*auditmanager.BatchDeleteDelegationByAssessmentOutput, error)
	BatchDisassociateAssessmentReportEvidence(ctx context.Context, params *auditmanager.BatchDisassociateAssessmentReportEvidenceInput, optFns ...func(*auditmanager.Options)) (*auditmanager.BatchDisassociateAssessmentReportEvidenceOutput, error)
	BatchImportEvidenceToAssessmentControl(ctx context.Context, params *auditmanager.BatchImportEvidenceToAssessmentControlInput, optFns ...func(*auditmanager.Options)) (*auditmanager.BatchImportEvidenceToAssessmentControlOutput, error)
	CreateAssessment(ctx context.Context, params *auditmanager.CreateAssessmentInput, optFns ...func(*auditmanager.Options)) (*auditmanager.CreateAssessmentOutput, error)
	CreateAssessmentFramework(ctx context.Context, params *auditmanager.CreateAssessmentFrameworkInput, optFns ...func(*auditmanager.Options)) (*auditmanager.CreateAssessmentFrameworkOutput, error)
	CreateAssessmentReport(ctx context.Context, params *auditmanager.CreateAssessmentReportInput, optFns ...func(*auditmanager.Options)) (*auditmanager.CreateAssessmentReportOutput, error)
	CreateControl(ctx context.Context, params *auditmanager.CreateControlInput, optFns ...func(*auditmanager.Options)) (*auditmanager.CreateControlOutput, error)
	DeleteAssessment(ctx context.Context, params *auditmanager.DeleteAssessmentInput, optFns ...func(*auditmanager.Options)) (*auditmanager.DeleteAssessmentOutput, error)
	DeleteAssessmentFramework(ctx context.Context, params *auditmanager.DeleteAssessmentFrameworkInput, optFns ...func(*auditmanager.Options)) (*auditmanager.DeleteAssessmentFrameworkOutput, error)
	DeleteAssessmentFrameworkShare(ctx context.Context, params *auditmanager.DeleteAssessmentFrameworkShareInput, optFns ...func(*auditmanager.Options)) (*auditmanager.DeleteAssessmentFrameworkShareOutput, error)
	DeleteAssessmentReport(ctx context.Context, params *auditmanager.DeleteAssessmentReportInput, optFns ...func(*auditmanager.Options)) (*auditmanager.DeleteAssessmentReportOutput, error)
	DeleteControl(ctx context.Context, params *auditmanager.DeleteControlInput, optFns ...func(*auditmanager.Options)) (*auditmanager.DeleteControlOutput, error)
	DeregisterAccount(ctx context.Context, params *auditmanager.DeregisterAccountInput, optFns ...func(*auditmanager.Options)) (*auditmanager.DeregisterAccountOutput, error)
	DeregisterOrganizationAdminAccount(ctx context.Context, params *auditmanager.DeregisterOrganizationAdminAccountInput, optFns ...func(*auditmanager.Options)) (*auditmanager.DeregisterOrganizationAdminAccountOutput, error)
	DisassociateAssessmentReportEvidenceFolder(ctx context.Context, params *auditmanager.DisassociateAssessmentReportEvidenceFolderInput, optFns ...func(*auditmanager.Options)) (*auditmanager.DisassociateAssessmentReportEvidenceFolderOutput, error)
	GetAccountStatus(ctx context.Context, params *auditmanager.GetAccountStatusInput, optFns ...func(*auditmanager.Options)) (*auditmanager.GetAccountStatusOutput, error)
	GetAssessment(ctx context.Context, params *auditmanager.GetAssessmentInput, optFns ...func(*auditmanager.Options)) (*auditmanager.GetAssessmentOutput, error)
	GetAssessmentFramework(ctx context.Context, params *auditmanager.GetAssessmentFrameworkInput, optFns ...func(*auditmanager.Options)) (*auditmanager.GetAssessmentFrameworkOutput, error)
	GetAssessmentReportUrl(ctx context.Context, params *auditmanager.GetAssessmentReportUrlInput, optFns ...func(*auditmanager.Options)) (*auditmanager.GetAssessmentReportUrlOutput, error)
	GetChangeLogs(ctx context.Context, params *auditmanager.GetChangeLogsInput, optFns ...func(*auditmanager.Options)) (*auditmanager.GetChangeLogsOutput, error)
	GetControl(ctx context.Context, params *auditmanager.GetControlInput, optFns ...func(*auditmanager.Options)) (*auditmanager.GetControlOutput, error)
	GetDelegations(ctx context.Context, params *auditmanager.GetDelegationsInput, optFns ...func(*auditmanager.Options)) (*auditmanager.GetDelegationsOutput, error)
	GetEvidence(ctx context.Context, params *auditmanager.GetEvidenceInput, optFns ...func(*auditmanager.Options)) (*auditmanager.GetEvidenceOutput, error)
	GetEvidenceByEvidenceFolder(ctx context.Context, params *auditmanager.GetEvidenceByEvidenceFolderInput, optFns ...func(*auditmanager.Options)) (*auditmanager.GetEvidenceByEvidenceFolderOutput, error)
	GetEvidenceFileUploadUrl(ctx context.Context, params *auditmanager.GetEvidenceFileUploadUrlInput, optFns ...func(*auditmanager.Options)) (*auditmanager.GetEvidenceFileUploadUrlOutput, error)
	GetEvidenceFolder(ctx context.Context, params *auditmanager.GetEvidenceFolderInput, optFns ...func(*auditmanager.Options)) (*auditmanager.GetEvidenceFolderOutput, error)
	GetEvidenceFoldersByAssessment(ctx context.Context, params *auditmanager.GetEvidenceFoldersByAssessmentInput, optFns ...func(*auditmanager.Options)) (*auditmanager.GetEvidenceFoldersByAssessmentOutput, error)
	GetEvidenceFoldersByAssessmentControl(ctx context.Context, params *auditmanager.GetEvidenceFoldersByAssessmentControlInput, optFns ...func(*auditmanager.Options)) (*auditmanager.GetEvidenceFoldersByAssessmentControlOutput, error)
	GetInsights(ctx context.Context, params *auditmanager.GetInsightsInput, optFns ...func(*auditmanager.Options)) (*auditmanager.GetInsightsOutput, error)
	GetInsightsByAssessment(ctx context.Context, params *auditmanager.GetInsightsByAssessmentInput, optFns ...func(*auditmanager.Options)) (*auditmanager.GetInsightsByAssessmentOutput, error)
	GetOrganizationAdminAccount(ctx context.Context, params *auditmanager.GetOrganizationAdminAccountInput, optFns ...func(*auditmanager.Options)) (*auditmanager.GetOrganizationAdminAccountOutput, error)
	GetServicesInScope(ctx context.Context, params *auditmanager.GetServicesInScopeInput, optFns ...func(*auditmanager.Options)) (*auditmanager.GetServicesInScopeOutput, error)
	GetSettings(ctx context.Context, params *auditmanager.GetSettingsInput, optFns ...func(*auditmanager.Options)) (*auditmanager.GetSettingsOutput, error)
	ListAssessmentControlInsightsByControlDomain(ctx context.Context, params *auditmanager.ListAssessmentControlInsightsByControlDomainInput, optFns ...func(*auditmanager.Options)) (*auditmanager.ListAssessmentControlInsightsByControlDomainOutput, error)
	ListAssessmentFrameworkShareRequests(ctx context.Context, params *auditmanager.ListAssessmentFrameworkShareRequestsInput, optFns ...func(*auditmanager.Options)) (*auditmanager.ListAssessmentFrameworkShareRequestsOutput, error)
	ListAssessmentFrameworks(ctx context.Context, params *auditmanager.ListAssessmentFrameworksInput, optFns ...func(*auditmanager.Options)) (*auditmanager.ListAssessmentFrameworksOutput, error)
	ListAssessmentReports(ctx context.Context, params *auditmanager.ListAssessmentReportsInput, optFns ...func(*auditmanager.Options)) (*auditmanager.ListAssessmentReportsOutput, error)
	ListAssessments(ctx context.Context, params *auditmanager.ListAssessmentsInput, optFns ...func(*auditmanager.Options)) (*auditmanager.ListAssessmentsOutput, error)
	ListControlDomainInsights(ctx context.Context, params *auditmanager.ListControlDomainInsightsInput, optFns ...func(*auditmanager.Options)) (*auditmanager.ListControlDomainInsightsOutput, error)
	ListControlDomainInsightsByAssessment(ctx context.Context, params *auditmanager.ListControlDomainInsightsByAssessmentInput, optFns ...func(*auditmanager.Options)) (*auditmanager.ListControlDomainInsightsByAssessmentOutput, error)
	ListControlInsightsByControlDomain(ctx context.Context, params *auditmanager.ListControlInsightsByControlDomainInput, optFns ...func(*auditmanager.Options)) (*auditmanager.ListControlInsightsByControlDomainOutput, error)
	ListControls(ctx context.Context, params *auditmanager.ListControlsInput, optFns ...func(*auditmanager.Options)) (*auditmanager.ListControlsOutput, error)
	ListKeywordsForDataSource(ctx context.Context, params *auditmanager.ListKeywordsForDataSourceInput, optFns ...func(*auditmanager.Options)) (*auditmanager.ListKeywordsForDataSourceOutput, error)
	ListNotifications(ctx context.Context, params *auditmanager.ListNotificationsInput, optFns ...func(*auditmanager.Options)) (*auditmanager.ListNotificationsOutput, error)
	ListTagsForResource(ctx context.Context, params *auditmanager.ListTagsForResourceInput, optFns ...func(*auditmanager.Options)) (*auditmanager.ListTagsForResourceOutput, error)
	RegisterAccount(ctx context.Context, params *auditmanager.RegisterAccountInput, optFns ...func(*auditmanager.Options)) (*auditmanager.RegisterAccountOutput, error)
	RegisterOrganizationAdminAccount(ctx context.Context, params *auditmanager.RegisterOrganizationAdminAccountInput, optFns ...func(*auditmanager.Options)) (*auditmanager.RegisterOrganizationAdminAccountOutput, error)
	StartAssessmentFrameworkShare(ctx context.Context, params *auditmanager.StartAssessmentFrameworkShareInput, optFns ...func(*auditmanager.Options)) (*auditmanager.StartAssessmentFrameworkShareOutput, error)
	TagResource(ctx context.Context, params *auditmanager.TagResourceInput, optFns ...func(*auditmanager.Options)) (*auditmanager.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *auditmanager.UntagResourceInput, optFns ...func(*auditmanager.Options)) (*auditmanager.UntagResourceOutput, error)
	UpdateAssessment(ctx context.Context, params *auditmanager.UpdateAssessmentInput, optFns ...func(*auditmanager.Options)) (*auditmanager.UpdateAssessmentOutput, error)
	UpdateAssessmentControl(ctx context.Context, params *auditmanager.UpdateAssessmentControlInput, optFns ...func(*auditmanager.Options)) (*auditmanager.UpdateAssessmentControlOutput, error)
	UpdateAssessmentControlSetStatus(ctx context.Context, params *auditmanager.UpdateAssessmentControlSetStatusInput, optFns ...func(*auditmanager.Options)) (*auditmanager.UpdateAssessmentControlSetStatusOutput, error)
	UpdateAssessmentFramework(ctx context.Context, params *auditmanager.UpdateAssessmentFrameworkInput, optFns ...func(*auditmanager.Options)) (*auditmanager.UpdateAssessmentFrameworkOutput, error)
	UpdateAssessmentFrameworkShare(ctx context.Context, params *auditmanager.UpdateAssessmentFrameworkShareInput, optFns ...func(*auditmanager.Options)) (*auditmanager.UpdateAssessmentFrameworkShareOutput, error)
	UpdateAssessmentStatus(ctx context.Context, params *auditmanager.UpdateAssessmentStatusInput, optFns ...func(*auditmanager.Options)) (*auditmanager.UpdateAssessmentStatusOutput, error)
	UpdateControl(ctx context.Context, params *auditmanager.UpdateControlInput, optFns ...func(*auditmanager.Options)) (*auditmanager.UpdateControlOutput, error)
	UpdateSettings(ctx context.Context, params *auditmanager.UpdateSettingsInput, optFns ...func(*auditmanager.Options)) (*auditmanager.UpdateSettingsOutput, error)
	ValidateAssessmentReportIntegrity(ctx context.Context, params *auditmanager.ValidateAssessmentReportIntegrityInput, optFns ...func(*auditmanager.Options)) (*auditmanager.ValidateAssessmentReportIntegrityOutput, error)
}
