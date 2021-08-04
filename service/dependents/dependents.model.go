package dependents

import (
	"sync"

	"github.com/0LuigiCode0/msa-core/helper"
	"github.com/0LuigiCode0/msa-core/service/client"
)

type ServiceDependentsManager interface {
	Add(key, addr string, group helper.GroupsType) error
	Get(group helper.GroupsType, key string) (client.ServiceClient, error)
	GetFirstByGroup(group helper.GroupsType) (client.ServiceClient, error)
	Delete(group helper.GroupsType, key string) error
}

type serviceDependents struct {
	dep map[helper.GroupsType]map[string]client.ServiceClient
	rw  sync.Mutex
}
