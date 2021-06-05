FROM golang:1.16.5-alpine3.13

WORKDIR /go/src/app

COPY --chown=1000:1000 src/ .

RUN go get -d -v ./... && \
      go install -v ./...

WORKDIR /go
USER 1000:1000
CMD ["fairwinds-pod-logger"]
