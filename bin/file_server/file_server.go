package main

import (
	"fmt"
	"net/http"
	"os"

	debug_handler "github.com/bborbe/http_handler/debug"

	"runtime"

	flag "github.com/bborbe/flagenv"
	"github.com/bborbe/http_handler/auth_basic"
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
	PARAMETER_DEBUG      = "debug"
)

var (
	logger          = log.DefaultLogger
	portPtr         = flag.Int(PARAMETER_PORT, 8080, "Port")
	documentRootPtr = flag.String(PARAMETER_ROOT, "", "Document root directory")
	logLevelPtr     = flag.String(PARAMETER_LOGLEVEL, log.INFO_STRING, log.FLAG_USAGE)
	authUserPtr     = flag.String(PARAMETER_AUTH_USER, "", "basic auth username")
	authPassPtr     = flag.String(PARAMETER_AUTH_PASS, "", "basic auth password")
	authRealmPtr    = flag.String(PARAMETER_AUTH_REALM, "", "basic auth realm")
	debugPtr        = flag.Bool(PARAMETER_DEBUG, false, "debug")
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
		*documentRootPtr,
		*authUserPtr,
		*authPassPtr,
		*authRealmPtr,
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
	documentRoot string,
	authUser string,
	authPass string,
	authRealm string,
) (*http.Server, error) {
	if port <= 0 {
		return nil, fmt.Errorf("parameter %s invalid", PARAMETER_PORT)
	}
	root, err := io_util.NormalizePath(documentRoot)
	if err != nil {
		return nil, err
	}
	var handler http.Handler = http.FileServer(http.Dir(root))
	if len(authUser) > 0 && len(authPass) > 0 && len(authRealm) > 0 {
		handler = auth_basic.New(handler.ServeHTTP, func(username string, password string) (bool, error) {
			return username == authUser && password == authPass, nil
		}, authRealm)
	}

	if debug {
		handler = debug_handler.New(handler)
	}

	return &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: handler}, nil
}
