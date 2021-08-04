package dependents

import (
	"msa-core/observer/client"
	"sync"
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
