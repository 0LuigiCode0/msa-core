package dependents

import (
	"sync"

	"github.com/0LuigiCode0/msa-core/observer/client"
)

type ObserverDependentsManager interface {
	Add(key, addr string) error
	Get(key string) (client.ObserverClient, error)
	GetFirst() (client.ObserverClient, error)
	Delete(key string) error
}

type observerDependents struct {
	dep map[string]client.ObserverClient
	rw  sync.Mutex
}
