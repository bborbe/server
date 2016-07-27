package main

import (
	"fmt"
	debug_handler "github.com/bborbe/http_handler/debug"
	"net/http"
	"os"

	"runtime"

	flag "github.com/bborbe/flagenv"
	"github.com/bborbe/http_handler/static"
	"github.com/bborbe/log"
	"github.com/facebookgo/grace/gracehttp"
)

const (
	PARAMETER_LOGLEVEL = "loglevel"
	PARAMETER_DEBUG    = "debug"
)

var (
	logger      = log.DefaultLogger
	portPtr     = flag.Int("port", 8080, "Port")
	logLevelPtr = flag.String(PARAMETER_LOGLEVEL, log.DEBUG_STRING, log.FLAG_USAGE)
	debugPtr    = flag.Bool(PARAMETER_DEBUG, true, "debug")
)

func main() {
	defer logger.Close()
	flag.Parse()

	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Debugf("set log level to %s", *logLevelPtr)

	runtime.GOMAXPROCS(runtime.NumCPU())

	server, err := createServer(
		*portPtr,
		*debugPtr,
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
	debug bool,
) (*http.Server, error) {
	handler := static.NewHandlerStaticContent("ok")

	if debug {
		handler = debug_handler.New(handler)
	}

	return &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: handler}, nil
}
