package g

import (
	"encoding/json"
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

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}

	configLock.Lock()
	defer configLock.Unlock()

	config = &c

	log.Println("read config file:", cfg, "successfully")
}
