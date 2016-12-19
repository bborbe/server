package main

import (
	flag "github.com/bborbe/flagenv"
	debug_handler "github.com/bborbe/http_handler/debug"
	"github.com/bborbe/http_handler/dump_request"
	"github.com/bborbe/server/model"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/golang/glog"
	"net/http"
	"runtime"
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
	port := model.Port(*portPtr)
	handler := debug_handler.New(dump.New())
	glog.V(2).Infof("create http server on %s", port.Address())
	return &http.Server{Addr: port.Address(), Handler: handler}, nil
}
