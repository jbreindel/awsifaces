// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/supplychain"
)

// SupplychainClient ...
type SupplychainClient interface {
	Options() supplychain.Options
	CreateBillOfMaterialsImportJob(ctx context.Context, params *supplychain.CreateBillOfMaterialsImportJobInput, optFns ...func(*supplychain.Options)) (*supplychain.CreateBillOfMaterialsImportJobOutput, error)
	GetBillOfMaterialsImportJob(ctx context.Context, params *supplychain.GetBillOfMaterialsImportJobInput, optFns ...func(*supplychain.Options)) (*supplychain.GetBillOfMaterialsImportJobOutput, error)
	SendDataIntegrationEvent(ctx context.Context, params *supplychain.SendDataIntegrationEventInput, optFns ...func(*supplychain.Options)) (*supplychain.SendDataIntegrationEventOutput, error)
}
