package main

import (
	"fmt"
	"net/http"
	"runtime"

	flag "github.com/bborbe/flagenv"
	debug_handler "github.com/bborbe/http_handler/debug"
	"github.com/bborbe/http_handler/dump_request"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/golang/glog"
)

var (
	portPtr = flag.Int("port", 8080, "Port")
)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	if err := do(); err != nil {
		glog.Exit(err)
	}
}

func do() error {
	server, err := createServer()
	if err != nil {
		return err
	}
	glog.V(2).Infof("start server")
	return gracehttp.Serve(server)
}

func createServer() (*http.Server, error) {
	address := fmt.Sprintf(":%d", *portPtr)
	handler := debug_handler.New(dump.New())
	glog.V(2).Infof("create http server on %s", address)
	return &http.Server{Addr: address, Handler: handler}, nil
}
