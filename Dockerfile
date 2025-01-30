## Build
FROM golang:1.23.0 AS build

WORKDIR /app

COPY . ./

RUN make install
RUN make generate
RUN make build

## Deploy
FROM ubuntu:22.04

WORKDIR /

COPY --from=build /app/configs /configs
COPY --from=build /app/output/server /server

EXPOSE 8082

ENTRYPOINT ["/server"]
