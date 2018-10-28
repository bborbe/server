# Collection of small servers

## Install

`go get github.com/bborbe/server/cmd/file-server`

`go get github.com/bborbe/server/cmd/overlay-server`

## Run

File Server

```
file_server \
-logtostderr \
-v=2 \
-port=8080 \
-root=/tmp \
-auth-user=user \
-auth-pass=pass \
-auth-realm=login-required
```

File Server with directory overlays

```
overlay_server \
-logtostderr \
-v=2 \
-port=8080 \
-root=/tmp \
-overlays=/a,/b,/c \ 
-auth-user=user \
-auth-pass=pass \
-auth-realm=login-required
```
