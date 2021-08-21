package client

import (
	"fmt"

	"github.com/0LuigiCode0/msa-core/grpc/msaMonitor"
	"github.com/0LuigiCode0/msa-core/helper"
)

func NewMonitorClient(addr string) (MonitorClient, error) {
	var err error
	client := &monitorClient{}

	client.conn, client.ctx, client.close, err = helper.CreateConn(addr)
	if err != nil {
		return nil, fmt.Errorf("cannot create conn: %v", err)
	}
	client.client = msaMonitor.NewMonitorClient(client.conn)

	return client, nil
}

func (c *monitorClient) Close() { c.close(); c.conn.Close() }

func (c *monitorClient) AddService(req *msaMonitor.RequestAddService) (*msaMonitor.Response, error) {
	return c.client.AddService(helper.Ctx, req)
}
func (c *monitorClient) DeleteService(req *msaMonitor.RequestDelService) (*msaMonitor.Response, error) {
	return c.client.DeleteService(helper.Ctx, req)
}
func (c *monitorClient) RebuildService(req *msaMonitor.RequestRebuildService) (*msaMonitor.Response, error) {
	return c.client.RebuildService(helper.Ctx, req)
}
func (c *monitorClient) SetService(req *msaMonitor.RequestSetService) (*msaMonitor.Response, error) {
	return c.client.SetService(helper.Ctx, req)
}
