package client

import (
	"context"

	"github.com/0LuigiCode0/msa-core/grpc/msaMonitor"

	"google.golang.org/grpc"
)

type MonitorClient interface {
	Close()

	AddService(req *msaMonitor.RequestAddService) (*msaMonitor.Response, error)
	DeleteService(req *msaMonitor.RequestDelService) (*msaMonitor.Response, error)
	RebuildService(req *msaMonitor.RequestRebuildService) (*msaMonitor.Response, error)
	SetService(req *msaMonitor.RequestSetService) (*msaMonitor.Response, error)
}

type monitorClient struct {
	client msaMonitor.MonitorClient
	conn   *grpc.ClientConn
	ctx    context.Context
	close  context.CancelFunc
}
