# Server

Package provide core server modules in Go

## Install

`go get github.com/bborbe/server/bin/file_server`

`go get github.com/bborbe/server/bin/overlay_server`

## Run

Debug Server always return ok  

```
debug_server \
-logtostderr \
-v=2 \
-port 8080
```

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

## Continuous integration

[Jenkins](https://jenkins.benjamin-borbe.de/job/Go-Server/)

## Copyright and license

    Copyright (c) 2016, Benjamin Borbe <bborbe@rocketnews.de>
    All rights reserved.
    
    Redistribution and use in source and binary forms, with or without
    modification, are permitted provided that the following conditions are
    met:
    
       * Redistributions of source code must retain the above copyright
         notice, this list of conditions and the following disclaimer.
       * Redistributions in binary form must reproduce the above
         copyright notice, this list of conditions and the following
         disclaimer in the documentation and/or other materials provided
         with the distribution.

    THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
    "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
    LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
    A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
    OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
    SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
    LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
    DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
    THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
    (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
    OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
