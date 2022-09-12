FROM golang:1.18-bullseye

RUN go install github.com/beego/bee/v2@latest

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

ENV APP_HOME /go/src/mathapp
RUN mkdir -p "$APP_HOME"

COPY ./src "$APP_HOME"
WORKDIR "$APP_HOME"
EXPOSE 8020
CMD ["bee", "run"]