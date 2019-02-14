FROM golang

ADD . /go/src/github.com/maanas-pm/web-hello/src

RUN go install github.com/maanas-pm/web-hello/src

EXPOSE 8082

ENTRYPOINT /go/bin/web-hello
