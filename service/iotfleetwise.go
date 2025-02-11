// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise"
)

// IotfleetwiseClient ...
type IotfleetwiseClient interface {
	Options() iotfleetwise.Options
	AssociateVehicleFleet(ctx context.Context, params *iotfleetwise.AssociateVehicleFleetInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.AssociateVehicleFleetOutput, error)
	BatchCreateVehicle(ctx context.Context, params *iotfleetwise.BatchCreateVehicleInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.BatchCreateVehicleOutput, error)
	BatchUpdateVehicle(ctx context.Context, params *iotfleetwise.BatchUpdateVehicleInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.BatchUpdateVehicleOutput, error)
	CreateCampaign(ctx context.Context, params *iotfleetwise.CreateCampaignInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.CreateCampaignOutput, error)
	CreateDecoderManifest(ctx context.Context, params *iotfleetwise.CreateDecoderManifestInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.CreateDecoderManifestOutput, error)
	CreateFleet(ctx context.Context, params *iotfleetwise.CreateFleetInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.CreateFleetOutput, error)
	CreateModelManifest(ctx context.Context, params *iotfleetwise.CreateModelManifestInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.CreateModelManifestOutput, error)
	CreateSignalCatalog(ctx context.Context, params *iotfleetwise.CreateSignalCatalogInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.CreateSignalCatalogOutput, error)
	CreateVehicle(ctx context.Context, params *iotfleetwise.CreateVehicleInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.CreateVehicleOutput, error)
	DeleteCampaign(ctx context.Context, params *iotfleetwise.DeleteCampaignInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.DeleteCampaignOutput, error)
	DeleteDecoderManifest(ctx context.Context, params *iotfleetwise.DeleteDecoderManifestInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.DeleteDecoderManifestOutput, error)
	DeleteFleet(ctx context.Context, params *iotfleetwise.DeleteFleetInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.DeleteFleetOutput, error)
	DeleteModelManifest(ctx context.Context, params *iotfleetwise.DeleteModelManifestInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.DeleteModelManifestOutput, error)
	DeleteSignalCatalog(ctx context.Context, params *iotfleetwise.DeleteSignalCatalogInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.DeleteSignalCatalogOutput, error)
	DeleteVehicle(ctx context.Context, params *iotfleetwise.DeleteVehicleInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.DeleteVehicleOutput, error)
	DisassociateVehicleFleet(ctx context.Context, params *iotfleetwise.DisassociateVehicleFleetInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.DisassociateVehicleFleetOutput, error)
	GetCampaign(ctx context.Context, params *iotfleetwise.GetCampaignInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.GetCampaignOutput, error)
	GetDecoderManifest(ctx context.Context, params *iotfleetwise.GetDecoderManifestInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.GetDecoderManifestOutput, error)
	GetEncryptionConfiguration(ctx context.Context, params *iotfleetwise.GetEncryptionConfigurationInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.GetEncryptionConfigurationOutput, error)
	GetFleet(ctx context.Context, params *iotfleetwise.GetFleetInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.GetFleetOutput, error)
	GetLoggingOptions(ctx context.Context, params *iotfleetwise.GetLoggingOptionsInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.GetLoggingOptionsOutput, error)
	GetModelManifest(ctx context.Context, params *iotfleetwise.GetModelManifestInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.GetModelManifestOutput, error)
	GetRegisterAccountStatus(ctx context.Context, params *iotfleetwise.GetRegisterAccountStatusInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.GetRegisterAccountStatusOutput, error)
	GetSignalCatalog(ctx context.Context, params *iotfleetwise.GetSignalCatalogInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.GetSignalCatalogOutput, error)
	GetVehicle(ctx context.Context, params *iotfleetwise.GetVehicleInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.GetVehicleOutput, error)
	GetVehicleStatus(ctx context.Context, params *iotfleetwise.GetVehicleStatusInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.GetVehicleStatusOutput, error)
	ImportDecoderManifest(ctx context.Context, params *iotfleetwise.ImportDecoderManifestInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.ImportDecoderManifestOutput, error)
	ImportSignalCatalog(ctx context.Context, params *iotfleetwise.ImportSignalCatalogInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.ImportSignalCatalogOutput, error)
	ListCampaigns(ctx context.Context, params *iotfleetwise.ListCampaignsInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.ListCampaignsOutput, error)
	ListDecoderManifestNetworkInterfaces(ctx context.Context, params *iotfleetwise.ListDecoderManifestNetworkInterfacesInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.ListDecoderManifestNetworkInterfacesOutput, error)
	ListDecoderManifestSignals(ctx context.Context, params *iotfleetwise.ListDecoderManifestSignalsInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.ListDecoderManifestSignalsOutput, error)
	ListDecoderManifests(ctx context.Context, params *iotfleetwise.ListDecoderManifestsInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.ListDecoderManifestsOutput, error)
	ListFleets(ctx context.Context, params *iotfleetwise.ListFleetsInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.ListFleetsOutput, error)
	ListFleetsForVehicle(ctx context.Context, params *iotfleetwise.ListFleetsForVehicleInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.ListFleetsForVehicleOutput, error)
	ListModelManifestNodes(ctx context.Context, params *iotfleetwise.ListModelManifestNodesInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.ListModelManifestNodesOutput, error)
	ListModelManifests(ctx context.Context, params *iotfleetwise.ListModelManifestsInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.ListModelManifestsOutput, error)
	ListSignalCatalogNodes(ctx context.Context, params *iotfleetwise.ListSignalCatalogNodesInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.ListSignalCatalogNodesOutput, error)
	ListSignalCatalogs(ctx context.Context, params *iotfleetwise.ListSignalCatalogsInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.ListSignalCatalogsOutput, error)
	ListTagsForResource(ctx context.Context, params *iotfleetwise.ListTagsForResourceInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.ListTagsForResourceOutput, error)
	ListVehicles(ctx context.Context, params *iotfleetwise.ListVehiclesInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.ListVehiclesOutput, error)
	ListVehiclesInFleet(ctx context.Context, params *iotfleetwise.ListVehiclesInFleetInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.ListVehiclesInFleetOutput, error)
	PutEncryptionConfiguration(ctx context.Context, params *iotfleetwise.PutEncryptionConfigurationInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.PutEncryptionConfigurationOutput, error)
	PutLoggingOptions(ctx context.Context, params *iotfleetwise.PutLoggingOptionsInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.PutLoggingOptionsOutput, error)
	RegisterAccount(ctx context.Context, params *iotfleetwise.RegisterAccountInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.RegisterAccountOutput, error)
	TagResource(ctx context.Context, params *iotfleetwise.TagResourceInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *iotfleetwise.UntagResourceInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.UntagResourceOutput, error)
	UpdateCampaign(ctx context.Context, params *iotfleetwise.UpdateCampaignInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.UpdateCampaignOutput, error)
	UpdateDecoderManifest(ctx context.Context, params *iotfleetwise.UpdateDecoderManifestInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.UpdateDecoderManifestOutput, error)
	UpdateFleet(ctx context.Context, params *iotfleetwise.UpdateFleetInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.UpdateFleetOutput, error)
	UpdateModelManifest(ctx context.Context, params *iotfleetwise.UpdateModelManifestInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.UpdateModelManifestOutput, error)
	UpdateSignalCatalog(ctx context.Context, params *iotfleetwise.UpdateSignalCatalogInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.UpdateSignalCatalogOutput, error)
	UpdateVehicle(ctx context.Context, params *iotfleetwise.UpdateVehicleInput, optFns ...func(*iotfleetwise.Options)) (*iotfleetwise.UpdateVehicleOutput, error)
}
