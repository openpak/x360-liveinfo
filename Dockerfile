# Multi-stage, cross-compiled: the build stage runs on the runner's native arch and Go
# emits the target arch, so linux/arm64 builds need no QEMU. x360-common comes from vendor/,
# so the build needs nothing outside this repository.
FROM --platform=$BUILDPLATFORM golang:1.26-alpine AS build
ARG TARGETOS TARGETARCH
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -mod=vendor -trimpath -ldflags="-s -w" -o /out/x360-liveinfo ./cmd/x360-liveinfo

FROM alpine:3.22
# The firmware dials TCP 25 (privileged): the binary carries the bind capability so the container stays unprivileged.
RUN addgroup -S app && adduser -S -G app app && apk add --no-cache libcap-utils
COPY --from=build /out/x360-liveinfo /usr/local/bin/x360-liveinfo
RUN setcap cap_net_bind_service=+ep /usr/local/bin/x360-liveinfo && apk del libcap-utils
USER app
EXPOSE 25/tcp
ENTRYPOINT ["/usr/local/bin/x360-liveinfo"]
