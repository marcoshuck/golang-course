FROM golang:1.16 AS builder

## Download dependencies
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ../03 .

## Build application
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o app ./cmd/app

## Move binary file to /dist
WORKDIR /dist
RUN cp /build/app .
RUN mkdir /data

#####################################################################
# Runner
FROM scratch
COPY --chown=0:0 --from=builder /dist /
COPY --chown=65534:0 --from=builder /data /data
USER 65534
WORKDIR /data
EXPOSE 3000
CMD ["/app"]