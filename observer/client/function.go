package client

import (
	"fmt"

	"github.com/0LuigiCode0/msa-core/grpc/msaObserver"
	"github.com/0LuigiCode0/msa-core/helper"
)

func NewObserverClient(addr string) (ObserverClient, error) {
	var err error
	client := &observerClient{}

	client.conn, client.ctx, client.close, err = helper.CreateConn(addr)
	if err != nil {
		return nil, fmt.Errorf("cannot create conn: %v", err)
	}
	client.client = msaObserver.NewObserverClient(client.conn)

	return client, nil
}

func (c *observerClient) Close() { c.close(); c.conn.Close() }

func (c *observerClient) PushFirst(req *msaObserver.RequestPushFirst) (*msaObserver.ResponsePushFirst, error) {
	return c.client.PushFirst(helper.Ctx, req)
}
func (c *observerClient) PushStats(req *msaObserver.RequestPushStats) (*msaObserver.Response, error) {
	return c.client.PushStats(helper.Ctx, req)
}
func (c *observerClient) PushStatus(req *msaObserver.RequestPushStatus) (*msaObserver.Response, error) {
	return c.client.PushStatus(helper.Ctx, req)
}
func (c *observerClient) Who(req *msaObserver.RequestWho) (*msaObserver.ResponseWho, error) {
	return c.client.Who(helper.Ctx, req)
}
func (c *observerClient) RestartService(req *msaObserver.RequestRestartService) (*msaObserver.Response, error) {
	return c.client.RestartService(helper.Ctx, req)
}
