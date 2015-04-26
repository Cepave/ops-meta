package http

import (
	"encoding/json"
	"gitcafe.com/ops/common/model"
	"gitcafe.com/ops/meta/g"
	"gitcafe.com/ops/meta/store"
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
			http.Error(w, "hostname is necessary", http.StatusBadRequest)
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

		if g.Config().Debug {
			log.Println("<<<<=====Heartbeat Response")
			log.Println(resp)
		}

		RenderJson(w, resp)

	})

}
