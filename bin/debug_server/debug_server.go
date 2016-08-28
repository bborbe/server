package main

import (
	"fmt"
	"net/http"

	debug_handler "github.com/bborbe/http_handler/debug"

	"runtime"

	flag "github.com/bborbe/flagenv"
	"github.com/bborbe/http_handler/static"
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

	err := do(
		*portPtr,
	)
	if err != nil {
		glog.Exit(err)
	}
}

func do(
	port int,
) error {
	server, err := createServer(
		port,
	)
	if err != nil {
		return err
	}
	glog.V(2).Infof("start server")
	return gracehttp.Serve(server)
}

func createServer(
	port int,
) (*http.Server, error) {
	handler := debug_handler.New(static.NewHandlerStaticContent("ok"))
	return &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: handler}, nil
}
