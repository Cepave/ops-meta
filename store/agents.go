package store

import (
	"gitcafe.com/ops/common/model"
	"sync"
)

type AgentsMap struct {
	sync.RWMutex
	M map[string]*model.RealAgent
}

func NewAgentsMap() *AgentsMap {
	return &AgentsMap{M: make(map[string]*model.RealAgent)}
}

type HostAgentsMap struct {
	sync.RWMutex
	M map[string]*AgentsMap
}

func NewHostAgentsMap() *HostAgentsMap {
	return &HostAgentsMap{M: make(map[string]*AgentsMap)}
}

var HostAgents = NewHostAgentsMap()

func (this *HostAgentsMap) Get(hostname string) (*AgentsMap, bool) {
	this.RLock()
	defer this.RUnlock()
	val, exists := this.M[hostname]
	return val, exists
}

func HandleHeartbeatRequest(req *model.HeartbeatRequest) {

}
