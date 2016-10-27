package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Jack47/hugger"
	"github.com/golang/glog"
)

func heartbreakerHandler(w http.ResponseWriter, req *http.Request) {
	io.Write(w, hugger.Hug())
}

var LISTENING_PORT = 1024

func main() {
	http.Handle("/heartbreaker", heartbreakerHandler)
	glog.Infof("Listening on port %d", LISTENING_PORT)
	err := http.ListenAndServe(":"+str(LISTENING_PORT), nil)
	if err != nil {
		glog.Fatal("ListenAndServe: " + err.Error())
	}
}
