package g

import (
	"github.com/Cepave/ops-common/model"
	"log"
	"strings"
)

func DesiredAgents(hostname string) (desiredAgents []*model.DesiredAgent) {
	config := Config()
	for _, inheritConfig := range config.Agents {
		defaultConfig := inheritConfig.Default
		if defaultConfig == nil {
			log.Println("default configuration is missed")
			continue
		}

		desiredAgent := &model.DesiredAgent{
			Name:    defaultConfig.Name,
			Version: defaultConfig.Version,
			Tarball: defaultConfig.Tarball,
			Md5:     defaultConfig.Md5,
			Cmd:     defaultConfig.Cmd,
		}

		others := inheritConfig.Others
		if others != nil && len(others) > 0 {
			for _, otherConfig := range inheritConfig.Others {
				if otherConfig == nil {
					continue
				}

				if !strings.HasPrefix(hostname, otherConfig.Prefix) {
					continue
				}

				if otherConfig.Version != "" {
					desiredAgent.Version = otherConfig.Version
				}

				if otherConfig.Tarball != "" {
					desiredAgent.Tarball = otherConfig.Tarball
				}

				if otherConfig.Md5 != "" {
					desiredAgent.Md5 = otherConfig.Md5
				}

				if otherConfig.Cmd != "" {
					desiredAgent.Cmd = otherConfig.Cmd
				}
			}
		}

		desiredAgents = append(desiredAgents, desiredAgent)
	}

	return
}
