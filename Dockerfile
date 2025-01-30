## Build
FROM golang:1.23.0 AS build

WORKDIR /app

COPY . ./

RUN make install
RUN go mod download
RUN make build

## Deploy
FROM ubuntu:22.04

WORKDIR /

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/configs /configs
COPY --from=build /app/output/server /server

EXPOSE 8082

ENTRYPOINT ["/server"]
