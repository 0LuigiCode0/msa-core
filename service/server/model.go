package server

import (
	"context"

	"github.com/0LuigiCode0/msa-core/grpc/msaService"
	dep_observer "github.com/0LuigiCode0/msa-core/observer/dependents"
	dep_service "github.com/0LuigiCode0/msa-core/service/dependents"

	"google.golang.org/grpc"
)

type ServiceServer interface {
	Start() error
	Close()

	Observers() dep_observer.ObserverDependentsManager
	Services() dep_service.ServiceDependentsManager
	GetAddr() string
	GetKey() string

	SetCall(f func(ctx context.Context, req *msaService.RequestCall) (*msaService.ResponseCall, error))
}

type serviceServer struct {
	msaService.ServiceServer

	server *grpc.Server
	key    string
	addr   string

	call func(ctx context.Context, req *msaService.RequestCall) (*msaService.ResponseCall, error)

	observers dep_observer.ObserverDependentsManager
	services  dep_service.ServiceDependentsManager
}
