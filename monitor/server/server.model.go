package server

import (
	"github.com/0LuigiCode0/msa-core/grpc/msa_monitor"
	dep_observer "github.com/0LuigiCode0/msa-core/observer/dependents"
	dep_service "github.com/0LuigiCode0/msa-core/service/dependents"

	"google.golang.org/grpc"
)

type MonitorServer interface {
	Start() error
	Close()

	Observers() dep_observer.ObserverDependentsManager
	Services() dep_service.ServiceDependentsManager
	GetAddr() string
	GetKey() string
}

type monitorServer struct {
	msa_monitor.MonitorServer

	server *grpc.Server
	key    string
	addr   string

	observers dep_observer.ObserverDependentsManager
	services  dep_service.ServiceDependentsManager
}
