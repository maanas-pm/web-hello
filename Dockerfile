FROM golang:latest

WORKDIR /go/src/github.com/maanas-pm/web-hello

ADD . .

WORKDIR /go/src/github.com/maanas-pm/web-hello

RUN go get github.com/go-chi/chi
RUN go get github.com/go-chi/render
RUN go get github.com/spf13/viper

RUN go install github.com/maanas-pm/web-hello

ENTRYPOINT /go/bin/web-hello

EXPOSE 8080

