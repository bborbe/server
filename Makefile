all: test install
install:
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install cmd/debug-server/*.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install cmd/file-server/*.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install cmd/overlay-server/*.go
test:
	go test -cover -race $(shell go list ./... | grep -v /vendor/)
vet:
	go tool vet .
	go tool vet --shadow .
lint:
	golint -min_confidence 1 ./...
errcheck:
	errcheck -ignore '(Close|Write)' ./...
check: lint vet errcheck
rundebug:
	debug_server \
	-logtostderr \
	-v=2 \
	-port 8080
runfileserver:
	file_server \
	-logtostderr \
	-v=2 \
	-port=8080 \
	-root=/tmp \
	-auth-user=user \
	-auth-pass=pass \
	-auth-realm=login-required
runoverlayserver:
	overlay_server \
	-logtostderr \
	-v=2 \
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
	go get -u golang.org/x/lint/golint
	go get -u github.com/kisielk/errcheck
	glide install
update:
	glide up
clean:
	rm -rf vendor
