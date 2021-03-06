package server

import (
	"context"
	"fmt"
	"net"

	"github.com/0LuigiCode0/msa-core/grpc/msaService"
	"github.com/0LuigiCode0/msa-core/helper"
	dep_observer "github.com/0LuigiCode0/msa-core/observer/dependents"
	dep_service "github.com/0LuigiCode0/msa-core/service/dependents"

	"google.golang.org/grpc"
)

func NewServiceServer(key, addr string) ServiceServer {
	server := &serviceServer{
		server:    grpc.NewServer(),
		key:       key,
		addr:      addr,
		observers: dep_observer.NewODM(),
		services:  dep_service.NewSDM(),
	}
	msaService.RegisterServiceServer(server.server, server)

	return server
}

func (s *serviceServer) Start() error {
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
func (s *serviceServer) Close() { s.server.GracefulStop() }

func (s *serviceServer) Services() dep_service.ServiceDependentsManager    { return s.services }
func (s *serviceServer) Observers() dep_observer.ObserverDependentsManager { return s.observers }
func (s *serviceServer) GetAddr() string                                   { return s.addr }
func (s *serviceServer) GetKey() string                                    { return s.key }

func (s *serviceServer) SetCall(f func(ctx context.Context, req *msaService.RequestCall) (*msaService.ResponseCall, error)) {
	s.call = f
}

func (s *serviceServer) Call(ctx context.Context, req *msaService.RequestCall) (*msaService.ResponseCall, error) {
	return s.call(ctx, req)
}

func (s *serviceServer) AddService(ctx context.Context, req *msaService.RequestAddService) (*msaService.Response, error) {
	if err := s.services.Add(req.Key, fmt.Sprintf("%v:%v", req.Host, req.Port), helper.GroupsType(req.GroupType)); err != nil {
		return nil, fmt.Errorf("cannot add service key %v in group %v addr %v: %v", req.Key, req.GroupType, fmt.Sprintf("%v:%v", req.Host, req.Port), err)
	}
	return &msaService.Response{
		Success: true,
	}, nil
}

func (s *serviceServer) DeleteService(ctx context.Context, req *msaService.RequestDelService) (*msaService.Response, error) {
	if err := s.services.Delete(helper.GroupsType(req.GroupType), req.Key); err != nil {
		return nil, fmt.Errorf("cannot remove service key %v in group %v: %v", req.Key, req.GroupType, err)
	}
	return &msaService.Response{
		Success: true,
	}, nil
}
