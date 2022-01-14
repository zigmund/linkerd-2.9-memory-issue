FROM golang:1.17.0-alpine3.14 as build_stage

WORKDIR /go/src

## Get dependencies
#ARG goproxy=https://proxy.golang.org|direct
#ENV GOPROXY=${goproxy}
#COPY go.mod go.sum ./
#RUN go mod download -x

COPY . .
RUN go build -v -o out/test-server

FROM alpine:3.14.2

COPY --from=build_stage /go/src/out/test-server /usr/local/bin/test-server
USER nobody
ENTRYPOINT ["/usr/local/bin/test-server"]
