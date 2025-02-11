// Code generated by ifacemaker; DO NOT EDIT.

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/arczonalshift"
)

// ArczonalshiftClient ...
type ArczonalshiftClient interface {
	Options() arczonalshift.Options
	CancelZonalShift(ctx context.Context, params *arczonalshift.CancelZonalShiftInput, optFns ...func(*arczonalshift.Options)) (*arczonalshift.CancelZonalShiftOutput, error)
	CreatePracticeRunConfiguration(ctx context.Context, params *arczonalshift.CreatePracticeRunConfigurationInput, optFns ...func(*arczonalshift.Options)) (*arczonalshift.CreatePracticeRunConfigurationOutput, error)
	DeletePracticeRunConfiguration(ctx context.Context, params *arczonalshift.DeletePracticeRunConfigurationInput, optFns ...func(*arczonalshift.Options)) (*arczonalshift.DeletePracticeRunConfigurationOutput, error)
	GetManagedResource(ctx context.Context, params *arczonalshift.GetManagedResourceInput, optFns ...func(*arczonalshift.Options)) (*arczonalshift.GetManagedResourceOutput, error)
	ListAutoshifts(ctx context.Context, params *arczonalshift.ListAutoshiftsInput, optFns ...func(*arczonalshift.Options)) (*arczonalshift.ListAutoshiftsOutput, error)
	ListManagedResources(ctx context.Context, params *arczonalshift.ListManagedResourcesInput, optFns ...func(*arczonalshift.Options)) (*arczonalshift.ListManagedResourcesOutput, error)
	ListZonalShifts(ctx context.Context, params *arczonalshift.ListZonalShiftsInput, optFns ...func(*arczonalshift.Options)) (*arczonalshift.ListZonalShiftsOutput, error)
	StartZonalShift(ctx context.Context, params *arczonalshift.StartZonalShiftInput, optFns ...func(*arczonalshift.Options)) (*arczonalshift.StartZonalShiftOutput, error)
	UpdatePracticeRunConfiguration(ctx context.Context, params *arczonalshift.UpdatePracticeRunConfigurationInput, optFns ...func(*arczonalshift.Options)) (*arczonalshift.UpdatePracticeRunConfigurationOutput, error)
	UpdateZonalAutoshiftConfiguration(ctx context.Context, params *arczonalshift.UpdateZonalAutoshiftConfigurationInput, optFns ...func(*arczonalshift.Options)) (*arczonalshift.UpdateZonalAutoshiftConfigurationOutput, error)
	UpdateZonalShift(ctx context.Context, params *arczonalshift.UpdateZonalShiftInput, optFns ...func(*arczonalshift.Options)) (*arczonalshift.UpdateZonalShiftOutput, error)
}
