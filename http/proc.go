package http

import (
	"fmt"
	"gitcafe.com/ops/meta/store"
	"net/http"
	"strings"
	"time"
)

func configProcRoutes() {

	http.HandleFunc("/status/json/", func(w http.ResponseWriter, r *http.Request) {
		agentName := r.URL.Path[len("/status/json/"):]
		if agentName == "" {
			http.Error(w, "agent name is blank", http.StatusBadRequest)
			return
		}

		data := store.HostAgents.Status(agentName)
		RenderJson(w, data)
	})

	http.HandleFunc("/status/text/", func(w http.ResponseWriter, r *http.Request) {
		agentName := r.URL.Path[len("/status/text/"):]
		if agentName == "" {
			http.Error(w, "agent name is blank", http.StatusBadRequest)
			return
		}

		data := store.HostAgents.Status(agentName)
		arr := make([]string, len(data))
		i := 0
		for hostname, ra := range data {
			if ra != nil {
				arr[i] = fmt.Sprintf(
					"%s %s %s %v %s\n",
					hostname,
					ra.Version,
					ra.Status,
					ra.Timestamp,
					time.Unix(ra.Timestamp, 0).Format("2006-01-02 15:04:05"),
				)
			} else {
				arr[i] = "no such agent"
			}

			i++
		}

		w.Write([]byte(strings.Join(arr, "")))

	})

}
