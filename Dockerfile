FROM golang:1.15 as build-env
WORKDIR /go/src/app
COPY go.mod .
RUN go mod download
COPY . /go/src/app
RUN now=$(date +'%Y-%m-%d_%T') && \
version=$(git describe --always) && \
flags="-X main.version=$version -X main.buildTime=$now" && \
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /go/bin/traffic-jam-api -ldflags "$flags"

# hadolint ignore=DL3006
FROM gcr.io/distroless/static
#FROM alpine:3.12
COPY --from=build-env /go/bin/traffic-jam-api /
COPY --from=build-env /go/src/app/static /static
ENTRYPOINT ["/traffic-jam-api"]
