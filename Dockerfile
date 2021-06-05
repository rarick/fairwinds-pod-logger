ARG IMAGE_NAME=golang
ARG IMAGE_TAG=1.16.5-alpine3.13

################################
FROM $IMAGE_NAME:$IMAGE_TAG as builder
################

WORKDIR /go/src/fairwinds-pod-logger

COPY --chown=1000:1000 src/ .

RUN go get -d -v ./...
RUN go install -v ./...

################################
FROM $IMAGE_IMAGE:$IMAGE_TAG as final
################

WORKDIR /go

COPY --chown=1000:1000 --from=builder /go/bin /go/bin
COPY --chown=1000:1000 --from=builder /go/pkg /go/pkg

USER 1000:1000
CMD ["fairwinds-pod-logger"]
