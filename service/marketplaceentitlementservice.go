// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/marketplaceentitlementservice"
)

// MarketplaceentitlementserviceClient ...
type MarketplaceentitlementserviceClient interface {
	Options() marketplaceentitlementservice.Options
	GetEntitlements(ctx context.Context, params *marketplaceentitlementservice.GetEntitlementsInput, optFns ...func(*marketplaceentitlementservice.Options)) (*marketplaceentitlementservice.GetEntitlementsOutput, error)
}
