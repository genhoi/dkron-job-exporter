# Image page: <https://hub.docker.com/_/golang>
FROM --platform=${TARGETPLATFORM:-linux/amd64} golang:1.21-alpine as builder

# app version must be passed during image building (version without any prefix).
# e.g.: `docker build --build-arg "APP_VERSION=1.2.3" .`
ARG APP_VERSION="undefined"

COPY . /src
WORKDIR /src

# arguments to pass on each go tool link invocation
ENV LDFLAGS="-s \
-X main.version=$APP_VERSION"

# compile binary file
RUN set -x
RUN go mod download
RUN go mod tidy -go 1.21
RUN CGO_ENABLED=0 go build -trimpath -ldflags "$LDFLAGS" -o ./dkron-job-exporter ./

FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine:3

RUN apk upgrade --update-cache --available && \
    apk add openssl && \
    rm -rf /var/cache/apk/*

# https://github.com/opencontainers/image-spec/blob/main/annotations.md
LABEL org.opencontainers.image.title="dkron-job-exporter"
LABEL org.opencontainers.image.description="Prometheus metrics exporter for dkron job"
LABEL org.opencontainers.image.source="https://github.com/genhoi/dkron-job-exporter"
LABEL org.opencontainers.image.vendor="genhoi"
LABEL org.opencontainers.image.licenses="MIT"

# copy required files from builder image
COPY --from=builder /src/dkron-job-exporter /usr/bin/dkron-job-exporter

# use roadrunner binary as image entrypoint
ENTRYPOINT ["/usr/bin/dkron-job-exporter"]