package client

import (
	"context"

	"github.com/0LuigiCode0/msa-core/grpc/msaService"
	"github.com/0LuigiCode0/msa-core/helper"

	"google.golang.org/grpc"
)

type ServiceClient interface {
	Close()

	Call(req *msaService.RequestCall) (*msaService.ResponseCall, error)
	AddService(req *msaService.RequestAddService) (*msaService.Response, error)
	DeleteService(req *msaService.RequestDelService) (*msaService.Response, error)
}

type serviceClient struct {
	client msaService.ServiceClient
	conn   *grpc.ClientConn
	group  helper.GroupsType
	ctx    context.Context
	close  context.CancelFunc
}
