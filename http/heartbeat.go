package http

import (
	"encoding/json"
	"github.com/Cepave/ops-common/model"
	"github.com/Cepave/ops-meta/g"
	"github.com/Cepave/ops-meta/store"
	"log"
	"net/http"
)

func configHeartbeatRoutes() {
	// post
	http.HandleFunc("/heartbeat", func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength == 0 {
			http.Error(w, "body is blank", http.StatusBadRequest)
			return
		}

		var req model.HeartbeatRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			http.Error(w, "body format error", http.StatusBadRequest)
			return
		}

		if req.Hostname == "" {
			http.Error(w, "hostname is blank", http.StatusBadRequest)
			return
		}

		if g.Config().Debug {
			log.Println("Heartbeat Request=====>>>>")
			log.Println(req)
		}

		store.ParseHeartbeatRequest(&req)

		resp := model.HeartbeatResponse{
			ErrorMessage:  "",
			DesiredAgents: g.DesiredAgents(req.Hostname),
		}

		// TODO: This is a workaround to ensure that one updater is resposible for only one agent,
		// so NQM agent(resp.DesiredAgents[1]) would be stopped if the request comes from owl-agent-updater.
		if len(resp.DesiredAgents) > 1 {
			resp.DesiredAgents[1].Cmd = "stop"
		}

		if g.Config().Debug {
			log.Println("<<<<=====Heartbeat Response")
			log.Println(resp)
		}

		RenderJson(w, resp)

	})

}
