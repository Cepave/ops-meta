package http

import (
	"encoding/json"
	"fmt"
	"github.com/Cepave/ops-common/model"
	"github.com/Cepave/ops-meta/g"
	"github.com/toolkits/net/httplib"
	"net/http"
	"testing"
	"time"
)

func sendHeartbeatReq(t *testing.T) {
	req := model.HeartbeatRequest{Hostname: "cnc-bj-123-123-123-123"}
	realAgents := []*model.RealAgent{}
	realAgent := &model.RealAgent{
		Name:      "installed-agent1",
		Version:   "5.1.4",
		Status:    "started",
		Timestamp: 1465874871,
	}
	realAgents = append(realAgents, realAgent)
	realAgent = &model.RealAgent{
		Name:      "installed-agent2",
		Version:   "0.9.1",
		Status:    "stoped",
		Timestamp: 1465874871,
	}
	realAgents = append(realAgents, realAgent)
	req.RealAgents = realAgents
	bs, err := json.Marshal(req)
	url := fmt.Sprintf("http://localhost:%s/heartbeat", "9002")
	httpRequest := httplib.Post(url).SetTimeout(time.Second*10, time.Minute)
	httpRequest.Body(bs)
	_, err = httpRequest.Bytes()
	if err != nil {
		fmt.Println(err)
	}
}

func start() {
	err := http.ListenAndServe(":9002", nil)
	if err != nil {
		fmt.Println("ListenAndServe failed: ", err)
	} else {
		fmt.Println("OK")
	}
}

func TestKordan(t *testing.T) {
	fmt.Println("test start")
	if err := g.ParseConfig("../cfg.example.json"); err != nil {
		fmt.Println(err)
	}
	go start()
	time.Sleep(500 * time.Millisecond)
	sendHeartbeatReq(t)
}
