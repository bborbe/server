package multi_fileserver

import (
	"net/http"
	"os"
	"path"

	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

type multiFileserverHandler struct {
	dirs []string
}

func NewMultiFileserverHandler(dirs ...string) *multiFileserverHandler {
	h := new(multiFileserverHandler)
	h.dirs = reverse(dirs)
	return h
}

func reverse(dirs []string) []string {
	result := make([]string, len(dirs))
	for i, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			logger.Warnf("dir %s not found", dir)
		}
		logger.Debugf("setup dir %s", dir)
		result[len(result) - i - 1] = dir
	}
	return result
}

func (h *multiFileserverHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	name := request.URL.Path
	for _, root := range h.dirs {
		logger.Debugf("search file %s in directory %s", name, root)
		file := path.Join(root, name)
		if _, err := os.Stat(file); os.IsNotExist(err) {
			logger.Debugf("file %s not found in directory %s", name, root)
		} else {
			logger.Debugf("found file %s in directory %s", name, root)
			http.ServeFile(responseWriter, request, file)
			return
		}
	}
	logger.Infof("file not found %s", name)
	http.NotFound(responseWriter, request)
}
