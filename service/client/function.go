package client

import (
	"fmt"

	"github.com/0LuigiCode0/msa-core/grpc/msaService"
	"github.com/0LuigiCode0/msa-core/helper"
)

func NewServiceClient(addr string, group helper.GroupsType) (ServiceClient, error) {
	var err error
	client := &serviceClient{
		group: group,
	}

	client.conn, client.ctx, client.close, err = helper.CreateConn(addr)
	if err != nil {
		return nil, fmt.Errorf("cannot create conn: %v", err)
	}
	client.client = msaService.NewServiceClient(client.conn)

	return client, nil
}

func (c *serviceClient) Close() { c.close(); c.conn.Close() }

func (c *serviceClient) Call(req *msaService.RequestCall) (*msaService.ResponseCall, error) {
	return c.client.Call(helper.Ctx, req)
}
func (c *serviceClient) AddService(req *msaService.RequestAddService) (*msaService.Response, error) {
	return c.client.AddService(helper.Ctx, req)
}
func (c *serviceClient) DeleteService(req *msaService.RequestDelService) (*msaService.Response, error) {
	return c.client.DeleteService(helper.Ctx, req)
}
