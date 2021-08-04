package dependents

import (
	"sync"

	"github.com/0LuigiCode0/msa-core/monitor/client"
)

type MonitorDependentsManager interface {
	Add(key, addr string) error
	Get(key string) (client.MonitorClient, error)
	GetFirst() (client.MonitorClient, error)
	Delete(key string) error
}

type monitorDependents struct {
	dep map[string]client.MonitorClient
	rw  sync.Mutex
}
