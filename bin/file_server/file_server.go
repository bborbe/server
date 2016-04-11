package main

import (
	"fmt"
	"net/http"
	"strings"

	flag "github.com/bborbe/flagenv"
	"github.com/bborbe/log"
	"github.com/bborbe/server/handler/multi_fileserver"
	"github.com/facebookgo/grace/gracehttp"
)

const (
	PARAMETER_LOGLEVEL = "loglevel"
)

var (
	logger          = log.DefaultLogger
	portPtr         = flag.Int("port", 8080, "Port")
	documentRootPtr = flag.String("root", "", "Document root directory")
	overlaysPtr     = flag.String("overlays", "", "Overlay directories separated by comma")
	logLevelPtr     = flag.String(PARAMETER_LOGLEVEL, log.INFO_STRING, log.FLAG_USAGE)
)

func main() {
	defer logger.Close()
	flag.Parse()

	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Debugf("set log level to %s", *logLevelPtr)

	gracehttp.Serve(createServer(*portPtr, *documentRootPtr, *overlaysPtr))
}

func createServer(port int, documentRoot string, overlays string) *http.Server {
	return &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: multi_fileserver.NewMultiFileserverHandler(toDirs(documentRoot, overlays)...)}
}

func toDirs(root string, overlays string) []string {
	result := []string{root}
	for _, dir := range strings.Split(overlays, ",") {
		result = append(result, dir)
	}
	return result
}
