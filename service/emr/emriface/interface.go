package emriface

import (
	"github.com/datacratic/aws-sdk-go/service/emr"
)

type EMRAPI interface {
	AddInstanceGroups(*emr.AddInstanceGroupsInput) (*emr.AddInstanceGroupsOutput, error)

	AddJobFlowSteps(*emr.AddJobFlowStepsInput) (*emr.AddJobFlowStepsOutput, error)

	AddTags(*emr.AddTagsInput) (*emr.AddTagsOutput, error)

	DescribeCluster(*emr.DescribeClusterInput) (*emr.DescribeClusterOutput, error)

	DescribeJobFlows(*emr.DescribeJobFlowsInput) (*emr.DescribeJobFlowsOutput, error)

	DescribeStep(*emr.DescribeStepInput) (*emr.DescribeStepOutput, error)

	ListBootstrapActions(*emr.ListBootstrapActionsInput) (*emr.ListBootstrapActionsOutput, error)

	ListClusters(*emr.ListClustersInput) (*emr.ListClustersOutput, error)

	ListInstanceGroups(*emr.ListInstanceGroupsInput) (*emr.ListInstanceGroupsOutput, error)

	ListInstances(*emr.ListInstancesInput) (*emr.ListInstancesOutput, error)

	ListSteps(*emr.ListStepsInput) (*emr.ListStepsOutput, error)

	ModifyInstanceGroups(*emr.ModifyInstanceGroupsInput) (*emr.ModifyInstanceGroupsOutput, error)

	RemoveTags(*emr.RemoveTagsInput) (*emr.RemoveTagsOutput, error)

	RunJobFlow(*emr.RunJobFlowInput) (*emr.RunJobFlowOutput, error)

	SetTerminationProtection(*emr.SetTerminationProtectionInput) (*emr.SetTerminationProtectionOutput, error)

	SetVisibleToAllUsers(*emr.SetVisibleToAllUsersInput) (*emr.SetVisibleToAllUsersOutput, error)

	TerminateJobFlows(*emr.TerminateJobFlowsInput) (*emr.TerminateJobFlowsOutput, error)
}