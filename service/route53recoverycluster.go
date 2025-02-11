// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/route53recoverycluster"
)

// Route53recoveryclusterClient ...
type Route53recoveryclusterClient interface {
	Options() route53recoverycluster.Options
	GetRoutingControlState(ctx context.Context, params *route53recoverycluster.GetRoutingControlStateInput, optFns ...func(*route53recoverycluster.Options)) (*route53recoverycluster.GetRoutingControlStateOutput, error)
	ListRoutingControls(ctx context.Context, params *route53recoverycluster.ListRoutingControlsInput, optFns ...func(*route53recoverycluster.Options)) (*route53recoverycluster.ListRoutingControlsOutput, error)
	UpdateRoutingControlState(ctx context.Context, params *route53recoverycluster.UpdateRoutingControlStateInput, optFns ...func(*route53recoverycluster.Options)) (*route53recoverycluster.UpdateRoutingControlStateOutput, error)
	UpdateRoutingControlStates(ctx context.Context, params *route53recoverycluster.UpdateRoutingControlStatesInput, optFns ...func(*route53recoverycluster.Options)) (*route53recoverycluster.UpdateRoutingControlStatesOutput, error)
}
