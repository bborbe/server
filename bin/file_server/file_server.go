package main

import (
	"fmt"
	"net/http"
	"os"
	flag "github.com/bborbe/flagenv"
	io_util "github.com/bborbe/io/util"
	"github.com/bborbe/log"
	"github.com/facebookgo/grace/gracehttp"
)

const (
	PARAMETER_LOGLEVEL = "loglevel"
)

var (
	logger = log.DefaultLogger
	portPtr = flag.Int("port", 8080, "Port")
	documentRootPtr = flag.String("root", "", "Document root directory")
	logLevelPtr = flag.String(PARAMETER_LOGLEVEL, log.INFO_STRING, log.FLAG_USAGE)
)

func main() {
	defer logger.Close()
	flag.Parse()

	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Debugf("set log level to %s", *logLevelPtr)

	server, err := createServer(*portPtr, *documentRootPtr)
	if err != nil {
		logger.Fatal(err)
		logger.Close()
		os.Exit(1)
	}
	logger.Debugf("start server")
	gracehttp.Serve(server)
}

func createServer(port int, documentRoot string) (*http.Server, error) {
	root, err := io_util.NormalizePath(documentRoot)
	if err != nil {
		return nil, err
	}
	return &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: http.FileServer(http.Dir(root))}, nil
}
