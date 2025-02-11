// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kinesisvideosignaling"
)

// KinesisvideosignalingClient ...
type KinesisvideosignalingClient interface {
	Options() kinesisvideosignaling.Options
	GetIceServerConfig(ctx context.Context, params *kinesisvideosignaling.GetIceServerConfigInput, optFns ...func(*kinesisvideosignaling.Options)) (*kinesisvideosignaling.GetIceServerConfigOutput, error)
	SendAlexaOfferToMaster(ctx context.Context, params *kinesisvideosignaling.SendAlexaOfferToMasterInput, optFns ...func(*kinesisvideosignaling.Options)) (*kinesisvideosignaling.SendAlexaOfferToMasterOutput, error)
}
