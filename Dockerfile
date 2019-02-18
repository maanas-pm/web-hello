FROM golang

ADD . /go/src/github.com/maanas-pm/web-hello
RUN go get github.com/spf13/viper
RUN go install github.com/maanas-pm/web-hello

EXPOSE 8082

ENTRYPOINT /go/bin/web-hello
