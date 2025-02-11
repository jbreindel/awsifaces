// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/verifiedpermissions"
)

// VerifiedpermissionsClient ...
type VerifiedpermissionsClient interface {
	Options() verifiedpermissions.Options
	BatchIsAuthorized(ctx context.Context, params *verifiedpermissions.BatchIsAuthorizedInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.BatchIsAuthorizedOutput, error)
	BatchIsAuthorizedWithToken(ctx context.Context, params *verifiedpermissions.BatchIsAuthorizedWithTokenInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.BatchIsAuthorizedWithTokenOutput, error)
	CreateIdentitySource(ctx context.Context, params *verifiedpermissions.CreateIdentitySourceInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.CreateIdentitySourceOutput, error)
	CreatePolicy(ctx context.Context, params *verifiedpermissions.CreatePolicyInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.CreatePolicyOutput, error)
	CreatePolicyStore(ctx context.Context, params *verifiedpermissions.CreatePolicyStoreInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.CreatePolicyStoreOutput, error)
	CreatePolicyTemplate(ctx context.Context, params *verifiedpermissions.CreatePolicyTemplateInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.CreatePolicyTemplateOutput, error)
	DeleteIdentitySource(ctx context.Context, params *verifiedpermissions.DeleteIdentitySourceInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.DeleteIdentitySourceOutput, error)
	DeletePolicy(ctx context.Context, params *verifiedpermissions.DeletePolicyInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.DeletePolicyOutput, error)
	DeletePolicyStore(ctx context.Context, params *verifiedpermissions.DeletePolicyStoreInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.DeletePolicyStoreOutput, error)
	DeletePolicyTemplate(ctx context.Context, params *verifiedpermissions.DeletePolicyTemplateInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.DeletePolicyTemplateOutput, error)
	GetIdentitySource(ctx context.Context, params *verifiedpermissions.GetIdentitySourceInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.GetIdentitySourceOutput, error)
	GetPolicy(ctx context.Context, params *verifiedpermissions.GetPolicyInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.GetPolicyOutput, error)
	GetPolicyStore(ctx context.Context, params *verifiedpermissions.GetPolicyStoreInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.GetPolicyStoreOutput, error)
	GetPolicyTemplate(ctx context.Context, params *verifiedpermissions.GetPolicyTemplateInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.GetPolicyTemplateOutput, error)
	GetSchema(ctx context.Context, params *verifiedpermissions.GetSchemaInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.GetSchemaOutput, error)
	IsAuthorized(ctx context.Context, params *verifiedpermissions.IsAuthorizedInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.IsAuthorizedOutput, error)
	IsAuthorizedWithToken(ctx context.Context, params *verifiedpermissions.IsAuthorizedWithTokenInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.IsAuthorizedWithTokenOutput, error)
	ListIdentitySources(ctx context.Context, params *verifiedpermissions.ListIdentitySourcesInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.ListIdentitySourcesOutput, error)
	ListPolicies(ctx context.Context, params *verifiedpermissions.ListPoliciesInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.ListPoliciesOutput, error)
	ListPolicyStores(ctx context.Context, params *verifiedpermissions.ListPolicyStoresInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.ListPolicyStoresOutput, error)
	ListPolicyTemplates(ctx context.Context, params *verifiedpermissions.ListPolicyTemplatesInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.ListPolicyTemplatesOutput, error)
	PutSchema(ctx context.Context, params *verifiedpermissions.PutSchemaInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.PutSchemaOutput, error)
	UpdateIdentitySource(ctx context.Context, params *verifiedpermissions.UpdateIdentitySourceInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.UpdateIdentitySourceOutput, error)
	UpdatePolicy(ctx context.Context, params *verifiedpermissions.UpdatePolicyInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.UpdatePolicyOutput, error)
	UpdatePolicyStore(ctx context.Context, params *verifiedpermissions.UpdatePolicyStoreInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.UpdatePolicyStoreOutput, error)
	UpdatePolicyTemplate(ctx context.Context, params *verifiedpermissions.UpdatePolicyTemplateInput, optFns ...func(*verifiedpermissions.Options)) (*verifiedpermissions.UpdatePolicyTemplateOutput, error)
}
