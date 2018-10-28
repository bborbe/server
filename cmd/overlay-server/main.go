package main

import (
	"fmt"
	flag "github.com/bborbe/flagenv"
	"github.com/bborbe/http_handler/auth_basic"
	debug_handler "github.com/bborbe/http_handler/debug"
	"github.com/bborbe/http_handler/multi_fileserver"
	io_util "github.com/bborbe/io/util"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/golang/glog"
	"net/http"
	"runtime"
	"strings"
)

var (
	portPtr         = flag.Int("port", 8080, "Port")
	documentRootPtr = flag.String("root", "", "Document root directory")
	overlaysPtr     = flag.String("overlays", "", "Overlay directories separated by comma")
	authUserPtr     = flag.String("auth-user", "", "basic auth username")
	authPassPtr     = flag.String("auth-pass", "", "basic auth password")
	authRealmPtr    = flag.String("auth-realm", "", "basic auth realm")
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
	documentRoot := *documentRootPtr
	overlays := *overlaysPtr
	authUser := *authUserPtr
	authPass := *authPassPtr
	authRealm := *authRealmPtr

	if *portPtr <= 0 {
		return nil, fmt.Errorf("parameter %s invalid", "port")
	}
	dirs, err := toDirs(documentRoot, overlays)
	if err != nil {
		return nil, err
	}
	var handler http.Handler = multi_fileserver.New(dirs...)
	if len(authUser) > 0 && len(authPass) > 0 && len(authRealm) > 0 {
		handler = auth_basic.New(handler.ServeHTTP, func(username string, password string) (bool, error) {
			return username == authUser && password == authPass, nil
		}, authRealm)
	}

	if glog.V(4) {
		handler = debug_handler.New(handler)
	}

	address := fmt.Sprintf(":%d", *portPtr)
	glog.V(2).Infof("create http server on %s", address)
	return &http.Server{Addr: address, Handler: handler}, nil
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
