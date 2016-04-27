package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	flag "github.com/bborbe/flagenv"
	io_util "github.com/bborbe/io/util"
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

	server, err := createServer(*portPtr, *documentRootPtr, *overlaysPtr)
	if err != nil {
		logger.Fatal(err)
		logger.Close()
		os.Exit(1)
	}
	logger.Debugf("start server")
	gracehttp.Serve(server)
}

func createServer(port int, documentRoot string, overlays string) (*http.Server, error) {
	dirs, err := toDirs(documentRoot, overlays)
	if err != nil {
		return nil, err
	}
	return &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: multi_fileserver.NewMultiFileserverHandler(dirs...)}, nil
}

func toDirs(root string, overlays string) ([]string, error) {
	root, err := io_util.NormalizePath(root)
	if err != nil {
		return nil, err
	}
	result := []string{root}
	for _, dir := range strings.Split(overlays, ",") {
		if len(dir) == 0 {
			continue
		}
		dir, err := io_util.NormalizePath(dir)
		if err != nil {
			return nil, err
		}
		result = append(result, dir)
	}
	return result, nil
}
