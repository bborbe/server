package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"runtime"

	flag "github.com/bborbe/flagenv"
	"github.com/bborbe/http_handler/auth_basic"
	"github.com/bborbe/http_handler/multi_fileserver"
	io_util "github.com/bborbe/io/util"
	"github.com/bborbe/log"
	"github.com/facebookgo/grace/gracehttp"
)

const (
	PARAMETER_ROOT       = "root"
	PARAMETER_PORT       = "port"
	PARAMETER_LOGLEVEL   = "loglevel"
	PARAMETER_AUTH_USER  = "auth-user"
	PARAMETER_AUTH_PASS  = "auth-pass"
	PARAMETER_AUTH_REALM = "auth-realm"
	PARAMETER_OVERLAYS   = "overlays"
)

var (
	logger          = log.DefaultLogger
	portPtr         = flag.Int(PARAMETER_PORT, 8080, "Port")
	documentRootPtr = flag.String(PARAMETER_ROOT, "", "Document root directory")
	overlaysPtr     = flag.String(PARAMETER_OVERLAYS, "", "Overlay directories separated by comma")
	logLevelPtr     = flag.String(PARAMETER_LOGLEVEL, log.INFO_STRING, log.FLAG_USAGE)
	authUserPtr     = flag.String(PARAMETER_AUTH_USER, "", "basic auth username")
	authPassPtr     = flag.String(PARAMETER_AUTH_PASS, "", "basic auth password")
	authRealmPtr    = flag.String(PARAMETER_AUTH_REALM, "", "basic auth realm")
)

func main() {
	defer logger.Close()
	flag.Parse()

	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Debugf("set log level to %s", *logLevelPtr)

	runtime.GOMAXPROCS(runtime.NumCPU())

	server, err := createServer(*portPtr, *documentRootPtr, *overlaysPtr, *authUserPtr, *authPassPtr, *authRealmPtr)
	if err != nil {
		logger.Fatal(err)
		logger.Close()
		os.Exit(1)
	}
	logger.Debugf("start server")
	gracehttp.Serve(server)
}

func createServer(port int, documentRoot string, overlays string, authUser string, authPass string, authRealm string) (*http.Server, error) {
	if port <= 0 {
		return nil, fmt.Errorf("parameter %s invalid", PARAMETER_PORT)
	}
	dirs, err := toDirs(documentRoot, overlays)
	if err != nil {
		return nil, err
	}
	var handler http.Handler = multi_fileserver.NewMultiFileserverHandler(dirs...)
	if len(authUser) > 0 && len(authPass) > 0 && len(authRealm) > 0 {
		handler = auth_basic.New(handler.ServeHTTP, func(username string, password string) (bool, error) {
			return username == authUser && password == authPass, nil
		}, authRealm)
	}
	return &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: handler}, nil
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
