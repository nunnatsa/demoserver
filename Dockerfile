FROM golang:1.17.1 as builder

WORKDIR /go/src/github.com/nunnatsa/demoserver

COPY go.mod main.go ./

RUN go build .

FROM centos:8
WORKDIR /app
COPY dist ./dist/
COPY --from=builder /go/src/github.com/nunnatsa/demoserver/demoserver ./
ENTRYPOINT ./demoserver
