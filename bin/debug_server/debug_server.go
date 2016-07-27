package main

import (
	"fmt"
	"net/http"
	"os"

	debug_handler "github.com/bborbe/http_handler/debug"

	"runtime"

	flag "github.com/bborbe/flagenv"
	"github.com/bborbe/http_handler/static"
	"github.com/bborbe/log"
	"github.com/facebookgo/grace/gracehttp"
)

const (
	PARAMETER_LOGLEVEL = "loglevel"
)

var (
	logger      = log.DefaultLogger
	portPtr     = flag.Int("port", 8080, "Port")
	logLevelPtr = flag.String(PARAMETER_LOGLEVEL, log.DEBUG_STRING, log.FLAG_USAGE)
)

func main() {
	defer logger.Close()
	flag.Parse()

	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Debugf("set log level to %s", *logLevelPtr)

	runtime.GOMAXPROCS(runtime.NumCPU())

	server, err := createServer(
		*portPtr,
	)
	if err != nil {
		logger.Fatal(err)
		logger.Close()
		os.Exit(1)
	}
	logger.Debugf("start server")
	gracehttp.Serve(server)
}

func createServer(
	port int,
) (*http.Server, error) {
	handler := debug_handler.New(static.NewHandlerStaticContent("ok"))
	return &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: handler}, nil
}
