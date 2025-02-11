// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/privatenetworks"
)

// PrivatenetworksClient ...
type PrivatenetworksClient interface {
	Options() privatenetworks.Options
	AcknowledgeOrderReceipt(ctx context.Context, params *privatenetworks.AcknowledgeOrderReceiptInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.AcknowledgeOrderReceiptOutput, error)
	ActivateDeviceIdentifier(ctx context.Context, params *privatenetworks.ActivateDeviceIdentifierInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.ActivateDeviceIdentifierOutput, error)
	ActivateNetworkSite(ctx context.Context, params *privatenetworks.ActivateNetworkSiteInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.ActivateNetworkSiteOutput, error)
	ConfigureAccessPoint(ctx context.Context, params *privatenetworks.ConfigureAccessPointInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.ConfigureAccessPointOutput, error)
	CreateNetwork(ctx context.Context, params *privatenetworks.CreateNetworkInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.CreateNetworkOutput, error)
	CreateNetworkSite(ctx context.Context, params *privatenetworks.CreateNetworkSiteInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.CreateNetworkSiteOutput, error)
	DeactivateDeviceIdentifier(ctx context.Context, params *privatenetworks.DeactivateDeviceIdentifierInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.DeactivateDeviceIdentifierOutput, error)
	DeleteNetwork(ctx context.Context, params *privatenetworks.DeleteNetworkInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.DeleteNetworkOutput, error)
	DeleteNetworkSite(ctx context.Context, params *privatenetworks.DeleteNetworkSiteInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.DeleteNetworkSiteOutput, error)
	GetDeviceIdentifier(ctx context.Context, params *privatenetworks.GetDeviceIdentifierInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.GetDeviceIdentifierOutput, error)
	GetNetwork(ctx context.Context, params *privatenetworks.GetNetworkInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.GetNetworkOutput, error)
	GetNetworkResource(ctx context.Context, params *privatenetworks.GetNetworkResourceInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.GetNetworkResourceOutput, error)
	GetNetworkSite(ctx context.Context, params *privatenetworks.GetNetworkSiteInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.GetNetworkSiteOutput, error)
	GetOrder(ctx context.Context, params *privatenetworks.GetOrderInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.GetOrderOutput, error)
	ListDeviceIdentifiers(ctx context.Context, params *privatenetworks.ListDeviceIdentifiersInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.ListDeviceIdentifiersOutput, error)
	ListNetworkResources(ctx context.Context, params *privatenetworks.ListNetworkResourcesInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.ListNetworkResourcesOutput, error)
	ListNetworkSites(ctx context.Context, params *privatenetworks.ListNetworkSitesInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.ListNetworkSitesOutput, error)
	ListNetworks(ctx context.Context, params *privatenetworks.ListNetworksInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.ListNetworksOutput, error)
	ListOrders(ctx context.Context, params *privatenetworks.ListOrdersInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.ListOrdersOutput, error)
	ListTagsForResource(ctx context.Context, params *privatenetworks.ListTagsForResourceInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.ListTagsForResourceOutput, error)
	Ping(ctx context.Context, params *privatenetworks.PingInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.PingOutput, error)
	StartNetworkResourceUpdate(ctx context.Context, params *privatenetworks.StartNetworkResourceUpdateInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.StartNetworkResourceUpdateOutput, error)
	TagResource(ctx context.Context, params *privatenetworks.TagResourceInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.TagResourceOutput, error)
	UntagResource(ctx context.Context, params *privatenetworks.UntagResourceInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.UntagResourceOutput, error)
	UpdateNetworkSite(ctx context.Context, params *privatenetworks.UpdateNetworkSiteInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.UpdateNetworkSiteOutput, error)
	UpdateNetworkSitePlan(ctx context.Context, params *privatenetworks.UpdateNetworkSitePlanInput, optFns ...func(*privatenetworks.Options)) (*privatenetworks.UpdateNetworkSitePlanOutput, error)
}
