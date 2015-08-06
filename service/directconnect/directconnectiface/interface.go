package directconnectiface

import (
	"github.com/datacratic/aws-sdk-go/service/directconnect"
)

type DirectConnectAPI interface {
	AllocateConnectionOnInterconnect(*directconnect.AllocateConnectionOnInterconnectInput) (*directconnect.Connection, error)

	AllocatePrivateVirtualInterface(*directconnect.AllocatePrivateVirtualInterfaceInput) (*directconnect.VirtualInterface, error)

	AllocatePublicVirtualInterface(*directconnect.AllocatePublicVirtualInterfaceInput) (*directconnect.VirtualInterface, error)

	ConfirmConnection(*directconnect.ConfirmConnectionInput) (*directconnect.ConfirmConnectionOutput, error)

	ConfirmPrivateVirtualInterface(*directconnect.ConfirmPrivateVirtualInterfaceInput) (*directconnect.ConfirmPrivateVirtualInterfaceOutput, error)

	ConfirmPublicVirtualInterface(*directconnect.ConfirmPublicVirtualInterfaceInput) (*directconnect.ConfirmPublicVirtualInterfaceOutput, error)

	CreateConnection(*directconnect.CreateConnectionInput) (*directconnect.Connection, error)

	CreateInterconnect(*directconnect.CreateInterconnectInput) (*directconnect.Interconnect, error)

	CreatePrivateVirtualInterface(*directconnect.CreatePrivateVirtualInterfaceInput) (*directconnect.VirtualInterface, error)

	CreatePublicVirtualInterface(*directconnect.CreatePublicVirtualInterfaceInput) (*directconnect.VirtualInterface, error)

	DeleteConnection(*directconnect.DeleteConnectionInput) (*directconnect.Connection, error)

	DeleteInterconnect(*directconnect.DeleteInterconnectInput) (*directconnect.DeleteInterconnectOutput, error)

	DeleteVirtualInterface(*directconnect.DeleteVirtualInterfaceInput) (*directconnect.DeleteVirtualInterfaceOutput, error)

	DescribeConnections(*directconnect.DescribeConnectionsInput) (*directconnect.Connections, error)

	DescribeConnectionsOnInterconnect(*directconnect.DescribeConnectionsOnInterconnectInput) (*directconnect.Connections, error)

	DescribeInterconnects(*directconnect.DescribeInterconnectsInput) (*directconnect.DescribeInterconnectsOutput, error)

	DescribeLocations(*directconnect.DescribeLocationsInput) (*directconnect.DescribeLocationsOutput, error)

	DescribeVirtualGateways(*directconnect.DescribeVirtualGatewaysInput) (*directconnect.DescribeVirtualGatewaysOutput, error)

	DescribeVirtualInterfaces(*directconnect.DescribeVirtualInterfacesInput) (*directconnect.DescribeVirtualInterfacesOutput, error)
}