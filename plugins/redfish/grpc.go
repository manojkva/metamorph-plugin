package redfish

import (
	"github.com/manojkva/metamorph-plugin/proto"
	"golang.org/x/net/context"
)

type GRPCClient struct{ client proto.RedfishClient }

func (m *GRPCClient) GetGUUID() ([]byte, error) {
	resp, err := m.client.GetGUUID(context.Background(), &proto.Empty{})
	if err != nil {
		return nil, err
	}

	return resp.Value, nil
}

func (m *GRPCClient) UpdateFirmware() error {
	_, err := m.client.UpdateFirmware(context.Background(), &proto.Empty{})
	return err
}
func (m *GRPCClient) ConfigureRAID() error {
	_, err := m.client.ConfigureRAID(context.Background(), &proto.Empty{})
	return err
}
func (m *GRPCClient) DeployISO() error {
	_, err := m.client.DeployISO(context.Background(), &proto.Empty{})
	return err
}

type GRPCServer struct {
	Impl Redfish
}

func (m *GRPCServer) GetGUUID(ctx context.Context, req *proto.Empty) (*proto.Response, error) {
	/* <TBD> Add check for required parameters and raise necessary errors if reqd*/
	v, err := m.Impl.GetGUUID()
	return &proto.Response{Value: v}, err
}
func (m *GRPCServer) UpdateFirmware(ctx context.Context, req *proto.Empty) (*proto.Empty, error) {
	/* <TBD> Add check for required parameters and raise necessary errors if reqd*/
	err := m.Impl.UpdateFirmware()
	return &proto.Empty{}, err
}
func (m *GRPCServer) ConfigureRAID(ctx context.Context, req *proto.Empty) (*proto.Empty, error) {
	/* <TBD> Add check for required parameters and raise necessary errors if reqd*/
	err := m.Impl.ConfigureRAID()
	return &proto.Empty{}, err
}
func (m *GRPCServer) DeployISO(ctx context.Context, req *proto.Empty) (*proto.Empty, error) {
	/* <TBD> Add check for required parameters and raise necessary errors if reqd*/
	err := m.Impl.DeployISO()
	return &proto.Empty{}, err
}
