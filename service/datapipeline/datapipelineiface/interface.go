package datapipelineiface

import (
	"github.com/datacratic/aws-sdk-go/service/datapipeline"
)

type DataPipelineAPI interface {
	ActivatePipeline(*datapipeline.ActivatePipelineInput) (*datapipeline.ActivatePipelineOutput, error)

	AddTags(*datapipeline.AddTagsInput) (*datapipeline.AddTagsOutput, error)

	CreatePipeline(*datapipeline.CreatePipelineInput) (*datapipeline.CreatePipelineOutput, error)

	DeletePipeline(*datapipeline.DeletePipelineInput) (*datapipeline.DeletePipelineOutput, error)

	DescribeObjects(*datapipeline.DescribeObjectsInput) (*datapipeline.DescribeObjectsOutput, error)

	DescribePipelines(*datapipeline.DescribePipelinesInput) (*datapipeline.DescribePipelinesOutput, error)

	EvaluateExpression(*datapipeline.EvaluateExpressionInput) (*datapipeline.EvaluateExpressionOutput, error)

	GetPipelineDefinition(*datapipeline.GetPipelineDefinitionInput) (*datapipeline.GetPipelineDefinitionOutput, error)

	ListPipelines(*datapipeline.ListPipelinesInput) (*datapipeline.ListPipelinesOutput, error)

	PollForTask(*datapipeline.PollForTaskInput) (*datapipeline.PollForTaskOutput, error)

	PutPipelineDefinition(*datapipeline.PutPipelineDefinitionInput) (*datapipeline.PutPipelineDefinitionOutput, error)

	QueryObjects(*datapipeline.QueryObjectsInput) (*datapipeline.QueryObjectsOutput, error)

	RemoveTags(*datapipeline.RemoveTagsInput) (*datapipeline.RemoveTagsOutput, error)

	ReportTaskProgress(*datapipeline.ReportTaskProgressInput) (*datapipeline.ReportTaskProgressOutput, error)

	ReportTaskRunnerHeartbeat(*datapipeline.ReportTaskRunnerHeartbeatInput) (*datapipeline.ReportTaskRunnerHeartbeatOutput, error)

	SetStatus(*datapipeline.SetStatusInput) (*datapipeline.SetStatusOutput, error)

	SetTaskStatus(*datapipeline.SetTaskStatusInput) (*datapipeline.SetTaskStatusOutput, error)

	ValidatePipelineDefinition(*datapipeline.ValidatePipelineDefinitionInput) (*datapipeline.ValidatePipelineDefinitionOutput, error)
}