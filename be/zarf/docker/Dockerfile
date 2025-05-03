ARG GO_BASE_VERSION=1.24.2
ARG GO_BASE_IMAGE=alpine

FROM golang:${GO_BASE_VERSION}-${GO_BASE_IMAGE} AS build
ARG SERVICE_NAME=sn

WORKDIR /src

COPY . .

# Build
RUN go build -o ./${SERVICE_NAME} ./cmd/main.go

# Run service
FROM alpine
ARG SERVICE_NAME=sn

COPY --from=build --chown=nobody:nogroup /src/${SERVICE_NAME} ./app/${SERVICE_NAME}

USER nobody:nogroup
WORKDIR /app

ENTRYPOINT ./sn run