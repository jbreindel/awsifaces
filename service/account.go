// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/account"
)

// AccountClient ...
type AccountClient interface {
	Options() account.Options
	DeleteAlternateContact(ctx context.Context, params *account.DeleteAlternateContactInput, optFns ...func(*account.Options)) (*account.DeleteAlternateContactOutput, error)
	DisableRegion(ctx context.Context, params *account.DisableRegionInput, optFns ...func(*account.Options)) (*account.DisableRegionOutput, error)
	EnableRegion(ctx context.Context, params *account.EnableRegionInput, optFns ...func(*account.Options)) (*account.EnableRegionOutput, error)
	GetAlternateContact(ctx context.Context, params *account.GetAlternateContactInput, optFns ...func(*account.Options)) (*account.GetAlternateContactOutput, error)
	GetContactInformation(ctx context.Context, params *account.GetContactInformationInput, optFns ...func(*account.Options)) (*account.GetContactInformationOutput, error)
	GetRegionOptStatus(ctx context.Context, params *account.GetRegionOptStatusInput, optFns ...func(*account.Options)) (*account.GetRegionOptStatusOutput, error)
	ListRegions(ctx context.Context, params *account.ListRegionsInput, optFns ...func(*account.Options)) (*account.ListRegionsOutput, error)
	PutAlternateContact(ctx context.Context, params *account.PutAlternateContactInput, optFns ...func(*account.Options)) (*account.PutAlternateContactOutput, error)
	PutContactInformation(ctx context.Context, params *account.PutContactInformationInput, optFns ...func(*account.Options)) (*account.PutContactInformationOutput, error)
}
