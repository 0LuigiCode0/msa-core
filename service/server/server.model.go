package server

import (
	"context"

	"github.com/0LuigiCode0/msa-core/grpc/msa_service"
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

	SetCall(f func(ctx context.Context, req *msa_service.RequestCall) (*msa_service.ResponseCall, error))
}

type serviceServer struct {
	msa_service.ServiceServer

	server *grpc.Server
	key    string
	addr   string

	call func(ctx context.Context, req *msa_service.RequestCall) (*msa_service.ResponseCall, error)

	observers dep_observer.ObserverDependentsManager
	services  dep_service.ServiceDependentsManager
}
