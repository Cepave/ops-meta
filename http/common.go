package http

import (
	"github.com/Cepave/ops-meta/g"
	"github.com/toolkits/file"
	"net/http"
	"strings"
)

func configCommonRoutes() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(g.VERSION))
	})

	http.HandleFunc("/workdir", func(w http.ResponseWriter, r *http.Request) {
		RenderDataJson(w, file.SelfDir())
	})

	http.HandleFunc("/config/reload", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.RemoteAddr, "127.0.0.1") {
			err := g.ParseConfig(g.ConfigFile)
			AutoRender(w, g.Config(), err)
		} else {
			w.Write([]byte("no privilege"))
		}
	})
}
