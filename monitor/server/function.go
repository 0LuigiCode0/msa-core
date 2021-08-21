package server

import (
	"context"
	"fmt"
	"net"

	"github.com/0LuigiCode0/msa-core/grpc/msaMonitor"
	dep_observer "github.com/0LuigiCode0/msa-core/observer/dependents"
	dep_service "github.com/0LuigiCode0/msa-core/service/dependents"

	"google.golang.org/grpc"
)

func NewMonitorServer(key, addr string) MonitorServer {
	server := &monitorServer{
		server:    grpc.NewServer(),
		key:       key,
		addr:      addr,
		observers: dep_observer.NewODM(),
		services:  dep_service.NewSDM(),
	}
	msaMonitor.RegisterMonitorServer(server.server, server)

	return server
}

func (s *monitorServer) Start() error {
	gListen, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("glistener error: %v", err)
	}
	if err := s.server.Serve(gListen); err != nil {
		if err == grpc.ErrServerStopped {
			return nil
		}
		return fmt.Errorf("gserve error: %v", err)
	}

	return nil
}
func (s *monitorServer) Close() { s.server.Stop() }

func (s *monitorServer) Observers() dep_observer.ObserverDependentsManager { return s.observers }
func (s *monitorServer) Services() dep_service.ServiceDependentsManager    { return s.services }
func (s *monitorServer) GetAddr() string                                   { return s.addr }
func (s *monitorServer) GetKey() string                                    { return s.key }

func (s *monitorServer) AddService(ctx context.Context, req *msaMonitor.RequestAddService) (*msaMonitor.Response, error) {
	fmt.Println(req)
	return &msaMonitor.Response{
		Success: true,
	}, nil
}

func (s *monitorServer) DeleteService(ctx context.Context, req *msaMonitor.RequestDelService) (*msaMonitor.Response, error) {
	fmt.Println(req)
	return &msaMonitor.Response{
		Success: true,
	}, nil
}

func (s *monitorServer) RebuildService(ctx context.Context, req *msaMonitor.RequestRebuildService) (*msaMonitor.Response, error) {
	fmt.Println(req)
	return &msaMonitor.Response{
		Success: true,
	}, nil
}

func (s *monitorServer) SetService(ctx context.Context, req *msaMonitor.RequestSetService) (*msaMonitor.Response, error) {
	fmt.Println(req)
	return &msaMonitor.Response{
		Success: true,
	}, nil
}
