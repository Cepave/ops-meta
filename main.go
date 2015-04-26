package main

import (
	"flag"
	"fmt"
	"gitcafe.com/ops/meta/g"
	"gitcafe.com/ops/meta/http"
	"gitcafe.com/ops/meta/store"
	"os"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)

	go http.Start()
	go store.CleanStaleHost()

	select {}
}
