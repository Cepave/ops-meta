package g

import (
	"encoding/json"
	"fmt"
	"github.com/toolkits/file"
	"log"
	"sync"
)

type HttpConfig struct {
	Enabled bool   `json:"enabled"`
	Listen  string `json:"listen"`
}

type AgentDefaultConfig struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Tarball string `json:"tarball"`
	Md5     string `json:"md5"`
	Cmd     string `json:"cmd"`
}

type AgentOtherConfig struct {
	Prefix  string `json:"prefix"`
	Version string `json:"version"`
	Tarball string `json:"tarball"`
	Md5     string `json:"md5"`
	Cmd     string `json:"cmd"`
}

type InheritConfig struct {
	Default *AgentDefaultConfig `json:"default"`
	Others  []*AgentOtherConfig `json:"others"`
}

type GlobalConfig struct {
	Debug      bool             `json:"debug"`
	TarballDir string           `json:"tarballDir"`
	Http       *HttpConfig      `json:"http"`
	Agents     []*InheritConfig `json:"agents"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	configLock = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func ParseConfig(cfg string) error {
	if cfg == "" {
		return fmt.Errorf("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		return fmt.Errorf("config file %s is nonexistent", cfg)
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		return fmt.Errorf("read config file %s fail %s", cfg, err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		return fmt.Errorf("parse config file %s fail %s", cfg, err)
	}

	configLock.Lock()
	defer configLock.Unlock()

	config = &c

	log.Println("read config file:", cfg, "successfully")
	return nil
}
