package server

import (
	"context"

	"github.com/0LuigiCode0/msa-core/grpc/msaObserver"
	"github.com/0LuigiCode0/msa-core/monitor/dependents"

	"google.golang.org/grpc"
)

type ObserverServer interface {
	Start() error
	Close()

	Monitors() dependents.MonitorDependentsManager
	GetAddr() string
	GetKey() string

	SetPushFirst(f func(ctx context.Context, req *msaObserver.RequestPushFirst) (*msaObserver.ResponsePushFirst, error))
	SetPushStats(f func(ctx context.Context, req *msaObserver.RequestPushStats) (*msaObserver.Response, error))
	SetPushStatus(f func(ctx context.Context, req *msaObserver.RequestPushStatus) (*msaObserver.Response, error))
	SetWho(f func(ctx context.Context, req *msaObserver.RequestWho) (*msaObserver.ResponseWho, error))
	SetRestartService(f func(ctx context.Context, req *msaObserver.RequestRestartService) (*msaObserver.Response, error))
}

type observerServer struct {
	msaObserver.ObserverServer

	server *grpc.Server
	key    string
	addr   string

	pushFirst      func(context.Context, *msaObserver.RequestPushFirst) (*msaObserver.ResponsePushFirst, error)
	pushStats      func(context.Context, *msaObserver.RequestPushStats) (*msaObserver.Response, error)
	pushStatus     func(context.Context, *msaObserver.RequestPushStatus) (*msaObserver.Response, error)
	who            func(context.Context, *msaObserver.RequestWho) (*msaObserver.ResponseWho, error)
	restartService func(context.Context, *msaObserver.RequestRestartService) (*msaObserver.Response, error)

	monitors dependents.MonitorDependentsManager
}
