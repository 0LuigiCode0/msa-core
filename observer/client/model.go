package client

import (
	"context"

	"github.com/0LuigiCode0/msa-core/grpc/msaObserver"

	"google.golang.org/grpc"
)

type ObserverClient interface {
	Close()

	PushFirst(req *msaObserver.RequestPushFirst) (*msaObserver.ResponsePushFirst, error)
	PushStats(req *msaObserver.RequestPushStats) (*msaObserver.Response, error)
	PushStatus(req *msaObserver.RequestPushStatus) (*msaObserver.Response, error)
	Who(req *msaObserver.RequestWho) (*msaObserver.ResponseWho, error)
	RestartService(req *msaObserver.RequestRestartService) (*msaObserver.Response, error)
}

type observerClient struct {
	client msaObserver.ObserverClient
	conn   *grpc.ClientConn
	ctx    context.Context
	close  context.CancelFunc
}
