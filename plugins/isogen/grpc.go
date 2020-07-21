package common

import (
	"github.com/manojkva/metamorph-plugin/proto"
	"golang.org/x/net/context"
)

type GRPCClient struct{ client proto.IsogenClient }


func (m *GRPCClient)  PrepareISO() error {
	 return nil
}

type GRPCServer struct {
	Impl ISOgen
}

}
func (m *GRPCServer)  PrepareISO(ctx context.Context, req *proto.Empty) (*proto.Empty,error) {
	/* <TBD> Add check for required parameters and raise necessary errors if reqd*/
	 return nil,nil
}
