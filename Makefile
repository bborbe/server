install:
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/debug_server/debug_server.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/file_server/file_server.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/overlay_server/overlay_server.go
test:
	GO15VENDOREXPERIMENT=1 go test `glide novendor`
check:
	golint ./...
	errcheck ./...
rundebug:
	debug_server \
	-loglevel=INFO \
	-port 8080
runfileserver:
	file_server \
	-loglevel=INFO \
	-port=8080 \
	-root=/tmp \
	-auth-user=user \
	-auth-pass=pass \
	-auth-realm=login-required
runoverlayserver:
	overlay_server \
	-loglevel=INFO \
	-port=8080 \
	-root=/tmp \
	-overlays=/a,/b,/c \
	-auth-user=user \
	-auth-pass=pass \
	-auth-realm=login-required
open:
	open http://localhost:8080/
format:
	find . -name "*.go" -exec gofmt -w "{}" \;
	goimports -w=true .
prepare:
	npm install
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/Masterminds/glide
	go get -u github.com/golang/lint/golint
	go get -u github.com/kisielk/errcheck
	glide install
update:
	glide up
clean:
	rm -rf vendor
