## Build
FROM golang:1.23.0 AS build

WORKDIR /app

COPY . ./

RUN make install
RUN make generate
RUN make build

## Deploy
FROM ubuntu:22.04

WORKDIR /pet

COPY --from=build /app/configs /pet/configs
COPY --from=build /app/output/server /pet/server

EXPOSE 8082

ENTRYPOINT ["/pet/server"]
