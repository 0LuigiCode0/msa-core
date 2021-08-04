package server

import (
	"msa-core/grpc/msa_monitor"
	dep_observer "msa-core/observer/dependents"
	dep_service "msa-core/service/dependents"

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
