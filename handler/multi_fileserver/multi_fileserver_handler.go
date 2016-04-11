package multi_fileserver

import (
	"net/http"
	"strings"
	"path"
	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

type multiFileserverHandler struct {
	dirs []http.Dir
}

func NewMultiFileserverHandler(dirs... http.Dir) *multiFileserverHandler {
	h := new(multiFileserverHandler)
	h.dirs = dirs
	return h
}

func (h *multiFileserverHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	for _, root := range h.dirs {
		logger.Debugf("search in dir: %v", root)
		upath := request.URL.Path
		if !strings.HasPrefix(upath, "/") {
			upath = "/" + upath
		}
		_, err := root.Open(path.Clean(upath))
		if err != nil {
			logger.Debugf("open %s failed", path.Clean(upath))
		} else {
			fileServer := http.FileServer(root)
			fileServer.ServeHTTP(responseWriter, request)
			return
		}
	}
	http.NotFound(responseWriter, request)
}
