package http

import (
	"encoding/base64"
	"github.com/Cepave/ops-meta/g"
	"io/ioutil"
	"net/http"
	"strings"
)

func configTarballRoutes() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		auth := strings.SplitN(r.Header["Authorization"][0], " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "bad syntax", http.StatusBadRequest)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !Validate(pair[0], pair[1]) {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		http.FileServer(http.Dir(g.Config().TarballDir)).ServeHTTP(w, r)
	})

}

func Validate(username, password string) bool {

	content, err := ioutil.ReadFile("./username")
	fUsername := strings.Trim(string(content), "\n")
	if err != nil {
		panic(err)
	}
	content, err = ioutil.ReadFile("./password")
	fPassword := strings.Trim(string(content), "\n")
	if err != nil {
		panic(err)
	}

	if username == fUsername && password == fPassword {
		return true
	}
	return false
}
