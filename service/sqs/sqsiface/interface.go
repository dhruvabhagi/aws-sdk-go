package sqsiface

import (
	"github.com/datacratic/aws-sdk-go/service/sqs"
)

type SQSAPI interface {
	AddPermission(*sqs.AddPermissionInput) (*sqs.AddPermissionOutput, error)

	ChangeMessageVisibility(*sqs.ChangeMessageVisibilityInput) (*sqs.ChangeMessageVisibilityOutput, error)

	ChangeMessageVisibilityBatch(*sqs.ChangeMessageVisibilityBatchInput) (*sqs.ChangeMessageVisibilityBatchOutput, error)

	CreateQueue(*sqs.CreateQueueInput) (*sqs.CreateQueueOutput, error)

	DeleteMessage(*sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error)

	DeleteMessageBatch(*sqs.DeleteMessageBatchInput) (*sqs.DeleteMessageBatchOutput, error)

	DeleteQueue(*sqs.DeleteQueueInput) (*sqs.DeleteQueueOutput, error)

	GetQueueAttributes(*sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error)

	GetQueueURL(*sqs.GetQueueURLInput) (*sqs.GetQueueURLOutput, error)

	ListDeadLetterSourceQueues(*sqs.ListDeadLetterSourceQueuesInput) (*sqs.ListDeadLetterSourceQueuesOutput, error)

	ListQueues(*sqs.ListQueuesInput) (*sqs.ListQueuesOutput, error)

	PurgeQueue(*sqs.PurgeQueueInput) (*sqs.PurgeQueueOutput, error)

	ReceiveMessage(*sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error)

	RemovePermission(*sqs.RemovePermissionInput) (*sqs.RemovePermissionOutput, error)

	SendMessage(*sqs.SendMessageInput) (*sqs.SendMessageOutput, error)

	SendMessageBatch(*sqs.SendMessageBatchInput) (*sqs.SendMessageBatchOutput, error)

	SetQueueAttributes(*sqs.SetQueueAttributesInput) (*sqs.SetQueueAttributesOutput, error)
}