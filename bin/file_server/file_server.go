package main

import (
	"flag"
	"net/http"

	"github.com/bborbe/log"
	"github.com/facebookgo/grace/gracehttp"
	"strings"
	"github.com/bborbe/server/handler/multi_fileserver"
)

const (
	PARAMETER_LOGLEVEL = "loglevel"
)

var (
	logger = log.DefaultLogger
	addressPtr = flag.String("a0", ":48568", "Zero address to bind to.")
	documentRootPtr = flag.String("root", "", "Document root directory")
	overlaysPtr = flag.String("overlays", "", "Overlay directories separated by comma")
	logLevelPtr = flag.String(PARAMETER_LOGLEVEL, log.INFO_STRING, log.FLAG_USAGE)
)

func main() {
	defer logger.Close()
	flag.Parse()
	gracehttp.Serve(createServer(*addressPtr, *documentRootPtr, *overlaysPtr))
}

func createServer(address string, documentRoot string, overlays string) *http.Server {
	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Debugf("set log level to %s", *logLevelPtr)
	return &http.Server{Addr: address, Handler: multi_fileserver.NewMultiFileserverHandler(toDirs(documentRoot, overlays)...)}
}

func toDirs(root string, overlays string) []string {
	result := []string{root}
	for _, dir := range strings.Split(overlays, ",") {
		result = append(result, dir)
	}
	return result
}

