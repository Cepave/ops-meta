package http

import (
	"encoding/json"
	"gitcafe.com/ops/common/model"
	"gitcafe.com/ops/meta/store"
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

		store.HandleHeartbeatRequest(&req)

		// write desired state
		w.Write([]byte(""))

	})

}
