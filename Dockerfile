FROM golang:1.3
ADD . /hub-hook
WORKDIR /hub-hook/cli
RUN go get && go build cli.go
ENTRYPOINT ["/hub-hook/cli/cli"]
CMD [":80"]
