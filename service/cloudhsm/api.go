// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudhsm provides a client for Amazon CloudHSM.
package cloudhsm

import (
	"sync"

	"github.com/datacratic/aws-sdk-go/aws"
)

var oprw sync.Mutex

// CreateHAPGRequest generates a request for the CreateHAPG operation.
func (c *CloudHSM) CreateHAPGRequest(input *CreateHAPGInput) (req *aws.Request, output *CreateHAPGOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opCreateHAPG == nil {
		opCreateHAPG = &aws.Operation{
			Name:       "CreateHapg",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &CreateHAPGInput{}
	}

	req = c.newRequest(opCreateHAPG, input, output)
	output = &CreateHAPGOutput{}
	req.Data = output
	return
}

// Creates a high-availability partition group. A high-availability partition
// group is a group of partitions that spans multiple physical HSMs.
func (c *CloudHSM) CreateHAPG(input *CreateHAPGInput) (output *CreateHAPGOutput, err error) {
	req, out := c.CreateHAPGRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreateHAPG *aws.Operation

// CreateHSMRequest generates a request for the CreateHSM operation.
func (c *CloudHSM) CreateHSMRequest(input *CreateHSMInput) (req *aws.Request, output *CreateHSMOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opCreateHSM == nil {
		opCreateHSM = &aws.Operation{
			Name:       "CreateHsm",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &CreateHSMInput{}
	}

	req = c.newRequest(opCreateHSM, input, output)
	output = &CreateHSMOutput{}
	req.Data = output
	return
}

// Creates an uninitialized HSM instance. Running this command provisions an
// HSM appliance and will result in charges to your AWS account for the HSM.
func (c *CloudHSM) CreateHSM(input *CreateHSMInput) (output *CreateHSMOutput, err error) {
	req, out := c.CreateHSMRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreateHSM *aws.Operation

// CreateLunaClientRequest generates a request for the CreateLunaClient operation.
func (c *CloudHSM) CreateLunaClientRequest(input *CreateLunaClientInput) (req *aws.Request, output *CreateLunaClientOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opCreateLunaClient == nil {
		opCreateLunaClient = &aws.Operation{
			Name:       "CreateLunaClient",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &CreateLunaClientInput{}
	}

	req = c.newRequest(opCreateLunaClient, input, output)
	output = &CreateLunaClientOutput{}
	req.Data = output
	return
}

// Creates an HSM client.
func (c *CloudHSM) CreateLunaClient(input *CreateLunaClientInput) (output *CreateLunaClientOutput, err error) {
	req, out := c.CreateLunaClientRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreateLunaClient *aws.Operation

// DeleteHAPGRequest generates a request for the DeleteHAPG operation.
func (c *CloudHSM) DeleteHAPGRequest(input *DeleteHAPGInput) (req *aws.Request, output *DeleteHAPGOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDeleteHAPG == nil {
		opDeleteHAPG = &aws.Operation{
			Name:       "DeleteHapg",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &DeleteHAPGInput{}
	}

	req = c.newRequest(opDeleteHAPG, input, output)
	output = &DeleteHAPGOutput{}
	req.Data = output
	return
}

// Deletes a high-availability partition group.
func (c *CloudHSM) DeleteHAPG(input *DeleteHAPGInput) (output *DeleteHAPGOutput, err error) {
	req, out := c.DeleteHAPGRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteHAPG *aws.Operation

// DeleteHSMRequest generates a request for the DeleteHSM operation.
func (c *CloudHSM) DeleteHSMRequest(input *DeleteHSMInput) (req *aws.Request, output *DeleteHSMOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDeleteHSM == nil {
		opDeleteHSM = &aws.Operation{
			Name:       "DeleteHsm",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &DeleteHSMInput{}
	}

	req = c.newRequest(opDeleteHSM, input, output)
	output = &DeleteHSMOutput{}
	req.Data = output
	return
}

// Deletes an HSM. Once complete, this operation cannot be undone and your key
// material cannot be recovered.
func (c *CloudHSM) DeleteHSM(input *DeleteHSMInput) (output *DeleteHSMOutput, err error) {
	req, out := c.DeleteHSMRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteHSM *aws.Operation

// DeleteLunaClientRequest generates a request for the DeleteLunaClient operation.
func (c *CloudHSM) DeleteLunaClientRequest(input *DeleteLunaClientInput) (req *aws.Request, output *DeleteLunaClientOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDeleteLunaClient == nil {
		opDeleteLunaClient = &aws.Operation{
			Name:       "DeleteLunaClient",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &DeleteLunaClientInput{}
	}

	req = c.newRequest(opDeleteLunaClient, input, output)
	output = &DeleteLunaClientOutput{}
	req.Data = output
	return
}

// Deletes a client.
func (c *CloudHSM) DeleteLunaClient(input *DeleteLunaClientInput) (output *DeleteLunaClientOutput, err error) {
	req, out := c.DeleteLunaClientRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteLunaClient *aws.Operation

// DescribeHAPGRequest generates a request for the DescribeHAPG operation.
func (c *CloudHSM) DescribeHAPGRequest(input *DescribeHAPGInput) (req *aws.Request, output *DescribeHAPGOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDescribeHAPG == nil {
		opDescribeHAPG = &aws.Operation{
			Name:       "DescribeHapg",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &DescribeHAPGInput{}
	}

	req = c.newRequest(opDescribeHAPG, input, output)
	output = &DescribeHAPGOutput{}
	req.Data = output
	return
}

// Retrieves information about a high-availability partition group.
func (c *CloudHSM) DescribeHAPG(input *DescribeHAPGInput) (output *DescribeHAPGOutput, err error) {
	req, out := c.DescribeHAPGRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeHAPG *aws.Operation

// DescribeHSMRequest generates a request for the DescribeHSM operation.
func (c *CloudHSM) DescribeHSMRequest(input *DescribeHSMInput) (req *aws.Request, output *DescribeHSMOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDescribeHSM == nil {
		opDescribeHSM = &aws.Operation{
			Name:       "DescribeHsm",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &DescribeHSMInput{}
	}

	req = c.newRequest(opDescribeHSM, input, output)
	output = &DescribeHSMOutput{}
	req.Data = output
	return
}

// Retrieves information about an HSM. You can identify the HSM by its ARN or
// its serial number.
func (c *CloudHSM) DescribeHSM(input *DescribeHSMInput) (output *DescribeHSMOutput, err error) {
	req, out := c.DescribeHSMRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeHSM *aws.Operation

// DescribeLunaClientRequest generates a request for the DescribeLunaClient operation.
func (c *CloudHSM) DescribeLunaClientRequest(input *DescribeLunaClientInput) (req *aws.Request, output *DescribeLunaClientOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDescribeLunaClient == nil {
		opDescribeLunaClient = &aws.Operation{
			Name:       "DescribeLunaClient",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &DescribeLunaClientInput{}
	}

	req = c.newRequest(opDescribeLunaClient, input, output)
	output = &DescribeLunaClientOutput{}
	req.Data = output
	return
}

// Retrieves information about an HSM client.
func (c *CloudHSM) DescribeLunaClient(input *DescribeLunaClientInput) (output *DescribeLunaClientOutput, err error) {
	req, out := c.DescribeLunaClientRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeLunaClient *aws.Operation

// GetConfigRequest generates a request for the GetConfig operation.
func (c *CloudHSM) GetConfigRequest(input *GetConfigInput) (req *aws.Request, output *GetConfigOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opGetConfig == nil {
		opGetConfig = &aws.Operation{
			Name:       "GetConfig",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &GetConfigInput{}
	}

	req = c.newRequest(opGetConfig, input, output)
	output = &GetConfigOutput{}
	req.Data = output
	return
}

// Gets the configuration files necessary to connect to all high availability
// partition groups the client is associated with.
func (c *CloudHSM) GetConfig(input *GetConfigInput) (output *GetConfigOutput, err error) {
	req, out := c.GetConfigRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetConfig *aws.Operation

// ListAvailableZonesRequest generates a request for the ListAvailableZones operation.
func (c *CloudHSM) ListAvailableZonesRequest(input *ListAvailableZonesInput) (req *aws.Request, output *ListAvailableZonesOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opListAvailableZones == nil {
		opListAvailableZones = &aws.Operation{
			Name:       "ListAvailableZones",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &ListAvailableZonesInput{}
	}

	req = c.newRequest(opListAvailableZones, input, output)
	output = &ListAvailableZonesOutput{}
	req.Data = output
	return
}

// Lists the Availability Zones that have available AWS CloudHSM capacity.
func (c *CloudHSM) ListAvailableZones(input *ListAvailableZonesInput) (output *ListAvailableZonesOutput, err error) {
	req, out := c.ListAvailableZonesRequest(input)
	output = out
	err = req.Send()
	return
}

var opListAvailableZones *aws.Operation

// ListHSMsRequest generates a request for the ListHSMs operation.
func (c *CloudHSM) ListHSMsRequest(input *ListHSMsInput) (req *aws.Request, output *ListHSMsOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opListHSMs == nil {
		opListHSMs = &aws.Operation{
			Name:       "ListHsms",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &ListHSMsInput{}
	}

	req = c.newRequest(opListHSMs, input, output)
	output = &ListHSMsOutput{}
	req.Data = output
	return
}

// Retrieves the identifiers of all of the HSMs provisioned for the current
// customer.
//
// This operation supports pagination with the use of the NextToken member.
// If more results are available, the NextToken member of the response contains
// a token that you pass in the next call to ListHsms to retrieve the next set
// of items.
func (c *CloudHSM) ListHSMs(input *ListHSMsInput) (output *ListHSMsOutput, err error) {
	req, out := c.ListHSMsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListHSMs *aws.Operation

// ListHapgsRequest generates a request for the ListHapgs operation.
func (c *CloudHSM) ListHapgsRequest(input *ListHapgsInput) (req *aws.Request, output *ListHapgsOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opListHapgs == nil {
		opListHapgs = &aws.Operation{
			Name:       "ListHapgs",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &ListHapgsInput{}
	}

	req = c.newRequest(opListHapgs, input, output)
	output = &ListHapgsOutput{}
	req.Data = output
	return
}

// Lists the high-availability partition groups for the account.
//
// This operation supports pagination with the use of the NextToken member.
// If more results are available, the NextToken member of the response contains
// a token that you pass in the next call to ListHapgs to retrieve the next
// set of items.
func (c *CloudHSM) ListHapgs(input *ListHapgsInput) (output *ListHapgsOutput, err error) {
	req, out := c.ListHapgsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListHapgs *aws.Operation

// ListLunaClientsRequest generates a request for the ListLunaClients operation.
func (c *CloudHSM) ListLunaClientsRequest(input *ListLunaClientsInput) (req *aws.Request, output *ListLunaClientsOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opListLunaClients == nil {
		opListLunaClients = &aws.Operation{
			Name:       "ListLunaClients",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &ListLunaClientsInput{}
	}

	req = c.newRequest(opListLunaClients, input, output)
	output = &ListLunaClientsOutput{}
	req.Data = output
	return
}

// Lists all of the clients.
//
// This operation supports pagination with the use of the NextToken member.
// If more results are available, the NextToken member of the response contains
// a token that you pass in the next call to ListLunaClients to retrieve the
// next set of items.
func (c *CloudHSM) ListLunaClients(input *ListLunaClientsInput) (output *ListLunaClientsOutput, err error) {
	req, out := c.ListLunaClientsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListLunaClients *aws.Operation

// ModifyHAPGRequest generates a request for the ModifyHAPG operation.
func (c *CloudHSM) ModifyHAPGRequest(input *ModifyHAPGInput) (req *aws.Request, output *ModifyHAPGOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opModifyHAPG == nil {
		opModifyHAPG = &aws.Operation{
			Name:       "ModifyHapg",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &ModifyHAPGInput{}
	}

	req = c.newRequest(opModifyHAPG, input, output)
	output = &ModifyHAPGOutput{}
	req.Data = output
	return
}

// Modifies an existing high-availability partition group.
func (c *CloudHSM) ModifyHAPG(input *ModifyHAPGInput) (output *ModifyHAPGOutput, err error) {
	req, out := c.ModifyHAPGRequest(input)
	output = out
	err = req.Send()
	return
}

var opModifyHAPG *aws.Operation

// ModifyHSMRequest generates a request for the ModifyHSM operation.
func (c *CloudHSM) ModifyHSMRequest(input *ModifyHSMInput) (req *aws.Request, output *ModifyHSMOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opModifyHSM == nil {
		opModifyHSM = &aws.Operation{
			Name:       "ModifyHsm",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &ModifyHSMInput{}
	}

	req = c.newRequest(opModifyHSM, input, output)
	output = &ModifyHSMOutput{}
	req.Data = output
	return
}

// Modifies an HSM.
func (c *CloudHSM) ModifyHSM(input *ModifyHSMInput) (output *ModifyHSMOutput, err error) {
	req, out := c.ModifyHSMRequest(input)
	output = out
	err = req.Send()
	return
}

var opModifyHSM *aws.Operation

// ModifyLunaClientRequest generates a request for the ModifyLunaClient operation.
func (c *CloudHSM) ModifyLunaClientRequest(input *ModifyLunaClientInput) (req *aws.Request, output *ModifyLunaClientOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opModifyLunaClient == nil {
		opModifyLunaClient = &aws.Operation{
			Name:       "ModifyLunaClient",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &ModifyLunaClientInput{}
	}

	req = c.newRequest(opModifyLunaClient, input, output)
	output = &ModifyLunaClientOutput{}
	req.Data = output
	return
}

// Modifies the certificate used by the client.
//
// This action can potentially start a workflow to install the new certificate
// on the client's HSMs.
func (c *CloudHSM) ModifyLunaClient(input *ModifyLunaClientInput) (output *ModifyLunaClientOutput, err error) {
	req, out := c.ModifyLunaClientRequest(input)
	output = out
	err = req.Send()
	return
}

var opModifyLunaClient *aws.Operation

// Contains the inputs for the CreateHapgRequest action.
type CreateHAPGInput struct {
	// The label of the new high-availability partition group.
	Label *string `type:"string" required:"true"`

	metadataCreateHAPGInput `json:"-", xml:"-"`
}

type metadataCreateHAPGInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Contains the output of the CreateHAPartitionGroup action.
type CreateHAPGOutput struct {
	// The ARN of the high-availability partition group.
	HAPGARN *string `locationName:"HapgArn" type:"string"`

	metadataCreateHAPGOutput `json:"-", xml:"-"`
}

type metadataCreateHAPGOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Contains the inputs for the CreateHsm action.
type CreateHSMInput struct {
	// A user-defined token to ensure idempotence. Subsequent calls to this action
	// with the same token will be ignored.
	ClientToken *string `locationName:"ClientToken" type:"string"`

	// The IP address to assign to the HSM's ENI.
	ENIIP *string `locationName:"EniIp" type:"string"`

	// The external ID from IamRoleArn, if present.
	ExternalID *string `locationName:"ExternalId" type:"string"`

	// The ARN of an IAM role to enable the AWS CloudHSM service to allocate an
	// ENI on your behalf.
	IAMRoleARN *string `locationName:"IamRoleArn" type:"string" required:"true"`

	// The SSH public key to install on the HSM.
	SSHKey *string `locationName:"SshKey" type:"string" required:"true"`

	// The identifier of the subnet in your VPC in which to place the HSM.
	SubnetID *string `locationName:"SubnetId" type:"string" required:"true"`

	// The subscription type.
	SubscriptionType *string `locationName:"SubscriptionType" type:"string" required:"true"`

	// The IP address for the syslog monitoring server.
	SyslogIP *string `locationName:"SyslogIp" type:"string"`

	metadataCreateHSMInput `json:"-", xml:"-"`
}

type metadataCreateHSMInput struct {
	SDKShapeTraits bool `locationName:"CreateHsmRequest" type:"structure"`
}

// Contains the output of the CreateHsm action.
type CreateHSMOutput struct {
	// The ARN of the HSM.
	HSMARN *string `locationName:"HsmArn" type:"string"`

	metadataCreateHSMOutput `json:"-", xml:"-"`
}

type metadataCreateHSMOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Contains the inputs for the CreateLunaClient action.
type CreateLunaClientInput struct {
	// The contents of a Base64-Encoded X.509 v3 certificate to be installed on
	// the HSMs used by this client.
	Certificate *string `type:"string" required:"true"`

	// The label for the client.
	Label *string `type:"string"`

	metadataCreateLunaClientInput `json:"-", xml:"-"`
}

type metadataCreateLunaClientInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Contains the output of the CreateLunaClient action.
type CreateLunaClientOutput struct {
	// The ARN of the client.
	ClientARN *string `locationName:"ClientArn" type:"string"`

	metadataCreateLunaClientOutput `json:"-", xml:"-"`
}

type metadataCreateLunaClientOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Contains the inputs for the DeleteHapg action.
type DeleteHAPGInput struct {
	// The ARN of the high-availability partition group to delete.
	HAPGARN *string `locationName:"HapgArn" type:"string" required:"true"`

	metadataDeleteHAPGInput `json:"-", xml:"-"`
}

type metadataDeleteHAPGInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Contains the output of the DeleteHapg action.
type DeleteHAPGOutput struct {
	// The status of the action.
	Status *string `type:"string" required:"true"`

	metadataDeleteHAPGOutput `json:"-", xml:"-"`
}

type metadataDeleteHAPGOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Contains the inputs for the DeleteHsm action.
type DeleteHSMInput struct {
	// The ARN of the HSM to delete.
	HSMARN *string `locationName:"HsmArn" type:"string" required:"true"`

	metadataDeleteHSMInput `json:"-", xml:"-"`
}

type metadataDeleteHSMInput struct {
	SDKShapeTraits bool `locationName:"DeleteHsmRequest" type:"structure"`
}

// Contains the output of the DeleteHsm action.
type DeleteHSMOutput struct {
	// The status of the action.
	Status *string `type:"string" required:"true"`

	metadataDeleteHSMOutput `json:"-", xml:"-"`
}

type metadataDeleteHSMOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteLunaClientInput struct {
	// The ARN of the client to delete.
	ClientARN *string `locationName:"ClientArn" type:"string" required:"true"`

	metadataDeleteLunaClientInput `json:"-", xml:"-"`
}

type metadataDeleteLunaClientInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteLunaClientOutput struct {
	// The status of the action.
	Status *string `type:"string" required:"true"`

	metadataDeleteLunaClientOutput `json:"-", xml:"-"`
}

type metadataDeleteLunaClientOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Contains the inputs for the DescribeHapg action.
type DescribeHAPGInput struct {
	// The ARN of the high-availability partition group to describe.
	HAPGARN *string `locationName:"HapgArn" type:"string" required:"true"`

	metadataDescribeHAPGInput `json:"-", xml:"-"`
}

type metadataDescribeHAPGInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Contains the output of the DescribeHapg action.
type DescribeHAPGOutput struct {
	// The ARN of the high-availability partition group.
	HAPGARN *string `locationName:"HapgArn" type:"string"`

	// The serial number of the high-availability partition group.
	HAPGSerial *string `locationName:"HapgSerial" type:"string"`

	// Contains a list of ARNs that identify the HSMs.
	HSMsLastActionFailed []*string `locationName:"HsmsLastActionFailed" type:"list"`

	// Contains a list of ARNs that identify the HSMs.
	HSMsPendingDeletion []*string `locationName:"HsmsPendingDeletion" type:"list"`

	// Contains a list of ARNs that identify the HSMs.
	HSMsPendingRegistration []*string `locationName:"HsmsPendingRegistration" type:"list"`

	// The label for the high-availability partition group.
	Label *string `type:"string"`

	// The date and time the high-availability partition group was last modified.
	LastModifiedTimestamp *string `type:"string"`

	// The list of partition serial numbers that belong to the high-availability
	// partition group.
	PartitionSerialList []*string `type:"list"`

	// The state of the high-availability partition group.
	State *string `type:"string"`

	metadataDescribeHAPGOutput `json:"-", xml:"-"`
}

type metadataDescribeHAPGOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Contains the inputs for the DescribeHsm action.
type DescribeHSMInput struct {
	// The ARN of the HSM. Either the HsmArn or the SerialNumber parameter must
	// be specified.
	HSMARN *string `locationName:"HsmArn" type:"string"`

	// The serial number of the HSM. Either the HsmArn or the HsmSerialNumber parameter
	// must be specified.
	HSMSerialNumber *string `locationName:"HsmSerialNumber" type:"string"`

	metadataDescribeHSMInput `json:"-", xml:"-"`
}

type metadataDescribeHSMInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Contains the output of the DescribeHsm action.
type DescribeHSMOutput struct {
	// The Availability Zone that the HSM is in.
	AvailabilityZone *string `type:"string"`

	// The identifier of the elastic network interface (ENI) attached to the HSM.
	ENIID *string `locationName:"EniId" type:"string"`

	// The IP address assigned to the HSM's ENI.
	ENIIP *string `locationName:"EniIp" type:"string"`

	// The ARN of the HSM.
	HSMARN *string `locationName:"HsmArn" type:"string"`

	// The HSM model type.
	HSMType *string `locationName:"HsmType" type:"string"`

	// The ARN of the IAM role assigned to the HSM.
	IAMRoleARN *string `locationName:"IamRoleArn" type:"string"`

	// The list of partitions on the HSM.
	Partitions []*string `type:"list"`

	// The date and time the SSH key was last updated.
	SSHKeyLastUpdated *string `locationName:"SshKeyLastUpdated" type:"string"`

	// The public SSH key.
	SSHPublicKey *string `locationName:"SshPublicKey" type:"string"`

	// The serial number of the HSM.
	SerialNumber *string `type:"string"`

	// The date and time the server certificate was last updated.
	ServerCertLastUpdated *string `type:"string"`

	// The URI of the certificate server.
	ServerCertURI *string `locationName:"ServerCertUri" type:"string"`

	// The HSM software version.
	SoftwareVersion *string `type:"string"`

	// The status of the HSM.
	Status *string `type:"string"`

	// Contains additional information about the status of the HSM.
	StatusDetails *string `type:"string"`

	// The identifier of the subnet the HSM is in.
	SubnetID *string `locationName:"SubnetId" type:"string"`

	// The subscription end date.
	SubscriptionEndDate *string `type:"string"`

	// The subscription start date.
	SubscriptionStartDate *string `type:"string"`

	// The subscription type.
	SubscriptionType *string `type:"string"`

	// The identifier of the VPC that the HSM is in.
	VPCID *string `locationName:"VpcId" type:"string"`

	// The name of the HSM vendor.
	VendorName *string `type:"string"`

	metadataDescribeHSMOutput `json:"-", xml:"-"`
}

type metadataDescribeHSMOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeLunaClientInput struct {
	// The certificate fingerprint.
	CertificateFingerprint *string `type:"string"`

	// The ARN of the client.
	ClientARN *string `locationName:"ClientArn" type:"string"`

	metadataDescribeLunaClientInput `json:"-", xml:"-"`
}

type metadataDescribeLunaClientInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeLunaClientOutput struct {
	// The certificate installed on the HSMs used by this client.
	Certificate *string `type:"string"`

	// The certificate fingerprint.
	CertificateFingerprint *string `type:"string"`

	// The ARN of the client.
	ClientARN *string `locationName:"ClientArn" type:"string"`

	// The label of the client.
	Label *string `type:"string"`

	// The date and time the client was last modified.
	LastModifiedTimestamp *string `type:"string"`

	metadataDescribeLunaClientOutput `json:"-", xml:"-"`
}

type metadataDescribeLunaClientOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetConfigInput struct {
	// The ARN of the client.
	ClientARN *string `locationName:"ClientArn" type:"string" required:"true"`

	// The client version.
	ClientVersion *string `type:"string" required:"true"`

	// A list of ARNs that identify the high-availability partition groups that
	// are associated with the client.
	HAPGList []*string `locationName:"HapgList" type:"list" required:"true"`

	metadataGetConfigInput `json:"-", xml:"-"`
}

type metadataGetConfigInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetConfigOutput struct {
	// The certificate file containing the server.pem files of the HSMs.
	ConfigCred *string `type:"string"`

	// The chrystoki.conf configuration file.
	ConfigFile *string `type:"string"`

	// The type of credentials.
	ConfigType *string `type:"string"`

	metadataGetConfigOutput `json:"-", xml:"-"`
}

type metadataGetConfigOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Contains the inputs for the ListAvailableZones action.
type ListAvailableZonesInput struct {
	metadataListAvailableZonesInput `json:"-", xml:"-"`
}

type metadataListAvailableZonesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListAvailableZonesOutput struct {
	// The list of Availability Zones that have available AWS CloudHSM capacity.
	AZList []*string `type:"list"`

	metadataListAvailableZonesOutput `json:"-", xml:"-"`
}

type metadataListAvailableZonesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListHSMsInput struct {
	// The NextToken value from a previous call to ListHsms. Pass null if this is
	// the first call.
	NextToken *string `type:"string"`

	metadataListHSMsInput `json:"-", xml:"-"`
}

type metadataListHSMsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Contains the output of the ListHsms action.
type ListHSMsOutput struct {
	// The list of ARNs that identify the HSMs.
	HSMList []*string `locationName:"HsmList" type:"list"`

	// If not null, more results are available. Pass this value to ListHsms to retrieve
	// the next set of items.
	NextToken *string `type:"string"`

	metadataListHSMsOutput `json:"-", xml:"-"`
}

type metadataListHSMsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListHapgsInput struct {
	// The NextToken value from a previous call to ListHapgs. Pass null if this
	// is the first call.
	NextToken *string `type:"string"`

	metadataListHapgsInput `json:"-", xml:"-"`
}

type metadataListHapgsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListHapgsOutput struct {
	// The list of high-availability partition groups.
	HAPGList []*string `locationName:"HapgList" type:"list" required:"true"`

	// If not null, more results are available. Pass this value to ListHapgs to
	// retrieve the next set of items.
	NextToken *string `type:"string"`

	metadataListHapgsOutput `json:"-", xml:"-"`
}

type metadataListHapgsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListLunaClientsInput struct {
	// The NextToken value from a previous call to ListLunaClients. Pass null if
	// this is the first call.
	NextToken *string `type:"string"`

	metadataListLunaClientsInput `json:"-", xml:"-"`
}

type metadataListLunaClientsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListLunaClientsOutput struct {
	// The list of clients.
	ClientList []*string `type:"list" required:"true"`

	// If not null, more results are available. Pass this to ListLunaClients to
	// retrieve the next set of items.
	NextToken *string `type:"string"`

	metadataListLunaClientsOutput `json:"-", xml:"-"`
}

type metadataListLunaClientsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ModifyHAPGInput struct {
	// The ARN of the high-availability partition group to modify.
	HAPGARN *string `locationName:"HapgArn" type:"string" required:"true"`

	// The new label for the high-availability partition group.
	Label *string `type:"string"`

	// The list of partition serial numbers to make members of the high-availability
	// partition group.
	PartitionSerialList []*string `type:"list"`

	metadataModifyHAPGInput `json:"-", xml:"-"`
}

type metadataModifyHAPGInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ModifyHAPGOutput struct {
	// The ARN of the high-availability partition group.
	HAPGARN *string `locationName:"HapgArn" type:"string"`

	metadataModifyHAPGOutput `json:"-", xml:"-"`
}

type metadataModifyHAPGOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Contains the inputs for the ModifyHsm action.
type ModifyHSMInput struct {
	// The new IP address for the elastic network interface attached to the HSM.
	ENIIP *string `locationName:"EniIp" type:"string"`

	// The new external ID.
	ExternalID *string `locationName:"ExternalId" type:"string"`

	// The ARN of the HSM to modify.
	HSMARN *string `locationName:"HsmArn" type:"string" required:"true"`

	// The new IAM role ARN.
	IAMRoleARN *string `locationName:"IamRoleArn" type:"string"`

	// The new identifier of the subnet that the HSM is in.
	SubnetID *string `locationName:"SubnetId" type:"string"`

	// The new IP address for the syslog monitoring server.
	SyslogIP *string `locationName:"SyslogIp" type:"string"`

	metadataModifyHSMInput `json:"-", xml:"-"`
}

type metadataModifyHSMInput struct {
	SDKShapeTraits bool `locationName:"ModifyHsmRequest" type:"structure"`
}

// Contains the output of the ModifyHsm action.
type ModifyHSMOutput struct {
	// The ARN of the HSM.
	HSMARN *string `locationName:"HsmArn" type:"string"`

	metadataModifyHSMOutput `json:"-", xml:"-"`
}

type metadataModifyHSMOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ModifyLunaClientInput struct {
	// The new certificate for the client.
	Certificate *string `type:"string" required:"true"`

	// The ARN of the client.
	ClientARN *string `locationName:"ClientArn" type:"string" required:"true"`

	metadataModifyLunaClientInput `json:"-", xml:"-"`
}

type metadataModifyLunaClientInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ModifyLunaClientOutput struct {
	// The ARN of the client.
	ClientARN *string `locationName:"ClientArn" type:"string"`

	metadataModifyLunaClientOutput `json:"-", xml:"-"`
}

type metadataModifyLunaClientOutput struct {
	SDKShapeTraits bool `type:"structure"`
}