ARG DOCKER_BUILDER_IMAGE=golang:1.25-alpine
FROM ${DOCKER_BUILDER_IMAGE} AS builder

WORKDIR /src
ADD . /src
RUN mkdir -p /out
RUN CGO_ENABLED=0 go build -mod=vendor -o=/out/example-api .

FROM scratch AS release
COPY --from=builder /out/example-api /example-api
CMD ["/example-api"]
