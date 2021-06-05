FROM golang:1.16.5-alpine3.13

WORKDIR /go/src/app

COPY --chown=1000:1000 src/ .

RUN go get -d -v ./... && \
      go install -v ./...

USER 1000:1000
