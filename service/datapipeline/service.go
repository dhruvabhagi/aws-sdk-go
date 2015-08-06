package datapipeline

import (
	"github.com/datacratic/aws-sdk-go/aws"
	"github.com/datacratic/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/datacratic/aws-sdk-go/internal/signer/v4"
)

// DataPipeline is a client for AWS Data Pipeline.
type DataPipeline struct {
	*aws.Service
}

// Used for custom service initialization logic
var initService func(*aws.Service)

// Used for custom request initialization logic
var initRequest func(*aws.Request)

// New returns a new DataPipeline client.
func New(config *aws.Config) *DataPipeline {
	if config == nil {
		config = &aws.Config{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config),
		ServiceName:  "datapipeline",
		APIVersion:   "2012-10-29",
		JSONVersion:  "1.1",
		TargetPrefix: "DataPipeline",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &DataPipeline{service}
}

// newRequest creates a new request for a DataPipeline operation and runs any
// custom request initialization.
func (c *DataPipeline) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := aws.NewRequest(c.Service, op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
