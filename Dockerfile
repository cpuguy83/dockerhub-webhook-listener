FROM golang:1.8

ADD . /go/src/github.com/cpuguy83/dockerhub-webhook-listener

WORKDIR /go/src/github.com/cpuguy83/dockerhub-webhook-listener/hub-listener

RUN go get && go build

ENTRYPOINT ["/go/src/github.com/cpuguy83/dockerhub-webhook-listener/hub-listener/hub-listener"]
CMD ["-listen", "0.0.0.0:80"]

