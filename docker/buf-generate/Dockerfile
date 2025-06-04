ARG GOLANG_VERSION=1.24.3 \
    NODE_VERSION=20 \
    BUF_VERSION=1.28.1 \
    ALPINE_VERSION=3.19

FROM golang:${GOLANG_VERSION}-alpine AS golang
# install Go plugin
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6 && \
    go install connectrpc.com/connect/cmd/protoc-gen-connect-go@v1.18.1

FROM node:${NODE_VERSION}-alpine AS node
# install TypeScript plugin for Connect-ES 2.0
RUN npm install --global \
    @bufbuild/protoc-gen-es@2.5.1 \
    @connectrpc/protoc-gen-connect-query@2.1.0

FROM bufbuild/buf:${BUF_VERSION} AS buf
FROM alpine:${ALPINE_VERSION}

COPY --from=golang /go/bin/ /usr/local/bin/
COPY --from=node /usr/local/bin/ /usr/local/bin/
COPY --from=node /usr/local/lib/node_modules/ /usr/local/lib/node_modules/
COPY --from=buf /usr/local/bin/buf /usr/local/bin/

RUN apk add --no-cache \
    libc6-compat \
    libgcc \
    libstdc++

WORKDIR /workspace
