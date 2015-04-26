package store

import (
	"time"
)

func CleanStaleHost() {
	d := time.Duration(24) * time.Hour
	for {
		time.Sleep(d)
		cleanStaleHost()
	}
}

func cleanStaleHost() {
	// three days ago
	before := time.Now().Unix() - 3600*24*3

	hostnames := HostAgents.Hostnames()
	count := len(hostnames)
	if count == 0 {
		return
	}

	for i := 0; i < count; i++ {
		agentsMap, exists := HostAgents.Get(hostnames[i])
		if !exists {
			continue
		}

		if agentsMap == nil || agentsMap.Len() == 0 {
			HostAgents.Delete(hostnames[i])
		}

		if agentsMap.IsStale(before) {
			HostAgents.Delete(hostnames[i])
		}
	}
}
